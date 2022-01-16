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

// QHolderRewardProxyMetaData contains all meta data concerning the QHolderRewardProxy contract.
var QHolderRewardProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061076d806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610041578063c4d66de81461005657600080fd5b3661003c5761003a610076565b005b600080fd5b34801561004d57600080fd5b5061003a610076565b34801561006257600080fd5b5061003a610071366004610616565b610529565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060238152602001610715602391396040518263ffffffff1660e01b81526004016100d0919061063a565b60206040518083038186803b1580156100e857600080fd5b505afa1580156100fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610120919061068f565b60008054604080518082018252601c81527f746f6b656e65636f6e6f6d6963732e73797374656d526573657276650000000060208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb902719161018c919060040161063a565b60206040518083038186803b1580156101a457600080fd5b505afa1580156101b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101dc919061068f565b6001600160a01b03166b033b2e3c9fd0803ce800000060405162498bff60e81b815260206004820152601d60248201527f676f7665726e65642e45505146492e515f726573657276655368617265000000604482015247906001600160a01b0386169063498bff009060640160206040518083038186803b15801561026057600080fd5b505afa158015610274573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029891906106ac565b6102a291906106c5565b6102ac91906106f2565b604051600081818185875af1925050503d80600081146102e8576040519150601f19603f3d011682016040523d82523d6000602084013e6102ed565b606091505b50509050806103825760405162461bcd60e51b815260206004820152605060248201527f5b5145432d3031343030305d2d4661696c656420746f207472616e736665722060448201527f746865207265736572766520736861726520746f207468652053797374656d5260648201526f32b9b2b93b329031b7b73a3930b1ba1760811b608482015260a4015b60405180910390fd5b60005460408051808201825260208082527f746f6b656e65636f6e6f6d6963732e71486f6c646572526577617264506f6f6c908201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb90271916103e79160040161063a565b60206040518083038186803b1580156103ff57600080fd5b505afa158015610413573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610437919061068f565b6001600160a01b03164760405160006040518083038185875af1925050503d8060008114610481576040519150601f19603f3d011682016040523d82523d6000602084013e610486565b606091505b505080915050806105255760405162461bcd60e51b815260206004820152605c60248201527f5b5145432d3031343030315d2d4661696c656420746f207472616e736665722060448201527f746865205120746f6b656e20686f6c6465722072657761726420746f2074686560648201527f2051486f6c646572526577617264506f6f6c20636f6e74726163742e00000000608482015260a401610379565b5050565b600054610100900460ff1680610542575060005460ff16155b6105a55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610379565b600054610100900460ff161580156105c7576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610525576000805461ff00191690555050565b6001600160a01b038116811461061357600080fd5b50565b60006020828403121561062857600080fd5b8135610633816105fe565b9392505050565b600060208083528351808285015260005b818110156106675785810183015185820160400152820161064b565b81811115610679576000604083870101525b50601f01601f1916929092016040019392505050565b6000602082840312156106a157600080fd5b8151610633816105fe565b6000602082840312156106be57600080fd5b5051919050565b60008160001904831182151516156106ed57634e487b7160e01b600052601160045260246000fd5b500290565b60008261070f57634e487b7160e01b600052601260045260246000fd5b50049056fe676f7665726e616e63652e657870657274732e45505146492e706172616d6574657273a2646970667358221220d9e91695fbcaf869925b9d447c211fa6d683a346df815df46f848b3c59d60b4264736f6c63430008090033",
}

// QHolderRewardProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use QHolderRewardProxyMetaData.ABI instead.
var QHolderRewardProxyABI = QHolderRewardProxyMetaData.ABI

// QHolderRewardProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QHolderRewardProxyMetaData.Bin instead.
var QHolderRewardProxyBin = QHolderRewardProxyMetaData.Bin

// DeployQHolderRewardProxy deploys a new Ethereum contract, binding an instance of QHolderRewardProxy to it.
func DeployQHolderRewardProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QHolderRewardProxy, error) {
	parsed, err := QHolderRewardProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QHolderRewardProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QHolderRewardProxy{QHolderRewardProxyCaller: QHolderRewardProxyCaller{contract: contract}, QHolderRewardProxyTransactor: QHolderRewardProxyTransactor{contract: contract}, QHolderRewardProxyFilterer: QHolderRewardProxyFilterer{contract: contract}}, nil
}

// QHolderRewardProxy is an auto generated Go binding around an Ethereum contract.
type QHolderRewardProxy struct {
	QHolderRewardProxyCaller     // Read-only binding to the contract
	QHolderRewardProxyTransactor // Write-only binding to the contract
	QHolderRewardProxyFilterer   // Log filterer for contract events
}

// QHolderRewardProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type QHolderRewardProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QHolderRewardProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QHolderRewardProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QHolderRewardProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QHolderRewardProxySession struct {
	Contract     *QHolderRewardProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QHolderRewardProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QHolderRewardProxyCallerSession struct {
	Contract *QHolderRewardProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// QHolderRewardProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QHolderRewardProxyTransactorSession struct {
	Contract     *QHolderRewardProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// QHolderRewardProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type QHolderRewardProxyRaw struct {
	Contract *QHolderRewardProxy // Generic contract binding to access the raw methods on
}

// QHolderRewardProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QHolderRewardProxyCallerRaw struct {
	Contract *QHolderRewardProxyCaller // Generic read-only contract binding to access the raw methods on
}

// QHolderRewardProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QHolderRewardProxyTransactorRaw struct {
	Contract *QHolderRewardProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQHolderRewardProxy creates a new instance of QHolderRewardProxy, bound to a specific deployed contract.
func NewQHolderRewardProxy(address common.Address, backend bind.ContractBackend) (*QHolderRewardProxy, error) {
	contract, err := bindQHolderRewardProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardProxy{QHolderRewardProxyCaller: QHolderRewardProxyCaller{contract: contract}, QHolderRewardProxyTransactor: QHolderRewardProxyTransactor{contract: contract}, QHolderRewardProxyFilterer: QHolderRewardProxyFilterer{contract: contract}}, nil
}

// NewQHolderRewardProxyCaller creates a new read-only instance of QHolderRewardProxy, bound to a specific deployed contract.
func NewQHolderRewardProxyCaller(address common.Address, caller bind.ContractCaller) (*QHolderRewardProxyCaller, error) {
	contract, err := bindQHolderRewardProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardProxyCaller{contract: contract}, nil
}

// NewQHolderRewardProxyTransactor creates a new write-only instance of QHolderRewardProxy, bound to a specific deployed contract.
func NewQHolderRewardProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*QHolderRewardProxyTransactor, error) {
	contract, err := bindQHolderRewardProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardProxyTransactor{contract: contract}, nil
}

// NewQHolderRewardProxyFilterer creates a new log filterer instance of QHolderRewardProxy, bound to a specific deployed contract.
func NewQHolderRewardProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*QHolderRewardProxyFilterer, error) {
	contract, err := bindQHolderRewardProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QHolderRewardProxyFilterer{contract: contract}, nil
}

// bindQHolderRewardProxy binds a generic wrapper to an already deployed contract.
func bindQHolderRewardProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QHolderRewardProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QHolderRewardProxy *QHolderRewardProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QHolderRewardProxy.Contract.QHolderRewardProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QHolderRewardProxy *QHolderRewardProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.QHolderRewardProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QHolderRewardProxy *QHolderRewardProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.QHolderRewardProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QHolderRewardProxy *QHolderRewardProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QHolderRewardProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QHolderRewardProxy *QHolderRewardProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QHolderRewardProxy *QHolderRewardProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.contract.Transact(opts, method, params...)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactor) Allocate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardProxy.contract.Transact(opts, "allocate")
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_QHolderRewardProxy *QHolderRewardProxySession) Allocate() (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Allocate(&_QHolderRewardProxy.TransactOpts)
}

// Allocate is a paid mutator transaction binding the contract method 0xabaa9916.
//
// Solidity: function allocate() returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactorSession) Allocate() (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Allocate(&_QHolderRewardProxy.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardProxy.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardProxy *QHolderRewardProxySession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Initialize(&_QHolderRewardProxy.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Initialize(&_QHolderRewardProxy.TransactOpts, _registry)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QHolderRewardProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardProxy *QHolderRewardProxySession) Receive() (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Receive(&_QHolderRewardProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_QHolderRewardProxy *QHolderRewardProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _QHolderRewardProxy.Contract.Receive(&_QHolderRewardProxy.TransactOpts)
}
