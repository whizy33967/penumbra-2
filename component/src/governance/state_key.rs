use penumbra_crypto::{stake::IdentityKey, Nullifier};

pub fn next_proposal_id() -> &'static str {
    "governance/next_proposal_id"
}

pub fn proposal_definition(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/data")
}

pub fn proposal_state(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/state")
}

pub fn proposal_deposit_amount(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/deposit_amount")
}

pub fn proposal_voting_start(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/voting_start")
}

pub fn proposal_voting_start_position(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/voting_start_position")
}

pub fn proposal_voting_end(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/voting_end")
}

pub fn unfinished_proposals() -> &'static str {
    "governance/unfinished_proposals"
}

pub fn voting_validators_list(proposal_id: u64) -> String {
    format!("governance/proposal/{proposal_id}/validator_vote/")
}

pub fn per_proposal_voted_nullifier_lookup(proposal_id: u64, nullifier: &Nullifier) -> String {
    format!("governance/proposal/{proposal_id}/voted_nullifiers/{nullifier}")
}

pub fn rate_data_at_proposal_start(proposal_id: u64, identity_key: IdentityKey) -> String {
    format!("governance/proposal/{proposal_id}/rate_data_at_start/{identity_key}")
}

pub fn all_rate_data_at_proposal_start(proposal_id: u64) -> String {
    // Note: this has to be the prefix of the `rate_data_at_proposal_start` function above.
    format!("governance/proposal/{proposal_id}/rate_data_at_start/")
}

pub fn voting_power_at_proposal_start(proposal_id: u64, identity_key: IdentityKey) -> String {
    format!("governance/proposal/{proposal_id}/voting_power_at_start/{identity_key}")
}

pub fn all_voting_power_at_proposal_start(proposal_id: u64) -> String {
    // Note: this has to be the prefix of the `voting_power_at_proposal_start` function above.
    format!("governance/proposal/{proposal_id}/voting_power_at_start/")
}

pub fn validator_vote(proposal_id: u64, identity_key: IdentityKey) -> String {
    format!("governance/proposal/{proposal_id}/validator_vote/{identity_key}")
}

pub fn all_validator_votes(proposal_id: u64) -> String {
    // Note: this has to be the prefix of the `validator_vote` function above.
    format!("governance/proposal/{proposal_id}/validator_vote/")
}

pub fn untallied_delegator_vote(
    proposal_id: u64,
    identity_key: IdentityKey,
    nullifier: &Nullifier,
) -> String {
    format!("governance/proposal/{proposal_id}/untallied_delegator_vote/{identity_key}/{nullifier}")
}

pub fn all_untallied_delegator_votes(proposal_id: u64) -> String {
    // Note: this has to be the prefix of the `untallied_delegator_vote` function above.
    format!("governance/proposal/{proposal_id}/untallied_delegator_vote/")
}

pub fn tallied_delegator_votes(proposal_id: u64, identity_key: IdentityKey) -> String {
    format!("governance/proposal/{proposal_id}/tallied_delegator_votes/{identity_key}")
}

pub fn all_tallied_delegator_votes(proposal_id: u64) -> String {
    // Note: this has to be the prefix of the `tallied_delegator_votes` function above.
    format!("governance/proposal/{proposal_id}/tallied_delegator_votes/")
}
