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

// BaseVotingWeightInfo is an auto generated low-level Go binding around an user-defined struct.
type BaseVotingWeightInfo struct {
	OwnWeight        *big.Int
	VotingAgent      common.Address
	DelegationStatus uint8
	LockedUntil      *big.Int
}

// VotingLockInfo is an auto generated low-level Go binding around an user-defined struct.
type VotingLockInfo struct {
	LockedAmount        *big.Int
	LockedUntil         *big.Int
	PendingUnlockAmount *big.Int
	PendingUnlockTime   *big.Int
}

// VotingWeightProxyMetaData contains all meta data concerning the VotingWeightProxy contract.
var VotingWeightProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newLockedAmount\",\"type\":\"uint256\"}],\"name\":\"LockedAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prevVotingAgent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newVotingAgent\",\"type\":\"address\"}],\"name\":\"NewVotingAgentAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockAmount\",\"type\":\"uint256\"}],\"name\":\"PendingUnlockAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockTime\",\"type\":\"uint256\"}],\"name\":\"PendingUnlockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_votingAgent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_delegatedAmount\",\"type\":\"uint256\"}],\"name\":\"VotingAgentChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAgent\",\"type\":\"address\"}],\"name\":\"announceNewVotingAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newUnlockTime\",\"type\":\"uint256\"}],\"name\":\"announceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegationInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingAgentPassOverTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockNeededUntil\",\"type\":\"uint256\"}],\"name\":\"extendLocking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"forceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_votingEndTime\",\"type\":\"uint256\"}],\"name\":\"getBaseVotingWeightInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ownWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"enumDelegationStatus\",\"name\":\"delegationStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVotingWeightInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structVotingLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getLockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockNeededUntil\",\"type\":\"uint256\"}],\"name\":\"getVotingWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_tokenLockSourcesKeys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_votingContractsKeys\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setNewVotingAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061245e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c55760003560e01c8063061c9040146100ca57806310874b5b14610119578063282d3fdf146101395780632f4294741461014e5780634cb77f97146101ad57806350696346146101ce5780635c7700df146101e15780636b164f7d146101f45780637bf9f918146102075780637eee288d1461021a578063929ec5371461022d578063d1e0fea114610240578063e97d651c146102b3578063efe86a5e146102c6578063f927c6f2146102d9575b600080fd5b6100dd6100d8366004611ecf565b6102e1565b60405161011091908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b60405180910390f35b61012c610127366004611f08565b61047c565b6040516101109190611f4a565b61014c610147366004611f08565b610577565b005b61018d61015c366004611ecf565b6003602081815260009384526040808520909152918352912080546001820154600283015492909301549092919084565b604080519485526020850193909352918301526060820152608001610110565b6101c06101bb366004611f08565b61084a565b604051908152602001610110565b6101c06101dc366004611fa0565b610a38565b61014c6101ef366004611f08565b610a53565b61014c610202366004611fa0565b610ca7565b6101c0610215366004611f08565b610e13565b61014c610228366004611f08565b610fae565b6101c061023b366004611fa0565b611293565b61028561024e366004611fa0565b60056020526000908152604090208054600182015460029092015490916001600160a01b03811691600160a01b90910460ff169084565b604080519485526001600160a01b039093166020850152901515918301919091526060820152608001610110565b61014c6102c1366004611fbd565b61139f565b61014c6102d436600461212e565b611558565b61014c611658565b6102e9611d1e565b6000805b6001548110156103c657600054600180546001600160a01b03808916936201000090041691633fb902719185908110610328576103286121a4565b906000526020600020016040518263ffffffff1660e01b815260040161034e91906121f5565b60206040518083038186803b15801561036657600080fd5b505afa15801561037a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061039e919061229d565b6001600160a01b031614156103b657600191506103c6565b6103bf816122d0565b90506102ed565b506000604051806060016040528060278152602001612402602791399050808261040c5760405162461bcd60e51b815260040161040391906122eb565b60405180910390fd5b506001600160a01b0380861660009081526003602081815260408084209489168452938152918390208351608081018552815481526001820154938101939093526002810154938301939093529190910154606082015261046c856118b2565b6020820152925050505b92915050565b610484611d46565b6001600160a01b03838116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff1615159181019190915260029091015460608201526104e3611d46565b6104ed858561084a565b81526020808301516001600160a01b031690820181905261050d90610a38565b606082015260408201511561053e576040810160025b9081600281111561053657610536611f34565b90525061056f565b846001600160a01b031681602001516001600160a01b0316141561056757604081016000610523565b600160408201525b949350505050565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb9027191859081106105b3576105b36121a4565b906000526020600020016040518263ffffffff1660e01b81526004016105d991906121f5565b60206040518083038186803b1580156105f157600080fd5b505afa158015610605573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610629919061229d565b6001600160a01b031614156107b957600082116106585760405162461bcd60e51b815260040161040390612340565b3360009081526003602081815260408084206001600160a01b038816855282529283902083516080810185528154808252600183015493820193909352600282015494810194909452909101546060830152839082906106b9908390612399565b9052503360008181526003602081815260408084206001600160a01b038a168086529083529381902086518082559287015160018201558187015160028201556060870151930192909255905191926000805160206123e28339815191529261072292906123b1565b60405180910390a26001600160a01b038085166000908152600560205260409020600181015490911661076d576001810180546001600160a01b0319166001600160a01b0387161790555b6001810154600160a01b900460ff166107b25760018101546001600160a01b0316600090815260056020526040812080548692906107ac908490612399565b90915550505b5050505050565b6107c2816122d0565b905061057a565b5060405162461bcd60e51b815260206004820152604960248201527f5b5145432d3032383030375d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c792074686520746f6b656e206c6f636b20736f757263657320686176606482015268329030b1b1b2b9b99760b91b608482015260a401610403565b6000610854611d1e565b6001600160a01b0384811660008181526005602090815260408083208151608081018352815481526001820154968716938101849052600160a01b90960460ff1615159186019190915260020154606085015290911480156108b857508160400151155b15610a225760005b600154811015610a2057600360008060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027160018581548110610904576109046121a4565b906000526020600020016040518263ffffffff1660e01b815260040161092a91906121f5565b60206040518083038186803b15801561094257600080fd5b505afa158015610956573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097a919061229d565b6001600160a01b03166001600160a01b031681526020019081526020016000206000886001600160a01b03166001600160a01b0316815260200190815260200160002060405180608001604052908160008201548152602001600182015481526020016002820154815260200160038201548152505093508584606001511115610a10576040840151610a0d9083612399565b91505b610a19816122d0565b90506108c0565b505b8151610a2e9082612399565b9695505050505050565b6001600160a01b031660009081526004602052604090205490565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb902719185908110610a8f57610a8f6121a4565b906000526020600020016040518263ffffffff1660e01b8152600401610ab591906121f5565b60206040518083038186803b158015610acd57600080fd5b505afa158015610ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b05919061229d565b6001600160a01b03161415610c975760008211610b345760405162461bcd60e51b815260040161040390612340565b3360009081526003602090815260408083206001600160a01b0387168452909152812060028101548154919291610b6b9190612399565b905083811015610bf15760405162461bcd60e51b8152602060048201526044602482018190527f5b5145432d3032383030365d2d43616e6e6f7420656e666f72636520746f2075908201527f6e6c6f636b206d6f7265207468616e2069732063757272656e746c79206c6f6360648201526335b2b21760e11b608482015260a401610403565b6002820154801515908511610c1a57848360020154610c1091906123ca565b6002840155610c41565b6000836002015486610c2c91906123ca565b600060028601559050610c3f8782611907565b505b8015610c8f57856001600160a01b03167f70601d6dee302c74325828267b70767d970eddb01567778dfc92cd02aa7671bd338560020154604051610c869291906123b1565b60405180910390a25b505050505050565b610ca0816122d0565b9050610a56565b6001600160a01b038116610d195760405162461bcd60e51b815260206004820152603360248201527f5b3032383031315d2d4e657720766f74696e67206167656e742063616e6e6f746044820152721031329030903d32b9379030b2323932b9b99760691b6064820152608401610403565b33600090815260056020526040902060018101546001600160a01b03811690600160a01b900460ff16610daf57610d4f33611293565b6001600160a01b038216600090815260056020526040902054610d7291906123ca565b6001600160a01b038216600090815260056020526040902055610d9481610a38565b600283015560018201805460ff60a01b1916600160a01b1790555b6001820180546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915233917f3b955a76d9960f085ab98c0fad7fad089e2fc0506bcb86b069c300112625df9d910160405180910390a2505050565b6000805b600254811015610f33576000546002805433926201000090046001600160a01b031691633fb902719185908110610e5057610e506121a4565b906000526020600020016040518263ffffffff1660e01b8152600401610e7691906121f5565b60206040518083038186803b158015610e8e57600080fd5b505afa158015610ea2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec6919061229d565b6001600160a01b03161415610f23576001600160a01b038416600090815260046020526040902054610ef89084611b41565b6001600160a01b038516600090815260046020526040902055610f1b848461084a565b915050610476565b610f2c816122d0565b9050610e17565b5060405162461bcd60e51b815260206004820152604360248201527f5b5145432d3032383031305d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920766f74696e6720636f6e7472616374732068617665206163636560648201526239b99760e91b608482015260a401610403565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb902719185908110610fea57610fea6121a4565b906000526020600020016040518263ffffffff1660e01b815260040161101091906121f5565b60206040518083038186803b15801561102857600080fd5b505afa15801561103c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611060919061229d565b6001600160a01b03161415611283576000821161108f5760405162461bcd60e51b815260040161040390612340565b3360009081526003602081815260408084206001600160a01b03881685529091529091209081015442116111425760405162461bcd60e51b815260206004820152604e60248201527f5b5145432d3032383030345d2d4e6f7420656e6f7567682074696d652068617360448201527f20656c61707365642073696e63652074686520616e6e6f756e63656d656e742060648201526d37b3103a3432903ab73637b1b59760911b608482015260a401610403565b6002810154815460009161115591612399565b9050838110156111cd5760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3032383030365d2d417661696c61626c65206c6f636b656420616d60448201527f6f756e74206c657373207468616e20756e6c6f636b20616d6f756e742e0000006064820152608401610403565b60028201548015159085116111ec57848360020154610c1091906123ca565b426111f6876118b2565b10610c1a5760405162461bcd60e51b815260206004820152605160248201527f5b5145432d3032383030355d2d536d61727420756e6c6f636b206e6f7420706f60448201527f737369626c652c20746f6b656e7320617265207374696c6c206c6f636b656420606482015270313c903932b1b2b73a103b37ba34b7339760791b608482015260a401610403565b61128c816122d0565b9050610fb1565b600080805b60015481101561139857600360008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271600185815481106112dc576112dc6121a4565b906000526020600020016040518263ffffffff1660e01b815260040161130291906121f5565b60206040518083038186803b15801561131a57600080fd5b505afa15801561132e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611352919061229d565b6001600160a01b03908116825260208083019390935260409182016000908120918816815292529020546113869083612399565b9150611391816122d0565b9050611298565b5092915050565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb9027191859081106113db576113db6121a4565b906000526020600020016040518263ffffffff1660e01b815260040161140191906121f5565b60206040518083038186803b15801561141957600080fd5b505afa15801561142d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611451919061229d565b6001600160a01b03161415611548573360009081526003602090815260408083206001600160a01b0388168452909152902060028101548410156114b35760008482600201546114a191906123ca565b90506114ad8682611b57565b506114d3565b60008160020154856114c591906123ca565b90506114d18682611907565b505b6002810184905560006114e5866118b2565b9050808410156114f3578093505b60038201849055600282015460408051338152602081019290925281018590526001600160a01b038716907f531e4f82f00c1943a0c7d64474e788195d4cf400182d1dbf9d7a3fd9787a4d4590606001610c86565b611551816122d0565b90506113a2565b600054610100900460ff1680611571575060005460ff16155b6115d45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610403565b600054610100900460ff161580156115f6576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b03871602179055825161162b906001906020860190611d6e565b50815161163f906002906020850190611d6e565b508015611652576000805461ff00191690555b50505050565b3360009081526005602090815260409182902082516080810184528154815260018201546001600160a01b03811693820193909352600160a01b90920460ff161515928201839052600201546060820152906117405760405162461bcd60e51b815260206004820152605b60248201527f5b5145432d3032383030385d2d4368616e67696e672074686520766f74696e6760448201527f206167656e742068617320746f20626520616e6e6f756e6365642c206661696c60648201527a32b2103a379039b2ba103732bb903b37ba34b7339030b3b2b73a1760291b608482015260a401610403565b42816060015111156117cb5760405162461bcd60e51b815260206004820152604860248201527f5b5145432d3032383030395d2d43616e6e6f74206368616e676520766f74696e60448201527f67206167656e74206265666f726520706173736f7665722074696d65206973206064820152673932b0b1b432b21760c11b608482015260a401610403565b600060408083018281523380845260056020908152928420855181559285015160018401805493511515600160a01b026001600160a81b03199094166001600160a01b03929092169190911792909217909155606084015160029092019190915561183590611293565b6020808401516001600160a01b038116600090815260059092526040822080549394509092849290611868908490612399565b90915550506040518281526001600160a01b0382169033907f395be45122e29f50a1f27a1ffd8913eac762e146c7f7214d01b52f2bc1dd34879060200160405180910390a3505050565b6001600160a01b0380821660009081526005602090815260408083206004909252808320546001830154909416835282205460028201549293919261190092916118fb91611b41565b611b41565b9392505050565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb902719185908110611943576119436121a4565b906000526020600020016040518263ffffffff1660e01b815260040161196991906121f5565b60206040518083038186803b15801561198157600080fd5b505afa158015611995573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b9919061229d565b6001600160a01b03161415611b31573360009081526003602090815260408083206001600160a01b038716845290915290208054831115611a5e5760405162461bcd60e51b815260206004820152603960248201527f5b5145432d3032383030335d2d43616e6e6f7420756e6c6f636b206d6f7265206044820152783a3430b71034b99031bab93932b73a363c903637b1b5b2b21760391b6064820152608401610403565b82816000016000828254611a7291906123ca565b909155505080546040516001600160a01b038616916000805160206123e283398151915291611aa29133916123b1565b60405180910390a26001600160a01b03848116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff16151591810182905260029092015460608301526107b2576020808201516001600160a01b0316600090815260059091526040812080548692906107ac9084906123ca565b611b3a816122d0565b905061190a565b6000818311611b505781611900565b5090919050565b60005b6001548110156107c9576000546001805433926201000090046001600160a01b031691633fb902719185908110611b9357611b936121a4565b906000526020600020016040518263ffffffff1660e01b8152600401611bb991906121f5565b60206040518083038186803b158015611bd157600080fd5b505afa158015611be5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c09919061229d565b6001600160a01b03161415611d0e573360009081526003602090815260408083206001600160a01b03871684529091528120805490918491839190611c4f908490612399565b909155505080546040516001600160a01b038616916000805160206123e283398151915291611c7f9133916123b1565b60405180910390a26001600160a01b03848116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff16151591810182905260029092015460608301526107b2576020808201516001600160a01b0316600090815260059091526040812080548692906107ac908490612399565b611d17816122d0565b9050611b5a565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260008082526020820181905290918201908152602001600081525090565b828054828255906000526020600020908101928215611dbb579160200282015b82811115611dbb5782518051611dab918491602090910190611dcb565b5091602001919060010190611d8e565b50611dc7929150611e4b565b5090565b828054611dd7906121ba565b90600052602060002090601f016020900481019282611df95760008555611e3f565b82601f10611e1257805160ff1916838001178555611e3f565b82800160010185558215611e3f579182015b82811115611e3f578251825591602001919060010190611e24565b50611dc7929150611e68565b80821115611dc7576000611e5f8282611e7d565b50600101611e4b565b5b80821115611dc75760008155600101611e69565b508054611e89906121ba565b6000825580601f10611e99575050565b601f016020900490600052602060002090810190611eb79190611e68565b50565b6001600160a01b0381168114611eb757600080fd5b60008060408385031215611ee257600080fd5b8235611eed81611eba565b91506020830135611efd81611eba565b809150509250929050565b60008060408385031215611f1b57600080fd5b8235611f2681611eba565b946020939093013593505050565b634e487b7160e01b600052602160045260246000fd5b815181526020808301516001600160a01b0316908201526040820151608082019060038110611f8957634e487b7160e01b600052602160045260246000fd5b806040840152506060830151606083015292915050565b600060208284031215611fb257600080fd5b813561190081611eba565b600080600060608486031215611fd257600080fd5b8335611fdd81611eba565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561203157612031611ff2565b604052919050565b6000601f838184011261204b57600080fd5b8235602067ffffffffffffffff8083111561206857612068611ff2565b8260051b612077838201612008565b938452868101830193838101908986111561209157600080fd5b84890192505b85831015612121578235848111156120af5760008081fd5b8901603f81018b136120c15760008081fd5b858101356040868211156120d7576120d7611ff2565b6120e8828b01601f19168901612008565b8281528d828486010111156120fd5760008081fd5b828285018a8301376000928101890192909252508352509184019190840190612097565b9998505050505050505050565b60008060006060848603121561214357600080fd5b833561214e81611eba565b9250602084013567ffffffffffffffff8082111561216b57600080fd5b61217787838801612039565b9350604086013591508082111561218d57600080fd5b5061219a86828701612039565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806121ce57607f821691505b602082108114156121ef57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208083526000845481600182811c91508083168061221757607f831692505b85831081141561223557634e487b7160e01b85526022600452602485fd5b87860183815260200181801561225257600181146122635761228e565b60ff1986168252878201965061228e565b60008b81526020902060005b868110156122885781548482015290850190890161226f565b83019750505b50949998505050505050505050565b6000602082840312156122af57600080fd5b815161190081611eba565b634e487b7160e01b600052601160045260246000fd5b60006000198214156122e4576122e46122ba565b5060010190565b600060208083528351808285015260005b81811015612318578581018301518582016040015282016122fc565b8181111561232a576000604083870101525b50601f01601f1916929092016040019392505050565b60208082526039908201527f5b5145432d3032383030325d2d496e76616c696420616d6f756e742076616c7560408201527832961030b6b7bab73a1031b0b73737ba103132903d32b9379760391b606082015260800190565b600082198211156123ac576123ac6122ba565b500190565b6001600160a01b03929092168252602082015260400190565b6000828210156123dc576123dc6122ba565b50039056fe48f899e185e3470300612fd03b36ae16768c2212cc0fdc94f44cf5d60bf3227e5b5145432d3032383030315d2d556e6b6e6f776e20746f6b656e206c6f636b20736f757263652ea264697066735822122017e0abd90cc3592490cdec77fd5d9cf08d61804c328b938ac391a9a6e240e4fa64736f6c63430008090033",
}

// VotingWeightProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingWeightProxyMetaData.ABI instead.
var VotingWeightProxyABI = VotingWeightProxyMetaData.ABI

// VotingWeightProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingWeightProxyMetaData.Bin instead.
var VotingWeightProxyBin = VotingWeightProxyMetaData.Bin

// DeployVotingWeightProxy deploys a new Ethereum contract, binding an instance of VotingWeightProxy to it.
func DeployVotingWeightProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VotingWeightProxy, error) {
	parsed, err := VotingWeightProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingWeightProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VotingWeightProxy{VotingWeightProxyCaller: VotingWeightProxyCaller{contract: contract}, VotingWeightProxyTransactor: VotingWeightProxyTransactor{contract: contract}, VotingWeightProxyFilterer: VotingWeightProxyFilterer{contract: contract}}, nil
}

// VotingWeightProxy is an auto generated Go binding around an Ethereum contract.
type VotingWeightProxy struct {
	VotingWeightProxyCaller     // Read-only binding to the contract
	VotingWeightProxyTransactor // Write-only binding to the contract
	VotingWeightProxyFilterer   // Log filterer for contract events
}

// VotingWeightProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingWeightProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingWeightProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingWeightProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingWeightProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingWeightProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingWeightProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingWeightProxySession struct {
	Contract     *VotingWeightProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VotingWeightProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingWeightProxyCallerSession struct {
	Contract *VotingWeightProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// VotingWeightProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingWeightProxyTransactorSession struct {
	Contract     *VotingWeightProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// VotingWeightProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingWeightProxyRaw struct {
	Contract *VotingWeightProxy // Generic contract binding to access the raw methods on
}

// VotingWeightProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingWeightProxyCallerRaw struct {
	Contract *VotingWeightProxyCaller // Generic read-only contract binding to access the raw methods on
}

// VotingWeightProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingWeightProxyTransactorRaw struct {
	Contract *VotingWeightProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingWeightProxy creates a new instance of VotingWeightProxy, bound to a specific deployed contract.
func NewVotingWeightProxy(address common.Address, backend bind.ContractBackend) (*VotingWeightProxy, error) {
	contract, err := bindVotingWeightProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxy{VotingWeightProxyCaller: VotingWeightProxyCaller{contract: contract}, VotingWeightProxyTransactor: VotingWeightProxyTransactor{contract: contract}, VotingWeightProxyFilterer: VotingWeightProxyFilterer{contract: contract}}, nil
}

// NewVotingWeightProxyCaller creates a new read-only instance of VotingWeightProxy, bound to a specific deployed contract.
func NewVotingWeightProxyCaller(address common.Address, caller bind.ContractCaller) (*VotingWeightProxyCaller, error) {
	contract, err := bindVotingWeightProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyCaller{contract: contract}, nil
}

// NewVotingWeightProxyTransactor creates a new write-only instance of VotingWeightProxy, bound to a specific deployed contract.
func NewVotingWeightProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingWeightProxyTransactor, error) {
	contract, err := bindVotingWeightProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyTransactor{contract: contract}, nil
}

// NewVotingWeightProxyFilterer creates a new log filterer instance of VotingWeightProxy, bound to a specific deployed contract.
func NewVotingWeightProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingWeightProxyFilterer, error) {
	contract, err := bindVotingWeightProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyFilterer{contract: contract}, nil
}

// bindVotingWeightProxy binds a generic wrapper to an already deployed contract.
func bindVotingWeightProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingWeightProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingWeightProxy *VotingWeightProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingWeightProxy.Contract.VotingWeightProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingWeightProxy *VotingWeightProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.VotingWeightProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingWeightProxy *VotingWeightProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.VotingWeightProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingWeightProxy *VotingWeightProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VotingWeightProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingWeightProxy *VotingWeightProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingWeightProxy *VotingWeightProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.contract.Transact(opts, method, params...)
}

// DelegationInfos is a free data retrieval call binding the contract method 0xd1e0fea1.
//
// Solidity: function delegationInfos(address ) view returns(uint256 receivedWeight, address votingAgent, bool isPending, uint256 votingAgentPassOverTime)
func (_VotingWeightProxy *VotingWeightProxyCaller) DelegationInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	ReceivedWeight          *big.Int
	VotingAgent             common.Address
	IsPending               bool
	VotingAgentPassOverTime *big.Int
}, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "delegationInfos", arg0)

	outstruct := new(struct {
		ReceivedWeight          *big.Int
		VotingAgent             common.Address
		IsPending               bool
		VotingAgentPassOverTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceivedWeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VotingAgent = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.IsPending = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.VotingAgentPassOverTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegationInfos is a free data retrieval call binding the contract method 0xd1e0fea1.
//
// Solidity: function delegationInfos(address ) view returns(uint256 receivedWeight, address votingAgent, bool isPending, uint256 votingAgentPassOverTime)
func (_VotingWeightProxy *VotingWeightProxySession) DelegationInfos(arg0 common.Address) (struct {
	ReceivedWeight          *big.Int
	VotingAgent             common.Address
	IsPending               bool
	VotingAgentPassOverTime *big.Int
}, error) {
	return _VotingWeightProxy.Contract.DelegationInfos(&_VotingWeightProxy.CallOpts, arg0)
}

// DelegationInfos is a free data retrieval call binding the contract method 0xd1e0fea1.
//
// Solidity: function delegationInfos(address ) view returns(uint256 receivedWeight, address votingAgent, bool isPending, uint256 votingAgentPassOverTime)
func (_VotingWeightProxy *VotingWeightProxyCallerSession) DelegationInfos(arg0 common.Address) (struct {
	ReceivedWeight          *big.Int
	VotingAgent             common.Address
	IsPending               bool
	VotingAgentPassOverTime *big.Int
}, error) {
	return _VotingWeightProxy.Contract.DelegationInfos(&_VotingWeightProxy.CallOpts, arg0)
}

// GetBaseVotingWeightInfo is a free data retrieval call binding the contract method 0x10874b5b.
//
// Solidity: function getBaseVotingWeightInfo(address _user, uint256 _votingEndTime) view returns((uint256,address,uint8,uint256))
func (_VotingWeightProxy *VotingWeightProxyCaller) GetBaseVotingWeightInfo(opts *bind.CallOpts, _user common.Address, _votingEndTime *big.Int) (BaseVotingWeightInfo, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "getBaseVotingWeightInfo", _user, _votingEndTime)

	if err != nil {
		return *new(BaseVotingWeightInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseVotingWeightInfo)).(*BaseVotingWeightInfo)

	return out0, err

}

// GetBaseVotingWeightInfo is a free data retrieval call binding the contract method 0x10874b5b.
//
// Solidity: function getBaseVotingWeightInfo(address _user, uint256 _votingEndTime) view returns((uint256,address,uint8,uint256))
func (_VotingWeightProxy *VotingWeightProxySession) GetBaseVotingWeightInfo(_user common.Address, _votingEndTime *big.Int) (BaseVotingWeightInfo, error) {
	return _VotingWeightProxy.Contract.GetBaseVotingWeightInfo(&_VotingWeightProxy.CallOpts, _user, _votingEndTime)
}

// GetBaseVotingWeightInfo is a free data retrieval call binding the contract method 0x10874b5b.
//
// Solidity: function getBaseVotingWeightInfo(address _user, uint256 _votingEndTime) view returns((uint256,address,uint8,uint256))
func (_VotingWeightProxy *VotingWeightProxyCallerSession) GetBaseVotingWeightInfo(_user common.Address, _votingEndTime *big.Int) (BaseVotingWeightInfo, error) {
	return _VotingWeightProxy.Contract.GetBaseVotingWeightInfo(&_VotingWeightProxy.CallOpts, _user, _votingEndTime)
}

// GetLockInfo is a free data retrieval call binding the contract method 0x061c9040.
//
// Solidity: function getLockInfo(address _tokenLockSource, address _who) view returns((uint256,uint256,uint256,uint256))
func (_VotingWeightProxy *VotingWeightProxyCaller) GetLockInfo(opts *bind.CallOpts, _tokenLockSource common.Address, _who common.Address) (VotingLockInfo, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "getLockInfo", _tokenLockSource, _who)

	if err != nil {
		return *new(VotingLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VotingLockInfo)).(*VotingLockInfo)

	return out0, err

}

// GetLockInfo is a free data retrieval call binding the contract method 0x061c9040.
//
// Solidity: function getLockInfo(address _tokenLockSource, address _who) view returns((uint256,uint256,uint256,uint256))
func (_VotingWeightProxy *VotingWeightProxySession) GetLockInfo(_tokenLockSource common.Address, _who common.Address) (VotingLockInfo, error) {
	return _VotingWeightProxy.Contract.GetLockInfo(&_VotingWeightProxy.CallOpts, _tokenLockSource, _who)
}

// GetLockInfo is a free data retrieval call binding the contract method 0x061c9040.
//
// Solidity: function getLockInfo(address _tokenLockSource, address _who) view returns((uint256,uint256,uint256,uint256))
func (_VotingWeightProxy *VotingWeightProxyCallerSession) GetLockInfo(_tokenLockSource common.Address, _who common.Address) (VotingLockInfo, error) {
	return _VotingWeightProxy.Contract.GetLockInfo(&_VotingWeightProxy.CallOpts, _tokenLockSource, _who)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x929ec537.
//
// Solidity: function getLockedAmount(address _who) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCaller) GetLockedAmount(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "getLockedAmount", _who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedAmount is a free data retrieval call binding the contract method 0x929ec537.
//
// Solidity: function getLockedAmount(address _who) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxySession) GetLockedAmount(_who common.Address) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetLockedAmount(&_VotingWeightProxy.CallOpts, _who)
}

// GetLockedAmount is a free data retrieval call binding the contract method 0x929ec537.
//
// Solidity: function getLockedAmount(address _who) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCallerSession) GetLockedAmount(_who common.Address) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetLockedAmount(&_VotingWeightProxy.CallOpts, _who)
}

// GetLockedUntil is a free data retrieval call binding the contract method 0x50696346.
//
// Solidity: function getLockedUntil(address _user) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCaller) GetLockedUntil(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "getLockedUntil", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedUntil is a free data retrieval call binding the contract method 0x50696346.
//
// Solidity: function getLockedUntil(address _user) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxySession) GetLockedUntil(_user common.Address) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetLockedUntil(&_VotingWeightProxy.CallOpts, _user)
}

// GetLockedUntil is a free data retrieval call binding the contract method 0x50696346.
//
// Solidity: function getLockedUntil(address _user) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCallerSession) GetLockedUntil(_user common.Address) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetLockedUntil(&_VotingWeightProxy.CallOpts, _user)
}

// GetVotingWeight is a free data retrieval call binding the contract method 0x4cb77f97.
//
// Solidity: function getVotingWeight(address _who, uint256 _lockNeededUntil) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCaller) GetVotingWeight(opts *bind.CallOpts, _who common.Address, _lockNeededUntil *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "getVotingWeight", _who, _lockNeededUntil)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingWeight is a free data retrieval call binding the contract method 0x4cb77f97.
//
// Solidity: function getVotingWeight(address _who, uint256 _lockNeededUntil) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxySession) GetVotingWeight(_who common.Address, _lockNeededUntil *big.Int) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetVotingWeight(&_VotingWeightProxy.CallOpts, _who, _lockNeededUntil)
}

// GetVotingWeight is a free data retrieval call binding the contract method 0x4cb77f97.
//
// Solidity: function getVotingWeight(address _who, uint256 _lockNeededUntil) view returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyCallerSession) GetVotingWeight(_who common.Address, _lockNeededUntil *big.Int) (*big.Int, error) {
	return _VotingWeightProxy.Contract.GetVotingWeight(&_VotingWeightProxy.CallOpts, _who, _lockNeededUntil)
}

// LockInfos is a free data retrieval call binding the contract method 0x2f429474.
//
// Solidity: function lockInfos(address , address ) view returns(uint256 lockedAmount, uint256 lockedUntil, uint256 pendingUnlockAmount, uint256 pendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxyCaller) LockInfos(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	LockedAmount        *big.Int
	LockedUntil         *big.Int
	PendingUnlockAmount *big.Int
	PendingUnlockTime   *big.Int
}, error) {
	var out []interface{}
	err := _VotingWeightProxy.contract.Call(opts, &out, "lockInfos", arg0, arg1)

	outstruct := new(struct {
		LockedAmount        *big.Int
		LockedUntil         *big.Int
		PendingUnlockAmount *big.Int
		PendingUnlockTime   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockedUntil = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingUnlockAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingUnlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockInfos is a free data retrieval call binding the contract method 0x2f429474.
//
// Solidity: function lockInfos(address , address ) view returns(uint256 lockedAmount, uint256 lockedUntil, uint256 pendingUnlockAmount, uint256 pendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxySession) LockInfos(arg0 common.Address, arg1 common.Address) (struct {
	LockedAmount        *big.Int
	LockedUntil         *big.Int
	PendingUnlockAmount *big.Int
	PendingUnlockTime   *big.Int
}, error) {
	return _VotingWeightProxy.Contract.LockInfos(&_VotingWeightProxy.CallOpts, arg0, arg1)
}

// LockInfos is a free data retrieval call binding the contract method 0x2f429474.
//
// Solidity: function lockInfos(address , address ) view returns(uint256 lockedAmount, uint256 lockedUntil, uint256 pendingUnlockAmount, uint256 pendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxyCallerSession) LockInfos(arg0 common.Address, arg1 common.Address) (struct {
	LockedAmount        *big.Int
	LockedUntil         *big.Int
	PendingUnlockAmount *big.Int
	PendingUnlockTime   *big.Int
}, error) {
	return _VotingWeightProxy.Contract.LockInfos(&_VotingWeightProxy.CallOpts, arg0, arg1)
}

// AnnounceNewVotingAgent is a paid mutator transaction binding the contract method 0x6b164f7d.
//
// Solidity: function announceNewVotingAgent(address _newAgent) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) AnnounceNewVotingAgent(opts *bind.TransactOpts, _newAgent common.Address) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "announceNewVotingAgent", _newAgent)
}

// AnnounceNewVotingAgent is a paid mutator transaction binding the contract method 0x6b164f7d.
//
// Solidity: function announceNewVotingAgent(address _newAgent) returns()
func (_VotingWeightProxy *VotingWeightProxySession) AnnounceNewVotingAgent(_newAgent common.Address) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.AnnounceNewVotingAgent(&_VotingWeightProxy.TransactOpts, _newAgent)
}

// AnnounceNewVotingAgent is a paid mutator transaction binding the contract method 0x6b164f7d.
//
// Solidity: function announceNewVotingAgent(address _newAgent) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) AnnounceNewVotingAgent(_newAgent common.Address) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.AnnounceNewVotingAgent(&_VotingWeightProxy.TransactOpts, _newAgent)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xe97d651c.
//
// Solidity: function announceUnlock(address _who, uint256 _amount, uint256 _newUnlockTime) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) AnnounceUnlock(opts *bind.TransactOpts, _who common.Address, _amount *big.Int, _newUnlockTime *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "announceUnlock", _who, _amount, _newUnlockTime)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xe97d651c.
//
// Solidity: function announceUnlock(address _who, uint256 _amount, uint256 _newUnlockTime) returns()
func (_VotingWeightProxy *VotingWeightProxySession) AnnounceUnlock(_who common.Address, _amount *big.Int, _newUnlockTime *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.AnnounceUnlock(&_VotingWeightProxy.TransactOpts, _who, _amount, _newUnlockTime)
}

// AnnounceUnlock is a paid mutator transaction binding the contract method 0xe97d651c.
//
// Solidity: function announceUnlock(address _who, uint256 _amount, uint256 _newUnlockTime) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) AnnounceUnlock(_who common.Address, _amount *big.Int, _newUnlockTime *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.AnnounceUnlock(&_VotingWeightProxy.TransactOpts, _who, _amount, _newUnlockTime)
}

// ExtendLocking is a paid mutator transaction binding the contract method 0x7bf9f918.
//
// Solidity: function extendLocking(address _who, uint256 _lockNeededUntil) returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyTransactor) ExtendLocking(opts *bind.TransactOpts, _who common.Address, _lockNeededUntil *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "extendLocking", _who, _lockNeededUntil)
}

// ExtendLocking is a paid mutator transaction binding the contract method 0x7bf9f918.
//
// Solidity: function extendLocking(address _who, uint256 _lockNeededUntil) returns(uint256)
func (_VotingWeightProxy *VotingWeightProxySession) ExtendLocking(_who common.Address, _lockNeededUntil *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.ExtendLocking(&_VotingWeightProxy.TransactOpts, _who, _lockNeededUntil)
}

// ExtendLocking is a paid mutator transaction binding the contract method 0x7bf9f918.
//
// Solidity: function extendLocking(address _who, uint256 _lockNeededUntil) returns(uint256)
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) ExtendLocking(_who common.Address, _lockNeededUntil *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.ExtendLocking(&_VotingWeightProxy.TransactOpts, _who, _lockNeededUntil)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0x5c7700df.
//
// Solidity: function forceUnlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) ForceUnlock(opts *bind.TransactOpts, _who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "forceUnlock", _who, _amount)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0x5c7700df.
//
// Solidity: function forceUnlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxySession) ForceUnlock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.ForceUnlock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// ForceUnlock is a paid mutator transaction binding the contract method 0x5c7700df.
//
// Solidity: function forceUnlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) ForceUnlock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.ForceUnlock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _tokenLockSourcesKeys, string[] _votingContractsKeys) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _tokenLockSourcesKeys []string, _votingContractsKeys []string) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "initialize", _registry, _tokenLockSourcesKeys, _votingContractsKeys)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _tokenLockSourcesKeys, string[] _votingContractsKeys) returns()
func (_VotingWeightProxy *VotingWeightProxySession) Initialize(_registry common.Address, _tokenLockSourcesKeys []string, _votingContractsKeys []string) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Initialize(&_VotingWeightProxy.TransactOpts, _registry, _tokenLockSourcesKeys, _votingContractsKeys)
}

// Initialize is a paid mutator transaction binding the contract method 0xefe86a5e.
//
// Solidity: function initialize(address _registry, string[] _tokenLockSourcesKeys, string[] _votingContractsKeys) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) Initialize(_registry common.Address, _tokenLockSourcesKeys []string, _votingContractsKeys []string) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Initialize(&_VotingWeightProxy.TransactOpts, _registry, _tokenLockSourcesKeys, _votingContractsKeys)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) Lock(opts *bind.TransactOpts, _who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "lock", _who, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxySession) Lock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Lock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) Lock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Lock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// SetNewVotingAgent is a paid mutator transaction binding the contract method 0xf927c6f2.
//
// Solidity: function setNewVotingAgent() returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) SetNewVotingAgent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "setNewVotingAgent")
}

// SetNewVotingAgent is a paid mutator transaction binding the contract method 0xf927c6f2.
//
// Solidity: function setNewVotingAgent() returns()
func (_VotingWeightProxy *VotingWeightProxySession) SetNewVotingAgent() (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.SetNewVotingAgent(&_VotingWeightProxy.TransactOpts)
}

// SetNewVotingAgent is a paid mutator transaction binding the contract method 0xf927c6f2.
//
// Solidity: function setNewVotingAgent() returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) SetNewVotingAgent() (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.SetNewVotingAgent(&_VotingWeightProxy.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactor) Unlock(opts *bind.TransactOpts, _who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.contract.Transact(opts, "unlock", _who, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxySession) Unlock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Unlock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x7eee288d.
//
// Solidity: function unlock(address _who, uint256 _amount) returns()
func (_VotingWeightProxy *VotingWeightProxyTransactorSession) Unlock(_who common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VotingWeightProxy.Contract.Unlock(&_VotingWeightProxy.TransactOpts, _who, _amount)
}

// VotingWeightProxyLockedAmountChangedIterator is returned from FilterLockedAmountChanged and is used to iterate over the raw logs and unpacked data for LockedAmountChanged events raised by the VotingWeightProxy contract.
type VotingWeightProxyLockedAmountChangedIterator struct {
	Event *VotingWeightProxyLockedAmountChanged // Event containing the contract specifics and raw log

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
func (it *VotingWeightProxyLockedAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingWeightProxyLockedAmountChanged)
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
		it.Event = new(VotingWeightProxyLockedAmountChanged)
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
func (it *VotingWeightProxyLockedAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingWeightProxyLockedAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingWeightProxyLockedAmountChanged represents a LockedAmountChanged event raised by the VotingWeightProxy contract.
type VotingWeightProxyLockedAmountChanged struct {
	TokenLockSource common.Address
	Who             common.Address
	NewLockedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockedAmountChanged is a free log retrieval operation binding the contract event 0x48f899e185e3470300612fd03b36ae16768c2212cc0fdc94f44cf5d60bf3227e.
//
// Solidity: event LockedAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newLockedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) FilterLockedAmountChanged(opts *bind.FilterOpts, _who []common.Address) (*VotingWeightProxyLockedAmountChangedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.FilterLogs(opts, "LockedAmountChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyLockedAmountChangedIterator{contract: _VotingWeightProxy.contract, event: "LockedAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLockedAmountChanged is a free log subscription operation binding the contract event 0x48f899e185e3470300612fd03b36ae16768c2212cc0fdc94f44cf5d60bf3227e.
//
// Solidity: event LockedAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newLockedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) WatchLockedAmountChanged(opts *bind.WatchOpts, sink chan<- *VotingWeightProxyLockedAmountChanged, _who []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.WatchLogs(opts, "LockedAmountChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingWeightProxyLockedAmountChanged)
				if err := _VotingWeightProxy.contract.UnpackLog(event, "LockedAmountChanged", log); err != nil {
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

// ParseLockedAmountChanged is a log parse operation binding the contract event 0x48f899e185e3470300612fd03b36ae16768c2212cc0fdc94f44cf5d60bf3227e.
//
// Solidity: event LockedAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newLockedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) ParseLockedAmountChanged(log types.Log) (*VotingWeightProxyLockedAmountChanged, error) {
	event := new(VotingWeightProxyLockedAmountChanged)
	if err := _VotingWeightProxy.contract.UnpackLog(event, "LockedAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingWeightProxyNewVotingAgentAnnouncedIterator is returned from FilterNewVotingAgentAnnounced and is used to iterate over the raw logs and unpacked data for NewVotingAgentAnnounced events raised by the VotingWeightProxy contract.
type VotingWeightProxyNewVotingAgentAnnouncedIterator struct {
	Event *VotingWeightProxyNewVotingAgentAnnounced // Event containing the contract specifics and raw log

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
func (it *VotingWeightProxyNewVotingAgentAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingWeightProxyNewVotingAgentAnnounced)
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
		it.Event = new(VotingWeightProxyNewVotingAgentAnnounced)
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
func (it *VotingWeightProxyNewVotingAgentAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingWeightProxyNewVotingAgentAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingWeightProxyNewVotingAgentAnnounced represents a NewVotingAgentAnnounced event raised by the VotingWeightProxy contract.
type VotingWeightProxyNewVotingAgentAnnounced struct {
	Who             common.Address
	PrevVotingAgent common.Address
	NewVotingAgent  common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVotingAgentAnnounced is a free log retrieval operation binding the contract event 0x3b955a76d9960f085ab98c0fad7fad089e2fc0506bcb86b069c300112625df9d.
//
// Solidity: event NewVotingAgentAnnounced(address indexed _who, address _prevVotingAgent, address _newVotingAgent)
func (_VotingWeightProxy *VotingWeightProxyFilterer) FilterNewVotingAgentAnnounced(opts *bind.FilterOpts, _who []common.Address) (*VotingWeightProxyNewVotingAgentAnnouncedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.FilterLogs(opts, "NewVotingAgentAnnounced", _whoRule)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyNewVotingAgentAnnouncedIterator{contract: _VotingWeightProxy.contract, event: "NewVotingAgentAnnounced", logs: logs, sub: sub}, nil
}

// WatchNewVotingAgentAnnounced is a free log subscription operation binding the contract event 0x3b955a76d9960f085ab98c0fad7fad089e2fc0506bcb86b069c300112625df9d.
//
// Solidity: event NewVotingAgentAnnounced(address indexed _who, address _prevVotingAgent, address _newVotingAgent)
func (_VotingWeightProxy *VotingWeightProxyFilterer) WatchNewVotingAgentAnnounced(opts *bind.WatchOpts, sink chan<- *VotingWeightProxyNewVotingAgentAnnounced, _who []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.WatchLogs(opts, "NewVotingAgentAnnounced", _whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingWeightProxyNewVotingAgentAnnounced)
				if err := _VotingWeightProxy.contract.UnpackLog(event, "NewVotingAgentAnnounced", log); err != nil {
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

// ParseNewVotingAgentAnnounced is a log parse operation binding the contract event 0x3b955a76d9960f085ab98c0fad7fad089e2fc0506bcb86b069c300112625df9d.
//
// Solidity: event NewVotingAgentAnnounced(address indexed _who, address _prevVotingAgent, address _newVotingAgent)
func (_VotingWeightProxy *VotingWeightProxyFilterer) ParseNewVotingAgentAnnounced(log types.Log) (*VotingWeightProxyNewVotingAgentAnnounced, error) {
	event := new(VotingWeightProxyNewVotingAgentAnnounced)
	if err := _VotingWeightProxy.contract.UnpackLog(event, "NewVotingAgentAnnounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingWeightProxyPendingUnlockAmountChangedIterator is returned from FilterPendingUnlockAmountChanged and is used to iterate over the raw logs and unpacked data for PendingUnlockAmountChanged events raised by the VotingWeightProxy contract.
type VotingWeightProxyPendingUnlockAmountChangedIterator struct {
	Event *VotingWeightProxyPendingUnlockAmountChanged // Event containing the contract specifics and raw log

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
func (it *VotingWeightProxyPendingUnlockAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingWeightProxyPendingUnlockAmountChanged)
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
		it.Event = new(VotingWeightProxyPendingUnlockAmountChanged)
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
func (it *VotingWeightProxyPendingUnlockAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingWeightProxyPendingUnlockAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingWeightProxyPendingUnlockAmountChanged represents a PendingUnlockAmountChanged event raised by the VotingWeightProxy contract.
type VotingWeightProxyPendingUnlockAmountChanged struct {
	TokenLockSource        common.Address
	Who                    common.Address
	NewPendingUnlockAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPendingUnlockAmountChanged is a free log retrieval operation binding the contract event 0x70601d6dee302c74325828267b70767d970eddb01567778dfc92cd02aa7671bd.
//
// Solidity: event PendingUnlockAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) FilterPendingUnlockAmountChanged(opts *bind.FilterOpts, _who []common.Address) (*VotingWeightProxyPendingUnlockAmountChangedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.FilterLogs(opts, "PendingUnlockAmountChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyPendingUnlockAmountChangedIterator{contract: _VotingWeightProxy.contract, event: "PendingUnlockAmountChanged", logs: logs, sub: sub}, nil
}

// WatchPendingUnlockAmountChanged is a free log subscription operation binding the contract event 0x70601d6dee302c74325828267b70767d970eddb01567778dfc92cd02aa7671bd.
//
// Solidity: event PendingUnlockAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) WatchPendingUnlockAmountChanged(opts *bind.WatchOpts, sink chan<- *VotingWeightProxyPendingUnlockAmountChanged, _who []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.WatchLogs(opts, "PendingUnlockAmountChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingWeightProxyPendingUnlockAmountChanged)
				if err := _VotingWeightProxy.contract.UnpackLog(event, "PendingUnlockAmountChanged", log); err != nil {
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

// ParsePendingUnlockAmountChanged is a log parse operation binding the contract event 0x70601d6dee302c74325828267b70767d970eddb01567778dfc92cd02aa7671bd.
//
// Solidity: event PendingUnlockAmountChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) ParsePendingUnlockAmountChanged(log types.Log) (*VotingWeightProxyPendingUnlockAmountChanged, error) {
	event := new(VotingWeightProxyPendingUnlockAmountChanged)
	if err := _VotingWeightProxy.contract.UnpackLog(event, "PendingUnlockAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingWeightProxyPendingUnlockChangedIterator is returned from FilterPendingUnlockChanged and is used to iterate over the raw logs and unpacked data for PendingUnlockChanged events raised by the VotingWeightProxy contract.
type VotingWeightProxyPendingUnlockChangedIterator struct {
	Event *VotingWeightProxyPendingUnlockChanged // Event containing the contract specifics and raw log

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
func (it *VotingWeightProxyPendingUnlockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingWeightProxyPendingUnlockChanged)
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
		it.Event = new(VotingWeightProxyPendingUnlockChanged)
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
func (it *VotingWeightProxyPendingUnlockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingWeightProxyPendingUnlockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingWeightProxyPendingUnlockChanged represents a PendingUnlockChanged event raised by the VotingWeightProxy contract.
type VotingWeightProxyPendingUnlockChanged struct {
	TokenLockSource        common.Address
	Who                    common.Address
	NewPendingUnlockAmount *big.Int
	NewPendingUnlockTime   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPendingUnlockChanged is a free log retrieval operation binding the contract event 0x531e4f82f00c1943a0c7d64474e788195d4cf400182d1dbf9d7a3fd9787a4d45.
//
// Solidity: event PendingUnlockChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount, uint256 _newPendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxyFilterer) FilterPendingUnlockChanged(opts *bind.FilterOpts, _who []common.Address) (*VotingWeightProxyPendingUnlockChangedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.FilterLogs(opts, "PendingUnlockChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyPendingUnlockChangedIterator{contract: _VotingWeightProxy.contract, event: "PendingUnlockChanged", logs: logs, sub: sub}, nil
}

// WatchPendingUnlockChanged is a free log subscription operation binding the contract event 0x531e4f82f00c1943a0c7d64474e788195d4cf400182d1dbf9d7a3fd9787a4d45.
//
// Solidity: event PendingUnlockChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount, uint256 _newPendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxyFilterer) WatchPendingUnlockChanged(opts *bind.WatchOpts, sink chan<- *VotingWeightProxyPendingUnlockChanged, _who []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.WatchLogs(opts, "PendingUnlockChanged", _whoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingWeightProxyPendingUnlockChanged)
				if err := _VotingWeightProxy.contract.UnpackLog(event, "PendingUnlockChanged", log); err != nil {
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

// ParsePendingUnlockChanged is a log parse operation binding the contract event 0x531e4f82f00c1943a0c7d64474e788195d4cf400182d1dbf9d7a3fd9787a4d45.
//
// Solidity: event PendingUnlockChanged(address _tokenLockSource, address indexed _who, uint256 _newPendingUnlockAmount, uint256 _newPendingUnlockTime)
func (_VotingWeightProxy *VotingWeightProxyFilterer) ParsePendingUnlockChanged(log types.Log) (*VotingWeightProxyPendingUnlockChanged, error) {
	event := new(VotingWeightProxyPendingUnlockChanged)
	if err := _VotingWeightProxy.contract.UnpackLog(event, "PendingUnlockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingWeightProxyVotingAgentChangedIterator is returned from FilterVotingAgentChanged and is used to iterate over the raw logs and unpacked data for VotingAgentChanged events raised by the VotingWeightProxy contract.
type VotingWeightProxyVotingAgentChangedIterator struct {
	Event *VotingWeightProxyVotingAgentChanged // Event containing the contract specifics and raw log

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
func (it *VotingWeightProxyVotingAgentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingWeightProxyVotingAgentChanged)
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
		it.Event = new(VotingWeightProxyVotingAgentChanged)
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
func (it *VotingWeightProxyVotingAgentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingWeightProxyVotingAgentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingWeightProxyVotingAgentChanged represents a VotingAgentChanged event raised by the VotingWeightProxy contract.
type VotingWeightProxyVotingAgentChanged struct {
	Who             common.Address
	VotingAgent     common.Address
	DelegatedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingAgentChanged is a free log retrieval operation binding the contract event 0x395be45122e29f50a1f27a1ffd8913eac762e146c7f7214d01b52f2bc1dd3487.
//
// Solidity: event VotingAgentChanged(address indexed _who, address indexed _votingAgent, uint256 _delegatedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) FilterVotingAgentChanged(opts *bind.FilterOpts, _who []common.Address, _votingAgent []common.Address) (*VotingWeightProxyVotingAgentChangedIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _votingAgentRule []interface{}
	for _, _votingAgentItem := range _votingAgent {
		_votingAgentRule = append(_votingAgentRule, _votingAgentItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.FilterLogs(opts, "VotingAgentChanged", _whoRule, _votingAgentRule)
	if err != nil {
		return nil, err
	}
	return &VotingWeightProxyVotingAgentChangedIterator{contract: _VotingWeightProxy.contract, event: "VotingAgentChanged", logs: logs, sub: sub}, nil
}

// WatchVotingAgentChanged is a free log subscription operation binding the contract event 0x395be45122e29f50a1f27a1ffd8913eac762e146c7f7214d01b52f2bc1dd3487.
//
// Solidity: event VotingAgentChanged(address indexed _who, address indexed _votingAgent, uint256 _delegatedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) WatchVotingAgentChanged(opts *bind.WatchOpts, sink chan<- *VotingWeightProxyVotingAgentChanged, _who []common.Address, _votingAgent []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _votingAgentRule []interface{}
	for _, _votingAgentItem := range _votingAgent {
		_votingAgentRule = append(_votingAgentRule, _votingAgentItem)
	}

	logs, sub, err := _VotingWeightProxy.contract.WatchLogs(opts, "VotingAgentChanged", _whoRule, _votingAgentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingWeightProxyVotingAgentChanged)
				if err := _VotingWeightProxy.contract.UnpackLog(event, "VotingAgentChanged", log); err != nil {
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

// ParseVotingAgentChanged is a log parse operation binding the contract event 0x395be45122e29f50a1f27a1ffd8913eac762e146c7f7214d01b52f2bc1dd3487.
//
// Solidity: event VotingAgentChanged(address indexed _who, address indexed _votingAgent, uint256 _delegatedAmount)
func (_VotingWeightProxy *VotingWeightProxyFilterer) ParseVotingAgentChanged(log types.Log) (*VotingWeightProxyVotingAgentChanged, error) {
	event := new(VotingWeightProxyVotingAgentChanged)
	if err := _VotingWeightProxy.contract.UnpackLog(event, "VotingAgentChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
