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
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6109e08061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806308f52c561461007d578063524ec3ec146100a6578063715018a6146100bb5780638da5cb5b146100c3578063d1c21d65146100e3578063db13183e146100f6578063e150d64514610109578063f2fde38b1461011c575b600080fd5b61009061008b3660046105f2565b61012f565b60405161009d9190610614565b60405180910390f35b6100b96100b43660046105f2565b610212565b005b6100b961023d565b6100cb610278565b6040516001600160a01b03909116815260200161009d565b6100b96100f136600461071e565b610287565b6100906101043660046105f2565b6102e1565b6100b961011736600461081d565b61037b565b6100b961012a3660046108d2565b610455565b63ffffffff81166000908152600160205260409020805460609190610153906108fb565b80601f016020809104026020016040519081016040528092919081815260200182805461017f906108fb565b80156101cc5780601f106101a1576101008083540402835291602001916101cc565b820191906000526020600020905b8154815290600101906020018083116101af57829003601f168201915b5050505050905080516000141561020d57506040805180820190915260178152764572726f72207769746820756e6b6e6f776e20636f646560481b60208201525b919050565b61021b8161012f565b60405162461bcd60e51b81526004016102349190610614565b60405180910390fd5b33610246610278565b6001600160a01b03161461026c5760405162461bcd60e51b815260040161023490610936565b61027660006104f5565b565b6000546001600160a01b031690565b33610290610278565b6001600160a01b0316146102b65760405162461bcd60e51b815260040161023490610936565b63ffffffff8216600090815260016020908152604090912082516102dc92840190610545565b505050565b600160205260009081526040902080546102fa906108fb565b80601f0160208091040260200160405190810160405280929190818152602001828054610326906108fb565b80156103735780601f1061034857610100808354040283529160200191610373565b820191906000526020600020905b81548152906001019060200180831161035657829003601f168201915b505050505081565b33610384610278565b6001600160a01b0316146103aa5760405162461bcd60e51b815260040161023490610936565b80518251146103fb5760405162461bcd60e51b815260206004820152601c60248201527f6172726179206c656e6774686573206d75737420626520657175616c000000006044820152606401610234565b60005b82518110156102dc5761044383828151811061041c5761041c61096b565b60200260200101518383815181106104365761043661096b565b6020026020010151610287565b8061044d81610981565b9150506103fe565b3361045e610278565b6001600160a01b0316146104845760405162461bcd60e51b815260040161023490610936565b6001600160a01b0381166104e95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610234565b6104f2816104f5565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054610551906108fb565b90600052602060002090601f01602090048101928261057357600085556105b9565b82601f1061058c57805160ff19168380011785556105b9565b828001600101855582156105b9579182015b828111156105b957825182559160200191906001019061059e565b506105c59291506105c9565b5090565b5b808211156105c557600081556001016105ca565b803563ffffffff8116811461020d57600080fd5b60006020828403121561060457600080fd5b61060d826105de565b9392505050565b600060208083528351808285015260005b8181101561064157858101830151858201604001528201610625565b81811115610653576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156106a7576106a7610669565b604052919050565b600082601f8301126106c057600080fd5b81356001600160401b038111156106d9576106d9610669565b6106ec601f8201601f191660200161067f565b81815284602083860101111561070157600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561073157600080fd5b61073a836105de565b915060208301356001600160401b0381111561075557600080fd5b610761858286016106af565b9150509250929050565b60006001600160401b0382111561078457610784610669565b5060051b60200190565b600082601f83011261079f57600080fd5b813560206107b46107af8361076b565b61067f565b82815260059290921b840181019181810190868411156107d357600080fd5b8286015b848110156108125780356001600160401b038111156107f65760008081fd5b6108048986838b01016106af565b8452509183019183016107d7565b509695505050505050565b6000806040838503121561083057600080fd5b82356001600160401b038082111561084757600080fd5b818501915085601f83011261085b57600080fd5b8135602061086b6107af8361076b565b82815260059290921b8401810191818101908984111561088a57600080fd5b948201945b838610156108af576108a0866105de565b8252948201949082019061088f565b965050860135925050808211156108c557600080fd5b506107618582860161078e565b6000602082840312156108e457600080fd5b81356001600160a01b038116811461060d57600080fd5b600181811c9082168061090f57607f821691505b6020821081141561093057634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156109a357634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122079185f17dee1805544d4949e958ee8f2429dd80fa393118dc05c1920132f9bba64736f6c63430008090033",
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
