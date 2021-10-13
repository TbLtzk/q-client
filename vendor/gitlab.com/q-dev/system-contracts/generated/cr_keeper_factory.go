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

// CompoundRateKeeperFactoryMetaData contains all meta data concerning the CompoundRateKeeperFactory contract.
var CompoundRateKeeperFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_impl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061066a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063c4d66de81461003b578063efc81a8c14610063575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b0316610087565b005b61006b61014c565b604080516001600160a01b039092168252519081900360200190f35b600054610100900460ff16806100a057506100a0610265565b806100ae575060005460ff16155b6100e95760405162461bcd60e51b815260040180806020018281038252602e815260200180610607602e913960400191505060405180910390fd5b600054610100900460ff16158015610114576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610148576000805461ff00191690555b5050565b6000805460405182916201000090046001600160a01b03169061016e9061027c565b6001600160a01b03909116815260406020820181905260008183018190529051918290036060019190f0801580156101aa573d6000803e3d6000fd5b509050806001600160a01b0316638129fc1c6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156101e857600080fd5b505af11580156101fc573d6000803e3d6000fd5b50506040805163f2fde38b60e01b815233600482015290516001600160a01b038516935063f2fde38b9250602480830192600092919082900301818387803b15801561024757600080fd5b505af115801561025b573d6000803e3d6000fd5b5092935050505090565b600061027030610276565b15905090565b3b151590565b61037d8061028a8339019056fe608060405260405161037d38038061037d8339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b50604052506100e3915050565b6100ec826101ab565b8051156101a4576000826001600160a01b0316826040518082805190602001908083835b6020831061012f5780518252601f199092019160209182019101610110565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461018f576040519150601f19603f3d011682016040523d82523d6000602084013e610194565b606091505b50509050806101a257600080fd5b505b5050610259565b6101be8161021d60201b6100271760201c565b6101f95760405162461bcd60e51b81526004018080602001828103825260368152602001806103476036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061025157508115155b949350505050565b60e0806102676000396000f3fe608060405236601057600e6013565b005b600e5b60196025565b602560216062565b6087565b565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590605a57508115155b949350505050565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801560a5573d6000f35b3d6000fdfea2646970667358221220600916e0ab8f6ded5d6ef5f7f8cdcbbeb03a64f2d4aecb8227e4873d863bf17064736f6c634300070600335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220cca2a90cff3986c0170b4d7652c1f95554f3ae3be529fd2668ea335a9c642e8864736f6c63430007060033",
}

// CompoundRateKeeperFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundRateKeeperFactoryMetaData.ABI instead.
var CompoundRateKeeperFactoryABI = CompoundRateKeeperFactoryMetaData.ABI

// CompoundRateKeeperFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CompoundRateKeeperFactoryMetaData.Bin instead.
var CompoundRateKeeperFactoryBin = CompoundRateKeeperFactoryMetaData.Bin

// DeployCompoundRateKeeperFactory deploys a new Ethereum contract, binding an instance of CompoundRateKeeperFactory to it.
func DeployCompoundRateKeeperFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundRateKeeperFactory, error) {
	parsed, err := CompoundRateKeeperFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CompoundRateKeeperFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompoundRateKeeperFactory{CompoundRateKeeperFactoryCaller: CompoundRateKeeperFactoryCaller{contract: contract}, CompoundRateKeeperFactoryTransactor: CompoundRateKeeperFactoryTransactor{contract: contract}, CompoundRateKeeperFactoryFilterer: CompoundRateKeeperFactoryFilterer{contract: contract}}, nil
}

// CompoundRateKeeperFactory is an auto generated Go binding around an Ethereum contract.
type CompoundRateKeeperFactory struct {
	CompoundRateKeeperFactoryCaller     // Read-only binding to the contract
	CompoundRateKeeperFactoryTransactor // Write-only binding to the contract
	CompoundRateKeeperFactoryFilterer   // Log filterer for contract events
}

// CompoundRateKeeperFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundRateKeeperFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundRateKeeperFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundRateKeeperFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundRateKeeperFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundRateKeeperFactorySession struct {
	Contract     *CompoundRateKeeperFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CompoundRateKeeperFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundRateKeeperFactoryCallerSession struct {
	Contract *CompoundRateKeeperFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CompoundRateKeeperFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundRateKeeperFactoryTransactorSession struct {
	Contract     *CompoundRateKeeperFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CompoundRateKeeperFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRateKeeperFactoryRaw struct {
	Contract *CompoundRateKeeperFactory // Generic contract binding to access the raw methods on
}

// CompoundRateKeeperFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundRateKeeperFactoryCallerRaw struct {
	Contract *CompoundRateKeeperFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundRateKeeperFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundRateKeeperFactoryTransactorRaw struct {
	Contract *CompoundRateKeeperFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundRateKeeperFactory creates a new instance of CompoundRateKeeperFactory, bound to a specific deployed contract.
func NewCompoundRateKeeperFactory(address common.Address, backend bind.ContractBackend) (*CompoundRateKeeperFactory, error) {
	contract, err := bindCompoundRateKeeperFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperFactory{CompoundRateKeeperFactoryCaller: CompoundRateKeeperFactoryCaller{contract: contract}, CompoundRateKeeperFactoryTransactor: CompoundRateKeeperFactoryTransactor{contract: contract}, CompoundRateKeeperFactoryFilterer: CompoundRateKeeperFactoryFilterer{contract: contract}}, nil
}

// NewCompoundRateKeeperFactoryCaller creates a new read-only instance of CompoundRateKeeperFactory, bound to a specific deployed contract.
func NewCompoundRateKeeperFactoryCaller(address common.Address, caller bind.ContractCaller) (*CompoundRateKeeperFactoryCaller, error) {
	contract, err := bindCompoundRateKeeperFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperFactoryCaller{contract: contract}, nil
}

// NewCompoundRateKeeperFactoryTransactor creates a new write-only instance of CompoundRateKeeperFactory, bound to a specific deployed contract.
func NewCompoundRateKeeperFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundRateKeeperFactoryTransactor, error) {
	contract, err := bindCompoundRateKeeperFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperFactoryTransactor{contract: contract}, nil
}

// NewCompoundRateKeeperFactoryFilterer creates a new log filterer instance of CompoundRateKeeperFactory, bound to a specific deployed contract.
func NewCompoundRateKeeperFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundRateKeeperFactoryFilterer, error) {
	contract, err := bindCompoundRateKeeperFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundRateKeeperFactoryFilterer{contract: contract}, nil
}

// bindCompoundRateKeeperFactory binds a generic wrapper to an already deployed contract.
func bindCompoundRateKeeperFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundRateKeeperFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundRateKeeperFactory.Contract.CompoundRateKeeperFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.CompoundRateKeeperFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.CompoundRateKeeperFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompoundRateKeeperFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address)
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactor) Create(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.contract.Transact(opts, "create")
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address)
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactorySession) Create() (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Create(&_CompoundRateKeeperFactory.TransactOpts)
}

// Create is a paid mutator transaction binding the contract method 0xefc81a8c.
//
// Solidity: function create() returns(address)
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactorSession) Create() (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Create(&_CompoundRateKeeperFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactor) Initialize(opts *bind.TransactOpts, _impl common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.contract.Transact(opts, "initialize", _impl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactorySession) Initialize(_impl common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Initialize(&_CompoundRateKeeperFactory.TransactOpts, _impl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactorSession) Initialize(_impl common.Address) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Initialize(&_CompoundRateKeeperFactory.TransactOpts, _impl)
}
