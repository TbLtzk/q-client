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

// EPDRMembershipMetaData contains all meta data concerning the EPDRMembership contract.
var EPDRMembershipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertLimitKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertVotingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"experts\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldExp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newExp\",\"type\":\"address\"}],\"name\":\"swapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_experts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506115a0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630b1ca49a146100bf5780633e021814146100d45780635313263a146100f2578063715018a6146100fa5780638da5cb5b14610102578063946d9204146101175780639eab52531461012a578063a230c5241461013f578063b278982d14610162578063b295a00e14610175578063c68368251461018b578063ca6d56dc1461019e578063de8fa431146101b1578063f2fde38b146101b9575b600080fd5b6100d26100cd36600461102a565b6101cc565b005b6100dc61030c565b6040516100e9919061104e565b60405180910390f35b6100dc61039a565b6100d26103a7565b61010a6103e2565b6040516100e991906110a3565b6100d2610125366004611120565b6103f1565b610132610472565b6040516100e991906111d3565b61015261014d36600461102a565b6104f8565b60405190151581526020016100e9565b6100d2610170366004611220565b61057f565b61017d610739565b6040519081526020016100e9565b60655461010a906001600160a01b031681565b6100d26101ac36600461102a565b6108db565b61017d610af7565b6100d26101c736600461102a565b610b74565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb90271906101fd90606890600401611294565b60206040518083038186803b15801561021557600080fd5b505afa158015610229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024d919061133c565b6001600160a01b0316336001600160a01b0316146102865760405162461bcd60e51b815260040161027d90611359565b60405180910390fd5b606554604051631484968760e11b81526001600160a01b03909116906329092d0e906102b69084906004016110a3565b602060405180830381600087803b1580156102d057600080fd5b505af11580156102e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030891906113c9565b5050565b6067805461031990611259565b80601f016020809104026020016040519081016040528092919081815260200182805461034590611259565b80156103925780601f1061036757610100808354040283529160200191610392565b820191906000526020600020905b81548152906001019060200180831161037557829003601f168201915b505050505081565b6068805461031990611259565b336103b06103e2565b6001600160a01b0316146103d65760405162461bcd60e51b815260040161027d906113eb565b6103e06000610c14565b565b6033546001600160a01b031690565b6103fb8282610c66565b60408051808201909152601d8082527f636f6e737469747574696f6e2e455044522e6d61784e45787065727473000000602090920191825261043f91606791610f7c565b5060405180606001604052806028815260200161154360289139805161046d91606891602090910190610f7c565b505050565b606554604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b1580156104b757600080fd5b505afa1580156104cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104f39190810190611420565b905090565b606554604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e8906105299085906004016110a3565b60206040518083038186803b15801561054157600080fd5b505afa158015610555573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057991906113c9565b92915050565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb90271906105b090606890600401611294565b60206040518083038186803b1580156105c857600080fd5b505afa1580156105dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610600919061133c565b6001600160a01b0316336001600160a01b0316146106305760405162461bcd60e51b815260040161027d90611359565b606554604051631484968760e11b81526001600160a01b039091169081906329092d0e906106629086906004016110a3565b602060405180830381600087803b15801561067c57600080fd5b505af1158015610690573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b491906113c9565b50604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f906106e19085906004016110a3565b602060405180830381600087803b1580156106fb57600080fd5b505af115801561070f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073391906113c9565b50505050565b6066546040805160608101909152602280825260009283926001600160a01b0390911691633fb90271919061152160208301396040518263ffffffff1660e01b8152600401610788919061104e565b60206040518083038186803b1580156107a057600080fd5b505afa1580156107b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d8919061133c565b6001600160a01b031663498bff0060676040518263ffffffff1660e01b81526004016108049190611294565b60206040518083038186803b15801561081c57600080fd5b505afa158015610830573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085491906114b9565b9050600081116108d65760405162461bcd60e51b815260206004820152604160248201527f5b5145432d3030373030315d2d457870657274206c696d697420706172616d6560448201527f74657220697320696e76616c6964206f7220646f6573206e6f742065786973746064820152601760f91b608482015260a40161027d565b919050565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb902719061090c90606890600401611294565b60206040518083038186803b15801561092457600080fd5b505afa158015610938573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095c919061133c565b6001600160a01b0316336001600160a01b03161461098c5760405162461bcd60e51b815260040161027d90611359565b6065546001600160a01b03166109a0610739565b816001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d957600080fd5b505afa1580156109ed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1191906114b9565b10610a795760405162461bcd60e51b815260206004820152603260248201527f5b5145432d3030373030305d2d4d6178696d756d206e756d626572206f6620656044820152713c3832b93a399034b9903932b0b1b432b21760711b606482015260840161027d565b604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610aa59085906004016110a3565b602060405180830381600087803b158015610abf57600080fd5b505af1158015610ad3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046d91906113c9565b6065546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d916004808301926020929190829003018186803b158015610b3c57600080fd5b505afa158015610b50573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f391906114b9565b33610b7d6103e2565b6001600160a01b031614610ba35760405162461bcd60e51b815260040161027d906113eb565b6001600160a01b038116610c085760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161027d565b610c1181610c14565b50565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680610c7f575060005460ff16155b610c9b5760405162461bcd60e51b815260040161027d906114d2565b600054610100900460ff16158015610cbd576000805461ffff19166101011790555b606680546001600160a01b0319166001600160a01b038516908117909155604080518082018252601d81527f636f6d6d6f6e2e666163746f72792e6164647265737353746f7261676500000060208201529051633fb9027160e01b8152633fb9027191610d2c9160040161104e565b60206040518083038186803b158015610d4457600080fd5b505afa158015610d58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7c919061133c565b6001600160a01b03166362c54dbd836040518263ffffffff1660e01b8152600401610da791906111d3565b602060405180830381600087803b158015610dc157600080fd5b505af1158015610dd5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df9919061133c565b606580546001600160a01b0319166001600160a01b0392909216919091179055610e21610e37565b801561046d576000805461ff0019169055505050565b600054610100900460ff1680610e50575060005460ff16155b610e6c5760405162461bcd60e51b815260040161027d906114d2565b600054610100900460ff16158015610e8e576000805461ffff19166101011790555b610e96610eb2565b610e9e610f1c565b8015610c11576000805461ff001916905550565b600054610100900460ff1680610ecb575060005460ff16155b610ee75760405162461bcd60e51b815260040161027d906114d2565b600054610100900460ff16158015610e9e576000805461ffff19166101011790558015610c11576000805461ff001916905550565b600054610100900460ff1680610f35575060005460ff16155b610f515760405162461bcd60e51b815260040161027d906114d2565b600054610100900460ff16158015610f73576000805461ffff19166101011790555b610e9e33610c14565b828054610f8890611259565b90600052602060002090601f016020900481019282610faa5760008555610ff0565b82601f10610fc357805160ff1916838001178555610ff0565b82800160010185558215610ff0579182015b82811115610ff0578251825591602001919060010190610fd5565b50610ffc929150611000565b5090565b5b80821115610ffc5760008155600101611001565b6001600160a01b0381168114610c1157600080fd5b60006020828403121561103c57600080fd5b813561104781611015565b9392505050565b600060208083528351808285015260005b8181101561107b5785810183015185820160400152820161105f565b8181111561108d576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b0391909116815260200190565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156110f5576110f56110b7565b604052919050565b60006001600160401b03821115611116576111166110b7565b5060051b60200190565b6000806040838503121561113357600080fd5b823561113e81611015565b91506020838101356001600160401b0381111561115a57600080fd5b8401601f8101861361116b57600080fd5b803561117e611179826110fd565b6110cd565b81815260059190911b8201830190838101908883111561119d57600080fd5b928401925b828410156111c45783356111b581611015565b825292840192908401906111a2565b80955050505050509250929050565b6020808252825182820181905260009190848201906040850190845b818110156112145783516001600160a01b0316835292840192918401916001016111ef565b50909695505050505050565b6000806040838503121561123357600080fd5b823561123e81611015565b9150602083013561124e81611015565b809150509250929050565b600181811c9082168061126d57607f821691505b6020821081141561128e57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208083526000845481600182811c9150808316806112b657607f831692505b8583108114156112d457634e487b7160e01b85526022600452602485fd5b8786018381526020018180156112f157600181146113025761132d565b60ff1986168252878201965061132d565b60008b81526020902060005b868110156113275781548482015290850190890161130e565b83019750505b50949998505050505050505050565b60006020828403121561134e57600080fd5b815161104781611015565b6020808252604a908201527f5b5145432d3030373030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920657870657274732066726f6d20746869732070616e656c2068616060820152693b329030b1b1b2b9b99760b11b608082015260a00190565b6000602082840312156113db57600080fd5b8151801515811461104757600080fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000602080838503121561143357600080fd5b82516001600160401b0381111561144957600080fd5b8301601f8101851361145a57600080fd5b8051611468611179826110fd565b81815260059190911b8201830190838101908783111561148757600080fd5b928401925b828410156114ae57835161149f81611015565b8252928401929084019061148c565b979650505050505050565b6000602082840312156114cb57600080fd5b5051919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b60608201526080019056fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273676f7665726e616e63652e657870657274732e455044522e6d656d62657273686970566f74696e67a2646970667358221220f25cf16a65f3c4cfb054e47f88818aa7cb775f8d2326b969c7bf73bfac42ca4b64736f6c63430008090033",
}

// EPDRMembershipABI is the input ABI used to generate the binding from.
// Deprecated: Use EPDRMembershipMetaData.ABI instead.
var EPDRMembershipABI = EPDRMembershipMetaData.ABI

// EPDRMembershipBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EPDRMembershipMetaData.Bin instead.
var EPDRMembershipBin = EPDRMembershipMetaData.Bin

// DeployEPDRMembership deploys a new Ethereum contract, binding an instance of EPDRMembership to it.
func DeployEPDRMembership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EPDRMembership, error) {
	parsed, err := EPDRMembershipMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EPDRMembershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// EPDRMembership is an auto generated Go binding around an Ethereum contract.
type EPDRMembership struct {
	EPDRMembershipCaller     // Read-only binding to the contract
	EPDRMembershipTransactor // Write-only binding to the contract
	EPDRMembershipFilterer   // Log filterer for contract events
}

// EPDRMembershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPDRMembershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPDRMembershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPDRMembershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPDRMembershipSession struct {
	Contract     *EPDRMembership   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPDRMembershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPDRMembershipCallerSession struct {
	Contract *EPDRMembershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EPDRMembershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPDRMembershipTransactorSession struct {
	Contract     *EPDRMembershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EPDRMembershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPDRMembershipRaw struct {
	Contract *EPDRMembership // Generic contract binding to access the raw methods on
}

// EPDRMembershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPDRMembershipCallerRaw struct {
	Contract *EPDRMembershipCaller // Generic read-only contract binding to access the raw methods on
}

// EPDRMembershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPDRMembershipTransactorRaw struct {
	Contract *EPDRMembershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPDRMembership creates a new instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembership(address common.Address, backend bind.ContractBackend) (*EPDRMembership, error) {
	contract, err := bindEPDRMembership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPDRMembership{EPDRMembershipCaller: EPDRMembershipCaller{contract: contract}, EPDRMembershipTransactor: EPDRMembershipTransactor{contract: contract}, EPDRMembershipFilterer: EPDRMembershipFilterer{contract: contract}}, nil
}

// NewEPDRMembershipCaller creates a new read-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipCaller(address common.Address, caller bind.ContractCaller) (*EPDRMembershipCaller, error) {
	contract, err := bindEPDRMembership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipCaller{contract: contract}, nil
}

// NewEPDRMembershipTransactor creates a new write-only instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipTransactor(address common.Address, transactor bind.ContractTransactor) (*EPDRMembershipTransactor, error) {
	contract, err := bindEPDRMembership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipTransactor{contract: contract}, nil
}

// NewEPDRMembershipFilterer creates a new log filterer instance of EPDRMembership, bound to a specific deployed contract.
func NewEPDRMembershipFilterer(address common.Address, filterer bind.ContractFilterer) (*EPDRMembershipFilterer, error) {
	contract, err := bindEPDRMembership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipFilterer{contract: contract}, nil
}

// bindEPDRMembership binds a generic wrapper to an already deployed contract.
func bindEPDRMembership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EPDRMembershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.EPDRMembershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.EPDRMembershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPDRMembership *EPDRMembershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPDRMembership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPDRMembership *EPDRMembershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPDRMembership.Contract.contract.Transact(opts, method, params...)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertLimitKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "expertLimitKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertLimitKey() (string, error) {
	return _EPDRMembership.Contract.ExpertLimitKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCaller) ExpertVotingKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "expertVotingKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPDRMembership *EPDRMembershipCallerSession) ExpertVotingKey() (string, error) {
	return _EPDRMembership.Contract.ExpertVotingKey(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Experts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "experts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Experts() (common.Address, error) {
	return _EPDRMembership.Contract.Experts(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetLimit() (*big.Int, error) {
	return _EPDRMembership.Contract.GetLimit(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPDRMembership *EPDRMembershipCallerSession) GetMembers() ([]common.Address, error) {
	return _EPDRMembership.Contract.GetMembers(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "getSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPDRMembership *EPDRMembershipCallerSession) GetSize() (*big.Int, error) {
	return _EPDRMembership.Contract.GetSize(&_EPDRMembership.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCaller) IsMember(opts *bind.CallOpts, _exp common.Address) (bool, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "isMember", _exp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPDRMembership *EPDRMembershipCallerSession) IsMember(_exp common.Address) (bool, error) {
	return _EPDRMembership.Contract.IsMember(&_EPDRMembership.CallOpts, _exp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPDRMembership.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPDRMembership *EPDRMembershipCallerSession) Owner() (common.Address, error) {
	return _EPDRMembership.Contract.Owner(&_EPDRMembership.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) AddMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "addMember", _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.AddMember(&_EPDRMembership.TransactOpts, _exp)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "initialize", _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.Initialize(&_EPDRMembership.TransactOpts, _registry, _experts)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) RemoveMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "removeMember", _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.RemoveMember(&_EPDRMembership.TransactOpts, _exp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPDRMembership.Contract.RenounceOwnership(&_EPDRMembership.TransactOpts)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactor) SwapMember(opts *bind.TransactOpts, _oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "swapMember", _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.SwapMember(&_EPDRMembership.TransactOpts, _oldExp, _newExp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPDRMembership *EPDRMembershipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPDRMembership.Contract.TransferOwnership(&_EPDRMembership.TransactOpts, newOwner)
}

// EPDRMembershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferredIterator struct {
	Event *EPDRMembershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPDRMembershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPDRMembershipOwnershipTransferred)
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
		it.Event = new(EPDRMembershipOwnershipTransferred)
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
func (it *EPDRMembershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPDRMembershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPDRMembershipOwnershipTransferred represents a OwnershipTransferred event raised by the EPDRMembership contract.
type EPDRMembershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPDRMembershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPDRMembershipOwnershipTransferredIterator{contract: _EPDRMembership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPDRMembership *EPDRMembershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPDRMembershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPDRMembership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPDRMembershipOwnershipTransferred)
				if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EPDRMembership *EPDRMembershipFilterer) ParseOwnershipTransferred(log types.Log) (*EPDRMembershipOwnershipTransferred, error) {
	event := new(EPDRMembershipOwnershipTransferred)
	if err := _EPDRMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
