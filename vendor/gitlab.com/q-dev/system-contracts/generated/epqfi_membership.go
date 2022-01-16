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

// EPQFIMembershipMetaData contains all meta data concerning the EPQFIMembership contract.
var EPQFIMembershipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertLimitKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertVotingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"experts\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldExp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newExp\",\"type\":\"address\"}],\"name\":\"swapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_experts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506115a2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630b1ca49a146100bf5780633e021814146100d45780635313263a146100f2578063715018a6146100fa5780638da5cb5b14610102578063946d9204146101175780639eab52531461012a578063a230c5241461013f578063b278982d14610162578063b295a00e14610175578063c68368251461018b578063ca6d56dc1461019e578063de8fa431146101b1578063f2fde38b146101b9575b600080fd5b6100d26100cd36600461102b565b6101cc565b005b6100dc61030c565b6040516100e9919061104f565b60405180910390f35b6100dc61039a565b6100d26103a7565b61010a6103e2565b6040516100e991906110a4565b6100d2610125366004611121565b6103f1565b610132610472565b6040516100e991906111d4565b61015261014d36600461102b565b6104f8565b60405190151581526020016100e9565b6100d2610170366004611221565b61057f565b61017d610739565b6040519081526020016100e9565b60655461010a906001600160a01b031681565b6100d26101ac36600461102b565b6108db565b61017d610af8565b6100d26101c736600461102b565b610b75565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb90271906101fd90606890600401611295565b60206040518083038186803b15801561021557600080fd5b505afa158015610229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024d919061133d565b6001600160a01b0316336001600160a01b0316146102865760405162461bcd60e51b815260040161027d9061135a565b60405180910390fd5b606554604051631484968760e11b81526001600160a01b03909116906329092d0e906102b69084906004016110a4565b602060405180830381600087803b1580156102d057600080fd5b505af11580156102e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030891906113ca565b5050565b606780546103199061125a565b80601f01602080910402602001604051908101604052809291908181526020018280546103459061125a565b80156103925780601f1061036757610100808354040283529160200191610392565b820191906000526020600020905b81548152906001019060200180831161037557829003601f168201915b505050505081565b606880546103199061125a565b336103b06103e2565b6001600160a01b0316146103d65760405162461bcd60e51b815260040161027d906113ec565b6103e06000610c15565b565b6033546001600160a01b031690565b6103fb8282610c67565b60408051808201909152601e8082527f636f6e737469747574696f6e2e45505146492e6d61784e457870657274730000602090920191825261043f91606791610f7d565b5060405180606001604052806029815260200161154460299139805161046d91606891602090910190610f7d565b505050565b606554604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b1580156104b757600080fd5b505afa1580156104cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104f39190810190611421565b905090565b606554604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e8906105299085906004016110a4565b60206040518083038186803b15801561054157600080fd5b505afa158015610555573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057991906113ca565b92915050565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb90271906105b090606890600401611295565b60206040518083038186803b1580156105c857600080fd5b505afa1580156105dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610600919061133d565b6001600160a01b0316336001600160a01b0316146106305760405162461bcd60e51b815260040161027d9061135a565b606554604051631484968760e11b81526001600160a01b039091169081906329092d0e906106629086906004016110a4565b602060405180830381600087803b15801561067c57600080fd5b505af1158015610690573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b491906113ca565b50604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f906106e19085906004016110a4565b602060405180830381600087803b1580156106fb57600080fd5b505af115801561070f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073391906113ca565b50505050565b6066546040805160608101909152602280825260009283926001600160a01b0390911691633fb90271919061152260208301396040518263ffffffff1660e01b8152600401610788919061104f565b60206040518083038186803b1580156107a057600080fd5b505afa1580156107b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d8919061133d565b6001600160a01b031663498bff0060676040518263ffffffff1660e01b81526004016108049190611295565b60206040518083038186803b15801561081c57600080fd5b505afa158015610830573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085491906114ba565b9050600081116108d65760405162461bcd60e51b815260206004820152604160248201527f5b5145432d3030373030315d2d457870657274206c696d697420706172616d6560448201527f74657220697320696e76616c6964206f7220646f6573206e6f742065786973746064820152601760f91b608482015260a40161027d565b919050565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb902719061090c90606890600401611295565b60206040518083038186803b15801561092457600080fd5b505afa158015610938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095c919061133d565b6001600160a01b0316336001600160a01b03161461098c5760405162461bcd60e51b815260040161027d9061135a565b6065546001600160a01b03166109a0610739565b816001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d957600080fd5b505afa1580156109ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1191906114ba565b1115610a7a5760405162461bcd60e51b815260206004820152603260248201527f5b5145432d3030373030305d2d4d6178696d756d206e756d626572206f6620656044820152713c3832b93a399034b9903932b0b1b432b21760711b606482015260840161027d565b604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610aa69085906004016110a4565b602060405180830381600087803b158015610ac057600080fd5b505af1158015610ad4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046d91906113ca565b6065546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d916004808301926020929190829003018186803b158015610b3d57600080fd5b505afa158015610b51573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f391906114ba565b33610b7e6103e2565b6001600160a01b031614610ba45760405162461bcd60e51b815260040161027d906113ec565b6001600160a01b038116610c095760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161027d565b610c1281610c15565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610c80575060005460ff16155b610c9c5760405162461bcd60e51b815260040161027d906114d3565b600054610100900460ff16158015610cbe576000805461ffff19166101011790555b606680546001600160a01b0319166001600160a01b038516908117909155604080518082018252601d81527f636f6d6d6f6e2e666163746f72792e6164647265737353746f7261676500000060208201529051633fb9027160e01b8152633fb9027191610d2d9160040161104f565b60206040518083038186803b158015610d4557600080fd5b505afa158015610d59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7d919061133d565b6001600160a01b03166362c54dbd836040518263ffffffff1660e01b8152600401610da891906111d4565b602060405180830381600087803b158015610dc257600080fd5b505af1158015610dd6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfa919061133d565b606580546001600160a01b0319166001600160a01b0392909216919091179055610e22610e38565b801561046d576000805461ff0019169055505050565b600054610100900460ff1680610e51575060005460ff16155b610e6d5760405162461bcd60e51b815260040161027d906114d3565b600054610100900460ff16158015610e8f576000805461ffff19166101011790555b610e97610eb3565b610e9f610f1d565b8015610c12576000805461ff001916905550565b600054610100900460ff1680610ecc575060005460ff16155b610ee85760405162461bcd60e51b815260040161027d906114d3565b600054610100900460ff16158015610e9f576000805461ffff19166101011790558015610c12576000805461ff001916905550565b600054610100900460ff1680610f36575060005460ff16155b610f525760405162461bcd60e51b815260040161027d906114d3565b600054610100900460ff16158015610f74576000805461ffff19166101011790555b610e9f33610c15565b828054610f899061125a565b90600052602060002090601f016020900481019282610fab5760008555610ff1565b82601f10610fc457805160ff1916838001178555610ff1565b82800160010185558215610ff1579182015b82811115610ff1578251825591602001919060010190610fd6565b50610ffd929150611001565b5090565b5b80821115610ffd5760008155600101611002565b6001600160a01b0381168114610c1257600080fd5b60006020828403121561103d57600080fd5b813561104881611016565b9392505050565b600060208083528351808285015260005b8181101561107c57858101830151858201604001528201611060565b8181111561108e576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b0391909116815260200190565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156110f6576110f66110b8565b604052919050565b60006001600160401b03821115611117576111176110b8565b5060051b60200190565b6000806040838503121561113457600080fd5b823561113f81611016565b91506020838101356001600160401b0381111561115b57600080fd5b8401601f8101861361116c57600080fd5b803561117f61117a826110fe565b6110ce565b81815260059190911b8201830190838101908883111561119e57600080fd5b928401925b828410156111c55783356111b681611016565b825292840192908401906111a3565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156112155783516001600160a01b0316835292840192918401916001016111f0565b50909695505050505050565b6000806040838503121561123457600080fd5b823561123f81611016565b9150602083013561124f81611016565b809150509250929050565b600181811c9082168061126e57607f821691505b6020821081141561128f57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208083526000845481600182811c9150808316806112b757607f831692505b8583108114156112d557634e487b7160e01b85526022600452602485fd5b8786018381526020018180156112f257600181146113035761132e565b60ff1986168252878201965061132e565b60008b81526020902060005b868110156113285781548482015290850190890161130f565b83019750505b50949998505050505050505050565b60006020828403121561134f57600080fd5b815161104881611016565b6020808252604a908201527f5b5145432d3030373030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920657870657274732066726f6d20746869732070616e656c2068616060820152693b329030b1b1b2b9b99760b11b608082015260a00190565b6000602082840312156113dc57600080fd5b8151801515811461104857600080fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000602080838503121561143457600080fd5b82516001600160401b0381111561144a57600080fd5b8301601f8101851361145b57600080fd5b805161146961117a826110fe565b81815260059190911b8201830190838101908783111561148857600080fd5b928401925b828410156114af5783516114a081611016565b8252928401929084019061148d565b979650505050505050565b6000602082840312156114cc57600080fd5b5051919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b60608201526080019056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273676f7665726e616e63652e657870657274732e45505146492e6d656d62657273686970566f74696e67a2646970667358221220696e1af54b4ecfc9777de44f9715d822cd914faf0628cddf9b105070e81ee42564736f6c63430008090033",
}

// EPQFIMembershipABI is the input ABI used to generate the binding from.
// Deprecated: Use EPQFIMembershipMetaData.ABI instead.
var EPQFIMembershipABI = EPQFIMembershipMetaData.ABI

// EPQFIMembershipBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EPQFIMembershipMetaData.Bin instead.
var EPQFIMembershipBin = EPQFIMembershipMetaData.Bin

// DeployEPQFIMembership deploys a new Ethereum contract, binding an instance of EPQFIMembership to it.
func DeployEPQFIMembership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EPQFIMembership, error) {
	parsed, err := EPQFIMembershipMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EPQFIMembershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EPQFIMembership{EPQFIMembershipCaller: EPQFIMembershipCaller{contract: contract}, EPQFIMembershipTransactor: EPQFIMembershipTransactor{contract: contract}, EPQFIMembershipFilterer: EPQFIMembershipFilterer{contract: contract}}, nil
}

// EPQFIMembership is an auto generated Go binding around an Ethereum contract.
type EPQFIMembership struct {
	EPQFIMembershipCaller     // Read-only binding to the contract
	EPQFIMembershipTransactor // Write-only binding to the contract
	EPQFIMembershipFilterer   // Log filterer for contract events
}

// EPQFIMembershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPQFIMembershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPQFIMembershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPQFIMembershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPQFIMembershipSession struct {
	Contract     *EPQFIMembership  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPQFIMembershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPQFIMembershipCallerSession struct {
	Contract *EPQFIMembershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EPQFIMembershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPQFIMembershipTransactorSession struct {
	Contract     *EPQFIMembershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EPQFIMembershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPQFIMembershipRaw struct {
	Contract *EPQFIMembership // Generic contract binding to access the raw methods on
}

// EPQFIMembershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPQFIMembershipCallerRaw struct {
	Contract *EPQFIMembershipCaller // Generic read-only contract binding to access the raw methods on
}

// EPQFIMembershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPQFIMembershipTransactorRaw struct {
	Contract *EPQFIMembershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPQFIMembership creates a new instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembership(address common.Address, backend bind.ContractBackend) (*EPQFIMembership, error) {
	contract, err := bindEPQFIMembership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembership{EPQFIMembershipCaller: EPQFIMembershipCaller{contract: contract}, EPQFIMembershipTransactor: EPQFIMembershipTransactor{contract: contract}, EPQFIMembershipFilterer: EPQFIMembershipFilterer{contract: contract}}, nil
}

// NewEPQFIMembershipCaller creates a new read-only instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipCaller(address common.Address, caller bind.ContractCaller) (*EPQFIMembershipCaller, error) {
	contract, err := bindEPQFIMembership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipCaller{contract: contract}, nil
}

// NewEPQFIMembershipTransactor creates a new write-only instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipTransactor(address common.Address, transactor bind.ContractTransactor) (*EPQFIMembershipTransactor, error) {
	contract, err := bindEPQFIMembership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipTransactor{contract: contract}, nil
}

// NewEPQFIMembershipFilterer creates a new log filterer instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipFilterer(address common.Address, filterer bind.ContractFilterer) (*EPQFIMembershipFilterer, error) {
	contract, err := bindEPQFIMembership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipFilterer{contract: contract}, nil
}

// bindEPQFIMembership binds a generic wrapper to an already deployed contract.
func bindEPQFIMembership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EPQFIMembershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPQFIMembership *EPQFIMembershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPQFIMembership.Contract.EPQFIMembershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPQFIMembership *EPQFIMembershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.EPQFIMembershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPQFIMembership *EPQFIMembershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.EPQFIMembershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPQFIMembership *EPQFIMembershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPQFIMembership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPQFIMembership *EPQFIMembershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPQFIMembership *EPQFIMembershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.contract.Transact(opts, method, params...)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCaller) ExpertLimitKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "expertLimitKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipSession) ExpertLimitKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertLimitKey(&_EPQFIMembership.CallOpts)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCallerSession) ExpertLimitKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertLimitKey(&_EPQFIMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCaller) ExpertVotingKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "expertVotingKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipSession) ExpertVotingKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertVotingKey(&_EPQFIMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCallerSession) ExpertVotingKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertVotingKey(&_EPQFIMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCaller) Experts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "experts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipSession) Experts() (common.Address, error) {
	return _EPQFIMembership.Contract.Experts(&_EPQFIMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCallerSession) Experts() (common.Address, error) {
	return _EPQFIMembership.Contract.Experts(&_EPQFIMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipSession) GetLimit() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetLimit(&_EPQFIMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetLimit() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetLimit(&_EPQFIMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipSession) GetMembers() ([]common.Address, error) {
	return _EPQFIMembership.Contract.GetMembers(&_EPQFIMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetMembers() ([]common.Address, error) {
	return _EPQFIMembership.Contract.GetMembers(&_EPQFIMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipSession) GetSize() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetSize(&_EPQFIMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetSize() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetSize(&_EPQFIMembership.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipCaller) IsMember(opts *bind.CallOpts, _exp common.Address) (bool, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "isMember", _exp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipSession) IsMember(_exp common.Address) (bool, error) {
	return _EPQFIMembership.Contract.IsMember(&_EPQFIMembership.CallOpts, _exp)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipCallerSession) IsMember(_exp common.Address) (bool, error) {
	return _EPQFIMembership.Contract.IsMember(&_EPQFIMembership.CallOpts, _exp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipSession) Owner() (common.Address, error) {
	return _EPQFIMembership.Contract.Owner(&_EPQFIMembership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCallerSession) Owner() (common.Address, error) {
	return _EPQFIMembership.Contract.Owner(&_EPQFIMembership.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) AddMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "addMember", _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.AddMember(&_EPQFIMembership.TransactOpts, _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.AddMember(&_EPQFIMembership.TransactOpts, _exp)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "initialize", _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.Initialize(&_EPQFIMembership.TransactOpts, _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.Initialize(&_EPQFIMembership.TransactOpts, _registry, _experts)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) RemoveMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "removeMember", _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RemoveMember(&_EPQFIMembership.TransactOpts, _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RemoveMember(&_EPQFIMembership.TransactOpts, _exp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RenounceOwnership(&_EPQFIMembership.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RenounceOwnership(&_EPQFIMembership.TransactOpts)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) SwapMember(opts *bind.TransactOpts, _oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "swapMember", _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.SwapMember(&_EPQFIMembership.TransactOpts, _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.SwapMember(&_EPQFIMembership.TransactOpts, _oldExp, _newExp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.TransferOwnership(&_EPQFIMembership.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.TransferOwnership(&_EPQFIMembership.TransactOpts, newOwner)
}

// EPQFIMembershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPQFIMembership contract.
type EPQFIMembershipOwnershipTransferredIterator struct {
	Event *EPQFIMembershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPQFIMembershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPQFIMembershipOwnershipTransferred)
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
		it.Event = new(EPQFIMembershipOwnershipTransferred)
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
func (it *EPQFIMembershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPQFIMembershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPQFIMembershipOwnershipTransferred represents a OwnershipTransferred event raised by the EPQFIMembership contract.
type EPQFIMembershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPQFIMembership *EPQFIMembershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPQFIMembershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPQFIMembership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipOwnershipTransferredIterator{contract: _EPQFIMembership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPQFIMembership *EPQFIMembershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPQFIMembershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPQFIMembership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPQFIMembershipOwnershipTransferred)
				if err := _EPQFIMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EPQFIMembership *EPQFIMembershipFilterer) ParseOwnershipTransferred(log types.Log) (*EPQFIMembershipOwnershipTransferred, error) {
	event := new(EPQFIMembershipOwnershipTransferred)
	if err := _EPQFIMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
