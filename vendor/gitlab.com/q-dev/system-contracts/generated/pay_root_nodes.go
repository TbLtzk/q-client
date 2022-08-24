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

// RootNodeRewardProxyMetaData contains all meta data concerning the RootNodeRewardProxy contract.
var RootNodeRewardProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061066a806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057600080fd5b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610425565b610333565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb90271916100cd91600401610449565b60206040518083038186803b1580156100e557600080fd5b505afa1580156100f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011d91906104ae565b6001600160a01b0316639eab52536040518163ffffffff1660e01b815260040160006040518083038186803b15801561015557600080fd5b505afa158015610169573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261019191908101906104e1565b905060008151476101a291906105a5565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb902719161020c9190600401610449565b60206040518083038186803b15801561022457600080fd5b505afa158015610238573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025c91906104ae565b90506000805b8451816001600160401b0316101561032c5784816001600160401b03168151811061028f5761028f6105c7565b6020908102919091010151604051632377789f60e21b81526001600160a01b03808316600483015291935090841690638ddde27c9086906024016020604051808303818588803b1580156102e257600080fd5b505af11580156102f6573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061031b91906105dd565b50610325816105ff565b9050610262565b5050505050565b600054610100900460ff168061034c575060005460ff16155b6103b35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff161580156103d5576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610409576000805461ff00191690555b5050565b6001600160a01b038116811461042257600080fd5b50565b60006020828403121561043757600080fd5b81356104428161040d565b9392505050565b600060208083528351808285015260005b818110156104765785810183015185820160400152820161045a565b81811115610488576000604083870101525b50601f01601f1916929092016040019392505050565b80516104a98161040d565b919050565b6000602082840312156104c057600080fd5b81516104428161040d565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156104f457600080fd5b82516001600160401b038082111561050b57600080fd5b818501915085601f83011261051f57600080fd5b815181811115610531576105316104cb565b8060051b604051601f19603f83011681018181108582111715610556576105566104cb565b60405291825284820192508381018501918883111561057457600080fd5b938501935b828510156105995761058a8561049e565b84529385019392850192610579565b98975050505050505050565b6000826105c257634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156105ef57600080fd5b8151801515811461044257600080fd5b60006001600160401b038083168181141561062a57634e487b7160e01b600052601160045260246000fd5b600101939250505056fea2646970667358221220de9a52c48babd2c5a77017efc9975fc1f42c565076107ef1da118156efec75c764736f6c63430008090033",
}

// RootNodeRewardProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use RootNodeRewardProxyMetaData.ABI instead.
var RootNodeRewardProxyABI = RootNodeRewardProxyMetaData.ABI

// RootNodeRewardProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RootNodeRewardProxyMetaData.Bin instead.
var RootNodeRewardProxyBin = RootNodeRewardProxyMetaData.Bin

// DeployRootNodeRewardProxy deploys a new Ethereum contract, binding an instance of RootNodeRewardProxy to it.
func DeployRootNodeRewardProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootNodeRewardProxy, error) {
	parsed, err := RootNodeRewardProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RootNodeRewardProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootNodeRewardProxy{RootNodeRewardProxyCaller: RootNodeRewardProxyCaller{contract: contract}, RootNodeRewardProxyTransactor: RootNodeRewardProxyTransactor{contract: contract}, RootNodeRewardProxyFilterer: RootNodeRewardProxyFilterer{contract: contract}}, nil
}

// RootNodeRewardProxy is an auto generated Go binding around an Ethereum contract.
type RootNodeRewardProxy struct {
	RootNodeRewardProxyCaller     // Read-only binding to the contract
	RootNodeRewardProxyTransactor // Write-only binding to the contract
	RootNodeRewardProxyFilterer   // Log filterer for contract events
}

// RootNodeRewardProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootNodeRewardProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootNodeRewardProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootNodeRewardProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootNodeRewardProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootNodeRewardProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootNodeRewardProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootNodeRewardProxySession struct {
	Contract     *RootNodeRewardProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootNodeRewardProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootNodeRewardProxyCallerSession struct {
	Contract *RootNodeRewardProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RootNodeRewardProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootNodeRewardProxyTransactorSession struct {
	Contract     *RootNodeRewardProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RootNodeRewardProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootNodeRewardProxyRaw struct {
	Contract *RootNodeRewardProxy // Generic contract binding to access the raw methods on
}

// RootNodeRewardProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootNodeRewardProxyCallerRaw struct {
	Contract *RootNodeRewardProxyCaller // Generic read-only contract binding to access the raw methods on
}

// RootNodeRewardProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootNodeRewardProxyTransactorRaw struct {
	Contract *RootNodeRewardProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootNodeRewardProxy creates a new instance of RootNodeRewardProxy, bound to a specific deployed contract.
func NewRootNodeRewardProxy(address common.Address, backend bind.ContractBackend) (*RootNodeRewardProxy, error) {
	contract, err := bindRootNodeRewardProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootNodeRewardProxy{RootNodeRewardProxyCaller: RootNodeRewardProxyCaller{contract: contract}, RootNodeRewardProxyTransactor: RootNodeRewardProxyTransactor{contract: contract}, RootNodeRewardProxyFilterer: RootNodeRewardProxyFilterer{contract: contract}}, nil
}

// NewRootNodeRewardProxyCaller creates a new read-only instance of RootNodeRewardProxy, bound to a specific deployed contract.
func NewRootNodeRewardProxyCaller(address common.Address, caller bind.ContractCaller) (*RootNodeRewardProxyCaller, error) {
	contract, err := bindRootNodeRewardProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootNodeRewardProxyCaller{contract: contract}, nil
}

// NewRootNodeRewardProxyTransactor creates a new write-only instance of RootNodeRewardProxy, bound to a specific deployed contract.
func NewRootNodeRewardProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*RootNodeRewardProxyTransactor, error) {
	contract, err := bindRootNodeRewardProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootNodeRewardProxyTransactor{contract: contract}, nil
}

// NewRootNodeRewardProxyFilterer creates a new log filterer instance of RootNodeRewardProxy, bound to a specific deployed contract.
func NewRootNodeRewardProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*RootNodeRewardProxyFilterer, error) {
	contract, err := bindRootNodeRewardProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootNodeRewardProxyFilterer{contract: contract}, nil
}

// bindRootNodeRewardProxy binds a generic wrapper to an already deployed contract.
func bindRootNodeRewardProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootNodeRewardProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootNodeRewardProxy *RootNodeRewardProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootNodeRewardProxy.Contract.RootNodeRewardProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootNodeRewardProxy *RootNodeRewardProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.RootNodeRewardProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootNodeRewardProxy *RootNodeRewardProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.RootNodeRewardProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootNodeRewardProxy *RootNodeRewardProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RootNodeRewardProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.contract.Transact(opts, method, params...)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootNodeRewardProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_RootNodeRewardProxy *RootNodeRewardProxySession) Allocate() (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Allocate(&_RootNodeRewardProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Allocate(&_RootNodeRewardProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _RootNodeRewardProxy.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_RootNodeRewardProxy *RootNodeRewardProxySession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Initialize(&_RootNodeRewardProxy.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Initialize(&_RootNodeRewardProxy.TransactOpts, _registry)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootNodeRewardProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootNodeRewardProxy *RootNodeRewardProxySession) Receive() (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Receive(&_RootNodeRewardProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RootNodeRewardProxy *RootNodeRewardProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _RootNodeRewardProxy.Contract.Receive(&_RootNodeRewardProxy.TransactOpts)
}
