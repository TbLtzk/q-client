// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
	"errors"
	"math/big"
	"strings"

	ethereum "gitlab.com/q-dev/q-client"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LiquidationAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type LiquidationAuctionAuctionInfo struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
}

// LiquidationAuctionMetaData contains all meta data concerning the LiquidationAuction contract.
var LiquidationAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumLiquidationAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLiquidationAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"enumLiquidationAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vaultId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122ed806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630f41ba4b146100675780633b89bb861461008f57806344f91c1e146100a257806356dd4755146100fd5780636e256b3714610110578063f399e22e14610131575b600080fd5b61007a610075366004611c0c565b610146565b60405190151581526020015b60405180910390f35b61007a61009d366004611c41565b61057c565b6100ed6100b0366004611c41565b6001602081815260009384526040808520909152918352912080549181015460029091015460ff83169261010090046001600160a01b0316919084565b6040516100869493929190611ca5565b61007a61010b366004611c0c565b611009565b61012361011e366004611c41565b611522565b604051908152602001610086565b61014461013f366004611d44565b6115c8565b005b6001600160a01b0383166000908152600160209081526040808320858452909152808220815160808101909252805483929190829060ff16600281111561018f5761018f611c6d565b60028111156101a0576101a0611c6d565b8152815461010090046001600160a01b031660208201526001808301546040830152600290920154606090910152909150815160028111156101e4576101e4611c6d565b146102555760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3032343030315d2d5468652061756374696f6e206973206e6f742060448201527530b1ba34bb3296103330b4b632b2103a37903134b21760511b60648201526084015b60405180910390fd5b42816060015110156102db5760405162461bcd60e51b815260206004820152604360248201527f5b5145432d3032343030325d2d5468652061756374696f6e20666f722074686960448201527f73207661756c742069732066696e69736865642c206661696c656420746f206260648201526234b21760e91b608482015260a40161024c565b6102e585856116ac565b83101561038e5760405162461bcd60e51b8152602060048201526064602482018190527f5b5145432d3032343030335d2d5468652062696420616d6f756e74206d75737460448301527f206578636565642074686520686967686573742062696420627920746865206d908201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60848201526337b9329760e11b60a482015260c40161024c565b60006103986117db565b9050806001600160a01b031663a9059cbb836020015184604001516040518363ffffffff1660e01b81526004016103d0929190611dd6565b602060405180830381600087803b1580156103ea57600080fd5b505af11580156103fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104229190611e04565b506040516323b872dd60e01b81526001600160a01b038216906323b872dd9061045390339030908990600401611e1f565b602060405180830381600087803b15801561046d57600080fd5b505af1158015610481573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a59190611e04565b503360208084019190915260408084018690526001600160a01b0388166000908152600180845282822089835290935220835181548593839160ff1916908360028111156104f5576104f5611c6d565b02179055506020828101518254610100600160a81b0319166101006001600160a01b039283160217835560408085015160018501556060909401516002909301929092559151868152339288928a16917f7f863cab1b412077be133ef69af7df0e6aca4c67ff7d89db843d601f124a025d910160405180910390a450600195945050505050565b6001600160a01b0382166000908152600160209081526040808320848452909152808220815160808101909252805483929190829060ff1660028111156105c5576105c5611c6d565b60028111156105d6576105d6611c6d565b8152815461010090046001600160a01b031660208083019190915260018301546040808401919091526002909301546060909201919091528151808301909252600080835290820152909150600254600380546001600160a01b0390921691633fb90271916106cb9161064890611e43565b80601f016020809104026020016040519081016040528092919081815260200182805461067490611e43565b80156106c15780601f10610696576101008083540402835291602001916106c1565b820191906000526020600020905b8154815290600101906020018083116106a457829003601f168201915b50505050506118f7565b6040518263ffffffff1660e01b81526004016106e79190611eae565b60206040518083038186803b1580156106ff57600080fd5b505afa158015610713573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107379190611ee1565b6001600160a01b039081168252600254600380549190921691633fb90271916107e7919061076490611e43565b80601f016020809104026020016040519081016040528092919081815260200182805461079090611e43565b80156107dd5780601f106107b2576101008083540402835291602001916107dd565b820191906000526020600020905b8154815290600101906020018083116107c057829003601f168201915b5050505050611962565b6040518263ffffffff1660e01b81526004016108039190611eae565b60206040518083038186803b15801561081b57600080fd5b505afa15801561082f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108539190611ee1565b6001600160a01b03166020820152606082015142116108f35760405162461bcd60e51b815260206004820152605060248201527f5b5145432d3032343030345d2d5468652061756374696f6e206973206e6f742060448201527f66696e6973686564207965742c206661696c656420746f20657865637574652060648201526f3a3432903634b8bab4b230ba34b7b71760811b608482015260a40161024c565b60018251600281111561090857610908611c6d565b146109715760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3032343030355d2d5468652061756374696f6e2068617320616c7260448201527232b0b23c903132b2b71032bc32b1baba32b21760691b606482015260840161024c565b600061097b6117db565b90506000816001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016109ab9190611efe565b60206040518083038186803b1580156109c357600080fd5b505afa1580156109d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fb9190611f12565b9050600080600085600001516001600160a01b0316639c841fef8b8b6040518363ffffffff1660e01b8152600401610a34929190611dd6565b60006040518083038186803b158015610a4c57600080fd5b505afa158015610a60573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a889190810190611f2b565b9550509450505092506000610aa5676765c793fa10079d601b1b90565b610aad6119b7565b6001600160a01b031663498bff00610b5860038054610acb90611e43565b80601f0160208091040260200160405190810160405280929190818152602001828054610af790611e43565b8015610b445780601f10610b1957610100808354040283529160200191610b44565b820191906000526020600020905b815481529060010190602001808311610b2757829003601f168201915b505050505088611a0190919063ffffffff16565b6040518263ffffffff1660e01b8152600401610b749190611eae565b60206040518083038186803b158015610b8c57600080fd5b505afa158015610ba0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc49190611f12565b610bce9084611fef565b610bd8919061200e565b905081610be58282612030565b89604001511115610d1557866001600160a01b031663a9059cbb8d84868d60400151610c119190612048565b610c1b9190612048565b6040518363ffffffff1660e01b8152600401610c38929190611dd6565b602060405180830381600087803b158015610c5257600080fd5b505af1158015610c66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8a9190611e04565b50602088015160405163a9059cbb60e01b81526001600160a01b0389169163a9059cbb91610cbd91908690600401611dd6565b602060405180830381600087803b158015610cd757600080fd5b505af1158015610ceb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0f9190611e04565b50610e05565b8289604001511115610d6257866001600160a01b031663a9059cbb8960200151858c60400151610d459190612048565b6040518363ffffffff1660e01b8152600401610cbd929190611dd6565b50604088015183811015610e055787602001516001600160a01b0316632b7c7b118a6040015186610d939190612048565b6040518263ffffffff1660e01b8152600401610db191815260200190565b602060405180830381600087803b158015610dcb57600080fd5b505af1158015610ddf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e039190611e04565b505b6001600160a01b038c1660009081526001602090815260408083208e84529091529020805460ff19166002179055875160405163095ea7b360e01b81526001600160a01b0389169163095ea7b391610e6291908590600401611dd6565b602060405180830381600087803b158015610e7c57600080fd5b505af1158015610e90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb49190611e04565b50875160208a0151604051630bf9efcd60e01b81526001600160a01b038f81166004830152602482018f9052604482018590529182166064820152911690630bf9efcd90608401600060405180830381600087803b158015610f1557600080fd5b505af1158015610f29573d6000803e3d6000fd5b50505050610fb486886001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610f5d9190611efe565b60206040518083038186803b158015610f7557600080fd5b505afa158015610f89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fad9190611f12565b8e8e611a34565b8b6001600160a01b03167fe41a48e6ab980d2332baf65c0861a8264376c6bffdd1cece3e97272c446779a88c8b604051610fef92919061205f565b60405180910390a25060019b9a5050505050505050505050565b6000806110146119b7565b600254600380549293506000926001600160a01b0390921691633fb90271916110409161064890611e43565b6040518263ffffffff1660e01b815260040161105c9190611eae565b60206040518083038186803b15801561107457600080fd5b505afa158015611088573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ac9190611ee1565b90506000816001600160a01b0316639c841fef88886040518363ffffffff1660e01b81526004016110de929190611dd6565b60006040518083038186803b1580156110f657600080fd5b505afa15801561110a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111329190810190611f2b565b506001600160a01b038c1660009081526001602090815260408083208e845290915280822081516080810190925280549398509196509450925083915060ff16600281111561118357611183611c6d565b600281111561119457611194611c6d565b8152815461010090046001600160a01b031660208201526001820154604082015260029091015460609091015290506000815160028111156111d8576111d8611c6d565b146112695760405162461bcd60e51b815260206004820152605560248201527f5b5145432d3032343030305d2d41206c69717569646174696f6e20617563746960448201527f6f6e20666f72207468697320626f72726f77696e67207661756c74206861732060648201527430b63932b0b23c903132b2b71039ba30b93a32b21760591b608482015260a40161024c565b816112f05760405163bcbaf48760e01b81526001600160a01b0384169063bcbaf4879061129c908b908b90600401611dd6565b602060405180830381600087803b1580156112b657600080fd5b505af11580156112ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ee9190611e04565b505b6112f86117db565b6001600160a01b03166323b872dd3330896040518463ffffffff1660e01b815260040161132793929190611e1f565b602060405180830381600087803b15801561134157600080fd5b505af1158015611355573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113799190611e04565b506001818190525033602082015260408082018790525162498bff60e81b81526001600160a01b0385169063498bff00906113f29060040160208082526021908201527f676f7665726e65642e455044522e6c69717569646174696f6e41756374696f6e6040820152600560fc1b606082015260800190565b60206040518083038186803b15801561140a57600080fd5b505afa15801561141e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114429190611f12565b61144c9042612030565b60608201526001600160a01b03881660009081526001602081815260408084208b8552909152909120825181548493839160ff19169083600281111561149457611494611c6d565b02179055506020828101518254610100600160a81b0319166101006001600160a01b0392831602178355604080850151600185015560609094015160029093019290925582518a81529081018990523392918b16917f3f56989faf73e0fa48db704c98688b152de2669b2c98b59b7e6c22a7998fe5f9910160405180910390a3506001979650505050505050565b6001600160a01b038216600090815260016020818152604080842085855290915282205460ff16600281111561155a5761155a611c6d565b146115b75760405162461bcd60e51b815260206004820152602760248201527f5b5145432d3032343030365d2d5468652061756374696f6e206973206e6f742060448201526630b1ba34bb329760c91b606482015260840161024c565b6115c183836116ac565b9392505050565b600054610100900460ff16806115e1575060005460ff16155b6116445760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161024c565b600054610100900460ff16158015611666576000805461ffff19166101011790555b600280546001600160a01b0319166001600160a01b0385161790558151611694906003906020850190611b5b565b5080156116a7576000805461ff00191690555b505050565b6000806116b76119b7565b60405162498bff60e81b815260206004820152602160248201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6044820152601d60fa1b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b15801561172f57600080fd5b505afa158015611743573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117679190611f12565b6001600160a01b03851660009081526001602081815260408084208885529091528220015491925081676765c793fa10079d601b1b6117a68585611fef565b6117b0919061200e565b6117ba9190612030565b9050818114156117d257806117ce816120a5565b9150505b95945050505050565b600254600380546000926001600160a01b031691633fb9027191611886919061180390611e43565b80601f016020809104026020016040519081016040528092919081815260200182805461182f90611e43565b801561187c5780601f106118515761010080835404028352916020019161187c565b820191906000526020600020905b81548152906001019060200180831161185f57829003601f168201915b5050505050611a79565b6040518263ffffffff1660e01b81526004016118a29190611eae565b60206040518083038186803b1580156118ba57600080fd5b505afa1580156118ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f29190611ee1565b905090565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461194c94938792016120c0565b6040516020818303038152906040529050919050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461194c949387920161212b565b600254604080516060810190915260228082526000926001600160a01b031691633fb902719161229660208301396040518263ffffffff1660e01b81526004016118a29190611eae565b60606115c183836040518060400160405280600e81526020016d6c69717569646174696f6e46656560901b815250611ace565b6001600160a01b0382166000908152600160208181526040808420858552909152909120015483611a658287612048565b14611a7257611a7261219a565b5050505050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461194c94938792016121b0565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525084604051806040016040528060018152602001605f60f81b81525085604051806040016040528060018152602001605f60f81b81525086604051602001611b4396959493929190612216565b60405160208183030381529060405290509392505050565b828054611b6790611e43565b90600052602060002090601f016020900481019282611b895760008555611bcf565b82601f10611ba257805160ff1916838001178555611bcf565b82800160010185558215611bcf579182015b82811115611bcf578251825591602001919060010190611bb4565b50611bdb929150611bdf565b5090565b5b80821115611bdb5760008155600101611be0565b6001600160a01b0381168114611c0957600080fd5b50565b600080600060608486031215611c2157600080fd5b8335611c2c81611bf4565b95602085013595506040909401359392505050565b60008060408385031215611c5457600080fd5b8235611c5f81611bf4565b946020939093013593505050565b634e487b7160e01b600052602160045260246000fd5b60038110611ca157634e487b7160e01b600052602160045260246000fd5b9052565b60808101611cb38287611c83565b6001600160a01b039490941660208201526040810192909252606090910152919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611d1557611d15611cd7565b604052919050565b60006001600160401b03821115611d3657611d36611cd7565b50601f01601f191660200190565b60008060408385031215611d5757600080fd5b8235611d6281611bf4565b915060208301356001600160401b03811115611d7d57600080fd5b8301601f81018513611d8e57600080fd5b8035611da1611d9c82611d1d565b611ced565b818152866020838501011115611db657600080fd5b816020840160208301376000602083830101528093505050509250929050565b6001600160a01b03929092168252602082015260400190565b80518015158114611dff57600080fd5b919050565b600060208284031215611e1657600080fd5b6115c182611def565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600181811c90821680611e5757607f821691505b60208210811415611e7857634e487b7160e01b600052602260045260246000fd5b50919050565b60005b83811015611e99578181015183820152602001611e81565b83811115611ea8576000848401525b50505050565b6020815260008251806020840152611ecd816040850160208701611e7e565b601f01601f19169190910160400192915050565b600060208284031215611ef357600080fd5b81516115c181611bf4565b6001600160a01b0391909116815260200190565b600060208284031215611f2457600080fd5b5051919050565b60008060008060008060c08789031215611f4457600080fd5b86516001600160401b03811115611f5a57600080fd5b8701601f81018913611f6b57600080fd5b8051611f79611d9c82611d1d565b8181528a6020838501011115611f8e57600080fd5b611f9f826020830160208601611e7e565b809850505050602087015194506040870151935060608701519250611fc660808801611def565b915060a087015190509295509295509295565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561200957612009611fd9565b500290565b60008261202b57634e487b7160e01b600052601260045260246000fd5b500490565b6000821982111561204357612043611fd9565b500190565b60008282101561205a5761205a611fd9565b500390565b600060a082019050838252612078602083018451611c83565b60018060a01b03602084015116604083015260408301516060830152606083015160808301529392505050565b60006000198214156120b9576120b9611fd9565b5060010190565b600085516120d2818460208a01611e7e565b8551908301906120e6818360208a01611e7e565b85519101906120f9818360208901611e7e565b845191019061210c818360208801611e7e565b68626f72726f77696e6760b81b91019081526009019695505050505050565b6000855161213d818460208a01611e7e565b855190830190612151818360208a01611e7e565b8551910190612164818360208901611e7e565b8451910190612177818360208801611e7e565b6c73797374656d42616c616e636560981b9101908152600d019695505050505050565b634e487b7160e01b600052600160045260246000fd5b600085516121c2818460208a01611e7e565b8551908301906121d6818360208a01611e7e565b85519101906121e9818360208901611e7e565b84519101906121fc818360208801611e7e565b6331b7b4b760e11b91019081526004019695505050505050565b6000875160206122298285838d01611e7e565b88519184019161223c8184848d01611e7e565b885192019161224e8184848c01611e7e565b87519201916122608184848b01611e7e565b86519201916122728184848a01611e7e565b85519201916122848184848901611e7e565b91909101999850505050505050505056fe676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a26469706673582212209225e2a3e0d5e6511efd8b0bcbf417931d413af139a956a61d6a9bc371f405ae64736f6c63430008090033",
}

// LiquidationAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidationAuctionMetaData.ABI instead.
var LiquidationAuctionABI = LiquidationAuctionMetaData.ABI

// LiquidationAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidationAuctionMetaData.Bin instead.
var LiquidationAuctionBin = LiquidationAuctionMetaData.Bin

// DeployLiquidationAuction deploys a new Ethereum contract, binding an instance of LiquidationAuction to it.
func DeployLiquidationAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LiquidationAuction, error) {
	parsed, err := LiquidationAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidationAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidationAuction{LiquidationAuctionCaller: LiquidationAuctionCaller{contract: contract}, LiquidationAuctionTransactor: LiquidationAuctionTransactor{contract: contract}, LiquidationAuctionFilterer: LiquidationAuctionFilterer{contract: contract}}, nil
}

// LiquidationAuction is an auto generated Go binding around an Ethereum contract.
type LiquidationAuction struct {
	LiquidationAuctionCaller     // Read-only binding to the contract
	LiquidationAuctionTransactor // Write-only binding to the contract
	LiquidationAuctionFilterer   // Log filterer for contract events
}

// LiquidationAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidationAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidationAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidationAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidationAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidationAuctionSession struct {
	Contract     *LiquidationAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LiquidationAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidationAuctionCallerSession struct {
	Contract *LiquidationAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LiquidationAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidationAuctionTransactorSession struct {
	Contract     *LiquidationAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LiquidationAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidationAuctionRaw struct {
	Contract *LiquidationAuction // Generic contract binding to access the raw methods on
}

// LiquidationAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidationAuctionCallerRaw struct {
	Contract *LiquidationAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidationAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidationAuctionTransactorRaw struct {
	Contract *LiquidationAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidationAuction creates a new instance of LiquidationAuction, bound to a specific deployed contract.
func NewLiquidationAuction(address common.Address, backend bind.ContractBackend) (*LiquidationAuction, error) {
	contract, err := bindLiquidationAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuction{LiquidationAuctionCaller: LiquidationAuctionCaller{contract: contract}, LiquidationAuctionTransactor: LiquidationAuctionTransactor{contract: contract}, LiquidationAuctionFilterer: LiquidationAuctionFilterer{contract: contract}}, nil
}

// NewLiquidationAuctionCaller creates a new read-only instance of LiquidationAuction, bound to a specific deployed contract.
func NewLiquidationAuctionCaller(address common.Address, caller bind.ContractCaller) (*LiquidationAuctionCaller, error) {
	contract, err := bindLiquidationAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionCaller{contract: contract}, nil
}

// NewLiquidationAuctionTransactor creates a new write-only instance of LiquidationAuction, bound to a specific deployed contract.
func NewLiquidationAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidationAuctionTransactor, error) {
	contract, err := bindLiquidationAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionTransactor{contract: contract}, nil
}

// NewLiquidationAuctionFilterer creates a new log filterer instance of LiquidationAuction, bound to a specific deployed contract.
func NewLiquidationAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidationAuctionFilterer, error) {
	contract, err := bindLiquidationAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionFilterer{contract: contract}, nil
}

// bindLiquidationAuction binds a generic wrapper to an already deployed contract.
func bindLiquidationAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidationAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationAuction *LiquidationAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationAuction.Contract.LiquidationAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationAuction *LiquidationAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.LiquidationAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationAuction *LiquidationAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.LiquidationAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidationAuction *LiquidationAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidationAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidationAuction *LiquidationAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidationAuction *LiquidationAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime)
func (_LiquidationAuction *LiquidationAuctionCaller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
}, error) {
	var out []interface{}
	err := _LiquidationAuction.contract.Call(opts, &out, "auctions", arg0, arg1)

	outstruct := new(struct {
		Status     uint8
		Bidder     common.Address
		HighestBid *big.Int
		EndTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime)
func (_LiquidationAuction *LiquidationAuctionSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
}, error) {
	return _LiquidationAuction.Contract.Auctions(&_LiquidationAuction.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime)
func (_LiquidationAuction *LiquidationAuctionCallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
}, error) {
	return _LiquidationAuction.Contract.Auctions(&_LiquidationAuction.CallOpts, arg0, arg1)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0x6e256b37.
//
// Solidity: function getRaisingBid(address _user, uint256 _vaultId) view returns(uint256)
func (_LiquidationAuction *LiquidationAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _user common.Address, _vaultId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidationAuction.contract.Call(opts, &out, "getRaisingBid", _user, _vaultId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaisingBid is a free data retrieval call binding the contract method 0x6e256b37.
//
// Solidity: function getRaisingBid(address _user, uint256 _vaultId) view returns(uint256)
func (_LiquidationAuction *LiquidationAuctionSession) GetRaisingBid(_user common.Address, _vaultId *big.Int) (*big.Int, error) {
	return _LiquidationAuction.Contract.GetRaisingBid(&_LiquidationAuction.CallOpts, _user, _vaultId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0x6e256b37.
//
// Solidity: function getRaisingBid(address _user, uint256 _vaultId) view returns(uint256)
func (_LiquidationAuction *LiquidationAuctionCallerSession) GetRaisingBid(_user common.Address, _vaultId *big.Int) (*big.Int, error) {
	return _LiquidationAuction.Contract.GetRaisingBid(&_LiquidationAuction.CallOpts, _user, _vaultId)
}

// Bid is a paid mutator transaction binding the contract method 0x0f41ba4b.
//
// Solidity: function bid(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactor) Bid(opts *bind.TransactOpts, _user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.contract.Transact(opts, "bid", _user, _vaultId, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x0f41ba4b.
//
// Solidity: function bid(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionSession) Bid(_user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Bid(&_LiquidationAuction.TransactOpts, _user, _vaultId, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x0f41ba4b.
//
// Solidity: function bid(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactorSession) Bid(_user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Bid(&_LiquidationAuction.TransactOpts, _user, _vaultId, _bid)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _user, uint256 _vaultId) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactor) Execute(opts *bind.TransactOpts, _user common.Address, _vaultId *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.contract.Transact(opts, "execute", _user, _vaultId)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _user, uint256 _vaultId) returns(bool)
func (_LiquidationAuction *LiquidationAuctionSession) Execute(_user common.Address, _vaultId *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Execute(&_LiquidationAuction.TransactOpts, _user, _vaultId)
}

// Execute is a paid mutator transaction binding the contract method 0x3b89bb86.
//
// Solidity: function execute(address _user, uint256 _vaultId) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactorSession) Execute(_user common.Address, _vaultId *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Execute(&_LiquidationAuction.TransactOpts, _user, _vaultId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_LiquidationAuction *LiquidationAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _LiquidationAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_LiquidationAuction *LiquidationAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Initialize(&_LiquidationAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_LiquidationAuction *LiquidationAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.Initialize(&_LiquidationAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x56dd4755.
//
// Solidity: function startAuction(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactor) StartAuction(opts *bind.TransactOpts, _user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.contract.Transact(opts, "startAuction", _user, _vaultId, _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x56dd4755.
//
// Solidity: function startAuction(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionSession) StartAuction(_user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.StartAuction(&_LiquidationAuction.TransactOpts, _user, _vaultId, _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x56dd4755.
//
// Solidity: function startAuction(address _user, uint256 _vaultId, uint256 _bid) returns(bool)
func (_LiquidationAuction *LiquidationAuctionTransactorSession) StartAuction(_user common.Address, _vaultId *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _LiquidationAuction.Contract.StartAuction(&_LiquidationAuction.TransactOpts, _user, _vaultId, _bid)
}

// LiquidationAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the LiquidationAuction contract.
type LiquidationAuctionAuctionStartedIterator struct {
	Event *LiquidationAuctionAuctionStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationAuctionAuctionStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationAuctionAuctionStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationAuctionAuctionStarted represents a AuctionStarted event raised by the LiquidationAuction contract.
type LiquidationAuctionAuctionStarted struct {
	User    common.Address
	VaultId *big.Int
	Bidder  common.Address
	Bid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x3f56989faf73e0fa48db704c98688b152de2669b2c98b59b7e6c22a7998fe5f9.
//
// Solidity: event AuctionStarted(address indexed _user, uint256 _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _user []common.Address, _bidder []common.Address) (*LiquidationAuctionAuctionStartedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _LiquidationAuction.contract.FilterLogs(opts, "AuctionStarted", _userRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionAuctionStartedIterator{contract: _LiquidationAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x3f56989faf73e0fa48db704c98688b152de2669b2c98b59b7e6c22a7998fe5f9.
//
// Solidity: event AuctionStarted(address indexed _user, uint256 _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *LiquidationAuctionAuctionStarted, _user []common.Address, _bidder []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _LiquidationAuction.contract.WatchLogs(opts, "AuctionStarted", _userRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationAuctionAuctionStarted)
				if err := _LiquidationAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuctionStarted is a log parse operation binding the contract event 0x3f56989faf73e0fa48db704c98688b152de2669b2c98b59b7e6c22a7998fe5f9.
//
// Solidity: event AuctionStarted(address indexed _user, uint256 _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) ParseAuctionStarted(log types.Log) (*LiquidationAuctionAuctionStarted, error) {
	event := new(LiquidationAuctionAuctionStarted)
	if err := _LiquidationAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the LiquidationAuction contract.
type LiquidationAuctionBidIterator struct {
	Event *LiquidationAuctionBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationAuctionBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationAuctionBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationAuctionBid represents a Bid event raised by the LiquidationAuction contract.
type LiquidationAuctionBid struct {
	User    common.Address
	VaultId *big.Int
	Bidder  common.Address
	Bid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x7f863cab1b412077be133ef69af7df0e6aca4c67ff7d89db843d601f124a025d.
//
// Solidity: event Bid(address indexed _user, uint256 indexed _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) FilterBid(opts *bind.FilterOpts, _user []common.Address, _vaultId []*big.Int, _bidder []common.Address) (*LiquidationAuctionBidIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _vaultIdRule []interface{}
	for _, _vaultIdItem := range _vaultId {
		_vaultIdRule = append(_vaultIdRule, _vaultIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _LiquidationAuction.contract.FilterLogs(opts, "Bid", _userRule, _vaultIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionBidIterator{contract: _LiquidationAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x7f863cab1b412077be133ef69af7df0e6aca4c67ff7d89db843d601f124a025d.
//
// Solidity: event Bid(address indexed _user, uint256 indexed _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *LiquidationAuctionBid, _user []common.Address, _vaultId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _vaultIdRule []interface{}
	for _, _vaultIdItem := range _vaultId {
		_vaultIdRule = append(_vaultIdRule, _vaultIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _LiquidationAuction.contract.WatchLogs(opts, "Bid", _userRule, _vaultIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationAuctionBid)
				if err := _LiquidationAuction.contract.UnpackLog(event, "Bid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBid is a log parse operation binding the contract event 0x7f863cab1b412077be133ef69af7df0e6aca4c67ff7d89db843d601f124a025d.
//
// Solidity: event Bid(address indexed _user, uint256 indexed _vaultId, address indexed _bidder, uint256 _bid)
func (_LiquidationAuction *LiquidationAuctionFilterer) ParseBid(log types.Log) (*LiquidationAuctionBid, error) {
	event := new(LiquidationAuctionBid)
	if err := _LiquidationAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidationAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the LiquidationAuction contract.
type LiquidationAuctionExecutedIterator struct {
	Event *LiquidationAuctionExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidationAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidationAuctionExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidationAuctionExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidationAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidationAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidationAuctionExecuted represents a Executed event raised by the LiquidationAuction contract.
type LiquidationAuctionExecuted struct {
	User    common.Address
	VaultId *big.Int
	Info    LiquidationAuctionAuctionInfo
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xe41a48e6ab980d2332baf65c0861a8264376c6bffdd1cece3e97272c446779a8.
//
// Solidity: event Executed(address indexed _user, uint256 _vaultId, (uint8,address,uint256,uint256) _info)
func (_LiquidationAuction *LiquidationAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _user []common.Address) (*LiquidationAuctionExecutedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LiquidationAuction.contract.FilterLogs(opts, "Executed", _userRule)
	if err != nil {
		return nil, err
	}
	return &LiquidationAuctionExecutedIterator{contract: _LiquidationAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xe41a48e6ab980d2332baf65c0861a8264376c6bffdd1cece3e97272c446779a8.
//
// Solidity: event Executed(address indexed _user, uint256 _vaultId, (uint8,address,uint256,uint256) _info)
func (_LiquidationAuction *LiquidationAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *LiquidationAuctionExecuted, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LiquidationAuction.contract.WatchLogs(opts, "Executed", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidationAuctionExecuted)
				if err := _LiquidationAuction.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuted is a log parse operation binding the contract event 0xe41a48e6ab980d2332baf65c0861a8264376c6bffdd1cece3e97272c446779a8.
//
// Solidity: event Executed(address indexed _user, uint256 _vaultId, (uint8,address,uint256,uint256) _info)
func (_LiquidationAuction *LiquidationAuctionFilterer) ParseExecuted(log types.Log) (*LiquidationAuctionExecuted, error) {
	event := new(LiquidationAuctionExecuted)
	if err := _LiquidationAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
