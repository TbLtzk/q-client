package params

import (
	"gitlab.com/q-dev/q-client/common"
)

// DevnetRootNodes - initial root nodes list on devnet.
var DevnetRootNodes = common.RootList{
	Timestamp: 0x6131c7b9,
	Nodes: []common.Address{
		common.HexToAddress("0x64D4edeFE8bA86d3588B213b0A053e7B910Cad68"), // Miner 1
		common.HexToAddress("0x4a14D788D86D021670EBcecE1196631d66595984"), // Miner 3
	},
}

// TestnetRootNodes - initial root nodes list on testnet.
var TestnetRootNodes = common.RootList{
	Timestamp: 0x607843b9,
	Nodes: []common.Address{
		common.HexToAddress("0x64D4edeFE8bA86d3588B213b0A053e7B910Cad68"), // Miner 1
		common.HexToAddress("0x4a14D788D86D021670EBcecE1196631d66595984"), // Miner 3
	},
}

var MainnetRootNodes = TestnetRootNodes
