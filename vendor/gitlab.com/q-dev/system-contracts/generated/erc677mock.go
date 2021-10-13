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
	Bin: "0x60806040523480156200001157600080fd5b506040516200132038038062001320833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604052505050818181818160039080519060200190620001bb9291906200037c565b508051620001d19060049060208401906200037c565b50506005805460ff1916601217905550620001fe91503390506c7e37be2022c0914b268000000062000206565b505062000428565b6001600160a01b03821662000262576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b620002706000838362000315565b6200028c816002546200031a60201b620008d21790919060201c565b6002556001600160a01b03821660009081526020818152604090912054620002bf918390620008d26200031a821b17901c565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b505050565b60008282018381101562000375576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282620003b45760008555620003ff565b82601f10620003cf57805160ff1916838001178555620003ff565b82800160010185558215620003ff579182015b82811115620003ff578251825591602001919060010190620003e2565b506200040d92915062000411565b5090565b5b808211156200040d576000815560010162000412565b610ee880620004386000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c806306fdde03146100bf578063095ea7b31461013c57806318160ddd1461017c57806323b872dd14610196578063313ce567146101cc57806339509351146101ea5780634000aea01461021657806342966c68146102cf57806370a08231146102ee57806379cc67901461031457806395d89b4114610340578063a457c2d714610348578063a9059cbb14610374578063dd62ed3e146103a0575b600080fd5b6100c76103ce565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101015781810151838201526020016100e9565b50505050905090810190601f16801561012e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101686004803603604081101561015257600080fd5b506001600160a01b038135169060200135610464565b604080519115158252519081900360200190f35b610184610481565b60408051918252519081900360200190f35b610168600480360360608110156101ac57600080fd5b506001600160a01b03813581169160208101359091169060400135610487565b6101d461050f565b6040805160ff9092168252519081900360200190f35b6101686004803603604081101561020057600080fd5b506001600160a01b038135169060200135610518565b6101686004803603606081101561022c57600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561025b57600080fd5b82018360208201111561026d57600080fd5b803590602001918460018302840111600160201b8311171561028e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610566945050505050565b6102ec600480360360208110156102e557600080fd5b5035610741565b005b6101846004803603602081101561030457600080fd5b50356001600160a01b0316610755565b6102ec6004803603604081101561032a57600080fd5b506001600160a01b038135169060200135610770565b6100c76107ca565b6101686004803603604081101561035e57600080fd5b506001600160a01b03813516906020013561082b565b6101686004803603604081101561038a57600080fd5b506001600160a01b038135169060200135610893565b610184600480360360408110156103b657600080fd5b506001600160a01b03813581169160200135166108a7565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045a5780601f1061042f5761010080835404028352916020019161045a565b820191906000526020600020905b81548152906001019060200180831161043d57829003601f168201915b5050505050905090565b600061047861047161092a565b848461092e565b50600192915050565b60025490565b6000610494848484610a1a565b610504846104a061092a565b6104ff85604051806060016040528060288152602001610dd8602891396001600160a01b038a166000908152600160205260408120906104de61092a565b6001600160a01b031681526020810191909152604001600020549190610b75565b61092e565b5060015b9392505050565b60055460ff1690565b600061047861052561092a565b846104ff856001600061053661092a565b6001600160a01b03908116825260208083019390935260409182016000908120918c1681529252902054906108d2565b6000806105738585610893565b905080610584576000915050610508565b846001600160a01b031661059661092a565b6001600160a01b03167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1686866040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106055781810151838201526020016105ed565b50505050905090810190601f1680156106325780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3846001600160a01b03811663c0ee0b8a61065861092a565b87876040518463ffffffff1660e01b815260040180846001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106bb5781810151838201526020016106a3565b50505050905090810190601f1680156106e85780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561070957600080fd5b505af115801561071d573d6000803e3d6000fd5b505050506040513d602081101561073357600080fd5b506001979650505050505050565b61075261074c61092a565b82610c0c565b50565b6001600160a01b031660009081526020819052604090205490565b60006107a782604051806060016040528060248152602001610e00602491396107a08661079b61092a565b6108a7565b9190610b75565b90506107bb836107b561092a565b8361092e565b6107c58383610c0c565b505050565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561045a5780601f1061042f5761010080835404028352916020019161045a565b600061047861083861092a565b846104ff85604051806060016040528060258152602001610e8e602591396001600061086261092a565b6001600160a01b03908116825260208083019390935260409182016000908120918d16815292529020549190610b75565b60006104786108a061092a565b8484610a1a565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b600082820183811015610508576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b3390565b6001600160a01b0383166109735760405162461bcd60e51b8152600401808060200182810382526024815260200180610e6a6024913960400191505060405180910390fd5b6001600160a01b0382166109b85760405162461bcd60e51b8152600401808060200182810382526022815260200180610d906022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610a5f5760405162461bcd60e51b8152600401808060200182810382526025815260200180610e456025913960400191505060405180910390fd5b6001600160a01b038216610aa45760405162461bcd60e51b8152600401808060200182810382526023815260200180610d4b6023913960400191505060405180910390fd5b610aaf8383836107c5565b610aec81604051806060016040528060268152602001610db2602691396001600160a01b0386166000908152602081905260409020549190610b75565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610b1b90826108d2565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610c045760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610bc9578181015183820152602001610bb1565b50505050905090810190601f168015610bf65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6001600160a01b038216610c515760405162461bcd60e51b8152600401808060200182810382526021815260200180610e246021913960400191505060405180910390fd5b610c5d826000836107c5565b610c9a81604051806060016040528060228152602001610d6e602291396001600160a01b0385166000908152602081905260409020549190610b75565b6001600160a01b038316600090815260208190526040902055600254610cc09082610d08565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600061050883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610b7556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212202c91eedc7f2f3c0db0db1efdf3dbb333a40dfd874fb21c8227c2b4838ed203a864736f6c63430007060033",
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
