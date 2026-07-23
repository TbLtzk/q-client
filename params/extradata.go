package params

import (
	"runtime"

	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/rlp"
)

// MakeExtraData returns miner extradata for Clique vanity.
// If extra is empty, it builds the default RLP fingerprint: packed QVersion,
// "QGOV Client", and GOOS. runtime.Version() is omitted so the blob fits in
// MaximumExtraDataSize (32). Non-empty extra is returned as-is unless it exceeds
// the size limit (then nil, with a warning).
func MakeExtraData(extra []byte) []byte {
	if len(extra) == 0 {
		extra, _ = rlp.EncodeToBytes([]interface{}{
			uint(QVersionMajor<<16 | QVersionMinor<<8 | QVersionPatch),
			"QGOV Client",
			runtime.GOOS,
		})
	}
	if uint64(len(extra)) > MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", MaximumExtraDataSize)
		return nil
	}
	return extra
}
