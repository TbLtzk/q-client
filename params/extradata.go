package params

import (
	"runtime"
	"strings"

	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/rlp"
)

// MakeExtraData returns miner extradata for Clique vanity.
// If extra is empty, it builds the default RLP fingerprint matching the
// historical 4-field geth shape:
//
//	[packedQVersion, "QGOV Client", shortGo, GOOS]
//
// shortGo is runtime.Version() truncated to go<major>.<minor> so the blob
// stays within MaximumExtraDataSize (32) on linux/darwin/windows. If it still
// exceeds the limit, the Go field is cleared (arity preserved) before giving up.
// Non-empty extra is returned as-is unless it exceeds the size limit (then nil).
func MakeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		extra = encodeDefaultExtraData(shortGoVersion())
		if uint64(len(extra)) > MaximumExtraDataSize {
			extra = encodeDefaultExtraData("")
		}
	}
	if uint64(len(extra)) > MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", MaximumExtraDataSize)
		return nil
	}
	return extra
}

func encodeDefaultExtraData(goTag string) []byte {
	extra, _ := rlp.EncodeToBytes([]interface{}{
		uint(QVersionMajor<<16 | QVersionMinor<<8 | QVersionPatch),
		"QGOV Client",
		goTag,
		runtime.GOOS,
	})
	return extra
}

// shortGoVersion returns a bounded Go toolchain tag (e.g. "go1.26" from "go1.26.1")
// so default vanity RLP fits Clique's 32-byte limit with the "QGOV Client" name.
func shortGoVersion() string {
	v := runtime.Version()
	if !strings.HasPrefix(v, "go") {
		return v
	}
	rest := strings.TrimPrefix(v, "go")
	parts := strings.Split(rest, ".")
	if len(parts) >= 2 {
		return "go" + parts[0] + "." + parts[1]
	}
	return v
}
