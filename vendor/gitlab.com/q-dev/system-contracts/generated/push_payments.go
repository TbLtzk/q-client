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

// PushPaymentsMetaData contains all meta data concerning the PushPayments contract.
var PushPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"safeTransferTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102ff806100206000396000f3fe6080604052600436106100345760003560e01c80633ccfd60b1461003957806370a08231146100505780638ddde27c14610083575b600080fd5b34801561004557600080fd5b5061004e6100a6565b005b34801561005c57600080fd5b5061007061006b366004610273565b610203565b6040519081526020015b60405180910390f35b610096610091366004610273565b61021e565b604051901515815260200161007a565b3360006100b282610203565b90506000811161012f5760405162461bcd60e51b815260206004820152603e60248201527f5b5145432d3033333030315d2d5468652063616c6c657220646f6573206e6f7460448201527f206861766520616e792062616c616e636520746f2077697468647261772e000060648201526084015b60405180910390fd5b6001600160a01b0382166000818152600160205260408082208290555190919083908381818185875af1925050503d8060008114610189576040519150601f19603f3d011682016040523d82523d6000602084013e61018e565b606091505b50509050806101fe5760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3033333030325d2d5472616e7366657220746f207468652077697460448201527534323930bbb0b61031b0b63632b9103330b4b632b21760511b6064820152608401610126565b505050565b6001600160a01b031660009081526001602052604090205490565b6000806000806000803487617530f190508061023d5761023d83610243565b92915050565b6001600160a01b0381166000908152600160205260408120805434929061026b9084906102a3565b909155505050565b60006020828403121561028557600080fd5b81356001600160a01b038116811461029c57600080fd5b9392505050565b600082198211156102c457634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220a931940c69eca94ebfa35dec13c492fadafb933164a973dce2783e455f48d8ae64736f6c63430008090033",
}

// PushPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use PushPaymentsMetaData.ABI instead.
var PushPaymentsABI = PushPaymentsMetaData.ABI

// PushPaymentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PushPaymentsMetaData.Bin instead.
var PushPaymentsBin = PushPaymentsMetaData.Bin

// DeployPushPayments deploys a new Ethereum contract, binding an instance of PushPayments to it.
func DeployPushPayments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PushPayments, error) {
	parsed, err := PushPaymentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PushPaymentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PushPayments{PushPaymentsCaller: PushPaymentsCaller{contract: contract}, PushPaymentsTransactor: PushPaymentsTransactor{contract: contract}, PushPaymentsFilterer: PushPaymentsFilterer{contract: contract}}, nil
}

// PushPayments is an auto generated Go binding around an Ethereum contract.
type PushPayments struct {
	PushPaymentsCaller     // Read-only binding to the contract
	PushPaymentsTransactor // Write-only binding to the contract
	PushPaymentsFilterer   // Log filterer for contract events
}

// PushPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PushPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PushPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PushPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PushPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PushPaymentsSession struct {
	Contract     *PushPayments     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PushPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PushPaymentsCallerSession struct {
	Contract *PushPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PushPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PushPaymentsTransactorSession struct {
	Contract     *PushPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PushPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PushPaymentsRaw struct {
	Contract *PushPayments // Generic contract binding to access the raw methods on
}

// PushPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PushPaymentsCallerRaw struct {
	Contract *PushPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// PushPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PushPaymentsTransactorRaw struct {
	Contract *PushPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPushPayments creates a new instance of PushPayments, bound to a specific deployed contract.
func NewPushPayments(address common.Address, backend bind.ContractBackend) (*PushPayments, error) {
	contract, err := bindPushPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PushPayments{PushPaymentsCaller: PushPaymentsCaller{contract: contract}, PushPaymentsTransactor: PushPaymentsTransactor{contract: contract}, PushPaymentsFilterer: PushPaymentsFilterer{contract: contract}}, nil
}

// NewPushPaymentsCaller creates a new read-only instance of PushPayments, bound to a specific deployed contract.
func NewPushPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PushPaymentsCaller, error) {
	contract, err := bindPushPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PushPaymentsCaller{contract: contract}, nil
}

// NewPushPaymentsTransactor creates a new write-only instance of PushPayments, bound to a specific deployed contract.
func NewPushPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PushPaymentsTransactor, error) {
	contract, err := bindPushPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PushPaymentsTransactor{contract: contract}, nil
}

// NewPushPaymentsFilterer creates a new log filterer instance of PushPayments, bound to a specific deployed contract.
func NewPushPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PushPaymentsFilterer, error) {
	contract, err := bindPushPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PushPaymentsFilterer{contract: contract}, nil
}

// bindPushPayments binds a generic wrapper to an already deployed contract.
func bindPushPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PushPaymentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PushPayments *PushPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PushPayments.Contract.PushPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PushPayments *PushPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PushPayments.Contract.PushPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PushPayments *PushPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PushPayments.Contract.PushPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PushPayments *PushPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PushPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PushPayments *PushPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PushPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PushPayments *PushPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PushPayments.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PushPayments *PushPaymentsCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PushPayments.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PushPayments *PushPaymentsSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _PushPayments.Contract.BalanceOf(&_PushPayments.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_PushPayments *PushPaymentsCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _PushPayments.Contract.BalanceOf(&_PushPayments.CallOpts, _account)
}

// SafeTransferTo is a paid mutator transaction binding the contract method 0x8ddde27c.
//
// Solidity: function safeTransferTo(address _account) payable returns(bool)
func (_PushPayments *PushPaymentsTransactor) SafeTransferTo(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PushPayments.contract.Transact(opts, "safeTransferTo", _account)
}

// SafeTransferTo is a paid mutator transaction binding the contract method 0x8ddde27c.
//
// Solidity: function safeTransferTo(address _account) payable returns(bool)
func (_PushPayments *PushPaymentsSession) SafeTransferTo(_account common.Address) (*types.Transaction, error) {
	return _PushPayments.Contract.SafeTransferTo(&_PushPayments.TransactOpts, _account)
}

// SafeTransferTo is a paid mutator transaction binding the contract method 0x8ddde27c.
//
// Solidity: function safeTransferTo(address _account) payable returns(bool)
func (_PushPayments *PushPaymentsTransactorSession) SafeTransferTo(_account common.Address) (*types.Transaction, error) {
	return _PushPayments.Contract.SafeTransferTo(&_PushPayments.TransactOpts, _account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PushPayments *PushPaymentsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PushPayments.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PushPayments *PushPaymentsSession) Withdraw() (*types.Transaction, error) {
	return _PushPayments.Contract.Withdraw(&_PushPayments.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PushPayments *PushPaymentsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PushPayments.Contract.Withdraw(&_PushPayments.TransactOpts)
}
