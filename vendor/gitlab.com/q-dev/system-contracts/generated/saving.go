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

// SavingBalanceDetails is an auto generated low-level Go binding around an user-defined struct.
type SavingBalanceDetails struct {
	CurrentBalance           *big.Int
	NormalizedBalance        *big.Int
	CompoundRate             *big.Int
	LastUpdateOfCompoundRate *big.Int
	InterestRate             *big.Int
}

// SavingMetaData contains all meta data concerning the Saving contract.
var SavingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"UserDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"UserWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"aggregatedNormalizedCapital\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compoundRateKeeper\",\"outputs\":[{\"internalType\":\"contractCompoundRateKeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"normalizedCapitals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCompoundRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"normalizedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"compoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateOfCompoundRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"internalType\":\"structSaving.BalanceDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e78806100206000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806312065fe0146100885780632e1a7d4d146100a35780633d4d5d4d146100c657806364eb4369146100e65780636fa0782414610130578063b6b55f2514610139578063c0ab9cbc1461014c578063cb5e79e014610154578063f399e22e14610174575b600080fd5b610090610189565b6040519081526020015b60405180910390f35b6100b66100b136600461196b565b610219565b604051901515815260200161009a565b6002546100d9906001600160a01b031681565b60405161009a9190611984565b6100ee610556565b60405161009a9190600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b61009060035481565b6100b661014736600461196b565b6107a4565b610090610a53565b6100906101623660046119b0565b60046020526000908152604090205481565b610187610182366004611a1b565b611260565b005b600254336000908152600460208190526040808320549051631d335d6560e21b81529182015290916001600160a01b0316906374cd75949060240160206040518083038186803b1580156101dc57600080fd5b505afa1580156101f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102149190611ac3565b905090565b600080610224610189565b90508061029e5760405162461bcd60e51b815260206004820152603e60248201527f5b5145432d3032323030315d2d5468652063616c6c657220646f6573206e6f7460448201527f206861766520616e792062616c616e636520746f2077697468647261772e000060648201526084015b60405180910390fd5b82808210156102aa5750805b6002546000906001600160a01b0316638dc3311d6102c88486611af2565b6040518263ffffffff1660e01b81526004016102e691815260200190565b60206040518083038186803b1580156102fe57600080fd5b505afa158015610312573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103369190611ac3565b33600090815260046020526040902054909150610354908290611af2565b6003546103619190611af2565b6003553360009081526004602052604081208290555460018054620100009092046001600160a01b031691633fb90271916104229161039f90611b09565b80601f01602080910402602001604051908101604052809291908181526020018280546103cb90611b09565b80156104185780601f106103ed57610100808354040283529160200191610418565b820191906000526020600020905b8154815290600101906020018083116103fb57829003601f168201915b505050505061148c565b6040518263ffffffff1660e01b815260040161043e9190611b74565b60206040518083038186803b15801561045657600080fd5b505afa15801561046a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048e9190611ba7565b6001600160a01b031663a9059cbb33846040518363ffffffff1660e01b81526004016104bb929190611bc4565b602060405180830381600087803b1580156104d557600080fd5b505af11580156104e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050d9190611bdd565b5060405182815233907fe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb9060200160405180910390a261054b6114f7565b506001949350505050565b61055e6118a3565b6105666118a3565b60025460408051630f7fb07b60e41b815290516001600160a01b0390921691829163f7fb07b0916004808301926020929190829003018186803b1580156105ac57600080fd5b505afa1580156105c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e49190611ac3565b826040018181525050806001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062657600080fd5b505afa15801561063a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065e9190611ac3565b606083015261066b610189565b8252336000908152600460209081526040909120549083015261068c6116ac565b6001600160a01b031663498bff0061072d600180546106aa90611b09565b80601f01602080910402602001604051908101604052809291908181526020018280546106d690611b09565b80156107235780601f106106f857610100808354040283529160200191610723565b820191906000526020600020905b81548152906001019060200180831161070657829003601f168201915b5050505050611756565b6040518263ffffffff1660e01b81526004016107499190611b74565b60206040518083038186803b15801561076157600080fd5b505afa158015610775573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107999190611ac3565b608083015250919050565b600080821161080b5760405162461bcd60e51b815260206004820152602d60248201527f5b5145432d3032323030305d2d4465706f73697420616d6f756e74206d75737460448201526c103737ba103132903d32b9379760991b6064820152608401610295565b600060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027161083f6001805461039f90611b09565b6040518263ffffffff1660e01b815260040161085b9190611b74565b60206040518083038186803b15801561087357600080fd5b505afa158015610887573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ab9190611ba7565b6040516323b872dd60e01b8152336004820152306024820152604481018490526001600160a01b0391909116906323b872dd90606401602060405180830381600087803b1580156108fb57600080fd5b505af115801561090f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109339190611bdd565b506002546000906001600160a01b0316638dc3311d84610951610189565b61095b9190611bff565b6040518263ffffffff1660e01b815260040161097991815260200190565b60206040518083038186803b15801561099157600080fd5b505afa1580156109a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c99190611ac3565b336000908152600460205260409020549091506109e69082611af2565b6003546109f39190611bff565b6003553360008181526004602052604090819020839055517f951fdc61d6a98f96098a17ea6ac287a6fd38aea6bef73083c93b274cb830107d90610a3a9086815260200190565b60405180910390a2610a4a6114f7565b50600192915050565b600254600180546000926001600160a01b0316918391610a7290611b09565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9e90611b09565b8015610aeb5780601f10610ac057610100808354040283529160200191610aeb565b820191906000526020600020905b815481529060010190602001808311610ace57829003601f168201915b505050505090506000610afc6116ac565b6001600160a01b031663498bff00610b1384611756565b6040518263ffffffff1660e01b8152600401610b2f9190611b74565b60206040518083038186803b158015610b4757600080fd5b505afa158015610b5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7f9190611ac3565b90506000836001600160a01b03166374cd75946003546040518263ffffffff1660e01b8152600401610bb391815260200190565b60206040518083038186803b158015610bcb57600080fd5b505afa158015610bdf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c039190611ac3565b604051634155c48560e11b8152600481018490529091506000906001600160a01b038616906382ab890a90602401602060405180830381600087803b158015610c4b57600080fd5b505af1158015610c5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c839190611ac3565b90506000856001600160a01b03166374cd75946003546040518263ffffffff1660e01b8152600401610cb791815260200190565b60206040518083038186803b158015610ccf57600080fd5b505afa158015610ce3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d079190611ac3565b60008054919250906201000090046001600160a01b0316633fb90271610d2c8861178a565b6040518263ffffffff1660e01b8152600401610d489190611b74565b60206040518083038186803b158015610d6057600080fd5b505afa158015610d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d989190611ba7565b90506000816001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b158015610dd557600080fd5b505afa158015610de9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e0d9190611ac3565b90506000610e1b8685611af2565b9050808210156111cf57600080546201000090046001600160a01b0316633fb90271610e468b6117df565b6040518263ffffffff1660e01b8152600401610e629190611b74565b60206040518083038186803b158015610e7a57600080fd5b505afa158015610e8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb29190611ba7565b6001600160a01b0316635ac2b2ee6040518163ffffffff1660e01b815260040160606040518083038186803b158015610eea57600080fd5b505afa158015610efe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f229190611c17565b90506000610f308484611af2565b9050816040015181866001600160a01b03166314a6bf0f6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f7157600080fd5b505afa158015610f85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa99190611ac3565b610fb39190611bff565b11156110455760405162461bcd60e51b815260206004820152605560248201527f5b5145432d3032323030325d2d53797374656d2064656274206578636565647360448201527f206f77656420626f72726f77696e6720666565732c206661696c656420746f206064820152743ab83230ba329031b7b6b837bab732103930ba329760591b608482015260a401610295565b6000546201000090046001600160a01b0316633fb902716110658c61148c565b6040518263ffffffff1660e01b81526004016110819190611b74565b60206040518083038186803b15801561109957600080fd5b505afa1580156110ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d19190611ba7565b6001600160a01b03166340c10f1986836040518363ffffffff1660e01b81526004016110fe929190611bc4565b602060405180830381600087803b15801561111857600080fd5b505af115801561112c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111509190611bdd565b50604051632b7c7b1160e01b8152600481018290526001600160a01b03861690632b7c7b1190602401602060405180830381600087803b15801561119357600080fd5b505af11580156111a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111cb9190611bdd565b5050505b60405163effd85ad60e01b8152600481018290526001600160a01b0384169063effd85ad90602401602060405180830381600087803b15801561121157600080fd5b505af1158015611225573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112499190611bdd565b506112526114f7565b509298975050505050505050565b600054610100900460ff1680611279575060005460ff16155b6112dc5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610295565b600054610100900460ff161580156112fe576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905581516113339060019060208501906118d2565b50600054604080518082018252601781527631b7b6b6b7b7173330b1ba37b93c9731b925b2b2b832b960491b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161139391600401611b74565b60206040518083038186803b1580156113ab57600080fd5b505afa1580156113bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e39190611ba7565b6001600160a01b031663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561141d57600080fd5b505af1158015611431573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114559190611ba7565b600280546001600160a01b0319166001600160a01b03929092169190911790558015611487576000805461ff00191690555b505050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946114e19493879201611c73565b6040516020818303038152906040529050919050565b600254600354604051631d335d6560e21b815260048101919091526000916001600160a01b0316906374cd75949060240160206040518083038186803b15801561154057600080fd5b505afa158015611554573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115789190611ac3565b905060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb902716115af6001805461039f90611b09565b6040518263ffffffff1660e01b81526004016115cb9190611b74565b60206040518083038186803b1580156115e357600080fd5b505afa1580156115f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161b9190611ba7565b6001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016116469190611984565b60206040518083038186803b15801561165e57600080fd5b505afa158015611672573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116969190611ac3565b9050818110156116a8576116a8611cd9565b5050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611e21602291396040518263ffffffff1660e01b81526004016117069190611b74565b60206040518083038186803b15801561171e57600080fd5b505afa158015611732573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102149190611ba7565b6060611784826040518060400160405280600a815260200169736176696e675261746560b01b815250611834565b92915050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946114e19493879201611cef565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b828401819052855180870187529182528184015293516060946114e19493879201611d5e565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b8152508460405160200161188c9493929190611dc9565b604051602081830303815290604052905092915050565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b8280546118de90611b09565b90600052602060002090601f0160209004810192826119005760008555611946565b82601f1061191957805160ff1916838001178555611946565b82800160010185558215611946579182015b8281111561194657825182559160200191906001019061192b565b50611952929150611956565b5090565b5b808211156119525760008155600101611957565b60006020828403121561197d57600080fd5b5035919050565b6001600160a01b0391909116815260200190565b6001600160a01b03811681146119ad57600080fd5b50565b6000602082840312156119c257600080fd5b81356119cd81611998565b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611a1357611a136119d4565b604052919050565b60008060408385031215611a2e57600080fd5b8235611a3981611998565b915060208381013567ffffffffffffffff80821115611a5757600080fd5b818601915086601f830112611a6b57600080fd5b813581811115611a7d57611a7d6119d4565b611a8f601f8201601f191685016119ea565b91508082528784828501011115611aa557600080fd5b80848401858401376000848284010152508093505050509250929050565b600060208284031215611ad557600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600082821015611b0457611b04611adc565b500390565b600181811c90821680611b1d57607f821691505b60208210811415611b3e57634e487b7160e01b600052602260045260246000fd5b50919050565b60005b83811015611b5f578181015183820152602001611b47565b83811115611b6e576000848401525b50505050565b6020815260008251806020840152611b93816040850160208701611b44565b601f01601f19169190910160400192915050565b600060208284031215611bb957600080fd5b81516119cd81611998565b6001600160a01b03929092168252602082015260400190565b600060208284031215611bef57600080fd5b815180151581146119cd57600080fd5b60008219821115611c1257611c12611adc565b500190565b600060608284031215611c2957600080fd5b6040516060810181811067ffffffffffffffff82111715611c4c57611c4c6119d4565b80604052508251815260208301516020820152604083015160408201528091505092915050565b60008551611c85818460208a01611b44565b855190830190611c99818360208a01611b44565b8551910190611cac818360208901611b44565b8451910190611cbf818360208801611b44565b6331b7b4b760e11b91019081526004019695505050505050565b634e487b7160e01b600052600160045260246000fd5b60008551611d01818460208a01611b44565b855190830190611d15818360208a01611b44565b8551910190611d28818360208901611b44565b8451910190611d3b818360208801611b44565b6c73797374656d42616c616e636560981b9101908152600d019695505050505050565b60008551611d70818460208a01611b44565b855190830190611d84818360208a01611b44565b8551910190611d97818360208901611b44565b8451910190611daa818360208801611b44565b68626f72726f77696e6760b81b91019081526009019695505050505050565b60008551611ddb818460208a01611b44565b855190830190611def818360208a01611b44565b8551910190611e02818360208901611b44565b8451910190611e15818360208801611b44565b01969550505050505056fe676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a2646970667358221220a51ee5a4a45ca235ab33d16540058c1e869d58411ac525547ccc7e8bea2b495564736f6c63430008090033",
}

// SavingABI is the input ABI used to generate the binding from.
// Deprecated: Use SavingMetaData.ABI instead.
var SavingABI = SavingMetaData.ABI

// SavingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SavingMetaData.Bin instead.
var SavingBin = SavingMetaData.Bin

// DeploySaving deploys a new Ethereum contract, binding an instance of Saving to it.
func DeploySaving(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Saving, error) {
	parsed, err := SavingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SavingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Saving{SavingCaller: SavingCaller{contract: contract}, SavingTransactor: SavingTransactor{contract: contract}, SavingFilterer: SavingFilterer{contract: contract}}, nil
}

// Saving is an auto generated Go binding around an Ethereum contract.
type Saving struct {
	SavingCaller     // Read-only binding to the contract
	SavingTransactor // Write-only binding to the contract
	SavingFilterer   // Log filterer for contract events
}

// SavingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SavingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SavingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SavingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SavingSession struct {
	Contract     *Saving           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SavingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SavingCallerSession struct {
	Contract *SavingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SavingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SavingTransactorSession struct {
	Contract     *SavingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SavingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SavingRaw struct {
	Contract *Saving // Generic contract binding to access the raw methods on
}

// SavingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SavingCallerRaw struct {
	Contract *SavingCaller // Generic read-only contract binding to access the raw methods on
}

// SavingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SavingTransactorRaw struct {
	Contract *SavingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaving creates a new instance of Saving, bound to a specific deployed contract.
func NewSaving(address common.Address, backend bind.ContractBackend) (*Saving, error) {
	contract, err := bindSaving(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Saving{SavingCaller: SavingCaller{contract: contract}, SavingTransactor: SavingTransactor{contract: contract}, SavingFilterer: SavingFilterer{contract: contract}}, nil
}

// NewSavingCaller creates a new read-only instance of Saving, bound to a specific deployed contract.
func NewSavingCaller(address common.Address, caller bind.ContractCaller) (*SavingCaller, error) {
	contract, err := bindSaving(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SavingCaller{contract: contract}, nil
}

// NewSavingTransactor creates a new write-only instance of Saving, bound to a specific deployed contract.
func NewSavingTransactor(address common.Address, transactor bind.ContractTransactor) (*SavingTransactor, error) {
	contract, err := bindSaving(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SavingTransactor{contract: contract}, nil
}

// NewSavingFilterer creates a new log filterer instance of Saving, bound to a specific deployed contract.
func NewSavingFilterer(address common.Address, filterer bind.ContractFilterer) (*SavingFilterer, error) {
	contract, err := bindSaving(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SavingFilterer{contract: contract}, nil
}

// bindSaving binds a generic wrapper to an already deployed contract.
func bindSaving(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SavingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Saving *SavingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Saving.Contract.SavingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Saving *SavingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Saving.Contract.SavingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Saving *SavingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Saving.Contract.SavingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Saving *SavingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Saving.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Saving *SavingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Saving.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Saving *SavingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Saving.Contract.contract.Transact(opts, method, params...)
}

// AggregatedNormalizedCapital is a free data retrieval call binding the contract method 0x6fa07824.
//
// Solidity: function aggregatedNormalizedCapital() view returns(uint256)
func (_Saving *SavingCaller) AggregatedNormalizedCapital(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Saving.contract.Call(opts, &out, "aggregatedNormalizedCapital")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregatedNormalizedCapital is a free data retrieval call binding the contract method 0x6fa07824.
//
// Solidity: function aggregatedNormalizedCapital() view returns(uint256)
func (_Saving *SavingSession) AggregatedNormalizedCapital() (*big.Int, error) {
	return _Saving.Contract.AggregatedNormalizedCapital(&_Saving.CallOpts)
}

// AggregatedNormalizedCapital is a free data retrieval call binding the contract method 0x6fa07824.
//
// Solidity: function aggregatedNormalizedCapital() view returns(uint256)
func (_Saving *SavingCallerSession) AggregatedNormalizedCapital() (*big.Int, error) {
	return _Saving.Contract.AggregatedNormalizedCapital(&_Saving.CallOpts)
}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Saving *SavingCaller) CompoundRateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Saving.contract.Call(opts, &out, "compoundRateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Saving *SavingSession) CompoundRateKeeper() (common.Address, error) {
	return _Saving.Contract.CompoundRateKeeper(&_Saving.CallOpts)
}

// CompoundRateKeeper is a free data retrieval call binding the contract method 0x3d4d5d4d.
//
// Solidity: function compoundRateKeeper() view returns(address)
func (_Saving *SavingCallerSession) CompoundRateKeeper() (common.Address, error) {
	return _Saving.Contract.CompoundRateKeeper(&_Saving.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Saving *SavingCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Saving.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Saving *SavingSession) GetBalance() (*big.Int, error) {
	return _Saving.Contract.GetBalance(&_Saving.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Saving *SavingCallerSession) GetBalance() (*big.Int, error) {
	return _Saving.Contract.GetBalance(&_Saving.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Saving *SavingCaller) GetBalanceDetails(opts *bind.CallOpts) (SavingBalanceDetails, error) {
	var out []interface{}
	err := _Saving.contract.Call(opts, &out, "getBalanceDetails")

	if err != nil {
		return *new(SavingBalanceDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(SavingBalanceDetails)).(*SavingBalanceDetails)

	return out0, err

}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Saving *SavingSession) GetBalanceDetails() (SavingBalanceDetails, error) {
	return _Saving.Contract.GetBalanceDetails(&_Saving.CallOpts)
}

// GetBalanceDetails is a free data retrieval call binding the contract method 0x64eb4369.
//
// Solidity: function getBalanceDetails() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Saving *SavingCallerSession) GetBalanceDetails() (SavingBalanceDetails, error) {
	return _Saving.Contract.GetBalanceDetails(&_Saving.CallOpts)
}

// NormalizedCapitals is a free data retrieval call binding the contract method 0xcb5e79e0.
//
// Solidity: function normalizedCapitals(address ) view returns(uint256)
func (_Saving *SavingCaller) NormalizedCapitals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Saving.contract.Call(opts, &out, "normalizedCapitals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalizedCapitals is a free data retrieval call binding the contract method 0xcb5e79e0.
//
// Solidity: function normalizedCapitals(address ) view returns(uint256)
func (_Saving *SavingSession) NormalizedCapitals(arg0 common.Address) (*big.Int, error) {
	return _Saving.Contract.NormalizedCapitals(&_Saving.CallOpts, arg0)
}

// NormalizedCapitals is a free data retrieval call binding the contract method 0xcb5e79e0.
//
// Solidity: function normalizedCapitals(address ) view returns(uint256)
func (_Saving *SavingCallerSession) NormalizedCapitals(arg0 common.Address) (*big.Int, error) {
	return _Saving.Contract.NormalizedCapitals(&_Saving.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_Saving *SavingTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Saving.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_Saving *SavingSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Saving.Contract.Deposit(&_Saving.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns(bool)
func (_Saving *SavingTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Saving.Contract.Deposit(&_Saving.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_Saving *SavingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _Saving.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_Saving *SavingSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _Saving.Contract.Initialize(&_Saving.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_Saving *SavingTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _Saving.Contract.Initialize(&_Saving.TransactOpts, _registry, _stc)
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Saving *SavingTransactor) UpdateCompoundRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Saving.contract.Transact(opts, "updateCompoundRate")
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Saving *SavingSession) UpdateCompoundRate() (*types.Transaction, error) {
	return _Saving.Contract.UpdateCompoundRate(&_Saving.TransactOpts)
}

// UpdateCompoundRate is a paid mutator transaction binding the contract method 0xc0ab9cbc.
//
// Solidity: function updateCompoundRate() returns(uint256)
func (_Saving *SavingTransactorSession) UpdateCompoundRate() (*types.Transaction, error) {
	return _Saving.Contract.UpdateCompoundRate(&_Saving.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Saving *SavingTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Saving.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Saving *SavingSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Saving.Contract.Withdraw(&_Saving.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns(bool)
func (_Saving *SavingTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Saving.Contract.Withdraw(&_Saving.TransactOpts, _amount)
}

// SavingUserDepositedIterator is returned from FilterUserDeposited and is used to iterate over the raw logs and unpacked data for UserDeposited events raised by the Saving contract.
type SavingUserDepositedIterator struct {
	Event *SavingUserDeposited // Event containing the contract specifics and raw log

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
func (it *SavingUserDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingUserDeposited)
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
		it.Event = new(SavingUserDeposited)
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
func (it *SavingUserDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingUserDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingUserDeposited represents a UserDeposited event raised by the Saving contract.
type SavingUserDeposited struct {
	User          common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserDeposited is a free log retrieval operation binding the contract event 0x951fdc61d6a98f96098a17ea6ac287a6fd38aea6bef73083c93b274cb830107d.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount)
func (_Saving *SavingFilterer) FilterUserDeposited(opts *bind.FilterOpts, _user []common.Address) (*SavingUserDepositedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Saving.contract.FilterLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return &SavingUserDepositedIterator{contract: _Saving.contract, event: "UserDeposited", logs: logs, sub: sub}, nil
}

// WatchUserDeposited is a free log subscription operation binding the contract event 0x951fdc61d6a98f96098a17ea6ac287a6fd38aea6bef73083c93b274cb830107d.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount)
func (_Saving *SavingFilterer) WatchUserDeposited(opts *bind.WatchOpts, sink chan<- *SavingUserDeposited, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Saving.contract.WatchLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingUserDeposited)
				if err := _Saving.contract.UnpackLog(event, "UserDeposited", log); err != nil {
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

// ParseUserDeposited is a log parse operation binding the contract event 0x951fdc61d6a98f96098a17ea6ac287a6fd38aea6bef73083c93b274cb830107d.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount)
func (_Saving *SavingFilterer) ParseUserDeposited(log types.Log) (*SavingUserDeposited, error) {
	event := new(SavingUserDeposited)
	if err := _Saving.contract.UnpackLog(event, "UserDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SavingUserWithdrawnIterator is returned from FilterUserWithdrawn and is used to iterate over the raw logs and unpacked data for UserWithdrawn events raised by the Saving contract.
type SavingUserWithdrawnIterator struct {
	Event *SavingUserWithdrawn // Event containing the contract specifics and raw log

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
func (it *SavingUserWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SavingUserWithdrawn)
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
		it.Event = new(SavingUserWithdrawn)
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
func (it *SavingUserWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SavingUserWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SavingUserWithdrawn represents a UserWithdrawn event raised by the Saving contract.
type SavingUserWithdrawn struct {
	User           common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserWithdrawn is a free log retrieval operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawAmount)
func (_Saving *SavingFilterer) FilterUserWithdrawn(opts *bind.FilterOpts, _user []common.Address) (*SavingUserWithdrawnIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Saving.contract.FilterLogs(opts, "UserWithdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return &SavingUserWithdrawnIterator{contract: _Saving.contract, event: "UserWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserWithdrawn is a free log subscription operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawAmount)
func (_Saving *SavingFilterer) WatchUserWithdrawn(opts *bind.WatchOpts, sink chan<- *SavingUserWithdrawn, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Saving.contract.WatchLogs(opts, "UserWithdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SavingUserWithdrawn)
				if err := _Saving.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
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

// ParseUserWithdrawn is a log parse operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed _user, uint256 _withdrawAmount)
func (_Saving *SavingFilterer) ParseUserWithdrawn(log types.Log) (*SavingUserWithdrawn, error) {
	event := new(SavingUserWithdrawn)
	if err := _Saving.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
