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

// EPQFIMembershipMetaData contains all meta data concerning the EPQFIMembership contract.
var EPQFIMembershipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertLimitKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expertVotingKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"experts\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exp\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldExp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newExp\",\"type\":\"address\"}],\"name\":\"swapMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_experts\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611755806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ba5760003560e01c80630b1ca49a146100bf5780633e021814146100d45780635313263a146100f2578063715018a6146100fa5780638da5cb5b14610102578063946d9204146101175780639eab52531461012a578063a230c5241461013f578063b278982d1461015f578063b295a00e14610172578063c683682514610187578063ca6d56dc1461018f578063de8fa431146101a2578063f2fde38b146101aa575b600080fd5b6100d26100cd36600461118b565b6101bd565b005b6100dc6102fd565b6040516100e991906113f0565b60405180910390f35b6100dc61038b565b6100d26103e6565b61010a610492565b6040516100e99190611384565b6100d2610125366004611202565b6104a1565b610132610522565b6040516100e99190611398565b61015261014d36600461118b565b6105a8565b6040516100e991906113e5565b6100d261016d3660046111ca565b61062f565b61017a6107e9565b6040516100e991906115f6565b61010a610926565b6100d261019d36600461118b565b610935565b61017a610b07565b6100d26101b836600461118b565b610b84565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb90271906101ee90606890600401611443565b60206040518083038186803b15801561020657600080fd5b505afa15801561021a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023e91906111ae565b6001600160a01b0316336001600160a01b0316146102775760405162461bcd60e51b815260040161026e90611586565b60405180910390fd5b606554604051631484968760e11b81526001600160a01b03909116906329092d0e906102a7908490600401611384565b602060405180830381600087803b1580156102c157600080fd5b505af11580156102d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f9919061134c565b5050565b6067805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103835780601f1061035857610100808354040283529160200191610383565b820191906000526020600020905b81548152906001019060200180831161036657829003601f168201915b505050505081565b6068805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103835780601f1061035857610100808354040283529160200191610383565b6103ee610c87565b6001600160a01b03166103ff610492565b6001600160a01b03161461045a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316906000805160206116b5833981519152908390a3603380546001600160a01b0319169055565b6033546001600160a01b031690565b6104ab8282610c8b565b60408051808201909152601e8082527f636f6e737469747574696f6e2e45505146492e6d61784e45787065727473000060209092019182526104ef916067916110ea565b506040518060600160405280602981526020016116f760299139805161051d916068916020909101906110ea565b505050565b606554604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac12916004808301926000929190829003018186803b15801561056757600080fd5b505afa15801561057b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105a391908101906112b4565b905090565b606554604051630bb7c8fd60e31b81526000916001600160a01b031690635dbe47e8906105d9908590600401611384565b60206040518083038186803b1580156105f157600080fd5b505afa158015610605573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610629919061134c565b92915050565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb902719061066090606890600401611443565b60206040518083038186803b15801561067857600080fd5b505afa15801561068c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b091906111ae565b6001600160a01b0316336001600160a01b0316146106e05760405162461bcd60e51b815260040161026e90611586565b606554604051631484968760e11b81526001600160a01b039091169081906329092d0e90610712908690600401611384565b602060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610764919061134c565b50604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610791908590600401611384565b602060405180830381600087803b1580156107ab57600080fd5b505af11580156107bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e3919061134c565b50505050565b6066546040805160608101909152602280825260009283926001600160a01b0390911691633fb9027191906116d560208301396040518263ffffffff1660e01b815260040161083891906113f0565b60206040518083038186803b15801561085057600080fd5b505afa158015610864573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088891906111ae565b6001600160a01b031663498bff0060676040518263ffffffff1660e01b81526004016108b49190611443565b60206040518083038186803b1580156108cc57600080fd5b505afa1580156108e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610904919061136c565b9050600081116105a35760405162461bcd60e51b815260040161026e9061151f565b6065546001600160a01b031681565b606654604051633fb9027160e01b81526001600160a01b0390911690633fb902719061096690606890600401611443565b60206040518083038186803b15801561097e57600080fd5b505afa158015610992573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b691906111ae565b6001600160a01b0316336001600160a01b0316146109e65760405162461bcd60e51b815260040161026e90611586565b6065546001600160a01b03166109fa6107e9565b816001600160a01b031663949d225d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3357600080fd5b505afa158015610a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6b919061136c565b1115610a895760405162461bcd60e51b815260040161026e906114cd565b604051630a3b0a4f60e01b81526001600160a01b03821690630a3b0a4f90610ab5908590600401611384565b602060405180830381600087803b158015610acf57600080fd5b505af1158015610ae3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051d919061134c565b6065546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d916004808301926020929190829003018186803b158015610b4c57600080fd5b505afa158015610b60573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a3919061136c565b610b8c610c87565b6001600160a01b0316610b9d610492565b6001600160a01b031614610bf8576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b038116610c3d5760405162461bcd60e51b81526004018080602001828103825260268152602001806116616026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216906000805160206116b583398151915290600090a3603380546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b600054610100900460ff1680610ca45750610ca4610e9a565b80610cb2575060005460ff16155b610ced5760405162461bcd60e51b815260040180806020018281038252602e815260200180611687602e913960400191505060405180910390fd5b600054610100900460ff16158015610d18576000805460ff1961ff0019909116610100171660011790555b606680546001600160a01b0319166001600160a01b038581169190911791829055604080518082018252601d81527f636f6d6d6f6e2e666163746f72792e6164647265737353746f7261676500000060208201529051633fb9027160e01b81529290911691633fb9027191610d8f916004016113f0565b60206040518083038186803b158015610da757600080fd5b505afa158015610dbb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ddf91906111ae565b6001600160a01b03166362c54dbd836040518263ffffffff1660e01b8152600401610e0a9190611398565b602060405180830381600087803b158015610e2457600080fd5b505af1158015610e38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5c91906111ae565b606580546001600160a01b0319166001600160a01b0392909216919091179055610e84610eab565b801561051d576000805461ff0019169055505050565b6000610ea530610f5d565b15905090565b600054610100900460ff1680610ec45750610ec4610e9a565b80610ed2575060005460ff16155b610f0d5760405162461bcd60e51b815260040180806020018281038252602e815260200180611687602e913960400191505060405180910390fd5b600054610100900460ff16158015610f38576000805460ff1961ff0019909116610100171660011790555b610f40610f63565b610f48611003565b8015610f5a576000805461ff00191690555b50565b3b151590565b600054610100900460ff1680610f7c5750610f7c610e9a565b80610f8a575060005460ff16155b610fc55760405162461bcd60e51b815260040180806020018281038252602e815260200180611687602e913960400191505060405180910390fd5b600054610100900460ff16158015610f48576000805460ff1961ff0019909116610100171660011790558015610f5a576000805461ff001916905550565b600054610100900460ff168061101c575061101c610e9a565b8061102a575060005460ff16155b6110655760405162461bcd60e51b815260040180806020018281038252602e815260200180611687602e913960400191505060405180910390fd5b600054610100900460ff16158015611090576000805460ff1961ff0019909116610100171660011790555b600061109a610c87565b603380546001600160a01b0319166001600160a01b038316908117909155604051919250906000906000805160206116b5833981519152908290a3508015610f5a576000805461ff001916905550565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826111205760008555611166565b82601f1061113957805160ff1916838001178555611166565b82800160010185558215611166579182015b8281111561116657825182559160200191906001019061114b565b50611172929150611176565b5090565b5b808211156111725760008155600101611177565b60006020828403121561119c578081fd5b81356111a78161164b565b9392505050565b6000602082840312156111bf578081fd5b81516111a78161164b565b600080604083850312156111dc578081fd5b82356111e78161164b565b915060208301356111f78161164b565b809150509250929050565b60008060408385031215611214578182fd5b823561121f8161164b565b91506020838101356001600160401b0381111561123a578283fd5b8401601f8101861361124a578283fd5b803561125d61125882611622565b6115ff565b81815283810190838501858402850186018a1015611279578687fd5b8694505b838510156112a45780356112908161164b565b83526001949094019391850191850161127d565b5080955050505050509250929050565b600060208083850312156112c6578182fd5b82516001600160401b038111156112db578283fd5b8301601f810185136112eb578283fd5b80516112f961125882611622565b8181528381019083850185840285018601891015611315578687fd5b8694505b8385101561134057805161132c8161164b565b835260019490940193918501918501611319565b50979650505050505050565b60006020828403121561135d578081fd5b815180151581146111a7578182fd5b60006020828403121561137d578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156113d95783516001600160a01b0316835292840192918401916001016113b4565b50909695505050505050565b901515815260200190565b6000602080835283518082850152825b8181101561141c57858101830151858201604001528201611400565b8181111561142d5783604083870101525b50601f01601f1916929092016040019392505050565b6000602080830181845282855460018082166000811461146a5760018114611488576114c0565b60028304607f16855260ff19831660408901526060880193506114c0565b600283048086526114988a61163f565b885b828110156114b65781548b82016040015290840190880161149a565b8a01604001955050505b5091979650505050505050565b60208082526032908201527f5b5145432d3030373030305d2d4d6178696d756d206e756d626572206f6620656040820152713c3832b93a399034b9903932b0b1b432b21760711b606082015260800190565b60208082526041908201527f5b5145432d3030373030315d2d457870657274206c696d697420706172616d6560408201527f74657220697320696e76616c6964206f7220646f6573206e6f742065786973746060820152601760f91b608082015260a00190565b6020808252604a908201527f5b5145432d3030373030325d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920657870657274732066726f6d20746869732070616e656c2068616060820152693b329030b1b1b2b9b99760b11b608082015260a00190565b90815260200190565b6040518181016001600160401b038111828210171561161a57fe5b604052919050565b60006001600160401b0382111561163557fe5b5060209081020190565b60009081526020902090565b6001600160a01b0381168114610f5a57600080fdfe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65648be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273676f7665726e616e63652e657870657274732e45505146492e6d656d62657273686970566f74696e67a2646970667358221220ebff49e5f98d1e2fef4979ebe856b0f37a64a3749171cc17f1b725bc34f9d62c64736f6c63430007060033",
}

// EPQFIMembershipABI is the input ABI used to generate the binding from.
// Deprecated: Use EPQFIMembershipMetaData.ABI instead.
var EPQFIMembershipABI = EPQFIMembershipMetaData.ABI

// EPQFIMembershipBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EPQFIMembershipMetaData.Bin instead.
var EPQFIMembershipBin = EPQFIMembershipMetaData.Bin

// DeployEPQFIMembership deploys a new Ethereum contract, binding an instance of EPQFIMembership to it.
func DeployEPQFIMembership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EPQFIMembership, error) {
	parsed, err := EPQFIMembershipMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EPQFIMembershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EPQFIMembership{EPQFIMembershipCaller: EPQFIMembershipCaller{contract: contract}, EPQFIMembershipTransactor: EPQFIMembershipTransactor{contract: contract}, EPQFIMembershipFilterer: EPQFIMembershipFilterer{contract: contract}}, nil
}

// EPQFIMembership is an auto generated Go binding around an Ethereum contract.
type EPQFIMembership struct {
	EPQFIMembershipCaller     // Read-only binding to the contract
	EPQFIMembershipTransactor // Write-only binding to the contract
	EPQFIMembershipFilterer   // Log filterer for contract events
}

// EPQFIMembershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type EPQFIMembershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EPQFIMembershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EPQFIMembershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EPQFIMembershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EPQFIMembershipSession struct {
	Contract     *EPQFIMembership  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EPQFIMembershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EPQFIMembershipCallerSession struct {
	Contract *EPQFIMembershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EPQFIMembershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EPQFIMembershipTransactorSession struct {
	Contract     *EPQFIMembershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EPQFIMembershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type EPQFIMembershipRaw struct {
	Contract *EPQFIMembership // Generic contract binding to access the raw methods on
}

// EPQFIMembershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EPQFIMembershipCallerRaw struct {
	Contract *EPQFIMembershipCaller // Generic read-only contract binding to access the raw methods on
}

// EPQFIMembershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EPQFIMembershipTransactorRaw struct {
	Contract *EPQFIMembershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEPQFIMembership creates a new instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembership(address common.Address, backend bind.ContractBackend) (*EPQFIMembership, error) {
	contract, err := bindEPQFIMembership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembership{EPQFIMembershipCaller: EPQFIMembershipCaller{contract: contract}, EPQFIMembershipTransactor: EPQFIMembershipTransactor{contract: contract}, EPQFIMembershipFilterer: EPQFIMembershipFilterer{contract: contract}}, nil
}

// NewEPQFIMembershipCaller creates a new read-only instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipCaller(address common.Address, caller bind.ContractCaller) (*EPQFIMembershipCaller, error) {
	contract, err := bindEPQFIMembership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipCaller{contract: contract}, nil
}

// NewEPQFIMembershipTransactor creates a new write-only instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipTransactor(address common.Address, transactor bind.ContractTransactor) (*EPQFIMembershipTransactor, error) {
	contract, err := bindEPQFIMembership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipTransactor{contract: contract}, nil
}

// NewEPQFIMembershipFilterer creates a new log filterer instance of EPQFIMembership, bound to a specific deployed contract.
func NewEPQFIMembershipFilterer(address common.Address, filterer bind.ContractFilterer) (*EPQFIMembershipFilterer, error) {
	contract, err := bindEPQFIMembership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipFilterer{contract: contract}, nil
}

// bindEPQFIMembership binds a generic wrapper to an already deployed contract.
func bindEPQFIMembership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EPQFIMembershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPQFIMembership *EPQFIMembershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPQFIMembership.Contract.EPQFIMembershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPQFIMembership *EPQFIMembershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.EPQFIMembershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPQFIMembership *EPQFIMembershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.EPQFIMembershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EPQFIMembership *EPQFIMembershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EPQFIMembership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EPQFIMembership *EPQFIMembershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EPQFIMembership *EPQFIMembershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.contract.Transact(opts, method, params...)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCaller) ExpertLimitKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "expertLimitKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipSession) ExpertLimitKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertLimitKey(&_EPQFIMembership.CallOpts)
}

// ExpertLimitKey is a free data retrieval call binding the contract method 0x3e021814.
//
// Solidity: function expertLimitKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCallerSession) ExpertLimitKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertLimitKey(&_EPQFIMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCaller) ExpertVotingKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "expertVotingKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipSession) ExpertVotingKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertVotingKey(&_EPQFIMembership.CallOpts)
}

// ExpertVotingKey is a free data retrieval call binding the contract method 0x5313263a.
//
// Solidity: function expertVotingKey() view returns(string)
func (_EPQFIMembership *EPQFIMembershipCallerSession) ExpertVotingKey() (string, error) {
	return _EPQFIMembership.Contract.ExpertVotingKey(&_EPQFIMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCaller) Experts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "experts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipSession) Experts() (common.Address, error) {
	return _EPQFIMembership.Contract.Experts(&_EPQFIMembership.CallOpts)
}

// Experts is a free data retrieval call binding the contract method 0xc6836825.
//
// Solidity: function experts() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCallerSession) Experts() (common.Address, error) {
	return _EPQFIMembership.Contract.Experts(&_EPQFIMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCaller) GetLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipSession) GetLimit() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetLimit(&_EPQFIMembership.CallOpts)
}

// GetLimit is a free data retrieval call binding the contract method 0xb295a00e.
//
// Solidity: function getLimit() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetLimit() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetLimit(&_EPQFIMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipCaller) GetMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipSession) GetMembers() ([]common.Address, error) {
	return _EPQFIMembership.Contract.GetMembers(&_EPQFIMembership.CallOpts)
}

// GetMembers is a free data retrieval call binding the contract method 0x9eab5253.
//
// Solidity: function getMembers() view returns(address[])
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetMembers() ([]common.Address, error) {
	return _EPQFIMembership.Contract.GetMembers(&_EPQFIMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "getSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipSession) GetSize() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetSize(&_EPQFIMembership.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() view returns(uint256)
func (_EPQFIMembership *EPQFIMembershipCallerSession) GetSize() (*big.Int, error) {
	return _EPQFIMembership.Contract.GetSize(&_EPQFIMembership.CallOpts)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipCaller) IsMember(opts *bind.CallOpts, _exp common.Address) (bool, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "isMember", _exp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipSession) IsMember(_exp common.Address) (bool, error) {
	return _EPQFIMembership.Contract.IsMember(&_EPQFIMembership.CallOpts, _exp)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address _exp) view returns(bool)
func (_EPQFIMembership *EPQFIMembershipCallerSession) IsMember(_exp common.Address) (bool, error) {
	return _EPQFIMembership.Contract.IsMember(&_EPQFIMembership.CallOpts, _exp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EPQFIMembership.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipSession) Owner() (common.Address, error) {
	return _EPQFIMembership.Contract.Owner(&_EPQFIMembership.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EPQFIMembership *EPQFIMembershipCallerSession) Owner() (common.Address, error) {
	return _EPQFIMembership.Contract.Owner(&_EPQFIMembership.CallOpts)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) AddMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "addMember", _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.AddMember(&_EPQFIMembership.TransactOpts, _exp)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) AddMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.AddMember(&_EPQFIMembership.TransactOpts, _exp)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "initialize", _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.Initialize(&_EPQFIMembership.TransactOpts, _registry, _experts)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address _registry, address[] _experts) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) Initialize(_registry common.Address, _experts []common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.Initialize(&_EPQFIMembership.TransactOpts, _registry, _experts)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) RemoveMember(opts *bind.TransactOpts, _exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "removeMember", _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RemoveMember(&_EPQFIMembership.TransactOpts, _exp)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address _exp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) RemoveMember(_exp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RemoveMember(&_EPQFIMembership.TransactOpts, _exp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RenounceOwnership(&_EPQFIMembership.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EPQFIMembership.Contract.RenounceOwnership(&_EPQFIMembership.TransactOpts)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) SwapMember(opts *bind.TransactOpts, _oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "swapMember", _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.SwapMember(&_EPQFIMembership.TransactOpts, _oldExp, _newExp)
}

// SwapMember is a paid mutator transaction binding the contract method 0xb278982d.
//
// Solidity: function swapMember(address _oldExp, address _newExp) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) SwapMember(_oldExp common.Address, _newExp common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.SwapMember(&_EPQFIMembership.TransactOpts, _oldExp, _newExp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.TransferOwnership(&_EPQFIMembership.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EPQFIMembership *EPQFIMembershipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EPQFIMembership.Contract.TransferOwnership(&_EPQFIMembership.TransactOpts, newOwner)
}

// EPQFIMembershipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EPQFIMembership contract.
type EPQFIMembershipOwnershipTransferredIterator struct {
	Event *EPQFIMembershipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EPQFIMembershipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EPQFIMembershipOwnershipTransferred)
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
		it.Event = new(EPQFIMembershipOwnershipTransferred)
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
func (it *EPQFIMembershipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EPQFIMembershipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EPQFIMembershipOwnershipTransferred represents a OwnershipTransferred event raised by the EPQFIMembership contract.
type EPQFIMembershipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPQFIMembership *EPQFIMembershipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EPQFIMembershipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPQFIMembership.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EPQFIMembershipOwnershipTransferredIterator{contract: _EPQFIMembership.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EPQFIMembership *EPQFIMembershipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EPQFIMembershipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EPQFIMembership.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EPQFIMembershipOwnershipTransferred)
				if err := _EPQFIMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EPQFIMembership *EPQFIMembershipFilterer) ParseOwnershipTransferred(log types.Log) (*EPQFIMembershipOwnershipTransferred, error) {
	event := new(EPQFIMembershipOwnershipTransferred)
	if err := _EPQFIMembership.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
