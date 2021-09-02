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

// ValidationRewardProxyMetaData contains all meta data concerning the ValidationRewardProxy contract.
var ValidationRewardProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610dbb806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057610034565b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610a39565b610727565b6040805160e081018252600060208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152478082526100b45750610725565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906100e890600401610c1f565b60206040518083038186803b15801561010057600080fd5b505afa158015610114573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101389190610a55565b60405162498bff60e81b81529091506001600160a01b0382169063498bff009061016490600401610bdd565b60206040518083038186803b15801561017c57600080fd5b505afa158015610190573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b49190610b6d565b60405162498bff60e81b81526001600160a01b0383169063498bff00906101dd90600401610c96565b60206040518083038186803b1580156101f557600080fd5b505afa158015610209573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022d9190610b6d565b01602083015260008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061026790600401610ccb565b60206040518083038186803b15801561027f57600080fd5b505afa158015610293573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b79190610a55565b90506000816001600160a01b0316630f5b8db36040518163ffffffff1660e01b815260040160006040518083038186803b1580156102f457600080fd5b505afa158015610308573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103309190810190610a71565b905083602001518151101561034757805160208501525b60005b84602001518110156103915761038482828151811061036557fe5b60200260200101516020015186604001516107ec90919063ffffffff16565b604086015260010161034a565b5060008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906103c690600401610b99565b60206040518083038186803b1580156103de57600080fd5b505afa1580156103f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104169190610a55565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb902719061044f90600401610c61565b60206040518083038186803b15801561046757600080fd5b505afa15801561047b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049f9190610a55565b905060006104ab61084d565b905060005b876020015181101561071c57836001600160a01b031663841c90648683815181106104d757fe5b6020026020010151600001516040518263ffffffff1660e01b81526004016104ff9190610b85565b60206040518083038186803b15801561051757600080fd5b505afa15801561052b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054f9190610b6d565b8860600181815250506105a6826105a08a604001516105a08c6060015161059a8b888151811061057b57fe5b6020026020010151602001518f6000015161085d90919063ffffffff16565b9061085d565b906108b6565b608089018190521561062d57836001600160a01b031663eab136a089608001518784815181106105d257fe5b6020026020010151600001516040518363ffffffff1660e01b81526004016105fa9190610b85565b6000604051808303818588803b15801561061357600080fd5b505af1158015610627573d6000803e3d6000fd5b50505050505b610662826105a08a604001516105a06106538d60600151886108f590919063ffffffff16565b61059a8b888151811061057b57fe5b60a0890181905261067257610714565b826001600160a01b0316638ddde27c8960a0015187848151811061069257fe5b6020026020010151600001516040518363ffffffff1660e01b81526004016106ba9190610b85565b6020604051808303818588803b1580156106d357600080fd5b505af11580156106e7573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061070c9190610b4d565b151560c08901525b6001016104b0565b50505050505050505b565b600054610100900460ff16806107405750610740610937565b8061074e575060005460ff16155b6107895760405162461bcd60e51b815260040180806020018281038252602e815260200180610d37602e913960400191505060405180910390fd5b600054610100900460ff161580156107b4576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0385160217905580156107e8576000805461ff00191690555b5050565b600082820183811015610844576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b90505b92915050565b6b033b2e3c9fd0803ce800000090565b60008261086c57506000610847565b8282028284828161087957fe5b04146108445760405162461bcd60e51b8152600401808060200182810382526021815260200180610d656021913960400191505060405180910390fd5b600061084483836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b81525061093d565b600061084483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506109df565b303b1590565b600081836109c95760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561098e578181015183820152602001610976565b50505050905090810190601f1680156109bb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816109d557fe5b0495945050505050565b60008184841115610a315760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561098e578181015183820152602001610976565b505050900390565b600060208284031215610a4a578081fd5b813561084481610d1e565b600060208284031215610a66578081fd5b815161084481610d1e565b60006020808385031215610a83578182fd5b825167ffffffffffffffff80821115610a9a578384fd5b818501915085601f830112610aad578384fd5b815181811115610ab957fe5b610ac68485830201610cfa565b818152848101908486016040808502870188018b1015610ae4578889fd5b8896505b84871015610b3e5780828c031215610afe578889fd5b80518181018181108882111715610b1157fe5b82528251610b1e81610d1e565b815282890151898201528452600196909601959287019290810190610ae8565b50909998505050505050505050565b600060208284031215610b5e578081fd5b81518015158114610844578182fd5b600060208284031215610b7e578081fd5b5051919050565b6001600160a01b0391909116815260200190565b60208082526024908201527f746f6b656e65636f6e6f6d6963732e76616c69646174696f6e526577617264506040820152636f6f6c7360e01b606082015260800190565b60208082526022908201527f636f6e737469747574696f6e2e6d61784e5374616e64627956616c696461746f604082015261727360f01b606082015260800190565b60208082526022908201527f676f7665726e616e63652e636f6e737469747574696f6e2e706172616d657465604082015261727360f01b606082015260800190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b6020808252601b908201527a636f6e737469747574696f6e2e6d61784e56616c696461746f727360281b604082015260600190565b602080825260159082015274676f7665726e616e63652e76616c696461746f727360581b604082015260600190565b60405181810167ffffffffffffffff81118282101715610d1657fe5b604052919050565b6001600160a01b0381168114610d3357600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122084e255903ccc583f190c68916283a4b778d57ca5e56aa304155db5ca61c9bc2a64736f6c63430007060033",
}

// ValidationRewardProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidationRewardProxyMetaData.ABI instead.
var ValidationRewardProxyABI = ValidationRewardProxyMetaData.ABI

// ValidationRewardProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidationRewardProxyMetaData.Bin instead.
var ValidationRewardProxyBin = ValidationRewardProxyMetaData.Bin

// DeployValidationRewardProxy deploys a new Ethereum contract, binding an instance of ValidationRewardProxy to it.
func DeployValidationRewardProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidationRewardProxy, error) {
	parsed, err := ValidationRewardProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidationRewardProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidationRewardProxy{ValidationRewardProxyCaller: ValidationRewardProxyCaller{contract: contract}, ValidationRewardProxyTransactor: ValidationRewardProxyTransactor{contract: contract}, ValidationRewardProxyFilterer: ValidationRewardProxyFilterer{contract: contract}}, nil
}

// ValidationRewardProxy is an auto generated Go binding around an Ethereum contract.
type ValidationRewardProxy struct {
	ValidationRewardProxyCaller     // Read-only binding to the contract
	ValidationRewardProxyTransactor // Write-only binding to the contract
	ValidationRewardProxyFilterer   // Log filterer for contract events
}

// ValidationRewardProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidationRewardProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidationRewardProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidationRewardProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidationRewardProxySession struct {
	Contract     *ValidationRewardProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ValidationRewardProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidationRewardProxyCallerSession struct {
	Contract *ValidationRewardProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ValidationRewardProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidationRewardProxyTransactorSession struct {
	Contract     *ValidationRewardProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ValidationRewardProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidationRewardProxyRaw struct {
	Contract *ValidationRewardProxy // Generic contract binding to access the raw methods on
}

// ValidationRewardProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidationRewardProxyCallerRaw struct {
	Contract *ValidationRewardProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ValidationRewardProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidationRewardProxyTransactorRaw struct {
	Contract *ValidationRewardProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidationRewardProxy creates a new instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxy(address common.Address, backend bind.ContractBackend) (*ValidationRewardProxy, error) {
	contract, err := bindValidationRewardProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxy{ValidationRewardProxyCaller: ValidationRewardProxyCaller{contract: contract}, ValidationRewardProxyTransactor: ValidationRewardProxyTransactor{contract: contract}, ValidationRewardProxyFilterer: ValidationRewardProxyFilterer{contract: contract}}, nil
}

// NewValidationRewardProxyCaller creates a new read-only instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyCaller(address common.Address, caller bind.ContractCaller) (*ValidationRewardProxyCaller, error) {
	contract, err := bindValidationRewardProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyCaller{contract: contract}, nil
}

// NewValidationRewardProxyTransactor creates a new write-only instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidationRewardProxyTransactor, error) {
	contract, err := bindValidationRewardProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyTransactor{contract: contract}, nil
}

// NewValidationRewardProxyFilterer creates a new log filterer instance of ValidationRewardProxy, bound to a specific deployed contract.
func NewValidationRewardProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidationRewardProxyFilterer, error) {
	contract, err := bindValidationRewardProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardProxyFilterer{contract: contract}, nil
}

// bindValidationRewardProxy binds a generic wrapper to an already deployed contract.
func bindValidationRewardProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRewardProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardProxy *ValidationRewardProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.ValidationRewardProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardProxy *ValidationRewardProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardProxy *ValidationRewardProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardProxy *ValidationRewardProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.contract.Transact(opts, method, params...)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Allocate() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Allocate(&_ValidationRewardProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Allocate(&_ValidationRewardProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Initialize(&_ValidationRewardProxy.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Initialize(&_ValidationRewardProxy.TransactOpts, _registry)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxySession) Receive() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Receive(&_ValidationRewardProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidationRewardProxy *ValidationRewardProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidationRewardProxy.Contract.Receive(&_ValidationRewardProxy.TransactOpts)
}
