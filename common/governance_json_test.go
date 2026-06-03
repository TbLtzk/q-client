package common

import (
	"encoding/json"
	"testing"

	"gitlab.com/q-dev/q-client/common/hexutil"
)

func TestRootListSignaturesJSON_hexOnly(t *testing.T) {
	sigHex := hexutil.Encode(make([]byte, 65))
	payload := `{
		"timestamp": 1715072400,
		"nodes": ["0xEB3B90FD1862B10D14D71881B32D80E530AD394B"],
		"hash": "0x0000000000000000000000000000000000000000000000000000000000000001",
		"signatures": ["` + sigHex + `"]
	}`

	var list RootList
	if err := json.Unmarshal([]byte(payload), &list); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(list.Signatures) != 1 {
		t.Fatalf("signatures len %d", len(list.Signatures))
	}
	if len(list.Signatures[0]) != 65 {
		t.Fatalf("signature bytes len %d", len(list.Signatures[0]))
	}

	out, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !json.Valid(out) {
		t.Fatalf("invalid json: %s", out)
	}
	var round RootList
	if err := json.Unmarshal(out, &round); err != nil {
		t.Fatalf("round-trip: %v", err)
	}
	if hexutil.Encode(round.Signatures[0]) != hexutil.Encode(list.Signatures[0]) {
		t.Fatalf("signature round-trip mismatch")
	}
}

func TestRootListSignaturesJSON_rejectsBase64(t *testing.T) {
	const payload = `{
		"timestamp": 1,
		"nodes": [],
		"hash": "0x0000000000000000000000000000000000000000000000000000000000000001",
		"signatures": ["AQID"]
	}`

	var list RootList
	if err := json.Unmarshal([]byte(payload), &list); err == nil {
		t.Fatalf("expected base64 signature to fail, got %d signatures", len(list.Signatures))
	}
}
