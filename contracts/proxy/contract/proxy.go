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

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorsFuncSigs = map[string]string{
	"4d238c8e": "addValidator(address)",
	"3d620628": "getValidatorsList()",
}

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// GetValidatorsList is a free data retrieval call binding the contract method 0x3d620628.
//
// Solidity: function getValidatorsList() view returns(address[])
func (_Validators *ValidatorsCaller) GetValidatorsList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidatorsList")
	return *ret0, err
}

// GetValidatorsList is a free data retrieval call binding the contract method 0x3d620628.
//
// Solidity: function getValidatorsList() view returns(address[])
func (_Validators *ValidatorsSession) GetValidatorsList() ([]common.Address, error) {
	return _Validators.Contract.GetValidatorsList(&_Validators.CallOpts)
}

// GetValidatorsList is a free data retrieval call binding the contract method 0x3d620628.
//
// Solidity: function getValidatorsList() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetValidatorsList() ([]common.Address, error) {
	return _Validators.Contract.GetValidatorsList(&_Validators.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address ) returns()
func (_Validators *ValidatorsTransactor) AddValidator(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addValidator", arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address ) returns()
func (_Validators *ValidatorsSession) AddValidator(arg0 common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, arg0)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address ) returns()
func (_Validators *ValidatorsTransactorSession) AddValidator(arg0 common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, arg0)
}

// ValidatorsProxyABI is the input ABI used to generate the binding from.
const ValidatorsProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"}],\"name\":\"changeDelegateContract\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorsProxyFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorsProxyFuncSigs = map[string]string{
	"4d238c8e": "addValidator(address)",
	"3b44eaa8": "changeDelegateContract(address)",
	"e35c0f7d": "getValidatorList()",
	"8da5cb5b": "owner()",
}

// ValidatorsProxyBin is the compiled bytecode used for deploying new contracts.
var ValidatorsProxyBin = "0x608060405234801561001057600080fd5b506040516103cc3803806103cc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b03191633179055610059816001600160e01b0361007f16565b50600180546001600160a01b0319166001600160a01b0392909216919091179055610085565b3b151590565b610338806100946000396000f3fe60806040526004361061003f5760003560e01c80633b44eaa8146100445780634d238c8e1461006c5780638da5cb5b14610092578063e35c0f7d146100c3575b600080fd5b61006a6004803603602081101561005a57600080fd5b50356001600160a01b0316610128565b005b61006a6004803603602081101561008257600080fd5b50356001600160a01b0316610161565b34801561009e57600080fd5b506100a76101e1565b604080516001600160a01b039092168252519081900360200190f35b3480156100cf57600080fd5b506100d86101f0565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101145781810151838201526020016100fc565b505050509050019250505060405180910390f35b6000546001600160a01b0316331461013f57600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316331461017857600080fd5b60015460408051632691c64760e11b81526001600160a01b03848116600483015291519190921691634d238c8e91602480830192600092919082900301818387803b1580156101c657600080fd5b505af11580156101da573d6000803e3d6000fd5b5050505050565b6000546001600160a01b031690565b600154604080516307ac40c560e31b815290516060926001600160a01b031691633d620628916004808301926000929190829003018186803b15801561023557600080fd5b505afa158015610249573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561027257600080fd5b810190808051604051939291908464010000000082111561029257600080fd5b9083019060208201858111156102a757600080fd5b82518660208202830111640100000000821117156102c457600080fd5b82525081516020918201928201910280838360005b838110156102f15781810151838201526020016102d9565b5050505090500160405250505090509056fea265627a7a723158204483e0e0b3c29356b459d92b5ae71c99d6688642a351dd67217e66e427afa1de64736f6c63430005110032"

// DeployValidatorsProxy deploys a new Ethereum contract, binding an instance of ValidatorsProxy to it.
func DeployValidatorsProxy(auth *bind.TransactOpts, backend bind.ContractBackend, contr common.Address) (common.Address, *types.Transaction, *ValidatorsProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorsProxyBin), backend, contr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorsProxy{ValidatorsProxyCaller: ValidatorsProxyCaller{contract: contract}, ValidatorsProxyTransactor: ValidatorsProxyTransactor{contract: contract}, ValidatorsProxyFilterer: ValidatorsProxyFilterer{contract: contract}}, nil
}

// ValidatorsProxy is an auto generated Go binding around an Ethereum contract.
type ValidatorsProxy struct {
	ValidatorsProxyCaller     // Read-only binding to the contract
	ValidatorsProxyTransactor // Write-only binding to the contract
	ValidatorsProxyFilterer   // Log filterer for contract events
}

// ValidatorsProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsProxySession struct {
	Contract     *ValidatorsProxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsProxyCallerSession struct {
	Contract *ValidatorsProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ValidatorsProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsProxyTransactorSession struct {
	Contract     *ValidatorsProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ValidatorsProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsProxyRaw struct {
	Contract *ValidatorsProxy // Generic contract binding to access the raw methods on
}

// ValidatorsProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsProxyCallerRaw struct {
	Contract *ValidatorsProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsProxyTransactorRaw struct {
	Contract *ValidatorsProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorsProxy creates a new instance of ValidatorsProxy, bound to a specific deployed contract.
func NewValidatorsProxy(address common.Address, backend bind.ContractBackend) (*ValidatorsProxy, error) {
	contract, err := bindValidatorsProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorsProxy{ValidatorsProxyCaller: ValidatorsProxyCaller{contract: contract}, ValidatorsProxyTransactor: ValidatorsProxyTransactor{contract: contract}, ValidatorsProxyFilterer: ValidatorsProxyFilterer{contract: contract}}, nil
}

// NewValidatorsProxyCaller creates a new read-only instance of ValidatorsProxy, bound to a specific deployed contract.
func NewValidatorsProxyCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsProxyCaller, error) {
	contract, err := bindValidatorsProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsProxyCaller{contract: contract}, nil
}

// NewValidatorsProxyTransactor creates a new write-only instance of ValidatorsProxy, bound to a specific deployed contract.
func NewValidatorsProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsProxyTransactor, error) {
	contract, err := bindValidatorsProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsProxyTransactor{contract: contract}, nil
}

// NewValidatorsProxyFilterer creates a new log filterer instance of ValidatorsProxy, bound to a specific deployed contract.
func NewValidatorsProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsProxyFilterer, error) {
	contract, err := bindValidatorsProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsProxyFilterer{contract: contract}, nil
}

// bindValidatorsProxy binds a generic wrapper to an already deployed contract.
func bindValidatorsProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorsProxy *ValidatorsProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorsProxy.Contract.ValidatorsProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorsProxy *ValidatorsProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.ValidatorsProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorsProxy *ValidatorsProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.ValidatorsProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorsProxy *ValidatorsProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorsProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorsProxy *ValidatorsProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorsProxy *ValidatorsProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.contract.Transact(opts, method, params...)
}

// GetValidatorList is a free data retrieval call binding the contract method 0xe35c0f7d.
//
// Solidity: function getValidatorList() view returns(address[])
func (_ValidatorsProxy *ValidatorsProxyCaller) GetValidatorList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ValidatorsProxy.contract.Call(opts, out, "getValidatorList")
	return *ret0, err
}

// GetValidatorList is a free data retrieval call binding the contract method 0xe35c0f7d.
//
// Solidity: function getValidatorList() view returns(address[])
func (_ValidatorsProxy *ValidatorsProxySession) GetValidatorList() ([]common.Address, error) {
	return _ValidatorsProxy.Contract.GetValidatorList(&_ValidatorsProxy.CallOpts)
}

// GetValidatorList is a free data retrieval call binding the contract method 0xe35c0f7d.
//
// Solidity: function getValidatorList() view returns(address[])
func (_ValidatorsProxy *ValidatorsProxyCallerSession) GetValidatorList() ([]common.Address, error) {
	return _ValidatorsProxy.Contract.GetValidatorList(&_ValidatorsProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorsProxy *ValidatorsProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorsProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorsProxy *ValidatorsProxySession) Owner() (common.Address, error) {
	return _ValidatorsProxy.Contract.Owner(&_ValidatorsProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorsProxy *ValidatorsProxyCallerSession) Owner() (common.Address, error) {
	return _ValidatorsProxy.Contract.Owner(&_ValidatorsProxy.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) payable returns()
func (_ValidatorsProxy *ValidatorsProxyTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.contract.Transact(opts, "addValidator", validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) payable returns()
func (_ValidatorsProxy *ValidatorsProxySession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.AddValidator(&_ValidatorsProxy.TransactOpts, validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) payable returns()
func (_ValidatorsProxy *ValidatorsProxyTransactorSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.AddValidator(&_ValidatorsProxy.TransactOpts, validator)
}

// ChangeDelegateContract is a paid mutator transaction binding the contract method 0x3b44eaa8.
//
// Solidity: function changeDelegateContract(address contr) payable returns()
func (_ValidatorsProxy *ValidatorsProxyTransactor) ChangeDelegateContract(opts *bind.TransactOpts, contr common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.contract.Transact(opts, "changeDelegateContract", contr)
}

// ChangeDelegateContract is a paid mutator transaction binding the contract method 0x3b44eaa8.
//
// Solidity: function changeDelegateContract(address contr) payable returns()
func (_ValidatorsProxy *ValidatorsProxySession) ChangeDelegateContract(contr common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.ChangeDelegateContract(&_ValidatorsProxy.TransactOpts, contr)
}

// ChangeDelegateContract is a paid mutator transaction binding the contract method 0x3b44eaa8.
//
// Solidity: function changeDelegateContract(address contr) payable returns()
func (_ValidatorsProxy *ValidatorsProxyTransactorSession) ChangeDelegateContract(contr common.Address) (*types.Transaction, error) {
	return _ValidatorsProxy.Contract.ChangeDelegateContract(&_ValidatorsProxy.TransactOpts, contr)
}
