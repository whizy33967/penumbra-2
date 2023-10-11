//! A module for grouping several setup ceremonies into a single one.
//!
//! This also matches the coordination strategy we have for phase2,
//! along with the corresponding protobufs.
use std::array;

use crate::single::{
    circuit_degree, group::F, log::ContributionHash, DLogProof, Phase1CRSElements,
    Phase1RawCRSElements, Phase2CRSElements, Phase2Contribution, Phase2RawCRSElements,
    Phase2RawContribution,
};
use anyhow::{anyhow, Result};
use ark_relations::r1cs::ConstraintMatrices;
use ark_serialize::{CanonicalDeserialize, CanonicalSerialize};
use penumbra_dex::{swap::proof::SwapCircuit, swap_claim::proof::SwapClaimCircuit};
use penumbra_governance::DelegatorVoteCircuit;
use penumbra_proof_params::generate_constraint_matrices;
use penumbra_proto::tools::summoning::v1alpha1::{self as pb};
use penumbra_shielded_pool::{NullifierDerivationCircuit, OutputCircuit, SpendCircuit};
use penumbra_stake::UndelegateClaimCircuit;

use rand_core::CryptoRngCore;

// Some helper functions since we have to use these seventeen billion times

fn to_bytes<T: CanonicalSerialize>(t: &T) -> Result<Vec<u8>> {
    let mut out = Vec::new();
    t.serialize_compressed(&mut out)?;
    Ok(out)
}

pub const NUM_CIRCUITS: usize = 7;

/// Generate all of the circuits as matrices.
fn circuits() -> [ConstraintMatrices<F>; NUM_CIRCUITS] {
    println!("GENERATING CIRCUITS?");
    [
        generate_constraint_matrices::<SpendCircuit>(),
        generate_constraint_matrices::<OutputCircuit>(),
        generate_constraint_matrices::<DelegatorVoteCircuit>(),
        generate_constraint_matrices::<UndelegateClaimCircuit>(),
        generate_constraint_matrices::<SwapCircuit>(),
        generate_constraint_matrices::<SwapClaimCircuit>(),
        generate_constraint_matrices::<NullifierDerivationCircuit>(),
    ]
}

/// Holds all of the CRS elements for phase2 in one struct, before validation.
#[derive(Clone, Debug)]
pub struct Phase2RawCeremonyCRS([Phase2RawCRSElements; NUM_CIRCUITS]);

impl Phase2RawCeremonyCRS {
    /// Skip validation, performing the conversion anyways.
    ///
    /// Useful when parsing known good data.
    pub fn assume_valid(self) -> Phase2CeremonyCRS {
        match self.0 {
            [x0, x1, x2, x3, x4, x5, x6] => Phase2CeremonyCRS([
                x0.assume_valid(),
                x1.assume_valid(),
                x2.assume_valid(),
                x3.assume_valid(),
                x4.assume_valid(),
                x5.assume_valid(),
                x6.assume_valid(),
            ]),
        }
    }
}

impl TryInto<pb::CeremonyCrs> for Phase2RawCeremonyCRS {
    type Error = anyhow::Error;

    fn try_into(self) -> Result<pb::CeremonyCrs> {
        Ok(pb::CeremonyCrs {
            spend: to_bytes(&self.0[0])?,
            output: to_bytes(&self.0[1])?,
            delegator_vote: to_bytes(&self.0[2])?,
            undelegate_claim: to_bytes(&self.0[3])?,
            swap: to_bytes(&self.0[4])?,
            swap_claim: to_bytes(&self.0[5])?,
            nullifer_derivation_crs: to_bytes(&self.0[6])?,
        })
    }
}

impl TryFrom<pb::CeremonyCrs> for Phase2RawCeremonyCRS {
    type Error = anyhow::Error;

    fn try_from(value: pb::CeremonyCrs) -> std::result::Result<Self, Self::Error> {
        Ok(Self([
            Phase2RawCRSElements::deserialize_compressed(value.spend.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.output.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.delegator_vote.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.undelegate_claim.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.swap.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.swap_claim.as_slice())?,
            Phase2RawCRSElements::deserialize_compressed(value.nullifer_derivation_crs.as_slice())?,
        ]))
    }
}

/// Holds all of the CRS elements for phase2 in one struct.
#[derive(Clone, Debug)]
pub struct Phase2CeremonyCRS([Phase2CRSElements; NUM_CIRCUITS]);

impl From<Phase2CeremonyCRS> for Phase2RawCeremonyCRS {
    fn from(value: Phase2CeremonyCRS) -> Self {
        Self(array::from_fn(|i| value.0[i].raw.clone()))
    }
}

impl TryFrom<Phase2CeremonyCRS> for pb::CeremonyCrs {
    type Error = anyhow::Error;

    fn try_from(data: Phase2CeremonyCRS) -> Result<pb::CeremonyCrs> {
        Phase2RawCeremonyCRS::from(data).try_into()
    }
}

impl Phase2CeremonyCRS {
    pub fn root() -> Result<Self> {
        let [c0, c1, c2, c3, c4, c5, c6] = circuits();
        println!("GENERATED CIRCUITS");
        Ok(Self([
            Phase2CRSElements::dummy_root(circuit_degree(&c0)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c1)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c2)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c3)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c4)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c5)?),
            Phase2CRSElements::dummy_root(circuit_degree(&c6)?),
        ]))
    }
}

/// All phase2 contributions, before they've been validated.
#[derive(Clone, Debug)]
pub struct Phase2RawCeremonyContribution([Phase2RawContribution; NUM_CIRCUITS]);

impl TryInto<pb::participate_request::Contribution> for Phase2RawCeremonyContribution {
    type Error = anyhow::Error;

    fn try_into(self) -> Result<pb::participate_request::Contribution> {
        Ok(pb::participate_request::Contribution {
            updated: Some(pb::CeremonyCrs {
                spend: to_bytes(&self.0[0].new_elements)?,
                output: to_bytes(&self.0[1].new_elements)?,
                delegator_vote: to_bytes(&self.0[2].new_elements)?,
                undelegate_claim: to_bytes(&self.0[3].new_elements)?,
                swap: to_bytes(&self.0[4].new_elements)?,
                swap_claim: to_bytes(&self.0[5].new_elements)?,
                nullifer_derivation_crs: to_bytes(&self.0[6].new_elements)?,
            }),
            update_proofs: Some(pb::CeremonyLinkingProof {
                spend: to_bytes(&self.0[0].linking_proof)?,
                output: to_bytes(&self.0[1].linking_proof)?,
                delegator_vote: to_bytes(&self.0[2].linking_proof)?,
                undelegate_claim: to_bytes(&self.0[3].linking_proof)?,
                swap: to_bytes(&self.0[4].linking_proof)?,
                swap_claim: to_bytes(&self.0[5].linking_proof)?,
                nullifer_derivation_crs: to_bytes(&self.0[6].linking_proof)?,
            }),
            parent_hashes: Some(pb::CeremonyParentHashes {
                spend: self.0[0].parent.0.to_vec(),
                output: self.0[1].parent.0.to_vec(),
                delegator_vote: self.0[2].parent.0.to_vec(),
                undelegate_claim: self.0[3].parent.0.to_vec(),
                swap: self.0[4].parent.0.to_vec(),
                swap_claim: self.0[5].parent.0.to_vec(),
                nullifer_derivation_crs: self.0[6].parent.0.to_vec(),
            }),
        })
    }
}

impl TryFrom<pb::participate_request::Contribution> for Phase2RawCeremonyContribution {
    type Error = anyhow::Error;

    fn try_from(value: pb::participate_request::Contribution) -> Result<Self> {
        Ok(Self([
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .spend
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .spend
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .spend
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .output
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .output
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .output
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .delegator_vote
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .delegator_vote
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .delegator_vote
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .undelegate_claim
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .undelegate_claim
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .undelegate_claim
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .swap
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .swap
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .swap
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .swap_claim
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .swap_claim
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .swap_claim
                        .as_slice(),
                )?,
            },
            Phase2RawContribution {
                parent: ContributionHash::try_from(
                    value
                        .parent_hashes
                        .as_ref()
                        .ok_or(anyhow!("no parent hashes"))?
                        .nullifer_derivation_crs
                        .as_slice(),
                )?,
                new_elements: Phase2RawCRSElements::deserialize_compressed(
                    value
                        .updated
                        .as_ref()
                        .ok_or(anyhow!("no updated"))?
                        .nullifer_derivation_crs
                        .as_slice(),
                )?,
                linking_proof: DLogProof::deserialize_compressed(
                    value
                        .update_proofs
                        .as_ref()
                        .ok_or(anyhow!("no update proofs"))?
                        .nullifer_derivation_crs
                        .as_slice(),
                )?,
            },
        ]))
    }
}

impl Phase2RawCeremonyContribution {
    /// Validate that this contribution is internally consistent.
    ///
    /// This doesn't check that it's connected to the right parent though, which is an additional
    /// step you want to do.
    pub fn validate(
        self,
        rng: &mut impl CryptoRngCore,
        root: &Phase2CeremonyCRS,
    ) -> Option<Phase2CeremonyContribution> {
        // Not happy at the need to copy here, but this avoids the need for a default impl or
        // unsafe.
        for (x, root_i) in self.0.iter().cloned().zip(root.0.iter()) {
            x.validate(rng, root_i)?;
        }
        Some(self.assume_valid())
    }

    /// Skip validation, performing the conversion anyways.
    ///
    /// Useful when parsing known good data.
    pub fn assume_valid(self) -> Phase2CeremonyContribution {
        // This avoids a copy, and will break if we change the size:
        match self.0 {
            [x0, x1, x2, x3, x4, x5, x6] => Phase2CeremonyContribution([
                x0.assume_valid(),
                x1.assume_valid(),
                x2.assume_valid(),
                x3.assume_valid(),
                x4.assume_valid(),
                x5.assume_valid(),
                x6.assume_valid(),
            ]),
        }
    }
}

/// Holds all of the phase2 contributions in a single package.
#[derive(Clone, Debug)]
pub struct Phase2CeremonyContribution([Phase2Contribution; NUM_CIRCUITS]);

impl From<Phase2CeremonyContribution> for Phase2RawCeremonyContribution {
    fn from(value: Phase2CeremonyContribution) -> Self {
        let out: [Phase2RawContribution; NUM_CIRCUITS] =
            array::from_fn(|i| Phase2RawContribution::from(value.0[i].clone()));
        Self(out)
    }
}

impl TryFrom<Phase2CeremonyContribution> for pb::participate_request::Contribution {
    type Error = anyhow::Error;

    fn try_from(data: Phase2CeremonyContribution) -> Result<pb::participate_request::Contribution> {
        Phase2RawCeremonyContribution::from(data).try_into()
    }
}

impl Phase2CeremonyContribution {
    /// Get the new elements contained in this contribution
    pub fn new_elements(&self) -> Phase2CeremonyCRS {
        Phase2CeremonyCRS(array::from_fn(|i| self.0[i].new_elements.clone()))
    }

    /// Check that this contribution is linked to some specific parent elements.
    #[must_use]
    pub fn is_linked_to(&self, parent: &Phase2CeremonyCRS) -> bool {
        self.0
            .iter()
            .zip(parent.0.iter())
            .all(|(x, y)| x.is_linked_to(y))
    }

    pub fn make<R: CryptoRngCore>(rng: &mut R, old: &Phase2CeremonyCRS) -> Self {
        Self(array::from_fn(|i| {
            Phase2Contribution::make(rng, ContributionHash::dummy(), &old.0.as_ref()[i])
        }))
    }
}

// TODO: Make the phase 1 and phase 2 functionality generic

/// Holds all of the CRS elements for phase1 in one struct, before validation.
#[derive(Clone, Debug)]
pub struct Phase1RawCeremonyCRS([Phase1RawCRSElements; NUM_CIRCUITS]);

impl Phase1RawCeremonyCRS {
    /// Skip validation, performing the conversion anyways.
    ///
    /// Useful when parsing known good data.
    pub fn assume_valid(self) -> Phase1CeremonyCRS {
        match self.0 {
            [x0, x1, x2, x3, x4, x5, x6] => Phase1CeremonyCRS([
                x0.assume_valid(),
                x1.assume_valid(),
                x2.assume_valid(),
                x3.assume_valid(),
                x4.assume_valid(),
                x5.assume_valid(),
                x6.assume_valid(),
            ]),
        }
    }
}

impl TryInto<pb::CeremonyCrs> for Phase1CeremonyCRS {
    type Error = anyhow::Error;

    fn try_into(self) -> Result<pb::CeremonyCrs> {
        Ok(pb::CeremonyCrs {
            spend: to_bytes(&self.0[0])?,
            output: to_bytes(&self.0[1])?,
            delegator_vote: to_bytes(&self.0[2])?,
            undelegate_claim: to_bytes(&self.0[3])?,
            swap: to_bytes(&self.0[4])?,
            swap_claim: to_bytes(&self.0[5])?,
            nullifer_derivation_crs: to_bytes(&self.0[6])?,
        })
    }
}

impl TryFrom<pb::CeremonyCrs> for Phase1RawCeremonyCRS {
    type Error = anyhow::Error;

    fn try_from(value: pb::CeremonyCrs) -> std::result::Result<Self, Self::Error> {
        Ok(Self([
            Phase1RawCRSElements::deserialize_compressed(value.spend.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.output.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.delegator_vote.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.undelegate_claim.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.swap.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.swap_claim.as_slice())?,
            Phase1RawCRSElements::deserialize_compressed(value.nullifer_derivation_crs.as_slice())?,
        ]))
    }
}

/// Holds all of the CRS elements for phase1 in one struct.
#[derive(Clone, Debug)]
pub struct Phase1CeremonyCRS([Phase1CRSElements; NUM_CIRCUITS]);

impl From<Phase1CeremonyCRS> for Phase1RawCeremonyCRS {
    fn from(value: Phase1CeremonyCRS) -> Self {
        Self(array::from_fn(|i| value.0[i].raw.clone()))
    }
}

// impl TryFrom<Phase1CeremonyCRS> for pb::CeremonyCrs {
//     type Error = anyhow::Error;

//     fn try_from(data: Phase1CeremonyCRS) -> Result<pb::CeremonyCrs> {
//         Phase1RawCeremonyCRS::from(data).try_into()
//     }
// }

impl Phase1CeremonyCRS {
    pub fn root() -> Result<Self> {
        let [c0, c1, c2, c3, c4, c5, c6] = circuits();
        println!("GENERATED CIRCUITS");
        Ok(Self([
            Phase1CRSElements::root(circuit_degree(&c0)?),
            Phase1CRSElements::root(circuit_degree(&c1)?),
            Phase1CRSElements::root(circuit_degree(&c2)?),
            Phase1CRSElements::root(circuit_degree(&c3)?),
            Phase1CRSElements::root(circuit_degree(&c4)?),
            Phase1CRSElements::root(circuit_degree(&c5)?),
            Phase1CRSElements::root(circuit_degree(&c6)?),
        ]))
    }
}
