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

// ParametersABI is the input ABI used to generate the binding from.
const ParametersABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQIDHoldersTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQTokenHoldersTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRootNodesTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesCoinbaseSubsidyShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesNativeAppsFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorNodesTransactionFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQIDHoldersTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setQTokenHoldersTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setRootNodesTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesCoinbaseSubsidyShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesNativeAppsFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"}],\"name\":\"setValidatorNodesTransactionFeeShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersionHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ParametersFuncSigs maps the 4-byte function signature to its string representation.
var ParametersFuncSigs = map[string]string{
	"34ce10c4": "getDecimal()",
	"fc72b8d3": "getQIDHoldersCoinbaseSubsidyShare()",
	"4aa027b3": "getQIDHoldersNativeAppsFeeShare()",
	"cbb636f8": "getQIDHoldersTransactionFeeShare()",
	"be54279e": "getQTokenHoldersCoinbaseSubsidyShare()",
	"99a8cf6a": "getQTokenHoldersNativeAppsFeeShare()",
	"03dfb900": "getQTokenHoldersTransactionFeeShare()",
	"3e344f6d": "getRootNodesCoinbaseSubsidyShare()",
	"ed219949": "getRootNodesNativeAppsFeeShare()",
	"20527af1": "getRootNodesTransactionFeeShare()",
	"539c907a": "getValidatorNodesCoinbaseSubsidyShare()",
	"e54ede97": "getValidatorNodesNativeAppsFeeShare()",
	"834b7b88": "getValidatorNodesTransactionFeeShare()",
	"c6590930": "getVersionHash()",
	"a3ffe985": "setQIDHoldersCoinbaseSubsidyShare(uint256)",
	"319089dc": "setQIDHoldersNativeAppsFeeShare(uint256)",
	"5c23da02": "setQIDHoldersTransactionFeeShare(uint256)",
	"fe361df7": "setQTokenHoldersCoinbaseSubsidyShare(uint256)",
	"8a278645": "setQTokenHoldersNativeAppsFeeShare(uint256)",
	"81f7f4c8": "setQTokenHoldersTransactionFeeShare(uint256)",
	"8cf5b300": "setRootNodesCoinbaseSubsidyShare(uint256)",
	"f443f349": "setRootNodesNativeAppsFeeShare(uint256)",
	"f6e3a1d7": "setRootNodesTransactionFeeShare(uint256)",
	"d2407598": "setValidatorNodesCoinbaseSubsidyShare(uint256)",
	"d0795a79": "setValidatorNodesNativeAppsFeeShare(uint256)",
	"891fefbb": "setValidatorNodesTransactionFeeShare(uint256)",
	"94c38179": "setVersionHash(string)",
}

// ParametersBin is the compiled bytecode used for deploying new contracts.
var ParametersBin = "0x608060405261271060025534801561001657600080fd5b5060408051808201909152601d8082527f5120436f6e737469747574696f6e20506172616d65746572204c697374000000602090920191825261005b916000916100cc565b50604051806060016040528060408152602001610823604091398051610089916001916020909101906100cc565b50610d056003819055600481905560058190556006819055600781905560088190556009819055600a819055600b819055600c819055600d819055600e55610167565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010d57805160ff191683800117855561013a565b8280016001018555821561013a579182015b8281111561013a57825182559160200191906001019061011f565b5061014692915061014a565b5090565b61016491905b808211156101465760008155600101610150565b90565b6106ad806101766000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c806394c38179116100f9578063d240759811610097578063f443f34911610071578063f443f3491461044a578063f6e3a1d714610467578063fc72b8d314610484578063fe361df71461048c576101a9565b8063d24075981461041d578063e54ede971461043a578063ed21994914610442576101a9565b8063be54279e116100d3578063be54279e14610373578063c65909301461037b578063cbb636f8146103f8578063d0795a7914610400576101a9565b806394c38179146102a857806399a8cf6a1461034e578063a3ffe98514610356576101a9565b8063539c907a11610166578063834b7b8811610140578063834b7b8814610249578063891fefbb146102515780638a2786451461026e5780638cf5b3001461028b576101a9565b8063539c907a146102075780635c23da021461020f57806381f7f4c81461022c576101a9565b806303dfb900146101ae57806320527af1146101c8578063319089dc146101d057806334ce10c4146101ef5780633e344f6d146101f75780634aa027b3146101ff575b600080fd5b6101b66104a9565b60408051918252519081900360200190f35b6101b66104b0565b6101ed600480360360208110156101e657600080fd5b50356104b6565b005b6101b66104bb565b6101b66104c1565b6101b66104c7565b6101b66104cd565b6101ed6004803603602081101561022557600080fd5b50356104d3565b6101ed6004803603602081101561024257600080fd5b50356104d8565b6101b66104dd565b6101ed6004803603602081101561026757600080fd5b50356104e3565b6101ed6004803603602081101561028457600080fd5b50356104e8565b6101ed600480360360208110156102a157600080fd5b50356104ed565b6101ed600480360360208110156102be57600080fd5b8101906020810181356401000000008111156102d957600080fd5b8201836020820111156102eb57600080fd5b8035906020019184600183028401116401000000008311171561030d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104f2945050505050565b6101b6610509565b6101ed6004803603602081101561036c57600080fd5b503561050f565b6101b6610514565b61038361051a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103bd5781810151838201526020016103a5565b50505050905090810190601f1680156103ea5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b66105af565b6101ed6004803603602081101561041657600080fd5b50356105b5565b6101ed6004803603602081101561043357600080fd5b50356105ba565b6101b66105bf565b6101b66105c5565b6101ed6004803603602081101561046057600080fd5b50356105cb565b6101ed6004803603602081101561047d57600080fd5b50356105d0565b6101b66105d5565b6101ed600480360360208110156104a257600080fd5b50356105db565b6003545b90565b600c5490565b600755565b60025490565b600e5490565b60075490565b600b5490565b600655565b600355565b60095490565b600955565b600455565b600e55565b80516105059060019060208401906105e0565b5050565b60045490565b600855565b60055490565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156105a55780601f1061057a576101008083540402835291602001916105a5565b820191906000526020600020905b81548152906001019060200180831161058857829003601f168201915b5050505050905090565b60065490565b600a55565b600b55565b600a5490565b600d5490565b600d55565b600c55565b60085490565b600555565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061062157805160ff191683800117855561064e565b8280016001018555821561064e579182015b8281111561064e578251825591602001919060010190610633565b5061065a92915061065e565b5090565b6104ad91905b8082111561065a576000815560010161066456fea265627a7a723158202347d01272016499d16be6ddb18ae76d5f1f02bc7ab7d32aa1a1e9cd08984f5964736f6c6343000511003230424142344244433743353045383746353132373341314344393544393538383530383334393739374333383433324637383839373036393346343042463837"

// DeployParameters deploys a new Ethereum contract, binding an instance of Parameters to it.
func DeployParameters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Parameters, error) {
	parsed, err := abi.JSON(strings.NewReader(ParametersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ParametersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Parameters{ParametersCaller: ParametersCaller{contract: contract}, ParametersTransactor: ParametersTransactor{contract: contract}, ParametersFilterer: ParametersFilterer{contract: contract}}, nil
}

// Parameters is an auto generated Go binding around an Ethereum contract.
type Parameters struct {
	ParametersCaller     // Read-only binding to the contract
	ParametersTransactor // Write-only binding to the contract
	ParametersFilterer   // Log filterer for contract events
}

// ParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParametersSession struct {
	Contract     *Parameters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParametersCallerSession struct {
	Contract *ParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParametersTransactorSession struct {
	Contract     *ParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParametersRaw struct {
	Contract *Parameters // Generic contract binding to access the raw methods on
}

// ParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParametersCallerRaw struct {
	Contract *ParametersCaller // Generic read-only contract binding to access the raw methods on
}

// ParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParametersTransactorRaw struct {
	Contract *ParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParameters creates a new instance of Parameters, bound to a specific deployed contract.
func NewParameters(address common.Address, backend bind.ContractBackend) (*Parameters, error) {
	contract, err := bindParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Parameters{ParametersCaller: ParametersCaller{contract: contract}, ParametersTransactor: ParametersTransactor{contract: contract}, ParametersFilterer: ParametersFilterer{contract: contract}}, nil
}

// NewParametersCaller creates a new read-only instance of Parameters, bound to a specific deployed contract.
func NewParametersCaller(address common.Address, caller bind.ContractCaller) (*ParametersCaller, error) {
	contract, err := bindParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersCaller{contract: contract}, nil
}

// NewParametersTransactor creates a new write-only instance of Parameters, bound to a specific deployed contract.
func NewParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*ParametersTransactor, error) {
	contract, err := bindParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParametersTransactor{contract: contract}, nil
}

// NewParametersFilterer creates a new log filterer instance of Parameters, bound to a specific deployed contract.
func NewParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*ParametersFilterer, error) {
	contract, err := bindParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParametersFilterer{contract: contract}, nil
}

// bindParameters binds a generic wrapper to an already deployed contract.
func bindParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.ParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.ParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Parameters *ParametersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Parameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Parameters *ParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Parameters *ParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Parameters.Contract.contract.Transact(opts, method, params...)
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersCaller) GetDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getDecimal")
	return *ret0, err
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersSession) GetDecimal() (*big.Int, error) {
	return _Parameters.Contract.GetDecimal(&_Parameters.CallOpts)
}

// GetDecimal is a free data retrieval call binding the contract method 0x34ce10c4.
//
// Solidity: function getDecimal() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetDecimal() (*big.Int, error) {
	return _Parameters.Contract.GetDecimal(&_Parameters.CallOpts)
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersCoinbaseSubsidyShare")
	return *ret0, err
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQIDHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xfc72b8d3.
//
// Solidity: function getQIDHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersNativeAppsFeeShare")
	return *ret0, err
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x4aa027b3.
//
// Solidity: function getQIDHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQIDHoldersTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQIDHoldersTransactionFeeShare")
	return *ret0, err
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQIDHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQIDHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0xcbb636f8.
//
// Solidity: function getQIDHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQIDHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQIDHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersCoinbaseSubsidyShare")
	return *ret0, err
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0xbe54279e.
//
// Solidity: function getQTokenHoldersCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersNativeAppsFeeShare")
	return *ret0, err
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersNativeAppsFeeShare is a free data retrieval call binding the contract method 0x99a8cf6a.
//
// Solidity: function getQTokenHoldersNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetQTokenHoldersTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getQTokenHoldersTransactionFeeShare")
	return *ret0, err
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetQTokenHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetQTokenHoldersTransactionFeeShare is a free data retrieval call binding the contract method 0x03dfb900.
//
// Solidity: function getQTokenHoldersTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetQTokenHoldersTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetQTokenHoldersTransactionFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesCoinbaseSubsidyShare")
	return *ret0, err
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetRootNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x3e344f6d.
//
// Solidity: function getRootNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesNativeAppsFeeShare")
	return *ret0, err
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xed219949.
//
// Solidity: function getRootNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetRootNodesTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getRootNodesTransactionFeeShare")
	return *ret0, err
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetRootNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetRootNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x20527af1.
//
// Solidity: function getRootNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetRootNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetRootNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesCoinbaseSubsidyShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesCoinbaseSubsidyShare")
	return *ret0, err
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetValidatorNodesCoinbaseSubsidyShare is a free data retrieval call binding the contract method 0x539c907a.
//
// Solidity: function getValidatorNodesCoinbaseSubsidyShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesCoinbaseSubsidyShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesCoinbaseSubsidyShare(&_Parameters.CallOpts)
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesNativeAppsFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesNativeAppsFeeShare")
	return *ret0, err
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesNativeAppsFeeShare is a free data retrieval call binding the contract method 0xe54ede97.
//
// Solidity: function getValidatorNodesNativeAppsFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesNativeAppsFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesNativeAppsFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCaller) GetValidatorNodesTransactionFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getValidatorNodesTransactionFeeShare")
	return *ret0, err
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersSession) GetValidatorNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetValidatorNodesTransactionFeeShare is a free data retrieval call binding the contract method 0x834b7b88.
//
// Solidity: function getValidatorNodesTransactionFeeShare() view returns(uint256)
func (_Parameters *ParametersCallerSession) GetValidatorNodesTransactionFeeShare() (*big.Int, error) {
	return _Parameters.Contract.GetValidatorNodesTransactionFeeShare(&_Parameters.CallOpts)
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersCaller) GetVersionHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Parameters.contract.Call(opts, out, "getVersionHash")
	return *ret0, err
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersSession) GetVersionHash() (string, error) {
	return _Parameters.Contract.GetVersionHash(&_Parameters.CallOpts)
}

// GetVersionHash is a free data retrieval call binding the contract method 0xc6590930.
//
// Solidity: function getVersionHash() view returns(string)
func (_Parameters *ParametersCallerSession) GetVersionHash() (string, error) {
	return _Parameters.Contract.GetVersionHash(&_Parameters.CallOpts)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersCoinbaseSubsidyShare", feeShare)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xa3ffe985.
//
// Solidity: function setQIDHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersNativeAppsFeeShare", feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x319089dc.
//
// Solidity: function setQIDHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQIDHoldersTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQIDHoldersTransactionFeeShare", feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQIDHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQIDHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x5c23da02.
//
// Solidity: function setQIDHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQIDHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQIDHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersCoinbaseSubsidyShare", feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xfe361df7.
//
// Solidity: function setQTokenHoldersCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersNativeAppsFeeShare", feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersNativeAppsFeeShare is a paid mutator transaction binding the contract method 0x8a278645.
//
// Solidity: function setQTokenHoldersNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetQTokenHoldersTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setQTokenHoldersTransactionFeeShare", feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetQTokenHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetQTokenHoldersTransactionFeeShare is a paid mutator transaction binding the contract method 0x81f7f4c8.
//
// Solidity: function setQTokenHoldersTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetQTokenHoldersTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetQTokenHoldersTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesCoinbaseSubsidyShare", feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0x8cf5b300.
//
// Solidity: function setRootNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesNativeAppsFeeShare", feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xf443f349.
//
// Solidity: function setRootNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetRootNodesTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setRootNodesTransactionFeeShare", feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetRootNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetRootNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0xf6e3a1d7.
//
// Solidity: function setRootNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetRootNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetRootNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesCoinbaseSubsidyShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesCoinbaseSubsidyShare", feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesCoinbaseSubsidyShare is a paid mutator transaction binding the contract method 0xd2407598.
//
// Solidity: function setValidatorNodesCoinbaseSubsidyShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesCoinbaseSubsidyShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesCoinbaseSubsidyShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesNativeAppsFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesNativeAppsFeeShare", feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesNativeAppsFeeShare is a paid mutator transaction binding the contract method 0xd0795a79.
//
// Solidity: function setValidatorNodesNativeAppsFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesNativeAppsFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesNativeAppsFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactor) SetValidatorNodesTransactionFeeShare(opts *bind.TransactOpts, feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setValidatorNodesTransactionFeeShare", feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersSession) SetValidatorNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetValidatorNodesTransactionFeeShare is a paid mutator transaction binding the contract method 0x891fefbb.
//
// Solidity: function setValidatorNodesTransactionFeeShare(uint256 feeShare) returns()
func (_Parameters *ParametersTransactorSession) SetValidatorNodesTransactionFeeShare(feeShare *big.Int) (*types.Transaction, error) {
	return _Parameters.Contract.SetValidatorNodesTransactionFeeShare(&_Parameters.TransactOpts, feeShare)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersTransactor) SetVersionHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _Parameters.contract.Transact(opts, "setVersionHash", hash)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersSession) SetVersionHash(hash string) (*types.Transaction, error) {
	return _Parameters.Contract.SetVersionHash(&_Parameters.TransactOpts, hash)
}

// SetVersionHash is a paid mutator transaction binding the contract method 0x94c38179.
//
// Solidity: function setVersionHash(string hash) returns()
func (_Parameters *ParametersTransactorSession) SetVersionHash(hash string) (*types.Transaction, error) {
	return _Parameters.Contract.SetVersionHash(&_Parameters.TransactOpts, hash)
}
