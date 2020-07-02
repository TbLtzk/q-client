// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "gitlab.com/q-dev/go-ethereum"
	"gitlab.com/q-dev/go-ethereum/accounts/abi"
	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/core/types"
	"gitlab.com/q-dev/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RewardReceiverABI is the input ABI used to generate the binding from.
const RewardReceiverABI = "[{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RewardReceiverFuncSigs maps the 4-byte function signature to its string representation.
var RewardReceiverFuncSigs = map[string]string{
	"9e1a00aa": "sendTo(address,uint256)",
}

// RewardReceiverBin is the compiled bytecode used for deploying new contracts.
var RewardReceiverBin = "0x608060405234801561001057600080fd5b5060b68061001f6000396000f3fe608060405260043610601c5760003560e01c80639e1a00aa14601e575b005b348015602957600080fd5b50601c60048036036040811015603e57600080fd5b5060405181356001600160a01b031691602001359082906108fc8315029083906000818181858888f19350505050158015607c573d6000803e3d6000fd5b50505056fea265627a7a7231582074e7195a0497f2fee65a07cd920928245e5f062fa1c199993c97724d14fca57964736f6c63430005110032"

// DeployRewardReceiver deploys a new Ethereum contract, binding an instance of RewardReceiver to it.
func DeployRewardReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RewardReceiver, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RewardReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardReceiver{RewardReceiverCaller: RewardReceiverCaller{contract: contract}, RewardReceiverTransactor: RewardReceiverTransactor{contract: contract}, RewardReceiverFilterer: RewardReceiverFilterer{contract: contract}}, nil
}

// RewardReceiver is an auto generated Go binding around an Ethereum contract.
type RewardReceiver struct {
	RewardReceiverCaller     // Read-only binding to the contract
	RewardReceiverTransactor // Write-only binding to the contract
	RewardReceiverFilterer   // Log filterer for contract events
}

// RewardReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardReceiverSession struct {
	Contract     *RewardReceiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardReceiverCallerSession struct {
	Contract *RewardReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RewardReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardReceiverTransactorSession struct {
	Contract     *RewardReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RewardReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardReceiverRaw struct {
	Contract *RewardReceiver // Generic contract binding to access the raw methods on
}

// RewardReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardReceiverCallerRaw struct {
	Contract *RewardReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// RewardReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardReceiverTransactorRaw struct {
	Contract *RewardReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardReceiver creates a new instance of RewardReceiver, bound to a specific deployed contract.
func NewRewardReceiver(address common.Address, backend bind.ContractBackend) (*RewardReceiver, error) {
	contract, err := bindRewardReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardReceiver{RewardReceiverCaller: RewardReceiverCaller{contract: contract}, RewardReceiverTransactor: RewardReceiverTransactor{contract: contract}, RewardReceiverFilterer: RewardReceiverFilterer{contract: contract}}, nil
}

// NewRewardReceiverCaller creates a new read-only instance of RewardReceiver, bound to a specific deployed contract.
func NewRewardReceiverCaller(address common.Address, caller bind.ContractCaller) (*RewardReceiverCaller, error) {
	contract, err := bindRewardReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardReceiverCaller{contract: contract}, nil
}

// NewRewardReceiverTransactor creates a new write-only instance of RewardReceiver, bound to a specific deployed contract.
func NewRewardReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardReceiverTransactor, error) {
	contract, err := bindRewardReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardReceiverTransactor{contract: contract}, nil
}

// NewRewardReceiverFilterer creates a new log filterer instance of RewardReceiver, bound to a specific deployed contract.
func NewRewardReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardReceiverFilterer, error) {
	contract, err := bindRewardReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardReceiverFilterer{contract: contract}, nil
}

// bindRewardReceiver binds a generic wrapper to an already deployed contract.
func bindRewardReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardReceiver *RewardReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardReceiver.Contract.RewardReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardReceiver *RewardReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardReceiver.Contract.RewardReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardReceiver *RewardReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardReceiver.Contract.RewardReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardReceiver *RewardReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RewardReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardReceiver *RewardReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardReceiver *RewardReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardReceiver.Contract.contract.Transact(opts, method, params...)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(address receiver, uint256 amount) returns()
func (_RewardReceiver *RewardReceiverTransactor) SendTo(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardReceiver.contract.Transact(opts, "sendTo", receiver, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(address receiver, uint256 amount) returns()
func (_RewardReceiver *RewardReceiverSession) SendTo(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardReceiver.Contract.SendTo(&_RewardReceiver.TransactOpts, receiver, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0x9e1a00aa.
//
// Solidity: function sendTo(address receiver, uint256 amount) returns()
func (_RewardReceiver *RewardReceiverTransactorSession) SendTo(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RewardReceiver.Contract.SendTo(&_RewardReceiver.TransactOpts, receiver, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RewardReceiver *RewardReceiverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RewardReceiver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RewardReceiver *RewardReceiverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RewardReceiver.Contract.Fallback(&_RewardReceiver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RewardReceiver *RewardReceiverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RewardReceiver.Contract.Fallback(&_RewardReceiver.TransactOpts, calldata)
}
