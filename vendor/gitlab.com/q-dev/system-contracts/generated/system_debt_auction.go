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
	Bin: "0x608060405234801561001057600080fd5b50611fe4806100206000396000f3fe60806040526004361061006f5760003560e01c8063065de74c1461007b5780630c0b86ca146100b0578063454a2ab3146100d4578063500ff7e6146100f4578063571a26a0146101095780636146195414610170578063a216592014610185578063f399e22e146101a557600080fd5b3661007657005b600080fd5b34801561008757600080fd5b5061009b610096366004611aaf565b6101c7565b60405190151581526020015b60405180910390f35b3480156100bc57600080fd5b506100c660035481565b6040519081526020016100a7565b3480156100e057600080fd5b5061009b6100ef366004611aaf565b610b0b565b34801561010057600080fd5b5061009b610f25565b34801561011557600080fd5b5061015f610124366004611aaf565b6002602081905260009182526040909120805460018201549282015460039092015460ff8216936101009092046001600160a01b0316929085565b6040516100a7959493929190611b00565b34801561017c57600080fd5b5061009b610f57565b34801561019157600080fd5b506100c66101a0366004611aaf565b6113ba565b3480156101b157600080fd5b506101c56101c0366004611b65565b61142e565b005b6000806101d2611519565b905060008060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027161028c6001805461020990611c29565b80601f016020809104026020016040519081016040528092919081815260200182805461023590611c29565b80156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b50505050506115c8565b6040518263ffffffff1660e01b81526004016102a89190611c94565b60206040518083038186803b1580156102c057600080fd5b505afa1580156102d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f89190611cc7565b60008054604080518082018252601c81527f746f6b656e65636f6e6f6d6963732e73797374656d526573657276650000000060208201529051633fb9027160e01b81529394509192620100009091046001600160a01b031691633fb90271916103649190600401611c94565b60206040518083038186803b15801561037c57600080fd5b505afa158015610390573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b49190611cc7565b6003546000818152600260208190526040808320815160a0810190925280549596509394929390929091839160ff16908111156103f3576103f3611ac8565b600281111561040457610404611ac8565b8152815461010090046001600160a01b031660208201526001808301546040830152600283015460608301526003909201546080909101529091508151600281111561045257610452611ac8565b14156104cb5760405162461bcd60e51b815260206004820152603c60248201527f5b5145432d3032363030305d2d4f6e6c79206f6e652073797374656d2064656260448201527f742061756374696f6e2063616e2072756e20617420612074696d652e0000000060648201526084015b60405180910390fd5b836001600160a01b0316633d03f7976040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561050657600080fd5b505af115801561051a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053e9190611ceb565b50846001600160a01b031663498bff006105e16001805461055e90611c29565b80601f016020809104026020016040519081016040528092919081815260200182805461058a90611c29565b80156105d75780601f106105ac576101008083540402835291602001916105d7565b820191906000526020600020905b8154815290600101906020018083116105ba57829003601f168201915b5050505050611633565b6040518263ffffffff1660e01b81526004016105fd9190611c94565b60206040518083038186803b15801561061557600080fd5b505afa158015610629573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064d9190611d0d565b846001600160a01b03166314a6bf0f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561068657600080fd5b505afa15801561069a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106be9190611d0d565b116107285760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3032363030315d2d53797374656d20646562742069732062656c6f6044820152733b9030bab1ba34b7b7103a343932b9b437b6321760611b60648201526084016104c2565b610730611664565b6001600160a01b03166323b872dd33308a6040518463ffffffff1660e01b815260040161075f93929190611d26565b602060405180830381600087803b15801561077957600080fd5b505af115801561078d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b19190611ceb565b5060405162498bff60e81b815260206004820152601860248201527719dbdd995c9b99590b915411148b9c995cd95c9d99531bdd60421b60448201526000906001600160a01b0387169063498bff009060640160206040518083038186803b15801561081c57600080fd5b505afa158015610830573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108549190611d0d565b604051632e1a7d4d60e01b8152600481018290529091506001600160a01b03851690632e1a7d4d90602401602060405180830381600087803b15801561089957600080fd5b505af11580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d19190611ceb565b6109695760405162461bcd60e51b815260206004820152605d60248201527f5b5145432d3032363030325d2d4661696c656420746f2077697468647261772060448201527f66726f6d207468652053797374656d5265736572766520636f6e74726163742c60648201527f206661696c656420746f207374617274207468652061756374696f6e2e000000608482015260a4016104c2565b600182819052503360208301526040808301899052608083018290525162498bff60e81b81526001600160a01b0387169063498bff00906109db906004016020808252601a90820152790676f7665726e65642e455044522e6465627441756374696f6e560341b604082015260600190565b60206040518083038186803b1580156109f357600080fd5b505afa158015610a07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2b9190611d0d565b610a359042611d60565b60608301526000838152600260208190526040909120835181548593839160ff1916906001908490811115610a6c57610a6c611ac8565b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339084907fd7d053b3d3fc9e8145896f8940adeb377f6735866f2161ef1660dae9f049247590610aec908c815260200190565b60405180910390a3610afd83611738565b506001979650505050505050565b6003546000818152600260208190526040808320815160a0810190925280549394938593839160ff1690811115610b4457610b44611ac8565b6002811115610b5557610b55611ac8565b8152815461010090046001600160a01b0316602082015260018083015460408301526002830154606083015260039092015460809091015290915081516002811115610ba357610ba3611ac8565b14610c0f5760405162461bcd60e51b815260206004820152603660248201527f5b5145432d3032363030335d2d5468652061756374696f6e206973206e6f742060448201527530b1ba34bb3296103330b4b632b2103a37903134b21760511b60648201526084016104c2565b4281606001511015610c805760405162461bcd60e51b815260206004820152603460248201527f5b5145432d3032363030345d2d5468652061756374696f6e2069732066696e6960448201527339b432b216103330b4b632b2103a37903134b21760611b60648201526084016104c2565b610c89826113ba565b841015610d325760405162461bcd60e51b8152602060048201526064602482018190527f5b5145432d3032363030355d2d5468652062696420616d6f756e74206d75737460448301527f206578636565642074686520686967686573742062696420627920746865206d908201527f696e696d756d20696e6372656d656e742070657263656e74616765206f72206d60848201526337b9329760e11b60a482015260c4016104c2565b6000610d3c611664565b9050806001600160a01b031663a9059cbb836020015184604001516040518363ffffffff1660e01b8152600401610d74929190611d78565b602060405180830381600087803b158015610d8e57600080fd5b505af1158015610da2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc69190611ceb565b506040516323b872dd60e01b81526001600160a01b038216906323b872dd90610df790339030908a90600401611d26565b602060405180830381600087803b158015610e1157600080fd5b505af1158015610e25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e499190611ceb565b50336020808401919091526040808401879052600085815260029283905220835181548593839160ff1916906001908490811115610e8957610e89611ac8565b0217905550602082015181546001600160a01b0390911661010002610100600160a81b031990911617815560408083015160018301556060830151600283015560809092015160039091015551339084907fdcd726e11f8b5e160f00290f0fe3a1abb547474e53a8e7a8f49a85e7b1ca319990610f099089815260200190565b60405180910390a3610f1a83611738565b506001949350505050565b6000600160035460009081526002602081905260409091205460ff1690811115610f5157610f51611ac8565b14905090565b6003546000818152600260208190526040808320815160a0810190925280549394938593839160ff1690811115610f9057610f90611ac8565b6002811115610fa157610fa1611ac8565b815281546001600160a01b036101009091048116602083015260018084015460408401526002840154606084015260039093015460809092019190915260008054835494955090936201000090910490911691633fb9027191611008919061020990611c29565b6040518263ffffffff1660e01b81526004016110249190611c94565b60206040518083038186803b15801561103c57600080fd5b505afa158015611050573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110749190611cc7565b90506000611080611664565b9050826060015142116110e75760405162461bcd60e51b815260206004820152602960248201527f5b5145432d3032363030365d2d5468652061756374696f6e206973206e6f74206044820152683334b734b9b432b21760b91b60648201526084016104c2565b6001835160028111156110fc576110fc611ac8565b1461116b5760405162461bcd60e51b815260206004820152603960248201527f5b5145432d3032363030375d2d5468652061756374696f6e206973206e6f742060448201527830b1ba34bb32961032bc32b1baba34b7b7103330b4b632b21760391b60648201526084016104c2565b61117484611738565b604080840151905163a9059cbb60e01b81526001600160a01b0383169163a9059cbb916111a5918691600401611d78565b602060405180830381600087803b1580156111bf57600080fd5b505af11580156111d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f79190611ceb565b506003805460009081526002602081905260408220805460ff191690911790558154919061122483611d91565b909155505060008054604080518082018252601b81527a746f6b656e65636f6e6f6d6963732e707573685061796d656e747360281b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161128d91600401611c94565b60206040518083038186803b1580156112a557600080fd5b505afa1580156112b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112dd9190611cc7565b60808501516020860151604051632377789f60e21b81529293506001600160a01b03841692638ddde27c929161131591600401611dac565b6020604051808303818588803b15801561132e57600080fd5b505af1158015611342573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906113679190611ceb565b5083602001516001600160a01b0316857f342522ce238e8878f82d1a020742d0d67f988f004878a3b9842b52b02d278d64866040516113a69190611dc0565b60405180910390a360019550505050505090565b600060035482111561141f5760405162461bcd60e51b815260206004820152602860248201527f5b5145432d3032363030395d2d5468652061756374696f6e20646f6573206e6f6044820152673a1032bc34b9ba1760c11b60648201526084016104c2565b61142882611836565b92915050565b600054610100900460ff1680611447575060005460ff16155b6114aa5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104c2565b600054610100900460ff161580156114cc576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038616021790558151611501906001906020850190611a16565b508015611514576000805461ff00191690555b505050565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611f8d602291396040518263ffffffff1660e01b81526004016115739190611c94565b60206040518083038186803b15801561158b57600080fd5b505afa15801561159f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c39190611cc7565b905090565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461161d9493879201611e09565b6040516020818303038152906040529050919050565b6060611428826040518060400160405280600d81526020016c1919589d151a1c995cda1bdb19609a1b815250611952565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb9027161171c6001805461169990611c29565b80601f01602080910402602001604051908101604052809291908181526020018280546116c590611c29565b80156117125780601f106116e757610100808354040283529160200191611712565b820191906000526020600020905b8154815290600101906020018083116116f557829003601f168201915b50505050506119c1565b6040518263ffffffff1660e01b81526004016115739190611c94565b600160008281526002602081905260409091205460ff169081111561175f5761175f611ac8565b141561183357600061176f611664565b600083815260026020526040908190206001015490516370a0823160e01b8152919250906001600160a01b038316906370a08231906117b2903090600401611dac565b60206040518083038186803b1580156117ca57600080fd5b505afa1580156117de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118029190611d0d565b101561181057611810611e78565b60008281526002602052604090206003015447101561183157611831611e78565b505b50565b600080611841611519565b60405162498bff60e81b815260206004820152602160248201527f676f7665726e65642e455044522e61756374696f6e4d696e496e6372656d656e6044820152601d60fa1b60648201526001600160a01b03919091169063498bff009060840160206040518083038186803b1580156118b957600080fd5b505afa1580156118cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f19190611d0d565b600084815260026020526040812060010154919250816b033b2e3c9fd0803ce800000061191e8585611e8e565b6119289190611ead565b6119329190611d60565b90508181141561194a578061194681611d91565b9150505b949350505050565b60606040518060400160405280600e81526020016d33b7bb32b93732b21722a822291760911b81525083604051806040016040528060018152602001605f60f81b815250846040516020016119aa9493929190611ecf565b604051602081830303815290604052905092915050565b60408051808201825260048152636465666960e01b602080830191909152825180840184526001808252601760f91b8284018190528551808701875291825281840152935160609461161d9493879201611f26565b828054611a2290611c29565b90600052602060002090601f016020900481019282611a445760008555611a8a565b82601f10611a5d57805160ff1916838001178555611a8a565b82800160010185558215611a8a579182015b82811115611a8a578251825591602001919060010190611a6f565b50611a96929150611a9a565b5090565b5b80821115611a965760008155600101611a9b565b600060208284031215611ac157600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60038110611afc57634e487b7160e01b600052602160045260246000fd5b9052565b60a08101611b0e8288611ade565b6001600160a01b0395909516602082015260408101939093526060830191909152608090910152919050565b6001600160a01b038116811461183357600080fd5b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611b7857600080fd5b8235611b8381611b3a565b9150602083013567ffffffffffffffff80821115611ba057600080fd5b818501915085601f830112611bb457600080fd5b813581811115611bc657611bc6611b4f565b604051601f8201601f19908116603f01168101908382118183101715611bee57611bee611b4f565b81604052828152886020848701011115611c0757600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600181811c90821680611c3d57607f821691505b60208210811415611c5e57634e487b7160e01b600052602260045260246000fd5b50919050565b60005b83811015611c7f578181015183820152602001611c67565b83811115611c8e576000848401525b50505050565b6020815260008251806020840152611cb3816040850160208701611c64565b601f01601f19169190910160400192915050565b600060208284031215611cd957600080fd5b8151611ce481611b3a565b9392505050565b600060208284031215611cfd57600080fd5b81518015158114611ce457600080fd5b600060208284031215611d1f57600080fd5b5051919050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611d7357611d73611d4a565b500190565b6001600160a01b03929092168252602082015260400190565b6000600019821415611da557611da5611d4a565b5060010190565b6001600160a01b0391909116815260200190565b600060a082019050611dd3828451611ade565b60018060a01b03602084015116602083015260408301516040830152606083015160608301526080830151608083015292915050565b60008551611e1b818460208a01611c64565b855190830190611e2f818360208a01611c64565b8551910190611e42818360208901611c64565b8451910190611e55818360208801611c64565b6c73797374656d42616c616e636560981b9101908152600d019695505050505050565b634e487b7160e01b600052600160045260246000fd5b6000816000190483118215151615611ea857611ea8611d4a565b500290565b600082611eca57634e487b7160e01b600052601260045260246000fd5b500490565b60008551611ee1818460208a01611c64565b855190830190611ef5818360208a01611c64565b8551910190611f08818360208901611c64565b8451910190611f1b818360208801611c64565b019695505050505050565b60008551611f38818460208a01611c64565b855190830190611f4c818360208a01611c64565b8551910190611f5f818360208901611c64565b8451910190611f72818360208801611c64565b6331b7b4b760e11b9101908152600401969550505050505056fe676f7665726e616e63652e657870657274732e455044522e706172616d6574657273a26469706673582212201e2c134553c5c2012d8f0fe4418cc002571ff85134846b4f4d44c8cc29be11d364736f6c63430008090033",
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
