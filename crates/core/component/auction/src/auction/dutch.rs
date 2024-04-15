use anyhow::anyhow;
use penumbra_asset::{asset, Value};
use penumbra_dex::lp::position::{self};
use penumbra_num::Amount;
use penumbra_proto::{core::component::auction::v1alpha1 as pb, DomainType};
use serde::{Deserialize, Serialize};

/// A deployed Dutch Auction, containing an immutable description
/// and stateful data about its current state.
#[derive(Clone, Serialize, Deserialize, PartialEq, Eq)]
#[serde(try_from = "pb::DutchAuction", into = "pb::DutchAuction")]
pub struct DutchAuction {
    pub description: DutchAuctionDescription,
    pub state: DutchAuctionState,
}

/* Protobuf impls for `DutchAuction` */
impl DomainType for DutchAuction {
    type Proto = pb::DutchAuction;
}

impl From<DutchAuction> for pb::DutchAuction {
    fn from(domain: DutchAuction) -> Self {
        pb::DutchAuction {
            description: Some(domain.description.into()),
            state: Some(domain.state.into()),
        }
    }
}

impl TryFrom<pb::DutchAuction> for DutchAuction {
    type Error = anyhow::Error;

    fn try_from(msg: pb::DutchAuction) -> Result<Self, Self::Error> {
        Ok(DutchAuction {
            description: msg
                .description
                .ok_or_else(|| anyhow!("DutchAuction is missing description"))?
                .try_into()?,
            state: msg
                .state
                .ok_or_else(|| anyhow!("DutchAuction is missing a state field"))?
                .try_into()?,
        })
    }
}
/* ********************************** */

/// A description of the immutable parts of a dutch auction.
#[derive(Clone, Serialize, Deserialize, PartialEq, Eq)]
#[serde(
    try_from = "pb::DutchAuctionDescription",
    into = "pb::DutchAuctionDescription"
)]
pub struct DutchAuctionDescription {
    pub input: Value,
    pub output_id: asset::Id,
    pub max_output: Amount,
    pub min_output: Amount,
    pub start_height: u64,
    pub end_height: u64,
    pub step_count: u64,
    pub nonce: [u8; 32],
}

/* Protobuf impls */
impl DomainType for DutchAuctionDescription {
    type Proto = pb::DutchAuctionDescription;
}

impl From<DutchAuctionDescription> for pb::DutchAuctionDescription {
    fn from(domain: DutchAuctionDescription) -> Self {
        Self {
            input: Some(domain.input.into()),
            output_id: Some(domain.output_id.into()),
            max_output: Some(domain.max_output.into()),
            min_output: Some(domain.min_output.into()),
            start_height: domain.start_height,
            end_height: domain.end_height,
            step_count: domain.step_count,
            nonce: domain.nonce.as_slice().to_vec(),
        }
    }
}

impl TryFrom<pb::DutchAuctionDescription> for DutchAuctionDescription {
    type Error = anyhow::Error;

    fn try_from(msg: pb::DutchAuctionDescription) -> Result<Self, Self::Error> {
        let d = DutchAuctionDescription {
            input: msg
                .input
                .ok_or_else(|| anyhow!("DutchAuctionDescription message is missing input"))?
                .try_into()?,
            output_id: msg
                .output_id
                .ok_or_else(|| {
                    anyhow!("DutchAuctionDescription message is missing an output identifier")
                })?
                .try_into()?,
            max_output: msg
                .max_output
                .ok_or_else(|| anyhow!("DutchAuctionDescription message is missing max output"))?
                .try_into()?,
            min_output: msg
                .min_output
                .ok_or_else(|| anyhow!("DutchAuctionDescription message is missing min output"))?
                .try_into()?,
            start_height: msg.start_height,
            end_height: msg.end_height,
            step_count: msg.step_count,
            nonce: msg.nonce.as_slice().try_into()?,
        };
        Ok(d)
    }
}
/* ********************************** */

/// A stateful description of a dutch auction, recording its state (via a sequence number),
/// the current position id associated to it (if any), and its amount IO.
/// # State
/// We record the state of the dutch auction via an untyped `u64` instead of an enum.
/// This futureproof support for auction types that have a richer state machine e.g. allows
/// claiming a withdrawn auction multiple times, burning and minting a new withdrawn auction
/// with an incremented sequence number.
///
/// For Dutch auctions:
///
///   ┌───┐            ┌───┐             ┌───┐
///   │ 0 │───Closed──▶│ 1 │──Withdrawn─▶│ 2 │
///   └───┘            └───┘             └───┘
///     ▲                                     
///     │                                     
///  Opened                                   
///     │                                     
///
#[derive(Clone, Serialize, Deserialize, PartialEq, Eq)]
#[serde(try_from = "pb::DutchAuctionState", into = "pb::DutchAuctionState")]
pub struct DutchAuctionState {
    pub sequence: u64,
    pub current_position: Option<position::Id>,
    pub next_trigger: u64,
    pub input_reserves: Amount,
    pub output_reserves: Amount,
}

/* Protobuf impls for `DutchAuctionState` */
impl DomainType for DutchAuctionState {
    type Proto = pb::DutchAuctionState;
}

impl From<DutchAuctionState> for pb::DutchAuctionState {
    fn from(domain: DutchAuctionState) -> Self {
        Self {
            seq: domain.sequence,
            current_position: domain.current_position.map(Into::into),
            next_trigger: domain.next_trigger,
            input_reserves: Some(domain.input_reserves.into()),
            output_reserves: Some(domain.output_reserves.into()),
        }
    }
}

impl TryFrom<pb::DutchAuctionState> for DutchAuctionState {
    type Error = anyhow::Error;

    fn try_from(msg: pb::DutchAuctionState) -> Result<Self, Self::Error> {
        let current_position = msg.current_position.map(TryInto::try_into).transpose()?;

        let domain_type = DutchAuctionState {
            sequence: msg.seq,
            current_position: domain_position_id,
            next_trigger: msg.next_trigger,
            input_reserves: msg
                .input_reserves
                .ok_or_else(|| anyhow!("DutchAuctionState message is missing input reserves"))?
                .try_into()?,
            output_reserves: msg
                .output_reserves
                .ok_or_else(|| anyhow!("DutchAuctionState message is missing output reserves"))?
                .try_into()?,
        };

        Ok(domain_type)
    }
}
/* ********************************** */
