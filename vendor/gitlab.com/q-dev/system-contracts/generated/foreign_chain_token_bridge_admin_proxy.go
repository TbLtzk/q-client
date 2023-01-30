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

// ForeignChainTokenBridgeAdminProxyMetaData contains all meta data concerning the ForeignChainTokenBridgeAdminProxy contract.
var ForeignChainTokenBridgeAdminProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeValidators\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_newValidatorsList\",\"type\":\"address[]\"}],\"name\":\"updateTokenbridgeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a12806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063278dcd62146100c0578063715018a6146100d55780638da5cb5b146100dd578063c4d66de8146100fb578063f2fde38b1461010e575b33610060610121565b6001600160a01b03161461008f5760405162461bcd60e51b815260040161008690610716565b60405180910390fd5b6065546001600160a01b0316366000803760008036600080855af13d6000803e8080156100bb573d6000f35b3d6000fd5b6100d36100ce3660046107c9565b610130565b005b6100d3610414565b6100e5610121565b6040516100f29190610867565b60405180910390f35b6100d361010936600461087b565b61044f565b6100d361011c36600461087b565b6104df565b6033546001600160a01b031690565b33610139610121565b6001600160a01b03161461015f5760405162461bcd60e51b815260040161008690610716565b60655460408051635890ef7960e01b815290516000926001600160a01b031691635890ef799160048083019286929190829003018186803b1580156101a357600080fd5b505afa1580156101b7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101df919081019061089f565b905060005b81518110156102e85760008282815181106102015761020161092d565b60200260200101519050600080600090505b855181101561026a5785818151811061022e5761022e61092d565b60200260200101516001600160a01b0316836001600160a01b03161415610258576001915061026a565b8061026281610943565b915050610213565b50806102d3576065546040516340a141ff60e01b81526001600160a01b03909116906340a141ff906102a0908590600401610867565b600060405180830381600087803b1580156102ba57600080fd5b505af11580156102ce573d6000803e3d6000fd5b505050505b505080806102e090610943565b9150506101e4565b5060005b825181101561040f5760008382815181106103095761030961092d565b602090810291909101015160655460405163facd743b60e01b81529192506001600160a01b03169063facd743b90610345908490600401610867565b60206040518083038186803b15801561035d57600080fd5b505afa158015610371573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610395919061096c565b6103fc57606554604051632691c64760e11b81526001600160a01b0390911690634d238c8e906103c9908490600401610867565b600060405180830381600087803b1580156103e357600080fd5b505af11580156103f7573d6000803e3d6000fd5b505050505b508061040781610943565b9150506102ec565b505050565b3361041d610121565b6001600160a01b0316146104435760405162461bcd60e51b815260040161008690610716565b61044d600061057f565b565b600054610100900460ff1680610468575060005460ff16155b6104845760405162461bcd60e51b81526004016100869061098e565b600054610100900460ff161580156104a6576000805461ffff19166101011790555b606580546001600160a01b0319166001600160a01b0384161790556104c96105d1565b80156104db576000805461ff00191690555b5050565b336104e8610121565b6001600160a01b03161461050e5760405162461bcd60e51b815260040161008690610716565b6001600160a01b0381166105735760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610086565b61057c8161057f565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16806105ea575060005460ff16155b6106065760405162461bcd60e51b81526004016100869061098e565b600054610100900460ff16158015610628576000805461ffff19166101011790555b61063061064c565b6106386106b6565b801561057c576000805461ff001916905550565b600054610100900460ff1680610665575060005460ff16155b6106815760405162461bcd60e51b81526004016100869061098e565b600054610100900460ff16158015610638576000805461ffff1916610101179055801561057c576000805461ff001916905550565b600054610100900460ff16806106cf575060005460ff16155b6106eb5760405162461bcd60e51b81526004016100869061098e565b600054610100900460ff1615801561070d576000805461ffff19166101011790555b6106383361057f565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156107895761078961074b565b604052919050565b60006001600160401b038211156107aa576107aa61074b565b5060051b60200190565b6001600160a01b038116811461057c57600080fd5b600060208083850312156107dc57600080fd5b82356001600160401b038111156107f257600080fd5b8301601f8101851361080357600080fd5b803561081661081182610791565b610761565b81815260059190911b8201830190838101908783111561083557600080fd5b928401925b8284101561085c57833561084d816107b4565b8252928401929084019061083a565b979650505050505050565b6001600160a01b0391909116815260200190565b60006020828403121561088d57600080fd5b8135610898816107b4565b9392505050565b600060208083850312156108b257600080fd5b82516001600160401b038111156108c857600080fd5b8301601f810185136108d957600080fd5b80516108e761081182610791565b81815260059190911b8201830190838101908783111561090657600080fd5b928401925b8284101561085c57835161091e816107b4565b8252928401929084019061090b565b634e487b7160e01b600052603260045260246000fd5b600060001982141561096557634e487b7160e01b600052601160045260246000fd5b5060010190565b60006020828403121561097e57600080fd5b8151801515811461089857600080fd5b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b60608201526080019056fea2646970667358221220b7215ea67040a2f1f9d310455eadc2964c69ac741f960a77ce2bff9b5588111c64736f6c63430008090033",
}

// ForeignChainTokenBridgeAdminProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ForeignChainTokenBridgeAdminProxyMetaData.ABI instead.
var ForeignChainTokenBridgeAdminProxyABI = ForeignChainTokenBridgeAdminProxyMetaData.ABI

// ForeignChainTokenBridgeAdminProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ForeignChainTokenBridgeAdminProxyMetaData.Bin instead.
var ForeignChainTokenBridgeAdminProxyBin = ForeignChainTokenBridgeAdminProxyMetaData.Bin

// DeployForeignChainTokenBridgeAdminProxy deploys a new Ethereum contract, binding an instance of ForeignChainTokenBridgeAdminProxy to it.
func DeployForeignChainTokenBridgeAdminProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ForeignChainTokenBridgeAdminProxy, error) {
	parsed, err := ForeignChainTokenBridgeAdminProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ForeignChainTokenBridgeAdminProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForeignChainTokenBridgeAdminProxy{ForeignChainTokenBridgeAdminProxyCaller: ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, ForeignChainTokenBridgeAdminProxyTransactor: ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, ForeignChainTokenBridgeAdminProxyFilterer: ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// ForeignChainTokenBridgeAdminProxy is an auto generated Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxy struct {
	ForeignChainTokenBridgeAdminProxyCaller     // Read-only binding to the contract
	ForeignChainTokenBridgeAdminProxyTransactor // Write-only binding to the contract
	ForeignChainTokenBridgeAdminProxyFilterer   // Log filterer for contract events
}

// ForeignChainTokenBridgeAdminProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForeignChainTokenBridgeAdminProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignChainTokenBridgeAdminProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForeignChainTokenBridgeAdminProxySession struct {
	Contract     *ForeignChainTokenBridgeAdminProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForeignChainTokenBridgeAdminProxyCallerSession struct {
	Contract *ForeignChainTokenBridgeAdminProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForeignChainTokenBridgeAdminProxyTransactorSession struct {
	Contract     *ForeignChainTokenBridgeAdminProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// ForeignChainTokenBridgeAdminProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxy // Generic contract binding to access the raw methods on
}

// ForeignChainTokenBridgeAdminProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyCallerRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ForeignChainTokenBridgeAdminProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForeignChainTokenBridgeAdminProxyTransactorRaw struct {
	Contract *ForeignChainTokenBridgeAdminProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForeignChainTokenBridgeAdminProxy creates a new instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxy(address common.Address, backend bind.ContractBackend) (*ForeignChainTokenBridgeAdminProxy, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxy{ForeignChainTokenBridgeAdminProxyCaller: ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, ForeignChainTokenBridgeAdminProxyTransactor: ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, ForeignChainTokenBridgeAdminProxyFilterer: ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}}, nil
}

// NewForeignChainTokenBridgeAdminProxyCaller creates a new read-only instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyCaller(address common.Address, caller bind.ContractCaller) (*ForeignChainTokenBridgeAdminProxyCaller, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyCaller{contract: contract}, nil
}

// NewForeignChainTokenBridgeAdminProxyTransactor creates a new write-only instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ForeignChainTokenBridgeAdminProxyTransactor, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyTransactor{contract: contract}, nil
}

// NewForeignChainTokenBridgeAdminProxyFilterer creates a new log filterer instance of ForeignChainTokenBridgeAdminProxy, bound to a specific deployed contract.
func NewForeignChainTokenBridgeAdminProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ForeignChainTokenBridgeAdminProxyFilterer, error) {
	contract, err := bindForeignChainTokenBridgeAdminProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyFilterer{contract: contract}, nil
}

// bindForeignChainTokenBridgeAdminProxy binds a generic wrapper to an already deployed contract.
func bindForeignChainTokenBridgeAdminProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForeignChainTokenBridgeAdminProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.ForeignChainTokenBridgeAdminProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ForeignChainTokenBridgeAdminProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Owner() (common.Address, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Owner(&_ForeignChainTokenBridgeAdminProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyCallerSession) Owner() (common.Address, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Owner(&_ForeignChainTokenBridgeAdminProxy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) Initialize(opts *bind.TransactOpts, _bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "initialize", _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Initialize(_bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Initialize(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _bridgeValidators)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridgeValidators) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) Initialize(_bridgeValidators common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Initialize(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _bridgeValidators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.RenounceOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.RenounceOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.TransferOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.TransferOwnership(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, newOwner)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) UpdateTokenbridgeValidators(opts *bind.TransactOpts, _newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.Transact(opts, "updateTokenbridgeValidators", _newValidatorsList)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) UpdateTokenbridgeValidators(_newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _newValidatorsList)
}

// UpdateTokenbridgeValidators is a paid mutator transaction binding the contract method 0x278dcd62.
//
// Solidity: function updateTokenbridgeValidators(address[] _newValidatorsList) returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) UpdateTokenbridgeValidators(_newValidatorsList []common.Address) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.UpdateTokenbridgeValidators(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, _newValidatorsList)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Fallback(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ForeignChainTokenBridgeAdminProxy.Contract.Fallback(&_ForeignChainTokenBridgeAdminProxy.TransactOpts, calldata)
}

// ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
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
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyOwnershipTransferred represents a OwnershipTransferred event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyOwnershipTransferredIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseOwnershipTransferred(log types.Log) (*ForeignChainTokenBridgeAdminProxyOwnershipTransferred, error) {
	event := new(ForeignChainTokenBridgeAdminProxyOwnershipTransferred)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForeignChainTokenBridgeAdminProxyValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorAddedIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyValidatorAdded // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
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
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyValidatorAdded represents a ValidatorAdded event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterValidatorAdded(opts *bind.FilterOpts, _validator []common.Address) (*ForeignChainTokenBridgeAdminProxyValidatorAddedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyValidatorAddedIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyValidatorAdded, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorAdded", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseValidatorAdded(log types.Log) (*ForeignChainTokenBridgeAdminProxyValidatorAdded, error) {
	event := new(ForeignChainTokenBridgeAdminProxyValidatorAdded)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator struct {
	Event *ForeignChainTokenBridgeAdminProxyValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
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
		it.Event = new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
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
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignChainTokenBridgeAdminProxyValidatorRemoved represents a ValidatorRemoved event raised by the ForeignChainTokenBridgeAdminProxy contract.
type ForeignChainTokenBridgeAdminProxyValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, _validator []common.Address) (*ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.FilterLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ForeignChainTokenBridgeAdminProxyValidatorRemovedIterator{contract: _ForeignChainTokenBridgeAdminProxy.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ForeignChainTokenBridgeAdminProxyValidatorRemoved, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _ForeignChainTokenBridgeAdminProxy.contract.WatchLogs(opts, "ValidatorRemoved", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
				if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed _validator)
func (_ForeignChainTokenBridgeAdminProxy *ForeignChainTokenBridgeAdminProxyFilterer) ParseValidatorRemoved(log types.Log) (*ForeignChainTokenBridgeAdminProxyValidatorRemoved, error) {
	event := new(ForeignChainTokenBridgeAdminProxyValidatorRemoved)
	if err := _ForeignChainTokenBridgeAdminProxy.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
