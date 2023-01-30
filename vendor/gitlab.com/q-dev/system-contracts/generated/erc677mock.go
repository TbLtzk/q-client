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

// ERC677MockMetaData contains all meta data concerning the ERC677Mock contract.
var ERC677MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200114a3803806200114a8339810160408190526200003491620002ed565b818181818160039080519060200190620000509291906200017a565b508051620000669060049060208401906200017a565b50505050506200008a336c7e37be2022c0914b26800000006200009260201b60201c565b5050620003bb565b6001600160a01b038216620000ed5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b806002600082825462000101919062000357565b90915550506001600160a01b038216600090815260208190526040812080548392906200013090849062000357565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b82805462000188906200037e565b90600052602060002090601f016020900481019282620001ac5760008555620001f7565b82601f10620001c757805160ff1916838001178555620001f7565b82800160010185558215620001f7579182015b82811115620001f7578251825591602001919060010190620001da565b506200020592915062000209565b5090565b5b808211156200020557600081556001016200020a565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200024857600080fd5b81516001600160401b038082111562000265576200026562000220565b604051601f8301601f19908116603f0116810190828211818310171562000290576200029062000220565b81604052838152602092508683858801011115620002ad57600080fd5b600091505b83821015620002d15785820183015181830184015290820190620002b2565b83821115620002e35760008385830101525b9695505050505050565b600080604083850312156200030157600080fd5b82516001600160401b03808211156200031957600080fd5b620003278683870162000236565b935060208501519150808211156200033e57600080fd5b506200034d8582860162000236565b9150509250929050565b600082198211156200037957634e487b7160e01b600052601160045260246000fd5b500190565b600181811c908216806200039357607f821691505b60208210811415620003b557634e487b7160e01b600052602260045260246000fd5b50919050565b610d7f80620003cb6000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c806306fdde03146100bf578063095ea7b3146100dd57806318160ddd1461010057806323b872dd14610112578063313ce5671461012557806339509351146101345780634000aea01461014757806342966c681461015a57806370a082311461016f57806379cc67901461019857806395d89b41146101ab578063a457c2d7146101b3578063a9059cbb146101c6578063dd62ed3e146101d9575b600080fd5b6100c7610212565b6040516100d49190610a79565b60405180910390f35b6100f06100eb366004610aa8565b6102a4565b60405190151581526020016100d4565b6002545b6040519081526020016100d4565b6100f0610120366004610ad2565b6102ba565b604051601281526020016100d4565b6100f0610142366004610aa8565b61036b565b6100f0610155366004610b24565b6103a7565b61016d610168366004610bef565b6104a3565b005b61010461017d366004610c08565b6001600160a01b031660009081526020819052604090205490565b61016d6101a6366004610aa8565b6104b0565b6100c7610536565b6100f06101c1366004610aa8565b610545565b6100f06101d4366004610aa8565b6105de565b6101046101e7366004610c23565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60606003805461022190610c56565b80601f016020809104026020016040519081016040528092919081815260200182805461024d90610c56565b801561029a5780601f1061026f5761010080835404028352916020019161029a565b820191906000526020600020905b81548152906001019060200180831161027d57829003601f168201915b5050505050905090565b60006102b13384846105eb565b50600192915050565b60006102c784848461070f565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156103515760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b61035e85338584036105eb565b60019150505b9392505050565b3360008181526001602090815260408083206001600160a01b038716845290915281205490916102b19185906103a2908690610ca7565b6105eb565b6000806103b485856105de565b9050806103c5576000915050610364565b6001600160a01b038516336001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16868660405161040a929190610cbf565b60405180910390a360405163607705c560e11b815285906001600160a01b0382169063c0ee0b8a9061044490339089908990600401610ce0565b602060405180830381600087803b15801561045e57600080fd5b505af1158015610472573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104969190610d10565b5060019695505050505050565b6104ad33826108de565b50565b60006104bc83336101e7565b90508181101561051a5760405162461bcd60e51b8152602060048201526024808201527f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f77604482015263616e636560e01b6064820152608401610348565b61052783338484036105eb565b61053183836108de565b505050565b60606004805461022190610c56565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156105c75760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610348565b6105d433858584036105eb565b5060019392505050565b60006102b133848461070f565b6001600160a01b03831661064d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610348565b6001600160a01b0382166106ae5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610348565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166107735760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610348565b6001600160a01b0382166107d55760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610348565b6001600160a01b0383166000908152602081905260409020548181101561084d5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610348565b6001600160a01b03808516600090815260208190526040808220858503905591851681529081208054849290610884908490610ca7565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108d091815260200190565b60405180910390a350505050565b6001600160a01b03821661093e5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610348565b6001600160a01b038216600090815260208190526040902054818110156109b25760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610348565b6001600160a01b03831660009081526020819052604081208383039055600280548492906109e1908490610d32565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3505050565b6000815180845260005b81811015610a5257602081850181015186830182015201610a36565b81811115610a64576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006103646020830184610a2c565b80356001600160a01b0381168114610aa357600080fd5b919050565b60008060408385031215610abb57600080fd5b610ac483610a8c565b946020939093013593505050565b600080600060608486031215610ae757600080fd5b610af084610a8c565b9250610afe60208501610a8c565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b600080600060608486031215610b3957600080fd5b610b4284610a8c565b925060208401359150604084013567ffffffffffffffff80821115610b6657600080fd5b818601915086601f830112610b7a57600080fd5b813581811115610b8c57610b8c610b0e565b604051601f8201601f19908116603f01168101908382118183101715610bb457610bb4610b0e565b81604052828152896020848701011115610bcd57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b600060208284031215610c0157600080fd5b5035919050565b600060208284031215610c1a57600080fd5b61036482610a8c565b60008060408385031215610c3657600080fd5b610c3f83610a8c565b9150610c4d60208401610a8c565b90509250929050565b600181811c90821680610c6a57607f821691505b60208210811415610c8b57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610cba57610cba610c91565b500190565b828152604060208201526000610cd86040830184610a2c565b949350505050565b60018060a01b0384168152826020820152606060408201526000610d076060830184610a2c565b95945050505050565b600060208284031215610d2257600080fd5b8151801515811461036457600080fd5b600082821015610d4457610d44610c91565b50039056fea2646970667358221220ff2ce118078c051ddaa88df70568102fc8d17b201586ddb5d66914f554d8b42364736f6c63430008090033",
}

// ERC677MockABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC677MockMetaData.ABI instead.
var ERC677MockABI = ERC677MockMetaData.ABI

// ERC677MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC677MockMetaData.Bin instead.
var ERC677MockBin = ERC677MockMetaData.Bin

// DeployERC677Mock deploys a new Ethereum contract, binding an instance of ERC677Mock to it.
func DeployERC677Mock(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *ERC677Mock, error) {
	parsed, err := ERC677MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC677MockBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC677Mock{ERC677MockCaller: ERC677MockCaller{contract: contract}, ERC677MockTransactor: ERC677MockTransactor{contract: contract}, ERC677MockFilterer: ERC677MockFilterer{contract: contract}}, nil
}

// ERC677Mock is an auto generated Go binding around an Ethereum contract.
type ERC677Mock struct {
	ERC677MockCaller     // Read-only binding to the contract
	ERC677MockTransactor // Write-only binding to the contract
	ERC677MockFilterer   // Log filterer for contract events
}

// ERC677MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC677MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC677MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC677MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC677MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC677MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC677MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC677MockSession struct {
	Contract     *ERC677Mock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC677MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC677MockCallerSession struct {
	Contract *ERC677MockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC677MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC677MockTransactorSession struct {
	Contract     *ERC677MockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC677MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC677MockRaw struct {
	Contract *ERC677Mock // Generic contract binding to access the raw methods on
}

// ERC677MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC677MockCallerRaw struct {
	Contract *ERC677MockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC677MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC677MockTransactorRaw struct {
	Contract *ERC677MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC677Mock creates a new instance of ERC677Mock, bound to a specific deployed contract.
func NewERC677Mock(address common.Address, backend bind.ContractBackend) (*ERC677Mock, error) {
	contract, err := bindERC677Mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC677Mock{ERC677MockCaller: ERC677MockCaller{contract: contract}, ERC677MockTransactor: ERC677MockTransactor{contract: contract}, ERC677MockFilterer: ERC677MockFilterer{contract: contract}}, nil
}

// NewERC677MockCaller creates a new read-only instance of ERC677Mock, bound to a specific deployed contract.
func NewERC677MockCaller(address common.Address, caller bind.ContractCaller) (*ERC677MockCaller, error) {
	contract, err := bindERC677Mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC677MockCaller{contract: contract}, nil
}

// NewERC677MockTransactor creates a new write-only instance of ERC677Mock, bound to a specific deployed contract.
func NewERC677MockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC677MockTransactor, error) {
	contract, err := bindERC677Mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC677MockTransactor{contract: contract}, nil
}

// NewERC677MockFilterer creates a new log filterer instance of ERC677Mock, bound to a specific deployed contract.
func NewERC677MockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC677MockFilterer, error) {
	contract, err := bindERC677Mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC677MockFilterer{contract: contract}, nil
}

// bindERC677Mock binds a generic wrapper to an already deployed contract.
func bindERC677Mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC677MockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC677Mock *ERC677MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC677Mock.Contract.ERC677MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC677Mock *ERC677MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC677Mock.Contract.ERC677MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC677Mock *ERC677MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC677Mock.Contract.ERC677MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC677Mock *ERC677MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC677Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC677Mock *ERC677MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC677Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC677Mock *ERC677MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC677Mock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC677Mock *ERC677MockCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC677Mock *ERC677MockSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC677Mock.Contract.Allowance(&_ERC677Mock.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC677Mock *ERC677MockCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC677Mock.Contract.Allowance(&_ERC677Mock.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC677Mock *ERC677MockCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC677Mock *ERC677MockSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC677Mock.Contract.BalanceOf(&_ERC677Mock.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC677Mock *ERC677MockCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC677Mock.Contract.BalanceOf(&_ERC677Mock.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC677Mock *ERC677MockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC677Mock *ERC677MockSession) Decimals() (uint8, error) {
	return _ERC677Mock.Contract.Decimals(&_ERC677Mock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC677Mock *ERC677MockCallerSession) Decimals() (uint8, error) {
	return _ERC677Mock.Contract.Decimals(&_ERC677Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC677Mock *ERC677MockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC677Mock *ERC677MockSession) Name() (string, error) {
	return _ERC677Mock.Contract.Name(&_ERC677Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC677Mock *ERC677MockCallerSession) Name() (string, error) {
	return _ERC677Mock.Contract.Name(&_ERC677Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC677Mock *ERC677MockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC677Mock *ERC677MockSession) Symbol() (string, error) {
	return _ERC677Mock.Contract.Symbol(&_ERC677Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC677Mock *ERC677MockCallerSession) Symbol() (string, error) {
	return _ERC677Mock.Contract.Symbol(&_ERC677Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC677Mock *ERC677MockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC677Mock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC677Mock *ERC677MockSession) TotalSupply() (*big.Int, error) {
	return _ERC677Mock.Contract.TotalSupply(&_ERC677Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC677Mock *ERC677MockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC677Mock.Contract.TotalSupply(&_ERC677Mock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Approve(&_ERC677Mock.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Approve(&_ERC677Mock.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC677Mock *ERC677MockTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC677Mock *ERC677MockSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Burn(&_ERC677Mock.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC677Mock *ERC677MockTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Burn(&_ERC677Mock.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC677Mock *ERC677MockTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC677Mock *ERC677MockSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.BurnFrom(&_ERC677Mock.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC677Mock *ERC677MockTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.BurnFrom(&_ERC677Mock.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC677Mock *ERC677MockSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.DecreaseAllowance(&_ERC677Mock.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.DecreaseAllowance(&_ERC677Mock.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC677Mock *ERC677MockSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.IncreaseAllowance(&_ERC677Mock.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.IncreaseAllowance(&_ERC677Mock.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Transfer(&_ERC677Mock.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.Transfer(&_ERC677Mock.TransactOpts, recipient, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) TransferAndCall(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "transferAndCall", _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_ERC677Mock *ERC677MockSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC677Mock.Contract.TransferAndCall(&_ERC677Mock.TransactOpts, _to, _value, _data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address _to, uint256 _value, bytes _data) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) TransferAndCall(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC677Mock.Contract.TransferAndCall(&_ERC677Mock.TransactOpts, _to, _value, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.TransferFrom(&_ERC677Mock.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC677Mock *ERC677MockTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC677Mock.Contract.TransferFrom(&_ERC677Mock.TransactOpts, sender, recipient, amount)
}

// ERC677MockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC677Mock contract.
type ERC677MockApprovalIterator struct {
	Event *ERC677MockApproval // Event containing the contract specifics and raw log

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
func (it *ERC677MockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC677MockApproval)
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
		it.Event = new(ERC677MockApproval)
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
func (it *ERC677MockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC677MockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC677MockApproval represents a Approval event raised by the ERC677Mock contract.
type ERC677MockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC677MockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC677Mock.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC677MockApprovalIterator{contract: _ERC677Mock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC677MockApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC677Mock.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC677MockApproval)
				if err := _ERC677Mock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) ParseApproval(log types.Log) (*ERC677MockApproval, error) {
	event := new(ERC677MockApproval)
	if err := _ERC677Mock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC677MockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC677Mock contract.
type ERC677MockTransferIterator struct {
	Event *ERC677MockTransfer // Event containing the contract specifics and raw log

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
func (it *ERC677MockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC677MockTransfer)
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
		it.Event = new(ERC677MockTransfer)
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
func (it *ERC677MockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC677MockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC677MockTransfer represents a Transfer event raised by the ERC677Mock contract.
type ERC677MockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_ERC677Mock *ERC677MockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC677MockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC677Mock.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC677MockTransferIterator{contract: _ERC677Mock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_ERC677Mock *ERC677MockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC677MockTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC677Mock.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC677MockTransfer)
				if err := _ERC677Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_ERC677Mock *ERC677MockFilterer) ParseTransfer(log types.Log) (*ERC677MockTransfer, error) {
	event := new(ERC677MockTransfer)
	if err := _ERC677Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC677MockTransfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the ERC677Mock contract.
type ERC677MockTransfer0Iterator struct {
	Event *ERC677MockTransfer0 // Event containing the contract specifics and raw log

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
func (it *ERC677MockTransfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC677MockTransfer0)
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
		it.Event = new(ERC677MockTransfer0)
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
func (it *ERC677MockTransfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC677MockTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC677MockTransfer0 represents a Transfer0 event raised by the ERC677Mock contract.
type ERC677MockTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC677MockTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC677Mock.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC677MockTransfer0Iterator{contract: _ERC677Mock.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *ERC677MockTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC677Mock.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC677MockTransfer0)
				if err := _ERC677Mock.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC677Mock *ERC677MockFilterer) ParseTransfer0(log types.Log) (*ERC677MockTransfer0, error) {
	event := new(ERC677MockTransfer0)
	if err := _ERC677Mock.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
