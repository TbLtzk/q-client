package governance

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

// protocolName is the official short name of the protocol used during capability negotiation.
const protocolName = "qgov"
//const protocolName = "qgov"

// Constants to match up protocol versions and messages
const (
	qgov2 = 2
	qgov3 = 3
)

// ProtocolVersions are the supported versions of the gov protocol (first is primary).
//var ProtocolVersions = []uint{qgov2}
var ProtocolVersions = []uint{qgov2,qgov3}

// protocolLengths are the number of implemented message corresponding to different protocol versions.
//var protocolLengths = map[uint]uint64{qgov2: 4}
var protocolLengths = map[uint]uint64{qgov2: 3,qgov3: 4}

// maximum possible number of root nodes
const maxNRootNodes = 101

// Maximum cap on the size of a protocol message.
// (doubled max cap list)
const protocolMaxMsgSize = 2*maxNRootNodes*(crypto.SignatureLength+common.AddressLength) + common.HashLength

// protocol message codes
const (
	StatusMsg        = 0x00
	RootListMsg      = 0x01
	ExclusionListMsg = 0x02
	ApprovalMsg      = 0x03
)

type statusMsgBody struct {
	CurrentRootList  common.RootList
	DesiredRootList  common.RootList
	ProposedRootList common.RootList

	CurrentExclusionList  common.ValidatorExclusionList
	DesiredExclusionList  common.ValidatorExclusionList
	ProposedExclusionList common.ValidatorExclusionList

	Network uint64
}
