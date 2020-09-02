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
	StatusMsg        = 0x00
	NewSignaturesMsg = 0x02
	NewRootListMsg   = 0x03
)

type statusMsgData struct {
	common.RootList
}

type newSignaturesMsg struct {
	Hash       common.Hash
	Signatures [][]byte
}

func initNewSignaturesMsg(hash common.Hash, sigs map[common.Address][]byte) newSignaturesMsg {
	var signatures [][]byte
	for _, sig := range sigs {
		signatures = append(signatures, sig)
	}

	return newSignaturesMsg{
		Hash:       hash,
		Signatures: signatures,
	}
}
