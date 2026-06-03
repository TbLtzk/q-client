package common

// GovernanceEIP712TypedData is the JSON shape returned by govPub signing-payload RPCs
// for eth_signTypedData_v4 (types, primaryType, domain, message).
type GovernanceEIP712TypedData struct {
	Types       map[string][]GovernanceEIP712TypeField `json:"types"`
	PrimaryType string                                 `json:"primaryType"`
	Domain      GovernanceEIP712Domain                 `json:"domain"`
	Message     map[string]interface{}                 `json:"message"`
}

// GovernanceEIP712TypeField is one field in an EIP-712 type definition.
type GovernanceEIP712TypeField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// GovernanceEIP712Domain is the EIP-712 domain separator for L0 governance.
type GovernanceEIP712Domain struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	ChainId           string `json:"chainId"`
	VerifyingContract string `json:"verifyingContract"`
}

// RootListSigningPayloadV1WithDigest bundles EIP-712 typed data and the wallet signing hash.
type RootListSigningPayloadV1WithDigest struct {
	TypedData GovernanceEIP712TypedData `json:"typedData"`
	Digest    Hash                      `json:"digest"`
}

// ExclusionListSigningPayloadV1WithDigest is the exclusion-list counterpart.
type ExclusionListSigningPayloadV1WithDigest struct {
	TypedData GovernanceEIP712TypedData `json:"typedData"`
	Digest    Hash                      `json:"digest"`
}
