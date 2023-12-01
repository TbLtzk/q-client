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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Allocated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d70806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057600080fd5b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610a44565b61089c565b6100a96040518060c001604052806000815260200160008152602001600081526020016000815260200160008152602001600081525090565b478082526100b45750565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610d19602291396040518263ffffffff1660e01b815260040161010e9190610a61565b60206040518083038186803b15801561012657600080fd5b505afa15801561013a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061015e9190610ab6565b60405162498bff60e81b815260206004820152602260248201527f636f6e737469747574696f6e2e6d61784e5374616e64627956616c696461746f604482015261727360f01b60648201529091506001600160a01b0382169063498bff009060840160206040518083038186803b1580156101d857600080fd5b505afa1580156101ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102109190610ad3565b60405162498bff60e81b815260206004820152601b60248201527a636f6e737469747574696f6e2e6d61784e56616c696461746f727360281b60448201526001600160a01b0383169063498bff009060640160206040518083038186803b15801561027a57600080fd5b505afa15801561028e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b29190610ad3565b6102bc9190610b02565b602083810191909152600080546040805180820182526015815274676f7665726e616e63652e76616c696461746f727360581b9481019490945251633fb9027160e01b81529192620100009091046001600160a01b031691633fb902719161032691600401610a61565b60206040518083038186803b15801561033e57600080fd5b505afa158015610352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103769190610ab6565b90506000816001600160a01b0316630f5b8db36040518163ffffffff1660e01b815260040160006040518083038186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103ef9190810190610b8a565b905083602001518151101561040657805160208501525b60005b84602001518110156104575781818151811061042757610427610c60565b60200260200101516020015185604001516104429190610b02565b604086015261045081610c76565b9050610409565b5060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060248152602001610cf5602491396040518263ffffffff1660e01b81526004016104b29190610a61565b60206040518083038186803b1580156104ca57600080fd5b505afa1580156104de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105029190610ab6565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb902719161056c9190600401610a61565b60206040518083038186803b15801561058457600080fd5b505afa158015610598573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105bc9190610ab6565b90506b033b2e3c9fd0803ce80000006000805b886020015181101561085e57846001600160a01b031663841c90648783815181106105fc576105fc610c60565b6020026020010151600001516040518263ffffffff1660e01b81526004016106249190610c91565b60206040518083038186803b15801561063c57600080fd5b505afa158015610650573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106749190610ad3565b8960600181815250506106c78960600151846106c089858151811061069b5761069b610c60565b6020026020010151602001518d604001518e600001516109769092919063ffffffff16565b9190610976565b60808a018190521561076657846001600160a01b031663eab136a08a608001518884815181106106f9576106f9610c60565b6020026020010151600001516040518363ffffffff1660e01b81526004016107219190610c91565b6000604051808303818588803b15801561073a57600080fd5b505af115801561074e573d6000803e3d6000fd5b50505050508860800151826107639190610b02565b91505b61078f8960600151846107799190610ca5565b846106c089858151811061069b5761069b610c60565b60a08a018190521561084e57836001600160a01b0316638ddde27c8a60a001518884815181106107c1576107c1610c60565b6020026020010151600001516040518363ffffffff1660e01b81526004016107e99190610c91565b6020604051808303818588803b15801561080257600080fd5b505af1158015610816573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061083b9190610cbc565b5060a089015161084b9083610b02565b91505b61085781610c76565b90506105cf565b506040518181527f60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be199060200160405180910390a15050505050505050565b600054610100900460ff16806108b5575060005460ff16155b61091c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff1615801561093e576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610972576000805461ff00191690555b5050565b6000808060001985870985870292508281108382030391505080600014156109b1578382816109a7576109a7610cde565b0492505050610a25565b8084116109bd57600080fd5b600084868809851960019081018716968790049682860381900495909211909303600082900391909104909201919091029190911760038402600290811880860282030280860282030280860282030280860282030280860282030280860290910302029150505b9392505050565b6001600160a01b0381168114610a4157600080fd5b50565b600060208284031215610a5657600080fd5b8135610a2581610a2c565b600060208083528351808285015260005b81811015610a8e57858101830151858201604001528201610a72565b81811115610aa0576000604083870101525b50601f01601f1916929092016040019392505050565b600060208284031215610ac857600080fd5b8151610a2581610a2c565b600060208284031215610ae557600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610b1557610b15610aec565b500190565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610b5357610b53610b1a565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b8257610b82610b1a565b604052919050565b60006020808385031215610b9d57600080fd5b825167ffffffffffffffff80821115610bb557600080fd5b818501915085601f830112610bc957600080fd5b815181811115610bdb57610bdb610b1a565b610be9848260051b01610b59565b818152848101925060069190911b830184019087821115610c0957600080fd5b928401925b81841015610c555760408489031215610c275760008081fd5b610c2f610b30565b8451610c3a81610a2c565b81528486015186820152835260409093019291840191610c0e565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600019821415610c8a57610c8a610aec565b5060010190565b6001600160a01b0391909116815260200190565b600082821015610cb757610cb7610aec565b500390565b600060208284031215610cce57600080fd5b81518015158114610a2557600080fd5b634e487b7160e01b600052601260045260246000fdfe746f6b656e65636f6e6f6d6963732e76616c69646174696f6e526577617264506f6f6c73676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a26469706673582212200be7f87312ac090faf9f2f653e31ee310756a72eba58ce1be85668fe8fba929564736f6c63430008090033",
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

// ValidationRewardProxyAllocatedIterator is returned from FilterAllocated and is used to iterate over the raw logs and unpacked data for Allocated events raised by the ValidationRewardProxy contract.
type ValidationRewardProxyAllocatedIterator struct {
	Event *ValidationRewardProxyAllocated // Event containing the contract specifics and raw log

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
func (it *ValidationRewardProxyAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRewardProxyAllocated)
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
		it.Event = new(ValidationRewardProxyAllocated)
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
func (it *ValidationRewardProxyAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRewardProxyAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRewardProxyAllocated represents a Allocated event raised by the ValidationRewardProxy contract.
type ValidationRewardProxyAllocated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllocated is a free log retrieval operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_ValidationRewardProxy *ValidationRewardProxyFilterer) FilterAllocated(opts *bind.FilterOpts) (*ValidationRewardProxyAllocatedIterator, error) {

	logs, sub, err := _ValidationRewardProxy.contract.FilterLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyAllocatedIterator{contract: _ValidationRewardProxy.contract, event: "Allocated", logs: logs, sub: sub}, nil
}

// WatchAllocated is a free log subscription operation binding the contract event 0x60a0c13c5aa0acfe5cd0c4f3ac6b9fe2742094a559183932b66e20108fc2be19.
//
// Solidity: event Allocated(uint256 amount)
func (_ValidationRewardProxy *ValidationRewardProxyFilterer) WatchAllocated(opts *bind.WatchOpts, sink chan<- *ValidationRewardProxyAllocated) (event.Subscription, error) {

	logs, sub, err := _ValidationRewardProxy.contract.WatchLogs(opts, "Allocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRewardProxyAllocated)
				if err := _ValidationRewardProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
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
func (_ValidationRewardProxy *ValidationRewardProxyFilterer) ParseAllocated(log types.Log) (*ValidationRewardProxyAllocated, error) {
	event := new(ValidationRewardProxyAllocated)
	if err := _ValidationRewardProxy.contract.UnpackLog(event, "Allocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
