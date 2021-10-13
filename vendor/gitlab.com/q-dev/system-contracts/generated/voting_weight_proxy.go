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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newLockedAmount\",\"type\":\"uint256\"}],\"name\":\"LockedAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prevVotingAgent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newVotingAgent\",\"type\":\"address\"}],\"name\":\"NewVotingAgentAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockAmount\",\"type\":\"uint256\"}],\"name\":\"PendingUnlockAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newPendingUnlockTime\",\"type\":\"uint256\"}],\"name\":\"PendingUnlockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_votingAgent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_delegatedAmount\",\"type\":\"uint256\"}],\"name\":\"VotingAgentChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegationInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receivedWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPending\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votingAgentPassOverTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_tokenLockSourcesKeys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_votingContractsKeys\",\"type\":\"string[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newUnlockTime\",\"type\":\"uint256\"}],\"name\":\"announceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"forceUnlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAgent\",\"type\":\"address\"}],\"name\":\"announceNewVotingAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setNewVotingAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockNeededUntil\",\"type\":\"uint256\"}],\"name\":\"extendLocking\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenLockSource\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingUnlockTime\",\"type\":\"uint256\"}],\"internalType\":\"structVotingLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_votingEndTime\",\"type\":\"uint256\"}],\"name\":\"getBaseVotingWeightInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ownWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"enumDelegationStatus\",\"name\":\"delegationStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVotingWeightInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getLockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockNeededUntil\",\"type\":\"uint256\"}],\"name\":\"getVotingWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612437806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c55760003560e01c8063061c9040146100ca57806310874b5b146100f3578063282d3fdf146101135780632f429474146101285780634cb77f971461014b578063506963461461016b5780635c7700df1461017e5780636b164f7d146101915780637bf9f918146101a45780637eee288d146101b7578063929ec537146101ca578063d1e0fea1146101dd578063e97d651c14610200578063efe86a5e14610213578063f927c6f214610226575b600080fd5b6100dd6100d8366004611c2a565b61022e565b6040516100ea91906122d2565b60405180910390f35b610106610101366004611cd5565b6103e0565b6040516100ea9190612290565b610126610121366004611cd5565b6104d5565b005b61013b610136366004611c2a565b610744565b6040516100ea949392919061232c565b61015e610159366004611cd5565b610775565b6040516100ea91906122fd565b61015e610179366004611bf2565b61093f565b61012661018c366004611cd5565b61095a565b61012661019f366004611bf2565b610b42565b61015e6101b2366004611cd5565b610c35565b6101266101c5366004611cd5565b610d3f565b61015e6101d8366004611bf2565b610ef3565b6101f06101eb366004611bf2565b610ff1565b6040516100ea9493929190612306565b61012661020e366004611d00565b611028565b610126610221366004611c62565b61121d565b61012661130c565b610236611982565b6000805b60015481101561030557600054600180546001600160a01b03808916936201000090041691633fb90271918590811061026f57fe5b906000526020600020016040518263ffffffff1660e01b81526004016102959190611ddb565b60206040518083038186803b1580156102ad57600080fd5b505afa1580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e59190611c0e565b6001600160a01b031614156102fd5760019150610305565b60010161023a565b5060006040518060600160405280602781526020016123db602791399050808261034b5760405162461bcd60e51b81526004016103429190611d88565b60405180910390fd5b506001600160a01b0380861660009081526003602081815260408084209489168085529482528084208151608081018352815481526001820154818501908152600283015482850152919094015460608501529484526004909152909120549151909111156103d5576001600160a01b038516600090815260046020908152604090912054908201525b925050505b92915050565b6103e86119aa565b6001600160a01b03838116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff1615159181019190915260029091015460608201526104476119aa565b6104518585610775565b81526020808301516001600160a01b03169082018190526104719061093f565b606082015260408201511561049c576040810160025b9081600281111561049457fe5b9052506104cd565b846001600160a01b031681602001516001600160a01b031614156104c557604081016000610487565b600160408201525b949350505050565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb90271918590811061050b57fe5b906000526020600020016040518263ffffffff1660e01b81526004016105319190611ddb565b60206040518083038186803b15801561054957600080fd5b505afa15801561055d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105819190611c0e565b6001600160a01b0316141561071f57600082116105b05760405162461bcd60e51b815260040161034290612157565b3360009081526003602081815260408084206001600160a01b03881685528252928390208351608081018552815480825260018301549382019390935260028201549481019490945290910154606083015261060c908461149c565b81523360008181526003602081815260408084206001600160a01b038a1680865290835293819020865180825592870151600182015581870151600282015560608701519301929092559051919260008051602061238d833981519152926106749290611d34565b60405180910390a26001600160a01b03808516600090815260056020526040902060018101549091166106bf576001810180546001600160a01b0319166001600160a01b0387161790555b6001810154600160a01b900460ff166107175760018101546001600160a01b03166000908152600560205260409020546106f9908561149c565b60018201546001600160a01b03166000908152600560205260409020555b505050610740565b6001016104d8565b5060405162461bcd60e51b815260040161034290611ec2565b5050565b6003602081815260009384526040808520909152918352912080546001820154600283015492909301549092919084565b600061077f611982565b6001600160a01b038481166000908152600560209081526040808320815160808101835281548152600182015495861693810193909352600160a01b90940460ff161515908201526002909201546060830152805b60015481101561092757600360008060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716001858154811061081257fe5b906000526020600020016040518263ffffffff1660e01b81526004016108389190611ddb565b60206040518083038186803b15801561085057600080fd5b505afa158015610864573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108889190611c0e565b6001600160a01b03166001600160a01b031681526020019081526020016000206000886001600160a01b03166001600160a01b031681526020019081526020016000206040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820154815250509350858460600151111561091f57604084015161091c90839061149c565b91505b6001016107d4565b50815161093590829061149c565b9695505050505050565b6001600160a01b031660009081526004602052604090205490565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb90271918590811061099057fe5b906000526020600020016040518263ffffffff1660e01b81526004016109b69190611ddb565b60206040518083038186803b1580156109ce57600080fd5b505afa1580156109e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a069190611c0e565b6001600160a01b03161415610b3a5760008211610a355760405162461bcd60e51b815260040161034290612157565b3360009081526003602090815260408083206001600160a01b0387168452909152812060028101548154919291610a6b9161149c565b905083811015610a8d5760405162461bcd60e51b815260040161034290611fa5565b6002820154801515908511610ab5576002830154610aab90866114fb565b6002840155610ae3565b6000610ace8460020154876114fb90919063ffffffff16565b600060028601559050610ae1878261153d565b505b8015610b3157856001600160a01b03167f70601d6dee302c74325828267b70767d970eddb01567778dfc92cd02aa7671bd338560020154604051610b28929190611d34565b60405180910390a25b50505050610740565b60010161095d565b33600090815260056020526040902060018101546001600160a01b03811690600160a01b900460ff16610bd757610b9a610b7b33610ef3565b6001600160a01b038316600090815260056020526040902054906114fb565b6001600160a01b038216600090815260056020526040902055610bbc8161093f565b600283015560018201805460ff60a01b1916600160a01b1790555b6001820180546001600160a01b0319166001600160a01b03851617905560405133907f3b955a76d9960f085ab98c0fad7fad089e2fc0506bcb86b069c300112625df9d90610c289084908790611d6e565b60405180910390a2505050565b6000805b600254811015610d26576000546002805433926201000090046001600160a01b031691633fb902719185908110610c6c57fe5b906000526020600020016040518263ffffffff1660e01b8152600401610c929190611ddb565b60206040518083038186803b158015610caa57600080fd5b505afa158015610cbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce29190611c0e565b6001600160a01b03161415610d1e576001600160a01b0384166000908152600460205260409020839055610d168484610775565b9150506103da565b600101610c39565b5060405162461bcd60e51b8152600401610342906121b0565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb902719185908110610d7557fe5b906000526020600020016040518263ffffffff1660e01b8152600401610d9b9190611ddb565b60206040518083038186803b158015610db357600080fd5b505afa158015610dc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610deb9190611c0e565b6001600160a01b03161415610eeb5760008211610e1a5760405162461bcd60e51b815260040161034290612157565b3360009081526003602081815260408084206001600160a01b0388168552909152909120908101544211610e605760405162461bcd60e51b815260040161034290611f31565b60028101548154600091610e74919061149c565b905083811015610e965760405162461bcd60e51b815260040161034290611e65565b6002820154801515908511610eb4576002830154610aab90866114fb565b6001600160a01b0386166000908152600460205260409020544211610ab55760405162461bcd60e51b815260040161034290612219565b600101610d42565b600080805b600154811015610fea57610fe0600360008060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027160018681548110610f3957fe5b906000526020600020016040518263ffffffff1660e01b8152600401610f5f9190611ddb565b60206040518083038186803b158015610f7757600080fd5b505afa158015610f8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faf9190611c0e565b6001600160a01b0390811682526020808301939093526040918201600090812091891681529252902054839061149c565b9150600101610ef8565b5092915050565b60056020526000908152604090208054600182015460029092015490916001600160a01b03811691600160a01b90910460ff169084565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb90271918590811061105e57fe5b906000526020600020016040518263ffffffff1660e01b81526004016110849190611ddb565b60206040518083038186803b15801561109c57600080fd5b505afa1580156110b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d49190611c0e565b6001600160a01b03161415611210573360009081526003602090815260408083206001600160a01b03881684529091529020600281015484101561113657600281015460009061112490866114fb565b9050611130868261172c565b5061115d565b600061114f8260020154866114fb90919063ffffffff16565b905061115b868261153d565b505b600281018490556001600160a01b0380861660009081526004602081815260408084205460058352818520600101549095168452919052812054908183116111a557816111a7565b825b9050808610156111b5578095505b6003840186905560028401546040516001600160a01b038a16917f531e4f82f00c1943a0c7d64474e788195d4cf400182d1dbf9d7a3fd9787a4d45916111fe9133918b90611d4d565b60405180910390a25050505050611218565b60010161102b565b505050565b600054610100900460ff168061123657506112366118d4565b80611244575060005460ff16155b61127f5760405162461bcd60e51b815260040180806020018281038252602e8152602001806123ad602e913960400191505060405180910390fd5b600054610100900460ff161580156112aa576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0387160217905582516112df9060019060208601906119d2565b5081516112f39060029060208501906119d2565b508015611306576000805461ff00191690555b50505050565b3360009081526005602090815260409182902082516080810184528154815260018201546001600160a01b03811693820193909352600160a01b90920460ff1615159282018390526002015460608201529061137a5760405162461bcd60e51b81526004016103429061207d565b428160600151111561139e5760405162461bcd60e51b81526004016103429061200f565b600060408083018281523380845260056020908152928420855181559285015160018401805493511515600160a01b0260ff60a01b196001600160a01b03939093166001600160a01b0319909516949094179190911692909217909155606084015160029092019190915561141290610ef3565b6020808401516001600160a01b0381166000908152600590925260409091205491925090611440908361149c565b6001600160a01b0382166000818152600560205260409081902092909255905133907f395be45122e29f50a1f27a1ffd8913eac762e146c7f7214d01b52f2bc1dd34879061148f9086906122fd565b60405180910390a3505050565b6000828201838110156114f4576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b60006114f483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506118e5565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb90271918590811061157357fe5b906000526020600020016040518263ffffffff1660e01b81526004016115999190611ddb565b60206040518083038186803b1580156115b157600080fd5b505afa1580156115c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115e99190611c0e565b6001600160a01b03161415611724573360009081526003602090815260408083206001600160a01b03871684529091529020805483111561163c5760405162461bcd60e51b8152600401610342906120fe565b805461164890846114fb565b8082556040516001600160a01b0386169160008051602061238d83398151915291611674913391611d34565b60405180910390a26001600160a01b03848116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff1615159181018290526002909201546060830152610717576020808201516001600160a01b03166000908152600590915260409020546116fe90856114fb565b6020808301516001600160a01b0316600090815260059091526040902055505050610740565b600101611540565b60005b600154811015610727576000546001805433926201000090046001600160a01b031691633fb90271918590811061176257fe5b906000526020600020016040518263ffffffff1660e01b81526004016117889190611ddb565b60206040518083038186803b1580156117a057600080fd5b505afa1580156117b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d89190611c0e565b6001600160a01b031614156118cc573360009081526003602090815260408083206001600160a01b038716845290915290208054611816908461149c565b8082556040516001600160a01b0386169160008051602061238d83398151915291611842913391611d34565b60405180910390a26001600160a01b03848116600090815260056020908152604091829020825160808101845281548152600182015494851692810192909252600160a01b90930460ff1615159181018290526002909201546060830152610717576020808201516001600160a01b03166000908152600590915260409020546116fe908561149c565b60010161172f565b60006118df3061197c565b15905090565b600081848411156119745760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611939578181015183820152602001611921565b50505050905090810190601f1680156119665780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b3b151590565b6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260008082526020820181905290918201908152602001600081525090565b828054828255906000526020600020908101928215611a1f579160200282015b82811115611a1f5782518051611a0f918491602090910190611a2f565b50916020019190600101906119f2565b50611a2b929150611ab7565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611a655760008555611aab565b82601f10611a7e57805160ff1916838001178555611aab565b82800160010185558215611aab579182015b82811115611aab578251825591602001919060010190611a90565b50611a2b929150611ad4565b80821115611a2b576000611acb8282611ae9565b50600101611ab7565b5b80821115611a2b5760008155600101611ad5565b50805460018160011615610100020316600290046000825580601f10611b0f5750611b2d565b601f016020900490600052602060002090810190611b2d9190611ad4565b50565b6000601f8381840112611b41578182fd5b8235602067ffffffffffffffff80831115611b5857fe5b611b658283850201612347565b83815282810190878401875b86811015611be35781358a018b603f820112611b8b57898afd5b86810135604087821115611b9b57fe5b611bac828c01601f19168a01612347565b8281528e82848601011115611bbf578c8dfd5b828285018b83013791820189018c9052508552509285019290850190600101611b71565b50909998505050505050505050565b600060208284031215611c03578081fd5b81356114f481612377565b600060208284031215611c1f578081fd5b81516114f481612377565b60008060408385031215611c3c578081fd5b8235611c4781612377565b91506020830135611c5781612377565b809150509250929050565b600080600060608486031215611c76578081fd5b8335611c8181612377565b9250602084013567ffffffffffffffff80821115611c9d578283fd5b611ca987838801611b30565b93506040860135915080821115611cbe578283fd5b50611ccb86828701611b30565b9150509250925092565b60008060408385031215611ce7578182fd5b8235611cf281612377565b946020939093013593505050565b600080600060608486031215611d14578283fd5b8335611d1f81612377565b95602085013595506040909401359392505050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0392831681529116602082015260400190565b6000602080835283518082850152825b81811015611db457858101830151858201604001528201611d98565b81811115611dc55783604083870101525b50601f01601f1916929092016040019392505050565b60006020808301818452828554600180821660008114611e025760018114611e2057611e58565b60028304607f16855260ff1983166040890152606088019350611e58565b60028304808652611e308a61236b565b885b82811015611e4e5781548b820160400152908401908801611e32565b8a01604001955050505b5091979650505050505050565b6020808252603d908201527f5b5145432d3032383030365d2d417661696c61626c65206c6f636b656420616d60408201527f6f756e74206c657373207468616e20756e6c6f636b20616d6f756e742e000000606082015260800190565b60208082526049908201527f5b5145432d3032383030375d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520746f6b656e206c6f636b20736f757263657320686176606082015268329030b1b1b2b9b99760b91b608082015260a00190565b6020808252604e908201527f5b5145432d3032383030345d2d4e6f7420656e6f7567682074696d652068617360408201527f20656c61707365642073696e63652074686520616e6e6f756e63656d656e742060608201526d37b3103a3432903ab73637b1b59760911b608082015260a00190565b60208082526044908201527f5b5145432d3032383030365d2d43616e6e6f7420656e666f72636520746f207560408201527f6e6c6f636b206d6f7265207468616e2069732063757272656e746c79206c6f6360608201526335b2b21760e11b608082015260a00190565b60208082526048908201527f5b5145432d3032383030395d2d43616e6e6f74206368616e676520766f74696e60408201527f67206167656e74206265666f726520706173736f7665722074696d65206973206060820152673932b0b1b432b21760c11b608082015260a00190565b6020808252605b908201527f5b5145432d3032383030385d2d4368616e67696e672074686520766f74696e6760408201527f206167656e742068617320746f20626520616e6e6f756e6365642c206661696c60608201527a32b2103a379039b2ba103732bb903b37ba34b7339030b3b2b73a1760291b608082015260a00190565b60208082526039908201527f5b5145432d3032383030335d2d43616e6e6f7420756e6c6f636b206d6f7265206040820152783a3430b71034b99031bab93932b73a363c903637b1b5b2b21760391b606082015260800190565b60208082526039908201527f5b5145432d3032383030325d2d496e76616c696420616d6f756e742076616c7560408201527832961030b6b7bab73a1031b0b73737ba103132903d32b9379760391b606082015260800190565b60208082526043908201527f5b5145432d3032383031305d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920766f74696e6720636f6e7472616374732068617665206163636560608201526239b99760e91b608082015260a00190565b60208082526051908201527f5b5145432d3032383030355d2d536d61727420756e6c6f636b206e6f7420706f60408201527f737369626c652c20746f6b656e7320617265207374696c6c206c6f636b656420606082015270313c903932b1b2b73a103b37ba34b7339760791b608082015260a00190565b815181526020808301516001600160a01b03169082015260408201516080820190600381106122bb57fe5b806040840152506060830151606083015292915050565b8151815260208083015190820152604080830151908201526060918201519181019190915260800190565b90815260200190565b9384526001600160a01b0392909216602084015215156040830152606082015260800190565b93845260208401929092526040830152606082015260800190565b60405181810167ffffffffffffffff8111828210171561236357fe5b604052919050565b60009081526020902090565b6001600160a01b0381168114611b2d57600080fdfe48f899e185e3470300612fd03b36ae16768c2212cc0fdc94f44cf5d60bf3227e496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65645b5145432d3032383030315d2d556e6b6e6f776e20746f6b656e206c6f636b20736f757263652ea26469706673582212206eef6fac646abf038ea8588968b4ebc27c2163505dfca6ae9f21508c0a5827e764736f6c63430007060033",
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
