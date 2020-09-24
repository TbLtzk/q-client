package params

import (
	"gitlab.com/q-dev/q-client/common"
)

// DevnetRootNodes - initial root nodes list on devnet.
var DevnetRootNodes = common.RootList{
	Timestamp: 1597927813,
	Nodes: []common.Address{
		common.HexToAddress("0x42b9e4ffa28bc4bd9c8f7207f3a296ce960a365d"),
		common.HexToAddress("0x7c7df968bb4aff69b6f613c713c404f88de394d3"),
	},
}
