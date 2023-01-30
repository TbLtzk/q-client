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

// AdminUpgradeabilityProxyMetaData contains all meta data concerning the AdminUpgradeabilityProxy contract.
var AdminUpgradeabilityProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052604051610d7b380380610d7b8339818101604052606081101561002657600080fd5b8101908080519060200190929190805190602001909291908051604051939291908464010000000082111561005a57600080fd5b8382019150602082018581111561007057600080fd5b825186600182028301116401000000008211171561008d57600080fd5b8083526020830192505050908051906020019080838360005b838110156100c15780820151818401526020810190506100a6565b50505050905090810190601f1680156100ee5780820380516001836020036101000a031916815260200191505b506040525050508281600160405180807f656970313936372e70726f78792e696d706c656d656e746174696f6e00000000815250601c019050604051809103902060001c0360001b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b1461016157fe5b6101708261032860201b60201c565b6000815111156102a55760008273ffffffffffffffffffffffffffffffffffffffff16826040518082805190602001908083835b602083106101c757805182526020820191506020810190506020830392506101a4565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610227576040519150601f19603f3d011682016040523d82523d6000602084013e61022c565b606091505b50509050806102a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f64656c6567617465642063616c6c206661696c6564000000000000000000000081525060200191505060405180910390fd5b505b5050600160405180807f656970313936372e70726f78792e61646d696e000000000000000000000000008152506013019050604051809103902060001c0360001b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b1461031157fe5b610320826103bf60201b60201c565b505050610401565b61033b816103ee60201b6108221760201c565b610390576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b815260200180610d40603b913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b600080823b905060008111915050919050565b610930806104106000396000f3fe60806040526004361061004a5760003560e01c80633659cfe6146100545780634f1ef286146100a55780635c60da1b1461013e5780638f28397014610195578063f851a440146101e6575b61005261023d565b005b34801561006057600080fd5b506100a36004803603602081101561007757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610257565b005b61013c600480360360408110156100bb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100f857600080fd5b82018360208201111561010a57600080fd5b8035906020019184600183028401116401000000008311171561012c57600080fd5b90919293919293905050506102ac565b005b34801561014a57600080fd5b506101536103ce565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101a157600080fd5b506101e4600480360360208110156101b857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610426565b005b3480156101f257600080fd5b506101fb61059f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102456105f7565b61025561025061068d565b6106be565b565b61025f6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102a05761029b81610715565b6102a9565b6102a861023d565b5b50565b6102b46106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103c0576102f083610715565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d806000811461035b576040519150601f19603f3d011682016040523d82523d6000602084013e610360565b606091505b50509050806103ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061089e6023913960400191505060405180910390fd5b506103c9565b6103c861023d565b5b505050565b60006103d86106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561041a5761041361068d565b9050610423565b61042261023d565b5b90565b61042e6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561059357600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806108686036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6105106106e4565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161058e81610764565b61059c565b61059b61023d565b5b50565b60006105a96106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105eb576105e46106e4565b90506105f4565b6105f361023d565b5b90565b6105ff6106e4565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610683576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806108366032913960400191505060405180910390fd5b61068b610793565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e80600081146106df573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b61071e81610795565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b565b61079e81610822565b6107f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806108c1603b913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737364656c6567617465642063616c6c206661696c6564206166746572207570677261646543616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a723158200cff5bad2f803a2127039eba3a278e1d70096ece8708d9ca347834fd0b404ae864736f6c6343000511003243616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373",
}

// AdminUpgradeabilityProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use AdminUpgradeabilityProxyMetaData.ABI instead.
var AdminUpgradeabilityProxyABI = AdminUpgradeabilityProxyMetaData.ABI

// AdminUpgradeabilityProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdminUpgradeabilityProxyMetaData.Bin instead.
var AdminUpgradeabilityProxyBin = AdminUpgradeabilityProxyMetaData.Bin

// DeployAdminUpgradeabilityProxy deploys a new Ethereum contract, binding an instance of AdminUpgradeabilityProxy to it.
func DeployAdminUpgradeabilityProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, _admin common.Address, _data []byte) (common.Address, *types.Transaction, *AdminUpgradeabilityProxy, error) {
	parsed, err := AdminUpgradeabilityProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdminUpgradeabilityProxyBin), backend, _logic, _admin, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// AdminUpgradeabilityProxy is an auto generated Go binding around an Ethereum contract.
type AdminUpgradeabilityProxy struct {
	AdminUpgradeabilityProxyCaller     // Read-only binding to the contract
	AdminUpgradeabilityProxyTransactor // Write-only binding to the contract
	AdminUpgradeabilityProxyFilterer   // Log filterer for contract events
}

// AdminUpgradeabilityProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminUpgradeabilityProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminUpgradeabilityProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminUpgradeabilityProxySession struct {
	Contract     *AdminUpgradeabilityProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminUpgradeabilityProxyCallerSession struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AdminUpgradeabilityProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminUpgradeabilityProxyTransactorSession struct {
	Contract     *AdminUpgradeabilityProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AdminUpgradeabilityProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyRaw struct {
	Contract *AdminUpgradeabilityProxy // Generic contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyCallerRaw struct {
	Contract *AdminUpgradeabilityProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AdminUpgradeabilityProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminUpgradeabilityProxyTransactorRaw struct {
	Contract *AdminUpgradeabilityProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminUpgradeabilityProxy creates a new instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxy(address common.Address, backend bind.ContractBackend) (*AdminUpgradeabilityProxy, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxy{AdminUpgradeabilityProxyCaller: AdminUpgradeabilityProxyCaller{contract: contract}, AdminUpgradeabilityProxyTransactor: AdminUpgradeabilityProxyTransactor{contract: contract}, AdminUpgradeabilityProxyFilterer: AdminUpgradeabilityProxyFilterer{contract: contract}}, nil
}

// NewAdminUpgradeabilityProxyCaller creates a new read-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyCaller(address common.Address, caller bind.ContractCaller) (*AdminUpgradeabilityProxyCaller, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyCaller{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyTransactor creates a new write-only instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminUpgradeabilityProxyTransactor, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyTransactor{contract: contract}, nil
}

// NewAdminUpgradeabilityProxyFilterer creates a new log filterer instance of AdminUpgradeabilityProxy, bound to a specific deployed contract.
func NewAdminUpgradeabilityProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminUpgradeabilityProxyFilterer, error) {
	contract, err := bindAdminUpgradeabilityProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyFilterer{contract: contract}, nil
}

// bindAdminUpgradeabilityProxy binds a generic wrapper to an already deployed contract.
func bindAdminUpgradeabilityProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminUpgradeabilityProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.AdminUpgradeabilityProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdminUpgradeabilityProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) Admin() (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Admin(&_AdminUpgradeabilityProxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) Admin() (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Admin(&_AdminUpgradeabilityProxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.ChangeAdmin(&_AdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.ChangeAdmin(&_AdminUpgradeabilityProxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) Implementation() (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Implementation(&_AdminUpgradeabilityProxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Implementation(&_AdminUpgradeabilityProxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.UpgradeTo(&_AdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.UpgradeTo(&_AdminUpgradeabilityProxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_AdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.UpgradeToAndCall(&_AdminUpgradeabilityProxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Fallback(&_AdminUpgradeabilityProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AdminUpgradeabilityProxy.Contract.Fallback(&_AdminUpgradeabilityProxy.TransactOpts, calldata)
}

// AdminUpgradeabilityProxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyAdminChangedIterator struct {
	Event *AdminUpgradeabilityProxyAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminUpgradeabilityProxyAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminUpgradeabilityProxyAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminUpgradeabilityProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminUpgradeabilityProxyAdminChanged represents a AdminChanged event raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AdminUpgradeabilityProxyAdminChangedIterator, error) {

	logs, sub, err := _AdminUpgradeabilityProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyAdminChangedIterator{contract: _AdminUpgradeabilityProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AdminUpgradeabilityProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AdminUpgradeabilityProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminUpgradeabilityProxyAdminChanged)
				if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) ParseAdminChanged(log types.Log) (*AdminUpgradeabilityProxyAdminChanged, error) {
	event := new(AdminUpgradeabilityProxyAdminChanged)
	if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdminUpgradeabilityProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyUpgradedIterator struct {
	Event *AdminUpgradeabilityProxyUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdminUpgradeabilityProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminUpgradeabilityProxyUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdminUpgradeabilityProxyUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdminUpgradeabilityProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminUpgradeabilityProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminUpgradeabilityProxyUpgraded represents a Upgraded event raised by the AdminUpgradeabilityProxy contract.
type AdminUpgradeabilityProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AdminUpgradeabilityProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AdminUpgradeabilityProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AdminUpgradeabilityProxyUpgradedIterator{contract: _AdminUpgradeabilityProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AdminUpgradeabilityProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AdminUpgradeabilityProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminUpgradeabilityProxyUpgraded)
				if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AdminUpgradeabilityProxy *AdminUpgradeabilityProxyFilterer) ParseUpgraded(log types.Log) (*AdminUpgradeabilityProxyUpgraded, error) {
	event := new(AdminUpgradeabilityProxyUpgraded)
	if err := _AdminUpgradeabilityProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
