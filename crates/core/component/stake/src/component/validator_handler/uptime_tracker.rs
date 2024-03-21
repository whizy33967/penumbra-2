use {
    super::{ValidatorDataRead, ValidatorDataWrite, ValidatorManager},
    crate::{
        component::{metrics, stake::ConsensusIndexRead, StateReadExt as _},
        validator, IdentityKey, Uptime,
    },
    anyhow::Result,
    async_trait::async_trait,
    cnidarium::StateWrite,
    futures::StreamExt as _,
    penumbra_sct::component::clock::EpochRead,
    sha2::{Digest as _, Sha256},
    std::collections::BTreeMap,
    tendermint::abci::types::CommitInfo,
    tokio::task::{AbortHandle, JoinSet},
    tracing::instrument,
};

/// A bundle of information about a validator used to track its uptime.
type ValidatorInformation = (IdentityKey, tendermint::PublicKey, Uptime);

/// A collection of tasks retrieving [`ValidatorInformation`].
type Lookups = JoinSet<anyhow::Result<Option<ValidatorInformation>>>;

/// Tracks validator uptimes.
///
/// Use [`track_uptime()`] to process a block's [`CommitInfo`] and update validator uptime
/// bookkeeping.
///
/// [`track_uptime()`]: Self::track_uptime
#[async_trait]
pub trait ValidatorUptimeTracker: StateWrite {
    #[instrument(skip(self, last_commit_info))]
    async fn track_uptime(&mut self, last_commit_info: &CommitInfo) -> Result<()> {
        // Note: this probably isn't the correct height for the LastCommitInfo,
        // which is about the *last* commit, but at least it'll be consistent,
        // which is all we need to count signatures.
        let height = self.get_block_height().await?;
        let params = self.get_stake_params().await?;

        // Build a mapping from addresses (20-byte truncated SHA256(pubkey)) to vote statuses.
        let did_address_vote = last_commit_info
            .votes
            .iter()
            .map(|vote| (vote.validator.address, vote.sig_info.is_signed()))
            .collect::<BTreeMap<[u8; 20], bool>>();

        // Since we don't have a lookup from "addresses" to identity keys,
        // iterate over our app's validators, and match them up with the vote data.
        // We can fetch all the data required for processing each validator concurrently:
        let mut lookups = Lookups::new();
        let mut validator_identity_stream = self.consensus_set_stream()?;
        while let Some(identity_key) = validator_identity_stream.next().await.transpose()? {
            self.spawn_validator_lookup_fut(identity_key, &mut lookups);
        }

        // Now process the data we fetched concurrently.
        // Note that this will process validator uptime changes in a random order, but because they are all
        // independent, this doesn't introduce any nondeterminism into the complete state change.
        while let Some(data) = lookups.join_next().await.transpose()? {
            if let Some((identity_key, consensus_key, mut uptime)) = data? {
                // for some reason last_commit_info has truncated sha256 hashes
                let addr: [u8; 20] =
                    Sha256::digest(&consensus_key.to_bytes()).as_slice()[0..20].try_into()?;

                let voted = did_address_vote
                    .get(&addr)
                    .cloned()
                    // If the height is `1`, then the `LastCommitInfo` refers to the genesis block,
                    // which has no signers -- so we'll mark all validators as having signed.
                    // https://github.com/penumbra-zone/penumbra/issues/1050
                    .unwrap_or(height == 1);

                tracing::debug!(
                    ?voted,
                    num_missed_blocks = ?uptime.num_missed_blocks(),
                    ?identity_key,
                    ?params.missed_blocks_maximum,
                    "recorded vote info"
                );
                metrics::gauge!(metrics::MISSED_BLOCKS, "identity_key" => identity_key.to_string())
                    .increment(uptime.num_missed_blocks() as f64);

                uptime.mark_height_as_signed(height, voted)?;
                if uptime.num_missed_blocks() as u64 >= params.missed_blocks_maximum {
                    self.set_validator_state(&identity_key, validator::State::Jailed)
                        .await?;
                } else {
                    self.set_validator_uptime(&identity_key, uptime);
                }
            }
        }

        Ok(())
    }

    /// Spawns a future that will retrieve validator information.
    ///
    /// NB: This function is synchronous, but the lookup will run asynchronously as part of the
    /// provided [`JoinSet`]. This permits us to fetch information about all of the validators
    /// in the consensus set in parallel.
    ///
    /// # Panics
    ///
    /// This will panic if there is no recorded state for a validator with the given
    /// [`IdentityKey`].
    fn spawn_validator_lookup_fut(
        &self,
        identity_key: crate::IdentityKey,
        lookups: &mut Lookups,
    ) -> AbortHandle {
        let state = self.get_validator_state(&identity_key);
        let uptime = self.get_validator_uptime(&identity_key);
        let consensus_key = self.fetch_validator_consensus_key(&identity_key);

        lookups.spawn(async move {
            let state = state
                .await?
                .expect("every known validator must have a recorded state");

            match state {
                validator::State::Active => {
                    // If the validator is active, we need its consensus key and current uptime data:
                    Ok(Some((
                        identity_key,
                        consensus_key
                            .await?
                            .expect("every known validator must have a recorded consensus key"),
                        uptime
                            .await?
                            .expect("every known validator must have a recorded uptime"),
                    )))
                }
                _ => {
                    // Otherwise, we don't need to track its uptime, and there's no data to fetch.
                    anyhow::Ok(None)
                }
            }
        })
    }
}

impl<T: StateWrite + ?Sized> ValidatorUptimeTracker for T {}
