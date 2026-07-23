package params

import (
	"runtime"
	"strings"
	"testing"

	"gitlab.com/q-dev/q-client/rlp"
)

type defaultExtraPayload struct {
	Version uint
	Name    string
	Go      string
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
	wantGo := shortGoVersion()
	if decoded.Go != wantGo && decoded.Go != "" {
		t.Fatalf("go tag=%q want %q or empty fallback", decoded.Go, wantGo)
	}
}

func TestShortGoVersion(t *testing.T) {
	got := shortGoVersion()
	if !strings.HasPrefix(got, "go") {
		t.Fatalf("shortGoVersion=%q missing go prefix", got)
	}
	full := runtime.Version()
	if strings.Count(got, ".") > strings.Count(full, ".") {
		t.Fatalf("shortGoVersion=%q longer than runtime.Version=%q", got, full)
	}
	// Prefer major.minor only when runtime reports a patch component.
	if parts := strings.Split(strings.TrimPrefix(full, "go"), "."); len(parts) >= 3 {
		if strings.Count(got, ".") != 1 {
			t.Fatalf("shortGoVersion=%q want major.minor form from %q", got, full)
		}
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

func TestEncodeDefaultExtraDataFitsCommonOS(t *testing.T) {
	goTag := shortGoVersion()
	for _, osName := range []string{"linux", "darwin", "windows"} {
		extra, err := rlp.EncodeToBytes([]interface{}{
			uint(QVersionMajor<<16 | QVersionMinor<<8 | QVersionPatch),
			"QGOV Client",
			goTag,
			osName,
		})
		if err != nil {
			t.Fatal(err)
		}
		if uint64(len(extra)) > MaximumExtraDataSize {
			t.Fatalf("os=%s go=%s len=%d exceeds %d", osName, goTag, len(extra), MaximumExtraDataSize)
		}
	}
}
