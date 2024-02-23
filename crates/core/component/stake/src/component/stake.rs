use crate::params::StakeParameters;
use crate::rate::BaseRateData;
use crate::validator::{self, Validator};
use crate::{
    state_key, CurrentConsensusKeys, Delegate, DelegationChanges, DelegationToken, FundingStreams,
    IdentityKey, Penalty, Undelegate,
};
use anyhow::Context;
use anyhow::{anyhow, Result};
use async_trait::async_trait;
use cnidarium::{StateRead, StateWrite};
use cnidarium_component::Component;
use futures::{StreamExt, TryStreamExt};
use penumbra_num::Amount;
use penumbra_proto::{StateReadProto, StateWriteProto};
use penumbra_sct::component::clock::EpochRead;
use penumbra_shielded_pool::component::SupplyRead;
use sha2::{Digest, Sha256};
use std::pin::Pin;
use std::str::FromStr;
use std::{collections::BTreeMap, sync::Arc};
use tendermint::v0_37::abci;
use tendermint::validator::Update;
use tendermint::{block, PublicKey};
use tracing::instrument;

use crate::component::epoch_handler::EpochHandler;
use crate::component::validator_handler::{ValidatorDataRead, ValidatorManager};

pub struct Staking {}

impl Staking {}

#[async_trait]
impl Component for Staking {
    type AppState = (
        crate::genesis::Content,
        penumbra_shielded_pool::genesis::Content,
    );

    #[instrument(name = "staking", skip(state, app_state))]
    async fn init_chain<S: StateWrite>(mut state: S, app_state: Option<&Self::AppState>) {
        match app_state {
            None => { /* perform upgrade specific check */ }
            Some((staking_genesis, sp_genesis)) => {
                state.put_stake_params(staking_genesis.stake_params.clone());

                let starting_height = state
                    .get_block_height()
                    .await
                    .expect("should be able to get initial block height");
                let starting_epoch = state
                    .get_epoch_by_height(starting_height)
                    .await
                    .expect("should be able to get initial epoch");
                let epoch_index = starting_epoch.index;

                let genesis_base_rate = BaseRateData {
                    epoch_index,
                    base_reward_rate: 0u128.into(),
                    base_exchange_rate: 1_0000_0000u128.into(),
                };
                state.set_base_rate(genesis_base_rate.clone());

                let mut genesis_allocations = BTreeMap::<_, Amount>::new();
                for allocation in &sp_genesis.allocations {
                    let value = allocation.value();
                    *genesis_allocations.entry(value.asset_id).or_default() += value.amount;
                }

                for validator in &staking_genesis.validators {
                    // Parse the proto into a domain type.
                    let validator = Validator::try_from(validator.clone())
                        .expect("should be able to parse genesis validator");

                    state
                        .add_genesis_validator(&genesis_allocations, &genesis_base_rate, validator)
                        .await
                        .expect("should be able to add genesis validator to state");
                }

                // First, "prime" the state with an empty set, so the build_ function can read it.
                state.put(
                    state_key::consensus_update::consensus_keys().to_owned(),
                    CurrentConsensusKeys::default(),
                );

                // Finally, record that there were no delegations in this block, so the data
                // isn't missing when we process the first epoch transition.
                state
                    .set_delegation_changes(
                        starting_height
                            .try_into()
                            .expect("should be able to convert u64 into block height"),
                        Default::default(),
                    )
                    .await;
            }
        }
        // Build the initial validator set update.
        state
            .build_cometbft_validator_updates()
            .await
            .expect("should be able to build initial tendermint validator updates");
    }

    #[instrument(name = "staking", skip(state, begin_block))]
    async fn begin_block<S: StateWrite + 'static>(
        state: &mut Arc<S>,
        begin_block: &abci::request::BeginBlock,
    ) {
        let state = Arc::get_mut(state).expect("state should be unique");
        // For each validator identified as byzantine by tendermint, update its
        // state to be slashed. If the validator is not tracked in the JMT, this
        // will be a no-op. See #2919 for more details.
        for evidence in begin_block.byzantine_validators.iter() {
            let _ = state.process_evidence(evidence).await.map_err(|e| {
                tracing::warn!(?e, "failed to process byzantine misbehavior evidence")
            });
        }

        state
            .track_uptime(&begin_block.last_commit_info)
            .await
            .expect("should be able to track uptime");
    }

    #[instrument(name = "staking", skip(state, end_block))]
    async fn end_block<S: StateWrite + 'static>(
        state: &mut Arc<S>,
        end_block: &abci::request::EndBlock,
    ) {
        let state = Arc::get_mut(state).expect("state should be unique");
        // Write the delegation changes for this block.
        state
            .set_delegation_changes(
                end_block
                    .height
                    .try_into()
                    .expect("should be able to convert i64 into block height"),
                state.get_delegation_changes_tally().clone(),
            )
            .await;
    }

    #[instrument(name = "staking", skip(state))]
    async fn end_epoch<S: StateWrite + 'static>(state: &mut Arc<S>) -> anyhow::Result<()> {
        let state = Arc::get_mut(state).context("state should be unique")?;
        let epoch_ending = state
            .get_current_epoch()
            .await
            .context("should be able to get current epoch during end_epoch")?;
        state.end_epoch(epoch_ending).await?;
        // Since we only update the validator set at epoch boundaries,
        // we only need to build the validator set updates here in end_epoch.
        state
            .build_cometbft_validator_updates()
            .await
            .context("should be able to build tendermint validator updates")?;
        Ok(())
    }
}

pub trait ConsensusUpdateRead: StateRead {
    /// Returns a list of validator updates to send to Tendermint.
    ///
    /// Set during `end_block`.
    fn cometbft_validator_updates(&self) -> Option<Vec<Update>> {
        self.object_get(state_key::internal::cometbft_validator_updates())
            .unwrap_or(None)
    }
}

impl<T: StateRead + ?Sized> ConsensusUpdateRead for T {}

pub(crate) trait ConsensusUpdateWrite: StateWrite {
    fn put_cometbft_validator_updates(&mut self, updates: Vec<Update>) {
        tracing::debug!(?updates);
        self.object_put(
            state_key::internal::cometbft_validator_updates(),
            Some(updates),
        )
    }
}

impl<T: StateWrite + ?Sized> ConsensusUpdateWrite for T {}

/// Extension trait providing read access to staking data.
#[async_trait]
pub trait StateReadExt: StateRead {
    /// Gets the stake parameters from the JMT.
    async fn get_stake_params(&self) -> Result<StakeParameters> {
        self.get(state_key::parameters::key())
            .await
            .expect("no deserialization error should happen")
            .ok_or_else(|| anyhow!("Missing StakeParameters"))
    }

    /// Indicates if the stake parameters have been updated in this block.
    fn stake_params_updated(&self) -> bool {
        self.object_get::<()>(state_key::parameters::updated_flag())
            .is_some()
    }

    async fn signed_blocks_window_len(&self) -> Result<u64> {
        Ok(self.get_stake_params().await?.signed_blocks_window_len)
    }

    async fn missed_blocks_maximum(&self) -> Result<u64> {
        Ok(self.get_stake_params().await?.missed_blocks_maximum)
    }

    /// Delegation changes accumulated over the course of this block, to be
    /// persisted at the end of the block for processing at the end of the next
    /// epoch.
    fn get_delegation_changes_tally(&self) -> DelegationChanges {
        self.object_get(state_key::chain::delegation_changes::key())
            .unwrap_or_default()
    }

    async fn get_current_base_rate(&self) -> Result<BaseRateData> {
        self.get(state_key::chain::base_rate::current())
            .await
            .map(|rate_data| rate_data.expect("rate data must be set after init_chain"))
    }

    fn get_previous_base_rate(&self) -> Option<BaseRateData> {
        self.object_get(state_key::chain::base_rate::previous())
    }

    /// Returns the funding queue from object storage (end-epoch).
    fn get_funding_queue(&self) -> Option<Vec<(IdentityKey, FundingStreams, Amount)>> {
        self.object_get(state_key::validators::rewards::staking())
    }

    async fn get_delegation_changes(&self, height: block::Height) -> Result<DelegationChanges> {
        Ok(self
            .get(&state_key::chain::delegation_changes::by_height(
                height.value(),
            ))
            .await?
            .ok_or_else(|| anyhow!("missing delegation changes for block {}", height))?)
    }
}

impl<T: StateRead + ?Sized> StateReadExt for T {}

/// Extension trait providing write access to staking data.
#[async_trait]
pub trait StateWriteExt: StateWrite {
    /// Writes the provided stake parameters to the JMT.
    fn put_stake_params(&mut self, params: StakeParameters) {
        // Note that the stake params have been updated:
        self.object_put(state_key::parameters::updated_flag(), ());

        // Change the stake parameters:
        self.put(state_key::parameters::key().into(), params)
    }

    /// Delegation changes accumulated over the course of this block, to be
    /// persisted at the end of the block for processing at the end of the next
    /// epoch.
    fn put_delegation_changes(&mut self, delegation_changes: DelegationChanges) {
        self.object_put(
            state_key::chain::delegation_changes::key(),
            delegation_changes,
        )
    }

    /// Push an entry in the delegation queue for the current block (object-storage).
    fn push_delegation(&mut self, delegation: Delegate) {
        let mut changes = self.get_delegation_changes_tally();
        changes.delegations.push(delegation);
        self.put_delegation_changes(changes);
    }

    /// Push an entry in the undelegation queue for the current block (object-storage).
    fn push_undelegation(&mut self, undelegation: Undelegate) {
        let mut changes = self.get_delegation_changes_tally();
        changes.undelegations.push(undelegation);
        self.put_delegation_changes(changes);
    }

    #[instrument(skip(self))]
    fn queue_staking_rewards(
        &mut self,
        staking_reward_queue: Vec<(IdentityKey, FundingStreams, Amount)>,
    ) {
        self.object_put(
            state_key::validators::rewards::staking(),
            staking_reward_queue,
        )
    }

    async fn register_consensus_key(
        &mut self,
        identity_key: &IdentityKey,
        consensus_key: &PublicKey,
    ) {
        /// Translates from consensus keys to the truncated sha256 hashes in last_commit_info
        /// This should really be a refined type upstream, but we can't currently upstream
        /// to tendermint-rs, for process reasons, and shouldn't do our own tendermint data
        /// modeling, so this is an interim hack.
        fn validator_address(ck: &PublicKey) -> [u8; 20] {
            let ck_bytes = ck.to_bytes();
            let addr: [u8; 20] = Sha256::digest(ck_bytes).as_slice()[0..20]
                .try_into()
                .expect("Sha256 digest should be 20-bytes long");

            addr
        }

        let address = validator_address(consensus_key);
        tracing::debug!(?identity_key, ?consensus_key, hash = ?hex::encode(address), "registering consensus key");
        self.put(
            state_key::validators::lookup_by::cometbft_address(&address),
            consensus_key.clone(),
        );
        self.put(
            state_key::validators::lookup_by::consensus_key(consensus_key),
            identity_key.clone(),
        );
    }
}

impl<T: StateWrite + ?Sized> StateWriteExt for T {}

#[async_trait]
pub trait SlashingData: StateRead {
    async fn get_penalty_in_epoch(&self, id: &IdentityKey, epoch_index: u64) -> Option<Penalty> {
        self.get(&state_key::penalty::for_id_in_epoch(id, epoch_index))
            .await
            .expect("serialization error cannot happen")
    }

    async fn get_penalty_for_range(&self, id: &IdentityKey, start: u64, end: u64) -> Vec<Penalty> {
        let prefix = state_key::penalty::prefix(id);
        let all_penalties: BTreeMap<String, Penalty> = self
            .prefix::<Penalty>(&prefix)
            .try_collect()
            .await
            .unwrap_or_default();
        let start_key = state_key::penalty::for_id_in_epoch(id, start);
        let end_key = state_key::penalty::for_id_in_epoch(id, end);
        all_penalties
            .range(start_key..end_key)
            .map(|(_k, v)| v.to_owned())
            .collect()
    }

    fn compute_compounded_penalty(penalties: Vec<Penalty>) -> Penalty {
        let compounded = Penalty::from_percent(0);
        penalties
            .into_iter()
            .fold(compounded, |acc, penalty| acc.compound(penalty))
    }

    /// Returns the compounded penalty for the given validator over the half-open range of epochs [start, end).
    async fn compounded_penalty_over_range(
        &self,
        id: &IdentityKey,
        start: u64,
        end: u64,
    ) -> Result<Penalty> {
        let range = self.get_penalty_for_range(id, start, end).await;
        let compounded_penalty = Self::compute_compounded_penalty(range);
        Ok(compounded_penalty)
    }
}

impl<T: StateRead + ?Sized> SlashingData for T {}

#[async_trait]
pub(crate) trait InternalStakingData: StateRead {
    /// Calculate the amount of stake that is delegated to the currently active validator set,
    /// denominated in the staking token.
    #[instrument(skip(self))]
    async fn total_active_stake(&self) -> Result<Amount> {
        let mut total_active_stake = Amount::zero();

        let mut validator_stream = self.consensus_set_stream()?;
        while let Some(validator_identity) = validator_stream.next().await {
            let validator_identity = validator_identity?;
            let validator_state = self
                .get_validator_state(&validator_identity)
                .await?
                .ok_or_else(|| {
                    anyhow::anyhow!("validator (identity_key={}) is in the consensus set index but its state was not found", validator_identity)
                })?;

            if validator_state != validator::State::Active {
                continue;
            }

            let delegation_token_supply = self
                .token_supply(&DelegationToken::from(validator_identity).id())
                .await?
                .expect("delegation token should be known");

            let validator_rate = self
                .get_validator_rate(&validator_identity)
                .await?
                .ok_or_else(|| {
                    anyhow::anyhow!("validator (identity_key={}) is in the consensus set index but its rate data was not found", validator_identity)
                })?;

            // Add the validator's unbonded amount to the total active stake
            total_active_stake = total_active_stake
                .checked_add(&validator_rate.unbonded_amount(delegation_token_supply))
                .ok_or_else(|| {
                    anyhow::anyhow!("total active stake overflowed `Amount` (128 bits)")
                })?;
        }

        Ok(total_active_stake)
    }
}

impl<T: StateRead + ?Sized> InternalStakingData for T {}

#[async_trait]
pub trait RateDataWrite: StateWrite {
    #[instrument(skip(self))]
    fn set_base_rate(&mut self, rate_data: BaseRateData) {
        tracing::debug!("setting base rate");
        self.put(state_key::chain::base_rate::current().to_owned(), rate_data);
    }

    #[instrument(skip(self))]
    fn set_prev_base_rate(&mut self, rate_data: BaseRateData) {
        self.object_put(state_key::chain::base_rate::previous(), rate_data);
    }

    async fn record_slashing_penalty(
        &mut self,
        identity_key: &IdentityKey,
        slashing_penalty: Penalty,
    ) {
        let current_epoch_index = self
            .get_current_epoch()
            .await
            .expect("epoch has been set")
            .index;

        let current_penalty = self
            .get_penalty_in_epoch(identity_key, current_epoch_index)
            .await
            .unwrap_or(Penalty::from_percent(0));

        let new_penalty = current_penalty.compound(slashing_penalty);

        self.put(
            state_key::penalty::for_id_in_epoch(identity_key, current_epoch_index),
            new_penalty,
        );
    }

    async fn set_delegation_changes(&mut self, height: block::Height, changes: DelegationChanges) {
        self.put(
            state_key::chain::delegation_changes::by_height(height.value()),
            changes,
        );
    }
}

impl<T: StateWrite + ?Sized> RateDataWrite for T {}

#[async_trait]
pub trait ConsensusIndexRead: StateRead {
    /// Returns a stream of [`IdentityKey`]s of validators that are currently in the consensus set.
    /// This only excludes validators that do not meet the minimum validator stake requirement
    /// (see [`StakeParameters::min_validator_stake`]).
    fn consensus_set_stream(
        &self,
    ) -> Result<Pin<Box<dyn futures::Stream<Item = Result<IdentityKey>> + Send + 'static>>> {
        Ok(self
            .nonverifiable_prefix_raw(
                state_key::validators::consensus_set_index::prefix().as_bytes(),
            )
            .map(|res| {
                res.map(|(_, raw_identity_key)| {
                    // TODO(erwan): is this an opportunity to extend the proto overlay?
                    let str_identity_key = std::str::from_utf8(raw_identity_key.as_slice())
                        .expect("state keys should only have valid identity keys");
                    IdentityKey::from_str(str_identity_key)
                        .expect("state keys should only have valid identity keys")
                })
            })
            .boxed())
    }

    /// Returns whether the given validator should be indexed in the consensus set.
    async fn belongs_in_index(&self, validator_id: &IdentityKey) -> bool {
        let Some(state) = self
            .get_validator_state(validator_id)
            .await
            .expect("no deserialization error")
        else {
            tracing::error!(
                %validator_id,
                "consensus set index check returned false because validator state was not found");
            return false;
        };

        let belongs_in_index =
            matches!(state, validator::State::Active | validator::State::Inactive);

        tracing::debug!(
            %validator_id,
            ?belongs_in_index,
            ?state,
            "does not belong in consensus set"
        );

        belongs_in_index
    }
}

impl<T: StateRead + ?Sized> ConsensusIndexRead for T {}

#[async_trait]
pub trait ConsensusIndexWrite: StateWrite {
    /// Add a validator identity to the consensus set index.
    /// The consensus set index includes any validator that has a delegation pool that
    /// is greater than [`StakeParameters::min_validator_stake`].
    fn add_consensus_set_index(&mut self, identity_key: &IdentityKey) {
        tracing::debug!(validator = %identity_key, "adding validator identity to consensus set index");
        self.nonverifiable_put_raw(
            state_key::validators::consensus_set_index::by_id(identity_key)
                .as_bytes()
                .to_vec(),
            identity_key.to_string().as_bytes().to_vec(),
        );
    }

    /// Remove a validator identity from the consensus set index.
    /// The consensus set index includes any validator that has a delegation pool that
    /// is greater than [`StakeParameters::min_validator_stake`].
    fn remove_consensus_set_index(&mut self, identity_key: &IdentityKey) {
        tracing::debug!(validator = %identity_key, "removing validator identity from consensus set index");
        self.nonverifiable_delete(
            state_key::validators::consensus_set_index::by_id(identity_key)
                .as_bytes()
                .to_vec(),
        );
    }
}

impl<T: StateWrite + ?Sized> ConsensusIndexWrite for T {}
