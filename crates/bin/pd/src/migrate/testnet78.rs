//! Contains functions related to the migration script of Testnet77.
//! The Testnet 77 release included several consensus-breaking changes,
//! but no state-breaking changes, so the migration is essentially a no-op,
//! other than resetting the halt bit.
use anyhow::Context;
use cnidarium::StateWrite;
use cnidarium::{StateDelta, Storage};
use jmt::RootHash;
use penumbra_app::app::StateReadExt as _;
use penumbra_dex::lp::position;
use penumbra_dex::lp::position::Position;
use penumbra_governance::StateWriteExt;
use penumbra_proto::StateReadProto;
use penumbra_sct::component::clock::EpochManager;
use penumbra_sct::component::clock::EpochRead;
use std::path::PathBuf;
use tracing::instrument;

use crate::testnet::generate::TestnetConfig;

/// Run the full migration, given an export path and a start time for genesis.
///
/// Menu:
/// - Reconstruct a correct VCB balance for the auction component.
#[instrument]
pub async fn migrate(
    storage: Storage,
    pd_home: PathBuf,
    genesis_start: Option<tendermint::time::Time>,
) -> anyhow::Result<()> {
    let initial_state = storage.latest_snapshot();
    let chain_id = initial_state.get_chain_id().await?;
    let root_hash = initial_state
        .root_hash()
        .await
        .expect("chain state has a root hash");
    let pre_upgrade_root_hash: RootHash = root_hash.into();
    let pre_upgrade_height = initial_state
        .get_block_height()
        .await
        .expect("chain state has a block height");
    let post_upgrade_height = pre_upgrade_height.wrapping_add(1);

    // We want to equip the price indices with corresponding positions.
    // To do this, we could either:
    // Iterate over each position in the DEX
    // Update the price index
    // This is tricky because they are internals.
    // Or
    // Iterate over each position in the dex
    // Store the positions
    // Close them
    // Open new positions.
    /* Migration */
    use futures::StreamExt;
    use penumbra_dex::component::PositionManager;

    let mut delta = StateDelta::new(initial_state);

    let prefix_lp = penumbra_dex::state_key::all_positions();
    let mut position_stream = delta.prefix::<Position>(&prefix_lp);

    while let Some(entry) = position_stream.next().await {
        let (_, lp) = entry?;
        if lp.state != position::State::Opened {
            continue;
        }

        // Re-hash the position, since the key is a bech32 string.
        let id = lp.id();
        // Close the position, adjusting all its index entries.
        delta.close_position_by_id(&id).await?;
        // Erase the position from the state, so that `update_position` allows the write.
        delta.delete(penumbra_dex::state_key::position_by_id(&id));
        // Open a position with the adjusted indexing logic.
        delta.open_position(lp).await?;
    }

    /* ********* */

    // Reset the halt flag, and application height.
    delta.ready_to_start();
    delta.put_block_height(0u64);
    let _ = storage
        .commit_in_place(delta)
        .await
        .context("failed to reset halt bit")?;
    storage.release().await;

    // The migration is complete, now we need to generate a genesis file. To do this, we need
    // to lookup a validator view from the chain, and specify the post-upgrade app hash and
    // initial height.
    let app_state = penumbra_app::genesis::Content {
        chain_id,
        ..Default::default()
    };
    let mut genesis = TestnetConfig::make_genesis(app_state.clone()).expect("can make genesis");
    genesis.app_hash = pre_upgrade_root_hash
        .0
        .to_vec()
        .try_into()
        .expect("infallible conversion");

    genesis.initial_height = post_upgrade_height as i64;
    genesis.genesis_time = genesis_start.unwrap_or_else(|| {
        let now = tendermint::time::Time::now();
        tracing::info!(%now, "no genesis time provided, detecting a testing setup");
        now
    });
    let checkpoint = pre_upgrade_root_hash.0.to_vec();
    let genesis = TestnetConfig::make_checkpoint(genesis, Some(checkpoint));
    let genesis_json = serde_json::to_string(&genesis).expect("can serialize genesis");
    tracing::info!("genesis: {}", genesis_json);
    let genesis_path = pd_home.join("genesis.json");
    std::fs::write(genesis_path, genesis_json).expect("can write genesis");

    let validator_state_path = pd_home.join("priv_validator_state.json");
    let fresh_validator_state = crate::testnet::generate::TestnetValidator::initial_state();
    std::fs::write(validator_state_path, fresh_validator_state).expect("can write validator state");

    tracing::info!(
        pre_upgrade_height,
        ?pre_upgrade_root_hash,
        "successful migration!"
    );

    Ok(())
}
