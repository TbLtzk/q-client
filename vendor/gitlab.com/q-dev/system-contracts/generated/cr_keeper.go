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

// CompoundRateKeeperMetaData contains all meta data concerning the CompoundRateKeeper contract.
var CompoundRateKeeperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"compoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_interestRate\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_oldAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_distributable\",\"type\":\"uint256\"}],\"name\":\"updateWithDistributableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetAmount\",\"type\":\"uint256\"}],\"name\":\"normalizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_normalizedAmount\",\"type\":\"uint256\"}],\"name\":\"denormalizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610019610037565b60015542600255600380546001600160a01b03191633179055610043565b670de0b6b3a764000090565b610855806100526000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806339c293f6146100885780634c89867f146100b157806366425d36146100b957806374cd7594146100cf5780638129fc1c146100e257806382ab890a146100ec5780638dc3311d146100ff578063f2fde38b14610112578063f7fb07b014610125575b600080fd5b61009b610096366004610772565b61012d565b6040516100a891906107b9565b60405180910390f35b61009b6101c1565b6100c16101c7565b6040516100a89291906107c2565b61009b6100dd36600461075a565b6101d0565b6100ea6101f3565b005b61009b6100fa36600461075a565b6102b6565b61009b61010d36600461075a565b61034d565b6100ea610120366004610733565b6103c2565b61009b61040e565b6003546000906001600160a01b031633146101635760405162461bcd60e51b815260040161015a90610793565b60405180910390fd5b81158061016e575082155b1561017c57506001546101bb565b6001546000906101a290859061019c906101968388610414565b9061046c565b906104c5565b60015490915081146101b8576001819055426002555b90505b92915050565b60025490565b60015460025482565b60006101eb6101dd610504565b60015461019c90859061046c565b90505b919050565b600054610100900460ff168061020c575061020c610514565b8061021a575060005460ff16155b6102555760405162461bcd60e51b815260040180806020018281038252602e8152602001806107d1602e913960400191505060405180910390fd5b600054610100900460ff16158015610280576000805460ff1961ff0019909116610100171660011790555b610288610525565b60015542600255600380546001600160a01b0319163317905580156102b3576000805461ff00191690555b50565b6003546000906001600160a01b031633146102e35760405162461bcd60e51b815260040161015a90610793565b60025442116102f557506001546101ee565b60006102ff610504565b600254909150600090610313904290610531565b9050600061033b8361019c61033261032b8984610414565b8688610573565b6001549061046c565b42600255600181905595945050505050565b60008061036b60016000015461019c610364610504565b869061046c565b9050600061037a6101dd610504565b9050838110156103bb576000610391836001610414565b90506103ac61039e610504565b60015461019c90849061046c565b91508482116103b9578092505b505b5092915050565b6003546001600160a01b031633146103ec5760405162461bcd60e51b815260040161015a90610793565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b60015490565b6000828201838110156101b8576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60008261047b575060006101bb565b8282028284828161048857fe5b04146101b85760405162461bcd60e51b81526004018080602001828103825260218152602001806107ff6021913960400191505060405180910390fd5b60006101b883836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250610631565b6b033b2e3c9fd0803ce800000090565b600061051f306106d3565b15905090565b670de0b6b3a764000090565b60006101b883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506106d9565b60008380156106135760018416801561058e57859250610592565b8392505b50600283046002850494505b841561060d5785860286878204146105b557600080fd5b818101818110156105c557600080fd5b85900496505060018516156106025785830283878204141587151516156105eb57600080fd5b818101818110156105fb57600080fd5b8590049350505b60028504945061059e565b50610629565b8380156106235760009250610627565b8392505b505b509392505050565b600081836106bd5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561068257818101518382015260200161066a565b50505050905090810190601f1680156106af5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816106c957fe5b0495945050505050565b3b151590565b6000818484111561072b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561068257818101518382015260200161066a565b505050900390565b600060208284031215610744578081fd5b81356001600160a01b03811681146101b8578182fd5b60006020828403121561076b578081fd5b5035919050565b60008060408385031215610784578081fd5b50508035926020909101359150565b6020808252600c908201526b155b985d5d1a1bdc9a5e995960a21b604082015260600190565b90815260200190565b91825260208201526040019056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a2646970667358221220f8491c10e0871ef1d9fc47889d250d1654a22580bfb8a88b4ff7d9c4ed641bc664736f6c63430007060033",
}

// CompoundRateKeeperABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundRateKeeperMetaData.ABI instead.
var CompoundRateKeeperABI = CompoundRateKeeperMetaData.ABI

// CompoundRateKeeperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompoundRateKeeperMetaData.Bin instead.
var CompoundRateKeeperBin = CompoundRateKeeperMetaData.Bin

// DeployCompoundRateKeeper deploys a new Ethereum contract, binding an instance of CompoundRateKeeper to it.
func DeployCompoundRateKeeper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundRateKeeper, error) {
	parsed, err := CompoundRateKeeperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompoundRateKeeperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompoundRateKeeper{CompoundRateKeeperCaller: CompoundRateKeeperCaller{contract: contract}, CompoundRateKeeperTransactor: CompoundRateKeeperTransactor{contract: contract}, CompoundRateKeeperFilterer: CompoundRateKeeperFilterer{contract: contract}}, nil
}

// CompoundRateKeeper is an auto generated Go binding around an Ethereum contract.
type CompoundRateKeeper struct {
	CompoundRateKeeperCaller     // Read-only binding to the contract
	CompoundRateKeeperTransactor // Write-only binding to the contract
	CompoundRateKeeperFilterer   // Log filterer for contract events
}

// CompoundRateKeeperCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundRateKeeperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundRateKeeperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundRateKeeperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundRateKeeperSession struct {
	Contract     *CompoundRateKeeper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompoundRateKeeperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundRateKeeperCallerSession struct {
	Contract *CompoundRateKeeperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CompoundRateKeeperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundRateKeeperTransactorSession struct {
	Contract     *CompoundRateKeeperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CompoundRateKeeperRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRateKeeperRaw struct {
	Contract *CompoundRateKeeper // Generic contract binding to access the raw methods on
}

// CompoundRateKeeperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundRateKeeperCallerRaw struct {
	Contract *CompoundRateKeeperCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundRateKeeperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundRateKeeperTransactorRaw struct {
	Contract *CompoundRateKeeperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundRateKeeper creates a new instance of CompoundRateKeeper, bound to a specific deployed contract.
func NewCompoundRateKeeper(address common.Address, backend bind.ContractBackend) (*CompoundRateKeeper, error) {
	contract, err := bindCompoundRateKeeper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeper{CompoundRateKeeperCaller: CompoundRateKeeperCaller{contract: contract}, CompoundRateKeeperTransactor: CompoundRateKeeperTransactor{contract: contract}, CompoundRateKeeperFilterer: CompoundRateKeeperFilterer{contract: contract}}, nil
}

// NewCompoundRateKeeperCaller creates a new read-only instance of CompoundRateKeeper, bound to a specific deployed contract.
func NewCompoundRateKeeperCaller(address common.Address, caller bind.ContractCaller) (*CompoundRateKeeperCaller, error) {
	contract, err := bindCompoundRateKeeper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperCaller{contract: contract}, nil
}

// NewCompoundRateKeeperTransactor creates a new write-only instance of CompoundRateKeeper, bound to a specific deployed contract.
func NewCompoundRateKeeperTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundRateKeeperTransactor, error) {
	contract, err := bindCompoundRateKeeper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperTransactor{contract: contract}, nil
}

// NewCompoundRateKeeperFilterer creates a new log filterer instance of CompoundRateKeeper, bound to a specific deployed contract.
func NewCompoundRateKeeperFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundRateKeeperFilterer, error) {
	contract, err := bindCompoundRateKeeper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperFilterer{contract: contract}, nil
}

// bindCompoundRateKeeper binds a generic wrapper to an already deployed contract.
func bindCompoundRateKeeper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundRateKeeperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundRateKeeper *CompoundRateKeeperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundRateKeeper.Contract.CompoundRateKeeperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundRateKeeper *CompoundRateKeeperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.CompoundRateKeeperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundRateKeeper *CompoundRateKeeperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.CompoundRateKeeperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundRateKeeper *CompoundRateKeeperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundRateKeeper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundRateKeeper *CompoundRateKeeperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundRateKeeper *CompoundRateKeeperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.contract.Transact(opts, method, params...)
}

// CompoundRate is a free data retrieval call binding the contract method 0x66425d36.
//
// Solidity: function compoundRate() view returns(uint256 rate, uint256 lastUpdate)
func (_CompoundRateKeeper *CompoundRateKeeperCaller) CompoundRate(opts *bind.CallOpts) (struct {
	Rate       *big.Int
	LastUpdate *big.Int
}, error) {
	var out []interface{}
	err := _CompoundRateKeeper.contract.Call(opts, &out, "compoundRate")

	outstruct := new(struct {
		Rate       *big.Int
		LastUpdate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CompoundRate is a free data retrieval call binding the contract method 0x66425d36.
//
// Solidity: function compoundRate() view returns(uint256 rate, uint256 lastUpdate)
func (_CompoundRateKeeper *CompoundRateKeeperSession) CompoundRate() (struct {
	Rate       *big.Int
	LastUpdate *big.Int
}, error) {
	return _CompoundRateKeeper.Contract.CompoundRate(&_CompoundRateKeeper.CallOpts)
}

// CompoundRate is a free data retrieval call binding the contract method 0x66425d36.
//
// Solidity: function compoundRate() view returns(uint256 rate, uint256 lastUpdate)
func (_CompoundRateKeeper *CompoundRateKeeperCallerSession) CompoundRate() (struct {
	Rate       *big.Int
	LastUpdate *big.Int
}, error) {
	return _CompoundRateKeeper.Contract.CompoundRate(&_CompoundRateKeeper.CallOpts)
}

// DenormalizeAmount is a free data retrieval call binding the contract method 0x74cd7594.
//
// Solidity: function denormalizeAmount(uint256 _normalizedAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCaller) DenormalizeAmount(opts *bind.CallOpts, _normalizedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CompoundRateKeeper.contract.Call(opts, &out, "denormalizeAmount", _normalizedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DenormalizeAmount is a free data retrieval call binding the contract method 0x74cd7594.
//
// Solidity: function denormalizeAmount(uint256 _normalizedAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) DenormalizeAmount(_normalizedAmount *big.Int) (*big.Int, error) {
	return _CompoundRateKeeper.Contract.DenormalizeAmount(&_CompoundRateKeeper.CallOpts, _normalizedAmount)
}

// DenormalizeAmount is a free data retrieval call binding the contract method 0x74cd7594.
//
// Solidity: function denormalizeAmount(uint256 _normalizedAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCallerSession) DenormalizeAmount(_normalizedAmount *big.Int) (*big.Int, error) {
	return _CompoundRateKeeper.Contract.DenormalizeAmount(&_CompoundRateKeeper.CallOpts, _normalizedAmount)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCaller) GetCurrentRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundRateKeeper.contract.Call(opts, &out, "getCurrentRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) GetCurrentRate() (*big.Int, error) {
	return _CompoundRateKeeper.Contract.GetCurrentRate(&_CompoundRateKeeper.CallOpts)
}

// GetCurrentRate is a free data retrieval call binding the contract method 0xf7fb07b0.
//
// Solidity: function getCurrentRate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCallerSession) GetCurrentRate() (*big.Int, error) {
	return _CompoundRateKeeper.Contract.GetCurrentRate(&_CompoundRateKeeper.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCaller) GetLastUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CompoundRateKeeper.contract.Call(opts, &out, "getLastUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) GetLastUpdate() (*big.Int, error) {
	return _CompoundRateKeeper.Contract.GetLastUpdate(&_CompoundRateKeeper.CallOpts)
}

// GetLastUpdate is a free data retrieval call binding the contract method 0x4c89867f.
//
// Solidity: function getLastUpdate() view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCallerSession) GetLastUpdate() (*big.Int, error) {
	return _CompoundRateKeeper.Contract.GetLastUpdate(&_CompoundRateKeeper.CallOpts)
}

// NormalizeAmount is a free data retrieval call binding the contract method 0x8dc3311d.
//
// Solidity: function normalizeAmount(uint256 _targetAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCaller) NormalizeAmount(opts *bind.CallOpts, _targetAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CompoundRateKeeper.contract.Call(opts, &out, "normalizeAmount", _targetAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizeAmount is a free data retrieval call binding the contract method 0x8dc3311d.
//
// Solidity: function normalizeAmount(uint256 _targetAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) NormalizeAmount(_targetAmount *big.Int) (*big.Int, error) {
	return _CompoundRateKeeper.Contract.NormalizeAmount(&_CompoundRateKeeper.CallOpts, _targetAmount)
}

// NormalizeAmount is a free data retrieval call binding the contract method 0x8dc3311d.
//
// Solidity: function normalizeAmount(uint256 _targetAmount) view returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperCallerSession) NormalizeAmount(_targetAmount *big.Int) (*big.Int, error) {
	return _CompoundRateKeeper.Contract.NormalizeAmount(&_CompoundRateKeeper.CallOpts, _targetAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeper *CompoundRateKeeperTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeper.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeper *CompoundRateKeeperSession) Initialize() (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.Initialize(&_CompoundRateKeeper.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeper *CompoundRateKeeperTransactorSession) Initialize() (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.Initialize(&_CompoundRateKeeper.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_CompoundRateKeeper *CompoundRateKeeperTransactor) TransferOwnership(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeper.contract.Transact(opts, "transferOwnership", _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_CompoundRateKeeper *CompoundRateKeeperSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.TransferOwnership(&_CompoundRateKeeper.TransactOpts, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _to) returns()
func (_CompoundRateKeeper *CompoundRateKeeperTransactorSession) TransferOwnership(_to common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.TransferOwnership(&_CompoundRateKeeper.TransactOpts, _to)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _interestRate) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperTransactor) Update(opts *bind.TransactOpts, _interestRate *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.contract.Transact(opts, "update", _interestRate)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _interestRate) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) Update(_interestRate *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.Update(&_CompoundRateKeeper.TransactOpts, _interestRate)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _interestRate) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperTransactorSession) Update(_interestRate *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.Update(&_CompoundRateKeeper.TransactOpts, _interestRate)
}

// UpdateWithDistributableAmount is a paid mutator transaction binding the contract method 0x39c293f6.
//
// Solidity: function updateWithDistributableAmount(uint256 _oldAmount, uint256 _distributable) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperTransactor) UpdateWithDistributableAmount(opts *bind.TransactOpts, _oldAmount *big.Int, _distributable *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.contract.Transact(opts, "updateWithDistributableAmount", _oldAmount, _distributable)
}

// UpdateWithDistributableAmount is a paid mutator transaction binding the contract method 0x39c293f6.
//
// Solidity: function updateWithDistributableAmount(uint256 _oldAmount, uint256 _distributable) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperSession) UpdateWithDistributableAmount(_oldAmount *big.Int, _distributable *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.UpdateWithDistributableAmount(&_CompoundRateKeeper.TransactOpts, _oldAmount, _distributable)
}

// UpdateWithDistributableAmount is a paid mutator transaction binding the contract method 0x39c293f6.
//
// Solidity: function updateWithDistributableAmount(uint256 _oldAmount, uint256 _distributable) returns(uint256)
func (_CompoundRateKeeper *CompoundRateKeeperTransactorSession) UpdateWithDistributableAmount(_oldAmount *big.Int, _distributable *big.Int) (*types.Transaction, error) {
	return _CompoundRateKeeper.Contract.UpdateWithDistributableAmount(&_CompoundRateKeeper.TransactOpts, _oldAmount, _distributable)
}
