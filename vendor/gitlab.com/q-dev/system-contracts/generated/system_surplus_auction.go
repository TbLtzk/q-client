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

// SystemSurplusAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemSurplusAuctionAuctionInfo struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}

// SystemSurplusAuctionMetaData contains all meta data concerning the SystemSurplusAuction contract.
var SystemSurplusAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSystemSurplusAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611af9806100206000396000f3fe6080604052600436106100605760003560e01c8063454a2ab314610065578063571a26a01461008d5780636b64c76914610115578063a21659201461012b578063b830620a1461014b578063f399e22e14610161578063fe0d94c114610183575b600080fd5b610078610073366004611683565b6101a3565b60405190151581526020015b60405180910390f35b34801561009957600080fd5b506100e36100a8366004611683565b6003602081905260009182526040909120805460018201546002830154929093015460ff8216936101009092046001600160a01b0316929085565b6040805195151586526001600160a01b039094166020860152928401919091526060830152608082015260a001610084565b61011d6104a8565b604051908152602001610084565b34801561013757600080fd5b5061011d610146366004611683565b610c99565b34801561015757600080fd5b5061011d60025481565b34801561016d57600080fd5b5061018161017c3660046116ca565b610ccd565b005b34801561018f57600080fd5b5061007861019e366004611683565b610db8565b60008160025481106101d05760405162461bcd60e51b81526004016101c79061178e565b60405180910390fd5b6000838152600360208190526040909120908101544211156102515760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3032353030355d2d5468652061756374696f6e2069732066696e6960448201527339b432b216103330b4b632b2103a37903134b21760611b60648201526084016101c7565b61025a84611288565b3410156103035760405162461bcd60e51b8152602060048201526064602482018190527f5b5145432d3032353030365d2d5468652062696420616d6f756e74206d75737460448301527f206578636565642074686520686967686573742062696420627920746865206d908201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60848201526337b9329760e11b60a482015260c4016101c7565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161036791600401611806565b60206040518083038186803b15801561037f57600080fd5b505afa158015610393573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b79190611839565b600183018054845433610100908102610100600160a81b0319831617875534909355604051632377789f60e21b81526001600160a01b0393909104831660048201819052939450909291841690638ddde27c9084906024016020604051808303818588803b15801561042857600080fd5b505af115801561043c573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906104619190611856565b50604051348152339088907fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca31999060200160405180910390a360019550505050505b50919050565b60003461051a5760405162461bcd60e51b815260206004820152603a60248201527f5b5145432d3032353030305d2d496e76616c69642062696420616d6f756e742c604482015279103330b4b632b21039ba30b93a103a34329030bab1ba34b7b71760311b60648201526084016101c7565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716105d26001805461054f90611878565b80601f016020809104026020016040519081016040528092919081815260200182805461057b90611878565b80156105c85780601f1061059d576101008083540402835291602001916105c8565b820191906000526020600020905b8154815290600101906020018083116105ab57829003601f168201915b50505050506113a4565b6040518263ffffffff1660e01b81526004016105ee9190611806565b60206040518083038186803b15801561060657600080fd5b505afa15801561061a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063e9190611839565b9050806001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561067b57600080fd5b505af115801561068f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b39190611856565b5060006106be61140f565b90506000816001600160a01b031663498bff00610764600180546106e190611878565b80601f016020809104026020016040519081016040528092919081815260200182805461070d90611878565b801561075a5780601f1061072f5761010080835404028352916020019161075a565b820191906000526020600020905b81548152906001019060200180831161073d57829003601f168201915b50505050506114be565b6040518263ffffffff1660e01b81526004016107809190611806565b60206040518083038186803b15801561079857600080fd5b505afa1580156107ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d091906118ad565b90506000826001600160a01b031663498bff00610876600180546107f390611878565b80601f016020809104026020016040519081016040528092919081815260200182805461081f90611878565b801561086c5780601f106108415761010080835404028352916020019161086c565b820191906000526020600020905b81548152906001019060200180831161084f57829003601f168201915b50505050506114f2565b6040518263ffffffff1660e01b81526004016108929190611806565b60206040518083038186803b1580156108aa57600080fd5b505afa1580156108be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e291906118ad565b90506000846001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b15801561091f57600080fd5b505afa158015610933573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095791906118ad565b9050818110156109de5760405162461bcd60e51b815260206004820152604660248201527f5b5145432d3032353030325d2d537572706c7573206166746572206e6574746960448201527f6e672069732062656c6f7720737572706c75732061756374696f6e207468726560648201526539b437b6321760d11b608482015260a4016101c7565b82811015610a4f5760405162461bcd60e51b815260206004820152603860248201527f5b5145432d3032353030335d2d4e6f7420656e6f75676820737572706c7573206044820152773a37903334b636103a34329030bab1ba34b7b7103637ba1760411b60648201526084016101c7565b60405163152353d960e01b8152600481018490526001600160a01b0386169063152353d990602401602060405180830381600087803b158015610a9157600080fd5b505af1158015610aa5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac99190611856565b50600280549081906000610adc836118dc565b9190505550610b1e6040518060a0016040528060001515815260200160006001600160a01b031681526020016000815260200160008152602001600081525090565b33602082015234604080830191909152606082018690525162498bff60e81b81526001600160a01b0387169063498bff0090610b8e906004016020808252601d908201527f676f7665726e65642e455044522e737572706c757341756374696f6e50000000604082015260600190565b60206040518083038186803b158015610ba657600080fd5b505afa158015610bba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bde91906118ad565b610be890426118f7565b6080820190815260008381526003602081815260409283902085518154838801516001600160a01b031661010002610100600160a81b0319921515929092166001600160a81b03199091161717815583860151600182015560608601516002820155935193909101929092558051848152349281019290925233917fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475910160405180910390a2509695505050505050565b6000816002548110610cbd5760405162461bcd60e51b81526004016101c79061178e565b610cc683611288565b9392505050565b600054610100900460ff1680610ce6575060005460ff16155b610d495760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016101c7565b600054610100900460ff16158015610d6b576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610da09060019060208501906115ea565b508015610db3576000805461ff00191690555b505050565b6000816002548110610ddc5760405162461bcd60e51b81526004016101c79061178e565b6000838152600360208190526040909120908101544211610e515760405162461bcd60e51b815260206004820152602960248201527f5b5145432d3032353030385d2d5468652061756374696f6e206973206e6f74206044820152683334b734b9b432b21760b91b60648201526084016101c7565b805460ff1615610ebf5760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3032353030395d2d5468652061756374696f6e2068617320616c7260448201527232b0b23c903132b2b71032bc32b1baba32b21760691b60648201526084016101c7565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271610f7760018054610ef490611878565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2090611878565b8015610f6d5780601f10610f4257610100808354040283529160200191610f6d565b820191906000526020600020905b815481529060010190602001808311610f5057829003601f168201915b5050505050611526565b6040518263ffffffff1660e01b8152600401610f939190611806565b60206040518083038186803b158015610fab57600080fd5b505afa158015610fbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe39190611839565b8254600284015460405163a9059cbb60e01b81526001600160a01b0361010090930483166004820152602481019190915291925082169063a9059cbb90604401602060405180830381600087803b15801561103d57600080fd5b505af1158015611051573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110759190611856565b508154600160ff199091161782556000805460408051606081019091526025808252620100009092046001600160a01b031691633fb902719190611a7d60208301396040518263ffffffff1660e01b81526004016110d39190611806565b60206040518083038186803b1580156110eb57600080fd5b505afa1580156110ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111239190611839565b6001600160a01b0316836001015460405160006040518083038185875af1925050503d8060008114611171576040519150601f19603f3d011682016040523d82523d6000602084013e611176565b606091505b50509050806111ff5760405162461bcd60e51b815260206004820152604960248201527f5145432d3032353031302d4661696c656420746f207472616e7366657220746860448201527f6520686967686573742062696420666f7220636f6d6d756e69747920646973746064820152683934b13aba34b7b71760b91b608482015260a4016101c7565b82546040805188815260ff8316151560208201526001600160a01b03600884901c8116928201929092526001860154606082015260028601546080820152600386015460a082015261010090920416907f2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b9060c00160405180910390a250600195945050505050565b60008061129361140f565b60405162498bff60e81b815260206004820152602160248201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6044820152601d60fa1b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b15801561130b57600080fd5b505afa15801561131f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134391906118ad565b600084815260036020526040812060010154919250816b033b2e3c9fd0803ce8000000611370858561190f565b61137a919061192e565b61138491906118f7565b90508181141561139c5780611398816118dc565b9150505b949350505050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946113f99493879201611950565b6040516020818303038152906040529050919050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611aa2602291396040518263ffffffff1660e01b81526004016114699190611806565b60206040518083038186803b15801561148157600080fd5b505afa158015611495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b99190611839565b905090565b60606114ec826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b81525061157b565b92915050565b60606114ec826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b81525061157b565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946113f994938792016119bf565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020016115d39493929190611a25565b604051602081830303815290604052905092915050565b8280546115f690611878565b90600052602060002090601f016020900481019282611618576000855561165e565b82601f1061163157805160ff191683800117855561165e565b8280016001018555821561165e579182015b8281111561165e578251825591602001919060010190611643565b5061166a92915061166e565b5090565b5b8082111561166a576000815560010161166f565b60006020828403121561169557600080fd5b5035919050565b6001600160a01b03811681146116b157600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156116dd57600080fd5b82356116e88161169c565b9150602083013567ffffffffffffffff8082111561170557600080fd5b818501915085601f83011261171957600080fd5b81358181111561172b5761172b6116b4565b604051601f8201601f19908116603f01168101908382118183101715611753576117536116b4565b8160405282815288602084870101111561176c57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60208082526028908201527f5b5145432d3032353030345d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60005b838110156117f15781810151838201526020016117d9565b83811115611800576000848401525b50505050565b60208152600082518060208401526118258160408501602087016117d6565b601f01601f19169190910160400192915050565b60006020828403121561184b57600080fd5b8151610cc68161169c565b60006020828403121561186857600080fd5b81518015158114610cc657600080fd5b600181811c9082168061188c57607f821691505b602082108114156104a257634e487b7160e01b600052602260045260246000fd5b6000602082840312156118bf57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006000198214156118f0576118f06118c6565b5060010190565b6000821982111561190a5761190a6118c6565b500190565b6000816000190483118215151615611929576119296118c6565b500290565b60008261194b57634e487b7160e01b600052601260045260246000fd5b500490565b60008551611962818460208a016117d6565b855190830190611976818360208a016117d6565b85519101906119898183602089016117d6565b845191019061199c8183602088016117d6565b6c73797374656d42616c616e636560981b9101908152600d019695505050505050565b600085516119d1818460208a016117d6565b8551908301906119e5818360208a016117d6565b85519101906119f88183602089016117d6565b8451910190611a0b8183602088016117d6565b6331b7b4b760e11b91019081526004019695505050505050565b60008551611a37818460208a016117d6565b855190830190611a4b818360208a016117d6565b8551910190611a5e8183602089016117d6565b8451910190611a718183602088016117d6565b01969550505050505056fe746f6b656e65636f6e6f6d6963732e64656661756c74416c6c6f636174696f6e50726f7879676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a26469706673582212209c36885d0ab3d170a903bce42384e584a733ae7cb55de63a92c31fd7a8667b6d64736f6c63430008090033",
}

// SystemSurplusAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemSurplusAuctionMetaData.ABI instead.
var SystemSurplusAuctionABI = SystemSurplusAuctionMetaData.ABI

// SystemSurplusAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemSurplusAuctionMetaData.Bin instead.
var SystemSurplusAuctionBin = SystemSurplusAuctionMetaData.Bin

// DeploySystemSurplusAuction deploys a new Ethereum contract, binding an instance of SystemSurplusAuction to it.
func DeploySystemSurplusAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemSurplusAuction, error) {
	parsed, err := SystemSurplusAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemSurplusAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// SystemSurplusAuction is an auto generated Go binding around an Ethereum contract.
type SystemSurplusAuction struct {
	SystemSurplusAuctionCaller     // Read-only binding to the contract
	SystemSurplusAuctionTransactor // Write-only binding to the contract
	SystemSurplusAuctionFilterer   // Log filterer for contract events
}

// SystemSurplusAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemSurplusAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSurplusAuctionSession struct {
	Contract     *SystemSurplusAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemSurplusAuctionCallerSession struct {
	Contract *SystemSurplusAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SystemSurplusAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemSurplusAuctionTransactorSession struct {
	Contract     *SystemSurplusAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemSurplusAuctionRaw struct {
	Contract *SystemSurplusAuction // Generic contract binding to access the raw methods on
}

// SystemSurplusAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCallerRaw struct {
	Contract *SystemSurplusAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SystemSurplusAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactorRaw struct {
	Contract *SystemSurplusAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemSurplusAuction creates a new instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuction(address common.Address, backend bind.ContractBackend) (*SystemSurplusAuction, error) {
	contract, err := bindSystemSurplusAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// NewSystemSurplusAuctionCaller creates a new read-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionCaller(address common.Address, caller bind.ContractCaller) (*SystemSurplusAuctionCaller, error) {
	contract, err := bindSystemSurplusAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionCaller{contract: contract}, nil
}

// NewSystemSurplusAuctionTransactor creates a new write-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemSurplusAuctionTransactor, error) {
	contract, err := bindSystemSurplusAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionTransactor{contract: contract}, nil
}

// NewSystemSurplusAuctionFilterer creates a new log filterer instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemSurplusAuctionFilterer, error) {
	contract, err := bindSystemSurplusAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionFilterer{contract: contract}, nil
}

// bindSystemSurplusAuction binds a generic wrapper to an already deployed contract.
func bindSystemSurplusAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemSurplusAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		IsExecuted bool
		Bidder     common.Address
		HighestBid *big.Int
		Lot        *big.Int
		EndTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsExecuted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) AuctionsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctionsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "getRaisingBid", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "bid", _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Execute(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "execute", _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// SystemSurplusAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStartedIterator struct {
	Event *SystemSurplusAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionAuctionStarted)
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
		it.Event = new(SystemSurplusAuctionAuctionStarted)
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
func (it *SystemSurplusAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionAuctionStarted represents a AuctionStarted event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStarted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionAuctionStartedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionAuctionStartedIterator{contract: _SystemSurplusAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionAuctionStarted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionAuctionStarted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseAuctionStarted(log types.Log) (*SystemSurplusAuctionAuctionStarted, error) {
	event := new(SystemSurplusAuctionAuctionStarted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemSurplusAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBidIterator struct {
	Event *SystemSurplusAuctionBid // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionBid)
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
		it.Event = new(SystemSurplusAuctionBid)
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
func (it *SystemSurplusAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionBid represents a Bid event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterBid(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemSurplusAuctionBidIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionBidIterator{contract: _SystemSurplusAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionBid, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionBid)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseBid(log types.Log) (*SystemSurplusAuctionBid, error) {
	event := new(SystemSurplusAuctionBid)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemSurplusAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecutedIterator struct {
	Event *SystemSurplusAuctionExecuted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionExecuted)
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
		it.Event = new(SystemSurplusAuctionExecuted)
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
func (it *SystemSurplusAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionExecuted represents a Executed event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecuted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Info      SystemSurplusAuctionAuctionInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionExecutedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionExecutedIterator{contract: _SystemSurplusAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionExecuted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionExecuted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseExecuted(log types.Log) (*SystemSurplusAuctionExecuted, error) {
	event := new(SystemSurplusAuctionExecuted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
