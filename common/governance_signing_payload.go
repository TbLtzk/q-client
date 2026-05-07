package common

import (
	"bytes"
	"encoding/binary"
	"sort"

	"golang.org/x/crypto/sha3"
)

const (
	// GovernanceSigningDomain is the EIP-712 domain name used by q-client for L0 governance typed payloads.
	GovernanceSigningDomain = "QGOV-L0-Governance"
	// GovernanceSigningVersionV1 is the first version of the typed signing payload.
	GovernanceSigningVersionV1 = "1"
	// GovernanceEIP712DomainTypeName is the expected EIP-712 domain struct name.
	GovernanceEIP712DomainTypeName = "EIP712Domain"
	// GovernanceEIP712RootListTypeName is the expected EIP-712 primary type for root-list proposals.
	GovernanceEIP712RootListTypeName = "QGOVL0RootListProposal"
	// GovernanceEIP712ExclusionListTypeName is the expected EIP-712 primary type for exclusion-list proposals.
	GovernanceEIP712ExclusionListTypeName = "QGOVL0ExclusionListProposal"
	// GovernanceProposalTypeRootList is the proposal type value for root-list updates.
	GovernanceProposalTypeRootList = "rootList"
	// GovernanceProposalTypeExclusionList is the proposal type value for exclusion-list updates.
	GovernanceProposalTypeExclusionList = "exclusionList"
)

type GovernanceSigningMetadata struct {
	Domain       string `json:"domain"`
	Version      string `json:"version"`
	ChainID      uint64 `json:"chainId"`
	ProposalType string `json:"proposalType"`
	Timestamp    uint64 `json:"timestamp"`
	PayloadHash  Hash   `json:"payloadHash"`
}

type RootListSigningPayload struct {
	Metadata GovernanceSigningMetadata `json:"metadata"`
	Nodes    []Address                 `json:"nodes"`
}

type ExclusionListSigningPayload struct {
	Metadata   GovernanceSigningMetadata `json:"metadata"`
	Validators []ExcludedValidator       `json:"validators"`
}

func NewRootListSigningPayloadV1(chainID uint64, list RootList) RootListSigningPayload {
	return RootListSigningPayload{
		Metadata: GovernanceSigningMetadata{
			Domain:       GovernanceSigningDomain,
			Version:      GovernanceSigningVersionV1,
			ChainID:      chainID,
			ProposalType: GovernanceProposalTypeRootList,
			Timestamp:    list.Timestamp,
			PayloadHash:  list.Hash,
		},
		Nodes: canonicalRootNodes(list.Nodes),
	}
}

func NewExclusionListSigningPayloadV1(chainID uint64, list ValidatorExclusionList) ExclusionListSigningPayload {
	return ExclusionListSigningPayload{
		Metadata: GovernanceSigningMetadata{
			Domain:       GovernanceSigningDomain,
			Version:      GovernanceSigningVersionV1,
			ChainID:      chainID,
			ProposalType: GovernanceProposalTypeExclusionList,
			Timestamp:    list.Timestamp,
			PayloadHash:  list.Hash,
		},
		Validators: canonicalExcludedValidators(list.Validators),
	}
}

func (p RootListSigningPayload) Hash() Hash {
	return keccakHash(p.canonicalBytes())
}

func (p ExclusionListSigningPayload) Hash() Hash {
	return keccakHash(p.canonicalBytes())
}

func (p RootListSigningPayload) canonicalBytes() []byte {
	var buf bytes.Buffer
	writeMetadata(&buf, p.Metadata)
	writeUint32(&buf, uint32(len(p.Nodes)))
	for _, node := range p.Nodes {
		buf.Write(node.Bytes())
	}
	return buf.Bytes()
}

func (p ExclusionListSigningPayload) canonicalBytes() []byte {
	var buf bytes.Buffer
	writeMetadata(&buf, p.Metadata)
	writeUint32(&buf, uint32(len(p.Validators)))
	for _, validator := range p.Validators {
		buf.Write(validator.Address.Bytes())
		writeUint64(&buf, validator.Block)
		writeUint64(&buf, validator.EndBlock)
	}
	return buf.Bytes()
}

func writeMetadata(buf *bytes.Buffer, meta GovernanceSigningMetadata) {
	writeString(buf, meta.Domain)
	writeString(buf, meta.Version)
	writeUint64(buf, meta.ChainID)
	writeString(buf, meta.ProposalType)
	writeUint64(buf, meta.Timestamp)
	buf.Write(meta.PayloadHash.Bytes())
}

func writeString(buf *bytes.Buffer, value string) {
	writeUint32(buf, uint32(len(value)))
	buf.WriteString(value)
}

func writeUint32(buf *bytes.Buffer, value uint32) {
	_ = binary.Write(buf, binary.BigEndian, value)
}

func writeUint64(buf *bytes.Buffer, value uint64) {
	_ = binary.Write(buf, binary.BigEndian, value)
}

func canonicalRootNodes(nodes []Address) []Address {
	unique := make(map[Address]struct{}, len(nodes))
	out := make([]Address, 0, len(nodes))
	for _, node := range nodes {
		if _, exists := unique[node]; exists {
			continue
		}
		unique[node] = struct{}{}
		out = append(out, node)
	}
	sort.SliceStable(out, func(i, j int) bool {
		return bytes.Compare(out[i].Bytes(), out[j].Bytes()) > 0
	})
	return out
}

func canonicalExcludedValidators(validators []ExcludedValidator) []ExcludedValidator {
	out := make([]ExcludedValidator, len(validators))
	copy(out, validators)
	sort.SliceStable(out, func(i, j int) bool {
		left, right := out[i], out[j]
		if left.Address != right.Address {
			return bytes.Compare(left.Address.Bytes(), right.Address.Bytes()) < 0
		}
		if left.Block != right.Block {
			return left.Block < right.Block
		}
		return left.EndBlock < right.EndBlock
	})
	return out
}

func keccakHash(data []byte) Hash {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	return BytesToHash(hasher.Sum(nil))
}
