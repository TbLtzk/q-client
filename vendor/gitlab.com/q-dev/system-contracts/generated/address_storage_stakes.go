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

// AddressStorageStakesMetaData contains all meta data concerning the AddressStorageStakes contract.
var AddressStorageStakesMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"AddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"DelegatedStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"StakeIncreased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addrList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedStake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_delegatedStake\",\"type\":\"uint256\"}],\"name\":\"setDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610c5e8061007e6000396000f3fe608060405234801561001057600080fd5b50600436106100995760003560e01c806326ffee081461009e5780635dbe47e8146100c65780636467b404146100d9578063715018a6146100ec57806380402564146100f65780638da5cb5b14610121578063949d225d14610129578063a39fac121461013a578063d90345031461014f578063f2fde38b14610198578063f5d82b6b146101ab575b600080fd5b6100b16100ac366004610aba565b6101be565b60405190151581526020015b60405180910390f35b6100b16100d4366004610ae4565b610387565b6100b16100e7366004610aba565b6103cc565b6100f4610498565b005b610109610104366004610b06565b6104d3565b6040516001600160a01b0390911681526020016100bd565b6101096104fd565b6002546040519081526020016100bd565b61014261050c565b6040516100bd9190610b1f565b61017d61015d366004610ae4565b600160208190526000918252604090912080549181015460029091015483565b604080519384526020840192909252908201526060016100bd565b6100f46101a6366004610ae4565b61056e565b6100b16101b9366004610aba565b61060e565b6000336101c96104fd565b6001600160a01b0316146101f85760405162461bcd60e51b81526004016101ef90610b6c565b60405180910390fd5b6001600160a01b038316600090815260016020526040902080546102935760405162461bcd60e51b815260206004820152604660248201527f5b5145432d3033363030305d2d546865206164647265737320646f6573206e6f60448201527f742065786973742c206661696c656420746f206465637265617365207468652060648201526539ba30b5b29760d11b608482015260a4016101ef565b828160010154101561031b5760405162461bcd60e51b815260206004820152604560248201527f5b5145432d3033363030325d2d546865207374616b6520746f2072656d6f766560448201527f2069732067726561746572207468616e2074686520617661696c61626c6520736064820152643a30b5b29760d91b608482015260a4016101ef565b8281600101600082825461032f9190610bb7565b90915550506040518381526001600160a01b038516907f700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c699060200160405180910390a261037b846106b7565b60019150505b92915050565b6001600160a01b0381166000908152600160208190526040822001541515806103815750506001600160a01b0316600090815260016020526040902060020154151590565b6000336103d76104fd565b6001600160a01b0316146103fd5760405162461bcd60e51b81526004016101ef90610b6c565b6001600160a01b038316600090815260016020526040902054158015610421575081155b1561042e57506001610381565b6001600160a01b03831660008181526001602052604090819020600201849055517fc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc9109061047e9085815260200190565b60405180910390a261048f836106b7565b50600192915050565b336104a16104fd565b6001600160a01b0316146104c75760405162461bcd60e51b81526004016101ef90610b6c565b6104d160006106d6565b565b600281815481106104e357600080fd5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b031690565b6060600280548060200260200160405190810160405280929190818152602001828054801561056457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610546575b5050505050905090565b336105776104fd565b6001600160a01b03161461059d5760405162461bcd60e51b81526004016101ef90610b6c565b6001600160a01b0381166106025760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101ef565b61060b816106d6565b50565b6000336106196104fd565b6001600160a01b03161461063f5760405162461bcd60e51b81526004016101ef90610b6c565b8161064c57506000610381565b6001600160a01b03831660009081526001602081905260408220018054849290610677908490610bce565b90915550506040518281526001600160a01b038416907f8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d9060200161047e565b6106c081610387565b6106cd5761060b81610726565b61060b816108f2565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116600090815260016020819052604082205461074b9190610bb7565b60025490915060009061076090600190610bb7565b905080821461081f57610774826001610bce565b600160006002848154811061078b5761078b610be6565b60009182526020808320909101546001600160a01b0316835282019290925260400190205560028054829081106107c4576107c4610be6565b600091825260209091200154600280546001600160a01b0390921691849081106107f0576107f0610be6565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600280548061083057610830610bfc565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03851682526001908190526040822082815590810182905560020155610882836109a5565b6040516001600160a01b038416907f24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb90600090a28082146108ed576108ed600283815481106108d3576108d3610be6565b6000918252602090912001546001600160a01b03166109a5565b505050565b6001600160a01b03811660009081526001602052604090205461060b576002805460018101825560009182527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b03841690811790915560405190917fa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f91a26002546001600160a01b03821660009081526001602052604090205561060b815b6001600160a01b0381166000908152600160205260409020546002548111156109d0576109d0610c12565b6109d982610387565b15610a4257600081116109ee576109ee610c12565b6001600160a01b03821660009081526001602081905260409091200154151580610a3257506001600160a01b03821660009081526001602052604090206002015415155b610a3e57610a3e610c12565b5050565b8015610a5057610a50610c12565b6001600160a01b03821660009081526001602081905260409091200154158015610a3257506001600160a01b03821660009081526001602052604090206002015415610a3e57610a3e610c12565b80356001600160a01b0381168114610ab557600080fd5b919050565b60008060408385031215610acd57600080fd5b610ad683610a9e565b946020939093013593505050565b600060208284031215610af657600080fd5b610aff82610a9e565b9392505050565b600060208284031215610b1857600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015610b605783516001600160a01b031683529284019291840191600101610b3b565b50909695505050505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082821015610bc957610bc9610ba1565b500390565b60008219821115610be157610be1610ba1565b500190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea2646970667358221220fd26de599087406d0da509094082bcdc828350cc3635085addf4464741fd0f5764736f6c63430008090033",
}

// AddressStorageStakesABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageStakesMetaData.ABI instead.
var AddressStorageStakesABI = AddressStorageStakesMetaData.ABI

// AddressStorageStakesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageStakesMetaData.Bin instead.
var AddressStorageStakesBin = AddressStorageStakesMetaData.Bin

// DeployAddressStorageStakes deploys a new Ethereum contract, binding an instance of AddressStorageStakes to it.
func DeployAddressStorageStakes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorageStakes, error) {
	parsed, err := AddressStorageStakesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageStakesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorageStakes{AddressStorageStakesCaller: AddressStorageStakesCaller{contract: contract}, AddressStorageStakesTransactor: AddressStorageStakesTransactor{contract: contract}, AddressStorageStakesFilterer: AddressStorageStakesFilterer{contract: contract}}, nil
}

// AddressStorageStakes is an auto generated Go binding around an Ethereum contract.
type AddressStorageStakes struct {
	AddressStorageStakesCaller     // Read-only binding to the contract
	AddressStorageStakesTransactor // Write-only binding to the contract
	AddressStorageStakesFilterer   // Log filterer for contract events
}

// AddressStorageStakesCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageStakesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageStakesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageStakesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageStakesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageStakesSession struct {
	Contract     *AddressStorageStakes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressStorageStakesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageStakesCallerSession struct {
	Contract *AddressStorageStakesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AddressStorageStakesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageStakesTransactorSession struct {
	Contract     *AddressStorageStakesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AddressStorageStakesRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageStakesRaw struct {
	Contract *AddressStorageStakes // Generic contract binding to access the raw methods on
}

// AddressStorageStakesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageStakesCallerRaw struct {
	Contract *AddressStorageStakesCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageStakesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageStakesTransactorRaw struct {
	Contract *AddressStorageStakesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorageStakes creates a new instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakes(address common.Address, backend bind.ContractBackend) (*AddressStorageStakes, error) {
	contract, err := bindAddressStorageStakes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakes{AddressStorageStakesCaller: AddressStorageStakesCaller{contract: contract}, AddressStorageStakesTransactor: AddressStorageStakesTransactor{contract: contract}, AddressStorageStakesFilterer: AddressStorageStakesFilterer{contract: contract}}, nil
}

// NewAddressStorageStakesCaller creates a new read-only instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageStakesCaller, error) {
	contract, err := bindAddressStorageStakes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesCaller{contract: contract}, nil
}

// NewAddressStorageStakesTransactor creates a new write-only instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageStakesTransactor, error) {
	contract, err := bindAddressStorageStakes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesTransactor{contract: contract}, nil
}

// NewAddressStorageStakesFilterer creates a new log filterer instance of AddressStorageStakes, bound to a specific deployed contract.
func NewAddressStorageStakesFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageStakesFilterer, error) {
	contract, err := bindAddressStorageStakes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesFilterer{contract: contract}, nil
}

// bindAddressStorageStakes binds a generic wrapper to an already deployed contract.
func bindAddressStorageStakes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressStorageStakesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakes *AddressStorageStakesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakes.Contract.AddressStorageStakesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakes *AddressStorageStakesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.AddressStorageStakesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakes *AddressStorageStakesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.AddressStorageStakesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorageStakes *AddressStorageStakesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorageStakes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorageStakes *AddressStorageStakesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorageStakes *AddressStorageStakesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.contract.Transact(opts, method, params...)
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCaller) AddrList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "addrList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _AddressStorageStakes.Contract.AddrList(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrList is a free data retrieval call binding the contract method 0x80402564.
//
// Solidity: function addrList(uint256 ) view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) AddrList(arg0 *big.Int) (common.Address, error) {
	return _AddressStorageStakes.Contract.AddrList(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesCaller) AddrStake(opts *bind.CallOpts, arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "addrStake", arg0)

	outstruct := new(struct {
		ListIndex      *big.Int
		Stake          *big.Int
		DelegatedStake *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelegatedStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesSession) AddrStake(arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	return _AddressStorageStakes.Contract.AddrStake(&_AddressStorageStakes.CallOpts, arg0)
}

// AddrStake is a free data retrieval call binding the contract method 0xd9034503.
//
// Solidity: function addrStake(address ) view returns(uint256 listIndex, uint256 stake, uint256 delegatedStake)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) AddrStake(arg0 common.Address) (struct {
	ListIndex      *big.Int
	Stake          *big.Int
	DelegatedStake *big.Int
}, error) {
	return _AddressStorageStakes.Contract.AddrStake(&_AddressStorageStakes.CallOpts, arg0)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesCaller) Contains(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "contains", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakes.Contract.Contains(&_AddressStorageStakes.CallOpts, _addr)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorageStakes.Contract.Contains(&_AddressStorageStakes.CallOpts, _addr)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakes.Contract.GetAddresses(&_AddressStorageStakes.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorageStakes *AddressStorageStakesCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorageStakes.Contract.GetAddresses(&_AddressStorageStakes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesSession) Owner() (common.Address, error) {
	return _AddressStorageStakes.Contract.Owner(&_AddressStorageStakes.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Owner() (common.Address, error) {
	return _AddressStorageStakes.Contract.Owner(&_AddressStorageStakes.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorageStakes.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesSession) Size() (*big.Int, error) {
	return _AddressStorageStakes.Contract.Size(&_AddressStorageStakes.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorageStakes *AddressStorageStakesCallerSession) Size() (*big.Int, error) {
	return _AddressStorageStakes.Contract.Size(&_AddressStorageStakes.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) Add(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "add", _addr, _stake)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Add(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Add(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) Add(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Add(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.RenounceOwnership(&_AddressStorageStakes.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.RenounceOwnership(&_AddressStorageStakes.TransactOpts)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) SetDelegated(opts *bind.TransactOpts, _addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "setDelegated", _addr, _delegatedStake)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) SetDelegated(_addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.SetDelegated(&_AddressStorageStakes.TransactOpts, _addr, _delegatedStake)
}

// SetDelegated is a paid mutator transaction binding the contract method 0x6467b404.
//
// Solidity: function setDelegated(address _addr, uint256 _delegatedStake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) SetDelegated(_addr common.Address, _delegatedStake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.SetDelegated(&_AddressStorageStakes.TransactOpts, _addr, _delegatedStake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactor) Sub(opts *bind.TransactOpts, _addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "sub", _addr, _stake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesSession) Sub(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Sub(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address _addr, uint256 _stake) returns(bool)
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) Sub(_addr common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.Sub(&_AddressStorageStakes.TransactOpts, _addr, _stake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.TransferOwnership(&_AddressStorageStakes.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorageStakes *AddressStorageStakesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorageStakes.Contract.TransferOwnership(&_AddressStorageStakes.TransactOpts, newOwner)
}

// AddressStorageStakesAddressAddedIterator is returned from FilterAddressAdded and is used to iterate over the raw logs and unpacked data for AddressAdded events raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressAddedIterator struct {
	Event *AddressStorageStakesAddressAdded // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesAddressAdded)
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
		it.Event = new(AddressStorageStakesAddressAdded)
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
func (it *AddressStorageStakesAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesAddressAdded represents a AddressAdded event raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressAdded is a free log retrieval operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterAddressAdded(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesAddressAddedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesAddressAddedIterator{contract: _AddressStorageStakes.contract, event: "AddressAdded", logs: logs, sub: sub}, nil
}

// WatchAddressAdded is a free log subscription operation binding the contract event 0xa226db3f664042183ee0281230bba26cbf7b5057e50aee7f25a175ff45ce4d7f.
//
// Solidity: event AddressAdded(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchAddressAdded(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesAddressAdded, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "AddressAdded", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesAddressAdded)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressAdded", log); err != nil {
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
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseAddressAdded(log types.Log) (*AddressStorageStakesAddressAdded, error) {
	event := new(AddressStorageStakesAddressAdded)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesAddressRemovedIterator is returned from FilterAddressRemoved and is used to iterate over the raw logs and unpacked data for AddressRemoved events raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressRemovedIterator struct {
	Event *AddressStorageStakesAddressRemoved // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesAddressRemoved)
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
		it.Event = new(AddressStorageStakesAddressRemoved)
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
func (it *AddressStorageStakesAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesAddressRemoved represents a AddressRemoved event raised by the AddressStorageStakes contract.
type AddressStorageStakesAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddressRemoved is a free log retrieval operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterAddressRemoved(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesAddressRemovedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesAddressRemovedIterator{contract: _AddressStorageStakes.contract, event: "AddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAddressRemoved is a free log subscription operation binding the contract event 0x24a12366c02e13fe4a9e03d86a8952e85bb74a456c16e4a18b6d8295700b74bb.
//
// Solidity: event AddressRemoved(address indexed _addr)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchAddressRemoved(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesAddressRemoved, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "AddressRemoved", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesAddressRemoved)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
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
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseAddressRemoved(log types.Log) (*AddressStorageStakesAddressRemoved, error) {
	event := new(AddressStorageStakesAddressRemoved)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "AddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesDelegatedStakeUpdatedIterator is returned from FilterDelegatedStakeUpdated and is used to iterate over the raw logs and unpacked data for DelegatedStakeUpdated events raised by the AddressStorageStakes contract.
type AddressStorageStakesDelegatedStakeUpdatedIterator struct {
	Event *AddressStorageStakesDelegatedStakeUpdated // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesDelegatedStakeUpdated)
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
		it.Event = new(AddressStorageStakesDelegatedStakeUpdated)
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
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesDelegatedStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesDelegatedStakeUpdated represents a DelegatedStakeUpdated event raised by the AddressStorageStakes contract.
type AddressStorageStakesDelegatedStakeUpdated struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDelegatedStakeUpdated is a free log retrieval operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterDelegatedStakeUpdated(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesDelegatedStakeUpdatedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "DelegatedStakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesDelegatedStakeUpdatedIterator{contract: _AddressStorageStakes.contract, event: "DelegatedStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegatedStakeUpdated is a free log subscription operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchDelegatedStakeUpdated(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesDelegatedStakeUpdated, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "DelegatedStakeUpdated", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesDelegatedStakeUpdated)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "DelegatedStakeUpdated", log); err != nil {
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

// ParseDelegatedStakeUpdated is a log parse operation binding the contract event 0xc85c322a22d4a5cd3d9611287b2abf5bc9e6426d3ed5dc8d77017a7fe8fdc910.
//
// Solidity: event DelegatedStakeUpdated(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseDelegatedStakeUpdated(log types.Log) (*AddressStorageStakesDelegatedStakeUpdated, error) {
	event := new(AddressStorageStakesDelegatedStakeUpdated)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "DelegatedStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressStorageStakes contract.
type AddressStorageStakesOwnershipTransferredIterator struct {
	Event *AddressStorageStakesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesOwnershipTransferred)
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
		it.Event = new(AddressStorageStakesOwnershipTransferred)
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
func (it *AddressStorageStakesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesOwnershipTransferred represents a OwnershipTransferred event raised by the AddressStorageStakes contract.
type AddressStorageStakesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressStorageStakesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesOwnershipTransferredIterator{contract: _AddressStorageStakes.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesOwnershipTransferred)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseOwnershipTransferred(log types.Log) (*AddressStorageStakesOwnershipTransferred, error) {
	event := new(AddressStorageStakesOwnershipTransferred)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesStakeDecreasedIterator is returned from FilterStakeDecreased and is used to iterate over the raw logs and unpacked data for StakeDecreased events raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeDecreasedIterator struct {
	Event *AddressStorageStakesStakeDecreased // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesStakeDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesStakeDecreased)
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
		it.Event = new(AddressStorageStakesStakeDecreased)
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
func (it *AddressStorageStakesStakeDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesStakeDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesStakeDecreased represents a StakeDecreased event raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeDecreased struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeDecreased is a free log retrieval operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterStakeDecreased(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesStakeDecreasedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "StakeDecreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesStakeDecreasedIterator{contract: _AddressStorageStakes.contract, event: "StakeDecreased", logs: logs, sub: sub}, nil
}

// WatchStakeDecreased is a free log subscription operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchStakeDecreased(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesStakeDecreased, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "StakeDecreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesStakeDecreased)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeDecreased", log); err != nil {
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

// ParseStakeDecreased is a log parse operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseStakeDecreased(log types.Log) (*AddressStorageStakesStakeDecreased, error) {
	event := new(AddressStorageStakesStakeDecreased)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageStakesStakeIncreasedIterator is returned from FilterStakeIncreased and is used to iterate over the raw logs and unpacked data for StakeIncreased events raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeIncreasedIterator struct {
	Event *AddressStorageStakesStakeIncreased // Event containing the contract specifics and raw log

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
func (it *AddressStorageStakesStakeIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageStakesStakeIncreased)
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
		it.Event = new(AddressStorageStakesStakeIncreased)
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
func (it *AddressStorageStakesStakeIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageStakesStakeIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageStakesStakeIncreased represents a StakeIncreased event raised by the AddressStorageStakes contract.
type AddressStorageStakesStakeIncreased struct {
	Addr  common.Address
	Stake *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeIncreased is a free log retrieval operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) FilterStakeIncreased(opts *bind.FilterOpts, _addr []common.Address) (*AddressStorageStakesStakeIncreasedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.FilterLogs(opts, "StakeIncreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageStakesStakeIncreasedIterator{contract: _AddressStorageStakes.contract, event: "StakeIncreased", logs: logs, sub: sub}, nil
}

// WatchStakeIncreased is a free log subscription operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) WatchStakeIncreased(opts *bind.WatchOpts, sink chan<- *AddressStorageStakesStakeIncreased, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _AddressStorageStakes.contract.WatchLogs(opts, "StakeIncreased", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageStakesStakeIncreased)
				if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
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

// ParseStakeIncreased is a log parse operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed _addr, uint256 _stake)
func (_AddressStorageStakes *AddressStorageStakesFilterer) ParseStakeIncreased(log types.Log) (*AddressStorageStakesStakeIncreased, error) {
	event := new(AddressStorageStakesStakeIncreased)
	if err := _AddressStorageStakes.contract.UnpackLog(event, "StakeIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
