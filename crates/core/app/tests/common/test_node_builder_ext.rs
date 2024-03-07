use {
    penumbra_genesis::AppState,
    penumbra_mock_consensus::builder::Builder,
    penumbra_proto::{
        core::keys::v1::{GovernanceKey, IdentityKey},
        penumbra::core::component::stake::v1::Validator as PenumbraValidator,
    },
    tap::Tap,
};

/// Penumbra-specific extensions to the mock consensus builder.
pub trait BuilderExt: Sized {
    /// The error thrown by [`with_penumbra_auto_app_state`]
    type Error;
    /// Add the provided Penumbra [`AppState`] to the builder.
    ///
    /// This will inject any configured validators into the state before serializing it into bytes.
    fn with_penumbra_auto_app_state(self, app_state: AppState) -> Result<Self, Self::Error>;
}

impl BuilderExt for Builder {
    type Error = anyhow::Error;
    fn with_penumbra_auto_app_state(self, app_state: AppState) -> Result<Self, Self::Error> {
        // Generate a penumbra validator using the test node's consensus keys (if they exist).
        // Eventually, we may wish to generate and inject additional definitions, but only a single
        // validator is supported for now.
        let app_state = match self
            .keys
            .as_ref()
            .map(|keys| generate_penumbra_validator(keys.consensus_verification_key))
            .inspect(log_validator)
            .map(std::iter::once)
        {
            Some(validator) => app_state_with_validators(app_state, validator)?,
            None => app_state,
        };

        // Serialize the app state into bytes, and add it to the builder.
        serde_json::to_vec(&app_state)
            .map_err(Self::Error::from)
            .map(|s| self.app_state(s))
    }
}

/// Injects the given collection of [`Validator`s][PenumbraValidator] into the app state.
fn app_state_with_validators<V>(
    app_state: AppState,
    validators: V,
) -> Result<AppState, anyhow::Error>
where
    V: IntoIterator<Item = PenumbraValidator>,
{
    use AppState::{Checkpoint, Content};
    match app_state {
        Checkpoint(_) => anyhow::bail!("checkpoint app state isn't supported"),
        Content(mut content) => {
            // Inject the builder's validators into the staking component's genesis state.
            std::mem::replace(
                &mut content.stake_content.validators,
                validators.into_iter().collect(),
            )
            .tap(|overwritten| {
                // Log a warning if this overwrote any validators already in the app state.
                if !overwritten.is_empty() {
                    tracing::warn!(
                        ?overwritten,
                        "`with_penumbra_auto_app_state` overwrote validators in the given AppState"
                    )
                }
            });
            Ok(Content(content))
        }
    }
}

/// Generates a [`Validator`][PenumbraValidator] given a set of consensus [`Keys`].
fn generate_penumbra_validator(
    consensus_verification_key: ed25519_consensus::VerificationKey,
) -> PenumbraValidator {
    /// A temporary stub for validator keys.
    ///
    /// An invalid key is intentionally provided here, until we have test coverage exercising the
    /// use of these keys. Once we need it we will:
    /// - generate a random signing key
    /// - get its verification key
    /// - use that for the identity key
    /// - throw the signing key away
    ///
    /// NB: for now, we will use the same key for governance. See the documentation of
    /// `GovernanceKey` for more information about cold storage of validator keys.
    const INVALID_KEY_BYTES: [u8; 32] = [0; 32];

    PenumbraValidator {
        identity_key: Some(IdentityKey {
            ik: INVALID_KEY_BYTES.to_vec().clone(),
        }),
        governance_key: Some(GovernanceKey {
            gk: INVALID_KEY_BYTES.to_vec().clone(),
        }),
        consensus_key: consensus_verification_key.as_bytes().to_vec(),
        enabled: true,
        sequence_number: 0,
        name: String::default(),
        website: String::default(),
        description: String::default(),
        funding_streams: Vec::default(),
    }
}

fn log_validator(
    PenumbraValidator {
        name,
        enabled,
        sequence_number,
        ..
    }: &PenumbraValidator,
) {
    tracing::trace!(
        %name,
        %enabled,
        %sequence_number,
        "injecting validator into app state"
    )
}
