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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_claimerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferedToQVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structValidationRewardPools.ValidatorProperties\",\"name\":\"_v\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_inc\",\"type\":\"uint256\"}],\"name\":\"UpdateRate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorsProperties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addCompoundRateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"requestRewardTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"updateValidatorsCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelegatorsShare\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getDelegatorsShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"addAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"subAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateOfCompoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"}],\"internalType\":\"structValidationRewardPools.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getDelegatedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getLastUpdateOfCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetAmount\",\"type\":\"uint256\"}],\"name\":\"getNormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_normalizedAmount\",\"type\":\"uint256\"}],\"name\":\"getDenormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611aeb806100206000396000f3fe6080604052600436106100de5760003560e01c8063882b5dc811610085578063882b5dc8146102195780638cfe7df314610246578063aaaa646d14610266578063c4d66de814610286578063e268281c146102a6578063eab136a0146102c6578063f38f2ce2146102d9578063f8b2cb4f146102f9578063fd5d904d14610319576100de565b806306bfa938146100e35780630ce98a301461011957806333da982e1461014657806343da20f014610166578063443817f5146101885780634a9b319a146101b95780637c47c16d146101d9578063841c9064146101f9575b600080fd5b3480156100ef57600080fd5b506101036100fe3660046115cc565b610339565b6040516101109190611992565b60405180910390f35b34801561012557600080fd5b506101396101343660046115cc565b6105ce565b6040516101109190611a2f565b34801561015257600080fd5b50610139610161366004611604565b610657565b34801561017257600080fd5b5061018661018136600461164f565b6106aa565b005b34801561019457600080fd5b506101a86101a33660046115cc565b610809565b604051610110959493929190611a46565b3480156101c557600080fd5b506101866101d4366004611604565b610840565b3480156101e557600080fd5b506101396101f43660046115cc565b61093c565b34801561020557600080fd5b506101396102143660046115cc565b610969565b34801561022557600080fd5b50610239610234366004611604565b610987565b6040516101109190611693565b34801561025257600080fd5b506101866102613660046115cc565b610b96565b34801561027257600080fd5b506101396102813660046115cc565b610c68565b34801561029257600080fd5b506101866102a13660046115cc565b610cb1565b3480156102b257600080fd5b506101396102c1366004611604565b610d76565b6101866102d43660046115cc565b610db1565b3480156102e557600080fd5b506101866102f4366004611604565b610e0d565b34801561030557600080fd5b506101396103143660046115cc565b610ee6565b34801561032557600080fd5b506101396103343660046115cc565b610f01565b61034161158f565b6001600160a01b038083166000908152600160208181526040808420815160a0810183528154815293810154928401929092526002820154908301526003810154606083015260040154909216608083015261039c84610ff6565b9050806103a761158f565b816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103e057600080fd5b505afa1580156103f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104189190611667565b8160a0018181525050816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561045a57600080fd5b505afa15801561046e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104929190611667565b6080820152600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906104ca906004016118c7565b60206040518083038186803b1580156104e257600080fd5b505afa1580156104f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051a91906115e8565b6001600160a01b031663b9a55be9876040518263ffffffff1660e01b8152600401610545919061167f565b60206040518083038186803b15801561055d57600080fd5b505afa158015610571573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105959190611667565b60608083019190915284015160408083019190915284015160c0820152835181526020938401519381019390935250909150505b919050565b6000806105da83610ff6565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561061757600080fd5b505afa15801561062b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064f9190611667565b949350505050565b60008061066384610ff6565b604051631d335d6560e21b81529091506001600160a01b038216906374cd759490610692908690600401611a2f565b60206040518083038186803b15801561061757600080fd5b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906106dd906004016118c7565b60206040518083038186803b1580156106f557600080fd5b505afa158015610709573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072d91906115e8565b6001600160a01b031663f2f4a2f1336040518263ffffffff1660e01b8152600401610758919061167f565b60206040518083038186803b15801561077057600080fd5b505afa158015610784573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a8919061162f565b6107cd5760405162461bcd60e51b81526004016107c49061169e565b60405180910390fd5b6107d5611033565b8111156107f45760405162461bcd60e51b81526004016107c49061173e565b33600090815260016020526040902060020155565b600160208190526000918252604090912080549181015460028201546003830154600490930154919290916001600160a01b031685565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906108739060040161170f565b60206040518083038186803b15801561088b57600080fd5b505afa15801561089f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c391906115e8565b6001600160a01b0316336001600160a01b0316146108f35760405162461bcd60e51b81526004016107c4906118f6565b6001600160a01b0382166000908152600160205260409020600301546109199082611043565b6001600160a01b0390921660009081526001602052604090206003019190915550565b6001600160a01b038116600090815260016020526040812060030154610963908390610657565b92915050565b6001600160a01b031660009081526001602052604090206002015490565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906109bb9060040161170f565b60206040518083038186803b1580156109d357600080fd5b505afa1580156109e7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0b91906115e8565b6001600160a01b0316336001600160a01b031614610a3b5760405162461bcd60e51b81526004016107c4906118f6565b6001600160a01b03831660009081526001602052604090208054610a5f90846110a2565b81556001810154610a7090846110a2565b6001820155600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610aa89060040161170f565b60206040518083038186803b158015610ac057600080fd5b505afa158015610ad4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af891906115e8565b6001600160a01b03166331989b58846040518263ffffffff1660e01b81526004016000604051808303818588803b158015610b3257600080fd5b505af1158015610b46573d6000803e3d6000fd5b5050505050836001600160a01b03167fa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c84604051610b849190611a2f565b60405180910390a25060019392505050565b6001600160a01b0381811660009081526001602052604090206004015416610c6557610bc06110e4565b6001600160a01b031663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610bfa57600080fd5b505af1158015610c0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3291906115e8565b6001600160a01b03828116600090815260016020526040902060040180546001600160a01b031916929091169190911790555b50565b600080610c7483610ff6565b90506000816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561061757600080fd5b600054610100900460ff1680610cca5750610cca61116d565b80610cd8575060005460ff16155b610d135760405162461bcd60e51b815260040180806020018281038252602e815260200180611a88602e913960400191505060405180910390fd5b600054610100900460ff16158015610d3e576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610d72576000805461ff00191690555b5050565b600080610d8284610ff6565b604051638dc3311d60e01b81529091506001600160a01b03821690638dc3311d90610692908690600401611a2f565b34610dce5760405162461bcd60e51b81526004016107c49061186f565b6001600160a01b038116600090815260016020526040902054610df19034611043565b6001600160a01b03909116600090815260016020526040902055565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610e409060040161170f565b60206040518083038186803b158015610e5857600080fd5b505afa158015610e6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9091906115e8565b6001600160a01b0316336001600160a01b031614610ec05760405162461bcd60e51b81526004016107c4906118f6565b6001600160a01b03821660009081526001602052604090206003015461091990826110a2565b6001600160a01b031660009081526001602052604090205490565b600080610f0d83611173565b600054604051633fb9027160e01b81529192506201000090046001600160a01b031690633fb9027190610f42906004016118c7565b60206040518083038186803b158015610f5a57600080fd5b505afa158015610f6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f9291906115e8565b6001600160a01b031663f27f73c9846040518263ffffffff1660e01b8152600401610fbd919061167f565b600060405180830381600087803b158015610fd757600080fd5b505af1158015610feb573d6000803e3d6000fd5b509295945050505050565b6001600160a01b03808216600090815260016020526040812060040154909116806109635760405162461bcd60e51b81526004016107c49061178d565b6b033b2e3c9fd0803ce800000090565b60008282018381101561109b576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b600061109b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506114f8565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061111890600401611961565b60206040518083038186803b15801561113057600080fd5b505afa158015611144573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116891906115e8565b905090565b303b1590565b6001600160a01b038082166000908152600160208181526040808420815160a0810183528154815293810154928401929092526002820154908301526003810154606083015260040154909216608083015290816111d084610ff6565b90506000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561120d57600080fd5b505afa158015611221573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112459190611667565b90506000826001600160a01b03166374cd759485606001516040518263ffffffff1660e01b81526004016112799190611a2f565b60206040518083038186803b15801561129157600080fd5b505afa1580156112a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c99190611667565b602085015185519192506000916112df916110a2565b9050600181116112f65782955050505050506105c9565b60006001600160a01b0385166339c293f6846113138560016110a2565b6040518363ffffffff1660e01b8152600401611330929190611a38565b602060405180830381600087803b15801561134a57600080fd5b505af115801561135e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113829190611667565b90506000856001600160a01b03166374cd759488606001516040518263ffffffff1660e01b81526004016113b69190611a2f565b60206040518083038186803b1580156113ce57600080fd5b505afa1580156113e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114069190611667565b9050600061141482866110a2565b90506000886060015111801561142a5750600081115b15611495576000611453600161144d8b602001518561104390919063ffffffff16565b90611043565b89519091508111156114775760405162461bcd60e51b81526004016107c4906117f7565b6001600160a01b038b16600090815260016020819052604090912001555b6001600160a01b038a166000908152600160205260409081902090517f02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782916114e2918990879086906119e0565b60405180910390a1509098975050505050505050565b600081848411156115875760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561154c578181015183820152602001611534565b50505050905090810190601f1680156115795780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6000602082840312156115dd578081fd5b813561109b81611a72565b6000602082840312156115f9578081fd5b815161109b81611a72565b60008060408385031215611616578081fd5b823561162181611a72565b946020939093013593505050565b600060208284031215611640578081fd5b8151801515811461109b578182fd5b600060208284031215611660578081fd5b5035919050565b600060208284031215611678578081fd5b5051919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252604b908201527f5b5145432d3031363030355d2d5468652063616c6c6572206973206e6f74206160408201527f2076616c696461746f722c206661696c656420746f207365742064656c65676160608201526a3a37b9399039b430b9329760a91b608082015260a00190565b6020808252601590820152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b604082015260600190565b6020808252602f908201527f5b5145432d3031363030335d2d496e76616c69642064656c656761746f72207360408201526e3430b932903830b930b6b2ba32b91760891b606082015260800190565b60208082526044908201527f5b5145432d3031363030375d2d436f6d706f756e64526174654b65657065722060408201527f6e6f7420696e697469616c697a656420666f7220676976656e2076616c6964616060820152633a37b91760e11b608082015260a00190565b60208082526052908201527f5b5145432d3031363030325d2d5468657265206973206e6f7420656e6f75676860408201527f20706f6f6c2062616c616e63652c206661696c656420746f20757064617465206060820152713a34329031b7b6b837bab732103930ba329760711b608082015260a00190565b60208082526038908201527f5b5145432d3031363030305d2d496e76616c69642076616c756520746f20696e60408201527731b932b0b9b2903a3432903837b7b6103130b630b731b29760411b606082015260800190565b602080825260159082015274676f7665726e616e63652e76616c696461746f727360581b604082015260600190565b60208082526045908201527f5b5145432d3031363030315d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c792074686520515661756c7420636f6e74726163742068617320616360608201526431b2b9b99760d91b608082015260a00190565b60208082526017908201527631b7b6b6b7b7173330b1ba37b93c9731b925b2b2b832b960491b604082015260600190565b600060e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b845481526001850154602082015260028501546040820152600385015460608201526004909401546001600160a01b0316608085015260a084019290925260c083015260e08201526101000190565b90815260200190565b918252602082015260400190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b6001600160a01b0381168114610c6557600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a264697066735822122070d677a85b6c82a333e0c9c289445919fc2e8ff02d6cc1ea51135bef22d5967a64736f6c63430007060033",
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardPools *ValidationRewardPoolsSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Initialize(&_ValidationRewardPools.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_ValidationRewardPools *ValidationRewardPoolsTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _ValidationRewardPools.Contract.Initialize(&_ValidationRewardPools.TransactOpts, _registry)
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
