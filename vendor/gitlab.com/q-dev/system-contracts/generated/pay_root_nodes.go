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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Allocated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106cc806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057600080fd5b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610467565b610375565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb90271916100cd9160040161048b565b60206040518083038186803b1580156100e557600080fd5b505afa1580156100f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011d91906104f0565b6001600160a01b0316639eab52536040518163ffffffff1660e01b815260040160006040518083038186803b15801561015557600080fd5b505afa158015610169573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101919190810190610523565b905060008151476101a291906105fd565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb902719161020c919060040161048b565b60206040518083038186803b15801561022457600080fd5b505afa158015610238573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025c91906104f0565b905060008060005b8551816001600160401b0316101561033a5785816001600160401b0316815181106102915761029161061f565b6020908102919091010151604051632377789f60e21b81526001600160a01b03808316600483015291945090851690638ddde27c9087906024016020604051808303818588803b1580156102e457600080fd5b505af11580156102f8573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061031d9190610635565b506103288583610657565b91506103338161066f565b9050610264565b506040518181527f60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be199060200160405180910390a15050505050565b600054610100900460ff168061038e575060005460ff16155b6103f55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015610417576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b03851602179055801561044b576000805461ff00191690555b5050565b6001600160a01b038116811461046457600080fd5b50565b60006020828403121561047957600080fd5b81356104848161044f565b9392505050565b600060208083528351808285015260005b818110156104b85785810183015185820160400152820161049c565b818111156104ca576000604083870101525b50601f01601f1916929092016040019392505050565b80516104eb8161044f565b919050565b60006020828403121561050257600080fd5b81516104848161044f565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561053657600080fd5b82516001600160401b038082111561054d57600080fd5b818501915085601f83011261056157600080fd5b8151818111156105735761057361050d565b8060051b604051601f19603f830116810181811085821117156105985761059861050d565b6040529182528482019250838101850191888311156105b657600080fd5b938501935b828510156105db576105cc856104e0565b845293850193928501926105bb565b98975050505050505050565b634e487b7160e01b600052601160045260246000fd5b60008261061a57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561064757600080fd5b8151801515811461048457600080fd5b6000821982111561066a5761066a6105e7565b500190565b60006001600160401b038083168181141561068c5761068c6105e7565b600101939250505056fea2646970667358221220d6a3ebb1655056f6032e0465819b36128901f8b5c2cfc04fba6e081778e221cb64736f6c63430008090033",
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

// RootNodeRewardProxyAllocatedIterator is returned from FilterAllocated and is used to iterate over the raw logs and unpacked data for Allocated events raised by the RootNodeRewardProxy contract.
type RootNodeRewardProxyAllocatedIterator struct {
	Event *RootNodeRewardProxyAllocated // Event containing the contract specifics and raw log

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
func (it *RootNodeRewardProxyAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootNodeRewardProxyAllocated)
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
		it.Event = new(RootNodeRewardProxyAllocated)
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
func (it *RootNodeRewardProxyAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootNodeRewardProxyAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootNodeRewardProxyAllocated represents a Allocated event raised by the RootNodeRewardProxy contract.
type RootNodeRewardProxyAllocated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocated is a free log retrieval operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_RootNodeRewardProxy *RootNodeRewardProxyFilterer) FilterAllocated(opts *bind.FilterOpts) (*RootNodeRewardProxyAllocatedIterator, error) {

	logs, sub, err := _RootNodeRewardProxy.contract.FilterLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return &RootNodeRewardProxyAllocatedIterator{contract: _RootNodeRewardProxy.contract, event: "Allocated", logs: logs, sub: sub}, nil
}

// WatchAllocated is a free log subscription operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_RootNodeRewardProxy *RootNodeRewardProxyFilterer) WatchAllocated(opts *bind.WatchOpts, sink chan<- *RootNodeRewardProxyAllocated) (event.Subscription, error) {

	logs, sub, err := _RootNodeRewardProxy.contract.WatchLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootNodeRewardProxyAllocated)
				if err := _RootNodeRewardProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
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

// ParseAllocated is a log parse operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_RootNodeRewardProxy *RootNodeRewardProxyFilterer) ParseAllocated(log types.Log) (*RootNodeRewardProxyAllocated, error) {
	event := new(RootNodeRewardProxyAllocated)
	if err := _RootNodeRewardProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
