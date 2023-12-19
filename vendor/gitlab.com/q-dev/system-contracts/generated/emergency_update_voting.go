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

// EmergencyUpdateVotingMetaData contains all meta data concerning the EmergencyUpdateVoting contract.
var EmergencyUpdateVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"_proposal\",\"type\":\"tuple\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIVoting.VotingOption\",\"name\":\"_votingOption\",\"type\":\"uint8\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVetoPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesAgainst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteAgainst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611cb0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c8063013cf08b146100b4578063307a064f146100e057806349c2a1a6146101355780635277b4ae146101565780635c622a0e14610194578063750e443a146101b457806386a50535146101c9578063b6f61f66146101dc578063c4d66de8146101ff578063c7f758a814610212578063d40b65eb14610232578063da35c66414610255578063fe0d94c11461025e575b600080fd5b6100c76100c2366004611762565b610271565b6040516100d79493929190611816565b60405180910390f35b6100f36100ee366004611762565b6103a2565b6040516100d79190600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b61014861014336600461187f565b610629565b6040519081526020016100d7565b610184610164366004611948565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100d7565b6101a76101a2366004611762565b610b80565b6040516100d7919061198e565b6101c76101c2366004611762565b610e39565b005b6101c76101d7366004611762565b610f22565b6101486101ea366004611762565b6000908152600160205260409020600a015490565b6101c761020d3660046119a8565b611007565b610225610220366004611762565b6110dc565b6040516100d79190611a30565b610148610240366004611762565b60009081526001602052604090206009015490565b61014860035481565b6101c761026c366004611762565b61126d565b60016020526000908152604090208054819061028c90611a43565b80601f01602080910402602001604051908101604052809291908181526020018280546102b890611a43565b80156103055780601f106102da57610100808354040283529160200191610305565b820191906000526020600020905b8154815290600101906020018083116102e857829003601f168201915b50506040805161010081018252600187015481526002870154602080830191909152600388015482840152600488015460608084019190915260058901546080840152600689015460a0840152600789015460c0840152600889015460e08401528351908101845260098901548152600a89015491810191909152600b88015492810192909252600c9096015494959490935060ff169150859050565b6103aa61160a565b8160006103b682610b80565b60078111156103c7576103c7611978565b14156103ee5760405162461bcd60e51b81526004016103e590611a78565b60405180910390fd5b60008381526001602052604080822081516080810190925280548290829061041590611a43565b80601f016020809104026020016040519081016040528092919081815260200182805461044190611a43565b801561048e5780601f106104635761010080835404028352916020019161048e565b820191906000526020600020905b81548152906001019060200180831161047157829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff161515910152905061053e61160a565b602080830180516080015183525160a0015160408084019190915283015190810151905160009161056e91611ad3565b9050600061057a61133f565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b1580156105b257600080fd5b505afa1580156105c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ea9190611aeb565b905060006105f883836113f1565b6020850181905260408601515190915060009061061590856113f1565b606086015250929550505050505b50919050565b600061063361133f565b6001600160a01b031663a230c524336040518263ffffffff1660e01b815260040161065e9190611b04565b60206040518083038186803b15801561067657600080fd5b505afa15801561068a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ae9190611b18565b6106ca5760405162461bcd60e51b81526004016103e590611b3a565b60008060029054906101000a90046001600160a01b03166001600160a01b0316633fb90271604051806060016040528060228152602001611c59602291396040518263ffffffff1660e01b81526004016107249190611b97565b60206040518083038186803b15801561073c57600080fd5b505afa158015610750573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107749190611baa565b905061077e611639565b838152602081015142905260405162498bff60e81b81526001600160a01b0383169063498bff00906107e2906004016020808252818101527f636f6e737469747574696f6e2e766f74696e672e656d67515570646174655650604082015260600190565b60206040518083038186803b1580156107fa57600080fd5b505afa15801561080e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108329190611aeb565b61083c9042611ad3565b602080830151015260405162498bff60e81b81526001600160a01b0383169063498bff00906108a99060040160208082526021908201527f636f6e737469747574696f6e2e766f74696e672e656d675155706461746551526040820152604d60f81b606082015260800190565b60206040518083038186803b1580156108c157600080fd5b505afa1580156108d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f99190611aeb565b60208201516080015260405162498bff60e81b81526001600160a01b0383169063498bff00906109689060040160208082526022908201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60408201526120a560f11b606082015260800190565b60206040518083038186803b15801561098057600080fd5b505afa158015610994573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b89190611aeb565b602082015160a0015260405162498bff60e81b81526001600160a01b0383169063498bff0090610a1c906004016020808252601f908201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e5000604082015260600190565b60206040518083038186803b158015610a3457600080fd5b505afa158015610a48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6c9190611aeb565b6020808301516060019190915260035460009081526001825260409020825180518493610a9d9284929101906116c9565b506020828101518051600184015580820151600284015560408082015160038086019190915560608084015160048701556080840151600587015560a0840151600687015560c0840151600787015560e0909301516008860155818601518051600987015593840151600a86015592810151600b850155930151600c909201805460ff1916921515929092179091555490517fa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e591610b5c918490611bc7565b60405180910390a160038054906000610b7483611be8565b90915550949350505050565b60008181526001602052604080822081516080810190925280548392919082908290610bab90611a43565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd790611a43565b8015610c245780601f10610bf957610100808354040283529160200191610c24565b820191906000526020600020905b815481529060010190602001808311610c0757829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015481830152600b8601548185015292840192909252600c9093015460ff1615159290910191909152818101510151909150610ce25750600092915050565b806060015115610cf55750600592915050565b806020015160200151421015610d0e5750600192915050565b604081015160208101519051600091610d2691611ad3565b90506000610d3261133f565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b158015610d6a57600080fd5b505afa158015610d7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da29190611aeb565b90506000610db083836113f1565b9050836020015160800151811015610dce5750600295945050505050565b6000610de2856040015160000151856113f1565b9050846020015160a001518111610e00575060029695505050505050565b6020808601516060810151910151610e189190611ad3565b421115610e2c575060079695505050505050565b5060049695505050505050565b610e4161133f565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401610e6c9190611b04565b60206040518083038186803b158015610e8457600080fd5b505afa158015610e98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebc9190611b18565b610ed85760405162461bcd60e51b81526004016103e590611b3a565b806000610ee482610b80565b6007811115610ef557610ef5611978565b1415610f135760405162461bcd60e51b81526004016103e590611a78565b610f1e82600261142a565b5050565b610f2a61133f565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401610f559190611b04565b60206040518083038186803b158015610f6d57600080fd5b505afa158015610f81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa59190611b18565b610fc15760405162461bcd60e51b81526004016103e590611b3a565b806000610fcd82610b80565b6007811115610fde57610fde611978565b1415610ffc5760405162461bcd60e51b81526004016103e590611a78565b610f1e82600161142a565b600054610100900460ff1680611020575060005460ff16155b6110835760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e5565b600054610100900460ff161580156110a5576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610f1e576000805461ff00191690555050565b6110e4611639565b8160006110f082610b80565b600781111561110157611101611978565b141561111f5760405162461bcd60e51b81526004016103e590611a78565b6000838152600160205260409081902081516080810190925280548290829061114790611a43565b80601f016020809104026020016040519081016040528092919081815260200182805461117390611a43565b80156111c05780601f10611195576101008083540402835291602001916111c0565b820191906000526020600020905b8154815290600101906020018083116111a357829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602080830191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff1615159101529392505050565b600461127882610b80565b600781111561128957611289611978565b146112f65760405162461bcd60e51b815260206004820152603760248201527f5b5145432d3033303030335d2d50726f706f73616c206d75737420626520504160448201527629a9a2a2103132b337b9329032bc31b2b1baba34b7339760491b60648201526084016103e5565b6000818152600160208190526040808320600c01805460ff19169092179091555182917f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f91a250565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161139c91600401611b97565b60206040518083038186803b1580156113b457600080fd5b505afa1580156113c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ec9190611baa565b905090565b60008161140057506000611424565b816114176b033b2e3c9fd0803ce800000085611c03565b6114219190611c22565b90505b92915050565b600161143583610b80565b600781111561144657611446611978565b146114b65760405162461bcd60e51b815260206004820152603a60248201527f5b5145432d3033303030305d2d566f74696e67206973206f6e6c7920706f737360448201527934b136329037b7102822a72224a72390383937b837b9b0b6399760311b60648201526084016103e5565b600082815260026020908152604080832033845290915290205460ff16156115445760405162461bcd60e51b815260206004820152603b60248201527f5b5145432d3033303030325d2d5468652063616c6c65722068617320616c726560448201527a30b23c903b37ba32b2103337b9103a343290383937b837b9b0b61760291b60648201526084016103e5565b600181600281111561155857611558611978565b141561158557600082815260016020526040812060090180549161157b83611be8565b91905055506115a8565b6000828152600160205260408120600a018054916115a283611be8565b91905055505b600082815260026020908152604080832033845290915290819020805460ff191660011790555182907fe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5906115fe908490611c44565b60405180910390a25050565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b60405180608001604052806060815260200161169360405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b81526020016116bc60405180606001604052806000815260200160008152602001600081525090565b8152600060209091015290565b8280546116d590611a43565b90600052602060002090601f0160209004810192826116f7576000855561173d565b82601f1061171057805160ff191683800117855561173d565b8280016001018555821561173d579182015b8281111561173d578251825591602001919060010190611722565b5061174992915061174d565b5090565b5b80821115611749576000815560010161174e565b60006020828403121561177457600080fd5b5035919050565b6000815180845260005b818110156117a157602081850181015186830182015201611785565b818111156117b3576000602083870101525b50601f01601f19169290920160200192915050565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b60006101a080835261182a8184018861177b565b91505061183a60208301866117c8565b835161012083015260208401516101408301526040909301516101608201529015156101809091015292915050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561189157600080fd5b813567ffffffffffffffff808211156118a957600080fd5b818401915084601f8301126118bd57600080fd5b8135818111156118cf576118cf611869565b604051601f8201601f19908116603f011681019083821181831017156118f7576118f7611869565b8160405282815287602084870101111561191057600080fd5b826020860160208301376000928101602001929092525095945050505050565b6001600160a01b038116811461194557600080fd5b50565b6000806040838503121561195b57600080fd5b82359150602083013561196d81611930565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60208101600883106119a2576119a2611978565b91905290565b6000602082840312156119ba57600080fd5b81356119c581611930565b9392505050565b60006101a082518185526119e28286018261177b565b91505060208301516119f760208601826117c8565b50604083810151805161012087015260208101516101408701520151610160850152606090920151151561018090930192909252919050565b60208152600061142160208301846119cc565b600181811c90821680611a5757607f821691505b6020821081141561062357634e487b7160e01b600052602260045260246000fd5b60208082526025908201527f5b5145432d3033303030355d2d50726f706f73616c20646f6573206e6f7420656040820152643c34b9ba1760d91b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611ae657611ae6611abd565b500190565b600060208284031215611afd57600080fd5b5051919050565b6001600160a01b0391909116815260200190565b600060208284031215611b2a57600080fd5b815180151581146119c557600080fd5b6020808252603d908201527f5b5145432d3033303030345d2d5065726d697373696f6e2064656e696564202d60408201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e000000606082015260800190565b602081526000611421602083018461177b565b600060208284031215611bbc57600080fd5b81516119c581611930565b828152604060208201526000611be060408301846119cc565b949350505050565b6000600019821415611bfc57611bfc611abd565b5060010190565b6000816000190483118215151615611c1d57611c1d611abd565b500290565b600082611c3f57634e487b7160e01b600052601260045260246000fd5b500490565b60208101600383106119a2576119a261197856fe676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a2646970667358221220d9d285539b85e03e782bfa14ab6504061bcf75bf41d57845b6eea388583f6e2864736f6c63430008090033",
}

// EmergencyUpdateVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use EmergencyUpdateVotingMetaData.ABI instead.
var EmergencyUpdateVotingABI = EmergencyUpdateVotingMetaData.ABI

// EmergencyUpdateVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EmergencyUpdateVotingMetaData.Bin instead.
var EmergencyUpdateVotingBin = EmergencyUpdateVotingMetaData.Bin

// DeployEmergencyUpdateVoting deploys a new Ethereum contract, binding an instance of EmergencyUpdateVoting to it.
func DeployEmergencyUpdateVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmergencyUpdateVoting, error) {
	parsed, err := EmergencyUpdateVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EmergencyUpdateVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmergencyUpdateVoting{EmergencyUpdateVotingCaller: EmergencyUpdateVotingCaller{contract: contract}, EmergencyUpdateVotingTransactor: EmergencyUpdateVotingTransactor{contract: contract}, EmergencyUpdateVotingFilterer: EmergencyUpdateVotingFilterer{contract: contract}}, nil
}

// EmergencyUpdateVoting is an auto generated Go binding around an Ethereum contract.
type EmergencyUpdateVoting struct {
	EmergencyUpdateVotingCaller     // Read-only binding to the contract
	EmergencyUpdateVotingTransactor // Write-only binding to the contract
	EmergencyUpdateVotingFilterer   // Log filterer for contract events
}

// EmergencyUpdateVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmergencyUpdateVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyUpdateVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmergencyUpdateVotingSession struct {
	Contract     *EmergencyUpdateVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EmergencyUpdateVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmergencyUpdateVotingCallerSession struct {
	Contract *EmergencyUpdateVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// EmergencyUpdateVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmergencyUpdateVotingTransactorSession struct {
	Contract     *EmergencyUpdateVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// EmergencyUpdateVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmergencyUpdateVotingRaw struct {
	Contract *EmergencyUpdateVoting // Generic contract binding to access the raw methods on
}

// EmergencyUpdateVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingCallerRaw struct {
	Contract *EmergencyUpdateVotingCaller // Generic read-only contract binding to access the raw methods on
}

// EmergencyUpdateVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmergencyUpdateVotingTransactorRaw struct {
	Contract *EmergencyUpdateVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmergencyUpdateVoting creates a new instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVoting(address common.Address, backend bind.ContractBackend) (*EmergencyUpdateVoting, error) {
	contract, err := bindEmergencyUpdateVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVoting{EmergencyUpdateVotingCaller: EmergencyUpdateVotingCaller{contract: contract}, EmergencyUpdateVotingTransactor: EmergencyUpdateVotingTransactor{contract: contract}, EmergencyUpdateVotingFilterer: EmergencyUpdateVotingFilterer{contract: contract}}, nil
}

// NewEmergencyUpdateVotingCaller creates a new read-only instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingCaller(address common.Address, caller bind.ContractCaller) (*EmergencyUpdateVotingCaller, error) {
	contract, err := bindEmergencyUpdateVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingCaller{contract: contract}, nil
}

// NewEmergencyUpdateVotingTransactor creates a new write-only instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*EmergencyUpdateVotingTransactor, error) {
	contract, err := bindEmergencyUpdateVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingTransactor{contract: contract}, nil
}

// NewEmergencyUpdateVotingFilterer creates a new log filterer instance of EmergencyUpdateVoting, bound to a specific deployed contract.
func NewEmergencyUpdateVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*EmergencyUpdateVotingFilterer, error) {
	contract, err := bindEmergencyUpdateVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingFilterer{contract: contract}, nil
}

// bindEmergencyUpdateVoting binds a generic wrapper to an already deployed contract.
func bindEmergencyUpdateVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmergencyUpdateVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.EmergencyUpdateVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyUpdateVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (IVotingBaseProposal, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getProposal", _proposalId)

	if err != nil {
		return *new(IVotingBaseProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _EmergencyUpdateVoting.Contract.GetProposal(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _EmergencyUpdateVoting.Contract.GetProposal(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetProposalStats(opts *bind.CallOpts, _proposalId *big.Int) (IVotingVotingStats, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getProposalStats", _proposalId)

	if err != nil {
		return *new(IVotingVotingStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingVotingStats)).(*IVotingVotingStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _EmergencyUpdateVoting.Contract.GetProposalStats(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _EmergencyUpdateVoting.Contract.GetProposalStats(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetStatus(opts *bind.CallOpts, _proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getStatus", _proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _EmergencyUpdateVoting.Contract.GetStatus(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _EmergencyUpdateVoting.Contract.GetStatus(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetVotesAgainst(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getVotesAgainst", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesAgainst(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesAgainst(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) GetVotesFor(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "getVotesFor", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesFor(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.GetVotesFor(&_EmergencyUpdateVoting.CallOpts, _proposalId)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) ProposalCount() (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.ProposalCount(&_EmergencyUpdateVoting.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) ProposalCount() (*big.Int, error) {
	return _EmergencyUpdateVoting.Contract.ProposalCount(&_EmergencyUpdateVoting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Remark   string
		Params   IVotingVotingParams
		Counters IVotingVotingCounters
		Executed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Remark = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Params = *abi.ConvertType(out[1], new(IVotingVotingParams)).(*IVotingVotingParams)
	outstruct.Counters = *abi.ConvertType(out[2], new(IVotingVotingCounters)).(*IVotingVotingCounters)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _EmergencyUpdateVoting.Contract.Proposals(&_EmergencyUpdateVoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _EmergencyUpdateVoting.Contract.Proposals(&_EmergencyUpdateVoting.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCaller) Voted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _EmergencyUpdateVoting.contract.Call(opts, &out, "voted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _EmergencyUpdateVoting.Contract.Voted(&_EmergencyUpdateVoting.CallOpts, arg0, arg1)
}

// Voted is a free data retrieval call binding the contract method 0x5277b4ae.
//
// Solidity: function voted(uint256 , address ) view returns(bool)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerSession) Voted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _EmergencyUpdateVoting.Contract.Voted(&_EmergencyUpdateVoting.CallOpts, arg0, arg1)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) CreateProposal(opts *bind.TransactOpts, _remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "createProposal", _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.CreateProposal(&_EmergencyUpdateVoting.TransactOpts, _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.CreateProposal(&_EmergencyUpdateVoting.TransactOpts, _remark)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) Execute(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "execute", _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Execute(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Execute(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Initialize(&_EmergencyUpdateVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.Initialize(&_EmergencyUpdateVoting.TransactOpts, _registry)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) VoteAgainst(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "voteAgainst", _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteAgainst(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteAgainst(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactor) VoteFor(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.contract.Transact(opts, "voteFor", _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteFor(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_EmergencyUpdateVoting *EmergencyUpdateVotingTransactorSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _EmergencyUpdateVoting.Contract.VoteFor(&_EmergencyUpdateVoting.TransactOpts, _proposalId)
}

// EmergencyUpdateVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalCreatedIterator struct {
	Event *EmergencyUpdateVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingProposalCreated)
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
		it.Event = new(EmergencyUpdateVotingProposalCreated)
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
func (it *EmergencyUpdateVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingProposalCreated represents a ProposalCreated event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalCreated struct {
	Id       *big.Int
	Proposal IVotingBaseProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*EmergencyUpdateVotingProposalCreatedIterator, error) {

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingProposalCreatedIterator{contract: _EmergencyUpdateVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingProposalCreated)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseProposalCreated(log types.Log) (*EmergencyUpdateVotingProposalCreated, error) {
	event := new(EmergencyUpdateVotingProposalCreated)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmergencyUpdateVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalExecutedIterator struct {
	Event *EmergencyUpdateVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingProposalExecuted)
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
		it.Event = new(EmergencyUpdateVotingProposalExecuted)
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
func (it *EmergencyUpdateVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingProposalExecuted represents a ProposalExecuted event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId []*big.Int) (*EmergencyUpdateVotingProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingProposalExecutedIterator{contract: _EmergencyUpdateVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingProposalExecuted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingProposalExecuted)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseProposalExecuted(log types.Log) (*EmergencyUpdateVotingProposalExecuted, error) {
	event := new(EmergencyUpdateVotingProposalExecuted)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EmergencyUpdateVotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingUserVotedIterator struct {
	Event *EmergencyUpdateVotingUserVoted // Event containing the contract specifics and raw log

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
func (it *EmergencyUpdateVotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyUpdateVotingUserVoted)
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
		it.Event = new(EmergencyUpdateVotingUserVoted)
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
func (it *EmergencyUpdateVotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyUpdateVotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyUpdateVotingUserVoted represents a UserVoted event raised by the EmergencyUpdateVoting contract.
type EmergencyUpdateVotingUserVoted struct {
	ProposalId   *big.Int
	VotingOption uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) FilterUserVoted(opts *bind.FilterOpts, _proposalId []*big.Int) (*EmergencyUpdateVotingUserVotedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.FilterLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &EmergencyUpdateVotingUserVotedIterator{contract: _EmergencyUpdateVoting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *EmergencyUpdateVotingUserVoted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _EmergencyUpdateVoting.contract.WatchLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyUpdateVotingUserVoted)
				if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
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

// ParseUserVoted is a log parse operation binding the contract event 0xe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption)
func (_EmergencyUpdateVoting *EmergencyUpdateVotingFilterer) ParseUserVoted(log types.Log) (*EmergencyUpdateVotingUserVoted, error) {
	event := new(EmergencyUpdateVotingUserVoted)
	if err := _EmergencyUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
