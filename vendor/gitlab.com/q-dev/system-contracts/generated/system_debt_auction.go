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

// SystemDebtAuctionAuctionInfo is an auto generated low-level Go binding around an user-defined struct.
type SystemDebtAuctionAuctionInfo struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}

// SystemDebtAuctionMetaData contains all meta data concerning the SystemDebtAuction contract.
var SystemDebtAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSystemDebtAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLot\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSystemDebtAuction.AuctionInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"enumSystemDebtAuction.AuctionStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAuctionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_stc\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"getRaisingBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasActiveAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612470806100206000396000f3fe60806040526004361061006f5760003560e01c8063065de74c1461007b5780630c0b86ca146100b1578063454a2ab3146100d3578063500ff7e6146100f3578063571a26a0146101085780636146195414610139578063a21659201461014e578063f399e22e1461016e57610076565b3661007657005b600080fd5b34801561008757600080fd5b5061009b610096366004611db0565b610190565b6040516100a89190611e3f565b60405180910390f35b3480156100bd57600080fd5b506100c6610990565b6040516100a891906123cd565b3480156100df57600080fd5b5061009b6100ee366004611db0565b610996565b3480156100ff57600080fd5b5061009b610c70565b34801561011457600080fd5b50610128610123366004611db0565b610c9c565b6040516100a8959493929190611e4a565b34801561014557600080fd5b5061009b610cd7565b34801561015a57600080fd5b506100c6610169366004611db0565b6110dd565b34801561017a57600080fd5b5061018e610189366004611ce6565b611110565b005b60008054604051633fb9027160e01b815282916201000090046001600160a01b031690633fb90271906101c59060040161206d565b60206040518083038186803b1580156101dd57600080fd5b505afa1580156101f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102159190611cca565b6000805460018054604080516020600261010085871615026000190190941693909304601f81018490048402820184019092528181529596509394620100009093046001600160a01b031693633fb90271936102cb9390929091908301828280156102c15780601f10610296576101008083540402835291602001916102c1565b820191906000526020600020905b8154815290600101906020018083116102a457829003601f168201915b50505050506111ea565b6040518263ffffffff1660e01b81526004016102e79190611e84565b60206040518083038186803b1580156102ff57600080fd5b505afa158015610313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103379190611cca565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb902719061037090600401611f5a565b60206040518083038186803b15801561038857600080fd5b505afa15801561039c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c09190611cca565b6003546000818152600260208190526040808320815160a0810190925280549596509394929390929091839160ff16908111156103f957fe5b600281111561040457fe5b8152815461010090046001600160a01b031660208201526001808301546040830152600283015460608301526003909201546080909101529091508151600281111561044c57fe5b14156104735760405162461bcd60e51b815260040161046a9061226a565b60405180910390fd5b836001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104ae57600080fd5b505af11580156104c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e69190611d90565b5060018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281526001600160a01b0389169363498bff009361058a938301828280156105805780601f1061055557610100808354040283529160200191610580565b820191906000526020600020905b81548152906001019060200180831161056357829003601f168201915b505050505061139c565b6040518263ffffffff1660e01b81526004016105a69190611e84565b60206040518083038186803b1580156105be57600080fd5b505afa1580156105d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f69190611dc8565b846001600160a01b03166314a6bf0f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561062f57600080fd5b505afa158015610643573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106679190611dc8565b116106845760405162461bcd60e51b815260040161046a90611f91565b61068c6113cd565b6001600160a01b03166323b872dd33308a6040518463ffffffff1660e01b81526004016106bb93929190611e02565b602060405180830381600087803b1580156106d557600080fd5b505af11580156106e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070d9190611d90565b5060405162498bff60e81b81526000906001600160a01b0387169063498bff009061073a906004016122fc565b60206040518083038186803b15801561075257600080fd5b505afa158015610766573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078a9190611dc8565b604051632e1a7d4d60e01b81529091506001600160a01b03851690632e1a7d4d906107b99084906004016123cd565b602060405180830381600087803b1580156107d357600080fd5b505af11580156107e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080b9190611d90565b6108275760405162461bcd60e51b815260040161046a90611ed7565b600182523360208301526040808301899052608083018290525162498bff60e81b81526108c2906001600160a01b0388169063498bff009061086b90600401612039565b60206040518083038186803b15801561088357600080fd5b505afa158015610897573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bb9190611dc8565b42906114ee565b60608301526000838152600260208190526040909120835181548593839160ff19169060019084908111156108f357fe5b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339084907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610971908c906123cd565b60405180910390a36109828361154d565b506001979650505050505050565b60035481565b6003546000818152600260208190526040808320815160a0810190925280549394938593839160ff16908111156109c957fe5b60028111156109d457fe5b8152815461010090046001600160a01b0316602082015260018083015460408301526002830154606083015260039092015460809091015290915081516002811115610a1c57fe5b14610a395760405162461bcd60e51b815260040161046a9061232e565b4281606001511015610a5d5760405162461bcd60e51b815260040161046a90611fe5565b610a66826110dd565b841015610a855760405162461bcd60e51b815260040161046a906121da565b6000610a8f6113cd565b9050806001600160a01b031663a9059cbb836020015184604001516040518363ffffffff1660e01b8152600401610ac7929190611e26565b602060405180830381600087803b158015610ae157600080fd5b505af1158015610af5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b199190611d90565b506040516323b872dd60e01b81526001600160a01b038216906323b872dd90610b4a90339030908a90600401611e02565b602060405180830381600087803b158015610b6457600080fd5b505af1158015610b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9c9190611d90565b50336020808401919091526040808401879052600085815260029283905220835181548593839160ff1916906001908490811115610bd657fe5b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339084907fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca319990610c549089906123cd565b60405180910390a3610c658361154d565b506001949350505050565b6000600160035460009081526002602081905260409091205460ff1690811115610c9657fe5b14905090565b6002602081905260009182526040909120805460018201549282015460039092015460ff8216936101009092046001600160a01b0316929085565b6003546000818152600260208190526040808320815160a0810190925280549394938593839160ff1690811115610d0a57fe5b6002811115610d1557fe5b81526020016000820160019054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016001820154815260200160028201548152602001600382015481525050905060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271610dfb60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c15780601f10610296576101008083540402835291602001916102c1565b6040518263ffffffff1660e01b8152600401610e179190611e84565b60206040518083038186803b158015610e2f57600080fd5b505afa158015610e43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e679190611cca565b90506000610e736113cd565b905082606001514211610e985760405162461bcd60e51b815260040161046a906120f0565b600183516002811115610ea757fe5b14610ec45760405162461bcd60e51b815260040161046a90612139565b610ecd8461154d565b604080840151905163a9059cbb60e01b81526001600160a01b0383169163a9059cbb91610efe918691600401611e26565b602060405180830381600087803b158015610f1857600080fd5b505af1158015610f2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f509190611d90565b506003546000908152600260208190526040909120805460ff1916600183021790555060038054600101905560008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610fb0906004016122c7565b60206040518083038186803b158015610fc857600080fd5b505afa158015610fdc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110009190611cca565b60808501516020860151604051632377789f60e21b81529293506001600160a01b03841692638ddde27c929161103891600401611dee565b6020604051808303818588803b15801561105157600080fd5b505af1158015611065573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061108a9190611d90565b5083602001516001600160a01b0316857f342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64866040516110c99190612384565b60405180910390a360019550505050505090565b60006003548211156111015760405162461bcd60e51b815260040161046a90612192565b61110a82611639565b92915050565b600054610100900460ff16806111295750611129611791565b80611137575060005460ff16155b6111725760405162461bcd60e51b815260040180806020018281038252602e8152602001806123ec602e913960400191505060405180910390fd5b600054610100900460ff1615801561119d576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b0386160217905581516111d2906001906020850190611c29565b5080156111e5576000805461ff00191690555b505050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106112705780518252601f199092019160209182019101611251565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106112b85780518252601f199092019160209182019101611299565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106113005780518252601f1990920191602091820191016112e1565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106113485780518252601f199092019160209182019101611329565b5181516020939093036101000a60001901801990911692169190911790526c73797374656d42616c616e636560981b92019182525060408051808303601219018152600d9092019052979650505050505050565b606061110a826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b8152506117a2565b6000805460018054604080516020600261010085871615026000190190941693909304601f8101849004840282018401909252818152620100009094046001600160a01b031693633fb902719361147d93919290918301828280156114735780601f1061144857610100808354040283529160200191611473565b820191906000526020600020905b81548152906001019060200180831161145657829003601f168201915b5050505050611930565b6040518263ffffffff1660e01b81526004016114999190611e84565b60206040518083038186803b1580156114b157600080fd5b505afa1580156114c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e99190611cca565b905090565b600082820183811015611546576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b9392505050565b600160008281526002602081905260409091205460ff169081111561156e57fe5b141561163657600061157e6113cd565b600083815260026020526040908190206001015490516370a0823160e01b8152919250906001600160a01b038316906370a08231906115c1903090600401611dee565b60206040518083038186803b1580156115d957600080fd5b505afa1580156115ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116119190611dc8565b101561161957fe5b60008281526002602052604090206003015447101561163457fe5b505b50565b60008054604051633fb9027160e01b815282916201000090046001600160a01b031690633fb902719061166e9060040161206d565b60206040518083038186803b15801561168657600080fd5b505afa15801561169a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116be9190611cca565b6001600160a01b031663498bff006040518163ffffffff1660e01b81526004016116e7906120af565b60206040518083038186803b1580156116ff57600080fd5b505afa158015611713573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117379190611dc8565b6000848152600260205260408120600101549192506117718261176b61175b611ad9565b6117658688611ae9565b90611b42565b906114ee565b905081811415611789576117868160016114ee565b90505b949350505050565b600061179c30611b81565b15905090565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020018085805190602001908083835b602083106118195780518252601f1990920191602091820191016117fa565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106118615780518252601f199092019160209182019101611842565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106118a95780518252601f19909201916020918201910161188a565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106118f15780518252601f1990920191602091820191016118d2565b6001836020036101000a038019825116818451168082178552505050505050905001945050505050604051602081830303815290604052905092915050565b6060604051806040016040528060048152602001636465666960e01b815250604051806040016040528060018152602001601760f91b81525083604051806040016040528060018152602001601760f91b8152506040516020018085805190602001908083835b602083106119b65780518252601f199092019160209182019101611997565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b602083106119fe5780518252601f1990920191602091820191016119df565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611a465780518252601f199092019160209182019101611a27565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611a8e5780518252601f199092019160209182019101611a6f565b5181516020939093036101000a60001901801990911692169190911790526331b7b4b760e11b92019182525060408051808303601b1901815260049092019052979650505050505050565b6b033b2e3c9fd0803ce800000090565b600082611af85750600061110a565b82820282848281611b0557fe5b04146115465760405162461bcd60e51b815260040180806020018281038252602181526020018061241a6021913960400191505060405180910390fd5b600061154683836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b815250611b87565b3b151590565b60008183611c135760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611bd8578181015183820152602001611bc0565b50505050905090810190601f168015611c055780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581611c1f57fe5b0495945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282611c5f5760008555611ca5565b82601f10611c7857805160ff1916838001178555611ca5565b82800160010185558215611ca5579182015b82811115611ca5578251825591602001919060010190611c8a565b50611cb1929150611cb5565b5090565b5b80821115611cb15760008155600101611cb6565b600060208284031215611cdb578081fd5b8151611546816123d6565b60008060408385031215611cf8578081fd5b8235611d03816123d6565b915060208381013567ffffffffffffffff80821115611d20578384fd5b818601915086601f830112611d33578384fd5b813581811115611d3f57fe5b604051601f8201601f1916810185018381118282101715611d5c57fe5b6040528181528382018501891015611d72578586fd5b81858501868301378585838301015280955050505050509250929050565b600060208284031215611da1578081fd5b81518015158114611546578182fd5b600060208284031215611dc1578081fd5b5035919050565b600060208284031215611dd9578081fd5b5051919050565b60038110611dea57fe5b9052565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b60a08101611e588288611de0565b6001600160a01b0395909516602082015260408101939093526060830191909152608090910152919050565b6000602080835283518082850152825b81811015611eb057858101830151858201604001528201611e94565b81811115611ec15783604083870101525b50601f01601f1916929092016040019392505050565b6020808252605d908201527f5b5145432d3032363030325d2d4661696c656420746f2077697468647261772060408201527f66726f6d207468652053797374656d5265736572766520636f6e74726163742c60608201527f206661696c656420746f207374617274207468652061756374696f6e2e000000608082015260a00190565b6020808252601c908201527f746f6b656e65636f6e6f6d6963732e73797374656d5265736572766500000000604082015260600190565b60208082526034908201527f5b5145432d3032363030315d2d53797374656d20646562742069732062656c6f6040820152733b9030bab1ba34b7b7103a343932b9b437b6321760611b606082015260800190565b60208082526034908201527f5b5145432d3032363030345d2d5468652061756374696f6e2069732066696e6960408201527339b432b216103330b4b632b2103a37903134b21760611b606082015260800190565b6020808252601a90820152790676f7665726e65642e455044522e6465627441756374696f6e560341b604082015260600190565b60208082526022908201527f676f7665726e616e63652e657870657274732e455044522e706172616d657465604082015261727360f01b606082015260800190565b60208082526021908201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6040820152601d60fa1b606082015260800190565b60208082526029908201527f5b5145432d3032363030365d2d5468652061756374696f6e206973206e6f74206040820152683334b734b9b432b21760b91b606082015260800190565b60208082526039908201527f5b5145432d3032363030375d2d5468652061756374696f6e206973206e6f742060408201527830b1ba34bb32961032bc32b1baba34b7b7103330b4b632b21760391b606082015260800190565b60208082526028908201527f5b5145432d3032363030395d2d5468652061756374696f6e20646f6573206e6f6040820152673a1032bc34b9ba1760c11b606082015260800190565b60208082526064908201527f5b5145432d3032363030355d2d5468652062696420616d6f756e74206d75737460408201527f206578636565642074686520686967686573742062696420627920746865206d60608201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60808201526337b9329760e11b60a082015260c00190565b6020808252603c908201527f5b5145432d3032363030305d2d4f6e6c79206f6e652073797374656d2064656260408201527f742061756374696f6e2063616e2072756e20617420612074696d652e00000000606082015260800190565b6020808252601b908201527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b604082015260600190565b60208082526018908201527719dbdd995c9b99590b915411148b9c995cd95c9d99531bdd60421b604082015260600190565b60208082526036908201527f5b5145432d3032363030335d2d5468652061756374696f6e206973206e6f742060408201527530b1ba34bb3296103330b4b632b2103a37903134b21760511b606082015260800190565b600060a082019050612397828451611de0565b60018060a01b03602084015116602083015260408301516040830152606083015160608301526080830151608083015292915050565b90815260200190565b6001600160a01b038116811461163657600080fdfe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122024c1fd94387b36b6662b7f6f6a9a231d9ed5573b2b0a1252f3896fd4098bb34164736f6c63430007060033",
}

// SystemDebtAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemDebtAuctionMetaData.ABI instead.
var SystemDebtAuctionABI = SystemDebtAuctionMetaData.ABI

// SystemDebtAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemDebtAuctionMetaData.Bin instead.
var SystemDebtAuctionBin = SystemDebtAuctionMetaData.Bin

// DeploySystemDebtAuction deploys a new Ethereum contract, binding an instance of SystemDebtAuction to it.
func DeploySystemDebtAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemDebtAuction, error) {
	parsed, err := SystemDebtAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemDebtAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemDebtAuction{SystemDebtAuctionCaller: SystemDebtAuctionCaller{contract: contract}, SystemDebtAuctionTransactor: SystemDebtAuctionTransactor{contract: contract}, SystemDebtAuctionFilterer: SystemDebtAuctionFilterer{contract: contract}}, nil
}

// SystemDebtAuction is an auto generated Go binding around an Ethereum contract.
type SystemDebtAuction struct {
	SystemDebtAuctionCaller     // Read-only binding to the contract
	SystemDebtAuctionTransactor // Write-only binding to the contract
	SystemDebtAuctionFilterer   // Log filterer for contract events
}

// SystemDebtAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemDebtAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemDebtAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemDebtAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemDebtAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemDebtAuctionSession struct {
	Contract     *SystemDebtAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SystemDebtAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemDebtAuctionCallerSession struct {
	Contract *SystemDebtAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SystemDebtAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemDebtAuctionTransactorSession struct {
	Contract     *SystemDebtAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SystemDebtAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemDebtAuctionRaw struct {
	Contract *SystemDebtAuction // Generic contract binding to access the raw methods on
}

// SystemDebtAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemDebtAuctionCallerRaw struct {
	Contract *SystemDebtAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SystemDebtAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemDebtAuctionTransactorRaw struct {
	Contract *SystemDebtAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemDebtAuction creates a new instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuction(address common.Address, backend bind.ContractBackend) (*SystemDebtAuction, error) {
	contract, err := bindSystemDebtAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuction{SystemDebtAuctionCaller: SystemDebtAuctionCaller{contract: contract}, SystemDebtAuctionTransactor: SystemDebtAuctionTransactor{contract: contract}, SystemDebtAuctionFilterer: SystemDebtAuctionFilterer{contract: contract}}, nil
}

// NewSystemDebtAuctionCaller creates a new read-only instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionCaller(address common.Address, caller bind.ContractCaller) (*SystemDebtAuctionCaller, error) {
	contract, err := bindSystemDebtAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionCaller{contract: contract}, nil
}

// NewSystemDebtAuctionTransactor creates a new write-only instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemDebtAuctionTransactor, error) {
	contract, err := bindSystemDebtAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionTransactor{contract: contract}, nil
}

// NewSystemDebtAuctionFilterer creates a new log filterer instance of SystemDebtAuction, bound to a specific deployed contract.
func NewSystemDebtAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemDebtAuctionFilterer, error) {
	contract, err := bindSystemDebtAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionFilterer{contract: contract}, nil
}

// bindSystemDebtAuction binds a generic wrapper to an already deployed contract.
func bindSystemDebtAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemDebtAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemDebtAuction.Contract.SystemDebtAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.SystemDebtAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemDebtAuction *SystemDebtAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.SystemDebtAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemDebtAuction *SystemDebtAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemDebtAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemDebtAuction *SystemDebtAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemDebtAuction *SystemDebtAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	var out []interface{}
	err := _SystemDebtAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Status     uint8
		Bidder     common.Address
		HighestBid *big.Int
		EndTime    *big.Int
		ReserveLot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveLot = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionSession) Auctions(arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	return _SystemDebtAuction.Contract.Auctions(&_SystemDebtAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(uint8 status, address bidder, uint256 highestBid, uint256 endTime, uint256 reserveLot)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) Auctions(arg0 *big.Int) (struct {
	Status     uint8
	Bidder     common.Address
	HighestBid *big.Int
	EndTime    *big.Int
	ReserveLot *big.Int
}, error) {
	return _SystemDebtAuction.Contract.Auctions(&_SystemDebtAuction.CallOpts, arg0)
}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCaller) CurrentAuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemDebtAuction.contract.Call(opts, &out, "currentAuctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionSession) CurrentAuctionId() (*big.Int, error) {
	return _SystemDebtAuction.Contract.CurrentAuctionId(&_SystemDebtAuction.CallOpts)
}

// CurrentAuctionId is a free data retrieval call binding the contract method 0x0c0b86ca.
//
// Solidity: function currentAuctionId() view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) CurrentAuctionId() (*big.Int, error) {
	return _SystemDebtAuction.Contract.CurrentAuctionId(&_SystemDebtAuction.CallOpts)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCaller) GetRaisingBid(opts *bind.CallOpts, _auctionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemDebtAuction.contract.Call(opts, &out, "getRaisingBid", _auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemDebtAuction.Contract.GetRaisingBid(&_SystemDebtAuction.CallOpts, _auctionId)
}

// GetRaisingBid is a free data retrieval call binding the contract method 0xa2165920.
//
// Solidity: function getRaisingBid(uint256 _auctionId) view returns(uint256)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) GetRaisingBid(_auctionId *big.Int) (*big.Int, error) {
	return _SystemDebtAuction.Contract.GetRaisingBid(&_SystemDebtAuction.CallOpts, _auctionId)
}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionCaller) HasActiveAuction(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemDebtAuction.contract.Call(opts, &out, "hasActiveAuction")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) HasActiveAuction() (bool, error) {
	return _SystemDebtAuction.Contract.HasActiveAuction(&_SystemDebtAuction.CallOpts)
}

// HasActiveAuction is a free data retrieval call binding the contract method 0x500ff7e6.
//
// Solidity: function hasActiveAuction() view returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionCallerSession) HasActiveAuction() (bool, error) {
	return _SystemDebtAuction.Contract.HasActiveAuction(&_SystemDebtAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Bid(opts *bind.TransactOpts, _bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "bid", _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) Bid(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Bid(&_SystemDebtAuction.TransactOpts, _bid)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Bid(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Bid(&_SystemDebtAuction.TransactOpts, _bid)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Execute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "execute")
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) Execute() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Execute(&_SystemDebtAuction.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x61461954.
//
// Solidity: function execute() returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Execute() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Execute(&_SystemDebtAuction.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "initialize", _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Initialize(&_SystemDebtAuction.TransactOpts, _registry, _stc)
}

// Initialize is a paid mutator transaction binding the contract method 0xf399e22e.
//
// Solidity: function initialize(address _registry, string _stc) returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Initialize(_registry common.Address, _stc string) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Initialize(&_SystemDebtAuction.TransactOpts, _registry, _stc)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactor) StartAuction(opts *bind.TransactOpts, _bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.Transact(opts, "startAuction", _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionSession) StartAuction(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.StartAuction(&_SystemDebtAuction.TransactOpts, _bid)
}

// StartAuction is a paid mutator transaction binding the contract method 0x065de74c.
//
// Solidity: function startAuction(uint256 _bid) returns(bool)
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) StartAuction(_bid *big.Int) (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.StartAuction(&_SystemDebtAuction.TransactOpts, _bid)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemDebtAuction.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionSession) Receive() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Receive(&_SystemDebtAuction.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SystemDebtAuction *SystemDebtAuctionTransactorSession) Receive() (*types.Transaction, error) {
	return _SystemDebtAuction.Contract.Receive(&_SystemDebtAuction.TransactOpts)
}

// SystemDebtAuctionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the SystemDebtAuction contract.
type SystemDebtAuctionAuctionStartedIterator struct {
	Event *SystemDebtAuctionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionAuctionStarted)
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
		it.Event = new(SystemDebtAuctionAuctionStarted)
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
func (it *SystemDebtAuctionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionAuctionStarted represents a AuctionStarted event raised by the SystemDebtAuction contract.
type SystemDebtAuctionAuctionStarted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionAuctionStartedIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "AuctionStarted", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionAuctionStartedIterator{contract: _SystemDebtAuction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f0492475.
//
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionAuctionStarted, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "AuctionStarted", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionAuctionStarted)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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
// Solidity: event AuctionStarted(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseAuctionStarted(log types.Log) (*SystemDebtAuctionAuctionStarted, error) {
	event := new(SystemDebtAuctionAuctionStarted)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemDebtAuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the SystemDebtAuction contract.
type SystemDebtAuctionBidIterator struct {
	Event *SystemDebtAuctionBid // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionBid)
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
		it.Event = new(SystemDebtAuctionBid)
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
func (it *SystemDebtAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionBid represents a Bid event raised by the SystemDebtAuction contract.
type SystemDebtAuctionBid struct {
	AuctionId *big.Int
	Bidder    common.Address
	Bid       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterBid(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionBidIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionBidIterator{contract: _SystemDebtAuction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca3199.
//
// Solidity: event Bid(uint256 indexed _auctionId, address indexed _bidder, uint256 _bid)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionBid, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "Bid", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionBid)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "Bid", log); err != nil {
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
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseBid(log types.Log) (*SystemDebtAuctionBid, error) {
	event := new(SystemDebtAuctionBid)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemDebtAuctionExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SystemDebtAuction contract.
type SystemDebtAuctionExecutedIterator struct {
	Event *SystemDebtAuctionExecuted // Event containing the contract specifics and raw log

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
func (it *SystemDebtAuctionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemDebtAuctionExecuted)
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
		it.Event = new(SystemDebtAuctionExecuted)
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
func (it *SystemDebtAuctionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemDebtAuctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemDebtAuctionExecuted represents a Executed event raised by the SystemDebtAuction contract.
type SystemDebtAuctionExecuted struct {
	AuctionId *big.Int
	Bidder    common.Address
	Info      SystemDebtAuctionAuctionInfo
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) FilterExecuted(opts *bind.FilterOpts, _auctionId []*big.Int, _bidder []common.Address) (*SystemDebtAuctionExecutedIterator, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.FilterLogs(opts, "Executed", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return &SystemDebtAuctionExecutedIterator{contract: _SystemDebtAuction.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SystemDebtAuctionExecuted, _auctionId []*big.Int, _bidder []common.Address) (event.Subscription, error) {

	var _auctionIdRule []interface{}
	for _, _auctionIdItem := range _auctionId {
		_auctionIdRule = append(_auctionIdRule, _auctionIdItem)
	}
	var _bidderRule []interface{}
	for _, _bidderItem := range _bidder {
		_bidderRule = append(_bidderRule, _bidderItem)
	}

	logs, sub, err := _SystemDebtAuction.contract.WatchLogs(opts, "Executed", _auctionIdRule, _bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemDebtAuctionExecuted)
				if err := _SystemDebtAuction.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64.
//
// Solidity: event Executed(uint256 indexed _auctionId, address indexed _bidder, (uint8,address,uint256,uint256,uint256) _info)
func (_SystemDebtAuction *SystemDebtAuctionFilterer) ParseExecuted(log types.Log) (*SystemDebtAuctionExecuted, error) {
	event := new(SystemDebtAuctionExecuted)
	if err := _SystemDebtAuction.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
