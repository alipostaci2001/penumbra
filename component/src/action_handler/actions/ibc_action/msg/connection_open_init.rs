use std::sync::Arc;

use anyhow::Result;
use async_trait::async_trait;
use ibc::core::ics03_connection::msgs::conn_open_init::MsgConnectionOpenInit;
use penumbra_storage::{State, StateTransaction};
use penumbra_transaction::Transaction;
use tracing::instrument;

use crate::action_handler::ActionHandler;
use crate::ibc::component::connection::execution::connection_open_init::ConnectionOpenInitExecute;
use crate::ibc::component::connection::stateful::connection_open_init::ConnectionOpenInitCheck;
use crate::ibc::component::connection::stateless::connection_open_init::version_is_supported;

#[async_trait]
impl ActionHandler for MsgConnectionOpenInit {
    #[instrument(name = "connection_open_init", skip(self, _context))]
    fn check_stateless(&self, _context: Arc<Transaction>) -> Result<()> {
        version_is_supported(self)?;

        Ok(())
    }

    #[instrument(name = "connection_open_init", skip(self, state, _context))]
    async fn check_stateful(&self, state: Arc<State>, _context: Arc<Transaction>) -> Result<()> {
        state.validate(self).await?;

        Ok(())
    }

    #[instrument(name = "connection_open_init", skip(self, state))]
    async fn execute(&self, state: &mut StateTransaction) -> Result<()> {
        state.execute(self).await;

        Ok(())
    }
}