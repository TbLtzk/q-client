package common

import "testing"

func TestRootListSigningPayloadV1_DeterministicAndSensitiveFields(t *testing.T) {
	base := RootList{
		Timestamp: 1715072400,
		Hash:      HexToHash("0x33f80a7cc3006cbd2f5fa4edca320042695f6f4cbf15daea73481f61ebf563f1"),
		Nodes: []Address{
			HexToAddress("0xEB3B90FD1862B10D14D71881B32D80E530AD394B"),
			HexToAddress("0x01FDCC35858C76C6ECD459DA0174116FB5A4BFF7"),
			HexToAddress("0xA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B"),
		},
	}
	sameDataDifferentOrder := RootList{
		Timestamp: base.Timestamp,
		Hash:      base.Hash,
		Nodes: []Address{
			base.Nodes[2],
			base.Nodes[1],
			base.Nodes[0],
			base.Nodes[2], // duplicate should not change payload
		},
	}

	baseHash := NewRootListSigningPayloadV1(999, base).Hash()
	sameDataHash := NewRootListSigningPayloadV1(999, sameDataDifferentOrder).Hash()
	if baseHash != sameDataHash {
		t.Fatalf("expected stable root-list payload hash for same data in different order")
	}

	changedChainIDHash := NewRootListSigningPayloadV1(1000, base).Hash()
	if baseHash == changedChainIDHash {
		t.Fatalf("expected hash change when chain ID changes")
	}

	changedTimestamp := base
	changedTimestamp.Timestamp = base.Timestamp + 1
	changedTimestampHash := NewRootListSigningPayloadV1(999, changedTimestamp).Hash()
	if baseHash == changedTimestampHash {
		t.Fatalf("expected hash change when timestamp changes")
	}

	changedPayloadHash := base
	changedPayloadHash.Hash = HexToHash("0x9b5f0c4ecf94af3142ac4f9de24fdb7b7df77ef6e2da94895cd6d95a6ca3d7dc")
	changedPayloadHashHash := NewRootListSigningPayloadV1(999, changedPayloadHash).Hash()
	if baseHash == changedPayloadHashHash {
		t.Fatalf("expected hash change when payload hash changes")
	}
}

func TestExclusionListSigningPayloadV1_DeterministicAndSensitiveFields(t *testing.T) {
	base := ValidatorExclusionList{
		Timestamp: 1715072400,
		Hash:      HexToHash("0x4209a8e2f099937261a9f4ef7ce7c6b2f3f16f79a0fb92ceb95f9fda4ea9d2ee"),
		Validators: []ExcludedValidator{
			{Address: HexToAddress("0xC5E8F30E914F3F9D88F2DC72F4C2D4E6FB08D5E2"), Block: 100, EndBlock: 120},
			{Address: HexToAddress("0x10AC88611B540D6E7725198EDB6B9B723E4A6E6D"), Block: 77, EndBlock: 0},
			{Address: HexToAddress("0x10AC88611B540D6E7725198EDB6B9B723E4A6E6D"), Block: 88, EndBlock: 91},
		},
	}
	sameDataDifferentOrder := ValidatorExclusionList{
		Timestamp: base.Timestamp,
		Hash:      base.Hash,
		Validators: []ExcludedValidator{
			base.Validators[2],
			base.Validators[0],
			base.Validators[1],
		},
	}

	baseHash := NewExclusionListSigningPayloadV1(999, base).Hash()
	sameDataHash := NewExclusionListSigningPayloadV1(999, sameDataDifferentOrder).Hash()
	if baseHash != sameDataHash {
		t.Fatalf("expected stable exclusion-list payload hash for same data in different order")
	}

	changedChainIDHash := NewExclusionListSigningPayloadV1(1000, base).Hash()
	if baseHash == changedChainIDHash {
		t.Fatalf("expected hash change when chain ID changes")
	}

	changedTimestamp := base
	changedTimestamp.Timestamp = base.Timestamp + 1
	changedTimestampHash := NewExclusionListSigningPayloadV1(999, changedTimestamp).Hash()
	if baseHash == changedTimestampHash {
		t.Fatalf("expected hash change when timestamp changes")
	}

	changedPayloadHash := base
	changedPayloadHash.Hash = HexToHash("0x7243f2609c76b876eff8a434522f810e0ca4f5912f6ea0f57c15d45f6bd8e61e")
	changedPayloadHashHash := NewExclusionListSigningPayloadV1(999, changedPayloadHash).Hash()
	if baseHash == changedPayloadHashHash {
		t.Fatalf("expected hash change when payload hash changes")
	}
}

func TestGovernanceSigningPayloads_DoNotCollideAcrossProposalTypes(t *testing.T) {
	hash := HexToHash("0x67b6914c8f2d6642eb5a2488b3fd3e6fec57f6f2de6f2c9e1fdbf0ef8e16b3fa")
	timestamp := uint64(1715072400)
	chainID := uint64(1337)
	addr := HexToAddress("0xA94F5374FCE5EDBC8E2A8697C15331677E6EBF0B")

	rootPayload := NewRootListSigningPayloadV1(chainID, RootList{
		Timestamp: timestamp,
		Hash:      hash,
		Nodes:     []Address{addr},
	})
	exclusionPayload := NewExclusionListSigningPayloadV1(chainID, ValidatorExclusionList{
		Timestamp: timestamp,
		Hash:      hash,
		Validators: []ExcludedValidator{
			{Address: addr, Block: 1, EndBlock: 0},
		},
	})

	if rootPayload.Hash() == exclusionPayload.Hash() {
		t.Fatalf("expected payload hashes to differ across proposal types")
	}
}
