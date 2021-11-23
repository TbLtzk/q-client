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

// ExceptionStorageMetaData contains all meta data concerning the ExceptionStorage contract.
var ExceptionStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"exceptions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"key\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"name\":\"setException\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"keys\",\"type\":\"uint32[]\"},{\"internalType\":\"string[]\",\"name\":\"notes\",\"type\":\"string[]\"}],\"name\":\"setExceptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"key\",\"type\":\"uint32\"}],\"name\":\"getException\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"note\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"key\",\"type\":\"uint32\"}],\"name\":\"throwException\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600061001b61006a565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006e565b3390565b6109d78061007d6000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806308f52c561461007d578063524ec3ec146100a6578063715018a6146100bb5780638da5cb5b146100c3578063d1c21d65146100d8578063db13183e146100eb578063e150d645146100fe578063f2fde38b14610111575b600080fd5b61009061008b366004610822565b610124565b60405161009d9190610891565b60405180910390f35b6100b96100b4366004610822565b610205565b005b6100b9610230565b6100cb6102d2565b60405161009d919061087d565b6100b96100e636600461083c565b6102e1565b6100906100f9366004610822565b610364565b6100b961010c366004610765565b6103fe565b6100b961011f366004610737565b6104bb565b63ffffffff811660009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156101bf5780601f10610194576101008083540402835291602001916101bf565b820191906000526020600020905b8154815290600101906020018083116101a257829003601f168201915b5050505050905080516000141561020057506040805180820190915260178152764572726f72207769746820756e6b6e6f776e20636f646560481b60208201525b919050565b61020e81610124565b60405162461bcd60e51b81526004016102279190610891565b60405180910390fd5b6102386105b3565b6000546001600160a01b03908116911614610288576040805162461bcd60e51b81526020600482018190526024820152600080516020610982833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b6102e96105b3565b6000546001600160a01b03908116911614610339576040805162461bcd60e51b81526020600482018190526024820152600080516020610982833981519152604482015290519081900360640190fd5b63ffffffff82166000908152600160209081526040909120825161035f928401906105b7565b505050565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156103f65780601f106103cb576101008083540402835291602001916103f6565b820191906000526020600020905b8154815290600101906020018083116103d957829003601f168201915b505050505081565b6104066105b3565b6000546001600160a01b03908116911614610456576040805162461bcd60e51b81526020600482018190526024820152600080516020610982833981519152604482015290519081900360640190fd5b80518251146104775760405162461bcd60e51b8152600401610227906108e4565b60005b825181101561035f576104b383828151811061049257fe5b60200260200101518383815181106104a657fe5b60200260200101516102e1565b60010161047a565b6104c36105b3565b6000546001600160a01b03908116911614610513576040805162461bcd60e51b81526020600482018190526024820152600080516020610982833981519152604482015290519081900360640190fd5b6001600160a01b0381166105585760405162461bcd60e51b815260040180806020018281038252602681526020018061095c6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826105ed5760008555610633565b82601f1061060657805160ff1916838001178555610633565b82800160010185558215610633579182015b82811115610633578251825591602001919060010190610618565b5061063f929150610643565b5090565b5b8082111561063f5760008155600101610644565b600082601f830112610668578081fd5b8135602061067d6106788361093e565b61091b565b82815281810190858301855b858110156106b2576106a0898684358b01016106bf565b84529284019290840190600101610689565b5090979650505050505050565b600082601f8301126106cf578081fd5b81356001600160401b038111156106e257fe5b6106f5601f8201601f191660200161091b565b818152846020838601011115610709578283fd5b816020850160208301379081016020019190915292915050565b803563ffffffff8116811461020057600080fd5b600060208284031215610748578081fd5b81356001600160a01b038116811461075e578182fd5b9392505050565b60008060408385031215610777578081fd5b82356001600160401b038082111561078d578283fd5b818501915085601f8301126107a0578283fd5b813560206107b06106788361093e565b82815281810190858301838502870184018b10156107cc578788fd5b8796505b848710156107f5576107e181610723565b8352600196909601959183019183016107d0565b509650508601359250508082111561080b578283fd5b5061081885828601610658565b9150509250929050565b600060208284031215610833578081fd5b61075e82610723565b6000806040838503121561084e578182fd5b61085783610723565b915060208301356001600160401b03811115610871578182fd5b610818858286016106bf565b6001600160a01b0391909116815260200190565b6000602080835283518082850152825b818110156108bd578581018301518582016040015282016108a1565b818111156108ce5783604083870101525b50601f01601f1916929092016040019392505050565b6020808252601c908201527f6172726179206c656e6774686573206d75737420626520657175616c00000000604082015260600190565b6040518181016001600160401b038111828210171561093657fe5b604052919050565b60006001600160401b0382111561095157fe5b506020908102019056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122000cf659349e826a8a60fa58eb1d2936f8579046aa8bd6c0c31147f7919f3f44b64736f6c63430007060033",
}

// ExceptionStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ExceptionStorageMetaData.ABI instead.
var ExceptionStorageABI = ExceptionStorageMetaData.ABI

// ExceptionStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExceptionStorageMetaData.Bin instead.
var ExceptionStorageBin = ExceptionStorageMetaData.Bin

// DeployExceptionStorage deploys a new Ethereum contract, binding an instance of ExceptionStorage to it.
func DeployExceptionStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExceptionStorage, error) {
	parsed, err := ExceptionStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExceptionStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExceptionStorage{ExceptionStorageCaller: ExceptionStorageCaller{contract: contract}, ExceptionStorageTransactor: ExceptionStorageTransactor{contract: contract}, ExceptionStorageFilterer: ExceptionStorageFilterer{contract: contract}}, nil
}

// ExceptionStorage is an auto generated Go binding around an Ethereum contract.
type ExceptionStorage struct {
	ExceptionStorageCaller     // Read-only binding to the contract
	ExceptionStorageTransactor // Write-only binding to the contract
	ExceptionStorageFilterer   // Log filterer for contract events
}

// ExceptionStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExceptionStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExceptionStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExceptionStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExceptionStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExceptionStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExceptionStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExceptionStorageSession struct {
	Contract     *ExceptionStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExceptionStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExceptionStorageCallerSession struct {
	Contract *ExceptionStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExceptionStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExceptionStorageTransactorSession struct {
	Contract     *ExceptionStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExceptionStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExceptionStorageRaw struct {
	Contract *ExceptionStorage // Generic contract binding to access the raw methods on
}

// ExceptionStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExceptionStorageCallerRaw struct {
	Contract *ExceptionStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ExceptionStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExceptionStorageTransactorRaw struct {
	Contract *ExceptionStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExceptionStorage creates a new instance of ExceptionStorage, bound to a specific deployed contract.
func NewExceptionStorage(address common.Address, backend bind.ContractBackend) (*ExceptionStorage, error) {
	contract, err := bindExceptionStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExceptionStorage{ExceptionStorageCaller: ExceptionStorageCaller{contract: contract}, ExceptionStorageTransactor: ExceptionStorageTransactor{contract: contract}, ExceptionStorageFilterer: ExceptionStorageFilterer{contract: contract}}, nil
}

// NewExceptionStorageCaller creates a new read-only instance of ExceptionStorage, bound to a specific deployed contract.
func NewExceptionStorageCaller(address common.Address, caller bind.ContractCaller) (*ExceptionStorageCaller, error) {
	contract, err := bindExceptionStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExceptionStorageCaller{contract: contract}, nil
}

// NewExceptionStorageTransactor creates a new write-only instance of ExceptionStorage, bound to a specific deployed contract.
func NewExceptionStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ExceptionStorageTransactor, error) {
	contract, err := bindExceptionStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExceptionStorageTransactor{contract: contract}, nil
}

// NewExceptionStorageFilterer creates a new log filterer instance of ExceptionStorage, bound to a specific deployed contract.
func NewExceptionStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ExceptionStorageFilterer, error) {
	contract, err := bindExceptionStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExceptionStorageFilterer{contract: contract}, nil
}

// bindExceptionStorage binds a generic wrapper to an already deployed contract.
func bindExceptionStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExceptionStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExceptionStorage *ExceptionStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExceptionStorage.Contract.ExceptionStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExceptionStorage *ExceptionStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.ExceptionStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExceptionStorage *ExceptionStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.ExceptionStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExceptionStorage *ExceptionStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExceptionStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExceptionStorage *ExceptionStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExceptionStorage *ExceptionStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.contract.Transact(opts, method, params...)
}

// Exceptions is a free data retrieval call binding the contract method 0xdb13183e.
//
// Solidity: function exceptions(uint32 ) view returns(string)
func (_ExceptionStorage *ExceptionStorageCaller) Exceptions(opts *bind.CallOpts, arg0 uint32) (string, error) {
	var out []interface{}
	err := _ExceptionStorage.contract.Call(opts, &out, "exceptions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Exceptions is a free data retrieval call binding the contract method 0xdb13183e.
//
// Solidity: function exceptions(uint32 ) view returns(string)
func (_ExceptionStorage *ExceptionStorageSession) Exceptions(arg0 uint32) (string, error) {
	return _ExceptionStorage.Contract.Exceptions(&_ExceptionStorage.CallOpts, arg0)
}

// Exceptions is a free data retrieval call binding the contract method 0xdb13183e.
//
// Solidity: function exceptions(uint32 ) view returns(string)
func (_ExceptionStorage *ExceptionStorageCallerSession) Exceptions(arg0 uint32) (string, error) {
	return _ExceptionStorage.Contract.Exceptions(&_ExceptionStorage.CallOpts, arg0)
}

// GetException is a free data retrieval call binding the contract method 0x08f52c56.
//
// Solidity: function getException(uint32 key) view returns(string note)
func (_ExceptionStorage *ExceptionStorageCaller) GetException(opts *bind.CallOpts, key uint32) (string, error) {
	var out []interface{}
	err := _ExceptionStorage.contract.Call(opts, &out, "getException", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetException is a free data retrieval call binding the contract method 0x08f52c56.
//
// Solidity: function getException(uint32 key) view returns(string note)
func (_ExceptionStorage *ExceptionStorageSession) GetException(key uint32) (string, error) {
	return _ExceptionStorage.Contract.GetException(&_ExceptionStorage.CallOpts, key)
}

// GetException is a free data retrieval call binding the contract method 0x08f52c56.
//
// Solidity: function getException(uint32 key) view returns(string note)
func (_ExceptionStorage *ExceptionStorageCallerSession) GetException(key uint32) (string, error) {
	return _ExceptionStorage.Contract.GetException(&_ExceptionStorage.CallOpts, key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExceptionStorage *ExceptionStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExceptionStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExceptionStorage *ExceptionStorageSession) Owner() (common.Address, error) {
	return _ExceptionStorage.Contract.Owner(&_ExceptionStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExceptionStorage *ExceptionStorageCallerSession) Owner() (common.Address, error) {
	return _ExceptionStorage.Contract.Owner(&_ExceptionStorage.CallOpts)
}

// ThrowException is a free data retrieval call binding the contract method 0x524ec3ec.
//
// Solidity: function throwException(uint32 key) view returns()
func (_ExceptionStorage *ExceptionStorageCaller) ThrowException(opts *bind.CallOpts, key uint32) error {
	var out []interface{}
	err := _ExceptionStorage.contract.Call(opts, &out, "throwException", key)

	if err != nil {
		return err
	}

	return err

}

// ThrowException is a free data retrieval call binding the contract method 0x524ec3ec.
//
// Solidity: function throwException(uint32 key) view returns()
func (_ExceptionStorage *ExceptionStorageSession) ThrowException(key uint32) error {
	return _ExceptionStorage.Contract.ThrowException(&_ExceptionStorage.CallOpts, key)
}

// ThrowException is a free data retrieval call binding the contract method 0x524ec3ec.
//
// Solidity: function throwException(uint32 key) view returns()
func (_ExceptionStorage *ExceptionStorageCallerSession) ThrowException(key uint32) error {
	return _ExceptionStorage.Contract.ThrowException(&_ExceptionStorage.CallOpts, key)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExceptionStorage *ExceptionStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExceptionStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExceptionStorage *ExceptionStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExceptionStorage.Contract.RenounceOwnership(&_ExceptionStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExceptionStorage *ExceptionStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExceptionStorage.Contract.RenounceOwnership(&_ExceptionStorage.TransactOpts)
}

// SetException is a paid mutator transaction binding the contract method 0xd1c21d65.
//
// Solidity: function setException(uint32 key, string note) returns()
func (_ExceptionStorage *ExceptionStorageTransactor) SetException(opts *bind.TransactOpts, key uint32, note string) (*types.Transaction, error) {
	return _ExceptionStorage.contract.Transact(opts, "setException", key, note)
}

// SetException is a paid mutator transaction binding the contract method 0xd1c21d65.
//
// Solidity: function setException(uint32 key, string note) returns()
func (_ExceptionStorage *ExceptionStorageSession) SetException(key uint32, note string) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.SetException(&_ExceptionStorage.TransactOpts, key, note)
}

// SetException is a paid mutator transaction binding the contract method 0xd1c21d65.
//
// Solidity: function setException(uint32 key, string note) returns()
func (_ExceptionStorage *ExceptionStorageTransactorSession) SetException(key uint32, note string) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.SetException(&_ExceptionStorage.TransactOpts, key, note)
}

// SetExceptions is a paid mutator transaction binding the contract method 0xe150d645.
//
// Solidity: function setExceptions(uint32[] keys, string[] notes) returns()
func (_ExceptionStorage *ExceptionStorageTransactor) SetExceptions(opts *bind.TransactOpts, keys []uint32, notes []string) (*types.Transaction, error) {
	return _ExceptionStorage.contract.Transact(opts, "setExceptions", keys, notes)
}

// SetExceptions is a paid mutator transaction binding the contract method 0xe150d645.
//
// Solidity: function setExceptions(uint32[] keys, string[] notes) returns()
func (_ExceptionStorage *ExceptionStorageSession) SetExceptions(keys []uint32, notes []string) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.SetExceptions(&_ExceptionStorage.TransactOpts, keys, notes)
}

// SetExceptions is a paid mutator transaction binding the contract method 0xe150d645.
//
// Solidity: function setExceptions(uint32[] keys, string[] notes) returns()
func (_ExceptionStorage *ExceptionStorageTransactorSession) SetExceptions(keys []uint32, notes []string) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.SetExceptions(&_ExceptionStorage.TransactOpts, keys, notes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExceptionStorage *ExceptionStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExceptionStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExceptionStorage *ExceptionStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.TransferOwnership(&_ExceptionStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExceptionStorage *ExceptionStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExceptionStorage.Contract.TransferOwnership(&_ExceptionStorage.TransactOpts, newOwner)
}

// ExceptionStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExceptionStorage contract.
type ExceptionStorageOwnershipTransferredIterator struct {
	Event *ExceptionStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExceptionStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExceptionStorageOwnershipTransferred)
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
		it.Event = new(ExceptionStorageOwnershipTransferred)
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
func (it *ExceptionStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExceptionStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExceptionStorageOwnershipTransferred represents a OwnershipTransferred event raised by the ExceptionStorage contract.
type ExceptionStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExceptionStorage *ExceptionStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExceptionStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExceptionStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExceptionStorageOwnershipTransferredIterator{contract: _ExceptionStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExceptionStorage *ExceptionStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExceptionStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExceptionStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExceptionStorageOwnershipTransferred)
				if err := _ExceptionStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExceptionStorage *ExceptionStorageFilterer) ParseOwnershipTransferred(log types.Log) (*ExceptionStorageOwnershipTransferred, error) {
	event := new(ExceptionStorageOwnershipTransferred)
	if err := _ExceptionStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
