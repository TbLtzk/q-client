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

// ValidationRewardPoolsPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidationRewardPoolsPoolInfo struct {
	PoolBalance               *big.Int
	ReservedForClaims         *big.Int
	AggregatedNormalizedStake *big.Int
	DelegatedStake            *big.Int
	CompoundRate              *big.Int
	LastUpdateOfCompoundRate  *big.Int
	DelegatorsShare           *big.Int
}

// ValidationRewardPoolsValidatorProperties is an auto generated low-level Go binding around an user-defined struct.
type ValidationRewardPoolsValidatorProperties struct {
	Balance                   *big.Int
	ReservedForClaim          *big.Int
	DelegatorsShare           *big.Int
	AggregatedNormalizedStake *big.Int
	CompoundRate              common.Address
}

// ValidationRewardPoolsMetaData contains all meta data concerning the ValidationRewardPools contract.
var ValidationRewardPoolsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newDelegatorsShare\",\"type\":\"uint256\"}],\"name\":\"DelegatorsShareChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_claimerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferedToQVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structValidationRewardPools.ValidatorProperties\",\"name\":\"_v\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_inc\",\"type\":\"uint256\"}],\"name\":\"UpdateRate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorsProperties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_crUpdateMinimumBase\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addCompoundRateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"requestRewardTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"updateValidatorsCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelegatorsShare\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"reserveAdditionalFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getDelegatorsShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"addAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"subAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateOfCompoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"}],\"internalType\":\"structValidationRewardPools.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getDelegatedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getLastUpdateOfCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetAmount\",\"type\":\"uint256\"}],\"name\":\"getNormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_normalizedAmount\",\"type\":\"uint256\"}],\"name\":\"getDenormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506117a8806100206000396000f3fe6080604052600436106100e95760003560e01c8063882b5dc811610085578063882b5dc8146102ec5780638cfe7df31461031c578063aaaa646d1461033c578063cd6dc6871461035c578063e268281c1461037c578063eab136a01461039c578063f38f2ce2146103af578063f8b2cb4f146103cf578063fd5d904d1461040557600080fd5b806306bfa938146100ee5780630ce98a301461016d57806313f7d4a51461019b57806333da982e146101b057806343da20f0146101d0578063443817f5146101f05780634a9b319a146102735780637c47c16d14610293578063841c9064146102b3575b600080fd5b3480156100fa57600080fd5b5061010e610109366004611540565b610425565b6040516101649190600060e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b60405180910390f35b34801561017957600080fd5b5061018d610188366004611540565b61063d565b604051908152602001610164565b6101ae6101a9366004611540565b6106c6565b005b3480156101bc57600080fd5b5061018d6101cb366004611564565b61074c565b3480156101dc57600080fd5b506101ae6101eb366004611590565b61079c565b3480156101fc57600080fd5b5061024261020b366004611540565b600160208190526000918252604090912080549181015460028201546003830154600490930154919290916001600160a01b031685565b6040805195865260208601949094529284019190915260608301526001600160a01b0316608082015260a001610164565b34801561027f57600080fd5b506101ae61028e366004611564565b610987565b34801561029f57600080fd5b5061018d6102ae366004611540565b6109f3565b3480156102bf57600080fd5b5061018d6102ce366004611540565b6001600160a01b031660009081526001602052604090206002015490565b3480156102f857600080fd5b5061030c610307366004611564565b610a20565b6040519015158152602001610164565b34801561032857600080fd5b506101ae610337366004611540565b610b4f565b34801561034857600080fd5b5061018d610357366004611540565b610c21565b34801561036857600080fd5b506101ae610377366004611564565b610c6a565b34801561038857600080fd5b5061018d610397366004611564565b610d46565b6101ae6103aa366004611540565b610d82565b3480156103bb57600080fd5b506101ae6103ca366004611564565b610dc7565b3480156103db57600080fd5b5061018d6103ea366004611540565b6001600160a01b031660009081526001602052604090205490565b34801561041157600080fd5b5061018d610420366004611540565b610e2a565b61042d6114ee565b6001600160a01b038083166000908152600160208181526040808420815160a0810183528154815293810154928401929092526002820154908301526003810154606083015260040154909216608083015261048884610ea4565b9050806104936114ee565b816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104cc57600080fd5b505afa1580156104e0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050491906115a9565b8160a0018181525050816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561054657600080fd5b505afa15801561055a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057e91906115a9565b608082015261058b610f45565b6001600160a01b031663b9a55be9876040518263ffffffff1660e01b81526004016105b691906115c2565b60206040518083038186803b1580156105ce57600080fd5b505afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060691906115a9565b60608083019190915284015160408083019190915284015160c0820152835181526020938401519381019390935250909392505050565b60008061064983610ea4565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561068657600080fd5b505afa15801561069a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106be91906115a9565b949350505050565b346106ec5760405162461bcd60e51b81526004016106e3906115d6565b60405180910390fd5b6001600160a01b03811660009081526001602052604081208054349290610714908490611644565b90915550506001600160a01b03811660009081526001602081905260408220018054349290610744908490611644565b909155505050565b60008061075884610ea4565b604051631d335d6560e21b8152600481018590529091506001600160a01b038216906374cd7594906024015b60206040518083038186803b15801561068657600080fd5b6107a4610f45565b6001600160a01b031663f2f4a2f1336040518263ffffffff1660e01b81526004016107cf91906115c2565b60206040518083038186803b1580156107e757600080fd5b505afa1580156107fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081f919061165c565b6108a55760405162461bcd60e51b815260206004820152604b60248201527f5b5145432d3031363030355d2d5468652063616c6c6572206973206e6f74206160448201527f2076616c696461746f722c206661696c656420746f207365742064656c65676160648201526a3a37b9399039b430b9329760a91b608482015260a4016106e3565b6b033b2e3c9fd0803ce80000008111156109195760405162461bcd60e51b815260206004820152602f60248201527f5b5145432d3031363030335d2d496e76616c69642064656c656761746f72207360448201526e3430b932903830b930b6b2ba32b91760891b60648201526084016106e3565b33600090815260016020526040902060020154818114610983573360008181526001602052604090819020600201849055517f474e82544a29fcdfa29d6a881106651c6cefaf1848e6973a122c1776f66de4169061097a9085815260200190565b60405180910390a25b5050565b61098f610ff8565b6001600160a01b0316336001600160a01b0316146109bf5760405162461bcd60e51b81526004016106e39061167e565b6001600160a01b038216600090815260016020526040812060030180548392906109ea908490611644565b90915550505050565b6001600160a01b038116600090815260016020526040812060030154610a1a90839061074c565b92915050565b6000610a2a610ff8565b6001600160a01b0316336001600160a01b031614610a5a5760405162461bcd60e51b81526004016106e39061167e565b6001600160a01b0383166000908152600160205260408120805490918491839190610a869084906116e9565b9250508190555082816001016000828254610aa191906116e9565b90915550610aaf9050610ff8565b6001600160a01b03166331989b58846040518263ffffffff1660e01b81526004016000604051808303818588803b158015610ae957600080fd5b505af1158015610afd573d6000803e3d6000fd5b5050505050836001600160a01b03167fa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c84604051610b3d91815260200190565b60405180910390a25060019392505050565b6001600160a01b0381811660009081526001602052604090206004015416610c1e57610b79611056565b6001600160a01b031663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610bb357600080fd5b505af1158015610bc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610beb9190611700565b6001600160a01b03828116600090815260016020526040902060040180546001600160a01b031916929091169190911790555b50565b600080610c2d83610ea4565b90506000816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561068657600080fd5b600054610100900460ff1680610c83575060005460ff16155b610ce65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106e3565b600054610100900460ff16158015610d08576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905560028290558015610d41576000805461ff00191690555b505050565b600080610d5284610ea4565b604051638dc3311d60e01b8152600481018590529091506001600160a01b03821690638dc3311d90602401610784565b34610d9f5760405162461bcd60e51b81526004016106e3906115d6565b6001600160a01b03811660009081526001602052604081208054349290610744908490611644565b610dcf610ff8565b6001600160a01b0316336001600160a01b031614610dff5760405162461bcd60e51b81526004016106e39061167e565b6001600160a01b038216600090815260016020526040812060030180548392906109ea9084906116e9565b600080610e36836110b6565b9050610e40610f45565b6001600160a01b031663f27f73c9846040518263ffffffff1660e01b8152600401610e6b91906115c2565b600060405180830381600087803b158015610e8557600080fd5b505af1158015610e99573d6000803e3d6000fd5b509295945050505050565b6001600160a01b0380821660009081526001602052604081206004015490911680610a1a5760405162461bcd60e51b8152602060048201526044602482018190527f5b5145432d3031363030375d2d436f6d706f756e64526174654b656570657220908201527f6e6f7420696e697469616c697a656420666f7220676976656e2076616c6964616064820152633a37b91760e11b608482015260a4016106e3565b600080546040805180820182526015815274676f7665726e616e63652e76616c696461746f727360581b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb9027191610fa39160040161171d565b60206040518083038186803b158015610fbb57600080fd5b505afa158015610fcf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff39190611700565b905090565b6000805460408051808201825260158152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb9027191610fa39160040161171d565b60008054604080518082018252601781527631b7b6b6b7b7173330b1ba37b93c9731b925b2b2b832b960491b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb9027191610fa39160040161171d565b6001600160a01b038082166000908152600160208181526040808420815160a08101835281548152938101549284019290925260028201549083015260038101546060830152600401549092166080830152908161111384610ea4565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561115057600080fd5b505afa158015611164573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118891906115a9565b90506000826001600160a01b03166374cd759485606001516040518263ffffffff1660e01b81526004016111be91815260200190565b60206040518083038186803b1580156111d657600080fd5b505afa1580156111ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120e91906115a9565b905060008460200151856000015161122691906116e9565b9050600181111580611239575060025482105b1561124957509095945050505050565b604051631ce149fb60e11b815260048101839052602481018290526000906001600160a01b038616906339c293f690604401602060405180830381600087803b15801561129557600080fd5b505af11580156112a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112cd91906115a9565b90506000856001600160a01b03166374cd759488606001516040518263ffffffff1660e01b815260040161130391815260200190565b60206040518083038186803b15801561131b57600080fd5b505afa15801561132f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135391906115a9565b9050600061136185836116e9565b9050600088606001511180156113775750600081115b1561144157600088602001518261138e9190611644565b89519091508111156114235760405162461bcd60e51b815260206004820152605260248201527f5b5145432d3031363030325d2d5468657265206973206e6f7420656e6f75676860448201527f20706f6f6c2062616c616e63652c206661696c656420746f20757064617465206064820152713a34329031b7b6b837bab732103930ba329760711b608482015260a4016106e3565b6001600160a01b038b16600090815260016020819052604090912001555b6001600160a01b038a166000908152600160205260409081902090517f02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782916114d891899087908690845481526001850154602082015260028501546040820152600385015460608201526004909401546001600160a01b0316608085015260a084019290925260c083015260e08201526101000190565b60405180910390a1509098975050505050505050565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6001600160a01b0381168114610c1e57600080fd5b60006020828403121561155257600080fd5b813561155d8161152b565b9392505050565b6000806040838503121561157757600080fd5b82356115828161152b565b946020939093013593505050565b6000602082840312156115a257600080fd5b5035919050565b6000602082840312156115bb57600080fd5b5051919050565b6001600160a01b0391909116815260200190565b60208082526038908201527f5b5145432d3031363030305d2d496e76616c69642076616c756520746f20696e60408201527731b932b0b9b2903a3432903837b7b6103130b630b731b29760411b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600082198211156116575761165761162e565b500190565b60006020828403121561166e57600080fd5b8151801515811461155d57600080fd5b60208082526045908201527f5b5145432d3031363030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520515661756c7420636f6e74726163742068617320616360608201526431b2b9b99760d91b608082015260a00190565b6000828210156116fb576116fb61162e565b500390565b60006020828403121561171257600080fd5b815161155d8161152b565b600060208083528351808285015260005b8181101561174a5785810183015185820160400152820161172e565b8181111561175c576000604083870101525b50601f01601f191692909201604001939250505056fea2646970667358221220f157481401fc9133c0cde04b5a979e71f531bb7e1dbf42824f7aa21a53b3816b64736f6c63430008090033",
}

// ValidationRewardPoolsABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidationRewardPoolsMetaData.ABI instead.
var ValidationRewardPoolsABI = ValidationRewardPoolsMetaData.ABI

// ValidationRewardPoolsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidationRewardPoolsMetaData.Bin instead.
var ValidationRewardPoolsBin = ValidationRewardPoolsMetaData.Bin

// DeployValidationRewardPools deploys a new Ethereum contract, binding an instance of ValidationRewardPools to it.
func DeployValidationRewardPools(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidationRewardPools, error) {
	parsed, err := ValidationRewardPoolsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidationRewardPoolsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidationRewardPools{ValidationRewardPoolsCaller: ValidationRewardPoolsCaller{contract: contract}, ValidationRewardPoolsTransactor: ValidationRewardPoolsTransactor{contract: contract}, ValidationRewardPoolsFilterer: ValidationRewardPoolsFilterer{contract: contract}}, nil
}

// ValidationRewardPools is an auto generated Go binding around an Ethereum contract.
type ValidationRewardPools struct {
	ValidationRewardPoolsCaller     // Read-only binding to the contract
	ValidationRewardPoolsTransactor // Write-only binding to the contract
	ValidationRewardPoolsFilterer   // Log filterer for contract events
}

// ValidationRewardPoolsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidationRewardPoolsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardPoolsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidationRewardPoolsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardPoolsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidationRewardPoolsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidationRewardPoolsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidationRewardPoolsSession struct {
	Contract     *ValidationRewardPools // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ValidationRewardPoolsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidationRewardPoolsCallerSession struct {
	Contract *ValidationRewardPoolsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ValidationRewardPoolsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidationRewardPoolsTransactorSession struct {
	Contract     *ValidationRewardPoolsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ValidationRewardPoolsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidationRewardPoolsRaw struct {
	Contract *ValidationRewardPools // Generic contract binding to access the raw methods on
}

// ValidationRewardPoolsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidationRewardPoolsCallerRaw struct {
	Contract *ValidationRewardPoolsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidationRewardPoolsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidationRewardPoolsTransactorRaw struct {
	Contract *ValidationRewardPoolsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidationRewardPools creates a new instance of ValidationRewardPools, bound to a specific deployed contract.
func NewValidationRewardPools(address common.Address, backend bind.ContractBackend) (*ValidationRewardPools, error) {
	contract, err := bindValidationRewardPools(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPools{ValidationRewardPoolsCaller: ValidationRewardPoolsCaller{contract: contract}, ValidationRewardPoolsTransactor: ValidationRewardPoolsTransactor{contract: contract}, ValidationRewardPoolsFilterer: ValidationRewardPoolsFilterer{contract: contract}}, nil
}

// NewValidationRewardPoolsCaller creates a new read-only instance of ValidationRewardPools, bound to a specific deployed contract.
func NewValidationRewardPoolsCaller(address common.Address, caller bind.ContractCaller) (*ValidationRewardPoolsCaller, error) {
	contract, err := bindValidationRewardPools(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsCaller{contract: contract}, nil
}

// NewValidationRewardPoolsTransactor creates a new write-only instance of ValidationRewardPools, bound to a specific deployed contract.
func NewValidationRewardPoolsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidationRewardPoolsTransactor, error) {
	contract, err := bindValidationRewardPools(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsTransactor{contract: contract}, nil
}

// NewValidationRewardPoolsFilterer creates a new log filterer instance of ValidationRewardPools, bound to a specific deployed contract.
func NewValidationRewardPoolsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidationRewardPoolsFilterer, error) {
	contract, err := bindValidationRewardPools(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsFilterer{contract: contract}, nil
}

// bindValidationRewardPools binds a generic wrapper to an already deployed contract.
func bindValidationRewardPools(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRewardPoolsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardPools *ValidationRewardPoolsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardPools.Contract.ValidationRewardPoolsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardPools *ValidationRewardPoolsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.ValidationRewardPoolsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardPools *ValidationRewardPoolsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.ValidationRewardPoolsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidationRewardPools *ValidationRewardPoolsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidationRewardPools.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidationRewardPools *ValidationRewardPoolsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidationRewardPools *ValidationRewardPoolsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetBalance(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getBalance", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetBalance(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetBalance(&_ValidationRewardPools.CallOpts, _validator)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetBalance(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetBalance(&_ValidationRewardPools.CallOpts, _validator)
}

// GetCompoundRate is a free data retrieval call binding the contract method 0x0ce98a30.
//
// Solidity: function getCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetCompoundRate(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getCompoundRate", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundRate is a free data retrieval call binding the contract method 0x0ce98a30.
//
// Solidity: function getCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetCompoundRate(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetCompoundRate(&_ValidationRewardPools.CallOpts, _validator)
}

// GetCompoundRate is a free data retrieval call binding the contract method 0x0ce98a30.
//
// Solidity: function getCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetCompoundRate(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetCompoundRate(&_ValidationRewardPools.CallOpts, _validator)
}

// GetDelegatedStake is a free data retrieval call binding the contract method 0x7c47c16d.
//
// Solidity: function getDelegatedStake(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetDelegatedStake(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getDelegatedStake", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatedStake is a free data retrieval call binding the contract method 0x7c47c16d.
//
// Solidity: function getDelegatedStake(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetDelegatedStake(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDelegatedStake(&_ValidationRewardPools.CallOpts, _validator)
}

// GetDelegatedStake is a free data retrieval call binding the contract method 0x7c47c16d.
//
// Solidity: function getDelegatedStake(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetDelegatedStake(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDelegatedStake(&_ValidationRewardPools.CallOpts, _validator)
}

// GetDelegatorsShare is a free data retrieval call binding the contract method 0x841c9064.
//
// Solidity: function getDelegatorsShare(address _addr) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetDelegatorsShare(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getDelegatorsShare", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorsShare is a free data retrieval call binding the contract method 0x841c9064.
//
// Solidity: function getDelegatorsShare(address _addr) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetDelegatorsShare(_addr common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDelegatorsShare(&_ValidationRewardPools.CallOpts, _addr)
}

// GetDelegatorsShare is a free data retrieval call binding the contract method 0x841c9064.
//
// Solidity: function getDelegatorsShare(address _addr) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetDelegatorsShare(_addr common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDelegatorsShare(&_ValidationRewardPools.CallOpts, _addr)
}

// GetDenormalizedAmount is a free data retrieval call binding the contract method 0x33da982e.
//
// Solidity: function getDenormalizedAmount(address _validator, uint256 _normalizedAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetDenormalizedAmount(opts *bind.CallOpts, _validator common.Address, _normalizedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getDenormalizedAmount", _validator, _normalizedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedAmount is a free data retrieval call binding the contract method 0x33da982e.
//
// Solidity: function getDenormalizedAmount(address _validator, uint256 _normalizedAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetDenormalizedAmount(_validator common.Address, _normalizedAmount *big.Int) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDenormalizedAmount(&_ValidationRewardPools.CallOpts, _validator, _normalizedAmount)
}

// GetDenormalizedAmount is a free data retrieval call binding the contract method 0x33da982e.
//
// Solidity: function getDenormalizedAmount(address _validator, uint256 _normalizedAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetDenormalizedAmount(_validator common.Address, _normalizedAmount *big.Int) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetDenormalizedAmount(&_ValidationRewardPools.CallOpts, _validator, _normalizedAmount)
}

// GetLastUpdateOfCompoundRate is a free data retrieval call binding the contract method 0xaaaa646d.
//
// Solidity: function getLastUpdateOfCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetLastUpdateOfCompoundRate(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getLastUpdateOfCompoundRate", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdateOfCompoundRate is a free data retrieval call binding the contract method 0xaaaa646d.
//
// Solidity: function getLastUpdateOfCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetLastUpdateOfCompoundRate(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetLastUpdateOfCompoundRate(&_ValidationRewardPools.CallOpts, _validator)
}

// GetLastUpdateOfCompoundRate is a free data retrieval call binding the contract method 0xaaaa646d.
//
// Solidity: function getLastUpdateOfCompoundRate(address _validator) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetLastUpdateOfCompoundRate(_validator common.Address) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetLastUpdateOfCompoundRate(&_ValidationRewardPools.CallOpts, _validator)
}

// GetNormalizedAmount is a free data retrieval call binding the contract method 0xe268281c.
//
// Solidity: function getNormalizedAmount(address _validator, uint256 _targetAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetNormalizedAmount(opts *bind.CallOpts, _validator common.Address, _targetAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getNormalizedAmount", _validator, _targetAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedAmount is a free data retrieval call binding the contract method 0xe268281c.
//
// Solidity: function getNormalizedAmount(address _validator, uint256 _targetAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetNormalizedAmount(_validator common.Address, _targetAmount *big.Int) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetNormalizedAmount(&_ValidationRewardPools.CallOpts, _validator, _targetAmount)
}

// GetNormalizedAmount is a free data retrieval call binding the contract method 0xe268281c.
//
// Solidity: function getNormalizedAmount(address _validator, uint256 _targetAmount) view returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetNormalizedAmount(_validator common.Address, _targetAmount *big.Int) (*big.Int, error) {
	return _ValidationRewardPools.Contract.GetNormalizedAmount(&_ValidationRewardPools.CallOpts, _validator, _targetAmount)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _validator) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ValidationRewardPools *ValidationRewardPoolsCaller) GetPoolInfo(opts *bind.CallOpts, _validator common.Address) (ValidationRewardPoolsPoolInfo, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "getPoolInfo", _validator)

	if err != nil {
		return *new(ValidationRewardPoolsPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidationRewardPoolsPoolInfo)).(*ValidationRewardPoolsPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _validator) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ValidationRewardPools *ValidationRewardPoolsSession) GetPoolInfo(_validator common.Address) (ValidationRewardPoolsPoolInfo, error) {
	return _ValidationRewardPools.Contract.GetPoolInfo(&_ValidationRewardPools.CallOpts, _validator)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _validator) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) GetPoolInfo(_validator common.Address) (ValidationRewardPoolsPoolInfo, error) {
	return _ValidationRewardPools.Contract.GetPoolInfo(&_ValidationRewardPools.CallOpts, _validator)
}

// ValidatorsProperties is a free data retrieval call binding the contract method 0x443817f5.
//
// Solidity: function validatorsProperties(address ) view returns(uint256 balance, uint256 reservedForClaim, uint256 delegatorsShare, uint256 aggregatedNormalizedStake, address compoundRate)
func (_ValidationRewardPools *ValidationRewardPoolsCaller) ValidatorsProperties(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance                   *big.Int
	ReservedForClaim          *big.Int
	DelegatorsShare           *big.Int
	AggregatedNormalizedStake *big.Int
	CompoundRate              common.Address
}, error) {
	var out []interface{}
	err := _ValidationRewardPools.contract.Call(opts, &out, "validatorsProperties", arg0)

	outstruct := new(struct {
		Balance                   *big.Int
		ReservedForClaim          *big.Int
		DelegatorsShare           *big.Int
		AggregatedNormalizedStake *big.Int
		CompoundRate              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReservedForClaim = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelegatorsShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AggregatedNormalizedStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CompoundRate = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ValidatorsProperties is a free data retrieval call binding the contract method 0x443817f5.
//
// Solidity: function validatorsProperties(address ) view returns(uint256 balance, uint256 reservedForClaim, uint256 delegatorsShare, uint256 aggregatedNormalizedStake, address compoundRate)
func (_ValidationRewardPools *ValidationRewardPoolsSession) ValidatorsProperties(arg0 common.Address) (struct {
	Balance                   *big.Int
	ReservedForClaim          *big.Int
	DelegatorsShare           *big.Int
	AggregatedNormalizedStake *big.Int
	CompoundRate              common.Address
}, error) {
	return _ValidationRewardPools.Contract.ValidatorsProperties(&_ValidationRewardPools.CallOpts, arg0)
}

// ValidatorsProperties is a free data retrieval call binding the contract method 0x443817f5.
//
// Solidity: function validatorsProperties(address ) view returns(uint256 balance, uint256 reservedForClaim, uint256 delegatorsShare, uint256 aggregatedNormalizedStake, address compoundRate)
func (_ValidationRewardPools *ValidationRewardPoolsCallerSession) ValidatorsProperties(arg0 common.Address) (struct {
	Balance                   *big.Int
	ReservedForClaim          *big.Int
	DelegatorsShare           *big.Int
	AggregatedNormalizedStake *big.Int
	CompoundRate              common.Address
}, error) {
	return _ValidationRewardPools.Contract.ValidatorsProperties(&_ValidationRewardPools.CallOpts, arg0)
}

// AddAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0x4a9b319a.
//
// Solidity: function addAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) AddAggregatedNormalizedStake(opts *bind.TransactOpts, _validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "addAggregatedNormalizedStake", _validator, _stake)
}

// AddAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0x4a9b319a.
//
// Solidity: function addAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) AddAggregatedNormalizedStake(_validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.AddAggregatedNormalizedStake(&_ValidationRewardPools.TransactOpts, _validator, _stake)
}

// AddAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0x4a9b319a.
//
// Solidity: function addAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) AddAggregatedNormalizedStake(_validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.AddAggregatedNormalizedStake(&_ValidationRewardPools.TransactOpts, _validator, _stake)
}

// AddCompoundRateKeeper is a paid mutator transaction binding the contract method 0x8cfe7df3.
//
// Solidity: function addCompoundRateKeeper(address _validator) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) AddCompoundRateKeeper(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "addCompoundRateKeeper", _validator)
}

// AddCompoundRateKeeper is a paid mutator transaction binding the contract method 0x8cfe7df3.
//
// Solidity: function addCompoundRateKeeper(address _validator) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) AddCompoundRateKeeper(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.AddCompoundRateKeeper(&_ValidationRewardPools.TransactOpts, _validator)
}

// AddCompoundRateKeeper is a paid mutator transaction binding the contract method 0x8cfe7df3.
//
// Solidity: function addCompoundRateKeeper(address _validator) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) AddCompoundRateKeeper(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.AddCompoundRateKeeper(&_ValidationRewardPools.TransactOpts, _validator)
}

// Increase is a paid mutator transaction binding the contract method 0xeab136a0.
//
// Solidity: function increase(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) Increase(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "increase", _validator)
}

// Increase is a paid mutator transaction binding the contract method 0xeab136a0.
//
// Solidity: function increase(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) Increase(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Increase(&_ValidationRewardPools.TransactOpts, _validator)
}

// Increase is a paid mutator transaction binding the contract method 0xeab136a0.
//
// Solidity: function increase(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) Increase(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Increase(&_ValidationRewardPools.TransactOpts, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _registry, uint256 _crUpdateMinimumBase) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _crUpdateMinimumBase *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "initialize", _registry, _crUpdateMinimumBase)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _registry, uint256 _crUpdateMinimumBase) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) Initialize(_registry common.Address, _crUpdateMinimumBase *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Initialize(&_ValidationRewardPools.TransactOpts, _registry, _crUpdateMinimumBase)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _registry, uint256 _crUpdateMinimumBase) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) Initialize(_registry common.Address, _crUpdateMinimumBase *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Initialize(&_ValidationRewardPools.TransactOpts, _registry, _crUpdateMinimumBase)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0x882b5dc8.
//
// Solidity: function requestRewardTransfer(address _validator, uint256 _rewardAmount) returns(bool)
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) RequestRewardTransfer(opts *bind.TransactOpts, _validator common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "requestRewardTransfer", _validator, _rewardAmount)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0x882b5dc8.
//
// Solidity: function requestRewardTransfer(address _validator, uint256 _rewardAmount) returns(bool)
func (_ValidationRewardPools *ValidationRewardPoolsSession) RequestRewardTransfer(_validator common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.RequestRewardTransfer(&_ValidationRewardPools.TransactOpts, _validator, _rewardAmount)
}

// RequestRewardTransfer is a paid mutator transaction binding the contract method 0x882b5dc8.
//
// Solidity: function requestRewardTransfer(address _validator, uint256 _rewardAmount) returns(bool)
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) RequestRewardTransfer(_validator common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.RequestRewardTransfer(&_ValidationRewardPools.TransactOpts, _validator, _rewardAmount)
}

// ReserveAdditionalFunds is a paid mutator transaction binding the contract method 0x13f7d4a5.
//
// Solidity: function reserveAdditionalFunds(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) ReserveAdditionalFunds(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "reserveAdditionalFunds", _validator)
}

// ReserveAdditionalFunds is a paid mutator transaction binding the contract method 0x13f7d4a5.
//
// Solidity: function reserveAdditionalFunds(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) ReserveAdditionalFunds(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.ReserveAdditionalFunds(&_ValidationRewardPools.TransactOpts, _validator)
}

// ReserveAdditionalFunds is a paid mutator transaction binding the contract method 0x13f7d4a5.
//
// Solidity: function reserveAdditionalFunds(address _validator) payable returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) ReserveAdditionalFunds(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.ReserveAdditionalFunds(&_ValidationRewardPools.TransactOpts, _validator)
}

// SetDelegatorsShare is a paid mutator transaction binding the contract method 0x43da20f0.
//
// Solidity: function setDelegatorsShare(uint256 _newDelegatorsShare) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) SetDelegatorsShare(opts *bind.TransactOpts, _newDelegatorsShare *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "setDelegatorsShare", _newDelegatorsShare)
}

// SetDelegatorsShare is a paid mutator transaction binding the contract method 0x43da20f0.
//
// Solidity: function setDelegatorsShare(uint256 _newDelegatorsShare) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) SetDelegatorsShare(_newDelegatorsShare *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.SetDelegatorsShare(&_ValidationRewardPools.TransactOpts, _newDelegatorsShare)
}

// SetDelegatorsShare is a paid mutator transaction binding the contract method 0x43da20f0.
//
// Solidity: function setDelegatorsShare(uint256 _newDelegatorsShare) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) SetDelegatorsShare(_newDelegatorsShare *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.SetDelegatorsShare(&_ValidationRewardPools.TransactOpts, _newDelegatorsShare)
}

// SubAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0xf38f2ce2.
//
// Solidity: function subAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) SubAggregatedNormalizedStake(opts *bind.TransactOpts, _validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "subAggregatedNormalizedStake", _validator, _stake)
}

// SubAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0xf38f2ce2.
//
// Solidity: function subAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) SubAggregatedNormalizedStake(_validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.SubAggregatedNormalizedStake(&_ValidationRewardPools.TransactOpts, _validator, _stake)
}

// SubAggregatedNormalizedStake is a paid mutator transaction binding the contract method 0xf38f2ce2.
//
// Solidity: function subAggregatedNormalizedStake(address _validator, uint256 _stake) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) SubAggregatedNormalizedStake(_validator common.Address, _stake *big.Int) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.SubAggregatedNormalizedStake(&_ValidationRewardPools.TransactOpts, _validator, _stake)
}

// UpdateValidatorsCompoundRate is a paid mutator transaction binding the contract method 0xfd5d904d.
//
// Solidity: function updateValidatorsCompoundRate(address _validator) returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) UpdateValidatorsCompoundRate(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "updateValidatorsCompoundRate", _validator)
}

// UpdateValidatorsCompoundRate is a paid mutator transaction binding the contract method 0xfd5d904d.
//
// Solidity: function updateValidatorsCompoundRate(address _validator) returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsSession) UpdateValidatorsCompoundRate(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.UpdateValidatorsCompoundRate(&_ValidationRewardPools.TransactOpts, _validator)
}

// UpdateValidatorsCompoundRate is a paid mutator transaction binding the contract method 0xfd5d904d.
//
// Solidity: function updateValidatorsCompoundRate(address _validator) returns(uint256)
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) UpdateValidatorsCompoundRate(_validator common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.UpdateValidatorsCompoundRate(&_ValidationRewardPools.TransactOpts, _validator)
}

// ValidationRewardPoolsDelegatorsShareChangedIterator is returned from FilterDelegatorsShareChanged and is used to iterate over the raw logs and unpacked data for DelegatorsShareChanged events raised by the ValidationRewardPools contract.
type ValidationRewardPoolsDelegatorsShareChangedIterator struct {
	Event *ValidationRewardPoolsDelegatorsShareChanged // Event containing the contract specifics and raw log

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
func (it *ValidationRewardPoolsDelegatorsShareChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRewardPoolsDelegatorsShareChanged)
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
		it.Event = new(ValidationRewardPoolsDelegatorsShareChanged)
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
func (it *ValidationRewardPoolsDelegatorsShareChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRewardPoolsDelegatorsShareChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRewardPoolsDelegatorsShareChanged represents a DelegatorsShareChanged event raised by the ValidationRewardPools contract.
type ValidationRewardPoolsDelegatorsShareChanged struct {
	ValidatorAddress   common.Address
	NewDelegatorsShare *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatorsShareChanged is a free log retrieval operation binding the contract event 0x474e82544a29fcdfa29d6a881106651c6cefaf1848e6973a122c1776f66de416.
//
// Solidity: event DelegatorsShareChanged(address indexed _validatorAddress, uint256 _newDelegatorsShare)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) FilterDelegatorsShareChanged(opts *bind.FilterOpts, _validatorAddress []common.Address) (*ValidationRewardPoolsDelegatorsShareChangedIterator, error) {

	var _validatorAddressRule []interface{}
	for _, _validatorAddressItem := range _validatorAddress {
		_validatorAddressRule = append(_validatorAddressRule, _validatorAddressItem)
	}

	logs, sub, err := _ValidationRewardPools.contract.FilterLogs(opts, "DelegatorsShareChanged", _validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsDelegatorsShareChangedIterator{contract: _ValidationRewardPools.contract, event: "DelegatorsShareChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorsShareChanged is a free log subscription operation binding the contract event 0x474e82544a29fcdfa29d6a881106651c6cefaf1848e6973a122c1776f66de416.
//
// Solidity: event DelegatorsShareChanged(address indexed _validatorAddress, uint256 _newDelegatorsShare)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) WatchDelegatorsShareChanged(opts *bind.WatchOpts, sink chan<- *ValidationRewardPoolsDelegatorsShareChanged, _validatorAddress []common.Address) (event.Subscription, error) {

	var _validatorAddressRule []interface{}
	for _, _validatorAddressItem := range _validatorAddress {
		_validatorAddressRule = append(_validatorAddressRule, _validatorAddressItem)
	}

	logs, sub, err := _ValidationRewardPools.contract.WatchLogs(opts, "DelegatorsShareChanged", _validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRewardPoolsDelegatorsShareChanged)
				if err := _ValidationRewardPools.contract.UnpackLog(event, "DelegatorsShareChanged", log); err != nil {
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

// ParseDelegatorsShareChanged is a log parse operation binding the contract event 0x474e82544a29fcdfa29d6a881106651c6cefaf1848e6973a122c1776f66de416.
//
// Solidity: event DelegatorsShareChanged(address indexed _validatorAddress, uint256 _newDelegatorsShare)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) ParseDelegatorsShareChanged(log types.Log) (*ValidationRewardPoolsDelegatorsShareChanged, error) {
	event := new(ValidationRewardPoolsDelegatorsShareChanged)
	if err := _ValidationRewardPools.contract.UnpackLog(event, "DelegatorsShareChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidationRewardPoolsRewardTransferedToQVaultIterator is returned from FilterRewardTransferedToQVault and is used to iterate over the raw logs and unpacked data for RewardTransferedToQVault events raised by the ValidationRewardPools contract.
type ValidationRewardPoolsRewardTransferedToQVaultIterator struct {
	Event *ValidationRewardPoolsRewardTransferedToQVault // Event containing the contract specifics and raw log

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
func (it *ValidationRewardPoolsRewardTransferedToQVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRewardPoolsRewardTransferedToQVault)
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
		it.Event = new(ValidationRewardPoolsRewardTransferedToQVault)
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
func (it *ValidationRewardPoolsRewardTransferedToQVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRewardPoolsRewardTransferedToQVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRewardPoolsRewardTransferedToQVault represents a RewardTransferedToQVault event raised by the ValidationRewardPools contract.
type ValidationRewardPoolsRewardTransferedToQVault struct {
	ClaimerAddress common.Address
	RewardAmount   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardTransferedToQVault is a free log retrieval operation binding the contract event 0xa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c.
//
// Solidity: event RewardTransferedToQVault(address indexed _claimerAddress, uint256 _rewardAmount)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) FilterRewardTransferedToQVault(opts *bind.FilterOpts, _claimerAddress []common.Address) (*ValidationRewardPoolsRewardTransferedToQVaultIterator, error) {

	var _claimerAddressRule []interface{}
	for _, _claimerAddressItem := range _claimerAddress {
		_claimerAddressRule = append(_claimerAddressRule, _claimerAddressItem)
	}

	logs, sub, err := _ValidationRewardPools.contract.FilterLogs(opts, "RewardTransferedToQVault", _claimerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsRewardTransferedToQVaultIterator{contract: _ValidationRewardPools.contract, event: "RewardTransferedToQVault", logs: logs, sub: sub}, nil
}

// WatchRewardTransferedToQVault is a free log subscription operation binding the contract event 0xa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c.
//
// Solidity: event RewardTransferedToQVault(address indexed _claimerAddress, uint256 _rewardAmount)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) WatchRewardTransferedToQVault(opts *bind.WatchOpts, sink chan<- *ValidationRewardPoolsRewardTransferedToQVault, _claimerAddress []common.Address) (event.Subscription, error) {

	var _claimerAddressRule []interface{}
	for _, _claimerAddressItem := range _claimerAddress {
		_claimerAddressRule = append(_claimerAddressRule, _claimerAddressItem)
	}

	logs, sub, err := _ValidationRewardPools.contract.WatchLogs(opts, "RewardTransferedToQVault", _claimerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRewardPoolsRewardTransferedToQVault)
				if err := _ValidationRewardPools.contract.UnpackLog(event, "RewardTransferedToQVault", log); err != nil {
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

// ParseRewardTransferedToQVault is a log parse operation binding the contract event 0xa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c.
//
// Solidity: event RewardTransferedToQVault(address indexed _claimerAddress, uint256 _rewardAmount)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) ParseRewardTransferedToQVault(log types.Log) (*ValidationRewardPoolsRewardTransferedToQVault, error) {
	event := new(ValidationRewardPoolsRewardTransferedToQVault)
	if err := _ValidationRewardPools.contract.UnpackLog(event, "RewardTransferedToQVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidationRewardPoolsUpdateRateIterator is returned from FilterUpdateRate and is used to iterate over the raw logs and unpacked data for UpdateRate events raised by the ValidationRewardPools contract.
type ValidationRewardPoolsUpdateRateIterator struct {
	Event *ValidationRewardPoolsUpdateRate // Event containing the contract specifics and raw log

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
func (it *ValidationRewardPoolsUpdateRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidationRewardPoolsUpdateRate)
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
		it.Event = new(ValidationRewardPoolsUpdateRate)
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
func (it *ValidationRewardPoolsUpdateRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidationRewardPoolsUpdateRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidationRewardPoolsUpdateRate represents a UpdateRate event raised by the ValidationRewardPools contract.
type ValidationRewardPoolsUpdateRate struct {
	V       ValidationRewardPoolsValidatorProperties
	OldRate *big.Int
	NewRate *big.Int
	Inc     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateRate is a free log retrieval operation binding the contract event 0x02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782.
//
// Solidity: event UpdateRate((uint256,uint256,uint256,uint256,address) _v, uint256 _oldRate, uint256 _newRate, uint256 _inc)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) FilterUpdateRate(opts *bind.FilterOpts) (*ValidationRewardPoolsUpdateRateIterator, error) {

	logs, sub, err := _ValidationRewardPools.contract.FilterLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return &ValidationRewardPoolsUpdateRateIterator{contract: _ValidationRewardPools.contract, event: "UpdateRate", logs: logs, sub: sub}, nil
}

// WatchUpdateRate is a free log subscription operation binding the contract event 0x02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782.
//
// Solidity: event UpdateRate((uint256,uint256,uint256,uint256,address) _v, uint256 _oldRate, uint256 _newRate, uint256 _inc)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) WatchUpdateRate(opts *bind.WatchOpts, sink chan<- *ValidationRewardPoolsUpdateRate) (event.Subscription, error) {

	logs, sub, err := _ValidationRewardPools.contract.WatchLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidationRewardPoolsUpdateRate)
				if err := _ValidationRewardPools.contract.UnpackLog(event, "UpdateRate", log); err != nil {
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

// ParseUpdateRate is a log parse operation binding the contract event 0x02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782.
//
// Solidity: event UpdateRate((uint256,uint256,uint256,uint256,address) _v, uint256 _oldRate, uint256 _newRate, uint256 _inc)
func (_ValidationRewardPools *ValidationRewardPoolsFilterer) ParseUpdateRate(log types.Log) (*ValidationRewardPoolsUpdateRate, error) {
	event := new(ValidationRewardPoolsUpdateRate)
	if err := _ValidationRewardPools.contract.UnpackLog(event, "UpdateRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
