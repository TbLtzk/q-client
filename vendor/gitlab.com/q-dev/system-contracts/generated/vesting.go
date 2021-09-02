// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TimeLockEntry is an auto generated low-level Go binding around an user-defined struct.
type TimeLockEntry struct {
	Amount       *big.Int
	ReleaseStart *big.Int
	ReleaseEnd   *big.Int
}

// VestingABI is the input ABI used to generate the binding from.
const VestingABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTimeLockEntry\",\"name\":\"_timelock\",\"type\":\"tuple\"}],\"name\":\"NewTimeLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTimeLockEntry\",\"name\":\"_timelock\",\"type\":\"tuple\"}],\"name\":\"Purged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releaseEnd\",\"type\":\"uint256\"}],\"name\":\"depositOnBehalfOf\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinimumBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTimeLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"releaseEnd\",\"type\":\"uint256\"}],\"internalType\":\"structTimeLockEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"purgeTimeLocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawableBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structVotingLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// VestingBin is the compiled bytecode used for deploying new contracts.
var VestingBin = "0x608060405234801561001057600080fd5b506115d5806100206000396000f3fe6080604052600436106100975760003560e01c80632e1a7d4d1461009c5780633aa6c4f9146100be5780636198e339146100f457806370a082311461011457806373983f4314610134578063a1c9e7c614610147578063afb60cff14610167578063b0426a3f14610187578063c4d66de8146101a9578063dd467064146101c9578063e7522fc0146101e9578063f0e8102414610209575b600080fd5b3480156100a857600080fd5b506100bc6100b736600461119b565b610236565b005b3480156100ca57600080fd5b506100de6100d93660046110a5565b6102cc565b6040516100eb9190611532565b60405180910390f35b34801561010057600080fd5b506100bc61010f36600461119b565b61031d565b34801561012057600080fd5b506100de61012f3660046110a5565b61032a565b6100bc610142366004611108565b610345565b34801561015357600080fd5b506100bc61016236600461119b565b610436565b34801561017357600080fd5b506100bc6101823660046110a5565b61051f565b34801561019357600080fd5b5061019c610528565b6040516100eb9190611507565b3480156101b557600080fd5b506100bc6101c43660046110a5565b61053d565b3480156101d557600080fd5b506100bc6101e436600461119b565b6105fb565b3480156101f557600080fd5b506100de6102043660046110dd565b610605565b34801561021557600080fd5b506102296102243660046110a5565b6106d4565b6040516100eb9190611223565b336102418183610767565b506000816001600160a01b03168360405161025b906111cc565b60006040518083038185875af1925050503d8060008114610298576040519150601f19603f3d011682016040523d82523d6000602084013e61029d565b606091505b50509050806102c75760405162461bcd60e51b81526004016102be90611333565b60405180910390fd5b505050565b6000806102d7610870565b905060006102e58442610605565b905080826000015111156103085781516102fe8561032a565b0392505050610318565b806103128561032a565b03925050505b919050565b610327338261097d565b50565b6001600160a01b031660009081526002602052604090205490565b8082106103645760405162461bcd60e51b81526004016102be906113f8565b600034116103845760405162461bcd60e51b81526004016102be90611380565b61038d83610a64565b4281111561042c5761039d61101a565b348152602080820184815260408084018581526001600160a01b03881660008181526001808752848220805480830182559083529690912087516003909702019586559351938501939093555160029093019290925590517f42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f906104229084906114f9565b60405180910390a2505b6102c78334610c06565b600354604051633fb9027160e01b81526000916001600160a01b031690633fb902719061046590600401611469565b60206040518083038186803b15801561047d57600080fd5b505afa158015610491573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b591906110c1565b604051633a5f594760e21b81529091506001600160a01b0382169063e97d651c906104e990339086906000906004016111cf565b600060405180830381600087803b15801561050357600080fd5b505af1158015610517573d6000803e3d6000fd5b505050505050565b61032781610a64565b61053061103b565b610538610870565b905090565b600054610100900460ff16806105565750610556610c4c565b80610564575060005460ff16155b61059f5760405162461bcd60e51b815260040180806020018281038252602e815260200180611551602e913960400191505060405180910390fd5b600054610100900460ff161580156105ca576000805460ff1961ff0019909116610100171660011790555b600380546001600160a01b0319166001600160a01b03841617905580156105f7576000805461ff00191690555b5050565b6103273382610c52565b6001600160a01b03821660009081526001602052604081205461062a575060006106ce565b6000805b6001600160a01b0385166000908152600160205260409020548110156106ca576001600160a01b038516600090815260016020526040812080548390811061067257fe5b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506106bf6106b88287610d88565b8490610de5565b92505060010161062e565b5090505b92915050565b6001600160a01b0381166000908152600160209081526040808320805482518185028101850190935280835260609492939192909184015b8282101561075c578382906000526020600020906003020160405180606001604052908160008201548152602001600182015481526020016002820154815250508152602001906001019061070c565b505050509050919050565b600081610776575060006106ce565b816107808461032a565b101561079e5760405162461bcd60e51b81526004016102be906112da565b6107a88342610605565b6107bb836107b58661032a565b90610e44565b10156107d95760405162461bcd60e51b81526004016102be906114a0565b60006107e3610870565b604081015181519192506000916107f991610de5565b9050600061080a826107b58861032a565b90508481101561082757610827866108228784610e44565b61097d565b6001600160a01b03861660009081526002602052604090205461084a9086610e44565b6001600160a01b0387166000908152600260205260409020555060019250505092915050565b61087861103b565b600354604051633fb9027160e01b81526000916001600160a01b031690633fb90271906108a790600401611469565b60206040518083038186803b1580156108bf57600080fd5b505afa1580156108d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f791906110c1565b6040516218724160e61b81529091506001600160a01b0382169063061c90409061092790309033906004016111f0565b60806040518083038186803b15801561093f57600080fd5b505afa158015610953573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610977919061113c565b91505090565b600354604051633fb9027160e01b81526000916001600160a01b031690633fb90271906109ac90600401611469565b60206040518083038186803b1580156109c457600080fd5b505afa1580156109d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109fc91906110c1565b604051637eee288d60e01b81529091506001600160a01b03821690637eee288d90610a2d908690869060040161120a565b600060405180830381600087803b158015610a4757600080fd5b505af1158015610a5b573d6000803e3d6000fd5b50505050505050565b6001600160a01b038116600090815260016020908152604080832080548251818502810185019093528083529192909190849084015b82821015610aea5783829060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505081526020019060010190610a9a565b5050506001600160a01b0384166000908152600160205260408120929350610b13929150611063565b60005b81518110156102c757818181518110610b2b57fe5b602002602001015160400151421015610ba9576001600160a01b03831660009081526001602052604090208251839083908110610b6457fe5b60209081029190910181015182546001818101855560009485529383902082516003909202019081559181015192820192909255604090910151600290910155610bfe565b826001600160a01b03167f646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93838381518110610be057fe5b6020026020010151604051610bf591906114f9565b60405180910390a25b600101610b16565b6001600160a01b038216600090815260026020526040902054610c299082610de5565b6001600160a01b0383166000908152600260205260409020556105f78282610c52565b303b1590565b600354604051633fb9027160e01b81526000916001600160a01b031690633fb9027190610c8190600401611469565b60206040518083038186803b158015610c9957600080fd5b505afa158015610cad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd191906110c1565b90506000610cdd610870565b60408101518151919250600091610cf391610de5565b905083610d03826107b58861032a565b1015610d215760405162461bcd60e51b81526004016102be90611271565b60405163282d3fdf60e01b81526001600160a01b0384169063282d3fdf90610d4f908890889060040161120a565b600060405180830381600087803b158015610d6957600080fd5b505af1158015610d7d573d6000803e3d6000fd5b505050505050505050565b60008260200151821015610d9e575081516106ce565b82604001518210610db1575060006106ce565b604083015160208401518451848303929190910390610ddc908290610dd69085610e86565b90610edf565b95945050505050565b600082820183811015610e3d576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b6000610e3d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610f1e565b600082610e95575060006106ce565b82820282848281610ea257fe5b0414610e3d5760405162461bcd60e51b815260040180806020018281038252602181526020018061157f6021913960400191505060405180910390fd5b6000610e3d83836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250610fb5565b60008184841115610fad5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f72578181015183820152602001610f5a565b50505050905090810190601f168015610f9f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600081836110045760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610f72578181015183820152602001610f5a565b50600083858161101057fe5b0495945050505050565b60405180606001604052806000815260200160008152602001600081525090565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b508054600082556003029060005260206000209081019061032791905b808211156110a1576000808255600182018190556002820155600301611080565b5090565b6000602082840312156110b6578081fd5b8135610e3d8161153b565b6000602082840312156110d2578081fd5b8151610e3d8161153b565b600080604083850312156110ef578081fd5b82356110fa8161153b565b946020939093013593505050565b60008060006060848603121561111c578081fd5b83356111278161153b565b95602085013595506040909401359392505050565b60006080828403121561114d578081fd5b6040516080810181811067ffffffffffffffff8211171561116a57fe5b8060405250825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b6000602082840312156111ac578081fd5b5035919050565b8051825260208082015190830152604090810151910152565b90565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0392831681529116602082015260400190565b6001600160a01b03929092168252602082015260400190565b6020808252825182820181905260009190848201906040850190845b81811015611265576112528385516111b3565b928401926060929092019160010161123f565b50909695505050505050565b60208082526043908201527f5b5145432d3031373031365d2d54686520617661696c61626c652062616c616e60408201527f6365206d7573742062652067726561746572207468616e2074686520616d6f75606082015262373a1760e91b608082015260a00190565b60208082526039908201527f5b5145432d3033313030315d2d496e73756666696369656e742062616c616e6360408201527832903337b9103b32b9ba34b733903bb4ba34323930bbb0b61760391b606082015260800190565b6020808252602d908201527f5b5145432d3033313030335d2d4661696c656420746f2077697468647261772060408201526c333937b6903b32b9ba34b7339760991b606082015260800190565b60208082526052908201527f5b5145432d3031373030305d2d496e76616c69642076616c756520666f72207460408201527f6865206465706f73697420616d6f756e742c206661696c656420746f2064657060608201527137b9b4ba1030903d32b937903b30b63ab29760711b608082015260a00190565b6020808252604b908201527f5b5145432d3033323030315d2d54696d65206c6f636b2072656c65617365206560408201527f6e642074696d652073686f756c642062652061667465722072656c656173652060608201526a39ba30b93a103a34b6b29760a91b608082015260a00190565b6020808252601c908201527f676f7665726e616e63652e766f74696e6757656967687450726f787900000000604082015260600190565b60208082526039908201527f5b5145432d3033313030325d2d56657374696e67207769746864726177616c20604082015278383932bb32b73a32b210313c903a34b6b2b637b1b59439949760391b606082015260800190565b606081016106ce82846111b3565b8151815260208083015190820152604080830151908201526060918201519181019190915260800190565b90815260200190565b6001600160a01b038116811461032757600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212204138091d735b4488599e511e04ece624210578ef66ce52ce388c6af4d1c8840c64736f6c63430007060033"

// DeployVesting deploys a new Ethereum contract, binding an instance of Vesting to it.
func DeployVesting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vesting, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VestingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// Vesting is an auto generated Go binding around an Ethereum contract.
type Vesting struct {
	VestingCaller     // Read-only binding to the contract
	VestingTransactor // Write-only binding to the contract
	VestingFilterer   // Log filterer for contract events
}

// VestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VestingSession struct {
	Contract     *Vesting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VestingCallerSession struct {
	Contract *VestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VestingTransactorSession struct {
	Contract     *VestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VestingRaw struct {
	Contract *Vesting // Generic contract binding to access the raw methods on
}

// VestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VestingCallerRaw struct {
	Contract *VestingCaller // Generic read-only contract binding to access the raw methods on
}

// VestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VestingTransactorRaw struct {
	Contract *VestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVesting creates a new instance of Vesting, bound to a specific deployed contract.
func NewVesting(address common.Address, backend bind.ContractBackend) (*Vesting, error) {
	contract, err := bindVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vesting{VestingCaller: VestingCaller{contract: contract}, VestingTransactor: VestingTransactor{contract: contract}, VestingFilterer: VestingFilterer{contract: contract}}, nil
}

// NewVestingCaller creates a new read-only instance of Vesting, bound to a specific deployed contract.
func NewVestingCaller(address common.Address, caller bind.ContractCaller) (*VestingCaller, error) {
	contract, err := bindVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VestingCaller{contract: contract}, nil
}

// NewVestingTransactor creates a new write-only instance of Vesting, bound to a specific deployed contract.
func NewVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*VestingTransactor, error) {
	contract, err := bindVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VestingTransactor{contract: contract}, nil
}

// NewVestingFilterer creates a new log filterer instance of Vesting, bound to a specific deployed contract.
func NewVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*VestingFilterer, error) {
	contract, err := bindVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VestingFilterer{contract: contract}, nil
}

// bindVesting binds a generic wrapper to an already deployed contract.
func bindVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.VestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.VestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vesting *VestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vesting *VestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vesting *VestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vesting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "balanceOf", _account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.BalanceOf(&_Vesting.CallOpts, _account)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingCaller) GetLockInfo(opts *bind.CallOpts) (VotingLockInfo, error) {
	var (
		ret0 = new(VotingLockInfo)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "getLockInfo")
	return *ret0, err
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingSession) GetLockInfo() (VotingLockInfo, error) {
	return _Vesting.Contract.GetLockInfo(&_Vesting.CallOpts)
}

// GetLockInfo is a free data retrieval call binding the contract method 0xb0426a3f.
//
// Solidity: function getLockInfo() view returns((uint256,uint256,uint256,uint256))
func (_Vesting *VestingCallerSession) GetLockInfo() (VotingLockInfo, error) {
	return _Vesting.Contract.GetLockInfo(&_Vesting.CallOpts)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingCaller) GetMinimumBalance(opts *bind.CallOpts, _account common.Address, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "getMinimumBalance", _account, _timestamp)
	return *ret0, err
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Vesting.Contract.GetMinimumBalance(&_Vesting.CallOpts, _account, _timestamp)
}

// GetMinimumBalance is a free data retrieval call binding the contract method 0xe7522fc0.
//
// Solidity: function getMinimumBalance(address _account, uint256 _timestamp) view returns(uint256)
func (_Vesting *VestingCallerSession) GetMinimumBalance(_account common.Address, _timestamp *big.Int) (*big.Int, error) {
	return _Vesting.Contract.GetMinimumBalance(&_Vesting.CallOpts, _account, _timestamp)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingCaller) GetTimeLocks(opts *bind.CallOpts, _account common.Address) ([]TimeLockEntry, error) {
	var (
		ret0 = new([]TimeLockEntry)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "getTimeLocks", _account)
	return *ret0, err
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Vesting.Contract.GetTimeLocks(&_Vesting.CallOpts, _account)
}

// GetTimeLocks is a free data retrieval call binding the contract method 0xf0e81024.
//
// Solidity: function getTimeLocks(address _account) view returns((uint256,uint256,uint256)[])
func (_Vesting *VestingCallerSession) GetTimeLocks(_account common.Address) ([]TimeLockEntry, error) {
	return _Vesting.Contract.GetTimeLocks(&_Vesting.CallOpts, _account)
}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCaller) WithdrawableBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vesting.contract.Call(opts, out, "withdrawableBalanceOf", _account)
	return *ret0, err
}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingSession) WithdrawableBalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.WithdrawableBalanceOf(&_Vesting.CallOpts, _account)
}

// WithdrawableBalanceOf is a free data retrieval call binding the contract method 0x3aa6c4f9.
//
// Solidity: function withdrawableBalanceOf(address _account) view returns(uint256)
func (_Vesting *VestingCallerSession) WithdrawableBalanceOf(_account common.Address) (*big.Int, error) {
	return _Vesting.Contract.WithdrawableBalanceOf(&_Vesting.CallOpts, _account)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) AnnounceUnlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "announceUnlock", _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AnnounceUnlock(&_Vesting.TransactOpts, _amount)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xa1c9e7c6.
//
// Solidity: function announceUnlock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) AnnounceUnlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.AnnounceUnlock(&_Vesting.TransactOpts, _amount)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingTransactor) DepositOnBehalfOf(opts *bind.TransactOpts, _account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "depositOnBehalfOf", _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.DepositOnBehalfOf(&_Vesting.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// DepositOnBehalfOf is a paid mutator transaction binding the contract method 0x73983f43.
//
// Solidity: function depositOnBehalfOf(address _account, uint256 _releaseStart, uint256 _releaseEnd) payable returns()
func (_Vesting *VestingTransactorSession) DepositOnBehalfOf(_account common.Address, _releaseStart *big.Int, _releaseEnd *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.DepositOnBehalfOf(&_Vesting.TransactOpts, _account, _releaseStart, _releaseEnd)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Initialize(&_Vesting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Vesting *VestingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.Initialize(&_Vesting.TransactOpts, _registry)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Lock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "lock", _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Lock(&_Vesting.TransactOpts, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Lock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Lock(&_Vesting.TransactOpts, _amount)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingTransactor) PurgeTimeLocks(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "purgeTimeLocks", _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PurgeTimeLocks(&_Vesting.TransactOpts, _account)
}

// PurgeTimeLocks is a paid mutator transaction binding the contract method 0xafb60cff.
//
// Solidity: function purgeTimeLocks(address _account) returns()
func (_Vesting *VestingTransactorSession) PurgeTimeLocks(_account common.Address) (*types.Transaction, error) {
	return _Vesting.Contract.PurgeTimeLocks(&_Vesting.TransactOpts, _account)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Unlock(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "unlock", _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Unlock(&_Vesting.TransactOpts, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Unlock(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Unlock(&_Vesting.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vesting.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Withdraw(&_Vesting.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Vesting *VestingTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Vesting.Contract.Withdraw(&_Vesting.TransactOpts, _amount)
}

// VestingNewTimeLockIterator is returned from FilterNewTimeLock and is used to iterate over the raw logs and unpacked data for NewTimeLock events raised by the Vesting contract.
type VestingNewTimeLockIterator struct {
	Event *VestingNewTimeLock // Event containing the contract specifics and raw log

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
func (it *VestingNewTimeLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingNewTimeLock)
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
		it.Event = new(VestingNewTimeLock)
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
func (it *VestingNewTimeLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingNewTimeLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingNewTimeLock represents a NewTimeLock event raised by the Vesting contract.
type VestingNewTimeLock struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewTimeLock is a free log retrieval operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) FilterNewTimeLock(opts *bind.FilterOpts, _account []common.Address) (*VestingNewTimeLockIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return &VestingNewTimeLockIterator{contract: _Vesting.contract, event: "NewTimeLock", logs: logs, sub: sub}, nil
}

// WatchNewTimeLock is a free log subscription operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) WatchNewTimeLock(opts *bind.WatchOpts, sink chan<- *VestingNewTimeLock, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "NewTimeLock", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingNewTimeLock)
				if err := _Vesting.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
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

// ParseNewTimeLock is a log parse operation binding the contract event 0x42dc784443c8161a28f39bb05a48bc8bac980ab9115d5b9cfef307f00204bd7f.
//
// Solidity: event NewTimeLock(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) ParseNewTimeLock(log types.Log) (*VestingNewTimeLock, error) {
	event := new(VestingNewTimeLock)
	if err := _Vesting.contract.UnpackLog(event, "NewTimeLock", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VestingPurgedIterator is returned from FilterPurged and is used to iterate over the raw logs and unpacked data for Purged events raised by the Vesting contract.
type VestingPurgedIterator struct {
	Event *VestingPurged // Event containing the contract specifics and raw log

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
func (it *VestingPurgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VestingPurged)
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
		it.Event = new(VestingPurged)
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
func (it *VestingPurgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VestingPurgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VestingPurged represents a Purged event raised by the Vesting contract.
type VestingPurged struct {
	Account  common.Address
	Timelock TimeLockEntry
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPurged is a free log retrieval operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) FilterPurged(opts *bind.FilterOpts, _account []common.Address) (*VestingPurgedIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.FilterLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return &VestingPurgedIterator{contract: _Vesting.contract, event: "Purged", logs: logs, sub: sub}, nil
}

// WatchPurged is a free log subscription operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) WatchPurged(opts *bind.WatchOpts, sink chan<- *VestingPurged, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Vesting.contract.WatchLogs(opts, "Purged", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VestingPurged)
				if err := _Vesting.contract.UnpackLog(event, "Purged", log); err != nil {
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

// ParsePurged is a log parse operation binding the contract event 0x646189d6820ad36442dfb379131756840be21178303b28f1f607ca4e45762e93.
//
// Solidity: event Purged(address indexed _account, (uint256,uint256,uint256) _timelock)
func (_Vesting *VestingFilterer) ParsePurged(log types.Log) (*VestingPurged, error) {
	event := new(VestingPurged)
	if err := _Vesting.contract.UnpackLog(event, "Purged", log); err != nil {
		return nil, err
	}
	return event, nil
}
