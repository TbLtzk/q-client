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

// AccountAliasesMetaData contains all meta data concerning the AccountAliases contract.
var AccountAliasesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"AliasUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"}],\"name\":\"Reserved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aliases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_main\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_role\",\"type\":\"uint256[]\"}],\"name\":\"resolveBatch\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_alias\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_alias\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_role\",\"type\":\"uint256[]\"}],\"name\":\"resolveBatchReverse\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_main\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"resolveReverse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reverseAliases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"setAlias\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506108ac806100206000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c80632425c56e1461007d5780639f0c6fd2146100cb578063a4dc3302146100e0578063abadc6f514610100578063d97b4de714610113578063e75179a414610126578063eccf800314610139578063eea840bd1461014c575b600080fd5b6100ae61008b36600461060b565b60006020818152928152604080822090935290815220546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100de6100d936600461060b565b610175565b005b6100f36100ee366004610709565b6102e8565b6040516100c291906107c8565b6100ae61010e36600461060b565b6103c6565b6100f3610121366004610709565b610408565b6100de610134366004610815565b6104df565b6100ae61014736600461060b565b610535565b6100ae61015a366004610815565b6001602052600090815260409020546001600160a01b031681565b6001600160a01b03828116600090815260016020526040902054166101ed5760405162461bcd60e51b815260206004820152602360248201527f5b5145432d3034303030305d2d416c696173206973206e6f742072657365727660448201526232b21760e91b60648201526084015b60405180910390fd5b6001600160a01b038281166000908152600160205260409020541633146102725760405162461bcd60e51b815260206004820152603360248201527f5b5145432d3034303030325d2d416c696173206973207265736572766564206660448201527237b91030b737ba3432b91030b1b1b7bab73a1760691b60648201526084016101e4565b33600081815260208181526040808320858452825280832080546001600160a01b0388166001600160a01b0319918216811790925581855260019093528184208054909316851790925551849391927f2aa9378fc153640a84c4851f2e9d5b9a9c6660640cc5f6c01b3574e69ab765cf91a45050565b60606102f683518351610576565b82516001600160401b0381111561030f5761030f610635565b604051908082528060200260200182016040528015610338578160200160208202803683370190505b50905060005b83518110156103bf5761038384828151811061035c5761035c610837565b602002602001015184838151811061037657610376610837565b60200260200101516103c6565b82828151811061039557610395610837565b6001600160a01b0390921660209283029190910190910152806103b78161084d565b91505061033e565b5092915050565b6001600160a01b038083166000818152602081815260408083208684528252808320548516808452600190925290912054909216146104025750815b92915050565b606061041683518351610576565b82516001600160401b0381111561042f5761042f610635565b604051908082528060200260200182016040528015610458578160200160208202803683370190505b50905060005b83518110156103bf576104a384828151811061047c5761047c610837565b602002602001015184838151811061049657610496610837565b6020026020010151610535565b8282815181106104b5576104b5610837565b6001600160a01b0390921660209283029190910190910152806104d78161084d565b91505061045e565b3360008181526001602052604080822080546001600160a01b0319166001600160a01b038616908117909155905190917fd0e1be92289836d3c01bb059be619bebe3f7818670bb0ba1b2656816330ffc6e91a350565b6001600160a01b0380831660008181526001602090815260408083205485168084528383528184208785529092529091205490921614610402575090919050565b8082146105eb5760405162461bcd60e51b815260206004820152603f60248201527f5b5145432d3034303030315d2d4c656e677468206f662061646472657373657360448201527f20616e6420726f6c6520617272617973206d75737420626520657175616c2e0060648201526084016101e4565b5050565b80356001600160a01b038116811461060657600080fd5b919050565b6000806040838503121561061e57600080fd5b610627836105ef565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561067357610673610635565b604052919050565b60006001600160401b0382111561069457610694610635565b5060051b60200190565b600082601f8301126106af57600080fd5b813560206106c46106bf8361067b565b61064b565b82815260059290921b840181019181810190868411156106e357600080fd5b8286015b848110156106fe57803583529183019183016106e7565b509695505050505050565b6000806040838503121561071c57600080fd5b82356001600160401b038082111561073357600080fd5b818501915085601f83011261074757600080fd5b813560206107576106bf8361067b565b82815260059290921b8401810191818101908984111561077657600080fd5b948201945b8386101561079b5761078c866105ef565b8252948201949082019061077b565b965050860135925050808211156107b157600080fd5b506107be8582860161069e565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156108095783516001600160a01b0316835292840192918401916001016107e4565b50909695505050505050565b60006020828403121561082757600080fd5b610830826105ef565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561086f57634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212207ade270f07cd16559b4785ee127b75151d0ccb68246c0b8f7224848ae6df10dd64736f6c63430008090033",
}

// AccountAliasesABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountAliasesMetaData.ABI instead.
var AccountAliasesABI = AccountAliasesMetaData.ABI

// AccountAliasesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountAliasesMetaData.Bin instead.
var AccountAliasesBin = AccountAliasesMetaData.Bin

// DeployAccountAliases deploys a new Ethereum contract, binding an instance of AccountAliases to it.
func DeployAccountAliases(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountAliases, error) {
	parsed, err := AccountAliasesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountAliasesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountAliases{AccountAliasesCaller: AccountAliasesCaller{contract: contract}, AccountAliasesTransactor: AccountAliasesTransactor{contract: contract}, AccountAliasesFilterer: AccountAliasesFilterer{contract: contract}}, nil
}

// AccountAliases is an auto generated Go binding around an Ethereum contract.
type AccountAliases struct {
	AccountAliasesCaller     // Read-only binding to the contract
	AccountAliasesTransactor // Write-only binding to the contract
	AccountAliasesFilterer   // Log filterer for contract events
}

// AccountAliasesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountAliasesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountAliasesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountAliasesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountAliasesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountAliasesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountAliasesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountAliasesSession struct {
	Contract     *AccountAliases   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountAliasesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountAliasesCallerSession struct {
	Contract *AccountAliasesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountAliasesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountAliasesTransactorSession struct {
	Contract     *AccountAliasesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountAliasesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountAliasesRaw struct {
	Contract *AccountAliases // Generic contract binding to access the raw methods on
}

// AccountAliasesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountAliasesCallerRaw struct {
	Contract *AccountAliasesCaller // Generic read-only contract binding to access the raw methods on
}

// AccountAliasesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountAliasesTransactorRaw struct {
	Contract *AccountAliasesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountAliases creates a new instance of AccountAliases, bound to a specific deployed contract.
func NewAccountAliases(address common.Address, backend bind.ContractBackend) (*AccountAliases, error) {
	contract, err := bindAccountAliases(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountAliases{AccountAliasesCaller: AccountAliasesCaller{contract: contract}, AccountAliasesTransactor: AccountAliasesTransactor{contract: contract}, AccountAliasesFilterer: AccountAliasesFilterer{contract: contract}}, nil
}

// NewAccountAliasesCaller creates a new read-only instance of AccountAliases, bound to a specific deployed contract.
func NewAccountAliasesCaller(address common.Address, caller bind.ContractCaller) (*AccountAliasesCaller, error) {
	contract, err := bindAccountAliases(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountAliasesCaller{contract: contract}, nil
}

// NewAccountAliasesTransactor creates a new write-only instance of AccountAliases, bound to a specific deployed contract.
func NewAccountAliasesTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountAliasesTransactor, error) {
	contract, err := bindAccountAliases(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountAliasesTransactor{contract: contract}, nil
}

// NewAccountAliasesFilterer creates a new log filterer instance of AccountAliases, bound to a specific deployed contract.
func NewAccountAliasesFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountAliasesFilterer, error) {
	contract, err := bindAccountAliases(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountAliasesFilterer{contract: contract}, nil
}

// bindAccountAliases binds a generic wrapper to an already deployed contract.
func bindAccountAliases(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountAliasesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountAliases *AccountAliasesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountAliases.Contract.AccountAliasesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountAliases *AccountAliasesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountAliases.Contract.AccountAliasesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountAliases *AccountAliasesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountAliases.Contract.AccountAliasesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountAliases *AccountAliasesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountAliases.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountAliases *AccountAliasesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountAliases.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountAliases *AccountAliasesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountAliases.Contract.contract.Transact(opts, method, params...)
}

// Aliases is a free data retrieval call binding the contract method 0x2425c56e.
//
// Solidity: function aliases(address , uint256 ) view returns(address)
func (_AccountAliases *AccountAliasesCaller) Aliases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "aliases", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aliases is a free data retrieval call binding the contract method 0x2425c56e.
//
// Solidity: function aliases(address , uint256 ) view returns(address)
func (_AccountAliases *AccountAliasesSession) Aliases(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.Aliases(&_AccountAliases.CallOpts, arg0, arg1)
}

// Aliases is a free data retrieval call binding the contract method 0x2425c56e.
//
// Solidity: function aliases(address , uint256 ) view returns(address)
func (_AccountAliases *AccountAliasesCallerSession) Aliases(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.Aliases(&_AccountAliases.CallOpts, arg0, arg1)
}

// Resolve is a free data retrieval call binding the contract method 0xabadc6f5.
//
// Solidity: function resolve(address _main, uint256 _role) view returns(address _alias)
func (_AccountAliases *AccountAliasesCaller) Resolve(opts *bind.CallOpts, _main common.Address, _role *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "resolve", _main, _role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0xabadc6f5.
//
// Solidity: function resolve(address _main, uint256 _role) view returns(address _alias)
func (_AccountAliases *AccountAliasesSession) Resolve(_main common.Address, _role *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.Resolve(&_AccountAliases.CallOpts, _main, _role)
}

// Resolve is a free data retrieval call binding the contract method 0xabadc6f5.
//
// Solidity: function resolve(address _main, uint256 _role) view returns(address _alias)
func (_AccountAliases *AccountAliasesCallerSession) Resolve(_main common.Address, _role *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.Resolve(&_AccountAliases.CallOpts, _main, _role)
}

// ResolveBatch is a free data retrieval call binding the contract method 0xa4dc3302.
//
// Solidity: function resolveBatch(address[] _main, uint256[] _role) view returns(address[] _alias)
func (_AccountAliases *AccountAliasesCaller) ResolveBatch(opts *bind.CallOpts, _main []common.Address, _role []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "resolveBatch", _main, _role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ResolveBatch is a free data retrieval call binding the contract method 0xa4dc3302.
//
// Solidity: function resolveBatch(address[] _main, uint256[] _role) view returns(address[] _alias)
func (_AccountAliases *AccountAliasesSession) ResolveBatch(_main []common.Address, _role []*big.Int) ([]common.Address, error) {
	return _AccountAliases.Contract.ResolveBatch(&_AccountAliases.CallOpts, _main, _role)
}

// ResolveBatch is a free data retrieval call binding the contract method 0xa4dc3302.
//
// Solidity: function resolveBatch(address[] _main, uint256[] _role) view returns(address[] _alias)
func (_AccountAliases *AccountAliasesCallerSession) ResolveBatch(_main []common.Address, _role []*big.Int) ([]common.Address, error) {
	return _AccountAliases.Contract.ResolveBatch(&_AccountAliases.CallOpts, _main, _role)
}

// ResolveBatchReverse is a free data retrieval call binding the contract method 0xd97b4de7.
//
// Solidity: function resolveBatchReverse(address[] _alias, uint256[] _role) view returns(address[] _main)
func (_AccountAliases *AccountAliasesCaller) ResolveBatchReverse(opts *bind.CallOpts, _alias []common.Address, _role []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "resolveBatchReverse", _alias, _role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ResolveBatchReverse is a free data retrieval call binding the contract method 0xd97b4de7.
//
// Solidity: function resolveBatchReverse(address[] _alias, uint256[] _role) view returns(address[] _main)
func (_AccountAliases *AccountAliasesSession) ResolveBatchReverse(_alias []common.Address, _role []*big.Int) ([]common.Address, error) {
	return _AccountAliases.Contract.ResolveBatchReverse(&_AccountAliases.CallOpts, _alias, _role)
}

// ResolveBatchReverse is a free data retrieval call binding the contract method 0xd97b4de7.
//
// Solidity: function resolveBatchReverse(address[] _alias, uint256[] _role) view returns(address[] _main)
func (_AccountAliases *AccountAliasesCallerSession) ResolveBatchReverse(_alias []common.Address, _role []*big.Int) ([]common.Address, error) {
	return _AccountAliases.Contract.ResolveBatchReverse(&_AccountAliases.CallOpts, _alias, _role)
}

// ResolveReverse is a free data retrieval call binding the contract method 0xeccf8003.
//
// Solidity: function resolveReverse(address _alias, uint256 _role) view returns(address _main)
func (_AccountAliases *AccountAliasesCaller) ResolveReverse(opts *bind.CallOpts, _alias common.Address, _role *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "resolveReverse", _alias, _role)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolveReverse is a free data retrieval call binding the contract method 0xeccf8003.
//
// Solidity: function resolveReverse(address _alias, uint256 _role) view returns(address _main)
func (_AccountAliases *AccountAliasesSession) ResolveReverse(_alias common.Address, _role *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.ResolveReverse(&_AccountAliases.CallOpts, _alias, _role)
}

// ResolveReverse is a free data retrieval call binding the contract method 0xeccf8003.
//
// Solidity: function resolveReverse(address _alias, uint256 _role) view returns(address _main)
func (_AccountAliases *AccountAliasesCallerSession) ResolveReverse(_alias common.Address, _role *big.Int) (common.Address, error) {
	return _AccountAliases.Contract.ResolveReverse(&_AccountAliases.CallOpts, _alias, _role)
}

// ReverseAliases is a free data retrieval call binding the contract method 0xeea840bd.
//
// Solidity: function reverseAliases(address ) view returns(address)
func (_AccountAliases *AccountAliasesCaller) ReverseAliases(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AccountAliases.contract.Call(opts, &out, "reverseAliases", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReverseAliases is a free data retrieval call binding the contract method 0xeea840bd.
//
// Solidity: function reverseAliases(address ) view returns(address)
func (_AccountAliases *AccountAliasesSession) ReverseAliases(arg0 common.Address) (common.Address, error) {
	return _AccountAliases.Contract.ReverseAliases(&_AccountAliases.CallOpts, arg0)
}

// ReverseAliases is a free data retrieval call binding the contract method 0xeea840bd.
//
// Solidity: function reverseAliases(address ) view returns(address)
func (_AccountAliases *AccountAliasesCallerSession) ReverseAliases(arg0 common.Address) (common.Address, error) {
	return _AccountAliases.Contract.ReverseAliases(&_AccountAliases.CallOpts, arg0)
}

// Reserve is a paid mutator transaction binding the contract method 0xe75179a4.
//
// Solidity: function reserve(address _newOwner) returns()
func (_AccountAliases *AccountAliasesTransactor) Reserve(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AccountAliases.contract.Transact(opts, "reserve", _newOwner)
}

// Reserve is a paid mutator transaction binding the contract method 0xe75179a4.
//
// Solidity: function reserve(address _newOwner) returns()
func (_AccountAliases *AccountAliasesSession) Reserve(_newOwner common.Address) (*types.Transaction, error) {
	return _AccountAliases.Contract.Reserve(&_AccountAliases.TransactOpts, _newOwner)
}

// Reserve is a paid mutator transaction binding the contract method 0xe75179a4.
//
// Solidity: function reserve(address _newOwner) returns()
func (_AccountAliases *AccountAliasesTransactorSession) Reserve(_newOwner common.Address) (*types.Transaction, error) {
	return _AccountAliases.Contract.Reserve(&_AccountAliases.TransactOpts, _newOwner)
}

// SetAlias is a paid mutator transaction binding the contract method 0x9f0c6fd2.
//
// Solidity: function setAlias(address _alias, uint256 _role) returns()
func (_AccountAliases *AccountAliasesTransactor) SetAlias(opts *bind.TransactOpts, _alias common.Address, _role *big.Int) (*types.Transaction, error) {
	return _AccountAliases.contract.Transact(opts, "setAlias", _alias, _role)
}

// SetAlias is a paid mutator transaction binding the contract method 0x9f0c6fd2.
//
// Solidity: function setAlias(address _alias, uint256 _role) returns()
func (_AccountAliases *AccountAliasesSession) SetAlias(_alias common.Address, _role *big.Int) (*types.Transaction, error) {
	return _AccountAliases.Contract.SetAlias(&_AccountAliases.TransactOpts, _alias, _role)
}

// SetAlias is a paid mutator transaction binding the contract method 0x9f0c6fd2.
//
// Solidity: function setAlias(address _alias, uint256 _role) returns()
func (_AccountAliases *AccountAliasesTransactorSession) SetAlias(_alias common.Address, _role *big.Int) (*types.Transaction, error) {
	return _AccountAliases.Contract.SetAlias(&_AccountAliases.TransactOpts, _alias, _role)
}

// AccountAliasesAliasUpdatedIterator is returned from FilterAliasUpdated and is used to iterate over the raw logs and unpacked data for AliasUpdated events raised by the AccountAliases contract.
type AccountAliasesAliasUpdatedIterator struct {
	Event *AccountAliasesAliasUpdated // Event containing the contract specifics and raw log

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
func (it *AccountAliasesAliasUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAliasesAliasUpdated)
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
		it.Event = new(AccountAliasesAliasUpdated)
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
func (it *AccountAliasesAliasUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAliasesAliasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAliasesAliasUpdated represents a AliasUpdated event raised by the AccountAliases contract.
type AccountAliasesAliasUpdated struct {
	Main  common.Address
	Alias common.Address
	Role  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAliasUpdated is a free log retrieval operation binding the contract event 0x2aa9378fc153640a84c4851f2e9d5b9a9c6660640cc5f6c01b3574e69ab765cf.
//
// Solidity: event AliasUpdated(address indexed _main, address indexed _alias, uint256 indexed role)
func (_AccountAliases *AccountAliasesFilterer) FilterAliasUpdated(opts *bind.FilterOpts, _main []common.Address, _alias []common.Address, role []*big.Int) (*AccountAliasesAliasUpdatedIterator, error) {

	var _mainRule []interface{}
	for _, _mainItem := range _main {
		_mainRule = append(_mainRule, _mainItem)
	}
	var _aliasRule []interface{}
	for _, _aliasItem := range _alias {
		_aliasRule = append(_aliasRule, _aliasItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccountAliases.contract.FilterLogs(opts, "AliasUpdated", _mainRule, _aliasRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &AccountAliasesAliasUpdatedIterator{contract: _AccountAliases.contract, event: "AliasUpdated", logs: logs, sub: sub}, nil
}

// WatchAliasUpdated is a free log subscription operation binding the contract event 0x2aa9378fc153640a84c4851f2e9d5b9a9c6660640cc5f6c01b3574e69ab765cf.
//
// Solidity: event AliasUpdated(address indexed _main, address indexed _alias, uint256 indexed role)
func (_AccountAliases *AccountAliasesFilterer) WatchAliasUpdated(opts *bind.WatchOpts, sink chan<- *AccountAliasesAliasUpdated, _main []common.Address, _alias []common.Address, role []*big.Int) (event.Subscription, error) {

	var _mainRule []interface{}
	for _, _mainItem := range _main {
		_mainRule = append(_mainRule, _mainItem)
	}
	var _aliasRule []interface{}
	for _, _aliasItem := range _alias {
		_aliasRule = append(_aliasRule, _aliasItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _AccountAliases.contract.WatchLogs(opts, "AliasUpdated", _mainRule, _aliasRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAliasesAliasUpdated)
				if err := _AccountAliases.contract.UnpackLog(event, "AliasUpdated", log); err != nil {
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

// ParseAliasUpdated is a log parse operation binding the contract event 0x2aa9378fc153640a84c4851f2e9d5b9a9c6660640cc5f6c01b3574e69ab765cf.
//
// Solidity: event AliasUpdated(address indexed _main, address indexed _alias, uint256 indexed role)
func (_AccountAliases *AccountAliasesFilterer) ParseAliasUpdated(log types.Log) (*AccountAliasesAliasUpdated, error) {
	event := new(AccountAliasesAliasUpdated)
	if err := _AccountAliases.contract.UnpackLog(event, "AliasUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountAliasesReservedIterator is returned from FilterReserved and is used to iterate over the raw logs and unpacked data for Reserved events raised by the AccountAliases contract.
type AccountAliasesReservedIterator struct {
	Event *AccountAliasesReserved // Event containing the contract specifics and raw log

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
func (it *AccountAliasesReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAliasesReserved)
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
		it.Event = new(AccountAliasesReserved)
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
func (it *AccountAliasesReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAliasesReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAliasesReserved represents a Reserved event raised by the AccountAliases contract.
type AccountAliasesReserved struct {
	Main  common.Address
	Alias common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReserved is a free log retrieval operation binding the contract event 0xd0e1be92289836d3c01bb059be619bebe3f7818670bb0ba1b2656816330ffc6e.
//
// Solidity: event Reserved(address indexed _main, address indexed _alias)
func (_AccountAliases *AccountAliasesFilterer) FilterReserved(opts *bind.FilterOpts, _main []common.Address, _alias []common.Address) (*AccountAliasesReservedIterator, error) {

	var _mainRule []interface{}
	for _, _mainItem := range _main {
		_mainRule = append(_mainRule, _mainItem)
	}
	var _aliasRule []interface{}
	for _, _aliasItem := range _alias {
		_aliasRule = append(_aliasRule, _aliasItem)
	}

	logs, sub, err := _AccountAliases.contract.FilterLogs(opts, "Reserved", _mainRule, _aliasRule)
	if err != nil {
		return nil, err
	}
	return &AccountAliasesReservedIterator{contract: _AccountAliases.contract, event: "Reserved", logs: logs, sub: sub}, nil
}

// WatchReserved is a free log subscription operation binding the contract event 0xd0e1be92289836d3c01bb059be619bebe3f7818670bb0ba1b2656816330ffc6e.
//
// Solidity: event Reserved(address indexed _main, address indexed _alias)
func (_AccountAliases *AccountAliasesFilterer) WatchReserved(opts *bind.WatchOpts, sink chan<- *AccountAliasesReserved, _main []common.Address, _alias []common.Address) (event.Subscription, error) {

	var _mainRule []interface{}
	for _, _mainItem := range _main {
		_mainRule = append(_mainRule, _mainItem)
	}
	var _aliasRule []interface{}
	for _, _aliasItem := range _alias {
		_aliasRule = append(_aliasRule, _aliasItem)
	}

	logs, sub, err := _AccountAliases.contract.WatchLogs(opts, "Reserved", _mainRule, _aliasRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAliasesReserved)
				if err := _AccountAliases.contract.UnpackLog(event, "Reserved", log); err != nil {
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

// ParseReserved is a log parse operation binding the contract event 0xd0e1be92289836d3c01bb059be619bebe3f7818670bb0ba1b2656816330ffc6e.
//
// Solidity: event Reserved(address indexed _main, address indexed _alias)
func (_AccountAliases *AccountAliasesFilterer) ParseReserved(log types.Log) (*AccountAliasesReserved, error) {
	event := new(AccountAliasesReserved)
	if err := _AccountAliases.contract.UnpackLog(event, "Reserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
