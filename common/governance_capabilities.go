package common

const (
	// L0GovernanceCapabilitiesSchemaVersion is the schema version of govPub_l0GovernanceCapabilities responses.
	L0GovernanceCapabilitiesSchemaVersion = 1

	// GovernanceSigningSchemeRawListHash is legacy ECDSA over the deterministic list hash.
	GovernanceSigningSchemeRawListHash = "raw-list-hash"
	// GovernanceSigningSchemeEIP712V1 is EIP-712 typed data (eth_signTypedData_v4).
	GovernanceSigningSchemeEIP712V1 = "eip712-v1"

	// GovernanceSigningPayloadVersionEIP712V1 is the signing-payload / typed-data version label.
	GovernanceSigningPayloadVersionEIP712V1 = "eip712-v1"
)

// L0GovernanceCapabilities is returned by govPub_l0GovernanceCapabilities.
type L0GovernanceCapabilities struct {
	SchemaVersion             uint   `json:"schemaVersion"`
	NetworkID                 uint64 `json:"networkId"`
	ExternalSubmissionEnabled bool   `json:"externalSubmissionEnabled"`
	DisableReason             string `json:"disableReason,omitempty"`

	// QgovTypedRelayVersion is set when the node registers qgov/6 typed attestation relay (separate from RPC flags).
	QgovTypedRelayVersion *uint `json:"qgovTypedRelayVersion,omitempty"`

	RootList      L0GovernanceProposalCapabilities `json:"rootList"`
	ExclusionList L0GovernanceProposalCapabilities `json:"exclusionList"`

	ProposalStatus                       bool `json:"proposalStatus"`
	ProposalStatusRequiredSigningAddress bool `json:"proposalStatusRequiredSigningAddress"`

	// AliasSigningRequired is true when the active root set has at least one root with a distinct alias address.
	AliasSigningRequired bool `json:"aliasSigningRequired"`

	SigningPayload L0GovernanceSigningPayloadCapabilities `json:"signingPayload"`
}

// L0GovernanceProposalCapabilities describes submission and verification options for one proposal kind.
type L0GovernanceProposalCapabilities struct {
	RawSubmit              bool     `json:"rawSubmit"`
	TypedSubmit            bool     `json:"typedSubmit"`
	SigningSchemes         []string `json:"signingSchemes"`
	SigningPayloadVersions []string `json:"signingPayloadVersions"`
}

// L0GovernanceSigningPayloadCapabilities documents EIP-712 domain metadata for wallet signing.
type L0GovernanceSigningPayloadCapabilities struct {
	Domain                   string `json:"domain"`
	Version                  string `json:"version"`
	VerifyingContract        string `json:"verifyingContract"`
	SigningPayloadWithDigest bool   `json:"signingPayloadWithDigest"`
}
