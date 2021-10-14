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

// QHolderRewardPoolMetaData contains all meta data concerning the QHolderRewardPool contract.
var QHolderRewardPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferedToQVault\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"requestRewardTransfer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061044e806100206000396000f3fe60806040526004361061002d5760003560e01c8063c4d66de814610039578063e9f891f51461006e57610034565b3661003457005b600080fd5b34801561004557600080fd5b5061006c6004803603602081101561005c57600080fd5b50356001600160a01b03166100aa565b005b34801561007a57600080fd5b506100986004803603602081101561009157600080fd5b503561016f565b60408051918252519081900360200190f35b600054610100900460ff16806100c357506100c361038e565b806100d1575060005460ff16155b61010c5760405162461bcd60e51b815260040180806020018281038252602e8152602001806103eb602e913960400191505060405180910390fd5b600054610100900460ff16158015610137576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b03851602179055801561016b576000805461ff00191690555b5050565b6000805460408051633fb9027160e01b815260206004820181905260156024830152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b60448301529151620100009093046001600160a01b031692633fb9027192606480840193919291829003018186803b1580156101e557600080fd5b505afa1580156101f9573d6000803e3d6000fd5b505050506040513d602081101561020f57600080fd5b50516001600160a01b031633146102575760405162461bcd60e51b81526004018080602001828103825260458152602001806103a66045913960600191505060405180910390fd5b60005460408051633fb9027160e01b815260206004820181905260156024830152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b60448301529151620100009093046001600160a01b031692633fb9027192606480840193919291829003018186803b1580156102cc57600080fd5b505afa1580156102e0573d6000803e3d6000fd5b505050506040513d60208110156102f657600080fd5b505160408051630633136b60e31b815290516001600160a01b03909216916331989b58918591600480830192600092919082900301818588803b15801561033c57600080fd5b505af1158015610350573d6000803e3d6000fd5b50506040805186815290517f4c3dc86535ba84f8434ea47b1c4fb201cfdd30398118d9ba50ce2164d176783c94509081900360200192509050a15090565b60006103993061039f565b15905090565b3b15159056fe5b5145432d3031353030305d2d5065726d697373696f6e2064656e696564202d206f6e6c792074686520515661756c7420636f6e747261637420686173206163636573732e496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220a1c5a16ba623e5e76d32c033628bb209d108d37bc7a47d53a2d3ee69972f5a8964736f6c63430007060033",
}

// QHolderRewardPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use QHolderRewardPoolMetaData.ABI instead.
var QHolderRewardPoolABI = QHolderRewardPoolMetaData.ABI

// QHolderRewardPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QHolderRewardPoolMetaData.Bin instead.
var QHolderRewardPoolBin = QHolderRewardPoolMetaData.Bin

// DeployQHolderRewardPool deploys a new Ethereum contract, binding an instance of QHolderRewardPool to it.
func DeployQHolderRewardPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QHolderRewardPool, error) {
	parsed, err := QHolderRewardPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QHolderRewardPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QHolderRewardPool{QHolderRewardPoolCaller: QHolderRewardPoolCaller{contract: contract}, QHolderRewardPoolTransactor: QHolderRewardPoolTransactor{contract: contract}, QHolderRewardPoolFilterer: QHolderRewardPoolFilterer{contract: contract}}, nil
}

// QHolderRewardPool is an auto generated Go binding around an Ethereum contract.
type QHolderRewardPool struct {
	QHolderRewardPoolCaller     // Read-only binding to the contract
	QHolderRewardPoolTransactor // Write-only binding to the contract
	QHolderRewardPoolFilterer   // Log filterer for contract events
}

// QHolderRewardPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type QHolderRewardPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QHolderRewardPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QHolderRewardPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QHolderRewardPoolSession struct {
	Contract     *QHolderRewardPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// QHolderRewardPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QHolderRewardPoolCallerSession struct {
	Contract *QHolderRewardPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// QHolderRewardPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QHolderRewardPoolTransactorSession struct {
	Contract     *QHolderRewardPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// QHolderRewardPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type QHolderRewardPoolRaw struct {
	Contract *QHolderRewardPool // Generic contract binding to access the raw methods on
}

// QHolderRewardPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QHolderRewardPoolCallerRaw struct {
	Contract *QHolderRewardPoolCaller // Generic read-only contract binding to access the raw methods on
}

// QHolderRewardPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QHolderRewardPoolTransactorRaw struct {
	Contract *QHolderRewardPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQHolderRewardPool creates a new instance of QHolderRewardPool, bound to a specific deployed contract.
func NewQHolderRewardPool(address common.Address, backend bind.ContractBackend) (*QHolderRewardPool, error) {
	contract, err := bindQHolderRewardPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardPool{QHolderRewardPoolCaller: QHolderRewardPoolCaller{contract: contract}, QHolderRewardPoolTransactor: QHolderRewardPoolTransactor{contract: contract}, QHolderRewardPoolFilterer: QHolderRewardPoolFilterer{contract: contract}}, nil
}

// NewQHolderRewardPoolCaller creates a new read-only instance of QHolderRewardPool, bound to a specific deployed contract.
func NewQHolderRewardPoolCaller(address common.Address, caller bind.ContractCaller) (*QHolderRewardPoolCaller, error) {
	contract, err := bindQHolderRewardPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardPoolCaller{contract: contract}, nil
}

// NewQHolderRewardPoolTransactor creates a new write-only instance of QHolderRewardPool, bound to a specific deployed contract.
func NewQHolderRewardPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*QHolderRewardPoolTransactor, error) {
	contract, err := bindQHolderRewardPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardPoolTransactor{contract: contract}, nil
}

// NewQHolderRewardPoolFilterer creates a new log filterer instance of QHolderRewardPool, bound to a specific deployed contract.
func NewQHolderRewardPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*QHolderRewardPoolFilterer, error) {
	contract, err := bindQHolderRewardPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardPoolFilterer{contract: contract}, nil
}

// bindQHolderRewardPool binds a generic wrapper to an already deployed contract.
func bindQHolderRewardPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QHolderRewardPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QHolderRewardPool *QHolderRewardPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QHolderRewardPool.Contract.QHolderRewardPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QHolderRewardPool *QHolderRewardPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.QHolderRewardPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QHolderRewardPool *QHolderRewardPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.QHolderRewardPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QHolderRewardPool *QHolderRewardPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QHolderRewardPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QHolderRewardPool *QHolderRewardPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QHolderRewardPool *QHolderRewardPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.contract.Transact(opts, method, params...)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardPool *QHolderRewardPoolTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardPool.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardPool *QHolderRewardPoolSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.Initialize(&_QHolderRewardPool.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardPool *QHolderRewardPoolTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.Initialize(&_QHolderRewardPool.TransactOpts, _registry)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0xe9f891f5.
//
// Solidity: function requestRewardTransfer(uint256 _rewardAmount) returns(uint256)
func (_QHolderRewardPool *QHolderRewardPoolTransactor) RequestRewardTransfer(opts *bind.TransactOpts, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _QHolderRewardPool.contract.Transact(opts, "requestRewardTransfer", _rewardAmount)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0xe9f891f5.
//
// Solidity: function requestRewardTransfer(uint256 _rewardAmount) returns(uint256)
func (_QHolderRewardPool *QHolderRewardPoolSession) RequestRewardTransfer(_rewardAmount *big.Int) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.RequestRewardTransfer(&_QHolderRewardPool.TransactOpts, _rewardAmount)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0xe9f891f5.
//
// Solidity: function requestRewardTransfer(uint256 _rewardAmount) returns(uint256)
func (_QHolderRewardPool *QHolderRewardPoolTransactorSession) RequestRewardTransfer(_rewardAmount *big.Int) (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.RequestRewardTransfer(&_QHolderRewardPool.TransactOpts, _rewardAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardPool *QHolderRewardPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardPool *QHolderRewardPoolSession) Receive() (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.Receive(&_QHolderRewardPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardPool *QHolderRewardPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _QHolderRewardPool.Contract.Receive(&_QHolderRewardPool.TransactOpts)
}

// QHolderRewardPoolRewardTransferedToQVaultIterator is returned from FilterRewardTransferedToQVault and is used to iterate over the raw logs and unpacked data for RewardTransferedToQVault events raised by the QHolderRewardPool contract.
type QHolderRewardPoolRewardTransferedToQVaultIterator struct {
	Event *QHolderRewardPoolRewardTransferedToQVault // Event containing the contract specifics and raw log

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
func (it *QHolderRewardPoolRewardTransferedToQVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QHolderRewardPoolRewardTransferedToQVault)
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
		it.Event = new(QHolderRewardPoolRewardTransferedToQVault)
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
func (it *QHolderRewardPoolRewardTransferedToQVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QHolderRewardPoolRewardTransferedToQVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QHolderRewardPoolRewardTransferedToQVault represents a RewardTransferedToQVault event raised by the QHolderRewardPool contract.
type QHolderRewardPoolRewardTransferedToQVault struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTransferedToQVault is a free log retrieval operation binding the contract event 0x4c3dc86535ba84f8434ea47b1c4fb201cfdd30398118d9ba50ce2164d176783c.
//
// Solidity: event RewardTransferedToQVault(uint256 _amount)
func (_QHolderRewardPool *QHolderRewardPoolFilterer) FilterRewardTransferedToQVault(opts *bind.FilterOpts) (*QHolderRewardPoolRewardTransferedToQVaultIterator, error) {

	logs, sub, err := _QHolderRewardPool.contract.FilterLogs(opts, "RewardTransferedToQVault")
	if err != nil {
		return nil, err
	}
	return &QHolderRewardPoolRewardTransferedToQVaultIterator{contract: _QHolderRewardPool.contract, event: "RewardTransferedToQVault", logs: logs, sub: sub}, nil
}

// WatchRewardTransferedToQVault is a free log subscription operation binding the contract event 0x4c3dc86535ba84f8434ea47b1c4fb201cfdd30398118d9ba50ce2164d176783c.
//
// Solidity: event RewardTransferedToQVault(uint256 _amount)
func (_QHolderRewardPool *QHolderRewardPoolFilterer) WatchRewardTransferedToQVault(opts *bind.WatchOpts, sink chan<- *QHolderRewardPoolRewardTransferedToQVault) (event.Subscription, error) {

	logs, sub, err := _QHolderRewardPool.contract.WatchLogs(opts, "RewardTransferedToQVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QHolderRewardPoolRewardTransferedToQVault)
				if err := _QHolderRewardPool.contract.UnpackLog(event, "RewardTransferedToQVault", log); err != nil {
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

// ParseRewardTransferedToQVault is a log parse operation binding the contract event 0x4c3dc86535ba84f8434ea47b1c4fb201cfdd30398118d9ba50ce2164d176783c.
//
// Solidity: event RewardTransferedToQVault(uint256 _amount)
func (_QHolderRewardPool *QHolderRewardPoolFilterer) ParseRewardTransferedToQVault(log types.Log) (*QHolderRewardPoolRewardTransferedToQVault, error) {
	event := new(QHolderRewardPoolRewardTransferedToQVault)
	if err := _QHolderRewardPool.contract.UnpackLog(event, "RewardTransferedToQVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
