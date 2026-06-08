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
	Timestamp: 0x6a26c125,
	Nodes: []common.Address{
		common.HexToAddress("0x64D4edeFE8bA86d3588B213b0A053e7B910Cad68"), // Miner 1
		common.HexToAddress("0x4a14D788D86D021670EBcecE1196631d66595984"), // Miner 3
	},
}

// MainnetRootNodes - initial root nodes list on mainnet.
var MainnetRootNodes = common.RootList{
	Timestamp: 0x64a417dd,
	Nodes: []common.Address{
		common.HexToAddress("0x87252d4c5831779af285fb657627216e05a4d7b2"),
		common.HexToAddress("0x7c2d7f8eec1d0884da541143b1db888009aa9fd5"),
		common.HexToAddress("0x764b072ef6363f4fea8c313c74d701952e30a14e"),
		common.HexToAddress("0x16ad35f8c92d9a4dd41d98aa6d68d4f6a35aa3d4"),
		common.HexToAddress("0x08cc588a4d445f5883dc6157ab366d240c3f5236"),
		common.HexToAddress("0xf6ce519a81a94138266c85866b351925bdda2aac"),
		common.HexToAddress("0xbe45d0967d9e05bb46c08a7d8635fe7595e454ce"),
		common.HexToAddress("0xb255fd9bef1b225f63e81f891c5ac861c1699edf"),
		common.HexToAddress("0xa11863a92a9a0870fda6ff1b200d1d50fa9d11ae"),
		common.HexToAddress("0x6b7559dc3113291d0a95114e48ab46673dcbb11b"),
		common.HexToAddress("0x659f8a787c2f06b07fa28911eab4c0bf54b9447c"),
		common.HexToAddress("0x65999d1299896e93b99823341b41840efa7a1734"),
		common.HexToAddress("0x62c64ca23d8252a2213dce623c918d5dc9dfe6c3"),
		common.HexToAddress("0x53a5bd93b6f32060aa35996deed9d79e4324b296"),
		common.HexToAddress("0x5295c5ae9376388c0b66f89fe91340d6caf2caf4"),
		common.HexToAddress("0x35ac9eca581dd7ede69d6a56ec7b793e2fe293dc"),
		common.HexToAddress("0x2ea7677ee90480ab62e968a473f40a614c05cb99"),
		common.HexToAddress("0x1f96984571abe475cde8699791bf8244707a4021"),
		common.HexToAddress("0x083af739c6c4bafb22a9c154925d761ce603b657"),
		common.HexToAddress("0x02ba6f1246fdb2296a539472ad5d7861eeb08f3a"),
		common.HexToAddress("0xcdbb36e2F844823a5EEffbc91A324b59fd40Fae4"),
		common.HexToAddress("0x48cC8D45859f578BC6A8E76B2FD719be80eDf4C5"),
		common.HexToAddress("0xB1FEDcC8BCA3f07702D3F92538f281a5095b0537"),
		common.HexToAddress("0xCb630e36DA08328d80FEFed490457c674B7a76B5"),
		common.HexToAddress("0x5C4acBD71879BFeD7322f0FdceF3CF86D9709543"),
		common.HexToAddress("0x6003e395038934fcac143c440d000ae204688d71"),
		common.HexToAddress("0xE0BbB63fDA7f1EbbEE2DAEd6BFD69E0313C15bfb"),
	},
}
