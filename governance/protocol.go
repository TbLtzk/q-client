package governance

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"
)

// Constants to match up protocol versions and messages
const (
	protocolName = "qgov" // protocolName is the official short name of the protocol used during capability negotiation.
	qgov5        = 5      // Exclusion sets new algorithm
)

// ProtocolVersions are the supported versions of the gov protocol (first is primary).
var ProtocolVersions = []uint{qgov5}

// protocolLengths are the number of implemented message corresponding to different protocol versions.
var protocolLengths = map[uint]uint64{qgov5: 7}

// maximum possible number of root nodes
const maxNRootNodes = 101

// Maximum cap on the size of a protocol message.
// (doubled max cap list)
const protocolMaxMsgSize = 2*maxNRootNodes*(crypto.SignatureLength+common.AddressLength) + common.HashLength

// maxConstitutionFileSize limits the size of the compressed governance payload
// that can be received over the wire.
const maxConstitutionFileSize = 512000 // TODO define correct value

// maxConstitutionDecompressedSize limits the size of a single decompressed
// constitution file. This prevents gzip "zip bombs" from exhausting memory.
const maxConstitutionDecompressedSize int64 = 1 * 1024 * 1024 // 1 MiB

// maxConstitutionTotalDecompressedSize limits the total size of all
// decompressed constitution files carried in a single message.
const maxConstitutionTotalDecompressedSize int64 = 2 * 1024 * 1024 // 2 MiB

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
