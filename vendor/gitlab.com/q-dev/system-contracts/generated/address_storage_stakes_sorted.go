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

// AddressStorageStakesSortedMetaData contains all meta data concerning the AddressStorageStakesSorted contract.
var AddressStorageStakesSortedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"addAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"linkedList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLast\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"updateStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a33610058565b600160005260026020527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b03191690556100a8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610e90806100b76000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c806340b7a9361461009e5780635dbe47e8146100c6578063715018a6146100d95780638da5cb5b146100e3578063972c5356146101035780639f65080b1461011a578063a39fac121461012d578063cf527e8b14610142578063d092e1861461016b578063d903450314610173578063f2fde38b14610193575b600080fd5b6100b16100ac366004610ceb565b6101a6565b60405190151581526020015b60405180910390f35b6100b16100d4366004610d15565b6104ac565b6100e16104c9565b005b6100eb610504565b6040516001600160a01b0390911681526020016100bd565b61010c60035481565b6040519081526020016100bd565b6100b1610128366004610ceb565b610513565b6101356107a5565b6040516100bd9190610d37565b6100eb610150366004610d15565b6002602052600090815260409020546001600160a01b031681565b6100eb61086d565b61010c610181366004610d15565b60016020526000908152604090205481565b6100e16101a1366004610d15565b6109d2565b6000336101b1610504565b6001600160a01b0316146101e05760405162461bcd60e51b81526004016101d790610d84565b60405180910390fd5b826001600160a01b0381161580159061020357506001600160a01b038116600114155b61026c5760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3033373030355d2d54686520616464726573732073686f756c64206044820152733737ba103132902422a0a21037b9102a20a4a61760611b60648201526084016101d7565b610275846104ac565b1561032a5760405162461bcd60e51b815260206004820152607360248201527f5b5145432d3033373030305d2d54686520616464726573732068617320616c7260448201527f65616479206265656e20616464656420746f207468652073746f726167652c2060648201527f6661696c656420746f2061646420746f2074686520616464726573732073746160848201527235b2b99039b7b93a32b21039ba37b930b3b29760691b60a482015260c4016101d7565b600083116103ae5760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3033373030315d2d496e76616c6964207374616b652c206661696c60448201527f656420746f2061646420746865206164647265737320746f207468652073746f6064820152643930b3b29760d91b608482015260a4016101d7565b60006103b984610a72565b6001600160a01b03808216600090815260026020526040902054919250161580610413576001600160a01b03808316600090815260026020526040808220548984168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020908152604080832080546001600160a01b031916948b16948517905592825260019052908120869055600380549161045e83610dcf565b90915550506040516001600160a01b038716907fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f90600090a26104a086610aeb565b50600195945050505050565b6001600160a01b0316600090815260016020526040902054151590565b336104d2610504565b6001600160a01b0316146104f85760405162461bcd60e51b81526004016101d790610d84565b6105026000610b9f565b565b6000546001600160a01b031690565b60003361051e610504565b6001600160a01b0316146105445760405162461bcd60e51b81526004016101d790610d84565b61054d836104ac565b6105d35760405162461bcd60e51b815260206004820152604b60248201527f5b5145432d3033373030325d2d5468652061646472657373206973206e6f742060448201527f696e207468652073746f726167652c206661696c656420746f2075706461746560648201526a103a34329039ba30b5b29760a91b608482015260a4016101d7565b6001600160a01b03808416600090815260026020819052604082205490921691906105fd86610bef565b6001600160a01b0390811682526020808301939093526040918201600090812080549583166001600160a01b031996871617905590871681526002909252902080549091169055816106a7576003805490600061065983610dea565b90915550506001600160a01b038316600081815260016020526040808220829055517f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb9190a250600161079f565b6001600160a01b03831660008181526001602052604090819020849055517fab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb906106f49085815260200190565b60405180910390a2600061070783610a72565b6001600160a01b03808216600090815260026020526040902054919250161580610761576001600160a01b03808316600090815260026020526040808220548884168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020526040902080546001600160a01b03191691871691909117905561079885610aeb565b6001925050505b92915050565b6060600060019050600060035467ffffffffffffffff8111156107ca576107ca610e01565b6040519080825280602002602001820160405280156107f3578160200160208202803683370190505b506003549091505b8015610866576001600160a01b0392831660009081526002602052604090205490921691828261082c600184610e17565b8151811061083c5761083c610e2e565b6001600160a01b03909216602092830291909101909101528061085e81610dea565b9150506107fb565b5092915050565b600033610878610504565b6001600160a01b03161461089e5760405162461bcd60e51b81526004016101d790610d84565b6000600354116109165760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3033373030335d2d546865206c69737420697320656d7074792c2060448201527f6661696c656420746f2072656d6f76652074686520616464726573732e00000060648201526084016101d7565b600260209081527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b038082166000818152604080822080549094166001600160a01b03199586161790955582549093169091556001909352908120819055600380549161098b83610dea565b90915550506040516001600160a01b038216907f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb90600090a26109cd81610aeb565b905090565b336109db610504565b6001600160a01b031614610a015760405162461bcd60e51b81526004016101d790610d84565b6001600160a01b038116610a665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101d7565b610a6f81610b9f565b50565b600060015b6001600160a01b03808216600090815260026020908152604080832054909316825260019052205483118015610ac657506001600160a01b038181166000908152600260205260409020541615155b1561079f576001600160a01b0390811660009081526002602052604090205416610a77565b610af4816104ac565b15610b23576001600160a01b038116600090815260016020526040902054610b1e57610b1e610e44565b610b49565b6001600160a01b03811660009081526001602052604090205415610b4957610b49610e44565b6001600160a01b03808216600090815260026020526040902054168015610b9b576001600160a01b038083166000908152600160205260408082205492841682529020541015610b9b57610b9b610e44565b5050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610bfa826104ac565b610c885760405162461bcd60e51b815260206004820152605360248201527f5b5145432d3033373030345d2d5468652061646472657373206973206e6f742060448201527f696e207468652073746f726167652c206661696c656420746f206765742074686064820152723290383932bb34b7bab99030b2323932b9b99760691b608482015260a4016101d7565b60015b6001600160a01b0381811660009081526002602052604090205481169084161461079f576001600160a01b0390811660009081526002602052604090205416610c8b565b80356001600160a01b0381168114610ce657600080fd5b919050565b60008060408385031215610cfe57600080fd5b610d0783610ccf565b946020939093013593505050565b600060208284031215610d2757600080fd5b610d3082610ccf565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610d785783516001600160a01b031683529284019291840191600101610d53565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000600019821415610de357610de3610db9565b5060010190565b600081610df957610df9610db9565b506000190190565b634e487b7160e01b600052604160045260246000fd5b600082821015610e2957610e29610db9565b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052600160045260246000fdfea2646970667358221220f16da91a38ec82b2a963da54b9873e02ac06b3dbae7278432d0f39dec8ba665964736f6c63430008090033",
}

// AddressStorageStakesSortedABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageStakesSortedMetaData.ABI instead.
var AddressStorageStakesSortedABI = AddressStorageStakesSortedMetaData.ABI

// AddressStorageStakesSortedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageStakesSortedMetaData.Bin instead.
var AddressStorageStakesSortedBin = AddressStorageStakesSortedMetaData.Bin

// DeployAddressStorageStakesSorted deploys a new Ethereum contract, binding an instance of AddressStorageStakesSorted to it.
func DeployAddressStorageStakesSorted(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorageStakesSorted, error) {
	parsed, err := AddressStorageStakesSortedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageStakesSortedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorageStakesSorted{AddressStorageStakesSortedCaller: AddressStorageStakesSortedCaller{contract: contract}, AddressStorageStakesSortedTransactor: AddressStorageStakesSortedTransactor{contract: contract}, AddressStorageStakesSortedFilterer: AddressStorageStakesSortedFilterer{contract: contract}}, nil
}

// AddressStorageStakesSorted is an auto generated Go binding around an Ethereum contract.
type AddressStorageStakesSorted struct {
	AddressStorageStakesSortedCaller     // Read-only binding to the contract
	AddressStorageStakesSortedTransactor // Write-only binding to the contract
	AddressStorageStakesSortedFilterer   // Log filterer for contract events
}

// AddressStorageStakesSortedCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageStakesSortedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesSortedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageStakesSortedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesSortedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageStakesSortedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesSortedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageStakesSortedSession struct {
	Contract     *AddressStorageStakesSorted // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AddressStorageStakesSortedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageStakesSortedCallerSession struct {
	Contract *AddressStorageStakesSortedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// AddressStorageStakesSortedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageStakesSortedTransactorSession struct {
	Contract     *AddressStorageStakesSortedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// AddressStorageStakesSortedRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageStakesSortedRaw struct {
	Contract *AddressStorageStakesSorted // Generic contract binding to access the raw methods on
}

// AddressStorageStakesSortedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageStakesSortedCallerRaw struct {
	Contract *AddressStorageStakesSortedCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageStakesSortedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageStakesSortedTransactorRaw struct {
	Contract *AddressStorageStakesSortedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorageStakesSorted creates a new instance of AddressStorageStakesSorted, bound to a specific deployed contract.
func NewAddressStorageStakesSorted(address common.Address, backend bind.ContractBackend) (*AddressStorageStakesSorted, error) {
	contract, err := bindAddressStorageStakesSorted(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSorted{AddressStorageStakesSortedCaller: AddressStorageStakesSortedCaller{contract: contract}, AddressStorageStakesSortedTransactor: AddressStorageStakesSortedTransactor{contract: contract}, AddressStorageStakesSortedFilterer: AddressStorageStakesSortedFilterer{contract: contract}}, nil
}

// NewAddressStorageStakesSortedCaller creates a new read-only instance of AddressStorageStakesSorted, bound to a specific deployed contract.
func NewAddressStorageStakesSortedCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageStakesSortedCaller, error) {
	contract, err := bindAddressStorageStakesSorted(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedCaller{contract: contract}, nil
}

// NewAddressStorageStakesSortedTransactor creates a new write-only instance of AddressStorageStakesSorted, bound to a specific deployed contract.
func NewAddressStorageStakesSortedTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageStakesSortedTransactor, error) {
	contract, err := bindAddressStorageStakesSorted(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedTransactor{contract: contract}, nil
}

// NewAddressStorageStakesSortedFilterer creates a new log filterer instance of AddressStorageStakesSorted, bound to a specific deployed contract.
func NewAddressStorageStakesSortedFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageStakesSortedFilterer, error) {
	contract, err := bindAddressStorageStakesSorted(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedFilterer{contract: contract}, nil
}

// bindAddressStorageStakesSorted binds a generic wrapper to an already deployed contract.
func bindAddressStorageStakesSorted(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressStorageStakesSortedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakesSorted.Contract.AddressStorageStakesSortedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.AddressStorageStakesSortedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.AddressStorageStakesSortedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakesSorted.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.contract.Transact(opts, method, params...)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) AddrStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "addrStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) AddrStake(arg0 common.Address) (*big.Int, error) {
	return _AddressStorageStakesSorted.Contract.AddrStake(&_AddressStorageStakesSorted.CallOpts, arg0)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) AddrStake(arg0 common.Address) (*big.Int, error) {
	return _AddressStorageStakesSorted.Contract.AddrStake(&_AddressStorageStakesSorted.CallOpts, arg0)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) Contains(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "contains", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakesSorted.Contract.Contains(&_AddressStorageStakesSorted.CallOpts, _addr)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakesSorted.Contract.Contains(&_AddressStorageStakesSorted.CallOpts, _addr)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakesSorted.Contract.GetAddresses(&_AddressStorageStakesSorted.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakesSorted.Contract.GetAddresses(&_AddressStorageStakesSorted.CallOpts)
}

// LinkedList is a free data retrieval call binding the contract method 0xcf527e8b.
//
// Solidity: function linkedList(address ) view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) LinkedList(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "linkedList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkedList is a free data retrieval call binding the contract method 0xcf527e8b.
//
// Solidity: function linkedList(address ) view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) LinkedList(arg0 common.Address) (common.Address, error) {
	return _AddressStorageStakesSorted.Contract.LinkedList(&_AddressStorageStakesSorted.CallOpts, arg0)
}

// LinkedList is a free data retrieval call binding the contract method 0xcf527e8b.
//
// Solidity: function linkedList(address ) view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) LinkedList(arg0 common.Address) (common.Address, error) {
	return _AddressStorageStakesSorted.Contract.LinkedList(&_AddressStorageStakesSorted.CallOpts, arg0)
}

// ListSize is a free data retrieval call binding the contract method 0x972c5356.
//
// Solidity: function listSize() view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) ListSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "listSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListSize is a free data retrieval call binding the contract method 0x972c5356.
//
// Solidity: function listSize() view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) ListSize() (*big.Int, error) {
	return _AddressStorageStakesSorted.Contract.ListSize(&_AddressStorageStakesSorted.CallOpts)
}

// ListSize is a free data retrieval call binding the contract method 0x972c5356.
//
// Solidity: function listSize() view returns(uint256)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) ListSize() (*big.Int, error) {
	return _AddressStorageStakesSorted.Contract.ListSize(&_AddressStorageStakesSorted.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakesSorted.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) Owner() (common.Address, error) {
	return _AddressStorageStakesSorted.Contract.Owner(&_AddressStorageStakesSorted.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedCallerSession) Owner() (common.Address, error) {
	return _AddressStorageStakesSorted.Contract.Owner(&_AddressStorageStakesSorted.CallOpts)
}

// AddAddr is a paid mutator transaction binding the contract method 0x40b7a936.
//
// Solidity: function addAddr(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactor) AddAddr(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.contract.Transact(opts, "addAddr", _addr, _stake)
}

// AddAddr is a paid mutator transaction binding the contract method 0x40b7a936.
//
// Solidity: function addAddr(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) AddAddr(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.AddAddr(&_AddressStorageStakesSorted.TransactOpts, _addr, _stake)
}

// AddAddr is a paid mutator transaction binding the contract method 0x40b7a936.
//
// Solidity: function addAddr(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorSession) AddAddr(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.AddAddr(&_AddressStorageStakesSorted.TransactOpts, _addr, _stake)
}

// RemoveLast is a paid mutator transaction binding the contract method 0xd092e186.
//
// Solidity: function removeLast() returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactor) RemoveLast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.contract.Transact(opts, "removeLast")
}

// RemoveLast is a paid mutator transaction binding the contract method 0xd092e186.
//
// Solidity: function removeLast() returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) RemoveLast() (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.RemoveLast(&_AddressStorageStakesSorted.TransactOpts)
}

// RemoveLast is a paid mutator transaction binding the contract method 0xd092e186.
//
// Solidity: function removeLast() returns(address)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorSession) RemoveLast() (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.RemoveLast(&_AddressStorageStakesSorted.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.RenounceOwnership(&_AddressStorageStakesSorted.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.RenounceOwnership(&_AddressStorageStakesSorted.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.TransferOwnership(&_AddressStorageStakesSorted.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.TransferOwnership(&_AddressStorageStakesSorted.TransactOpts, newOwner)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x9f65080b.
//
// Solidity: function updateStake(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactor) UpdateStake(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.contract.Transact(opts, "updateStake", _addr, _stake)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x9f65080b.
//
// Solidity: function updateStake(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedSession) UpdateStake(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.UpdateStake(&_AddressStorageStakesSorted.TransactOpts, _addr, _stake)
}

// UpdateStake is a paid mutator transaction binding the contract method 0x9f65080b.
//
// Solidity: function updateStake(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedTransactorSession) UpdateStake(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakesSorted.Contract.UpdateStake(&_AddressStorageStakesSorted.TransactOpts, _addr, _stake)
}

// AddressStorageStakesSortedAddressAddedIterator is returned from FilterAddressAdded and is used to iterate over the raw logs and unpacked data for AddressAdded events raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedAddressAddedIterator struct {
	Event *AddressStorageStakesSortedAddressAdded // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesSortedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesSortedAddressAdded)
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
		it.Event = new(AddressStorageStakesSortedAddressAdded)
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
func (it *AddressStorageStakesSortedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesSortedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesSortedAddressAdded represents a AddressAdded event raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressAdded is a free log retrieval operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) FilterAddressAdded(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesSortedAddressAddedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.FilterLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedAddressAddedIterator{contract: _AddressStorageStakesSorted.contract, event: "AddressAdded", logs: logs, sub: sub}, nil
}

// WatchAddressAdded is a free log subscription operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) WatchAddressAdded(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesSortedAddressAdded, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.WatchLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesSortedAddressAdded)
				if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "AddressAdded", log); err != nil {
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

// ParseAddressAdded is a log parse operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) ParseAddressAdded(log types.Log) (*AddressStorageStakesSortedAddressAdded, error) {
	event := new(AddressStorageStakesSortedAddressAdded)
	if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "AddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesSortedAddressRemovedIterator is returned from FilterAddressRemoved and is used to iterate over the raw logs and unpacked data for AddressRemoved events raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedAddressRemovedIterator struct {
	Event *AddressStorageStakesSortedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesSortedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesSortedAddressRemoved)
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
		it.Event = new(AddressStorageStakesSortedAddressRemoved)
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
func (it *AddressStorageStakesSortedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesSortedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesSortedAddressRemoved represents a AddressRemoved event raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressRemoved is a free log retrieval operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) FilterAddressRemoved(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesSortedAddressRemovedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.FilterLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedAddressRemovedIterator{contract: _AddressStorageStakesSorted.contract, event: "AddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAddressRemoved is a free log subscription operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) WatchAddressRemoved(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesSortedAddressRemoved, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.WatchLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesSortedAddressRemoved)
				if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
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

// ParseAddressRemoved is a log parse operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) ParseAddressRemoved(log types.Log) (*AddressStorageStakesSortedAddressRemoved, error) {
	event := new(AddressStorageStakesSortedAddressRemoved)
	if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesSortedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedOwnershipTransferredIterator struct {
	Event *AddressStorageStakesSortedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesSortedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesSortedOwnershipTransferred)
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
		it.Event = new(AddressStorageStakesSortedOwnershipTransferred)
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
func (it *AddressStorageStakesSortedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesSortedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesSortedOwnershipTransferred represents a OwnershipTransferred event raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressStorageStakesSortedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedOwnershipTransferredIterator{contract: _AddressStorageStakesSorted.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesSortedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesSortedOwnershipTransferred)
				if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) ParseOwnershipTransferred(log types.Log) (*AddressStorageStakesSortedOwnershipTransferred, error) {
	event := new(AddressStorageStakesSortedOwnershipTransferred)
	if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesSortedStakeUpdatedIterator is returned from FilterStakeUpdated and is used to iterate over the raw logs and unpacked data for StakeUpdated events raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedStakeUpdatedIterator struct {
	Event *AddressStorageStakesSortedStakeUpdated // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesSortedStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesSortedStakeUpdated)
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
		it.Event = new(AddressStorageStakesSortedStakeUpdated)
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
func (it *AddressStorageStakesSortedStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesSortedStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesSortedStakeUpdated represents a StakeUpdated event raised by the AddressStorageStakesSorted contract.
type AddressStorageStakesSortedStakeUpdated struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeUpdated is a free log retrieval operation binding the contract event 0xab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb.
//
// Solidity: event StakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) FilterStakeUpdated(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesSortedStakeUpdatedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.FilterLogs(opts, "StakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesSortedStakeUpdatedIterator{contract: _AddressStorageStakesSorted.contract, event: "StakeUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeUpdated is a free log subscription operation binding the contract event 0xab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb.
//
// Solidity: event StakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) WatchStakeUpdated(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesSortedStakeUpdated, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakesSorted.contract.WatchLogs(opts, "StakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesSortedStakeUpdated)
				if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "StakeUpdated", log); err != nil {
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

// ParseStakeUpdated is a log parse operation binding the contract event 0xab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb.
//
// Solidity: event StakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakesSorted *AddressStorageStakesSortedFilterer) ParseStakeUpdated(log types.Log) (*AddressStorageStakesSortedStakeUpdated, error) {
	event := new(AddressStorageStakesSortedStakeUpdated)
	if err := _AddressStorageStakesSorted.contract.UnpackLog(event, "StakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
