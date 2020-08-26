package governance

import (
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/crypto"
)

// protocolName is the official short name of the protocol used during capability negotiation.
const protocolName = "q"

// Constants to match up protocol versions and messages
const (
	q1 = 1
)

// ProtocolVersions are the supported versions of the eth protocol (first is primary).
var ProtocolVersions = []uint{q1}

// protocolLengths are the number of implemented message corresponding to different protocol versions.
var protocolLengths = map[uint]uint64{q1: 3}

// maximum possible number of root nodes
const maxNRootNodes = 101

// Maximum cap on the size of a protocol message.
// (doubled max cap list)
const protocolMaxMsgSize = 2*maxNRootNodes*(crypto.SignatureLength+common.AddressLength) + common.HashLength

// protocol message codes
const (
	GetRootListMsg = 0x01
	RootListMsg    = 0x02
	NewRootListMsg = 0x03
)

type rootListData struct {
	Timestamp  uint64
	RootList   []common.Address
	Digest     common.Hash

	Signatures [][]byte
}
