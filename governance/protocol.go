package governance

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

// Constants to match up protocol versions and messages
const (
	protocolName = "qgov" // protocolName is the official short name of the protocol used during capability negotiation.
	qgov2        = 2      //TODO decide when to remove this outdated protocol
	qgov3        = 3
	qgov4        = 4
)

// ProtocolVersions are the supported versions of the gov protocol (first is primary).
var ProtocolVersions = []uint{qgov2, qgov3, qgov4}

// protocolLengths are the number of implemented message corresponding to different protocol versions.
var protocolLengths = map[uint]uint64{qgov2: 3, qgov3: 4, qgov4: 7}

// maximum possible number of root nodes
const maxNRootNodes = 101

// Maximum cap on the size of a protocol message.
// (doubled max cap list)
const protocolMaxMsgSize = 2*maxNRootNodes*(crypto.SignatureLength+common.AddressLength) + common.HashLength
const maxConstitutionFileSize = 512000 //TODO define correct value

// protocol message codes
const (
	StatusMsg                  = 0x00
	RootListMsg                = 0x01
	ExclusionListMsg           = 0x02
	ApprovalMsg                = 0x03
	ConstitutionFileRequestMsg = 0x04
	ConstitutionFilesMsg       = 0x05
	KnownConstitutionFilesMsg  = 0x06
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
