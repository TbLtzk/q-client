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
	Bin: "0x608060405234801561001057600080fd5b50610da5806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610039578063c4d66de81461005057610034565b3661003457005b600080fd5b34801561004557600080fd5b5061004e610070565b005b34801561005c57600080fd5b5061004e61006b366004610a95565b61078c565b610078610a5f565b47808252610086575061078a565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001610d4e602291396040518263ffffffff1660e01b81526004016100e09190610bf5565b60206040518083038186803b1580156100f857600080fd5b505afa15801561010c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101309190610ab1565b60405162498bff60e81b81529091506001600160a01b0382169063498bff009061015c90600401610c48565b60206040518083038186803b15801561017457600080fd5b505afa158015610188573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ac9190610bc9565b60405162498bff60e81b81526001600160a01b0383169063498bff00906101d590600401610c8a565b60206040518083038186803b1580156101ed57600080fd5b505afa158015610201573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102259190610bc9565b01602083810191909152600080546040805180820182526015815274676f7665726e616e63652e76616c696461746f727360581b9481019490945251633fb9027160e01b81529192620100009091046001600160a01b031691633fb902719161029091600401610bf5565b60206040518083038186803b1580156102a857600080fd5b505afa1580156102bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e09190610ab1565b90506000816001600160a01b0316630f5b8db36040518163ffffffff1660e01b815260040160006040518083038186803b15801561031d57600080fd5b505afa158015610331573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103599190810190610acd565b905083602001518151101561037057805160208501525b60005b84602001518110156103ba576103ad82828151811061038e57fe5b602002602001015160200151866040015161085190919063ffffffff16565b6040860152600101610373565b5060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060248152602001610cfc602491396040518263ffffffff1660e01b81526004016104159190610bf5565b60206040518083038186803b15801561042d57600080fd5b505afa158015610441573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104659190610ab1565b60008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb90271916104cf9190600401610bf5565b60206040518083038186803b1580156104e757600080fd5b505afa1580156104fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051f9190610ab1565b9050600061052b6108b0565b905060005b876020015181101561078157836001600160a01b031663841c906486838151811061055757fe5b6020026020010151600001516040518263ffffffff1660e01b815260040161057f9190610be1565b60206040518083038186803b15801561059757600080fd5b505afa1580156105ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cf9190610bc9565b88606001818152505061061c8860600151836106158885815181106105f057fe5b6020026020010151602001518c604001518d600001516108c09092919063ffffffff16565b91906108c0565b60808901819052156106a357836001600160a01b031663eab136a0896080015187848151811061064857fe5b6020026020010151600001516040518363ffffffff1660e01b81526004016106709190610be1565b6000604051808303818588803b15801561068957600080fd5b505af115801561069d573d6000803e3d6000fd5b50505050505b6106cd6106bd89606001518461096f90919063ffffffff16565b836106158885815181106105f057fe5b60a089018190526106dd57610779565b826001600160a01b0316638ddde27c8960a001518784815181106106fd57fe5b6020026020010151600001516040518363ffffffff1660e01b81526004016107259190610be1565b6020604051808303818588803b15801561073e57600080fd5b505af1158015610752573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906107779190610ba9565b505b600101610530565b50505050505050505b565b600054610100900460ff16806107a557506107a56109b1565b806107b3575060005460ff16155b6107ee5760405162461bcd60e51b815260040180806020018281038252602e815260200180610d20602e913960400191505060405180910390fd5b600054610100900460ff16158015610819576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b03851602179055801561084d576000805461ff00191690555b5050565b6000828201838110156108a9576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b6b033b2e3c9fd0803ce800000090565b60008080600019858709868602925082811090839003039050806108f657600084116108eb57600080fd5b5082900490506108a9565b80841161090257600080fd5b6000848688096000868103871696879004966002600389028118808a02820302808a02820302808a02820302808a02820302808a02820302808a02909103029181900381900460010186841190950394909402919094039290920491909117919091029150509392505050565b60006108a983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506109c2565b60006109bc30610a59565b15905090565b60008184841115610a515760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a165781810151838201526020016109fe565b50505050905090810190601f168015610a435780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b3b151590565b6040518060c001604052806000815260200160008152602001600081526020016000815260200160008152602001600081525090565b600060208284031215610aa6578081fd5b81356108a981610ce3565b600060208284031215610ac2578081fd5b81516108a981610ce3565b60006020808385031215610adf578182fd5b825167ffffffffffffffff80821115610af6578384fd5b818501915085601f830112610b09578384fd5b815181811115610b1557fe5b610b228485830201610cbf565b818152848101908486016040808502870188018b1015610b40578889fd5b8896505b84871015610b9a5780828c031215610b5a578889fd5b80518181018181108882111715610b6d57fe5b82528251610b7a81610ce3565b815282890151898201528452600196909601959287019290810190610b44565b50909998505050505050505050565b600060208284031215610bba578081fd5b815180151581146108a9578182fd5b600060208284031215610bda578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6000602080835283518082850152825b81811015610c2157858101830151858201604001528201610c05565b81811115610c325783604083870101525b50601f01601f1916929092016040019392505050565b60208082526022908201527f636f6e737469747574696f6e2e6d61784e5374616e64627956616c696461746f604082015261727360f01b606082015260800190565b6020808252601b908201527a636f6e737469747574696f6e2e6d61784e56616c696461746f727360281b604082015260600190565b60405181810167ffffffffffffffff81118282101715610cdb57fe5b604052919050565b6001600160a01b0381168114610cf857600080fd5b5056fe746f6b656e65636f6e6f6d6963732e76616c69646174696f6e526577617264506f6f6c73496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a26469706673582212200e376064f6ac7d1655f7000f01694369a6dc1764ebc652b5e389d0cd65e797cc64736f6c63430007060033",
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
