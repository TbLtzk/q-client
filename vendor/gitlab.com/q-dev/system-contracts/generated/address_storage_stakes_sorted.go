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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"linkedList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"addAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"updateStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLast\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600061001b6100a3565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600160005260026020527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b03191690556100a7565b3390565b610eab806100b66000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c806340b7a9361461009e5780635dbe47e8146100de578063715018a6146101045780638da5cb5b1461010e578063972c5356146101325780639f65080b1461014c578063a39fac1214610178578063cf527e8b146101d0578063d092e186146101f6578063d9034503146101fe578063f2fde38b14610224575b600080fd5b6100ca600480360360408110156100b457600080fd5b506001600160a01b03813516906020013561024a565b604080519115158252519081900360200190f35b6100ca600480360360208110156100f457600080fd5b50356001600160a01b031661046e565b61010c61048b565b005b61011661052d565b604080516001600160a01b039092168252519081900360200190f35b61013a61053c565b60408051918252519081900360200190f35b6100ca6004803603604081101561016257600080fd5b506001600160a01b038135169060200135610542565b6101806107a2565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101bc5781810151838201526020016101a4565b505050509050019250505060405180910390f35b610116600480360360208110156101e657600080fd5b50356001600160a01b0316610856565b610116610871565b61013a6004803603602081101561021457600080fd5b50356001600160a01b03166109b2565b61010c6004803603602081101561023a57600080fd5b50356001600160a01b03166109c4565b6000610254610abc565b6000546001600160a01b039081169116146102a4576040805162461bcd60e51b81526020600482018190526024820152600080516020610dbe833981519152604482015290519081900360640190fd5b826001600160a01b038116158015906102c757506001600160a01b038116600114155b6103025760405162461bcd60e51b8152600401808060200182810382526034815260200180610cb46034913960400191505060405180910390fd5b61030b8461046e565b156103475760405162461bcd60e51b8152600401808060200182810382526073815260200180610d4b6073913960800191505060405180910390fd5b600083116103865760405162461bcd60e51b8152600401808060200182810382526045815260200180610dde6045913960600191505060405180910390fd5b600061039184610ac0565b6001600160a01b038082166000908152600260205260409020549192501615806103eb576001600160a01b03808316600090815260026020526040808220548984168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020908152604080832080546001600160a01b031916948b16948517905583835260019182905280832089905560038054909201909155517fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f9190a261046286610b39565b50600195945050505050565b6001600160a01b0316600090815260016020526040902054151590565b610493610abc565b6000546001600160a01b039081169116146104e3576040805162461bcd60e51b81526020600482018190526024820152600080516020610dbe833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b60035481565b600061054c610abc565b6000546001600160a01b0390811691161461059c576040805162461bcd60e51b81526020600482018190526024820152600080516020610dbe833981519152604482015290519081900360640190fd5b6105a58361046e565b6105e05760405162461bcd60e51b815260040180806020018281038252604b815260200180610c69604b913960600191505060405180910390fd5b6001600160a01b038084166000908152600260208190526040822054909216919061060a86610bdb565b6001600160a01b0390811682526020808301939093526040918201600090812080549583166001600160a01b031996871617905590871681526002909252902080549091169055816106a957600380546000190190556001600160a01b038316600081815260016020526040808220829055517f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb9190a250600161079c565b6001600160a01b038316600081815260016020908152604091829020859055815185815291517fab0e25dc39626189cfb41155020ba89e726b10244275733e9d7c63cf33ffccdb9281900390910190a2600061070483610ac0565b6001600160a01b0380821660009081526002602052604090205491925016158061075e576001600160a01b03808316600090815260026020526040808220548884168352912080546001600160a01b031916919092161790555b6001600160a01b03828116600090815260026020526040902080546001600160a01b03191691871691909117905561079585610b39565b6001925050505b92915050565b6060600060019050600060035467ffffffffffffffff811180156107c557600080fd5b506040519080825280602002602001820160405280156107ef578160200160208202803683370190505b506003549091505b801561084f576001600160a01b03928316600090815260026020526040902054825193169283908390600019840190811061082e57fe5b6001600160a01b0390921660209283029190910190910152600019016107f7565b5091505090565b6002602052600090815260409020546001600160a01b031681565b600061087b610abc565b6000546001600160a01b039081169116146108cb576040805162461bcd60e51b81526020600482018190526024820152600080516020610dbe833981519152604482015290519081900360640190fd5b60006003541161090c5760405162461bcd60e51b815260040180806020018281038252603d815260200180610d0e603d913960400191505060405180910390fd5b600260209081527fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e080546001600160a01b038082166000818152604080822080549094166001600160a01b0319958616179095558254909316909155600190935281812081905560038054600019019055905182917f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb91a26109ad81610b39565b905090565b60016020526000908152604090205481565b6109cc610abc565b6000546001600160a01b03908116911614610a1c576040805162461bcd60e51b81526020600482018190526024820152600080516020610dbe833981519152604482015290519081900360640190fd5b6001600160a01b038116610a615760405162461bcd60e51b8152600401808060200182810382526026815260200180610ce86026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b600060015b6001600160a01b03808216600090815260026020908152604080832054909316825260019052205483118015610b1457506001600160a01b038181166000908152600260205260409020541615155b1561079c576001600160a01b0390811660009081526002602052604090205416610ac5565b610b428161046e565b15610b6b576001600160a01b038116600090815260016020526040902054610b6657fe5b610b8b565b6001600160a01b03811660009081526001602052604090205415610b8b57fe5b6001600160a01b03808216600090815260026020526040902054168015610bd7576001600160a01b038083166000908152600160205260408082205492841682529020541015610bd757fe5b5050565b6000610be68261046e565b610c215760405162461bcd60e51b8152600401808060200182810382526053815260200180610e236053913960600191505060405180910390fd5b60015b6001600160a01b0381811660009081526002602052604090205481169084161461079c576001600160a01b0390811660009081526002602052604090205416610c2456fe5b5145432d3033373030325d2d5468652061646472657373206973206e6f7420696e207468652073746f726167652c206661696c656420746f2075706461746520746865207374616b652e5b5145432d3033373030355d2d54686520616464726573732073686f756c64206e6f742062652048454144206f72205441494c2e4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573735b5145432d3033373030335d2d546865206c69737420697320656d7074792c206661696c656420746f2072656d6f76652074686520616464726573732e5b5145432d3033373030305d2d54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746f207468652061646472657373207374616b657320736f727465642073746f726167652e4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725b5145432d3033373030315d2d496e76616c6964207374616b652c206661696c656420746f2061646420746865206164647265737320746f207468652073746f726167652e5b5145432d3033373030345d2d5468652061646472657373206973206e6f7420696e207468652073746f726167652c206661696c656420746f20676574207468652070726576696f757320616464726573732ea26469706673582212202e7303f66f99fcc4c890b09a871f44bc557a5e7c43eb071edcab4b368ce10f9f64736f6c63430007060033",
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
