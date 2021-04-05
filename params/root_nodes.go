package params

import (
	"gitlab.com/q-dev/q-client/common"
)

// DevnetRootNodes - initial root nodes list on devnet.
var DevnetRootNodes = common.RootList{
	Timestamp: 0x60646e00,
	Nodes: []common.Address{
		common.HexToAddress("0x66316FfA38490d4d072F34EF7D7BA64Ce6b4478e"),
		common.HexToAddress("0xd10a97806b8FdFC8E4CC83a49f35CCF513F0a1f3"),
		common.HexToAddress("0xcca19442F5b3e5Fa71aaE69C092aC280e81Fd39f"),
	},
}

// DarrowRootNodes.
var DarrowRootNodes = common.RootList{
	Timestamp: 0x602bd9a6,
	Nodes: []common.Address{
		common.HexToAddress("0x66316FfA38490d4d072F34EF7D7BA64Ce6b4478e"),
		common.HexToAddress("0xd10a97806b8FdFC8E4CC83a49f35CCF513F0a1f3"),
		common.HexToAddress("0xcca19442F5b3e5Fa71aaE69C092aC280e81Fd39f"),
	},
}

var MainnetRootNodes = DarrowRootNodes
