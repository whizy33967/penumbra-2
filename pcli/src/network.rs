use penumbra_proto::{
    client::v1alpha1::{
        oblivious_query_service_client::ObliviousQueryServiceClient,
        specific_query_service_client::SpecificQueryServiceClient,
        tendermint_proxy_service_client::TendermintProxyServiceClient,
    },
    DomainType,
};
use penumbra_transaction::{plan::TransactionPlan, Id as TransactionId, Transaction};
use penumbra_view::ViewClient;
use rand_core::OsRng;
use std::future::Future;
use tonic::transport::Channel;
use tracing::instrument;

use crate::App;

impl App {
    pub async fn build_and_submit_transaction(
        &mut self,
        plan: TransactionPlan,
    ) -> anyhow::Result<TransactionId> {
        let transaction = self.build_transaction(plan).await?;
        self.submit_transaction(transaction).await
    }

    pub fn build_transaction(
        &mut self,
        plan: TransactionPlan,
    ) -> impl Future<Output = anyhow::Result<Transaction>> + '_ {
        println!("building transaction...");
        let start = std::time::Instant::now();
        let tx = penumbra_wallet::build_transaction(
            &self.fvk,
            self.view.as_mut().unwrap(),
            &mut self.custody,
            OsRng,
            plan,
        );
        async move {
            let tx = tx.await?;
            let elapsed = start.elapsed();
            println!(
                "finished proving in {}.{:03} seconds [{} actions, {} proofs, {} bytes]",
                elapsed.as_secs(),
                elapsed.subsec_millis(),
                tx.actions().count(),
                tx.num_proofs(),
                tx.encode_to_vec().len()
            );
            Ok(tx)
        }
    }

    /// Submits a transaction to the network.
    pub async fn submit_transaction(
        &mut self,
        transaction: Transaction,
    ) -> Result<TransactionId, anyhow::Error> {
        println!("broadcasting transaction and awaiting confirmation...");
        let id = self.view().broadcast_transaction(transaction, true).await?;
        println!("transaction confirmed and detected: {}", id);
        Ok(id)
    }

    /// Submits a transaction to the network, returning `Ok` as soon as the
    /// transaction has been submitted, rather than waiting for confirmation.
    #[instrument(skip(self, transaction))]
    pub async fn submit_transaction_unconfirmed(
        &mut self,
        transaction: Transaction,
    ) -> Result<(), anyhow::Error> {
        println!("broadcasting transaction without confirmation...");
        self.view()
            .broadcast_transaction(transaction, false)
            .await?;

        Ok(())
    }

    pub async fn specific_client(
        &self,
    ) -> Result<SpecificQueryServiceClient<Channel>, anyhow::Error> {
        SpecificQueryServiceClient::connect(self.pd_url.as_ref().to_owned())
            .await
            .map_err(Into::into)
    }

    pub async fn oblivious_client(
        &self,
    ) -> Result<ObliviousQueryServiceClient<Channel>, anyhow::Error> {
        ObliviousQueryServiceClient::connect(self.pd_url.as_ref().to_owned())
            .await
            .map_err(Into::into)
    }

    pub async fn tendermint_proxy_client(
        &self,
    ) -> Result<TendermintProxyServiceClient<Channel>, anyhow::Error> {
        TendermintProxyServiceClient::connect(self.pd_url.as_ref().to_owned())
            .await
            .map_err(Into::into)
    }
}
