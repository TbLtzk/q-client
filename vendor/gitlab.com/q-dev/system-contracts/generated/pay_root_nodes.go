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
	Bin: "0x608060405234801561001057600080fd5b5061069b806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057610034565b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b36600461049b565b6102c8565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906100a4906004016105f1565b60206040518083038186803b1580156100bc57600080fd5b505afa1580156100d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100f491906104b7565b6001600160a01b0316639eab52536040518163ffffffff1660e01b815260040160006040518083038186803b15801561012c57600080fd5b505afa158015610140573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261016891908101906104d3565b9050600061017747835161038d565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb90271906101b0906004016105bc565b60206040518083038186803b1580156101c857600080fd5b505afa1580156101dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020091906104b7565b90506000805b84518167ffffffffffffffff1610156102c157848167ffffffffffffffff168151811061022f57fe5b60200260200101519150826001600160a01b0316638ddde27c85846040518363ffffffff1660e01b815260040161026691906105a8565b6020604051808303818588803b15801561027f57600080fd5b505af1158015610293573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906102b89190610588565b50600101610206565b5050505050565b600054610100900460ff16806102e157506102e16103d3565b806102ef575060005460ff16155b61032a5760405162461bcd60e51b815260040180806020018281038252602e815260200180610638602e913960400191505060405180910390fd5b600054610100900460ff16158015610355576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610389576000805461ff00191690555b5050565b60006103cc83836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506103e4565b9392505050565b60006103de30610486565b15905090565b600081836104705760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561043557818101518382015260200161041d565b50505050905090810190601f1680156104625780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161047c57fe5b0495945050505050565b803b15155b919050565b805161048b8161061f565b6000602082840312156104ac578081fd5b81356103cc8161061f565b6000602082840312156104c8578081fd5b81516103cc8161061f565b600060208083850312156104e5578182fd5b825167ffffffffffffffff808211156104fc578384fd5b818501915085601f83011261050f578384fd5b81518181111561051b57fe5b8381026040518582820101818110858211171561053457fe5b604052828152858101935084860182860187018a1015610552578788fd5b8795505b8386101561057b5761056781610490565b855260019590950194938601938601610556565b5098975050505050505050565b600060208284031215610599578081fd5b815180151581146103cc578182fd5b6001600160a01b0391909116815260200190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b602080825260149082015273676f7665726e616e63652e726f6f744e6f64657360601b604082015260600190565b6001600160a01b038116811461063457600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a26469706673582212206da16cd23d006848edddc49a30d45581fb70a2700c4a8351956e810e3df25e6164736f6c63430007060033",
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
