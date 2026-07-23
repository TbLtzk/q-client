package params

import (
	"runtime"
	"testing"

	"gitlab.com/q-dev/q-client/rlp"
)

type defaultExtraPayload struct {
	Version uint
	Name    string
	OS      string
}

func TestMakeExtraDataDefault(t *testing.T) {
	got := MakeExtraData(nil)
	if len(got) == 0 {
		t.Fatal("default extradata is empty")
	}
	if uint64(len(got)) > MaximumExtraDataSize {
		t.Fatalf("default extradata len=%d exceeds MaximumExtraDataSize=%d", len(got), MaximumExtraDataSize)
	}

	var decoded defaultExtraPayload
	if err := rlp.DecodeBytes(got, &decoded); err != nil {
		t.Fatalf("decode default extradata: %v", err)
	}

	wantPacked := uint(QVersionMajor<<16 | QVersionMinor<<8 | QVersionPatch)
	if decoded.Version != wantPacked {
		t.Fatalf("packed version=%d want=%d", decoded.Version, wantPacked)
	}
	if decoded.Name != "QGOV Client" {
		t.Fatalf("client name=%q want %q", decoded.Name, "QGOV Client")
	}
	if decoded.OS != runtime.GOOS {
		t.Fatalf("os=%q want %q", decoded.OS, runtime.GOOS)
	}
}

func TestMakeExtraDataOverride(t *testing.T) {
	custom := []byte("custom-extra")
	got := MakeExtraData(custom)
	if string(got) != string(custom) {
		t.Fatalf("override not preserved: %q", got)
	}
}

func TestMakeExtraDataTooLong(t *testing.T) {
	tooLong := make([]byte, MaximumExtraDataSize+1)
	if got := MakeExtraData(tooLong); got != nil {
		t.Fatalf("expected nil for oversized extradata, got len=%d", len(got))
	}
}
