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
	Bin: "0x608060405234801561001057600080fd5b506123f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100835760003560e01c806312065fe0146100885780632e1a7d4d146100a65780633d4d5d4d146100c657806364eb4369146100db5780636fa07824146100f0578063b6b55f25146100f8578063c0ab9cbc1461010b578063cb5e79e014610113578063f399e22e14610126575b600080fd5b61009061013b565b60405161009d919061234d565b60405180910390f35b6100b96100b436600461209c565b6101cf565b60405161009d919061211d565b6100ce6104c5565b60405161009d91906120cc565b6100e36104d4565b60405161009d9190612313565b6100906107ac565b6100b961010636600461209c565b6107b2565b610090610a56565b610090610121366004611f53565b61126d565b610139610134366004611f8b565b61127f565b005b600254336000908152600460208190526040808320549051631d335d6560e21b815292936001600160a01b0316926374cd75949261017a92910161234d565b60206040518083038186803b15801561019257600080fd5b505afa1580156101a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ca91906120b4565b905090565b6000806101da61013b565b9050806102025760405162461bcd60e51b81526004016101f990612238565b60405180910390fd5b828082101561020e5750805b6002546000906001600160a01b0316638dc3311d61022c858561146e565b6040518263ffffffff1660e01b8152600401610248919061234d565b60206040518083038186803b15801561026057600080fd5b505afa158015610274573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029891906120b4565b336000908152600460205260409020549091506102c2906102b9908361146e565b6003549061146e565b6003553360009081526004602090815260408083208490559154600180548451601f600260001984861615610100020190931692909204918201859004850281018501909552808552620100009092046001600160a01b031693633fb9027193610385939192919083018282801561037b5780601f106103505761010080835404028352916020019161037b565b820191906000526020600020905b81548152906001019060200180831161035e57829003601f168201915b50505050506114b7565b6040518263ffffffff1660e01b81526004016103a19190612128565b60206040518083038186803b1580156103b957600080fd5b505afa1580156103cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f19190611f6f565b6001600160a01b031663a9059cbb33846040518363ffffffff1660e01b815260040161041e929190612104565b602060405180830381600087803b15801561043857600080fd5b505af115801561044c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104709190612027565b50336001600160a01b03167fe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb836040516104aa919061234d565b60405180910390a26104ba611660565b506001949350505050565b6002546001600160a01b031681565b6104dc611e83565b6104e4611e83565b60025460408051630f7fb07b60e41b815290516001600160a01b0390921691829163f7fb07b0916004808301926020929190829003018186803b15801561052a57600080fd5b505afa15801561053e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056291906120b4565b826040018181525050806001600160a01b0316634c89867f6040518163ffffffff1660e01b815260040160206040518083038186803b1580156105a457600080fd5b505afa1580156105b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105dc91906120b4565b60608301526105e961013b565b8252336000908152600460208181526040808420549186019190915291549151633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161063491016121f6565b60206040518083038186803b15801561064c57600080fd5b505afa158015610660573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106849190611f6f565b6001600160a01b031663498bff0061073560018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561072b5780601f106107005761010080835404028352916020019161072b565b820191906000526020600020905b81548152906001019060200180831161070e57829003601f168201915b505050505061185b565b6040518263ffffffff1660e01b81526004016107519190612128565b60206040518083038186803b15801561076957600080fd5b505afa15801561077d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a191906120b4565b608083015250905090565b60035481565b60008082116107d35760405162461bcd60e51b81526004016101f990612295565b60005460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361084d939192909183018282801561037b5780601f106103505761010080835404028352916020019161037b565b6040518263ffffffff1660e01b81526004016108699190612128565b60206040518083038186803b15801561088157600080fd5b505afa158015610895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b99190611f6f565b6001600160a01b03166323b872dd3330856040518463ffffffff1660e01b81526004016108e8939291906120e0565b602060405180830381600087803b15801561090257600080fd5b505af1158015610916573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093a9190612027565b506002546000906001600160a01b0316638dc3311d6109618561095b61013b565b9061188f565b6040518263ffffffff1660e01b815260040161097d919061234d565b60206040518083038186803b15801561099557600080fd5b505afa1580156109a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cd91906120b4565b336000908152600460205260409020549091506109f8906109ef90839061146e565b6003549061188f565b6003553360008181526004602052604090819020839055517f951fdc61d6a98f96098a17ea6ac287a6fd38aea6bef73083c93b274cb830107d90610a3d90869061234d565b60405180910390a2610a4d611660565b50600192915050565b6002805460018054604080516020601f600019858716156101000201909416969096049283018690048602810186019091528181526000946001600160a01b03909416938593919290830182828015610af05780601f10610ac557610100808354040283529160200191610af0565b820191906000526020600020905b815481529060010190602001808311610ad357829003601f168201915b505060008054604051633fb9027160e01b81529596509094620100009091046001600160a01b03169350633fb902719250610b2e91506004016121f6565b60206040518083038186803b158015610b4657600080fd5b505afa158015610b5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7e9190611f6f565b6001600160a01b031663498bff00610b958461185b565b6040518263ffffffff1660e01b8152600401610bb19190612128565b60206040518083038186803b158015610bc957600080fd5b505afa158015610bdd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0191906120b4565b90506000836001600160a01b03166374cd75946003546040518263ffffffff1660e01b8152600401610c33919061234d565b60206040518083038186803b158015610c4b57600080fd5b505afa158015610c5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8391906120b4565b90506000846001600160a01b03166382ab890a846040518263ffffffff1660e01b8152600401610cb3919061234d565b602060405180830381600087803b158015610ccd57600080fd5b505af1158015610ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0591906120b4565b90506000856001600160a01b03166374cd75946003546040518263ffffffff1660e01b8152600401610d37919061234d565b60206040518083038186803b158015610d4f57600080fd5b505afa158015610d63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8791906120b4565b60008054919250906201000090046001600160a01b0316633fb90271610dac886118e7565b6040518263ffffffff1660e01b8152600401610dc89190612128565b60206040518083038186803b158015610de057600080fd5b505afa158015610df4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e189190611f6f565b90506000816001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b158015610e5557600080fd5b505afa158015610e69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8d91906120b4565b90506000610e9b848761146e565b9050808210156111d857600080546201000090046001600160a01b0316633fb90271610ec68b611a99565b6040518263ffffffff1660e01b8152600401610ee29190612128565b60206040518083038186803b158015610efa57600080fd5b505afa158015610f0e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f329190611f6f565b6001600160a01b0316635ac2b2ee6040518163ffffffff1660e01b815260040160606040518083038186803b158015610f6a57600080fd5b505afa158015610f7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa29190612047565b90506000610fb0838561146e565b9050816040015161102c82876001600160a01b03166314a6bf0f6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ff457600080fd5b505afa158015611008573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095b91906120b4565b111561104a5760405162461bcd60e51b81526004016101f99061217b565b6000546201000090046001600160a01b0316633fb9027161106a8c6114b7565b6040518263ffffffff1660e01b81526004016110869190612128565b60206040518083038186803b15801561109e57600080fd5b505afa1580156110b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d69190611f6f565b6001600160a01b03166340c10f1986836040518363ffffffff1660e01b8152600401611103929190612104565b602060405180830381600087803b15801561111d57600080fd5b505af1158015611131573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111559190612027565b50604051632b7c7b1160e01b81526001600160a01b03861690632b7c7b119061118290849060040161234d565b602060405180830381600087803b15801561119c57600080fd5b505af11580156111b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d49190612027565b5050505b60405163effd85ad60e01b81526001600160a01b0384169063effd85ad9061120490849060040161234d565b602060405180830381600087803b15801561121e57600080fd5b505af1158015611232573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112569190612027565b5061125f611660565b509297505050505050505090565b60046020526000908152604090205481565b600054610100900460ff16806112985750611298611c47565b806112a6575060005460ff16155b6112e15760405162461bcd60e51b815260040180806020018281038252602e815260200180612393602e913960400191505060405180910390fd5b600054610100900460ff1615801561130c576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151611341906001906020850190611eb2565b50600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190611375906004016122e2565b60206040518083038186803b15801561138d57600080fd5b505afa1580156113a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c59190611f6f565b6001600160a01b031663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156113ff57600080fd5b505af1158015611413573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114379190611f6f565b600280546001600160a01b0319166001600160a01b03929092169190911790558015611469576000805461ff00191690555b505050565b60006114b083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611c58565b9392505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b6020831061153d5780518252601f19909201916020918201910161151e565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106115855780518252601f199092019160209182019101611566565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106115cd5780518252601f1990920191602091820191016115ae565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106116155780518252601f1990920191602091820191016115f6565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b600254600354604051631d335d6560e21b81526000926001600160a01b0316916374cd759491611693919060040161234d565b60206040518083038186803b1580156116ab57600080fd5b505afa1580156116bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116e391906120b4565b6000805460018054604080516020600261010085871615026000190190941693909304601f81018490048402820184019092528181529596509394620100009093046001600160a01b031693633fb902719361176493909290919083018282801561037b5780601f106103505761010080835404028352916020019161037b565b6040518263ffffffff1660e01b81526004016117809190612128565b60206040518083038186803b15801561179857600080fd5b505afa1580156117ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d09190611f6f565b6001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016117fb91906120cc565b60206040518083038186803b15801561181357600080fd5b505afa158015611827573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184b91906120b4565b90508181101561185757fe5b5050565b6060611889826040518060400160405280600a815260200169736176696e675261746560b01b815250611cef565b92915050565b6000828201838110156114b0576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b6020831061196d5780518252601f19909201916020918201910161194e565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106119b55780518252601f199092019160209182019101611996565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106119fd5780518252601f1990920191602091820191016119de565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611a455780518252601f199092019160209182019101611a26565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b60208310611b1f5780518252601f199092019160209182019101611b00565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310611b675780518252601f199092019160209182019101611b48565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611baf5780518252601f199092019160209182019101611b90565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611bf75780518252601f199092019160209182019101611bd8565b5181516020939093036101000a600019018019909116921691909117905268626f72726f77696e6760b81b9201918252506040805180830360161901815260099092019052979650505050505050565b6000611c5230611e7d565b15905090565b60008184841115611ce75760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611cac578181015183820152602001611c94565b50505050905090810190601f168015611cd95780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b60208310611d665780518252601f199092019160209182019101611d47565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310611dae5780518252601f199092019160209182019101611d8f565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611df65780518252601f199092019160209182019101611dd7565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611e3e5780518252601f199092019160209182019101611e1f565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b3b151590565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611ee85760008555611f2e565b82601f10611f0157805160ff1916838001178555611f2e565b82800160010185558215611f2e579182015b82811115611f2e578251825591602001919060010190611f13565b50611f3a929150611f3e565b5090565b5b80821115611f3a5760008155600101611f3f565b600060208284031215611f64578081fd5b81356114b08161237a565b600060208284031215611f80578081fd5b81516114b08161237a565b60008060408385031215611f9d578081fd5b8235611fa88161237a565b915060208381013567ffffffffffffffff80821115611fc5578384fd5b818601915086601f830112611fd8578384fd5b813581811115611fe457fe5b611ff6601f8201601f19168501612356565b9150808252878482850101111561200b578485fd5b8084840185840137810190920192909252919491935090915050565b600060208284031215612038578081fd5b815180151581146114b0578182fd5b600060608284031215612058578081fd5b6040516060810181811067ffffffffffffffff8211171561207557fe5b80604052508251815260208301516020820152604083015160408201528091505092915050565b6000602082840312156120ad578081fd5b5035919050565b6000602082840312156120c5578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b6000602080835283518082850152825b8181101561215457858101830151858201604001528201612138565b818111156121655783604083870101525b50601f01601f1916929092016040019392505050565b60208082526055908201527f5b5145432d3032323030325d2d53797374656d2064656274206578636565647360408201527f206f77656420626f72726f77696e6720666565732c206661696c656420746f206060820152743ab83230ba329031b7b6b837bab732103930ba329760591b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b6020808252603e908201527f5b5145432d3032323030315d2d5468652063616c6c657220646f6573206e6f7460408201527f206861766520616e792062616c616e636520746f2077697468647261772e0000606082015260800190565b6020808252602d908201527f5b5145432d3032323030305d2d4465706f73697420616d6f756e74206d75737460408201526c103737ba103132903d32b9379760991b606082015260800190565b60208082526017908201527631b7b6b6b7b7173330b1ba37b93c9731b925b2b2b832b960491b604082015260600190565b600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b90815260200190565b60405181810167ffffffffffffffff8111828210171561237257fe5b604052919050565b6001600160a01b038116811461238f57600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564a2646970667358221220e35b963539755d648fb4d35a71f47d993445ee9e34e5531c770e336bae9c094464736f6c63430007060033",
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
