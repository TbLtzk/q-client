// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"gitlab.com/q-dev/q-client/accounts/abi"
	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorsList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorsFuncSigs = map[string]string{
	"4d238c8e": "addValidator(address)",
	"3d620628": "getValidatorsList()",
}

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
var ValidatorsBin = "0x60e06040527364d4edefe8ba86d3588b213b0a053e7b910cad686080908152736a39b688d591ea00c9ea69658438794204b5cc6260a052734a14d788d86d021670ebcece1196631d6659598460c05261005c90600090600361006f565b5034801561006957600080fd5b506100fb565b8280548282559060005260206000209081019282156100c4579160200282015b828111156100c45782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019061008f565b506100d09291506100d4565b5090565b6100f891905b808211156100d0578054600160a060020a03191681556001016100da565b90565b6102828061010a6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633d62062881146100505780634d238c8e1461007b575b600080fd5b34801561005c57600080fd5b5061006561009d565b6040516100729190610214565b60405180910390f35b34801561008757600080fd5b5061009b610096366004610188565b61010c565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561010257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116100d7575b5050505050905090565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000610181823561022f565b9392505050565b60006020828403121561019a57600080fd5b60006101a68484610175565b949350505050565b6101b78161022f565b82525050565b60006101c88261022b565b8084526020840193506101da83610225565b60005b8281101561020a576101f08683516101ae565b6101f982610225565b6020969096019591506001016101dd565b5093949350505050565b6020808252810161018181846101bd565b60200190565b5190565b73ffffffffffffffffffffffffffffffffffffffff16905600a265627a7a72305820aa395d36a868c31f84352d0f9d02c290a15c02b40f925d0f431d8290c08617016c6578706572696d656e74616cf50037"

// DeployValidators deploys a new Ethereum contract, binding an instance of Validators to it.
func DeployValidators(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validators, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
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
// Solidity: function addValidator(address validatorAddress) returns()
func (_Validators *ValidatorsTransactor) AddValidator(opts *bind.TransactOpts, validatorAddress common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addValidator", validatorAddress)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validatorAddress) returns()
func (_Validators *ValidatorsSession) AddValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, validatorAddress)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validatorAddress) returns()
func (_Validators *ValidatorsTransactorSession) AddValidator(validatorAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddValidator(&_Validators.TransactOpts, validatorAddress)
}
