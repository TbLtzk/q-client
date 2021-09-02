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

// ValidationRewardPoolsABI is the input ABI used to generate the binding from.
const ValidationRewardPoolsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_claimerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferedToQVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structValidationRewardPools.ValidatorProperties\",\"name\":\"_v\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_inc\",\"type\":\"uint256\"}],\"name\":\"UpdateRate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorsProperties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"compoundRate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getDelegatedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getLastUpdateOfCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"addAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"subAggregatedNormalizedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getDelegatorsShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetAmount\",\"type\":\"uint256\"}],\"name\":\"getNormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_normalizedAmount\",\"type\":\"uint256\"}],\"name\":\"getDenormalizedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelegatorsShare\",\"type\":\"uint256\"}],\"name\":\"setDelegatorsShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addCompoundRateKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"requestRewardTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"updateValidatorsCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"poolBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedForClaims\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregatedNormalizedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateOfCompoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsShare\",\"type\":\"uint256\"}],\"internalType\":\"structValidationRewardPools.PoolInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidationRewardPoolsBin is the compiled bytecode used for deploying new contracts.
var ValidationRewardPoolsBin = "0x608060405234801561001057600080fd5b506123fe806100206000396000f3fe6080604052600436106100de5760003560e01c8063882b5dc811610085578063882b5dc8146102195780638cfe7df314610246578063aaaa646d14610266578063c4d66de814610286578063e268281c146102a6578063eab136a0146102c6578063f38f2ce2146102d9578063f8b2cb4f146102f9578063fd5d904d14610319576100de565b806306bfa938146100e35780630ce98a301461011957806333da982e1461014657806343da20f014610166578063443817f5146101885780634a9b319a146101b95780637c47c16d146101d9578063841c9064146101f9575b600080fd5b3480156100ef57600080fd5b506101036100fe3660046116d4565b610339565b6040516101109190611b34565b60405180910390f35b34801561012557600080fd5b506101396101343660046116d4565b6105ef565b6040516101109190611bd1565b34801561015257600080fd5b5061013961016136600461170c565b6106a4565b34801561017257600080fd5b50610186610181366004611757565b61076e565b005b34801561019457600080fd5b506101a86101a33660046116d4565b6108c4565b604051610110959493929190611be8565b3480156101c557600080fd5b506101866101d436600461170c565b6108fb565b3480156101e557600080fd5b506101396101f43660046116d4565b6109f7565b34801561020557600080fd5b506101396102143660046116d4565b610a24565b34801561022557600080fd5b5061023961023436600461170c565b610a42565b604051610110919061179b565b34801561025257600080fd5b506101866102613660046116d4565b610c51565b34801561027257600080fd5b506101396102813660046116d4565b610d85565b34801561029257600080fd5b506101866102a13660046116d4565b610de5565b3480156102b257600080fd5b506101396102c136600461170c565b610eaa565b6101866102d43660046116d4565b610f13565b3480156102e557600080fd5b506101866102f436600461170c565b610f6f565b34801561030557600080fd5b506101396103143660046116d4565b611048565b34801561032557600080fd5b506101396103343660046116d4565b611063565b61034161168a565b6001600160a01b03808316600090815260016020818152604092839020835160a08101855281548152928101549183019190915260028101549282019290925260038201546060820152600490910154909116608082018190526103a361168a565b6001600160a01b0382166103c9574260a08201526103bf611158565b60808201526104ba565b816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561040257600080fd5b505afa158015610416573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043a919061176f565b8160a0018181525050816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561047c57600080fd5b505afa158015610490573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104b4919061176f565b60808201525b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906104ed90600401611b05565b60206040518083038186803b15801561050557600080fd5b505afa158015610519573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053d91906116f0565b6001600160a01b031663b9a55be9866040518263ffffffff1660e01b81526004016105689190611787565b60206040518083038186803b15801561058057600080fd5b505afa158015610594573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b8919061176f565b60608083019190915283015160408083019190915283015160c082015282518152602092830151928101929092525090505b919050565b6001600160a01b0380821660009081526001602052604081206004015490911681610618611168565b90506001600160a01b0382161561069d57816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561066257600080fd5b505afa158015610676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069a919061176f565b90505b9392505050565b6001600160a01b03808316600090815260016020526040812060040154909116806106ea5760405162461bcd60e51b81526004016106e190611a20565b60405180910390fd5b604051631d335d6560e21b81526001600160a01b038216906374cd759490610716908690600401611bd1565b60206040518083038186803b15801561072e57600080fd5b505afa158015610742573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610766919061176f565b949350505050565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906107a190600401611b05565b60206040518083038186803b1580156107b957600080fd5b505afa1580156107cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f191906116f0565b6001600160a01b031663f2f4a2f1336040518263ffffffff1660e01b815260040161081c9190611787565b60206040518083038186803b15801561083457600080fd5b505afa158015610848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086c9190611737565b6108885760405162461bcd60e51b81526004016106e1906119b4565b610890611158565b8111156108af5760405162461bcd60e51b81526004016106e190611848565b33600090815260016020526040902060020155565b600160208190526000918252604090912080549181015460028201546003830154600490930154919290916001600160a01b031685565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061092e90600401611819565b60206040518083038186803b15801561094657600080fd5b505afa15801561095a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097e91906116f0565b6001600160a01b0316336001600160a01b0316146109ae5760405162461bcd60e51b81526004016106e190611a8d565b6001600160a01b0382166000908152600160205260409020600301546109d49082611174565b6001600160a01b0390921660009081526001602052604090206003019190915550565b6001600160a01b038116600090815260016020526040812060030154610a1e9083906106a4565b92915050565b6001600160a01b031660009081526001602052604090206002015490565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610a7690600401611819565b60206040518083038186803b158015610a8e57600080fd5b505afa158015610aa2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac691906116f0565b6001600160a01b0316336001600160a01b031614610af65760405162461bcd60e51b81526004016106e190611a8d565b6001600160a01b03831660009081526001602052604090208054610b1a90846111cc565b81556001810154610b2b90846111cc565b6001820155600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610b6390600401611819565b60206040518083038186803b158015610b7b57600080fd5b505afa158015610b8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb391906116f0565b6001600160a01b03166331989b58846040518263ffffffff1660e01b81526004016000604051808303818588803b158015610bed57600080fd5b505af1158015610c01573d6000803e3d6000fd5b5050505050836001600160a01b03167fa3e662ac91ce446097ba1e12083429fc4016b94f6a31cd0dd9baaefd1f8afb6c84604051610c3f9190611bd1565b60405180910390a25060019392505050565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610c8490600401611b05565b60206040518083038186803b158015610c9c57600080fd5b505afa158015610cb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd491906116f0565b6001600160a01b0316336001600160a01b031614610d045760405162461bcd60e51b81526004016106e1906118c0565b6001600160a01b0381811660009081526001602052604090206004015416610d8257604051610d32906116c7565b604051809103906000f080158015610d4e573d6000803e3d6000fd5b506001600160a01b03828116600090815260016020526040902060040180546001600160a01b031916929091169190911790555b50565b6001600160a01b0380821660009081526001602052604081206004015490911642811561069d57816001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561066257600080fd5b600054610100900460ff1680610dfe5750610dfe61120e565b80610e0c575060005460ff16155b610e475760405162461bcd60e51b815260040180806020018281038252602e81526020018061239b602e913960400191505060405180910390fd5b600054610100900460ff16158015610e72576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610ea6576000805461ff00191690555b5050565b6001600160a01b0380831660009081526001602052604081206004015490911680610ee75760405162461bcd60e51b81526004016106e190611a20565b604051638dc3311d60e01b81526001600160a01b03821690638dc3311d90610716908690600401611bd1565b34610f305760405162461bcd60e51b81526004016106e1906117a6565b6001600160a01b038116600090815260016020526040902054610f539034611174565b6001600160a01b03909116600090815260016020526040902055565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610fa290600401611819565b60206040518083038186803b158015610fba57600080fd5b505afa158015610fce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff291906116f0565b6001600160a01b0316336001600160a01b0316146110225760405162461bcd60e51b81526004016106e190611a8d565b6001600160a01b0382166000908152600160205260409020600301546109d490826111cc565b6001600160a01b031660009081526001602052604090205490565b60008061106f83611214565b600054604051633fb9027160e01b81529192506201000090046001600160a01b031690633fb90271906110a490600401611b05565b60206040518083038186803b1580156110bc57600080fd5b505afa1580156110d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f491906116f0565b6001600160a01b031663f27f73c9846040518263ffffffff1660e01b815260040161111f9190611787565b600060405180830381600087803b15801561113957600080fd5b505af115801561114d573d6000803e3d6000fd5b509295945050505050565b6b033b2e3c9fd0803ce800000090565b670de0b6b3a764000090565b60008282018381101561069d576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b600061069d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506115f3565b303b1590565b6001600160a01b038082166000908152600160208181526040808420815160a081018352815481529381015492840192909252600282015490830152600381015460608301526004015490921660808301819052909190806112cd5760405161127c906116c7565b604051809103906000f080158015611298573d6000803e3d6000fd5b506001600160a01b03858116600090815260016020526040902060040180546001600160a01b03191691831691909117905590505b6000816001600160a01b031663f7fb07b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561130857600080fd5b505afa15801561131c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611340919061176f565b90506000826001600160a01b03166374cd759485606001516040518263ffffffff1660e01b81526004016113749190611bd1565b60206040518083038186803b15801561138c57600080fd5b505afa1580156113a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c4919061176f565b602085015185519192506000916113da916111cc565b9050600181116113f15782955050505050506105ea565b60006001600160a01b0385166339c293f68461140e8560016111cc565b6040518363ffffffff1660e01b815260040161142b929190611bda565b602060405180830381600087803b15801561144557600080fd5b505af1158015611459573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147d919061176f565b90506000856001600160a01b03166374cd759488606001516040518263ffffffff1660e01b81526004016114b19190611bd1565b60206040518083038186803b1580156114c957600080fd5b505afa1580156114dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611501919061176f565b9050600061150f82866111cc565b9050600088606001511180156115255750600081115b1561159057600061154e60016115488b602001518561117490919063ffffffff16565b90611174565b89519091508111156115725760405162461bcd60e51b81526004016106e19061193c565b6001600160a01b038b16600090815260016020819052604090912001555b6001600160a01b038a166000908152600160205260409081902090517f02f03399427b8ef86638c5f1c7a4eb80b07b6e158f61b5621aee1658638a5782916115dd91899087908690611b82565b60405180910390a1509098975050505050505050565b600081848411156116825760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561164757818101518382015260200161162f565b50505050905090810190601f1680156116745780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6040518060e00160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b61077180611c2a83390190565b6000602082840312156116e5578081fd5b813561069d81611c14565b600060208284031215611701578081fd5b815161069d81611c14565b6000806040838503121561171e578081fd5b823561172981611c14565b946020939093013593505050565b600060208284031215611748578081fd5b8151801515811461069d578182fd5b600060208284031215611768578081fd5b5035919050565b600060208284031215611780578081fd5b5051919050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020808252604d908201527f5b5145432d3031363030305d2d496e76616c69642076616c756520746f20696e60408201527f6372656173652074686520706f6f6c2062616c616e63652c2074686520696e6360608201526c3932b0b9b2903330b4b632b21760991b608082015260a00190565b6020808252601590820152741d1bdad95b9958dbdb9bdb5a58dccb9c55985d5b1d605a1b604082015260600190565b60208082526052908201527f5b5145432d3031363030335d2d496e76616c69642076616c756520666f72207460408201527f68652064656c656761746f722073686172652c206661696c656420746f2073656060820152713a103232b632b3b0ba37b91039b430b9329760711b608082015260a00190565b60208082526056908201527f5b5145432d3031363030355d2d5468652063616c6c657220646f6573206e6f7460408201527f2068617665206163636573732c206f6e6c79207468652056616c696461746f72606082015275399031b7b73a3930b1ba103430b99030b1b1b2b9b99760511b608082015260a00190565b60208082526052908201527f5b5145432d3031363030325d2d5468657265206973206e6f7420656e6f75676860408201527f20706f6f6c2062616c616e63652c206661696c656420746f20757064617465206060820152713a34329031b7b6b837bab732103930ba329760711b608082015260a00190565b60208082526046908201527f5b5145432d3031363030355d2d546865207374616b6520646f6573206e6f742060408201527f65786973742c206661696c656420746f207365742064656c656761746f72732060608201526539b430b9329760d11b608082015260a00190565b60208082526047908201527f5b5145432d3031363030375d2d54686520676976656e2061646472657373206960408201527f73206e6f742061207265676973746572656420436f6d706f756e64526174652060608201526625b2b2b832b91760c91b608082015260a00190565b60208082526052908201527f5b5145432d3031363030315d2d5468652063616c6c657220646f6573206e6f7460408201527f2068617665206163636573732c206f6e6c792074686520515661756c7420636f606082015271373a3930b1ba103430b99030b1b1b2b9b99760711b608082015260a00190565b602080825260159082015274676f7665726e616e63652e76616c696461746f727360581b604082015260600190565b600060e082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b845481526001850154602082015260028501546040820152600385015460608201526004909401546001600160a01b0316608085015260a084019290925260c083015260e08201526101000190565b90815260200190565b918252602082015260400190565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b6001600160a01b0381168114610d8257600080fdfe608060405234801561001057600080fd5b50610019610037565b60005542600155600280546001600160a01b03191633179055610043565b670de0b6b3a764000090565b61071f806100526000396000f3fe608060405234801561001057600080fd5b50600436106100785760003560e01c806339c293f61461007d5780634c89867f146100a657806366425d36146100ae57806374cd7594146100c457806382ab890a146100d75780638dc3311d146100ea578063f2fde38b146100fd578063f7fb07b014610112575b600080fd5b61009061008b36600461066a565b61011a565b60405161009d91906106b1565b60405180910390f35b61009061019c565b6100b66101a2565b60405161009d9291906106ba565b6100906100d2366004610652565b6101ab565b6100906100e5366004610652565b6101ce565b6100906100f8366004610652565b610271565b61011061010b36600461062b565b6102d6565b005b610090610322565b6002546000906001600160a01b031633146101505760405162461bcd60e51b81526004016101479061068b565b60405180910390fd5b81158061015b575082155b156101695750600054610196565b6000805461018d908590610187906101818388610328565b90610387565b906103e0565b60008190559150505b92915050565b60015490565b60005460015482565b60006101c66101b861041f565b600054610187908590610387565b90505b919050565b6002546000906001600160a01b031633146101fb5760405162461bcd60e51b81526004016101479061068b565b600154421161020d57506000546101c9565b600061021761041f565b60015490915060009061022b90429061042f565b905060006102538361018761024a6102438984610328565b8688610471565b60005490610387565b6000549091508114610269576000819055426001555b949350505050565b60008061028e600080015461018761028761041f565b8690610387565b90506000805b6102ad61029f61041f565b600054610187908690610387565b90508481116102ba578291505b6102c5836001610328565b925084811061029457509392505050565b6002546001600160a01b031633146103005760405162461bcd60e51b81526004016101479061068b565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b60005490565b600082820183811015610380576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b60008261039657506000610196565b828202828482816103a357fe5b04146103805760405162461bcd60e51b81526004018080602001828103825260218152602001806106c96021913960400191505060405180910390fd5b600061038083836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b81525061052f565b6b033b2e3c9fd0803ce800000090565b600061038083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506105d1565b60008380156105115760018416801561048c57859250610490565b8392505b50600283046002850494505b841561050b5785860286878204146104b357600080fd5b818101818110156104c357600080fd5b85900496505060018516156105005785830283878204141587151516156104e957600080fd5b818101818110156104f957600080fd5b8590049350505b60028504945061049c565b50610527565b8380156105215760009250610525565b8392505b505b509392505050565b600081836105bb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610580578181015183820152602001610568565b50505050905090810190601f1680156105ad5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816105c757fe5b0495945050505050565b600081848411156106235760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610580578181015183820152602001610568565b505050900390565b60006020828403121561063c578081fd5b81356001600160a01b0381168114610380578182fd5b600060208284031215610663578081fd5b5035919050565b6000806040838503121561067c578081fd5b50508035926020909101359150565b6020808252600c908201526b155b985d5d1a1bdc9a5e995960a21b604082015260600190565b90815260200190565b91825260208201526040019056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212200983c2c2c4b846af6d5b77e7cd6482c3d38a79f65aaef20e06d23b654994005864736f6c63430007060033496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a26469706673582212203a23097b8d4af1e851d93598340276eec66a065d876f96c881af0a51295a062f64736f6c63430007060033"

// DeployValidationRewardPools deploys a new Ethereum contract, binding an instance of ValidationRewardPools to it.
func DeployValidationRewardPools(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidationRewardPools, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidationRewardPoolsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidationRewardPoolsBin), backend)
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
func (_ValidationRewardPools *ValidationRewardPoolsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ValidationRewardPools *ValidationRewardPoolsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getBalance", _validator)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getCompoundRate", _validator)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getDelegatedStake", _validator)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getDelegatorsShare", _addr)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getDenormalizedAmount", _validator, _normalizedAmount)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getLastUpdateOfCompoundRate", _validator)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getNormalizedAmount", _validator, _targetAmount)
	return *ret0, err
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
	var (
		ret0 = new(ValidationRewardPoolsPoolInfo)
	)
	out := ret0
	err := _ValidationRewardPools.contract.Call(opts, out, "getPoolInfo", _validator)
	return *ret0, err
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
	ret := new(struct {
		Balance                   *big.Int
		ReservedForClaim          *big.Int
		DelegatorsShare           *big.Int
		AggregatedNormalizedStake *big.Int
		CompoundRate              common.Address
	})
	out := ret
	err := _ValidationRewardPools.contract.Call(opts, out, "validatorsProperties", arg0)
	return *ret, err
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
	return event, nil
}
