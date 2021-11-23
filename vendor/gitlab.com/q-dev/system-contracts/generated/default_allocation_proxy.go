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

// DefaultAllocationProxyMetaData contains all meta data concerning the DefaultAllocationProxy contract.
var DefaultAllocationProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beneficiaries\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_beneficiaries\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shares\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e35806100206000396000f3fe6080604052600436106100435760003560e01c806357a858fc1461004f578063abaa991614610085578063efe86a5e1461009c578063efeb5e58146100bc5761004a565b3661004a57005b600080fd5b34801561005b57600080fd5b5061006f61006a366004610c01565b6100dc565b60405161007c9190610c34565b60405180910390f35b34801561009157600080fd5b5061009a610185565b005b3480156100a857600080fd5b5061009a6100b7366004610b8f565b610682565b3480156100c857600080fd5b5061006f6100d7366004610c01565b610792565b600281815481106100ec57600080fd5b600091825260209182902001805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529350909183018282801561017d5780601f106101525761010080835404028352916020019161017d565b820191906000526020600020905b81548152906001019060200180831161016057829003601f168201915b505050505081565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610dde602291396040518263ffffffff1660e01b81526004016101df9190610c34565b60206040518083038186803b1580156101f757600080fd5b505afa15801561020b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022f9190610b73565b905060006002805480602002602001604051908101604052809291908181526020016000905b828210156103005760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156102ec5780601f106102c1576101008083540402835291602001916102ec565b820191906000526020600020905b8154815290600101906020018083116102cf57829003601f168201915b505050505081526020019060010190610255565b505050509050600081516001600160401b038111801561031f57600080fd5b50604051908082528060200260200182016040528015610349578160200160208202803683370190505b5090506000805b835181101561042757846001600160a01b031663498bff0085838151811061037457fe5b60200260200101516040518263ffffffff1660e01b81526004016103989190610c34565b60206040518083038186803b1580156103b057600080fd5b505afa1580156103c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e89190610c19565b8382815181106103f457fe5b60200260200101818152505061041d8284838151811061041057fe5b60200260200101516107a2565b9150600101610350565b5060006001805480602002602001604051908101604052809291908181526020016000905b828210156104f75760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156104e35780601f106104b8576101008083540402835291602001916104e3565b820191906000526020600020905b8154815290600101906020018083116104c657829003601f168201915b50505050508152602001906001019061044c565b50505050905060008047905060005b83518110156106785760006105418661053b89858151811061052457fe5b60200260200101518661080390919063ffffffff16565b9061085c565b9050600060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027186848151811061057357fe5b60200260200101516040518263ffffffff1660e01b81526004016105979190610c34565b60206040518083038186803b1580156105af57600080fd5b505afa1580156105c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e79190610b73565b93506000846001600160a01b03168260405161060290610c31565b60006040518083038185875af1925050503d806000811461063f576040519150601f19603f3d011682016040523d82523d6000602084013e610644565b606091505b505090508061066e5760405162461bcd60e51b815260040161066590610c87565b60405180910390fd5b5050600101610506565b5050505050505050565b600054610100900460ff168061069b575061069b61089b565b806106a9575060005460ff16155b6106e45760405162461bcd60e51b815260040180806020018281038252602e815260200180610d8f602e913960400191505060405180910390fd5b600054610100900460ff1615801561070f576000805460ff1961ff0019909116610100171660011790555b81518351146107305760405162461bcd60e51b815260040161066590610ce4565b8251610743906001906020860190610954565b508151610757906002906020850190610954565b506000805462010000600160b01b031916620100006001600160a01b03871602179055801561078c576000805461ff00191690555b50505050565b600181815481106100ec57600080fd5b6000828201838110156107fa576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b90505b92915050565b600082610812575060006107fd565b8282028284828161081f57fe5b04146107fa5760405162461bcd60e51b8152600401808060200182810382526021815260200180610dbd6021913960400191505060405180910390fd5b60006107fa83836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506108ac565b60006108a63061094e565b15905090565b600081836109385760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156108fd5781810151838201526020016108e5565b50505050905090810190601f16801561092a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161094457fe5b0495945050505050565b3b151590565b8280548282559060005260206000209081019282156109a1579160200282015b828111156109a157825180516109919184916020909101906109b1565b5091602001919060010190610974565b506109ad929150610a39565b5090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826109e75760008555610a2d565b82601f10610a0057805160ff1916838001178555610a2d565b82800160010185558215610a2d579182015b82811115610a2d578251825591602001919060010190610a12565b506109ad929150610a56565b808211156109ad576000610a4d8282610a6b565b50600101610a39565b5b808211156109ad5760008155600101610a57565b50805460018160011615610100020316600290046000825580601f10610a915750610aaf565b601f016020900490600052602060002090810190610aaf9190610a56565b50565b6000601f8381840112610ac3578182fd5b823560206001600160401b0380831115610ad957fe5b610ae68283850201610d56565b83815282810190878401875b86811015610b645781358a018b603f820112610b0c57898afd5b86810135604087821115610b1c57fe5b610b2d828c01601f19168a01610d56565b8281528e82848601011115610b40578c8dfd5b828285018b83013791820189018c9052508552509285019290850190600101610af2565b50909998505050505050505050565b600060208284031215610b84578081fd5b81516107fa81610d79565b600080600060608486031215610ba3578182fd5b8335610bae81610d79565b925060208401356001600160401b0380821115610bc9578384fd5b610bd587838801610ab2565b93506040860135915080821115610bea578283fd5b50610bf786828701610ab2565b9150509250925092565b600060208284031215610c12578081fd5b5035919050565b600060208284031215610c2a578081fd5b5051919050565b90565b6000602080835283518082850152825b81811015610c6057858101830151858201604001528201610c44565b81811115610c715783604083870101525b50601f01601f1916929092016040019392505050565b6020808252603e908201527f5b5145432d3031313030315d2d4661696c656420746f207472616e736665722060408201527f74686520616d6f756e742c20616c6c6f636174696f6e206661696c65642e0000606082015260800190565b6020808252604c908201527f5b5145432d3031313030305d2d4e756d626572206f662062656e65666963696160408201527f726965732073686f756c64206265207468652073616d65206173206e756d626560608201526b391037b31039b430b932b99760a11b608082015260a00190565b6040518181016001600160401b0381118282101715610d7157fe5b604052919050565b6001600160a01b0381168114610aaf57600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220b60c162cf3a0c4b306df626c4675d65feb35971de110b0e116815a276c6d11db64736f6c63430007060033",
}

// DefaultAllocationProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use DefaultAllocationProxyMetaData.ABI instead.
var DefaultAllocationProxyABI = DefaultAllocationProxyMetaData.ABI

// DefaultAllocationProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DefaultAllocationProxyMetaData.Bin instead.
var DefaultAllocationProxyBin = DefaultAllocationProxyMetaData.Bin

// DeployDefaultAllocationProxy deploys a new Ethereum contract, binding an instance of DefaultAllocationProxy to it.
func DeployDefaultAllocationProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DefaultAllocationProxy, error) {
	parsed, err := DefaultAllocationProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DefaultAllocationProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DefaultAllocationProxy{DefaultAllocationProxyCaller: DefaultAllocationProxyCaller{contract: contract}, DefaultAllocationProxyTransactor: DefaultAllocationProxyTransactor{contract: contract}, DefaultAllocationProxyFilterer: DefaultAllocationProxyFilterer{contract: contract}}, nil
}

// DefaultAllocationProxy is an auto generated Go binding around an Ethereum contract.
type DefaultAllocationProxy struct {
	DefaultAllocationProxyCaller     // Read-only binding to the contract
	DefaultAllocationProxyTransactor // Write-only binding to the contract
	DefaultAllocationProxyFilterer   // Log filterer for contract events
}

// DefaultAllocationProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefaultAllocationProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefaultAllocationProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefaultAllocationProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefaultAllocationProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefaultAllocationProxySession struct {
	Contract     *DefaultAllocationProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DefaultAllocationProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefaultAllocationProxyCallerSession struct {
	Contract *DefaultAllocationProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DefaultAllocationProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefaultAllocationProxyTransactorSession struct {
	Contract     *DefaultAllocationProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DefaultAllocationProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefaultAllocationProxyRaw struct {
	Contract *DefaultAllocationProxy // Generic contract binding to access the raw methods on
}

// DefaultAllocationProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefaultAllocationProxyCallerRaw struct {
	Contract *DefaultAllocationProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DefaultAllocationProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefaultAllocationProxyTransactorRaw struct {
	Contract *DefaultAllocationProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefaultAllocationProxy creates a new instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxy(address common.Address, backend bind.ContractBackend) (*DefaultAllocationProxy, error) {
	contract, err := bindDefaultAllocationProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxy{DefaultAllocationProxyCaller: DefaultAllocationProxyCaller{contract: contract}, DefaultAllocationProxyTransactor: DefaultAllocationProxyTransactor{contract: contract}, DefaultAllocationProxyFilterer: DefaultAllocationProxyFilterer{contract: contract}}, nil
}

// NewDefaultAllocationProxyCaller creates a new read-only instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyCaller(address common.Address, caller bind.ContractCaller) (*DefaultAllocationProxyCaller, error) {
	contract, err := bindDefaultAllocationProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyCaller{contract: contract}, nil
}

// NewDefaultAllocationProxyTransactor creates a new write-only instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DefaultAllocationProxyTransactor, error) {
	contract, err := bindDefaultAllocationProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyTransactor{contract: contract}, nil
}

// NewDefaultAllocationProxyFilterer creates a new log filterer instance of DefaultAllocationProxy, bound to a specific deployed contract.
func NewDefaultAllocationProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DefaultAllocationProxyFilterer, error) {
	contract, err := bindDefaultAllocationProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefaultAllocationProxyFilterer{contract: contract}, nil
}

// bindDefaultAllocationProxy binds a generic wrapper to an already deployed contract.
func bindDefaultAllocationProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefaultAllocationProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAllocationProxy *DefaultAllocationProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.DefaultAllocationProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DefaultAllocationProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.contract.Transact(opts, method, params...)
}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCaller) Beneficiaries(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DefaultAllocationProxy.contract.Call(opts, &out, "beneficiaries", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Beneficiaries(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Beneficiaries(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Beneficiaries is a free data retrieval call binding the contract method 0xefeb5e58.
//
// Solidity: function beneficiaries(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerSession) Beneficiaries(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Beneficiaries(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCaller) Shares(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _DefaultAllocationProxy.contract.Call(opts, &out, "shares", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Shares(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Shares(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Shares is a free data retrieval call binding the contract method 0x57a858fc.
//
// Solidity: function shares(uint256 ) view returns(string)
func (_DefaultAllocationProxy *DefaultAllocationProxyCallerSession) Shares(arg0 *big.Int) (string, error) {
	return _DefaultAllocationProxy.Contract.Shares(&_DefaultAllocationProxy.CallOpts, arg0)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Allocate() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Allocate(&_DefaultAllocationProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Allocate(&_DefaultAllocationProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.Transact(opts, "initialize", _registry, _beneficiaries, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Initialize(_registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Initialize(&_DefaultAllocationProxy.TransactOpts, _registry, _beneficiaries, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _beneficiaries, string[] _shares) returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Initialize(_registry common.Address, _beneficiaries []string, _shares []string) (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Initialize(&_DefaultAllocationProxy.TransactOpts, _registry, _beneficiaries, _shares)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DefaultAllocationProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxySession) Receive() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Receive(&_DefaultAllocationProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DefaultAllocationProxy *DefaultAllocationProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _DefaultAllocationProxy.Contract.Receive(&_DefaultAllocationProxy.TransactOpts)
}
