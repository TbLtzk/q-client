package governance

import (
	"gitlab.com/q-dev/q-client/common"
)

func (a *GovernancePublicAPI) L0GovernanceCapabilities() (common.L0GovernanceCapabilities, error) {
	rm := a.gov.RootManager
	externalEnabled := !rm.DisablePublicSubmission

	cap := common.L0GovernanceCapabilities{
		SchemaVersion:                        common.L0GovernanceCapabilitiesSchemaVersion,
		NetworkID:                            rm.networkId,
		ExternalSubmissionEnabled:            externalEnabled,
		RootList:                             proposalCapabilities(externalEnabled),
		ExclusionList:                        proposalCapabilities(externalEnabled),
		ProposalStatus:                       true,
		ProposalStatusRequiredSigningAddress: true,
		AliasSigningRequired:                 rm.hasDistinctRootAliases(),
		SigningPayload: common.L0GovernanceSigningPayloadCapabilities{
			Domain:                   common.GovernanceSigningDomain,
			Version:                  common.GovernanceSigningVersionV1,
			VerifyingContract:        governanceEIP712ZeroContract,
			SigningPayloadWithDigest: true,
		},
	}

	if !externalEnabled {
		cap.DisableReason = errPublicGovernanceSubmissionDisabled.Error()
	}

	if advertisesQgovTypedRelay() {
		v := uint(qgov6)
		cap.QgovTypedRelayVersion = &v
	}

	return cap, nil
}

func proposalCapabilities(externalSubmissionEnabled bool) common.L0GovernanceProposalCapabilities {
	schemes := []string{
		common.GovernanceSigningSchemeRawListHash,
		common.GovernanceSigningSchemeEIP712V1,
	}
	payloadVersions := []string{common.GovernanceSigningPayloadVersionEIP712V1}

	return common.L0GovernanceProposalCapabilities{
		RawSubmit:              externalSubmissionEnabled,
		TypedSubmit:            externalSubmissionEnabled,
		SigningSchemes:         schemes,
		SigningPayloadVersions: payloadVersions,
	}
}

func advertisesQgovTypedRelay() bool {
	for _, version := range ProtocolVersions {
		if version == qgov6 {
			return true
		}
	}
	return false
}

func (s *RootManager) hasDistinctRootAliases() bool {
	active := s.getActiveRootSet(true)
	if active == nil {
		return false
	}
	for root, alias := range active.aliases {
		if root != alias {
			return true
		}
	}
	return false
}
