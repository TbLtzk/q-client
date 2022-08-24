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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"AliasUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"}],\"name\":\"Reserved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aliases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reverseAliases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"setAlias\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_alias\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_role\",\"type\":\"uint256[]\"}],\"name\":\"resolveBatchReverse\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_main\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_main\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_role\",\"type\":\"uint256[]\"}],\"name\":\"resolveBatch\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_alias\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_alias\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_role\",\"type\":\"uint256\"}],\"name\":\"resolveReverse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_main\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061084b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c80632425c56e1461007d5780639f0c6fd2146100cb578063a4dc3302146100e0578063abadc6f514610100578063d97b4de714610113578063e75179a414610126578063eccf800314610139578063eea840bd1461014c575b600080fd5b6100ae61008b3660046105aa565b60006020818152928152604080822090935290815220546001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100de6100d93660046105aa565b610175565b005b6100f36100ee3660046106a8565b610287565b6040516100c29190610767565b6100ae61010e3660046105aa565b610365565b6100f36101213660046106a8565b6103a7565b6100de6101343660046107b4565b61047e565b6100ae6101473660046105aa565b6104d4565b6100ae61015a3660046107b4565b6001602052600090815260409020546001600160a01b031681565b6001600160a01b038281166000908152600160205260409020541615806101b557506001600160a01b038281166000908152600160205260409020541633145b6102115760405162461bcd60e51b815260206004820152602260248201527f5b5145432d3034303030305d2d416c69617320616c7265616479206578697374604482015261399760f11b60648201526084015b60405180910390fd5b33600081815260208181526040808320858452825280832080546001600160a01b0388166001600160a01b0319918216811790925581855260019093528184208054909316851790925551849391927f2aa9378fc153640a84c4851f2e9d5b9a9c6660640cc5f6c01b3574e69ab765cf91a45050565b606061029583518351610515565b82516001600160401b038111156102ae576102ae6105d4565b6040519080825280602002602001820160405280156102d7578160200160208202803683370190505b50905060005b835181101561035e576103228482815181106102fb576102fb6107d6565b6020026020010151848381518110610315576103156107d6565b6020026020010151610365565b828281518110610334576103346107d6565b6001600160a01b039092166020928302919091019091015280610356816107ec565b9150506102dd565b5092915050565b6001600160a01b038083166000818152602081815260408083208684528252808320548516808452600190925290912054909216146103a15750815b92915050565b60606103b583518351610515565b82516001600160401b038111156103ce576103ce6105d4565b6040519080825280602002602001820160405280156103f7578160200160208202803683370190505b50905060005b835181101561035e5761044284828151811061041b5761041b6107d6565b6020026020010151848381518110610435576104356107d6565b60200260200101516104d4565b828281518110610454576104546107d6565b6001600160a01b039092166020928302919091019091015280610476816107ec565b9150506103fd565b3360008181526001602052604080822080546001600160a01b0319166001600160a01b038616908117909155905190917fd0e1be92289836d3c01bb059be619bebe3f7818670bb0ba1b2656816330ffc6e91a350565b6001600160a01b03808316600081815260016020908152604080832054851680845283835281842087855290925290912054909216146103a1575090919050565b80821461058a5760405162461bcd60e51b815260206004820152603f60248201527f5b5145432d3034303030315d2d4c656e677468206f662061646472657373657360448201527f20616e6420726f6c6520617272617973206d75737420626520657175616c2e006064820152608401610208565b5050565b80356001600160a01b03811681146105a557600080fd5b919050565b600080604083850312156105bd57600080fd5b6105c68361058e565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610612576106126105d4565b604052919050565b60006001600160401b03821115610633576106336105d4565b5060051b60200190565b600082601f83011261064e57600080fd5b8135602061066361065e8361061a565b6105ea565b82815260059290921b8401810191818101908684111561068257600080fd5b8286015b8481101561069d5780358352918301918301610686565b509695505050505050565b600080604083850312156106bb57600080fd5b82356001600160401b03808211156106d257600080fd5b818501915085601f8301126106e657600080fd5b813560206106f661065e8361061a565b82815260059290921b8401810191818101908984111561071557600080fd5b948201945b8386101561073a5761072b8661058e565b8252948201949082019061071a565b9650508601359250508082111561075057600080fd5b5061075d8582860161063d565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156107a85783516001600160a01b031683529284019291840191600101610783565b50909695505050505050565b6000602082840312156107c657600080fd5b6107cf8261058e565b9392505050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561080e57634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212200211dadbb2a5e490739b087fdbead66f2978968b183cb8fe68b5e25493f9084f64736f6c63430008090033",
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
