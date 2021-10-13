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
	Bin: "0x608060405234801561001057600080fd5b50610898806100206000396000f3fe60806040526004361061002d5760003560e01c8063abaa991614610041578063c4d66de8146100565761003c565b3661003c5761003a610076565b005b600080fd5b34801561004d57600080fd5b5061003a610076565b34801561006257600080fd5b5061003a6100713660046105c9565b61039b565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906100aa9060040161061c565b60206040518083038186803b1580156100c257600080fd5b505afa1580156100d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fa91906105e5565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb90271906101339060040161065f565b60206040518083038186803b15801561014b57600080fd5b505afa15801561015f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061018391906105e5565b6001600160a01b031661022161019761045f565b61021b47866001600160a01b031663498bff006040518163ffffffff1660e01b81526004016101c590610696565b60206040518083038186803b1580156101dd57600080fd5b505afa1580156101f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102159190610601565b9061046f565b906104d1565b60405161022d90610619565b60006040518083038185875af1925050503d806000811461026a576040519150601f19603f3d011682016040523d82523d6000602084013e61026f565b606091505b50509050806102995760405162461bcd60e51b815260040161029090610785565b60405180910390fd5b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906102cc90600401610750565b60206040518083038186803b1580156102e457600080fd5b505afa1580156102f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031c91906105e5565b6001600160a01b03164760405161033290610619565b60006040518083038185875af1925050503d806000811461036f576040519150601f19603f3d011682016040523d82523d6000602084013e610374565b606091505b505080915050806103975760405162461bcd60e51b8152600401610290906106cd565b5050565b600054610100900460ff16806103b457506103b4610510565b806103c2575060005460ff16155b6103fd5760405162461bcd60e51b815260040180806020018281038252602e815260200180610814602e913960400191505060405180910390fd5b600054610100900460ff16158015610428576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610397576000805461ff00191690555050565b6b033b2e3c9fd0803ce800000090565b60008261047e575060006104cb565b8282028284828161048b57fe5b04146104c85760405162461bcd60e51b81526004018080602001828103825260218152602001806108426021913960400191505060405180910390fd5b90505b92915050565b60006104c883836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250610521565b600061051b306105c3565b15905090565b600081836105ad5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561057257818101518382015260200161055a565b50505050905090810190601f16801561059f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816105b957fe5b0495945050505050565b3b151590565b6000602082840312156105da578081fd5b81356104c8816107fb565b6000602082840312156105f6578081fd5b81516104c8816107fb565b600060208284031215610612578081fd5b5051919050565b90565b60208082526023908201527f676f7665726e616e63652e657870657274732e45505146492e706172616d657460408201526265727360e81b606082015260800190565b6020808252601c908201527f746f6b656e65636f6e6f6d6963732e73797374656d5265736572766500000000604082015260600190565b6020808252601d908201527f676f7665726e65642e45505146492e515f726573657276655368617265000000604082015260600190565b6020808252605c908201527f5b5145432d3031343030315d2d4661696c656420746f207472616e736665722060408201527f746865205120746f6b656e20686f6c6465722072657761726420746f2074686560608201527f2051486f6c646572526577617264506f6f6c20636f6e74726163742e00000000608082015260a00190565b6020808252818101527f746f6b656e65636f6e6f6d6963732e71486f6c646572526577617264506f6f6c604082015260600190565b60208082526050908201527f5b5145432d3031343030305d2d4661696c656420746f207472616e736665722060408201527f746865207265736572766520736861726520746f207468652053797374656d5260608201526f32b9b2b93b329031b7b73a3930b1ba1760811b608082015260a00190565b6001600160a01b038116811461081057600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212208c48cd83429ab15d0228cdd0ba1da36c8e6ed04a0832782a189d1a8e3cebbe8464736f6c63430007060033",
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
