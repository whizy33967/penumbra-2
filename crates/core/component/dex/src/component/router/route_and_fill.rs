use std::sync::Arc;

use anyhow::{Context, Result};
use async_trait::async_trait;
use cnidarium::StateWrite;
use penumbra_asset::{asset, Value};
use penumbra_num::Amount;
use penumbra_sct::component::clock::EpochRead;
use tracing::instrument;

use crate::{
    component::{
        chandelier::Chandelier,
        flow::SwapFlow,
        router::{FillRoute, PathSearch, RoutingParams},
        ExecutionCircuitBreaker, InternalDexWrite, PositionManager,
    },
    lp::position::MAX_RESERVE_AMOUNT,
    BatchSwapOutputData, SwapExecution, TradingPair,
};

use super::fill_route::FillError;

/// Ties together the routing and filling logic, to process
/// a block's batch swap flows.
#[async_trait]
pub trait HandleBatchSwaps: StateWrite + Sized {
    #[instrument(skip(self, trading_pair, batch_data, block_height, params))]
    async fn handle_batch_swaps(
        self: &mut Arc<Self>,
        trading_pair: TradingPair,
        batch_data: SwapFlow,
        // This will be read from the ABCI request
        block_height: u64,
        params: RoutingParams,
    ) -> Result<()>
    where
        Self: 'static,
    {
        let (delta_1, delta_2) = (batch_data.0, batch_data.1);

        tracing::debug!(?delta_1, ?delta_2, ?trading_pair, "decrypted batch swaps");

        let execution_circuit_breaker = ExecutionCircuitBreaker::default();

        let swap_execution_1_for_2 = if delta_1.value() > 0 {
            Some(
                self.route_and_fill(
                    trading_pair.asset_1(),
                    trading_pair.asset_2(),
                    delta_1,
                    params.clone(),
                    execution_circuit_breaker.clone(),
                )
                .await?,
            )
        } else {
            tracing::debug!("no input for asset 1, skipping 1=>2 routing and execution");
            None
        };

        let swap_execution_2_for_1 = if delta_2.value() > 0 {
            Some(
                self.route_and_fill(
                    trading_pair.asset_2(),
                    trading_pair.asset_1(),
                    delta_2,
                    params.clone(),
                    execution_circuit_breaker,
                )
                .await?,
            )
        } else {
            tracing::debug!("no input for asset 2, skipping 2=>1 execution");
            None
        };

        let (lambda_2, unfilled_1) = match &swap_execution_1_for_2 {
            Some(swap_execution) => (
                swap_execution.output.amount,
                delta_1 - swap_execution.input.amount,
            ),
            None => (0u64.into(), delta_1),
        };
        let (lambda_1, unfilled_2) = match &swap_execution_2_for_1 {
            Some(swap_execution) => (
                swap_execution.output.amount,
                delta_2 - swap_execution.input.amount,
            ),
            None => (0u64.into(), delta_2),
        };
        let epoch = self.get_current_epoch().await.expect("epoch is set");
        let output_data = BatchSwapOutputData {
            height: block_height,
            trading_pair,
            delta_1,
            delta_2,
            lambda_1,
            lambda_2,
            unfilled_1,
            unfilled_2,
            sct_position_prefix: (
                u16::try_from(epoch.index).expect("epoch index should be small enough"),
                // The block index is determined by looking at how many blocks have elapsed since
                // the start of the epoch.
                u16::try_from(block_height - epoch.start_height)
                    .expect("block index should be small enough"),
                0,
            )
                .into(),
        };

        tracing::debug!(
            ?output_data,
            ?swap_execution_1_for_2,
            ?swap_execution_2_for_1
        );

        // Update the candlestick tracking
        if let Some(se) = swap_execution_1_for_2.clone() {
            tracing::debug!("updating candlestick for 1=>2 swap");
            Arc::get_mut(self)
                .expect("expected state to have no other refs")
                .record_swap_execution(&se)
                .await;
        }
        if let Some(se) = &swap_execution_2_for_1 {
            tracing::debug!("updating candlestick for 2=>1 swap");
            Arc::get_mut(self)
                .expect("expected state to have no other refs")
                .record_swap_execution(se)
                .await;
        }

        // Fetch the swap execution object that should have been modified during the routing and filling.
        Arc::get_mut(self)
            .expect("expected state to have no other refs")
            .set_output_data(output_data, swap_execution_1_for_2, swap_execution_2_for_1)
            .await?;

        Ok(())
    }
}

impl<T: PositionManager> HandleBatchSwaps for T {}

/// Lower-level trait that ties together the routing and filling logic.
#[async_trait]
pub trait RouteAndFill: StateWrite + Sized {
    #[instrument(skip(self, asset_1, asset_2, input, params, execution_circuit_breaker))]
    async fn route_and_fill(
        self: &mut Arc<Self>,
        asset_1: asset::Id,
        asset_2: asset::Id,
        input: Amount,
        params: RoutingParams,
        mut execution_circuit_breaker: ExecutionCircuitBreaker,
    ) -> Result<SwapExecution>
    where
        Self: 'static,
    {
        tracing::debug!(?input, ?asset_1, ?asset_2, "starting route_and_fill");

        // Unfilled output of asset 1
        let mut total_unfilled_1 = input;
        // Output of asset 2
        let mut total_output_2 = 0u64.into();

        // An ordered list of execution traces that were used to fill the trade.
        let mut traces: Vec<Vec<Value>> = Vec::new();

        let max_delta_1: Amount = MAX_RESERVE_AMOUNT.into();

        let mut seen_position_sets = Vec::new();

        // Termination conditions:
        // 1. We have no more delta_1 remaining
        // 2. A path can no longer be found
        // 3. We have reached the `RoutingParams` specified price limit
        // 4. The execution circuit breaker has been triggered based on the number of path searches and executions

        loop {
            // Check if we have exceeded the execution circuit breaker limits.
            if execution_circuit_breaker.exceeded_limits() {
                tracing::debug!("execution circuit breaker triggered, exiting route_and_fill");
                break;
            }

            // Find the best route between the two assets in the trading pair.
            let (path, spill_price) = self
                .path_search(asset_1, asset_2, params.clone())
                .await
                .context("error finding best path")?;

            let Some(path) = path else {
                tracing::debug!("no path found, exiting route_and_fill");
                break;
            };

            if path.is_empty() {
                tracing::debug!("empty path found, exiting route_and_fill");
                break;
            }

            // Increment the execution circuit breaker path search counter.
            execution_circuit_breaker.current_path_searches += 1;

            let delta_1 = Value {
                amount: total_unfilled_1.min(max_delta_1),
                asset_id: asset_1,
            };

            tracing::debug!(?path, delta_1 = ?delta_1.amount, "found path, filling up to spill price");

            let execution = Arc::get_mut(self)
                .expect("expected state to have no other refs")
                .fill_route(delta_1, &path, spill_price)
                .await;
            println!("execution: {:?}", execution);

            let (execution, positions_executed) = match execution {
                Ok((execution, positions_executed)) => (execution, positions_executed),
                Err(FillError::ExecutionOverflow(position_id)) => {
                    // We have encountered an overflow during the execution of the route.
                    // To route around this, we will close the position and try to route and fill again.
                    tracing::debug!(culprit = ?position_id, "overflow detected during routing execution");
                    Arc::get_mut(self)
                        .expect("expected state to have no other refs")
                        .close_position_by_id(&position_id)
                        .await
                        .expect("the position still exists");
                    continue;
                }
                Err(e) => {
                    // We have encountered an error during the execution of the route,
                    // there are no clear ways to route around this, so we propagate the error.
                    // `fill_route` is transactional and will have rolled back the state.
                    anyhow::bail!("error filling route: {:?}", e);
                }
            };

            println!("execution positions: {:?}", positions_executed);
            // If we've seen this set of positions before, avoid executing against them again.
            // Ideally this would be fed into path search but that's not feasible.
            if seen_position_sets.contains(&positions_executed) {
                tracing::debug!("repeated position set");
                break;
            }
            seen_position_sets.push(positions_executed);

            // Immediately track the execution in the state.
            (total_output_2, total_unfilled_1) = {
                let lambda_2 = execution.output;
                let unfilled_1 = Value {
                    amount: total_unfilled_1
                        .checked_sub(&execution.input.amount)
                        .expect("unable to subtract unfilled input from total input"),
                    asset_id: asset_1,
                };
                tracing::debug!(input = ?delta_1.amount, output = ?lambda_2.amount, unfilled = ?unfilled_1.amount, "filled along best path");

                assert_eq!(lambda_2.asset_id, asset_2);
                assert_eq!(unfilled_1.asset_id, asset_1);

                // Append the traces from this execution to the outer traces.
                traces.append(&mut execution.traces.clone());

                (
                    total_output_2 + lambda_2.amount,
                    total_unfilled_1 - delta_1.amount + unfilled_1.amount,
                )
            };

            // Increment the execution circuit breaker execution counter.
            execution_circuit_breaker.current_executions += 1;

            if total_unfilled_1.value() == 0 {
                tracing::debug!("filled all input, exiting route_and_fill");
                break;
            }

            // Ensure that we've actually executed, or else bail out.
            let Some(accurate_max_price) = execution.max_price() else {
                tracing::debug!("no traces in execution, exiting route_and_fill");
                break;
            };

            // Check that the execution price is below the price limit, if one is set.
            if let Some(price_limit) = params.price_limit {
                if accurate_max_price >= price_limit {
                    tracing::debug!(
                        ?accurate_max_price,
                        ?price_limit,
                        "execution price above price limit, exiting route_and_fill"
                    );
                    break;
                }
            }
        }

        Ok(SwapExecution {
            traces,
            input: Value {
                asset_id: asset_1,
                amount: input - total_unfilled_1,
            },
            output: Value {
                asset_id: asset_2,
                amount: total_output_2,
            },
        })
    }
}

impl<T: HandleBatchSwaps> RouteAndFill for T {}
