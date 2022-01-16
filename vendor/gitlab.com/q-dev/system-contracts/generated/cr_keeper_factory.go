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
	Bin: "0x608060405234801561001057600080fd5b50610a73806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063c4d66de81461003b578063efc81a8c14610050575b600080fd5b61004e61004936600461026c565b61006e565b005b610058610148565b604051610065919061029c565b60405180910390f35b600054610100900460ff1680610087575060005460ff16155b6100ee5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015610110576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610144576000805461ff00191690555b5050565b6000805460405182916201000090046001600160a01b03169061016a9061025f565b6001600160a01b039091168152604060208201819052600090820152606001604051809103906000f0801580156101a5573d6000803e3d6000fd5b509050806001600160a01b0316638129fc1c6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156101e357600080fd5b505af11580156101f7573d6000803e3d6000fd5b505060405163f2fde38b60e01b81526001600160a01b038416925063f2fde38b915061022790339060040161029c565b600060405180830381600087803b15801561024157600080fd5b505af1158015610255573d6000803e3d6000fd5b5092949350505050565b61078d806102b183390190565b60006020828403121561027e57600080fd5b81356001600160a01b038116811461029557600080fd5b9392505050565b6001600160a01b039190911681526020019056fe608060405260405161078d38038061078d83398101604081905261002291610337565b61004d60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd610405565b600080516020610746833981519152146100695761006961042a565b6100758282600061007c565b505061048f565b610085836100b2565b6000825111806100925750805b156100ad576100ab83836100f260201b6100291760201c565b505b505050565b6100bb8161011e565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606101178383604051806060016040528060278152602001610766602791396101de565b9392505050565b610131816102b360201b6100551760201c565b6101985760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101bd60008051602061074683398151915260001b6102b960201b61005b1760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060833b61023d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161018f565b600080856001600160a01b0316856040516102589190610440565b600060405180830381855af49150503d8060008114610293576040519150601f19603f3d011682016040523d82523d6000602084013e610298565b606091505b5090925090506102a98282866102bc565b9695505050505050565b3b151590565b90565b606083156102cb575081610117565b8251156102db5782518084602001fd5b8160405162461bcd60e51b815260040161018f919061045c565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561032657818101518382015260200161030e565b838111156100ab5750506000910152565b6000806040838503121561034a57600080fd5b82516001600160a01b038116811461036157600080fd5b60208401519092506001600160401b038082111561037e57600080fd5b818501915085601f83011261039257600080fd5b8151818111156103a4576103a46102f5565b604051601f8201601f19908116603f011681019083821181831017156103cc576103cc6102f5565b816040528281528860208487010111156103e557600080fd5b6103f683602083016020880161030b565b80955050505050509250929050565b60008282101561042557634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b6000825161045281846020870161030b565b9190910192915050565b602081526000825180602084015261047b81604085016020870161030b565b601f01601f19169190910160400192915050565b6102a88061049e6000396000f3fe60806040523661001357610011610017565b005b6100115b61002761002261005e565b610096565b565b606061004e838360405180606001604052806027815260200161024c602791396100ba565b9392505050565b3b151590565b90565b60006100917f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100b5573d6000f35b3d6000fd5b6060833b61011e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b03168560405161013991906101fc565b600060405180830381855af49150503d8060008114610174576040519150601f19603f3d011682016040523d82523d6000602084013e610179565b606091505b5091509150610189828286610193565b9695505050505050565b606083156101a257508161004e565b8251156101b25782518084602001fd5b8160405162461bcd60e51b81526004016101159190610218565b60005b838110156101e75781810151838201526020016101cf565b838111156101f6576000848401525b50505050565b6000825161020e8184602087016101cc565b9190910192915050565b60208152600082518060208401526102378160408501602087016101cc565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e849187d041a1239355c10a41a65b28858052b2cac3f065bceb24af0a86ed8b964736f6c63430008090033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122068590ff60409451e80ed7d68826c870e183e68d810f65f94dc9eb84b6ad5450064736f6c63430008090033",
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
