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
	Bin: "0x608060405234801561001057600080fd5b50670de0b6b3a764000060015542600255600380546001600160a01b03191633179055610652806100426000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806339c293f6146100885780634c89867f146100ae57806366425d36146100b657806374cd7594146100d95780638129fc1c146100ec57806382ab890a146100f65780638dc3311d14610109578063f2fde38b1461011c578063f7fb07b01461012f575b600080fd5b61009b61009636600461050c565b610137565b6040519081526020015b60405180910390f35b60025461009b565b6001546002546100c4919082565b604080519283526020830191909152016100a5565b61009b6100e736600461052e565b6101cb565b6100f46101f0565b005b61009b61010436600461052e565b6102c5565b61009b61011736600461052e565b610362565b6100f461012a366004610547565b610402565b60015461009b565b6003546000906001600160a01b0316331461016d5760405162461bcd60e51b815260040161016490610570565b60405180910390fd5b811580610178575082155b1561018657506001546101c5565b600154600090849061019885836105ac565b6101a291906105c4565b6101ac91906105e3565b60015490915081146101c2576001819055426002555b90505b92915050565b6000676765c793fa10079d601b1b6001546101e690846105c4565b6101c591906105e3565b600054610100900460ff1680610209575060005460ff16155b61026c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610164565b600054610100900460ff1615801561028e576000805461ffff19166101011790555b670de0b6b3a764000060015542600255600380546001600160a01b0319163317905580156102c2576000805461ff00191690555b50565b6003546000906001600160a01b031633146102f25760405162461bcd60e51b815260040161016490610570565b600254421161030357505060015490565b600254676765c793fa10079d601b1b906000906103209042610605565b905060008261033961033282886105ac565b848661044e565b60015461034691906105c4565b61035091906105e3565b42600255600181905595945050505050565b600154600090819061037f676765c793fa10079d601b1b856105c4565b61038991906105e3565b90506000676765c793fa10079d601b1b6001546103a690846105c4565b6103b091906105e3565b9050838110156103fb5760006103c78360016105ac565b9050676765c793fa10079d601b1b6001546103e290836105c4565b6103ec91906105e3565b91508482116103f9578092505b505b5092915050565b6003546001600160a01b0316331461042c5760405162461bcd60e51b815260040161016490610570565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b60008380156104ee576001841680156104695785925061046d565b8392505b50600283046002850494505b84156104e857858602868782041461049057600080fd5b818101818110156104a057600080fd5b85900496505060018516156104dd5785830283878204141587151516156104c657600080fd5b818101818110156104d657600080fd5b8590049350505b600285049450610479565b50610504565b8380156104fe5760009250610502565b8392505b505b509392505050565b6000806040838503121561051f57600080fd5b50508035926020909101359150565b60006020828403121561054057600080fd5b5035919050565b60006020828403121561055957600080fd5b81356001600160a01b03811681146101c257600080fd5b6020808252600c908201526b155b985d5d1a1bdc9a5e995960a21b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082198211156105bf576105bf610596565b500190565b60008160001904831182151516156105de576105de610596565b500290565b60008261060057634e487b7160e01b600052601260045260246000fd5b500490565b60008282101561061757610617610596565b50039056fea2646970667358221220a43ad52c89657fef109bd72542481dc67b81796bfc5eb0f45adfe36e6a2636d964736f6c63430008090033",
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
