package params

import (
	"runtime"

	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/rlp"
)

// DefaultMinerExtraData builds the default miner extradata used as Clique vanity
// when --miner.extradata is unset: packed QVersion, "QGOV Client", and GOOS.
// runtime.Version() is omitted so the RLP fits in MaximumExtraDataSize (32).
func DefaultMinerExtraData() []byte {
	extra, _ := rlp.EncodeToBytes([]interface{}{
		uint(QVersionMajor<<16 | QVersionMinor<<8 | QVersionPatch),
		"QGOV Client",
		runtime.GOOS,
	})
	if uint64(len(extra)) > MaximumExtraDataSize {
		log.Warn("Miner extra data exceed limit", "extra", hexutil.Bytes(extra), "limit", MaximumExtraDataSize)
		return nil
	}
	return extra
}
