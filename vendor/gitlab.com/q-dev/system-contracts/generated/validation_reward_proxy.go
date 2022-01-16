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

// ValidationRewardProxyMetaData contains all meta data concerning the ValidationRewardProxy contract.
var ValidationRewardProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d8f806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057600080fd5b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610a5a565b610844565b6100a96040518060c001604052806000815260200160008152602001600081526020016000815260200160008152602001600081525090565b478082526100b45750565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610d38602291396040518263ffffffff1660e01b815260040161010e9190610a77565b60206040518083038186803b15801561012657600080fd5b505afa15801561013a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061015e9190610acc565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e6d61784e5374616e64627956616c696461746f604482015261727360f01b60648201529091506001600160a01b0382169063498bff009060840160206040518083038186803b1580156101d857600080fd5b505afa1580156101ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102109190610ae9565b60405162498bff60e81b815260206004820152601b60248201527a636f6e737469747574696f6e2e6d61784e56616c696461746f727360281b60448201526001600160a01b0383169063498bff009060640160206040518083038186803b15801561027a57600080fd5b505afa15801561028e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b29190610ae9565b6102bc9190610b18565b602083810191909152600080546040805180820182526015815274676f7665726e616e63652e76616c696461746f727360581b9481019490945251633fb9027160e01b81529192620100009091046001600160a01b031691633fb902719161032691600401610a77565b60206040518083038186803b15801561033e57600080fd5b505afa158015610352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103769190610acc565b90506000816001600160a01b0316630f5b8db36040518163ffffffff1660e01b815260040160006040518083038186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103ef9190810190610ba0565b905083602001518151101561040657805160208501525b60005b84602001518110156104575781818151811061042757610427610c76565b60200260200101516020015185604001516104429190610b18565b604086015261045081610c8c565b9050610409565b5060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060248152602001610d14602491396040518263ffffffff1660e01b81526004016104b29190610a77565b60206040518083038186803b1580156104ca57600080fd5b505afa1580156104de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105029190610acc565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb902719161056c9190600401610a77565b60206040518083038186803b15801561058457600080fd5b505afa158015610598573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105bc9190610acc565b90506b033b2e3c9fd0803ce800000060005b876020015181101561083a57836001600160a01b031663841c90648683815181106105fb576105fb610c76565b6020026020010151600001516040518263ffffffff1660e01b81526004016106239190610ca7565b60206040518083038186803b15801561063b57600080fd5b505afa15801561064f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106739190610ae9565b8860600181815250506106c68860600151836106bf88858151811061069a5761069a610c76565b6020026020010151602001518c604001518d6000015161091f9092919063ffffffff16565b919061091f565b608089018190521561075357836001600160a01b031663eab136a089608001518784815181106106f8576106f8610c76565b6020026020010151600001516040518363ffffffff1660e01b81526004016107209190610ca7565b6000604051808303818588803b15801561073957600080fd5b505af115801561074d573d6000803e3d6000fd5b50505050505b61077c8860600151836107669190610cbb565b836106bf88858151811061069a5761069a610c76565b60a089018190521561082a57826001600160a01b0316638ddde27c8960a001518784815181106107ae576107ae610c76565b6020026020010151600001516040518363ffffffff1660e01b81526004016107d69190610ca7565b6020604051808303818588803b1580156107ef57600080fd5b505af1158015610803573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906108289190610cd2565b505b61083381610c8c565b90506105ce565b5050505050505050565b600054610100900460ff168061085d575060005460ff16155b6108c55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff161580156108e7576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b03851602179055801561091b576000805461ff00191690555b5050565b6000808060001985870985870292508281108382030391505080600014156109af57600084116109a45760405162461bcd60e51b815260206004820152602a60248201527f46756c6c4d6174683a2064656e6f6d696e61746f722073686f756c6420626520604482015269036b7b932903d32b937960b51b60648201526084016108bc565b508290049050610a3b565b8084116109bb57600080fd5b6000848688098084039381119092039190506000856109dc81196001610b18565b16958690049560008181038290046001018581029290960491909101949150610a06876003610cf4565b600290811888810282030280890282030280890282030280890282030280890282030280890290910302949094029450505050505b9392505050565b6001600160a01b0381168114610a5757600080fd5b50565b600060208284031215610a6c57600080fd5b8135610a3b81610a42565b600060208083528351808285015260005b81811015610aa457858101830151858201604001528201610a88565b81811115610ab6576000604083870101525b50601f01601f1916929092016040019392505050565b600060208284031215610ade57600080fd5b8151610a3b81610a42565b600060208284031215610afb57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610b2b57610b2b610b02565b500190565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610b6957610b69610b30565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b9857610b98610b30565b604052919050565b60006020808385031215610bb357600080fd5b825167ffffffffffffffff80821115610bcb57600080fd5b818501915085601f830112610bdf57600080fd5b815181811115610bf157610bf1610b30565b610bff848260051b01610b6f565b818152848101925060069190911b830184019087821115610c1f57600080fd5b928401925b81841015610c6b5760408489031215610c3d5760008081fd5b610c45610b46565b8451610c5081610a42565b81528486015186820152835260409093019291840191610c24565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610ca057610ca0610b02565b5060010190565b6001600160a01b0391909116815260200190565b600082821015610ccd57610ccd610b02565b500390565b600060208284031215610ce457600080fd5b81518015158114610a3b57600080fd5b6000816000190483118215151615610d0e57610d0e610b02565b50029056fe746f6b656e65636f6e6f6d6963732e76616c69646174696f6e526577617264506f6f6c73676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220faa43c6afb601c9275a899418ab6734814ba01cdeb0de29e21dd5920fb73b05164736f6c63430008090033",
}

// ValidationRewardProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidationRewardProxyMetaData.ABI instead.
var ValidationRewardProxyABI = ValidationRewardProxyMetaData.ABI

// ValidationRewardProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidationRewardProxyMetaData.Bin instead.
var ValidationRewardProxyBin = ValidationRewardProxyMetaData.Bin

// DeployValidationRewardProxy deploys a new Ethereum contract, binding an instance of ValidationRewardProxy to it.
func DeployValidationRewardProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidationRewardProxy, error) {
	parsed, err := ValidationRewardProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidationRewardProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidationRewardProxy{ValidationRewardProxyCaller: ValidationRewardProxyCaller{contract: contract}, ValidationRewardProxyTransactor: ValidationRewardProxyTransactor{contract: contract}, ValidationRewardProxyFilterer: ValidationRewardProxyFilterer{contract: contract}}, nil
}

// ValidationRewardProxy is an auto generated Go binding around an Ethereum contract.
type ValidationRewardProxy struct {
	ValidationRewardProxyCaller     // Read-only binding to the contract
	ValidationRewardProxyTransactor // Write-only binding to the contract
	ValidationRewardProxyFilterer   // Log filterer for contract events
}

// ValidationRewardProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidationRewardProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidationRewardProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidationRewardProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidationRewardProxySession struct {
	Contract     *ValidationRewardProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ValidationRewardProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidationRewardProxyCallerSession struct {
	Contract *ValidationRewardProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ValidationRewardProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidationRewardProxyTransactorSession struct {
	Contract     *ValidationRewardProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ValidationRewardProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidationRewardProxyRaw struct {
	Contract *ValidationRewardProxy // Generic contract binding to access the raw methods on
}

// ValidationRewardProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidationRewardProxyCallerRaw struct {
	Contract *ValidationRewardProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ValidationRewardProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidationRewardProxyTransactorRaw struct {
	Contract *ValidationRewardProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidationRewardProxy creates a new instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxy(address common.Address, backend bind.ContractBackend) (*ValidationRewardProxy, error) {
	contract, err := bindValidationRewardProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxy{ValidationRewardProxyCaller: ValidationRewardProxyCaller{contract: contract}, ValidationRewardProxyTransactor: ValidationRewardProxyTransactor{contract: contract}, ValidationRewardProxyFilterer: ValidationRewardProxyFilterer{contract: contract}}, nil
}

// NewValidationRewardProxyCaller creates a new read-only instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyCaller(address common.Address, caller bind.ContractCaller) (*ValidationRewardProxyCaller, error) {
	contract, err := bindValidationRewardProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyCaller{contract: contract}, nil
}

// NewValidationRewardProxyTransactor creates a new write-only instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidationRewardProxyTransactor, error) {
	contract, err := bindValidationRewardProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyTransactor{contract: contract}, nil
}

// NewValidationRewardProxyFilterer creates a new log filterer instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidationRewardProxyFilterer, error) {
	contract, err := bindValidationRewardProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyFilterer{contract: contract}, nil
}

// bindValidationRewardProxy binds a generic wrapper to an already deployed contract.
func bindValidationRewardProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRewardProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardProxy *ValidationRewardProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardProxy *ValidationRewardProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardProxy *ValidationRewardProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.contract.Transact(opts, method, params...)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Allocate() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Allocate(&_ValidationRewardProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Allocate(&_ValidationRewardProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Initialize(&_ValidationRewardProxy.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Initialize(&_ValidationRewardProxy.TransactOpts, _registry)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Receive() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Receive(&_ValidationRewardProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Receive(&_ValidationRewardProxy.TransactOpts)
}
