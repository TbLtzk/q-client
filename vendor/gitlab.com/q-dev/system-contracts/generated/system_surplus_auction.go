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

// SystemSurplusAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemSurplusAuctionAuctionInfo struct {
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
	IsExecuted bool
}

// SystemSurplusAuctionABI is the input ABI used to generate the binding from.
const SystemSurplusAuctionABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structSystemSurplusAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemSurplusAuctionBin is the compiled bytecode used for deploying new contracts.
var SystemSurplusAuctionBin = "0x608060405234801561001057600080fd5b506120b5806100206000396000f3fe6080604052600436106100605760003560e01c8063454a2ab314610065578063571a26a01461008e5780636b64c769146100bf578063a2165920146100d4578063b830620a146100f4578063f399e22e14610109578063fe0d94c11461012b575b600080fd5b610078610073366004611996565b61014b565b6040516100859190611a26565b60405180910390f35b34801561009a57600080fd5b506100ae6100a9366004611996565b6103a5565b6040516100859594939291906119f6565b6100c76103e2565b6040516100859190611fbc565b3480156100e057600080fd5b506100c76100ef366004611996565b610ae4565b34801561010057600080fd5b506100c7610b18565b34801561011557600080fd5b506101296101243660046118cc565b610b1e565b005b34801561013757600080fd5b50610078610146366004611996565b610bf8565b60008160025481106101785760405162461bcd60e51b815260040161016f90611cb2565b60405180910390fd5b60008381526003602081905260409091209081015442106101ab5760405162461bcd60e51b815260040161016f90611f68565b6101b484610f84565b3410156101d35760405162461bcd60e51b815260040161016f90611a84565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061020790600401611ec6565b60206040518083038186803b15801561021f57600080fd5b505afa158015610233573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025791906118b0565b60018301548354604051632377789f60e21b81529293506001600160a01b0380851693638ddde27c939261028f9216906004016119c9565b6020604051808303818588803b1580156102a857600080fd5b505af11580156102bc573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906102e19190611976565b508154336001600160a01b0319918216811780855534600180870182815560008b81526003602081905260409182902080546001600160a01b03909716969098169590951787559054918601919091556002808801549086015582870154928501929092556004808701549401805460ff191660ff90951615159490941790935551909187917fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca31999161039291611fbc565b60405180910390a3506001949350505050565b6003602081905260009182526040909120805460018201546002830154938301546004909301546001600160a01b03909216939092909160ff1685565b6000346104015760405162461bcd60e51b815260040161016f90611b14565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb90271936104b193919290918301828280156104a75780601f1061047c576101008083540402835291602001916104a7565b820191906000526020600020905b81548152906001019060200180831161048a57829003601f168201915b50505050506110dc565b6040518263ffffffff1660e01b81526004016104cd9190611a31565b60206040518083038186803b1580156104e557600080fd5b505afa1580156104f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051d91906118b0565b9050806001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561055a57600080fd5b505af115801561056e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105929190611976565b6105ae5760405162461bcd60e51b815260040161016f90611e58565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb90271906105e290600401611cfa565b60206040518083038186803b1580156105fa57600080fd5b505afa15801561060e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063291906118b0565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0386169363498bff00936106db938301828280156106d15780601f106106a6576101008083540402835291602001916106d1565b820191906000526020600020905b8154815290600101906020018083116106b457829003601f168201915b505050505061128e565b6040518263ffffffff1660e01b81526004016106f79190611a31565b60206040518083038186803b15801561070f57600080fd5b505afa158015610723573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074791906119ae565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0387169363498bff00936107f0938301828280156107e65780601f106107bb576101008083540402835291602001916107e6565b820191906000526020600020905b8154815290600101906020018083116107c957829003601f168201915b50505050506112c2565b6040518263ffffffff1660e01b815260040161080c9190611a31565b60206040518083038186803b15801561082457600080fd5b505afa158015610838573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085c91906119ae565b90506000846001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b15801561089957600080fd5b505afa1580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d191906119ae565b9050818110156108f35760405162461bcd60e51b815260040161016f90611d7d565b828110156109135760405162461bcd60e51b815260040161016f90611b6e565b604051631dc7b5db60e31b81526001600160a01b0386169063ee3daed89061093f908690600401611fbc565b602060405180830381600087803b15801561095957600080fd5b505af115801561096d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109919190611976565b5060028054600181019091556109a56117d5565b33815234602082015260408082018690525162498bff60e81b8152610a38906001600160a01b0388169063498bff00906109e190600401611c7b565b60206040518083038186803b1580156109f957600080fd5b505afa158015610a0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3191906119ae565b42906112f6565b60608201908152600083815260036020818152604092839020855181546001600160a01b0319166001600160a01b0390911617815590850151600182015582850151600282015592519083015560808301516004909201805460ff1916921515929092179091555133907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610ad1908590349061200a565b60405180910390a2509550505050505090565b6000816002548110610b085760405162461bcd60e51b815260040161016f90611cb2565b610b1183610f84565b9392505050565b60025481565b600054610100900460ff1680610b375750610b3761134e565b80610b45575060005460ff16155b610b805760405162461bcd60e51b815260040180806020018281038252602e815260200180612031602e913960400191505060405180910390fd5b600054610100900460ff16158015610bab576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610be090600190602085019061180f565b508015610bf3576000805461ff00191690555b505050565b6000816002548110610c1c5760405162461bcd60e51b815260040161016f90611cb2565b600083815260036020819052604090912090810154421015610c505760405162461bcd60e51b815260040161016f90611de9565b600481015460ff1615610c755760405162461bcd60e51b815260040161016f90611c0b565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610d259391929091830182828015610d1b5780601f10610cf057610100808354040283529160200191610d1b565b820191906000526020600020905b815481529060010190602001808311610cfe57829003601f168201915b5050505050611354565b6040518263ffffffff1660e01b8152600401610d419190611a31565b60206040518083038186803b158015610d5957600080fd5b505afa158015610d6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9191906118b0565b8254600284015460405163a9059cbb60e01b81529293506001600160a01b038085169363a9059cbb93610dca93921691906004016119dd565b602060405180830381600087803b158015610de457600080fd5b505af1158015610df8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1c9190611976565b5060048281018054600160ff1990911617905560008054604051633fb9027160e01b81529192620100009091046001600160a01b031691633fb9027191610e639101611bc6565b60206040518083038186803b158015610e7b57600080fd5b505afa158015610e8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb391906118b0565b6001600160a01b03168360010154604051610ecd906119c6565b60006040518083038185875af1925050503d8060008114610f0a576040519150601f19603f3d011682016040523d82523d6000602084013e610f0f565b606091505b5050905080610f305760405162461bcd60e51b815260040161016f90611efb565b82546040516001600160a01b03909116907f082679bed23190a10049632ab265cbdc8daffaf8843809bb9f893aded2f4aa2090610f709089908790611fc5565b60405180910390a250600195945050505050565b60008054604051633fb9027160e01b815282916201000090046001600160a01b031690633fb9027190610fb990600401611cfa565b60206040518083038186803b158015610fd157600080fd5b505afa158015610fe5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100991906118b0565b6001600160a01b031663498bff006040518163ffffffff1660e01b815260040161103290611d3c565b60206040518083038186803b15801561104a57600080fd5b505afa15801561105e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108291906119ae565b6000848152600360205260408120600101549192506110bc826110b66110a66114fd565b6110b0868861150d565b90611566565b906112f6565b9050818114156110d4576110d18160016112f6565b90505b949350505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106111625780518252601f199092019160209182019101611143565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106111aa5780518252601f19909201916020918201910161118b565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106111f25780518252601f1990920191602091820191016111d3565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061123a5780518252601f19909201916020918201910161121b565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b60606112bc826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b8152506115a5565b92915050565b60606112bc826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b8152506115a5565b600082820183811015610b11576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b303b1590565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106113da5780518252601f1990920191602091820191016113bb565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106114225780518252601f199092019160209182019101611403565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061146a5780518252601f19909201916020918201910161144b565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114b25780518252601f199092019160209182019101611493565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b6b033b2e3c9fd0803ce800000090565b60008261151c575060006112bc565b8282028284828161152957fe5b0414610b115760405162461bcd60e51b815260040180806020018281038252602181526020018061205f6021913960400191505060405180910390fd5b6000610b1183836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250611733565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b6020831061161c5780518252601f1990920191602091820191016115fd565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106116645780518252601f199092019160209182019101611645565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106116ac5780518252601f19909201916020918201910161168d565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106116f45780518252601f1990920191602091820191016116d5565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b600081836117bf5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561178457818101518382015260200161176c565b50505050905090810190601f1680156117b15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816117cb57fe5b0495945050505050565b6040518060a0016040528060006001600160a01b031681526020016000815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611845576000855561188b565b82601f1061185e57805160ff191683800117855561188b565b8280016001018555821561188b579182015b8281111561188b578251825591602001919060010190611870565b5061189792915061189b565b5090565b5b80821115611897576000815560010161189c565b6000602082840312156118c1578081fd5b8151610b1181612018565b600080604083850312156118de578081fd5b82356118e981612018565b915060208381013567ffffffffffffffff80821115611906578384fd5b818601915086601f830112611919578384fd5b81358181111561192557fe5b604051601f8201601f191681018501838111828210171561194257fe5b6040528181528382018501891015611958578586fd5b81858501868301378585838301015280955050505050509250929050565b600060208284031215611987578081fd5b81518015158114610b11578182fd5b6000602082840312156119a7578081fd5b5035919050565b6000602082840312156119bf578081fd5b5051919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039590951685526020850193909352604084019190915260608301521515608082015260a00190565b901515815260200190565b6000602080835283518082850152825b81811015611a5d57858101830151858201604001528201611a41565b81811115611a6e5783604083870101525b50601f01601f1916929092016040019392505050565b60208082526064908201527f5b5145432d3032353030365d2d5468652062696420616d6f756e74206d75737460408201527f206578636565642074686520686967686573742062696420627920746865206d60608201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60808201526337b9329760e11b60a082015260c00190565b6020808252603a908201527f5b5145432d3032353030305d2d496e76616c69642062696420616d6f756e742c604082015279103330b4b632b21039ba30b93a103a34329030bab1ba34b7b71760311b606082015260800190565b60208082526038908201527f5b5145432d3032353030335d2d4e6f7420656e6f75676820737572706c7573206040820152773a37903334b636103a34329030bab1ba34b7b7103637ba1760411b606082015260800190565b60208082526025908201527f746f6b656e65636f6e6f6d6963732e64656661756c74416c6c6f636174696f6e60408201526450726f787960d81b606082015260800190565b6020808252604a908201527f5b5145432d3032353030395d2d5468652061756374696f6e2068617320616c7260408201527f656164792065786563757465642c206661696c656420746f207472616e73666560608201526939103a3432903637ba1760b11b608082015260a00190565b6020808252601d908201527f676f7665726e65642e455044522e737572706c757341756374696f6e50000000604082015260600190565b60208082526028908201527f5b5145432d3032353030345d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b60208082526021908201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6040820152601d60fa1b606082015260800190565b60208082526046908201527f5b5145432d3032353030325d2d537572706c7573206166746572206e6574746960408201527f6e672069732062656c6f7720737572706c75732061756374696f6e207468726560608201526539b437b6321760d11b608082015260a00190565b60208082526049908201527f5b5145432d3032353030385d2d5468652061756374696f6e206973206e6f742060408201527f66696e6973686564207965742c206661696c656420746f207472616e73666572606082015268103a3432903637ba1760b91b608082015260a00190565b60208082526048908201527f5b5145432d3032353030315d2d546865206e657474696e6720776173206e6f7460408201527f20706572666f726d65642c206661696c656420746f207374617274207468652060608201526730bab1ba34b7b71760c11b608082015260a00190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b60208082526047908201527f5145432d3032353031302d4661696c656420746f207472616e7366657220746860408201527f6520686967686573742062696420746f2074686564656661756c7420616c6c6f606082015266632070726f787960c81b608082015260a00190565b60208082526034908201527f5b5145432d3032353030355d2d5468652061756374696f6e2069732066696e6960408201527339b432b216103330b4b632b2103a37903134b21760611b606082015260800190565b90815260200190565b91825280546001600160a01b031660208301526001810154604083015260028101546060830152600381015460808301526004015460ff16151560a082015260c00190565b918252602082015260400190565b6001600160a01b038116811461202d57600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212202c2f2019ae4871981f6fb46b3d9793297ad880143e6e3a24bd1f0a8ffef382cd64736f6c63430007060033"

// DeploySystemSurplusAuction deploys a new Ethereum contract, binding an instance of SystemSurplusAuction to it.
func DeploySystemSurplusAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemSurplusAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemSurplusAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemSurplusAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// SystemSurplusAuction is an auto generated Go binding around an Ethereum contract.
type SystemSurplusAuction struct {
	SystemSurplusAuctionCaller     // Read-only binding to the contract
	SystemSurplusAuctionTransactor // Write-only binding to the contract
	SystemSurplusAuctionFilterer   // Log filterer for contract events
}

// SystemSurplusAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemSurplusAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemSurplusAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemSurplusAuctionSession struct {
	Contract     *SystemSurplusAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemSurplusAuctionCallerSession struct {
	Contract *SystemSurplusAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SystemSurplusAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemSurplusAuctionTransactorSession struct {
	Contract     *SystemSurplusAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SystemSurplusAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemSurplusAuctionRaw struct {
	Contract *SystemSurplusAuction // Generic contract binding to access the raw methods on
}

// SystemSurplusAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemSurplusAuctionCallerRaw struct {
	Contract *SystemSurplusAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SystemSurplusAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemSurplusAuctionTransactorRaw struct {
	Contract *SystemSurplusAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemSurplusAuction creates a new instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuction(address common.Address, backend bind.ContractBackend) (*SystemSurplusAuction, error) {
	contract, err := bindSystemSurplusAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuction{SystemSurplusAuctionCaller: SystemSurplusAuctionCaller{contract: contract}, SystemSurplusAuctionTransactor: SystemSurplusAuctionTransactor{contract: contract}, SystemSurplusAuctionFilterer: SystemSurplusAuctionFilterer{contract: contract}}, nil
}

// NewSystemSurplusAuctionCaller creates a new read-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionCaller(address common.Address, caller bind.ContractCaller) (*SystemSurplusAuctionCaller, error) {
	contract, err := bindSystemSurplusAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionCaller{contract: contract}, nil
}

// NewSystemSurplusAuctionTransactor creates a new write-only instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemSurplusAuctionTransactor, error) {
	contract, err := bindSystemSurplusAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionTransactor{contract: contract}, nil
}

// NewSystemSurplusAuctionFilterer creates a new log filterer instance of SystemSurplusAuction, bound to a specific deployed contract.
func NewSystemSurplusAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemSurplusAuctionFilterer, error) {
	contract, err := bindSystemSurplusAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionFilterer{contract: contract}, nil
}

// bindSystemSurplusAuction binds a generic wrapper to an already deployed contract.
func bindSystemSurplusAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemSurplusAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.SystemSurplusAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemSurplusAuction *SystemSurplusAuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SystemSurplusAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address bidder, uint256 highestBid, uint256 lot, uint256 endTime, bool isExecuted)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
	IsExecuted bool
}, error) {
	ret := new(struct {
		Bidder     common.Address
		HighestBid *big.Int
		Lot        *big.Int
		EndTime    *big.Int
		IsExecuted bool
	})
	out := ret
	err := _SystemSurplusAuction.contract.Call(opts, out, "auctions", arg0)
	return *ret, err
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address bidder, uint256 highestBid, uint256 lot, uint256 endTime, bool isExecuted)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Auctions(arg0 *big.Int) (struct {
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
	IsExecuted bool
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address bidder, uint256 highestBid, uint256 lot, uint256 endTime, bool isExecuted)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
	IsExecuted bool
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) AuctionsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemSurplusAuction.contract.Call(opts, out, "auctionsCount")
	return *ret0, err
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) AuctionsCount() (*big.Int, error) {
	return _SystemSurplusAuction.Contract.AuctionsCount(&_SystemSurplusAuction.CallOpts)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SystemSurplusAuction.contract.Call(opts, out, "getRaisingBid", _auctionId)
	return *ret0, err
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemSurplusAuction.Contract.GetRaisingBid(&_SystemSurplusAuction.CallOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "bid", _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _auctionId) payable returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Bid(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Bid(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Execute(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "execute", _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _auctionId) returns(bool)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Execute(_auctionId *big.Int) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Execute(&_SystemSurplusAuction.TransactOpts, _auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.Initialize(&_SystemSurplusAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemSurplusAuction.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() payable returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionTransactorSession) StartAuction() (*types.Transaction, error) {
	return _SystemSurplusAuction.Contract.StartAuction(&_SystemSurplusAuction.TransactOpts)
}

// SystemSurplusAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStartedIterator struct {
	Event *SystemSurplusAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionAuctionStarted)
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
		it.Event = new(SystemSurplusAuctionAuctionStarted)
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
func (it *SystemSurplusAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionAuctionStarted represents a AuctionStarted event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionAuctionStarted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionAuctionStartedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionAuctionStartedIterator{contract: _SystemSurplusAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionAuctionStarted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "AuctionStarted", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionAuctionStarted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseAuctionStarted(log types.Log) (*SystemSurplusAuctionAuctionStarted, error) {
	event := new(SystemSurplusAuctionAuctionStarted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemSurplusAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBidIterator struct {
	Event *SystemSurplusAuctionBid // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionBid)
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
		it.Event = new(SystemSurplusAuctionBid)
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
func (it *SystemSurplusAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionBid represents a Bid event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterBid(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemSurplusAuctionBidIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionBidIterator{contract: _SystemSurplusAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionBid, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionBid)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseBid(log types.Log) (*SystemSurplusAuctionBid, error) {
	event := new(SystemSurplusAuctionBid)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SystemSurplusAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecutedIterator struct {
	Event *SystemSurplusAuctionExecuted // Event containing the contract specifics and raw log

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
func (it *SystemSurplusAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemSurplusAuctionExecuted)
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
		it.Event = new(SystemSurplusAuctionExecuted)
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
func (it *SystemSurplusAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemSurplusAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemSurplusAuctionExecuted represents a Executed event raised by the SystemSurplusAuction contract.
type SystemSurplusAuctionExecuted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Info      SystemSurplusAuctionAuctionInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x082679bed23190a10049632ab265cbdc8daffaf8843809bb9f893aded2f4aa20.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (address,uint256,uint256,uint256,bool) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _bidder []common.Address) (*SystemSurplusAuctionExecutedIterator, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.FilterLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemSurplusAuctionExecutedIterator{contract: _SystemSurplusAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x082679bed23190a10049632ab265cbdc8daffaf8843809bb9f893aded2f4aa20.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (address,uint256,uint256,uint256,bool) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SystemSurplusAuctionExecuted, _bidder []common.Address) (event.Subscription, error) {

	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemSurplusAuction.contract.WatchLogs(opts, "Executed", _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemSurplusAuctionExecuted)
				if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x082679bed23190a10049632ab265cbdc8daffaf8843809bb9f893aded2f4aa20.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (address,uint256,uint256,uint256,bool) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseExecuted(log types.Log) (*SystemSurplusAuctionExecuted, error) {
	event := new(SystemSurplusAuctionExecuted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	return event, nil
}
