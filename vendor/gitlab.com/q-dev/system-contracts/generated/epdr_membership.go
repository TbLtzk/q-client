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

// EPDRMembershipMetaData contains all meta data concerning the EPDRMembership contract.
var EPDRMembershipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertLimitKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertVotingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"experts\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldExp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newExp\",\"type\":\"address\"}],\"name\":\"swapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_experts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b61150f8061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630b1ca49a146100bf5780633e021814146100d45780635313263a146100f2578063715018a6146100fa5780638da5cb5b14610102578063946d9204146101175780639eab52531461012a578063a230c5241461013f578063b278982d1461015f578063b295a00e14610172578063c683682514610187578063ca6d56dc1461018f578063de8fa431146101a2578063f2fde38b146101aa575b600080fd5b6100d26100cd366004610f0c565b6101bd565b005b6100dc6102fc565b6040516100e99190611171565b60405180910390f35b6100dc61038a565b6100d26103e5565b61010a610499565b6040516100e99190611105565b6100d2610125366004610f83565b6104a8565b610132610529565b6040516100e99190611119565b61015261014d366004610f0c565b6105af565b6040516100e99190611166565b6100d261016d366004610f4b565b610636565b61017a6107ef565b6040516100e991906113f0565b61010a610910565b6100d261019d366004610f0c565b61091f565b61017a610af0565b6100d26101b8366004610f0c565b610b6d565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb90271906101ed9060049081016111c4565b60206040518083038186803b15801561020557600080fd5b505afa158015610219573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023d9190610f2f565b6001600160a01b0316336001600160a01b0316146102765760405162461bcd60e51b815260040161026d9061133e565b60405180910390fd5b600154604051631484968760e11b81526001600160a01b03909116906329092d0e906102a6908490600401611105565b602060405180830381600087803b1580156102c057600080fd5b505af11580156102d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f891906110cd565b5050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103825780601f1061035757610100808354040283529160200191610382565b820191906000526020600020905b81548152906001019060200180831161036557829003601f168201915b505050505081565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103825780601f1061035757610100808354040283529160200191610382565b6103ed610c77565b6000546001600160a01b0390811691161461044f576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6104b28282610c7b565b60408051808201909152601d8082527f636f6e737469747574696f6e2e455044522e6d61784e4578706572747300000060209092019182526104f691600391610e6b565b506040518060600160405280602881526020016114b260289139805161052491600491602090910190610e6b565b505050565b600154604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801561056e57600080fd5b505afa158015610582573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105aa9190810190611035565b905090565b600154604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e8906105e0908590600401611105565b60206040518083038186803b1580156105f857600080fd5b505afa15801561060c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063091906110cd565b92915050565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb90271906106669060049081016111c4565b60206040518083038186803b15801561067e57600080fd5b505afa158015610692573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b69190610f2f565b6001600160a01b0316336001600160a01b0316146106e65760405162461bcd60e51b815260040161026d9061133e565b600154604051631484968760e11b81526001600160a01b039091169081906329092d0e90610718908690600401611105565b602060405180830381600087803b15801561073257600080fd5b505af1158015610746573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076a91906110cd565b50604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610797908590600401611105565b602060405180830381600087803b1580156107b157600080fd5b505af11580156107c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e991906110cd565b50505050565b600254604051633fb9027160e01b815260009182916001600160a01b0390911690633fb9027190610822906004016113ae565b60206040518083038186803b15801561083a57600080fd5b505afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108729190610f2f565b6001600160a01b031663498bff0060036040518263ffffffff1660e01b815260040161089e91906111c4565b60206040518083038186803b1580156108b657600080fd5b505afa1580156108ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ee91906110ed565b9050600081116105aa5760405162461bcd60e51b815260040161026d906112d7565b6001546001600160a01b031681565b600254604051633fb9027160e01b81526001600160a01b0390911690633fb902719061094f9060049081016111c4565b60206040518083038186803b15801561096757600080fd5b505afa15801561097b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099f9190610f2f565b6001600160a01b0316336001600160a01b0316146109cf5760405162461bcd60e51b815260040161026d9061133e565b6001546001600160a01b03166109e36107ef565b816001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a1c57600080fd5b505afa158015610a30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5491906110ed565b1115610a725760405162461bcd60e51b815260040161026d9061124e565b604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610a9e908590600401611105565b602060405180830381600087803b158015610ab857600080fd5b505af1158015610acc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052491906110cd565b6001546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d916004808301926020929190829003018186803b158015610b3557600080fd5b505afa158015610b49573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105aa91906110ed565b610b75610c77565b6000546001600160a01b03908116911614610bd7576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b038116610c1c5760405162461bcd60e51b815260040180806020018281038252602681526020018061145e6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b600054600160a81b900460ff1680610c965750610c96610e65565b80610cab5750600054600160a01b900460ff16155b610ce65760405162461bcd60e51b815260040180806020018281038252602e815260200180611484602e913960400191505060405180910390fd5b600054600160a81b900460ff16158015610d1d576000805460ff60a01b1960ff60a81b19909116600160a81b1716600160a01b1790555b600280546001600160a01b0319166001600160a01b038581169190911791829055604051633fb9027160e01b8152911690633fb9027190610d60906004016112a0565b60206040518083038186803b158015610d7857600080fd5b505afa158015610d8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db09190610f2f565b6001600160a01b03166362c54dbd836040518263ffffffff1660e01b8152600401610ddb9190611119565b602060405180830381600087803b158015610df557600080fd5b505af1158015610e09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2d9190610f2f565b600180546001600160a01b0319166001600160a01b03929092169190911790558015610524576000805460ff60a81b19169055505050565b303b1590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282610ea15760008555610ee7565b82601f10610eba57805160ff1916838001178555610ee7565b82800160010185558215610ee7579182015b82811115610ee7578251825591602001919060010190610ecc565b50610ef3929150610ef7565b5090565b5b80821115610ef35760008155600101610ef8565b600060208284031215610f1d578081fd5b8135610f2881611445565b9392505050565b600060208284031215610f40578081fd5b8151610f2881611445565b60008060408385031215610f5d578081fd5b8235610f6881611445565b91506020830135610f7881611445565b809150509250929050565b60008060408385031215610f95578182fd5b8235610fa081611445565b91506020838101356001600160401b03811115610fbb578283fd5b8401601f81018613610fcb578283fd5b8035610fde610fd98261141c565b6113f9565b81815283810190838501858402850186018a1015610ffa578687fd5b8694505b8385101561102557803561101181611445565b835260019490940193918501918501610ffe565b5080955050505050509250929050565b60006020808385031215611047578182fd5b82516001600160401b0381111561105c578283fd5b8301601f8101851361106c578283fd5b805161107a610fd98261141c565b8181528381019083850185840285018601891015611096578687fd5b8694505b838510156110c15780516110ad81611445565b83526001949094019391850191850161109a565b50979650505050505050565b6000602082840312156110de578081fd5b81518015158114610f28578182fd5b6000602082840312156110fe578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101561115a5783516001600160a01b031683529284019291840191600101611135565b50909695505050505050565b901515815260200190565b6000602080835283518082850152825b8181101561119d57858101830151858201604001528201611181565b818111156111ae5783604083870101525b50601f01601f1916929092016040019392505050565b600060208083018184528285546001808216600081146111eb576001811461120957611241565b60028304607f16855260ff1983166040890152606088019350611241565b600283048086526112198a611439565b885b828110156112375781548b82016040015290840190880161121b565b8a01604001955050505b5091979650505050505050565b60208082526032908201527f5b5145432d3030373030305d2d4d6178696d756d206e756d626572206f6620656040820152713c3832b93a399034b9903932b0b1b432b21760711b606082015260800190565b6020808252601d908201527f636f6d6d6f6e2e666163746f72792e6164647265737353746f72616765000000604082015260600190565b60208082526041908201527f5b5145432d3030373030315d2d457870657274206c696d697420706172616d6560408201527f74657220697320696e76616c6964206f7220646f6573206e6f742065786973746060820152601760f91b608082015260a00190565b6020808252604a908201527f5b5145432d3030373030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920657870657274732066726f6d20746869732070616e656c2068616060820152693b329030b1b1b2b9b99760b11b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e636f6e737469747574696f6e2e706172616d657465604082015261727360f01b606082015260800190565b90815260200190565b6040518181016001600160401b038111828210171561141457fe5b604052919050565b60006001600160401b0382111561142f57fe5b5060209081020190565b60009081526020902090565b6001600160a01b038116811461145a57600080fd5b5056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564676f7665726e616e63652e657870657274732e455044522e6d656d62657273686970566f74696e67a26469706673582212203d615d923bf055e46d0ce4f96b09613b77a549d97e316bf26cfd42aca762907f64736f6c63430007060033",
}

// EPDRMembershipABI is the input ABI used to generate the binding from.
// Deprecated: Use EPDRMembershipMetaData.ABI instead.
var EPDRMembershipABI = EPDRMembershipMetaData.ABI

// EPDRMembershipBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EPDRMembershipMetaData.Bin instead.
var EPDRMembershipBin = EPDRMembershipMetaData.Bin

// DeployEPDRMembership deploys a new Ethereum contract, binding an instance of EPDRMembership to it.
func DeployEPDRMembership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EPDRMembership, error) {
	parsed, err := EPDRMembershipMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EPDRMembershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// EPDRMembership is an auto generated Go binding around an Ethereum contract.
type EPDRMembership struct {
	EPDRMembershipCaller     // Read-only binding to the contract
	EPDRMembershipTransactor // Write-only binding to the contract
	EPDRMembershipFilterer   // Log filterer for contract events
}

// EPDRMembershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPDRMembershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPDRMembershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPDRMembershipSession struct {
	Contract     *EPDRMembership   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPDRMembershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPDRMembershipCallerSession struct {
	Contract *EPDRMembershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EPDRMembershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPDRMembershipTransactorSession struct {
	Contract     *EPDRMembershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EPDRMembershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPDRMembershipRaw struct {
	Contract *EPDRMembership // Generic contract binding to access the raw methods on
}

// EPDRMembershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPDRMembershipCallerRaw struct {
	Contract *EPDRMembershipCaller // Generic read-only contract binding to access the raw methods on
}

// EPDRMembershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactorRaw struct {
	Contract *EPDRMembershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPDRMembership creates a new instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembership(address common.Address, backend bind.ContractBackend) (*EPDRMembership, error) {
	contract, err := bindEPDRMembership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// NewEPDRMembershipCaller creates a new read-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipCaller(address common.Address, caller bind.ContractCaller) (*EPDRMembershipCaller, error) {
	contract, err := bindEPDRMembership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipCaller{contract: contract}, nil
}

// NewEPDRMembershipTransactor creates a new write-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipTransactor(address common.Address, transactor bind.ContractTransactor) (*EPDRMembershipTransactor, error) {
	contract, err := bindEPDRMembership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipTransactor{contract: contract}, nil
}

// NewEPDRMembershipFilterer creates a new log filterer instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipFilterer(address common.Address, filterer bind.ContractFilterer) (*EPDRMembershipFilterer, error) {
	contract, err := bindEPDRMembership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipFilterer{contract: contract}, nil
}

// bindEPDRMembership binds a generic wrapper to an already deployed contract.
func bindEPDRMembership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EPDRMembershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.EPDRMembershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transact(opts, method, params...)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertLimitKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "expertLimitKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertVotingKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "expertVotingKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Experts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "experts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCallerSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCaller) IsMember(opts *bind.CallOpts, _exp common.Address) (bool, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "isMember", _exp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCallerSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) AddMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "addMember", _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "initialize", _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) RemoveMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "removeMember", _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) SwapMember(opts *bind.TransactOpts, _oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "swapMember", _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// EPDRMembershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferredIterator struct {
	Event *EPDRMembershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPDRMembershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPDRMembershipOwnershipTransferred)
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
		it.Event = new(EPDRMembershipOwnershipTransferred)
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
func (it *EPDRMembershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPDRMembershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPDRMembershipOwnershipTransferred represents a OwnershipTransferred event raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPDRMembershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipOwnershipTransferredIterator{contract: _EPDRMembership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPDRMembershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPDRMembershipOwnershipTransferred)
				if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) ParseOwnershipTransferred(log types.Log) (*EPDRMembershipOwnershipTransferred, error) {
	event := new(EPDRMembershipOwnershipTransferred)
	if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
