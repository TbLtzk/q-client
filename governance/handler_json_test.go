package governance

import (
	"encoding/json"
	"testing"
	"time"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/crypto"
)

func TestSubmitTypedSignedRootList_hexJSONWire(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := gov.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}

	list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
	unsigned := list
	unsigned.Signatures = nil
	sig := signRootListEIP712(t, rm.networkId, list, rm.aliasPrivateKeys[0])
	list.Signatures = common.HexSignaturesFromBytes([][]byte{sig})

	raw, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !json.Valid(raw) {
		t.Fatalf("invalid json")
	}

	var decoded common.RootList
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(decoded.Signatures) != 1 {
		t.Fatalf("signatures len %d", len(decoded.Signatures))
	}

	hash, err := NewGovernancePublicAPI(gov).SubmitTypedSignedRootList(decoded)
	if err != nil {
		t.Fatalf("SubmitTypedSignedRootList: %v", err)
	}
	if hash != list.Hash {
		t.Fatalf("hash %s want %s", hash.Hex(), list.Hash.Hex())
	}
}

func TestSubmitTypedSignedRootList_walletRecoveryByte(t *testing.T) {
	rm := newTestRootManager(t, false, true)
	gov, err := New(rm.RootManager, tmpDirName(t))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if err := gov.Start(); err != nil {
		t.Fatalf("Start: %v", err)
	}

	list := randomRootList(t, rm, time.Now().Add(5*time.Minute).Unix(), 10, 0, true)
	sig := signRootListEIP712(t, rm.networkId, list, rm.aliasPrivateKeys[0])
	walletSig := append([]byte(nil), sig...)
	walletSig[crypto.RecoveryIDOffset] += 27
	list.Signatures = common.HexSignaturesFromBytes([][]byte{walletSig})

	if _, err := NewGovernancePublicAPI(gov).SubmitTypedSignedRootList(list); err != nil {
		t.Fatalf("SubmitTypedSignedRootList with v+27: %v", err)
	}
}

func TestRootListJSON_signatureMarshalsAsHex(t *testing.T) {
	sig := hexutil.Bytes(make([]byte, 65))
	raw, err := json.Marshal(common.RootList{Signatures: []hexutil.Bytes{sig}})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !json.Valid(raw) {
		t.Fatal("invalid json")
	}
	var probe struct {
		Signatures []string `json:"signatures"`
	}
	if err := json.Unmarshal(raw, &probe); err != nil {
		t.Fatalf("unmarshal probe: %v", err)
	}
	if len(probe.Signatures) != 1 || len(probe.Signatures[0]) < 4 || probe.Signatures[0][:2] != "0x" {
		t.Fatalf("expected 0x hex signature, got %q", probe.Signatures)
	}
}
