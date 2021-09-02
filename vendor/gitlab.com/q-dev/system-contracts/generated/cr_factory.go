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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610997806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638129fc1c1461003b578063efc81a8c14610045575b600080fd5b610043610069565b005b61004d61010b565b604080516001600160a01b039092168252519081900360200190f35b600054610100900460ff168061008257506100826101a6565b80610090575060005460ff16155b6100cb5760405162461bcd60e51b815260040180806020018281038252602e815260200180610934602e913960400191505060405180910390fd5b600054610100900460ff161580156100f6576000805460ff1961ff0019909116610100171660011790555b8015610108576000805461ff00191690555b50565b60008060405161011a906101ac565b604051809103906000f080158015610136573d6000803e3d6000fd5b509050806001600160a01b031663f2fde38b336040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561018857600080fd5b505af115801561019c573d6000803e3d6000fd5b5092935050505090565b303b1590565b61077a806101ba8339019056fe608060405234801561001057600080fd5b50610019610037565b60005542600155600280546001600160a01b03191633179055610043565b670de0b6b3a764000090565b610728806100526000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806339c293f61461007d5780634c89867f146100a657806366425d36146100ae57806374cd7594146100c457806382ab890a146100d75780638dc3311d146100ea578063f2fde38b146100fd578063f7fb07b014610112575b600080fd5b61009061008b366004610673565b61011a565b60405161009d91906106ba565b60405180910390f35b6100906101ac565b6100b66101b2565b60405161009d9291906106c3565b6100906100d236600461065b565b6101bb565b6100906100e536600461065b565b6101de565b6100906100f836600461065b565b610281565b61011061010b366004610634565b6102e6565b005b610090610332565b6002546000906001600160a01b031633146101505760405162461bcd60e51b815260040161014790610694565b60405180910390fd5b81158061015b575082155b1561016957506000546101a6565b6000805461018d908590610187906101818388610338565b90610390565b906103e9565b60005490915081146101a3576000819055426001555b90505b92915050565b60015490565b60005460015482565b60006101d66101c8610428565b600054610187908590610390565b90505b919050565b6002546000906001600160a01b0316331461020b5760405162461bcd60e51b815260040161014790610694565b600154421161021d57506000546101d9565b6000610227610428565b60015490915060009061023b904290610438565b905060006102638361018761025a6102538984610338565b868861047a565b60005490610390565b6000549091508114610279576000819055426001555b949350505050565b60008061029e6000800154610187610297610428565b8690610390565b90506000805b6102bd6102af610428565b600054610187908690610390565b90508481116102ca578291505b6102d5836001610338565b92508481106102a457509392505050565b6002546001600160a01b031633146103105760405162461bcd60e51b815260040161014790610694565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60005490565b6000828201838110156101a3576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60008261039f575060006101a6565b828202828482816103ac57fe5b04146101a35760405162461bcd60e51b81526004018080602001828103825260218152602001806106d26021913960400191505060405180910390fd5b60006101a383836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250610538565b6b033b2e3c9fd0803ce800000090565b60006101a383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506105da565b600083801561051a5760018416801561049557859250610499565b8392505b50600283046002850494505b84156105145785860286878204146104bc57600080fd5b818101818110156104cc57600080fd5b85900496505060018516156105095785830283878204141587151516156104f257600080fd5b8181018181101561050257600080fd5b8590049350505b6002850494506104a5565b50610530565b83801561052a576000925061052e565b8392505b505b509392505050565b600081836105c45760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610589578181015183820152602001610571565b50505050905090810190601f1680156105b65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816105d057fe5b0495945050505050565b6000818484111561062c5760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610589578181015183820152602001610571565b505050900390565b600060208284031215610645578081fd5b81356001600160a01b03811681146101a3578182fd5b60006020828403121561066c578081fd5b5035919050565b60008060408385031215610685578081fd5b50508035926020909101359150565b6020808252600c908201526b155b985d5d1a1bdc9a5e995960a21b604082015260600190565b90815260200190565b91825260208201526040019056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a2646970667358221220bafc25308439bd32102ea702a90b2cfcadecc50719aeca83abc2854e7868cac764736f6c63430007060033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220663b3747db20c23474e5bb8d32bf4eb5c93b94aec5fa11f97bc9faf162e2496764736f6c63430007060033",
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactorySession) Initialize() (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Initialize(&_CompoundRateKeeperFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CompoundRateKeeperFactory *CompoundRateKeeperFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _CompoundRateKeeperFactory.Contract.Initialize(&_CompoundRateKeeperFactory.TransactOpts)
}
