use std::sync::Arc;

use anyhow::Result;
use async_trait::async_trait;
use penumbra_storage::{StateRead, StateWrite};
use penumbra_transaction::{action::PositionOpen, Transaction};

use crate::action_handler::ActionHandler;
use crate::dex::PositionRead;

#[async_trait]
/// Debits the initial reserves and credits an opened position NFT.
impl ActionHandler for PositionOpen {
    async fn check_stateless(&self, _context: Arc<Transaction>) -> Result<()> {
        // The initial reserves must have a non-zero Amount for either `r1` or `r2`.
        if self.initial_reserves.r1.value() == 0 && self.initial_reserves.r2.value() == 0 {
            return Err(anyhow::anyhow!(
                "initial reserves must have a non-zero Amount for either `r1` or `r2`"
            ));
        }

        // The two assets in the position must be different.
        if self.position.phi.pair.asset_1() == self.position.phi.pair.asset_2() {
            return Err(anyhow::anyhow!(
                "the two assets in the position must be different"
            ));
        }

        // TODO: any other checks of the trading function that should be performed?

        Ok(())
    }

    async fn check_stateful<S: StateRead + 'static>(&self, _state: Arc<S>) -> Result<()> {
        // Validate that the position nonce is not already in use
        // TODO: this check is duplicated in `PositionManager::position_open`
        // but that's probably okay, checking here lets us skip out of execution sooner
        state.check_nonce_unused(&self.position).await?;

        // Validate that the position ID doesn't collide
        // It's important to reject all LP actions for now, to prevent
        // inflation / minting bugs until we implement all required checks
        // (e.g., minting tokens by withdrawing reserves we don't check)
        Err(anyhow::anyhow!("lp actions not supported yet"))
    }

    async fn execute<S: StateWrite>(&self, _state: S) -> Result<()> {
        // let position = self.position;
        // let initial_reserves = self.initial_reserves;
        // let lpnft = state.position_open(position, initial_reserves).await?;
        // TODO: implement

        Ok(())
    }
}
