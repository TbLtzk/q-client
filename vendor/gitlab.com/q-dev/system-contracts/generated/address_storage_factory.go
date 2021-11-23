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

// AddressStorageFactoryMetaData contains all meta data concerning the AddressStorageFactory contract.
var AddressStorageFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_impl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061074c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806362c54dbd1461003b578063c4d66de8146100f8575b600080fd5b6100dc6004803603602081101561005157600080fd5b810190602081018135600160201b81111561006b57600080fd5b82018360208201111561007d57600080fd5b803590602001918460208302840111600160201b8311171561009e57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610120945050505050565b604080516001600160a01b039092168252519081900360200190f35b61011e6004803603602081101561010e57600080fd5b50356001600160a01b0316610282565b005b6000805460405182916201000090046001600160a01b0316906101429061035e565b6001600160a01b03909116815260406020820181905260008183018190529051918290036060019190f08015801561017e573d6000803e3d6000fd5b5060405163a224cee760e01b81526020600482018181528651602484015286519394506001600160a01b0385169363a224cee79388938392604490920191818601910280838360005b838110156101df5781810151838201526020016101c7565b5050505090500192505050600060405180830381600087803b15801561020457600080fd5b505af1158015610218573d6000803e3d6000fd5b50506040805163f2fde38b60e01b815233600482015290516001600160a01b038516935063f2fde38b9250602480830192600092919082900301818387803b15801561026357600080fd5b505af1158015610277573d6000803e3d6000fd5b509295945050505050565b600054610100900460ff168061029b575061029b610347565b806102a9575060005460ff16155b6102e45760405162461bcd60e51b815260040180806020018281038252602e8152602001806106e9602e913960400191505060405180910390fd5b600054610100900460ff1615801561030f576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610343576000805461ff00191690555b5050565b600061035230610358565b15905090565b3b151590565b61037d8061036c8339019056fe608060405260405161037d38038061037d8339818101604052604081101561002657600080fd5b81516020830180516040519294929383019291908464010000000082111561004d57600080fd5b90830190602082018581111561006257600080fd5b825164010000000081118282018810171561007c57600080fd5b82525081516020918201929091019080838360005b838110156100a9578181015183820152602001610091565b50505050905090810190601f1680156100d65780820380516001836020036101000a031916815260200191505b50604052506100e3915050565b6100ec826101ab565b8051156101a4576000826001600160a01b0316826040518082805190602001908083835b6020831061012f5780518252601f199092019160209182019101610110565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811461018f576040519150601f19603f3d011682016040523d82523d6000602084013e610194565b606091505b50509050806101a257600080fd5b505b5050610259565b6101be8161021d60201b6100271760201c565b6101f95760405162461bcd60e51b81526004018080602001828103825260368152602001806103476036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081811480159061025157508115155b949350505050565b60e0806102676000396000f3fe608060405236601057600e6013565b005b600e5b60196025565b602560216062565b6087565b565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590605a57508115155b949350505050565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801560a5573d6000f35b3d6000fdfea2646970667358221220600916e0ab8f6ded5d6ef5f7f8cdcbbeb03a64f2d4aecb8227e4873d863bf17064736f6c634300070600335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220f7e428717458dcfe530cccd332a979e90689696f8aefcfa80a8aa435faa2e1b264736f6c63430007060033",
}

// AddressStorageFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageFactoryMetaData.ABI instead.
var AddressStorageFactoryABI = AddressStorageFactoryMetaData.ABI

// AddressStorageFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageFactoryMetaData.Bin instead.
var AddressStorageFactoryBin = AddressStorageFactoryMetaData.Bin

// DeployAddressStorageFactory deploys a new Ethereum contract, binding an instance of AddressStorageFactory to it.
func DeployAddressStorageFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorageFactory, error) {
	parsed, err := AddressStorageFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorageFactory{AddressStorageFactoryCaller: AddressStorageFactoryCaller{contract: contract}, AddressStorageFactoryTransactor: AddressStorageFactoryTransactor{contract: contract}, AddressStorageFactoryFilterer: AddressStorageFactoryFilterer{contract: contract}}, nil
}

// AddressStorageFactory is an auto generated Go binding around an Ethereum contract.
type AddressStorageFactory struct {
	AddressStorageFactoryCaller     // Read-only binding to the contract
	AddressStorageFactoryTransactor // Write-only binding to the contract
	AddressStorageFactoryFilterer   // Log filterer for contract events
}

// AddressStorageFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageFactorySession struct {
	Contract     *AddressStorageFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AddressStorageFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageFactoryCallerSession struct {
	Contract *AddressStorageFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AddressStorageFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageFactoryTransactorSession struct {
	Contract     *AddressStorageFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AddressStorageFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageFactoryRaw struct {
	Contract *AddressStorageFactory // Generic contract binding to access the raw methods on
}

// AddressStorageFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageFactoryCallerRaw struct {
	Contract *AddressStorageFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageFactoryTransactorRaw struct {
	Contract *AddressStorageFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorageFactory creates a new instance of AddressStorageFactory, bound to a specific deployed contract.
func NewAddressStorageFactory(address common.Address, backend bind.ContractBackend) (*AddressStorageFactory, error) {
	contract, err := bindAddressStorageFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFactory{AddressStorageFactoryCaller: AddressStorageFactoryCaller{contract: contract}, AddressStorageFactoryTransactor: AddressStorageFactoryTransactor{contract: contract}, AddressStorageFactoryFilterer: AddressStorageFactoryFilterer{contract: contract}}, nil
}

// NewAddressStorageFactoryCaller creates a new read-only instance of AddressStorageFactory, bound to a specific deployed contract.
func NewAddressStorageFactoryCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageFactoryCaller, error) {
	contract, err := bindAddressStorageFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFactoryCaller{contract: contract}, nil
}

// NewAddressStorageFactoryTransactor creates a new write-only instance of AddressStorageFactory, bound to a specific deployed contract.
func NewAddressStorageFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageFactoryTransactor, error) {
	contract, err := bindAddressStorageFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFactoryTransactor{contract: contract}, nil
}

// NewAddressStorageFactoryFilterer creates a new log filterer instance of AddressStorageFactory, bound to a specific deployed contract.
func NewAddressStorageFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageFactoryFilterer, error) {
	contract, err := bindAddressStorageFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFactoryFilterer{contract: contract}, nil
}

// bindAddressStorageFactory binds a generic wrapper to an already deployed contract.
func bindAddressStorageFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressStorageFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageFactory *AddressStorageFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageFactory.Contract.AddressStorageFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageFactory *AddressStorageFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.AddressStorageFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageFactory *AddressStorageFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.AddressStorageFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageFactory *AddressStorageFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageFactory *AddressStorageFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageFactory *AddressStorageFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0x62c54dbd.
//
// Solidity: function create(address[] _addrList) returns(address)
func (_AddressStorageFactory *AddressStorageFactoryTransactor) Create(opts *bind.TransactOpts, _addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.contract.Transact(opts, "create", _addrList)
}

// Create is a paid mutator transaction binding the contract method 0x62c54dbd.
//
// Solidity: function create(address[] _addrList) returns(address)
func (_AddressStorageFactory *AddressStorageFactorySession) Create(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.Create(&_AddressStorageFactory.TransactOpts, _addrList)
}

// Create is a paid mutator transaction binding the contract method 0x62c54dbd.
//
// Solidity: function create(address[] _addrList) returns(address)
func (_AddressStorageFactory *AddressStorageFactoryTransactorSession) Create(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.Create(&_AddressStorageFactory.TransactOpts, _addrList)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_AddressStorageFactory *AddressStorageFactoryTransactor) Initialize(opts *bind.TransactOpts, _impl common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.contract.Transact(opts, "initialize", _impl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_AddressStorageFactory *AddressStorageFactorySession) Initialize(_impl common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.Initialize(&_AddressStorageFactory.TransactOpts, _impl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _impl) returns()
func (_AddressStorageFactory *AddressStorageFactoryTransactorSession) Initialize(_impl common.Address) (*types.Transaction, error) {
	return _AddressStorageFactory.Contract.Initialize(&_AddressStorageFactory.TransactOpts, _impl)
}
