use std::convert::{TryFrom, TryInto};

use penumbra_crypto::balance;
use penumbra_proto::{
    core::ibc::v1alpha1 as pb_ibc, core::stake::v1alpha1 as pbs, core::transaction::v1alpha1 as pb,
    DomainType,
};

mod dao_deposit;
mod dao_output;
mod dao_spend;
mod delegate;
mod delegator_vote;
mod ics20_withdrawal;
pub mod output;
mod position;
mod proposal_deposit_claim;
mod proposal_submit;
mod proposal_withdraw;
pub mod spend;
pub mod swap;
pub mod swap_claim;
mod undelegate;
mod undelegate_claim;
mod validator_vote;

use crate::{ActionView, TransactionPerspective};

pub use self::ics20_withdrawal::Ics20Withdrawal;
pub use crate::proposal::{Proposal, ProposalKind, ProposalPayload};
pub use crate::vote::Vote;
pub use dao_deposit::DaoDeposit;
pub use dao_output::DaoOutput;
pub use dao_spend::DaoSpend;
pub use delegate::Delegate;
pub use delegator_vote::{DelegatorVote, DelegatorVoteBody};
pub use output::Output;
pub use position::{PositionClose, PositionOpen, PositionRewardClaim, PositionWithdraw};
pub use proposal_deposit_claim::ProposalDepositClaim;
pub use proposal_submit::ProposalSubmit;
pub use proposal_withdraw::ProposalWithdraw;
pub use spend::Spend;
pub use swap::Swap;
pub use swap_claim::SwapClaim;
pub use undelegate::Undelegate;
pub use undelegate_claim::{UndelegateClaim, UndelegateClaimBody};
pub use validator_vote::{ValidatorVote, ValidatorVoteBody};

/// Common behavior between Penumbra actions.
pub trait IsAction {
    fn balance_commitment(&self) -> balance::Commitment;
    fn view_from_perspective(&self, txp: &TransactionPerspective) -> ActionView;
}

/// An action performed by a Penumbra transaction.
#[derive(Clone, Debug)]
#[allow(clippy::large_enum_variant)]
pub enum Action {
    Output(output::Output),
    Spend(spend::Spend),
    ValidatorDefinition(pbs::ValidatorDefinition),
    IBCAction(pb_ibc::IbcAction),
    Swap(Swap),
    SwapClaim(SwapClaim),
    ProposalSubmit(ProposalSubmit),
    ProposalWithdraw(ProposalWithdraw),
    DelegatorVote(DelegatorVote),
    ValidatorVote(ValidatorVote),
    ProposalDepositClaim(ProposalDepositClaim),

    PositionOpen(PositionOpen),
    PositionClose(PositionClose),
    PositionWithdraw(PositionWithdraw),
    PositionRewardClaim(PositionRewardClaim),

    Delegate(Delegate),
    Undelegate(Undelegate),
    UndelegateClaim(UndelegateClaim),

    Ics20Withdrawal(Ics20Withdrawal),

    DaoSpend(DaoSpend),
    DaoOutput(DaoOutput),
    DaoDeposit(DaoDeposit),
}

impl Action {
    /// Create a tracing span to track execution related to this action.
    ///
    /// The `idx` parameter is the index of this action in the transaction.
    pub fn create_span(&self, idx: usize) -> tracing::Span {
        match self {
            Action::Output(_) => tracing::info_span!("Output", ?idx),
            Action::Spend(_) => tracing::info_span!("Spend", ?idx),
            Action::ValidatorDefinition(_) => {
                tracing::info_span!("ValidatorDefinition", ?idx)
            }
            Action::IBCAction(_) => tracing::info_span!("IbcAction", ?idx),
            Action::Swap(_) => tracing::info_span!("Swap", ?idx),
            Action::SwapClaim(_) => tracing::info_span!("SwapClaim", ?idx),
            Action::ProposalSubmit(_) => tracing::info_span!("ProposalSubmit", ?idx),
            Action::ProposalWithdraw(_) => {
                tracing::info_span!("ProposalWithdraw", ?idx)
            }
            Action::DelegatorVote(_) => tracing::info_span!("DelegatorVote", ?idx),
            Action::ValidatorVote(_) => tracing::info_span!("ValidatorVote", ?idx),
            Action::ProposalDepositClaim(_) => {
                tracing::info_span!("ProposalDepositClaim", ?idx)
            }
            Action::PositionOpen(_) => tracing::info_span!("PositionOpen", ?idx),
            Action::PositionClose(_) => tracing::info_span!("PositionClose", ?idx),
            Action::PositionWithdraw(_) => {
                tracing::info_span!("PositionWithdraw", ?idx)
            }
            Action::PositionRewardClaim(_) => {
                tracing::info_span!("PositionRewardClaim", ?idx)
            }
            Action::Delegate(_) => tracing::info_span!("Delegate", ?idx),
            Action::Undelegate(_) => tracing::info_span!("Undelegate", ?idx),
            Action::UndelegateClaim(_) => tracing::info_span!("UndelegateClaim", ?idx),
            Action::Ics20Withdrawal(_) => tracing::info_span!("Ics20Withdrawal", ?idx),
            Action::DaoDeposit(_) => tracing::info_span!("DaoDeposit", ?idx),
            Action::DaoSpend(_) => tracing::info_span!("DaoSpend", ?idx),
            Action::DaoOutput(_) => tracing::info_span!("DaoOutput", ?idx),
        }
    }
}

impl IsAction for Action {
    fn balance_commitment(&self) -> balance::Commitment {
        match self {
            Action::Output(output) => output.balance_commitment(),
            Action::Spend(spend) => spend.balance_commitment(),
            Action::Delegate(delegate) => delegate.balance_commitment(),
            Action::Undelegate(undelegate) => undelegate.balance_commitment(),
            Action::UndelegateClaim(undelegate_claim) => undelegate_claim.balance_commitment(),
            Action::Swap(swap) => swap.balance_commitment(),
            Action::SwapClaim(swap_claim) => swap_claim.balance_commitment(),
            Action::ProposalSubmit(submit) => submit.balance_commitment(),
            Action::ProposalWithdraw(withdraw) => withdraw.balance_commitment(),
            Action::DelegatorVote(delegator_vote) => delegator_vote.balance_commitment(),
            Action::ValidatorVote(validator_vote) => validator_vote.balance_commitment(),
            Action::ProposalDepositClaim(p) => p.balance_commitment(),
            Action::PositionOpen(p) => p.balance_commitment(),
            Action::PositionClose(p) => p.balance_commitment(),
            Action::PositionWithdraw(p) => p.balance_commitment(),
            Action::PositionRewardClaim(p) => p.balance_commitment(),
            Action::Ics20Withdrawal(withdrawal) => withdrawal.balance_commitment(),
            Action::DaoDeposit(deposit) => deposit.balance_commitment(),
            Action::DaoSpend(spend) => spend.balance_commitment(),
            Action::DaoOutput(output) => output.balance_commitment(),
            // These actions just post Protobuf data to the chain, and leave the
            // value balance unchanged.
            Action::ValidatorDefinition(_) => balance::Commitment::default(),
            Action::IBCAction(_) => balance::Commitment::default(),
        }
    }

    fn view_from_perspective(&self, txp: &TransactionPerspective) -> ActionView {
        match self {
            Action::Swap(x) => x.view_from_perspective(txp),
            Action::SwapClaim(x) => x.view_from_perspective(txp),
            Action::Output(x) => x.view_from_perspective(txp),
            Action::Spend(x) => x.view_from_perspective(txp),
            Action::Delegate(x) => x.view_from_perspective(txp),
            Action::Undelegate(x) => x.view_from_perspective(txp),
            Action::UndelegateClaim(x) => x.view_from_perspective(txp),
            Action::ProposalSubmit(x) => x.view_from_perspective(txp),
            Action::ProposalWithdraw(x) => x.view_from_perspective(txp),
            Action::DelegatorVote(x) => x.view_from_perspective(txp),
            Action::ValidatorVote(x) => x.view_from_perspective(txp),
            Action::ProposalDepositClaim(x) => x.view_from_perspective(txp),
            Action::PositionOpen(x) => x.view_from_perspective(txp),
            Action::PositionClose(x) => x.view_from_perspective(txp),
            Action::PositionWithdraw(x) => x.view_from_perspective(txp),
            Action::PositionRewardClaim(x) => x.view_from_perspective(txp),
            Action::Ics20Withdrawal(x) => x.view_from_perspective(txp),
            Action::DaoSpend(x) => x.view_from_perspective(txp),
            Action::DaoOutput(x) => x.view_from_perspective(txp),
            Action::DaoDeposit(x) => x.view_from_perspective(txp),
            // TODO: figure out where to implement the actual decryption methods for these? where are their action definitions?
            Action::ValidatorDefinition(x) => ActionView::ValidatorDefinition(x.to_owned()),
            Action::IBCAction(x) => ActionView::IBCAction(x.to_owned()),
        }
    }
}

impl DomainType for Action {
    type Proto = pb::Action;
}

impl From<Action> for pb::Action {
    fn from(msg: Action) -> Self {
        match msg {
            Action::Output(inner) => pb::Action {
                action: Some(pb::action::Action::Output(inner.into())),
            },
            Action::Spend(inner) => pb::Action {
                action: Some(pb::action::Action::Spend(inner.into())),
            },
            Action::Delegate(inner) => pb::Action {
                action: Some(pb::action::Action::Delegate(inner.into())),
            },
            Action::Undelegate(inner) => pb::Action {
                action: Some(pb::action::Action::Undelegate(inner.into())),
            },
            Action::UndelegateClaim(inner) => pb::Action {
                action: Some(pb::action::Action::UndelegateClaim(inner.into())),
            },
            Action::ValidatorDefinition(inner) => pb::Action {
                action: Some(pb::action::Action::ValidatorDefinition(inner)),
            },
            Action::SwapClaim(inner) => pb::Action {
                action: Some(pb::action::Action::SwapClaim(inner.into())),
            },
            Action::Swap(inner) => pb::Action {
                action: Some(pb::action::Action::Swap(inner.into())),
            },
            Action::IBCAction(inner) => pb::Action {
                action: Some(pb::action::Action::IbcAction(inner)),
            },
            Action::ProposalSubmit(inner) => pb::Action {
                action: Some(pb::action::Action::ProposalSubmit(inner.into())),
            },
            Action::ProposalWithdraw(inner) => pb::Action {
                action: Some(pb::action::Action::ProposalWithdraw(inner.into())),
            },
            Action::DelegatorVote(inner) => pb::Action {
                action: Some(pb::action::Action::DelegatorVote(inner.into())),
            },
            Action::ValidatorVote(inner) => pb::Action {
                action: Some(pb::action::Action::ValidatorVote(inner.into())),
            },
            Action::ProposalDepositClaim(inner) => pb::Action {
                action: Some(pb::action::Action::ProposalDepositClaim(inner.into())),
            },
            Action::PositionOpen(inner) => pb::Action {
                action: Some(pb::action::Action::PositionOpen(inner.into())),
            },
            Action::PositionClose(inner) => pb::Action {
                action: Some(pb::action::Action::PositionClose(inner.into())),
            },
            Action::PositionWithdraw(inner) => pb::Action {
                action: Some(pb::action::Action::PositionWithdraw(inner.into())),
            },
            Action::PositionRewardClaim(inner) => pb::Action {
                action: Some(pb::action::Action::PositionRewardClaim(inner.into())),
            },
            Action::Ics20Withdrawal(withdrawal) => pb::Action {
                action: Some(pb::action::Action::Ics20Withdrawal(withdrawal.into())),
            },
            Action::DaoSpend(inner) => pb::Action {
                action: Some(pb::action::Action::DaoSpend(inner.into())),
            },
            Action::DaoOutput(inner) => pb::Action {
                action: Some(pb::action::Action::DaoOutput(inner.into())),
            },
            Action::DaoDeposit(inner) => pb::Action {
                action: Some(pb::action::Action::DaoDeposit(inner.into())),
            },
        }
    }
}

impl TryFrom<pb::Action> for Action {
    type Error = anyhow::Error;
    fn try_from(proto: pb::Action) -> anyhow::Result<Self, Self::Error> {
        if proto.action.is_none() {
            return Err(anyhow::anyhow!("missing action content"));
        }
        match proto.action.unwrap() {
            pb::action::Action::Output(inner) => Ok(Action::Output(inner.try_into()?)),
            pb::action::Action::Spend(inner) => Ok(Action::Spend(inner.try_into()?)),
            pb::action::Action::Delegate(inner) => Ok(Action::Delegate(inner.try_into()?)),
            pb::action::Action::Undelegate(inner) => Ok(Action::Undelegate(inner.try_into()?)),
            pb::action::Action::UndelegateClaim(inner) => {
                Ok(Action::UndelegateClaim(inner.try_into()?))
            }
            pb::action::Action::ValidatorDefinition(inner) => {
                Ok(Action::ValidatorDefinition(inner))
            }
            pb::action::Action::SwapClaim(inner) => Ok(Action::SwapClaim(inner.try_into()?)),
            pb::action::Action::Swap(inner) => Ok(Action::Swap(inner.try_into()?)),
            pb::action::Action::IbcAction(inner) => Ok(Action::IBCAction(inner)),
            pb::action::Action::ProposalSubmit(inner) => {
                Ok(Action::ProposalSubmit(inner.try_into()?))
            }
            pb::action::Action::ProposalWithdraw(inner) => {
                Ok(Action::ProposalWithdraw(inner.try_into()?))
            }
            pb::action::Action::DelegatorVote(inner) => {
                Ok(Action::DelegatorVote(inner.try_into()?))
            }
            pb::action::Action::ValidatorVote(inner) => {
                Ok(Action::ValidatorVote(inner.try_into()?))
            }
            pb::action::Action::ProposalDepositClaim(inner) => {
                Ok(Action::ProposalDepositClaim(inner.try_into()?))
            }

            pb::action::Action::PositionOpen(inner) => Ok(Action::PositionOpen(inner.try_into()?)),
            pb::action::Action::PositionClose(inner) => {
                Ok(Action::PositionClose(inner.try_into()?))
            }
            pb::action::Action::PositionWithdraw(inner) => {
                Ok(Action::PositionWithdraw(inner.try_into()?))
            }
            pb::action::Action::PositionRewardClaim(inner) => {
                Ok(Action::PositionRewardClaim(inner.try_into()?))
            }
            pb::action::Action::Ics20Withdrawal(inner) => {
                Ok(Action::Ics20Withdrawal(inner.try_into()?))
            }
            pb::action::Action::DaoSpend(inner) => Ok(Action::DaoSpend(inner.try_into()?)),
            pb::action::Action::DaoOutput(inner) => Ok(Action::DaoOutput(inner.try_into()?)),
            pb::action::Action::DaoDeposit(inner) => Ok(Action::DaoDeposit(inner.try_into()?)),
        }
    }
}
