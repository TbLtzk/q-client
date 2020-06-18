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
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"validatorPubkey\",\"type\":\"string\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorsFuncSigs = map[string]string{
	"b5da04f5": "addValidator(string)",
	"b7ab4db5": "getValidators()",
}

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
var ValidatorsBin = "0x610140604052608060a08181527f306138346535363263396232326534333236396237646361323135636632656460c0527f386332306262663335646136376261653864356565383162333664386262623660e0527f3965336563373034623962366637353031303539666538363138343361383336610100527f62326662616236343166333636313663646437373336356231613532326435626101205281526100ae9060009060016100c1565b503480156100bb57600080fd5b5061021f565b82805482825590600052602060002090810192821561010e579160200282015b8281111561010e57825180516100fe91849160209091019061011e565b50916020019190600101906100e1565b5061011a929150610198565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061015f57805160ff191683800117855561018c565b8280016001018555821561018c579182015b8281111561018c578251825591602001919060010190610171565b5061011a9291506101be565b6101bb91905b8082111561011a5760006101b282826101d8565b5060010161019e565b90565b6101bb91905b8082111561011a57600081556001016101c4565b50805460018160011615610100020316600290046000825580601f106101fe575061021c565b601f01602090049060005260206000209081019061021c91906101be565b50565b6104728061022e6000396000f30060806040526004361061004b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663b5da04f58114610050578063b7ab4db514610072575b600080fd5b34801561005c57600080fd5b5061007061006b3660046102a8565b61009d565b005b34801561007e57600080fd5b506100876100e3565b6040516100949190610381565b60405180910390f35b600080546001810180835591805282516100de917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019060208501906101bc565b505050565b60606000805480602002602001604051908101604052809291908181526020016000905b828210156101b25760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561019e5780601f106101735761010080835404028352916020019161019e565b820191906000526020600020905b81548152906001019060200180831161018157829003601f168201915b505050505081526020019060010190610107565b5050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101fd57805160ff191683800117855561022a565b8280016001018555821561022a579182015b8281111561022a57825182559160200191906001019061020f565b5061023692915061023a565b5090565b6101b991905b808211156102365760008155600101610240565b6000601f8201831361026557600080fd5b8135610278610273826103c0565b610399565b9150808252602083016020830185838301111561029457600080fd5b61029f8382846103f2565b50505092915050565b6000602082840312156102ba57600080fd5b813567ffffffffffffffff8111156102d157600080fd5b6102dd84828501610254565b949350505050565b60006102f0826103ee565b80845260208401935083602082028501610309856103e8565b60005b8481101561034057838303885261032483835161034c565b925061032f826103e8565b60209890980197915060010161030c565b50909695505050505050565b6000610357826103ee565b80845261036b8160208601602086016103fe565b6103748161042e565b9093016020019392505050565b6020808252810161039281846102e5565b9392505050565b60405181810167ffffffffffffffff811182821017156103b857600080fd5b604052919050565b600067ffffffffffffffff8211156103d757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b82818337506000910152565b60005b83811015610419578181015183820152602001610401565b83811115610428576000848401525b50505050565b601f01601f1916905600a265627a7a72305820472f4af95520daa1e5330ed1beabc53a05f2eff7d536a3693ce16fbaec96385b6c6578706572696d656e74616cf50037"

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

// AddValidator is a free data retrieval call binding the contract method 0xb5da04f5.
//
// Solidity: function addValidator(string validatorPubkey) view returns()
func (_Validators *ValidatorsCaller) AddValidator(opts *bind.CallOpts, validatorPubkey string) error {
	var ()
	out := &[]interface{}{}
	err := _Validators.contract.Call(opts, out, "addValidator", validatorPubkey)
	return err
}

// AddValidator is a free data retrieval call binding the contract method 0xb5da04f5.
//
// Solidity: function addValidator(string validatorPubkey) view returns()
func (_Validators *ValidatorsSession) AddValidator(validatorPubkey string) error {
	return _Validators.Contract.AddValidator(&_Validators.CallOpts, validatorPubkey)
}

// AddValidator is a free data retrieval call binding the contract method 0xb5da04f5.
//
// Solidity: function addValidator(string validatorPubkey) view returns()
func (_Validators *ValidatorsCallerSession) AddValidator(validatorPubkey string) error {
	return _Validators.Contract.AddValidator(&_Validators.CallOpts, validatorPubkey)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(string[])
func (_Validators *ValidatorsCaller) GetValidators(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(string[])
func (_Validators *ValidatorsSession) GetValidators() ([]string, error) {
	return _Validators.Contract.GetValidators(&_Validators.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(string[])
func (_Validators *ValidatorsCallerSession) GetValidators() ([]string, error) {
	return _Validators.Contract.GetValidators(&_Validators.CallOpts)
}
