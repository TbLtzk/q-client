package governance

import (
	"crypto/ecdsa"
	"testing"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/signer/core/apitypes"
)

func signRootListEIP712(t *testing.T, networkID uint64, list common.RootList, key *ecdsa.PrivateKey) []byte {
	t.Helper()

	unsigned := list
	unsigned.Signatures = nil
	set, err := newRootSet(&unsigned)
	if err != nil {
		t.Fatalf("newRootSet: %v", err)
	}
	if set == nil {
		t.Fatal("newRootSet returned nil")
	}

	digest, err := rootListEIP712Digest(networkID, set.timestamp, set.hash, set.rootAddresses)
	if err != nil {
		t.Fatalf("rootListEIP712Digest: %v", err)
	}
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		t.Fatalf("crypto.Sign: %v", err)
	}
	return sig
}

func signRootListEIP712FromTypedData(t *testing.T, td apitypes.TypedData, key *ecdsa.PrivateKey) []byte {
	t.Helper()

	digest, err := eip712Digest(td)
	if err != nil {
		t.Fatalf("eip712Digest: %v", err)
	}
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		t.Fatalf("crypto.Sign: %v", err)
	}
	return sig
}

func signExclusionListEIP712FromTypedData(t *testing.T, td apitypes.TypedData, key *ecdsa.PrivateKey) []byte {
	t.Helper()

	digest, err := eip712Digest(td)
	if err != nil {
		t.Fatalf("eip712Digest: %v", err)
	}
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		t.Fatalf("crypto.Sign: %v", err)
	}
	return sig
}

func signExclusionListEIP712(t *testing.T, networkID uint64, list common.ValidatorExclusionList, key *ecdsa.PrivateKey) []byte {
	t.Helper()

	unsigned := list
	unsigned.Signatures = nil
	set, err := newExclusionSet(&unsigned)
	if err != nil {
		t.Fatalf("newExclusionSet: %v", err)
	}
	if set == nil {
		t.Fatal("newExclusionSet returned nil")
	}

	source := set.makeList()
	source.Signatures = nil
	validators := canonicalExcludedValidatorsForEIP712(source.Validators)

	digest, err := exclusionListEIP712Digest(networkID, set.timestamp, set.hash, validators)
	if err != nil {
		t.Fatalf("exclusionListEIP712Digest: %v", err)
	}
	sig, err := crypto.Sign(digest.Bytes(), key)
	if err != nil {
		t.Fatalf("crypto.Sign: %v", err)
	}
	return sig
}

func TestRootListEIP712_DeterministicAndDistinctFromCustomDigest(t *testing.T) {
	list := common.RootList{
		Timestamp: 1715072400,
		Nodes: []common.Address{
			common.HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
			common.HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
			common.HexToAddress("0xA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
		},
	}
	unsigned := list
	unsigned.Signatures = nil
	set, err := newRootSet(&unsigned)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}
	list.Hash = set.hash

	chainID := uint64(999)
	unsigned = list
	unsigned.Signatures = nil
	set, err = newRootSet(&unsigned)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}

	d1, err := rootListEIP712Digest(chainID, set.timestamp, set.hash, set.rootAddresses)
	if err != nil {
		t.Fatalf("digest: %v", err)
	}
	d2, err := rootListEIP712Digest(chainID, set.timestamp, set.hash, set.rootAddresses)
	if err != nil {
		t.Fatalf("digest: %v", err)
	}
	if d1 != d2 {
		t.Fatalf("expected deterministic EIP-712 digest")
	}

	custom := common.NewRootListSigningPayloadV1(chainID, list).Hash()
	if d1 == custom {
		t.Fatalf("EIP-712 digest must differ from legacy custom canonical digest")
	}
}

func TestExclusionListEIP712_DistinctProposalTypes(t *testing.T) {
	chainID := uint64(1337)
	hash := common.HexToHash("0x67b6914c8f2d6642eb5a2488b3fd3e6fec57f6f2de6f2c9e1fdbf0ef8e16b3fa")
	addr := common.HexToAddress("0xA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B")
	timestamp := uint64(1715072400)

	rootDigest, err := rootListEIP712Digest(chainID, timestamp, hash, []common.Address{addr})
	if err != nil {
		t.Fatalf("root digest: %v", err)
	}
	exDigest, err := exclusionListEIP712Digest(chainID, timestamp, hash, []common.ExcludedValidator{{Address: addr, Block: 1}})
	if err != nil {
		t.Fatalf("exclusion digest: %v", err)
	}
	if rootDigest == exDigest {
		t.Fatalf("root and exclusion EIP-712 digests must not collide")
	}
}
