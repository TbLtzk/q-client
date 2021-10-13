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

// SystemSurplusAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemSurplusAuctionAuctionInfo struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}

// SystemSurplusAuctionMetaData contains all meta data concerning the SystemSurplusAuction contract.
var SystemSurplusAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSystemSurplusAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611fb0806100206000396000f3fe6080604052600436106100605760003560e01c8063454a2ab314610065578063571a26a01461008e5780636b64c769146100bf578063a2165920146100d4578063b830620a146100f4578063f399e22e14610109578063fe0d94c11461012b575b600080fd5b610078610073366004611943565b61014b565b60405161008591906119a3565b60405180910390f35b34801561009a57600080fd5b506100ae6100a9366004611943565b610356565b6040516100859594939291906119ae565b6100c7610391565b6040516100859190611eb8565b3480156100e057600080fd5b506100c76100ef366004611943565b610a7e565b34801561010057600080fd5b506100c7610ab2565b34801561011557600080fd5b50610129610124366004611879565b610ab8565b005b34801561013757600080fd5b50610078610146366004611943565b610b92565b60008160025481106101785760405162461bcd60e51b815260040161016f90611bed565b60405180910390fd5b6000838152600360208190526040909120908101544211156101ac5760405162461bcd60e51b815260040161016f90611da2565b6101b584610f20565b3410156101d45760405162461bcd60e51b815260040161016f90611a2f565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061020890600401611d6d565b60206040518083038186803b15801561022057600080fd5b505afa158015610234573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610258919061185d565b600183018054845433610100908102610100600160a81b0319831617875534909355604051632377789f60e21b815293945090926001600160a01b0392909104821691841690638ddde27c9084906102b4908590600401611976565b6020604051808303818588803b1580156102cd57600080fd5b505af11580156102e1573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906103069190611923565b50336001600160a01b0316877fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199346040516103419190611eb8565b60405180910390a35060019695505050505050565b6003602081905260009182526040909120805460018201546002830154929093015460ff8216936101009092046001600160a01b0316929085565b6000346103b05760405162461bcd60e51b815260040161016f90611abf565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361046093919290918301828280156104565780601f1061042b57610100808354040283529160200191610456565b820191906000526020600020905b81548152906001019060200180831161043957829003601f168201915b5050505050611078565b6040518263ffffffff1660e01b815260040161047c91906119dc565b60206040518083038186803b15801561049457600080fd5b505afa1580156104a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cc919061185d565b9050806001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561050957600080fd5b505af115801561051d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105419190611923565b5060008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061057690600401611c35565b60206040518083038186803b15801561058e57600080fd5b505afa1580156105a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c6919061185d565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0386169363498bff009361066f938301828280156106655780601f1061063a57610100808354040283529160200191610665565b820191906000526020600020905b81548152906001019060200180831161064857829003601f168201915b505050505061122a565b6040518263ffffffff1660e01b815260040161068b91906119dc565b60206040518083038186803b1580156106a357600080fd5b505afa1580156106b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106db919061195b565b60018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529394506000936001600160a01b0387169363498bff00936107849383018282801561077a5780601f1061074f5761010080835404028352916020019161077a565b820191906000526020600020905b81548152906001019060200180831161075d57829003601f168201915b505050505061125e565b6040518263ffffffff1660e01b81526004016107a091906119dc565b60206040518083038186803b1580156107b857600080fd5b505afa1580156107cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107f0919061195b565b90506000846001600160a01b0316632383b0746040518163ffffffff1660e01b815260040160206040518083038186803b15801561082d57600080fd5b505afa158015610841573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610865919061195b565b9050818110156108875760405162461bcd60e51b815260040161016f90611cb8565b828110156108a75760405162461bcd60e51b815260040161016f90611b19565b60405163152353d960e01b81526001600160a01b0386169063152353d9906108d3908690600401611eb8565b602060405180830381600087803b1580156108ed57600080fd5b505af1158015610901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109259190611923565b506002805460018101909155610939611782565b33602082015234604080830191909152606082018690525162498bff60e81b81526109d2906001600160a01b0388169063498bff009061097b90600401611bb6565b60206040518083038186803b15801561099357600080fd5b505afa1580156109a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cb919061195b565b4290611292565b6080820190815260008381526003602081815260409283902085518154928701516001600160a01b031661010002610100600160a81b031991151560ff1990941693909317169190911781558285015160018201556060850151600282015592519201919091555133907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610a6b9085903490611f05565b60405180910390a2509550505050505090565b6000816002548110610aa25760405162461bcd60e51b815260040161016f90611bed565b610aab83610f20565b9392505050565b60025481565b600054610100900460ff1680610ad15750610ad16112ea565b80610adf575060005460ff16155b610b1a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611f2c602e913960400191505060405180910390fd5b600054610100900460ff16158015610b45576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151610b7a9060019060208501906117bc565b508015610b8d576000805461ff00191690555b505050565b6000816002548110610bb65760405162461bcd60e51b815260040161016f90611bed565b6000838152600360208190526040909120908101544211610be95760405162461bcd60e51b815260040161016f90611d24565b805460ff1615610c0b5760405162461bcd60e51b815260040161016f90611df6565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb9027193610cbb9391929091830182828015610cb15780601f10610c8657610100808354040283529160200191610cb1565b820191906000526020600020905b815481529060010190602001808311610c9457829003601f168201915b50505050506112fb565b6040518263ffffffff1660e01b8152600401610cd791906119dc565b60206040518083038186803b158015610cef57600080fd5b505afa158015610d03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d27919061185d565b8254600284015460405163a9059cbb60e01b81529293506001600160a01b038085169363a9059cbb93610d6793610100909104909216919060040161198a565b602060405180830381600087803b158015610d8157600080fd5b505af1158015610d95573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610db99190611923565b508154600160ff1990911617825560008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610dfb90600401611b71565b60206040518083038186803b158015610e1357600080fd5b505afa158015610e27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4b919061185d565b6001600160a01b03168360010154604051610e6590611973565b60006040518083038185875af1925050503d8060008114610ea2576040519150601f19603f3d011682016040523d82523d6000602084013e610ea7565b606091505b5050905080610ec85760405162461bcd60e51b815260040161016f90611e49565b82546040516101009091046001600160a01b0316907f2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b90610f0c9089908790611ec1565b60405180910390a250600195945050505050565b60008054604051633fb9027160e01b815282916201000090046001600160a01b031690633fb9027190610f5590600401611c35565b60206040518083038186803b158015610f6d57600080fd5b505afa158015610f81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa5919061185d565b6001600160a01b031663498bff006040518163ffffffff1660e01b8152600401610fce90611c77565b60206040518083038186803b158015610fe657600080fd5b505afa158015610ffa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101e919061195b565b600084815260036020526040812060010154919250611058826110526110426114a4565b61104c86886114b4565b9061150d565b90611292565b9050818114156110705761106d816001611292565b90505b949350505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106110fe5780518252601f1990920191602091820191016110df565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106111465780518252601f199092019160209182019101611127565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061118e5780518252601f19909201916020918201910161116f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106111d65780518252601f1990920191602091820191016111b7565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b6060611258826040518060400160405280600a8152602001691cdd5c9c1b1d5cd31bdd60b21b81525061154c565b92915050565b6060611258826040518060400160405280601081526020016f1cdd5c9c1b1d5cd51a1c995cda1bdb1960821b81525061154c565b600082820183811015610aab576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b60006112f5306116da565b15905090565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106113815780518252601f199092019160209182019101611362565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106113c95780518252601f1990920191602091820191016113aa565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106114115780518252601f1990920191602091820191016113f2565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114595780518252601f19909201916020918201910161143a565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b6b033b2e3c9fd0803ce800000090565b6000826114c357506000611258565b828202828482816114d057fe5b0414610aab5760405162461bcd60e51b8152600401808060200182810382526021815260200180611f5a6021913960400191505060405180910390fd5b6000610aab83836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506116e0565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b602083106115c35780518252601f1990920191602091820191016115a4565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b6020831061160b5780518252601f1990920191602091820191016115ec565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106116535780518252601f199092019160209182019101611634565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061169b5780518252601f19909201916020918201910161167c565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b3b151590565b6000818361176c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611731578181015183820152602001611719565b50505050905090810190601f16801561175e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161177857fe5b0495945050505050565b6040518060a0016040528060001515815260200160006001600160a01b031681526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826117f25760008555611838565b82601f1061180b57805160ff1916838001178555611838565b82800160010185558215611838579182015b8281111561183857825182559160200191906001019061181d565b50611844929150611848565b5090565b5b808211156118445760008155600101611849565b60006020828403121561186e578081fd5b8151610aab81611f13565b6000806040838503121561188b578081fd5b823561189681611f13565b915060208381013567ffffffffffffffff808211156118b3578384fd5b818601915086601f8301126118c6578384fd5b8135818111156118d257fe5b604051601f8201601f19168101850183811182821017156118ef57fe5b6040528181528382018501891015611905578586fd5b81858501868301378585838301015280955050505050509250929050565b600060208284031215611934578081fd5b81518015158114610aab578182fd5b600060208284031215611954578081fd5b5035919050565b60006020828403121561196c578081fd5b5051919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b94151585526001600160a01b0393909316602085015260408401919091526060830152608082015260a00190565b6000602080835283518082850152825b81811015611a08578581018301518582016040015282016119ec565b81811115611a195783604083870101525b50601f01601f1916929092016040019392505050565b60208082526064908201527f5b5145432d3032353030365d2d5468652062696420616d6f756e74206d75737460408201527f206578636565642074686520686967686573742062696420627920746865206d60608201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60808201526337b9329760e11b60a082015260c00190565b6020808252603a908201527f5b5145432d3032353030305d2d496e76616c69642062696420616d6f756e742c604082015279103330b4b632b21039ba30b93a103a34329030bab1ba34b7b71760311b606082015260800190565b60208082526038908201527f5b5145432d3032353030335d2d4e6f7420656e6f75676820737572706c7573206040820152773a37903334b636103a34329030bab1ba34b7b7103637ba1760411b606082015260800190565b60208082526025908201527f746f6b656e65636f6e6f6d6963732e64656661756c74416c6c6f636174696f6e60408201526450726f787960d81b606082015260800190565b6020808252601d908201527f676f7665726e65642e455044522e737572706c757341756374696f6e50000000604082015260600190565b60208082526028908201527f5b5145432d3032353030345d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b60208082526021908201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6040820152601d60fa1b606082015260800190565b60208082526046908201527f5b5145432d3032353030325d2d537572706c7573206166746572206e6574746960408201527f6e672069732062656c6f7720737572706c75732061756374696f6e207468726560608201526539b437b6321760d11b608082015260a00190565b60208082526029908201527f5b5145432d3032353030385d2d5468652061756374696f6e206973206e6f74206040820152683334b734b9b432b21760b91b606082015260800190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b60208082526034908201527f5b5145432d3032353030355d2d5468652061756374696f6e2069732066696e6960408201527339b432b216103330b4b632b2103a37903134b21760611b606082015260800190565b60208082526033908201527f5b5145432d3032353030395d2d5468652061756374696f6e2068617320616c7260408201527232b0b23c903132b2b71032bc32b1baba32b21760691b606082015260800190565b60208082526049908201527f5145432d3032353031302d4661696c656420746f207472616e7366657220746860408201527f6520686967686573742062696420666f7220636f6d6d756e69747920646973746060820152683934b13aba34b7b71760b91b608082015260a00190565b90815260200190565b918252805460ff81161515602084015260081c6001600160a01b0316604083015260018101546060830152600281015460808301526003015460a082015260c00190565b918252602082015260400190565b6001600160a01b0381168114611f2857600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212209fdffcd51a421260a516aaa75b8559cc63986ad2657aa11c51c5a1037a40ffc564736f6c63430007060033",
}

// SystemSurplusAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemSurplusAuctionMetaData.ABI instead.
var SystemSurplusAuctionABI = SystemSurplusAuctionMetaData.ABI

// SystemSurplusAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemSurplusAuctionMetaData.Bin instead.
var SystemSurplusAuctionBin = SystemSurplusAuctionMetaData.Bin

// DeploySystemSurplusAuction deploys a new Ethereum contract, binding an instance of SystemSurplusAuction to it.
func DeploySystemSurplusAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemSurplusAuction, error) {
	parsed, err := SystemSurplusAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemSurplusAuctionBin), backend)
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
func (_SystemSurplusAuction *SystemSurplusAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_SystemSurplusAuction *SystemSurplusAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		IsExecuted bool
		Bidder     common.Address
		HighestBid *big.Int
		Lot        *big.Int
		EndTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsExecuted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(bool isExecuted, address bidder, uint256 highestBid, uint256 lot, uint256 endTime)
func (_SystemSurplusAuction *SystemSurplusAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	IsExecuted bool
	Bidder     common.Address
	HighestBid *big.Int
	Lot        *big.Int
	EndTime    *big.Int
}, error) {
	return _SystemSurplusAuction.Contract.Auctions(&_SystemSurplusAuction.CallOpts, arg0)
}

// AuctionsCount is a free data retrieval call binding the contract method 0xb830620a.
//
// Solidity: function auctionsCount() view returns(uint256)
func (_SystemSurplusAuction *SystemSurplusAuctionCaller) AuctionsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "auctionsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _SystemSurplusAuction.contract.Call(opts, &out, "getRaisingBid", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	event.Raw = log
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
	event.Raw = log
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

// FilterExecuted is a free log retrieval operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
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

// WatchExecuted is a free log subscription operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
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

// ParseExecuted is a log parse operation binding the contract event 0x2510a5567d82839751607a8a4f3d96799802809b132d2364298542dc215ce49b.
//
// Solidity: event Executed(uint256 _auctionId, address indexed _bidder, (bool,address,uint256,uint256,uint256) _info)
func (_SystemSurplusAuction *SystemSurplusAuctionFilterer) ParseExecuted(log types.Log) (*SystemSurplusAuctionExecuted, error) {
	event := new(SystemSurplusAuctionExecuted)
	if err := _SystemSurplusAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
