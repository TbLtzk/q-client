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
	Bin: "0x608060405234801561001057600080fd5b50610bc1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806362c54dbd1461003b578063c4d66de814610064575b600080fd5b61004e6100493660046102b6565b610079565b60405161005b919061037b565b60405180910390f35b61007761007236600461038f565b61019d565b005b6000805460405182916201000090046001600160a01b03169061009b90610277565b6001600160a01b039091168152604060208201819052600090820152606001604051809103906000f0801580156100d6573d6000803e3d6000fd5b5060405163a224cee760e01b81529091506001600160a01b0382169063a224cee7906101069086906004016103b1565b600060405180830381600087803b15801561012057600080fd5b505af1158015610134573d6000803e3d6000fd5b505060405163f2fde38b60e01b81526001600160a01b038416925063f2fde38b915061016490339060040161037b565b600060405180830381600087803b15801561017e57600080fd5b505af1158015610192573d6000803e3d6000fd5b509295945050505050565b600054610100900460ff16806101b6575060005460ff16155b61021d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff1615801561023f576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610273576000805461ff00191690555b5050565b61078d806103ff83390190565b634e487b7160e01b600052604160045260246000fd5b80356001600160a01b03811681146102b157600080fd5b919050565b600060208083850312156102c957600080fd5b823567ffffffffffffffff808211156102e157600080fd5b818501915085601f8301126102f557600080fd5b81358181111561030757610307610284565b8060051b604051601f19603f8301168101818110858211171561032c5761032c610284565b60405291825284820192508381018501918883111561034a57600080fd5b938501935b8285101561036f576103608561029a565b8452938501939285019261034f565b98975050505050505050565b6001600160a01b0391909116815260200190565b6000602082840312156103a157600080fd5b6103aa8261029a565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156103f25783516001600160a01b0316835292840192918401916001016103cd565b5090969550505050505056fe608060405260405161078d38038061078d83398101604081905261002291610337565b61004d60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd610405565b600080516020610746833981519152146100695761006961042a565b6100758282600061007c565b505061048f565b610085836100b2565b6000825111806100925750805b156100ad576100ab83836100f260201b6100291760201c565b505b505050565b6100bb8161011e565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606101178383604051806060016040528060278152602001610766602791396101de565b9392505050565b610131816102b360201b6100551760201c565b6101985760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806101bd60008051602061074683398151915260001b6102b960201b61005b1760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060833b61023d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161018f565b600080856001600160a01b0316856040516102589190610440565b600060405180830381855af49150503d8060008114610293576040519150601f19603f3d011682016040523d82523d6000602084013e610298565b606091505b5090925090506102a98282866102bc565b9695505050505050565b3b151590565b90565b606083156102cb575081610117565b8251156102db5782518084602001fd5b8160405162461bcd60e51b815260040161018f919061045c565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561032657818101518382015260200161030e565b838111156100ab5750506000910152565b6000806040838503121561034a57600080fd5b82516001600160a01b038116811461036157600080fd5b60208401519092506001600160401b038082111561037e57600080fd5b818501915085601f83011261039257600080fd5b8151818111156103a4576103a46102f5565b604051601f8201601f19908116603f011681019083821181831017156103cc576103cc6102f5565b816040528281528860208487010111156103e557600080fd5b6103f683602083016020880161030b565b80955050505050509250929050565b60008282101561042557634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052600160045260246000fd5b6000825161045281846020870161030b565b9190910192915050565b602081526000825180602084015261047b81604085016020870161030b565b601f01601f19169190910160400192915050565b6102a88061049e6000396000f3fe60806040523661001357610011610017565b005b6100115b61002761002261005e565b610096565b565b606061004e838360405180606001604052806027815260200161024c602791396100ba565b9392505050565b3b151590565b90565b60006100917f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100b5573d6000f35b3d6000fd5b6060833b61011e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084015b60405180910390fd5b600080856001600160a01b03168560405161013991906101fc565b600060405180830381855af49150503d8060008114610174576040519150601f19603f3d011682016040523d82523d6000602084013e610179565b606091505b5091509150610189828286610193565b9695505050505050565b606083156101a257508161004e565b8251156101b25782518084602001fd5b8160405162461bcd60e51b81526004016101159190610218565b60005b838110156101e75781810151838201526020016101cf565b838111156101f6576000848401525b50505050565b6000825161020e8184602087016101cc565b9190910192915050565b60208152600082518060208401526102378160408501602087016101cc565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e849187d041a1239355c10a41a65b28858052b2cac3f065bceb24af0a86ed8b964736f6c63430008090033360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200b70453a967540fbb15f4aebc2ea54f90483898b69b5d8fa3b562cf6b41e572464736f6c63430008090033",
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
