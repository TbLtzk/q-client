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
	Bin: "0x608060405234801561001057600080fd5b50611bfb806100206000396000f3fe6080604052600436106100e95760003560e01c8063882b5dc811610085578063882b5dc8146102375780638cfe7df314610264578063aaaa646d14610284578063cd6dc687146102a4578063e268281c146102c4578063eab136a0146102e4578063f38f2ce2146102f7578063f8b2cb4f14610317578063fd5d904d14610337576100e9565b806306bfa938146100ee5780630ce98a301461012457806313f7d4a51461015157806333da982e1461016657806343da20f014610186578063443817f5146101a65780634a9b319a146101d75780637c47c16d146101f7578063841c906414610217575b600080fd5b3480156100fa57600080fd5b5061010e6101093660046116dc565b610357565b60405161011b9190611aa2565b60405180910390f35b34801561013057600080fd5b5061014461013f3660046116dc565b6105ec565b60405161011b9190611b3f565b61016461015f3660046116dc565b610675565b005b34801561017257600080fd5b50610144610181366004611714565b610708565b34801561019257600080fd5b506101646101a136600461175f565b61075b565b3480156101b257600080fd5b506101c66101c13660046116dc565b610908565b60405161011b959493929190611b56565b3480156101e357600080fd5b506101646101f2366004611714565b61093f565b34801561020357600080fd5b506101446102123660046116dc565b610a3b565b34801561022357600080fd5b506101446102323660046116dc565b610a68565b34801561024357600080fd5b50610257610252366004611714565b610a86565b60405161011b91906117a3565b34801561027057600080fd5b5061016461027f3660046116dc565b610c95565b34801561029057600080fd5b5061014461029f3660046116dc565b610d67565b3480156102b057600080fd5b506101646102bf366004611714565b610db0565b3480156102d057600080fd5b506101446102df366004611714565b610e7b565b6101646102f23660046116dc565b610eb6565b34801561030357600080fd5b50610164610312366004611714565b610f12565b34801561032357600080fd5b506101446103323660046116dc565b610feb565b34801561034357600080fd5b506101446103523660046116dc565b611006565b61035f61169f565b6001600160a01b038083166000908152600160208181526040808420815160a081018352815481529381015492840192909252600282015490830152600381015460608301526004015490921660808301526103ba846110fb565b9050806103c561169f565b816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103fe57600080fd5b505afa158015610412573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104369190611777565b8160a0018181525050816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561047857600080fd5b505afa15801561048c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b09190611777565b6080820152600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906104e8906004016119d7565b60206040518083038186803b15801561050057600080fd5b505afa158015610514573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053891906116f8565b6001600160a01b031663b9a55be9876040518263ffffffff1660e01b8152600401610563919061178f565b60206040518083038186803b15801561057b57600080fd5b505afa15801561058f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b39190611777565b60608083019190915284015160408083019190915284015160c0820152835181526020938401519381019390935250909150505b919050565b6000806105f8836110fb565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561063557600080fd5b505afa158015610649573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066d9190611777565b949350505050565b3461069b5760405162461bcd60e51b81526004016106929061197f565b60405180910390fd5b6001600160a01b0381166000908152600160205260409020546106be9034611138565b6001600160a01b038216600090815260016020819052604090912091825501546106e89034611138565b6001600160a01b0390911660009081526001602081905260409091200155565b600080610714846110fb565b604051631d335d6560e21b81529091506001600160a01b038216906374cd759490610743908690600401611b3f565b60206040518083038186803b15801561063557600080fd5b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061078e906004016119d7565b60206040518083038186803b1580156107a657600080fd5b505afa1580156107ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107de91906116f8565b6001600160a01b031663f2f4a2f1336040518263ffffffff1660e01b8152600401610809919061178f565b60206040518083038186803b15801561082157600080fd5b505afa158015610835573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610859919061173f565b6108755760405162461bcd60e51b8152600401610692906117ae565b61087d611197565b81111561089c5760405162461bcd60e51b81526004016106929061184e565b33600090815260016020526040902060020154818114610904573360008181526001602052604090819020600201849055517f474e82544a29fcdfa29d6a881106651c6cefaf1848e6973a122c1776f66de416906108fb908590611b3f565b60405180910390a25b5050565b600160208190526000918252604090912080549181015460028201546003830154600490930154919290916001600160a01b031685565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906109729060040161181f565b60206040518083038186803b15801561098a57600080fd5b505afa15801561099e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c291906116f8565b6001600160a01b0316336001600160a01b0316146109f25760405162461bcd60e51b815260040161069290611a06565b6001600160a01b038216600090815260016020526040902060030154610a189082611138565b6001600160a01b0390921660009081526001602052604090206003019190915550565b6001600160a01b038116600090815260016020526040812060030154610a62908390610708565b92915050565b6001600160a01b031660009081526001602052604090206002015490565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610aba9060040161181f565b60206040518083038186803b158015610ad257600080fd5b505afa158015610ae6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0a91906116f8565b6001600160a01b0316336001600160a01b031614610b3a5760405162461bcd60e51b815260040161069290611a06565b6001600160a01b03831660009081526001602052604090208054610b5e90846111a7565b81556001810154610b6f90846111a7565b6001820155600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610ba79060040161181f565b60206040518083038186803b158015610bbf57600080fd5b505afa158015610bd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bf791906116f8565b6001600160a01b03166331989b58846040518263ffffffff1660e01b81526004016000604051808303818588803b158015610c3157600080fd5b505af1158015610c45573d6000803e3d6000fd5b5050505050836001600160a01b03167fa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c84604051610c839190611b3f565b60405180910390a25060019392505050565b6001600160a01b0381811660009081526001602052604090206004015416610d6457610cbf6111e9565b6001600160a01b031663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610cf957600080fd5b505af1158015610d0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3191906116f8565b6001600160a01b03828116600090815260016020526040902060040180546001600160a01b031916929091169190911790555b50565b600080610d73836110fb565b90506000816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561063557600080fd5b600054610100900460ff1680610dc95750610dc9611272565b80610dd7575060005460ff16155b610e125760405162461bcd60e51b815260040180806020018281038252602e815260200180611b98602e913960400191505060405180910390fd5b600054610100900460ff16158015610e3d576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905560028290558015610e76576000805461ff00191690555b505050565b600080610e87846110fb565b604051638dc3311d60e01b81529091506001600160a01b03821690638dc3311d90610743908690600401611b3f565b34610ed35760405162461bcd60e51b81526004016106929061197f565b6001600160a01b038116600090815260016020526040902054610ef69034611138565b6001600160a01b03909116600090815260016020526040902055565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610f459060040161181f565b60206040518083038186803b158015610f5d57600080fd5b505afa158015610f71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f9591906116f8565b6001600160a01b0316336001600160a01b031614610fc55760405162461bcd60e51b815260040161069290611a06565b6001600160a01b038216600090815260016020526040902060030154610a1890826111a7565b6001600160a01b031660009081526001602052604090205490565b60008061101283611283565b600054604051633fb9027160e01b81529192506201000090046001600160a01b031690633fb9027190611047906004016119d7565b60206040518083038186803b15801561105f57600080fd5b505afa158015611073573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109791906116f8565b6001600160a01b031663f27f73c9846040518263ffffffff1660e01b81526004016110c2919061178f565b600060405180830381600087803b1580156110dc57600080fd5b505af11580156110f0573d6000803e3d6000fd5b509295945050505050565b6001600160a01b0380821660009081526001602052604081206004015490911680610a625760405162461bcd60e51b81526004016106929061189d565b600082820183811015611190576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b6b033b2e3c9fd0803ce800000090565b600061119083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611602565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061121d90600401611a71565b60206040518083038186803b15801561123557600080fd5b505afa158015611249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126d91906116f8565b905090565b600061127d30611699565b15905090565b6001600160a01b038082166000908152600160208181526040808420815160a0810183528154815293810154928401929092526002820154908301526003810154606083015260040154909216608083015290816112e0846110fb565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561131d57600080fd5b505afa158015611331573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113559190611777565b90506000826001600160a01b03166374cd759485606001516040518263ffffffff1660e01b81526004016113899190611b3f565b60206040518083038186803b1580156113a157600080fd5b505afa1580156113b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d99190611777565b602085015185519192506000916113ef916111a7565b9050600181111580611402575060025482105b156114145782955050505050506105e7565b604051631ce149fb60e11b81526000906001600160a01b038616906339c293f6906114459086908690600401611b48565b602060405180830381600087803b15801561145f57600080fd5b505af1158015611473573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114979190611777565b90506000856001600160a01b03166374cd759488606001516040518263ffffffff1660e01b81526004016114cb9190611b3f565b60206040518083038186803b1580156114e357600080fd5b505afa1580156114f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151b9190611777565b9050600061152982866111a7565b90506000886060015111801561153f5750600081115b1561159f57600061155d89602001518361113890919063ffffffff16565b89519091508111156115815760405162461bcd60e51b815260040161069290611907565b6001600160a01b038b16600090815260016020819052604090912001555b6001600160a01b038a166000908152600160205260409081902090517f02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782916115ec91899087908690611af0565b60405180910390a1509098975050505050505050565b600081848411156116915760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561165657818101518382015260200161163e565b50505050905090810190601f1680156116835780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b3b151590565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6000602082840312156116ed578081fd5b813561119081611b82565b600060208284031215611709578081fd5b815161119081611b82565b60008060408385031215611726578081fd5b823561173181611b82565b946020939093013593505050565b600060208284031215611750578081fd5b81518015158114611190578182fd5b600060208284031215611770578081fd5b5035919050565b600060208284031215611788578081fd5b5051919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252604b908201527f5b5145432d3031363030355d2d5468652063616c6c6572206973206e6f74206160408201527f2076616c696461746f722c206661696c656420746f207365742064656c65676160608201526a3a37b9399039b430b9329760a91b608082015260a00190565b6020808252601590820152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b604082015260600190565b6020808252602f908201527f5b5145432d3031363030335d2d496e76616c69642064656c656761746f72207360408201526e3430b932903830b930b6b2ba32b91760891b606082015260800190565b60208082526044908201527f5b5145432d3031363030375d2d436f6d706f756e64526174654b65657065722060408201527f6e6f7420696e697469616c697a656420666f7220676976656e2076616c6964616060820152633a37b91760e11b608082015260a00190565b60208082526052908201527f5b5145432d3031363030325d2d5468657265206973206e6f7420656e6f75676860408201527f20706f6f6c2062616c616e63652c206661696c656420746f20757064617465206060820152713a34329031b7b6b837bab732103930ba329760711b608082015260a00190565b60208082526038908201527f5b5145432d3031363030305d2d496e76616c69642076616c756520746f20696e60408201527731b932b0b9b2903a3432903837b7b6103130b630b731b29760411b606082015260800190565b602080825260159082015274676f7665726e616e63652e76616c696461746f727360581b604082015260600190565b60208082526045908201527f5b5145432d3031363030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520515661756c7420636f6e74726163742068617320616360608201526431b2b9b99760d91b608082015260a00190565b60208082526017908201527631b7b6b6b7b7173330b1ba37b93c9731b925b2b2b832b960491b604082015260600190565b600060e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b845481526001850154602082015260028501546040820152600385015460608201526004909401546001600160a01b0316608085015260a084019290925260c083015260e08201526101000190565b90815260200190565b918252602082015260400190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b6001600160a01b0381168114610d6457600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a26469706673582212201a5726f1a2d19dfcc9fa7baacaf7f73a66332d8a0ca140af8f9bc8a31b3c005064736f6c63430007060033",
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
