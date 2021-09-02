// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CompoundRateKeeperFactoryABI is the input ABI used to generate the binding from.
const CompoundRateKeeperFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CompoundRateKeeperFactoryBin is the compiled bytecode used for deploying new contracts.
var CompoundRateKeeperFactoryBin = "0x608060405234801561001057600080fd5b506108a3806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063efc81a8c14610030575b600080fd5b610038610054565b604080516001600160a01b039092168252519081900360200190f35b600080604051610063906100ef565b604051809103906000f08015801561007f573d6000803e3d6000fd5b509050806001600160a01b031663f2fde38b336040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b1580156100d157600080fd5b505af11580156100e5573d6000803e3d6000fd5b5092935050505090565b610771806100fd8339019056fe608060405234801561001057600080fd5b50610019610037565b60005542600155600280546001600160a01b03191633179055610043565b670de0b6b3a764000090565b61071f806100526000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806339c293f61461007d5780634c89867f146100a657806366425d36146100ae57806374cd7594146100c457806382ab890a146100d75780638dc3311d146100ea578063f2fde38b146100fd578063f7fb07b014610112575b600080fd5b61009061008b36600461066a565b61011a565b60405161009d91906106b1565b60405180910390f35b61009061019c565b6100b66101a2565b60405161009d9291906106ba565b6100906100d2366004610652565b6101ab565b6100906100e5366004610652565b6101ce565b6100906100f8366004610652565b610271565b61011061010b36600461062b565b6102d6565b005b610090610322565b6002546000906001600160a01b031633146101505760405162461bcd60e51b81526004016101479061068b565b60405180910390fd5b81158061015b575082155b156101695750600054610196565b6000805461018d908590610187906101818388610328565b90610387565b906103e0565b60008190559150505b92915050565b60015490565b60005460015482565b60006101c66101b861041f565b600054610187908590610387565b90505b919050565b6002546000906001600160a01b031633146101fb5760405162461bcd60e51b81526004016101479061068b565b600154421161020d57506000546101c9565b600061021761041f565b60015490915060009061022b90429061042f565b905060006102538361018761024a6102438984610328565b8688610471565b60005490610387565b6000549091508114610269576000819055426001555b949350505050565b60008061028e600080015461018761028761041f565b8690610387565b90506000805b6102ad61029f61041f565b600054610187908690610387565b90508481116102ba578291505b6102c5836001610328565b925084811061029457509392505050565b6002546001600160a01b031633146103005760405162461bcd60e51b81526004016101479061068b565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60005490565b600082820183811015610380576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b60008261039657506000610196565b828202828482816103a357fe5b04146103805760405162461bcd60e51b81526004018080602001828103825260218152602001806106c96021913960400191505060405180910390fd5b600061038083836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b81525061052f565b6b033b2e3c9fd0803ce800000090565b600061038083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506105d1565b60008380156105115760018416801561048c57859250610490565b8392505b50600283046002850494505b841561050b5785860286878204146104b357600080fd5b818101818110156104c357600080fd5b85900496505060018516156105005785830283878204141587151516156104e957600080fd5b818101818110156104f957600080fd5b8590049350505b60028504945061049c565b50610527565b8380156105215760009250610525565b8392505b505b509392505050565b600081836105bb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610580578181015183820152602001610568565b50505050905090810190601f1680156105ad5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816105c757fe5b0495945050505050565b600081848411156106235760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610580578181015183820152602001610568565b505050900390565b60006020828403121561063c578081fd5b81356001600160a01b0381168114610380578182fd5b600060208284031215610663578081fd5b5035919050565b6000806040838503121561067c578081fd5b50508035926020909101359150565b6020808252600c908201526b155b985d5d1a1bdc9a5e995960a21b604082015260600190565b90815260200190565b91825260208201526040019056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212200983c2c2c4b846af6d5b77e7cd6482c3d38a79f65aaef20e06d23b654994005864736f6c63430007060033a264697066735822122005ac5d21bc619071819f8bc92b6df5fcc535a704f6f4a0fd5e7622a757b466f064736f6c63430007060033"

// DeployCompoundRateKeeperFactory deploys a new Ethereum contract, binding an instance of CompoundRateKeeperFactory to it.
func DeployCompoundRateKeeperFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundRateKeeperFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundRateKeeperFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompoundRateKeeperFactoryBin), backend)
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
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
