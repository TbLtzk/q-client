package governance

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"gitlab.com/q-dev/q-client/common"
)

func TestRootSetFromDBJSON_legacyBase64Signatures(t *testing.T) {
	sig := make([]byte, 65)
	sig[64] = 1
	unsigned := common.RootList{
		Timestamp: 100,
		Nodes:     []common.Address{common.HexToAddress("0x01")},
	}
	set, err := newRootSet(&unsigned)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}
	legacy, err := json.Marshal(rootListDB{
		Timestamp:  unsigned.Timestamp,
		Nodes:      unsigned.Nodes,
		Hash:       set.hash,
		Signatures: [][]byte{sig},
	})
	if err != nil {
		t.Fatalf("marshal legacy: %v", err)
	}

	var probe struct {
		Signatures []string `json:"signatures"`
	}
	if err := json.Unmarshal(legacy, &probe); err != nil {
		t.Fatalf("probe: %v", err)
	}
	if _, err := base64.StdEncoding.DecodeString(probe.Signatures[0]); err != nil {
		t.Fatalf("expected base64 signature element, got %q", probe.Signatures[0])
	}

	loaded, err := rootSetFromDBJSON(legacy)
	if err != nil || loaded == nil {
		t.Fatalf("rootSetFromDBJSON: %v", err)
	}
	if loaded.hash != set.hash || len(loaded.rootAddresses) != 1 {
		t.Fatalf("unexpected set %+v", loaded)
	}
}

func TestRootSetDBRoundTrip_usesLegacyEncoding(t *testing.T) {
	unsigned := common.RootList{
		Timestamp: 200,
		Nodes:     []common.Address{common.HexToAddress("0x03")},
	}
	set, err := newRootSet(&unsigned)
	if err != nil || set == nil {
		t.Fatalf("newRootSet: %v", err)
	}
	set.signers = map[common.Address][]byte{
		common.HexToAddress("0x04"): make([]byte, 65),
	}

	raw, err := json.Marshal(rootSetToDBList(set))
	if err != nil {
		t.Fatalf("marshal db: %v", err)
	}

	loaded, err := rootSetFromDBJSON(raw)
	if err != nil || loaded == nil {
		t.Fatalf("load: %v", err)
	}
	if loaded.timestamp != set.timestamp {
		t.Fatalf("timestamp mismatch")
	}
}

func TestExclusionSetFromDBJSON_legacyBase64Signatures(t *testing.T) {
	sig := make([]byte, 65)
	unsigned := common.ValidatorExclusionList{
		Timestamp:  1,
		Validators: []common.ExcludedValidator{{Address: common.HexToAddress("0x05"), Block: 1}},
	}
	base, err := newExclusionSet(&unsigned)
	if err != nil || base == nil {
		t.Fatalf("newExclusionSet: %v", err)
	}
	legacy, err := json.Marshal(exclusionListDB{
		Timestamp:  unsigned.Timestamp,
		Validators: unsigned.Validators,
		Hash:       base.hash,
		Signatures: [][]byte{sig},
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if _, err := exclusionSetFromDBJSON(legacy); err != nil {
		t.Fatalf("exclusionSetFromDBJSON: %v", err)
	}
}
