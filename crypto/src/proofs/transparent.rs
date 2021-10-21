//! Transparent proofs for `MVP1` of the Penumbra system.

use ark_serialize::CanonicalSerialize;
use decaf377::FieldExt;
use std::convert::TryFrom;
use thiserror;

use crate::{ka, merkle, note, value, Fq, Fr, Value};

pub const OUTPUT_PROOF_LEN_BYTES: usize = 192;
// xx check the spend proof len
pub const SPEND_PROOF_LEN_BYTES: usize = 192;

#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("Invalid spend auth randomizer")]
    InvalidSpendAuthRandomizer,
}

pub struct SpendProof {
    pub spend_auth_randomizer: Fr,
    pub merkle_path: merkle::Path,
    // more TK
}

impl Into<[u8; SPEND_PROOF_LEN_BYTES]> for SpendProof {
    fn into(self) -> [u8; SPEND_PROOF_LEN_BYTES] {
        let mut spend_auth_randomizer_bytes = [0u8; 32];
        self.spend_auth_randomizer
            .serialize(&mut spend_auth_randomizer_bytes[..])
            .expect("serialization into array should be infallible");

        let mut bytes = [0u8; SPEND_PROOF_LEN_BYTES];
        bytes.copy_from_slice(&spend_auth_randomizer_bytes);

        // TODO: Merkle path serialization and add in here

        // When we put more stuff into this transparent spend proof, add here.
        bytes
    }
}

impl TryFrom<&[u8]> for SpendProof {
    type Error = Error;

    fn try_from(bytes: &[u8]) -> Result<SpendProof, Self::Error> {
        let mut spend_auth_randomizer_bytes = [0u8; 32];
        spend_auth_randomizer_bytes.copy_from_slice(&bytes[0..32]);

        let spend_auth_randomizer = Fr::from_bytes(spend_auth_randomizer_bytes)
            .map_err(|_| Error::InvalidSpendAuthRandomizer)?;

        // TODO! Merkle path serialization.
        Ok(SpendProof {
            spend_auth_randomizer,
            merkle_path: merkle::Path::default(),
        })
    }
}

/// Transparent proof for new note creation.
///
/// This structure keeps track of the auxiliary (private) inputs.
/// To generate the final proof, one calls `generate` and provides the
/// public inputs.
pub struct OutputProof {
    // The diversified base for the destination address.
    pub g_d: decaf377::Element,
    // The transmission key for the destination address.
    pub pk_d: ka::Public,
    // The value of the newly created note.
    pub value: Value,
    // The blinding factor used for generating the value commitment.
    pub v_blinding: Fr,
    // The blinding factor used for generating the note commitment.
    pub note_blinding: Fq,
    // The ephemeral secret key that corresponds to the public key.
    pub esk: ka::Secret,
}

impl OutputProof {
    /// Called to generate the proof using public inputs.
    ///
    /// The public inputs are:
    /// * value commitment of the new note,
    /// * note commitment of the new note,
    /// * the ephemeral public key used to generate the new note.
    pub fn generate(
        value_commitment: value::Commitment,
        note_commitment: note::Commitment,
        epk: ka::Public,
    ) {
        // This would return the generated proof.
        todo!()
    }
}

impl Into<[u8; OUTPUT_PROOF_LEN_BYTES]> for OutputProof {
    fn into(self) -> [u8; OUTPUT_PROOF_LEN_BYTES] {
        let bytes = [0u8; OUTPUT_PROOF_LEN_BYTES];
        // When we put more stuff into this transparent output proof, add here.
        bytes
    }
}
