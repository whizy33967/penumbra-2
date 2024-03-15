//! The Penumbra shielded pool [`Component`] and [`ActionHandler`] implementations.

mod action_handler;
mod metrics;
mod note_manager;
mod shielded_pool;
mod assets;
mod transfer;

pub use self::metrics::register_metrics;
pub use note_manager::NoteManager;
pub use shielded_pool::{ShieldedPool, StateReadExt, StateWriteExt};
pub use assets::{AssetRegistry, AssetRegistryRead};
pub use transfer::Ics20Transfer;

pub mod rpc;
