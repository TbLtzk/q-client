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

// GeneralUpdateVotingMetaData contains all meta data concerning the GeneralUpdateVoting contract.
var GeneralUpdateVotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"_proposal\",\"type\":\"tuple\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"QuorumReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIVoting.VotingOption\",\"name\":\"_votingOption\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VetoOccurred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasRootVetoed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteAgainst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"veto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesAgainst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVetosNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVetosPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVetoPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotingWeightInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasAlreadyVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"canVote\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ownWeight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"votingAgent\",\"type\":\"address\"},{\"internalType\":\"enumDelegationStatus\",\"name\":\"delegationStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structBaseVotingWeightInfo\",\"name\":\"base\",\"type\":\"tuple\"}],\"internalType\":\"structIQthVoting.VotingWeightInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061237c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f65760003560e01c8063bb1d689311610092578063bb1d68931461023b578063c4d66de81461024e578063c7f758a814610261578063d40b65eb14610281578063da35c664146102a4578063dc296ae1146102ad578063e8d2e442146102eb578063f99b395414610319578063fe0d94c11461032c57600080fd5b8063013cf08b146100fb5780631d28dec714610127578063307a064f1461013c57806349c2a1a6146101915780635c622a0e146101b2578063750e443a146101d257806386a50535146101e5578063ad0ccf4d146101f8578063b6f61f6614610218575b600080fd5b61010e610109366004611d88565b61033f565b60405161011e9493929190611e3c565b60405180910390f35b61013a610135366004611d88565b610470565b005b61014f61014a366004611d88565b610717565b60405161011e9190600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b6101a461019f366004611ea5565b610938565b60405190815260200161011e565b6101c56101c0366004611d88565b610e95565b60405161011e9190611f6c565b61013a6101e0366004611d88565b611131565b61013a6101f3366004611d88565b61117b565b61020b610206366004611d88565b6111c1565b60405161011e9190611f99565b6101a4610226366004611d88565b6000908152600160205260409020600a015490565b6101a4610249366004611d88565b6112ea565b61013a61025c366004612007565b61133d565b61027461026f366004611d88565b611412565b60405161011e9190612088565b6101a461028f366004611d88565b60009081526001602052604090206009015490565b6101a460045481565b6102db6102bb36600461209b565b600260209081526000928352604080842090915290825290205460ff1681565b604051901515815260200161011e565b6102db6102f936600461209b565b600360209081526000928352604080842090915290825290205460ff1681565b6101a4610327366004611d88565b6115a3565b61013a61033a366004611d88565b6115ef565b60016020526000908152604090208054819061035a906120cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610386906120cb565b80156103d35780601f106103a8576101008083540402835291602001916103d3565b820191906000526020600020905b8154815290600101906020018083116103b657829003601f168201915b50506040805161010081018252600187015481526002870154602080830191909152600388015482840152600488015460608084019190915260058901546080840152600689015460a0840152600789015460c0840152600889015460e08401528351908101845260098901548152600a89015491810191909152600b88015492810192909252600c9096015494959490935060ff169150859050565b6104786116c1565b60405163288c314960e21b81523360048201526001600160a01b03919091169063a230c5249060240160206040518083038186803b1580156104b957600080fd5b505afa1580156104cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f19190612100565b6105685760405162461bcd60e51b815260206004820152603d60248201527f5b5145432d3032393030375d2d5065726d697373696f6e2064656e696564202d60448201527f206f6e6c7920726f6f74206e6f6465732068617665206163636573732e00000060648201526084015b60405180910390fd5b80600061057482610e95565b600781111561058557610585611f56565b14156105a35760405162461bcd60e51b815260040161055f90612122565b60036105ae83610e95565b60078111156105bf576105bf611f56565b146106305760405162461bcd60e51b815260206004820152603b60248201527f5b5145432d3032393030345d2d50726f706f73616c20737461747573206d757360448201527a3a1031329030b1b1b2b83a32b216103b32ba37903330b4b632b21760291b606482015260840161055f565b600082815260036020908152604080832033845290915290205460ff16156106cf5760405162461bcd60e51b815260206004820152604660248201527f5b5145432d3032393030355d2d5468652073656e6465722068617320616c726560448201527f616479207665746f656420746869732070726f706f73616c2c207665746f206660648201526530b4b632b21760d11b608482015260a40161055f565b60008281526003602090815260408083203384528252808320805460ff191660019081179091558584529091528120600b01805490919061070f9061217d565b909155505050565b61071f611be9565b81600061072b82610e95565b600781111561073c5761073c611f56565b141561075a5760405162461bcd60e51b815260040161055f90612122565b600083815260016020526040808220815160808101909252805482908290610781906120cb565b80601f01602080910402602001604051908101604052809291908181526020018280546107ad906120cb565b80156107fa5780601f106107cf576101008083540402835291602001916107fa565b820191906000526020600020905b8154815290600101906020018083116107dd57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff16151591015290506108aa611be9565b602080830180516080015183525160a001516040808401919091528301519081015190516000916108da91612198565b905060006108e743611773565b905060006108f583836117a3565b6020850181905260408601515190915060009061091290856117a3565b606086018190529050610924896117d2565b608086015250929550505050505b50919050565b600080546040805160608101909152602280825283926201000090046001600160a01b031691633fb902719161232560208301396040518263ffffffff1660e01b815260040161098891906121b0565b60206040518083038186803b1580156109a057600080fd5b505afa1580156109b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d891906121c3565b90506109e2611c18565b838152602081015142905260405162498bff60e81b81526001600160a01b0383169063498bff0090610a4490600401602080825260259082015260008051602061230583398151915260408201526406e737456560dc1b606082015260800190565b60206040518083038186803b158015610a5c57600080fd5b505afa158015610a70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9491906121e0565b610a9e9042612198565b602080830151015260405162498bff60e81b81526001600160a01b0383169063498bff0090610afe9060040160208082526026908201526000805160206123058339815191526040820152656e737451524d60d01b606082015260800190565b60206040518083038186803b158015610b1657600080fd5b505afa158015610b2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b4e91906121e0565b60208201516080015260405162498bff60e81b81526001600160a01b0383169063498bff0090610bb09060040160208082526027908201526000805160206123058339815191526040820152663739ba2926a0a560c91b606082015260800190565b60206040518083038186803b158015610bc857600080fd5b505afa158015610bdc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0091906121e0565b602082015160a0015260405162498bff60e81b81526001600160a01b0383169063498bff0090610c6490600401602080825260299082015260008051602061230583398151915260408201526806e7374524e56414c560bc1b606082015260800190565b60206040518083038186803b158015610c7c57600080fd5b505afa158015610c90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb491906121e0565b816020015160200151610cc79190612198565b60208201516040908101919091525162498bff60e81b81526001600160a01b0383169063498bff0090610d2e906004016020808252601f908201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e5000604082015260600190565b60206040518083038186803b158015610d4657600080fd5b505afa158015610d5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7e91906121e0565b6020808301516060019190915260045460009081526001825260409020825180518493610daf928492910190611ca8565b506020828101518051600184015580820151600284015560408082015160038501556060808301516004808701919091556080840151600587015560a0840151600687015560c0840151600787015560e0909301516008860155818601518051600987015593840151600a86015592810151600b8501559190930151600c909201805460ff191692151592909217909155905490517fa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e591610e719184906121f9565b60405180910390a160048054906000610e898361217d565b90915550949350505050565b60008181526001602052604080822081516080810190925280548392919082908290610ec0906120cb565b80601f0160208091040260200160405190810160405280929190818152602001828054610eec906120cb565b8015610f395780601f10610f0e57610100808354040283529160200191610f39565b820191906000526020600020905b815481529060010190602001808311610f1c57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015481830152600b8601548185015292840192909252600c9093015460ff1615159290910191909152818101510151909150610ff75750600092915050565b80606001511561100a5750600592915050565b8060200151602001514210156110235750600192915050565b60408101516020810151905160009161103b91612198565b9050600061104843611773565b9050600061105683836117a3565b90508360200151608001518110156110745750600295945050505050565b6000611088856040015160000151856117a3565b9050846020015160a001518110156110a7575060029695505050505050565b6110bd6002676765c793fa10079d601b1b61221a565b6110c6886117d2565b11156110d9575060029695505050505050565b8460200151604001514210156110f6575060039695505050505050565b602085015160608101516040909101516111109190612198565b421115611124575060079695505050505050565b5060049695505050505050565b80600061113d82610e95565b600781111561114e5761114e611f56565b141561116c5760405162461bcd60e51b815260040161055f90612122565b61117782600261187e565b5050565b80600061118782610e95565b600781111561119857611198611f56565b14156111b65760405162461bcd60e51b815260040161055f90612122565b61117782600161187e565b6111c9611d2c565b8160006111d582610e95565b60078111156111e6576111e6611f56565b14156112045760405162461bcd60e51b815260040161055f90612122565b600061120e611b83565b9050611218611d2c565b600085815260016020526040908190206002015490516310874b5b60e01b81526001600160a01b038416916310874b5b9161125791339160040161223c565b60806040518083038186803b15801561126f57600080fd5b505afa158015611283573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a79190612255565b60408083019182526000878152600260209081528282203383529052205460ff16151582525151158015906112db57508051155b15156020820152949350505050565b600081816112f782610e95565b600781111561130857611308611f56565b14156113265760405162461bcd60e51b815260040161055f90612122565b50506000908152600160205260409020600b015490565b600054610100900460ff1680611356575060005460ff16155b6113b95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161055f565b600054610100900460ff161580156113db576000805461ffff19166101011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015611177576000805461ff00191690555050565b61141a611c18565b81600061142682610e95565b600781111561143757611437611f56565b14156114555760405162461bcd60e51b815260040161055f90612122565b6000838152600160205260409081902081516080810190925280548290829061147d906120cb565b80601f01602080910402602001604051908101604052809291908181526020018280546114a9906120cb565b80156114f65780601f106114cb576101008083540402835291602001916114f6565b820191906000526020600020905b8154815290600101906020018083116114d957829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602080830191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff1615159101529392505050565b600081816115b082610e95565b60078111156115c1576115c1611f56565b14156115df5760405162461bcd60e51b815260040161055f90612122565b6115e8836117d2565b9392505050565b60046115fa82610e95565b600781111561160b5761160b611f56565b146116785760405162461bcd60e51b815260206004820152603760248201527f5b5145432d3032393030305d2d50726f706f73616c206d75737420626520504160448201527629a9a2a2103132b337b9329032bc31b2b1baba34b7339760491b606482015260840161055f565b6000818152600160208190526040808320600c01805460ff19169092179091555182917f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f91a250565b600080546040805180820182526014815273676f7665726e616e63652e726f6f744e6f64657360601b60208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161171e916004016121b0565b60206040518083038186803b15801561173657600080fd5b505afa15801561174a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176e91906121c3565b905090565b6000676765c793fa10079d601b1b6117936714d1120d7b160000846122ce565b61179d9190612198565b92915050565b6000816117b25750600061179d565b816117c8676765c793fa10079d601b1b856122ce565b6115e8919061221a565b60006117dc6116c1565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b15801561181457600080fd5b505afa158015611828573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184c91906121e0565b676765c793fa10079d601b1b6000848152600160205260409020600b015461187491906122ce565b61179d919061221a565b600161188983610e95565b600781111561189a5761189a611f56565b1461190a5760405162461bcd60e51b815260206004820152603a60248201527f5b5145432d3032393030315d2d566f74696e67206973206f6e6c7920706f737360448201527934b136329037b7102822a72224a72390383937b837b9b0b6399760311b606482015260840161055f565b600082815260026020908152604080832033845290915290205460ff16156119985760405162461bcd60e51b815260206004820152603b60248201527f5b5145432d3032393030325d2d5468652063616c6c65722068617320616c726560448201527a30b23c903b37ba32b2103337b9103a343290383937b837b9b0b61760291b606482015260840161055f565b60008281526002602090815260408083203384528252808320805460ff191660019081179091558584529091528120906119d0611b83565b6002830154604051630f7f3f2360e31b81529192506000916001600160a01b03841691637bf9f91891611a0791339160040161223c565b602060405180830381600087803b158015611a2157600080fd5b505af1158015611a35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5991906121e0565b905060008111611ae95760405162461bcd60e51b815260206004820152604f60248201527f5b5145432d3032393030335d2d54686520746f74616c20766f74696e6720776560448201527f69676874206d7573742062652067726561746572207468616e207a65726f2c2060648201526e3330b4b632b2103a37903b37ba329760891b608482015260a40161055f565b6001846002811115611afd57611afd611f56565b1415611b255780836009016000016000828254611b1a9190612198565b90915550611b429050565b80836009016001016000828254611b3c9190612198565b90915550505b847f5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d48583604051611b749291906122ed565b60405180910390a25050505050565b60008054604080518082018252601c81527f676f7665726e616e63652e766f74696e6757656967687450726f78790000000060208201529051633fb9027160e01b8152620100009092046001600160a01b031691633fb902719161171e916004016121b0565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b604051806080016040528060608152602001611c7260405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b8152602001611c9b60405180606001604052806000815260200160008152602001600081525090565b8152600060209091015290565b828054611cb4906120cb565b90600052602060002090601f016020900481019282611cd65760008555611d1c565b82601f10611cef57805160ff1916838001178555611d1c565b82800160010185558215611d1c579182015b82811115611d1c578251825591602001919060010190611d01565b50611d28929150611d73565b5090565b6040805160608101825260008082526020820152908101611d6e6040805160808101825260008082526020820181905290918201908152602001600081525090565b905290565b5b80821115611d285760008155600101611d74565b600060208284031215611d9a57600080fd5b5035919050565b6000815180845260005b81811015611dc757602081850181015186830182015201611dab565b81811115611dd9576000602083870101525b50601f01601f19169290920160200192915050565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b60006101a0808352611e5081840188611da1565b915050611e606020830186611dee565b835161012083015260208401516101408301526040909301516101608201529015156101809091015292915050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215611eb757600080fd5b813567ffffffffffffffff80821115611ecf57600080fd5b818401915084601f830112611ee357600080fd5b813581811115611ef557611ef5611e8f565b604051601f8201601f19908116603f01168101908382118183101715611f1d57611f1d611e8f565b81604052828152876020848701011115611f3657600080fd5b826020860160208301376000928101602001929092525095945050505050565b634e487b7160e01b600052602160045260246000fd5b6020810160088310611f8057611f80611f56565b91905290565b60038110611f9657611f96611f56565b50565b815115158152602080830151151581830152604080840151805182850152918201516001600160a01b0316606084015281015160c083019190611fdb81611f86565b60808401526060015160a090920191909152919050565b6001600160a01b0381168114611f9657600080fd5b60006020828403121561201957600080fd5b81356115e881611ff2565b60006101a0825181855261203a82860182611da1565b915050602083015161204f6020860182611dee565b50604083810151805161012087015260208101516101408701520151610160850152606090920151151561018090930192909252919050565b6020815260006115e86020830184612024565b600080604083850312156120ae57600080fd5b8235915060208301356120c081611ff2565b809150509250929050565b600181811c908216806120df57607f821691505b6020821081141561093257634e487b7160e01b600052602260045260246000fd5b60006020828403121561211257600080fd5b815180151581146115e857600080fd5b60208082526025908201527f5b5145432d3032393030365d2d50726f706f73616c20646f6573206e6f7420656040820152643c34b9ba1760d91b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060001982141561219157612191612167565b5060010190565b600082198211156121ab576121ab612167565b500190565b6020815260006115e86020830184611da1565b6000602082840312156121d557600080fd5b81516115e881611ff2565b6000602082840312156121f257600080fd5b5051919050565b8281526040602082015260006122126040830184612024565b949350505050565b60008261223757634e487b7160e01b600052601260045260246000fd5b500490565b6001600160a01b03929092168252602082015260400190565b60006080828403121561226757600080fd5b6040516080810181811067ffffffffffffffff8211171561228a5761228a611e8f565b60405282518152602083015161229f81611ff2565b60208201526040830151600381106122b657600080fd5b60408201526060928301519281019290925250919050565b60008160001904831182151516156122e8576122e8612167565b500290565b604081016122fa84611f86565b928152602001529056fe636f6e737469747574696f6e2e766f74696e672e6368616e6765516e6f74436f676f7665726e616e63652e636f6e737469747574696f6e2e706172616d6574657273a26469706673582212200ad31daacbb0f7a29e079c50b9e6d7d8db192fccb9a2a105d540aa5193cc2ada64736f6c63430008090033",
}

// GeneralUpdateVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use GeneralUpdateVotingMetaData.ABI instead.
var GeneralUpdateVotingABI = GeneralUpdateVotingMetaData.ABI

// GeneralUpdateVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GeneralUpdateVotingMetaData.Bin instead.
var GeneralUpdateVotingBin = GeneralUpdateVotingMetaData.Bin

// DeployGeneralUpdateVoting deploys a new Ethereum contract, binding an instance of GeneralUpdateVoting to it.
func DeployGeneralUpdateVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GeneralUpdateVoting, error) {
	parsed, err := GeneralUpdateVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GeneralUpdateVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GeneralUpdateVoting{GeneralUpdateVotingCaller: GeneralUpdateVotingCaller{contract: contract}, GeneralUpdateVotingTransactor: GeneralUpdateVotingTransactor{contract: contract}, GeneralUpdateVotingFilterer: GeneralUpdateVotingFilterer{contract: contract}}, nil
}

// GeneralUpdateVoting is an auto generated Go binding around an Ethereum contract.
type GeneralUpdateVoting struct {
	GeneralUpdateVotingCaller     // Read-only binding to the contract
	GeneralUpdateVotingTransactor // Write-only binding to the contract
	GeneralUpdateVotingFilterer   // Log filterer for contract events
}

// GeneralUpdateVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneralUpdateVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneralUpdateVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneralUpdateVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralUpdateVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneralUpdateVotingSession struct {
	Contract     *GeneralUpdateVoting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GeneralUpdateVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneralUpdateVotingCallerSession struct {
	Contract *GeneralUpdateVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// GeneralUpdateVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneralUpdateVotingTransactorSession struct {
	Contract     *GeneralUpdateVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// GeneralUpdateVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneralUpdateVotingRaw struct {
	Contract *GeneralUpdateVoting // Generic contract binding to access the raw methods on
}

// GeneralUpdateVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneralUpdateVotingCallerRaw struct {
	Contract *GeneralUpdateVotingCaller // Generic read-only contract binding to access the raw methods on
}

// GeneralUpdateVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneralUpdateVotingTransactorRaw struct {
	Contract *GeneralUpdateVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneralUpdateVoting creates a new instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVoting(address common.Address, backend bind.ContractBackend) (*GeneralUpdateVoting, error) {
	contract, err := bindGeneralUpdateVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVoting{GeneralUpdateVotingCaller: GeneralUpdateVotingCaller{contract: contract}, GeneralUpdateVotingTransactor: GeneralUpdateVotingTransactor{contract: contract}, GeneralUpdateVotingFilterer: GeneralUpdateVotingFilterer{contract: contract}}, nil
}

// NewGeneralUpdateVotingCaller creates a new read-only instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingCaller(address common.Address, caller bind.ContractCaller) (*GeneralUpdateVotingCaller, error) {
	contract, err := bindGeneralUpdateVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingCaller{contract: contract}, nil
}

// NewGeneralUpdateVotingTransactor creates a new write-only instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneralUpdateVotingTransactor, error) {
	contract, err := bindGeneralUpdateVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingTransactor{contract: contract}, nil
}

// NewGeneralUpdateVotingFilterer creates a new log filterer instance of GeneralUpdateVoting, bound to a specific deployed contract.
func NewGeneralUpdateVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneralUpdateVotingFilterer, error) {
	contract, err := bindGeneralUpdateVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingFilterer{contract: contract}, nil
}

// bindGeneralUpdateVoting binds a generic wrapper to an already deployed contract.
func bindGeneralUpdateVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneralUpdateVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralUpdateVoting *GeneralUpdateVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.GeneralUpdateVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralUpdateVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (IVotingBaseProposal, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getProposal", _proposalId)

	if err != nil {
		return *new(IVotingBaseProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingBaseProposal)).(*IVotingBaseProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _GeneralUpdateVoting.Contract.GetProposal(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns((string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetProposal(_proposalId *big.Int) (IVotingBaseProposal, error) {
	return _GeneralUpdateVoting.Contract.GetProposal(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetProposalStats(opts *bind.CallOpts, _proposalId *big.Int) (IVotingVotingStats, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getProposalStats", _proposalId)

	if err != nil {
		return *new(IVotingVotingStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IVotingVotingStats)).(*IVotingVotingStats)

	return out0, err

}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _GeneralUpdateVoting.Contract.GetProposalStats(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetProposalStats is a free data retrieval call binding the contract method 0x307a064f.
//
// Solidity: function getProposalStats(uint256 _proposalId) view returns((uint256,uint256,uint256,uint256,uint256))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetProposalStats(_proposalId *big.Int) (IVotingVotingStats, error) {
	return _GeneralUpdateVoting.Contract.GetProposalStats(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetStatus(opts *bind.CallOpts, _proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getStatus", _proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _GeneralUpdateVoting.Contract.GetStatus(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _proposalId) view returns(uint8)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetStatus(_proposalId *big.Int) (uint8, error) {
	return _GeneralUpdateVoting.Contract.GetStatus(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVetosNumber(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVetosNumber", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosNumber(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosNumber is a free data retrieval call binding the contract method 0xbb1d6893.
//
// Solidity: function getVetosNumber(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVetosNumber(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosNumber(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVetosPercentage(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVetosPercentage", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosPercentage(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVetosPercentage is a free data retrieval call binding the contract method 0xf99b3954.
//
// Solidity: function getVetosPercentage(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVetosPercentage(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVetosPercentage(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotesAgainst(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotesAgainst", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesAgainst(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesAgainst is a free data retrieval call binding the contract method 0xb6f61f66.
//
// Solidity: function getVotesAgainst(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotesAgainst(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesAgainst(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotesFor(opts *bind.CallOpts, _proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotesFor", _proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesFor(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotesFor is a free data retrieval call binding the contract method 0xd40b65eb.
//
// Solidity: function getVotesFor(uint256 _proposalId) view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotesFor(_proposalId *big.Int) (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.GetVotesFor(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) GetVotingWeightInfo(opts *bind.CallOpts, _proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "getVotingWeightInfo", _proposalId)

	if err != nil {
		return *new(IQthVotingVotingWeightInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IQthVotingVotingWeightInfo)).(*IQthVotingVotingWeightInfo)

	return out0, err

}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _GeneralUpdateVoting.Contract.GetVotingWeightInfo(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// GetVotingWeightInfo is a free data retrieval call binding the contract method 0xad0ccf4d.
//
// Solidity: function getVotingWeightInfo(uint256 _proposalId) view returns((bool,bool,(uint256,address,uint8,uint256)))
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) GetVotingWeightInfo(_proposalId *big.Int) (IQthVotingVotingWeightInfo, error) {
	return _GeneralUpdateVoting.Contract.GetVotingWeightInfo(&_GeneralUpdateVoting.CallOpts, _proposalId)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) HasRootVetoed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "hasRootVetoed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasRootVetoed(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasRootVetoed is a free data retrieval call binding the contract method 0xe8d2e442.
//
// Solidity: function hasRootVetoed(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) HasRootVetoed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasRootVetoed(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) HasUserVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "hasUserVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasUserVoted(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// HasUserVoted is a free data retrieval call binding the contract method 0xdc296ae1.
//
// Solidity: function hasUserVoted(uint256 , address ) view returns(bool)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) HasUserVoted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GeneralUpdateVoting.Contract.HasUserVoted(&_GeneralUpdateVoting.CallOpts, arg0, arg1)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) ProposalCount() (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.ProposalCount(&_GeneralUpdateVoting.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) ProposalCount() (*big.Int, error) {
	return _GeneralUpdateVoting.Contract.ProposalCount(&_GeneralUpdateVoting.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_GeneralUpdateVoting *GeneralUpdateVotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	var out []interface{}
	err := _GeneralUpdateVoting.contract.Call(opts, &out, "proposals", arg0)

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
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _GeneralUpdateVoting.Contract.Proposals(&_GeneralUpdateVoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string remark, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) params, (uint256,uint256,uint256) counters, bool executed)
func (_GeneralUpdateVoting *GeneralUpdateVotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Remark   string
	Params   IVotingVotingParams
	Counters IVotingVotingCounters
	Executed bool
}, error) {
	return _GeneralUpdateVoting.Contract.Proposals(&_GeneralUpdateVoting.CallOpts, arg0)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) CreateProposal(opts *bind.TransactOpts, _remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "createProposal", _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.CreateProposal(&_GeneralUpdateVoting.TransactOpts, _remark)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _remark) returns(uint256)
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) CreateProposal(_remark string) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.CreateProposal(&_GeneralUpdateVoting.TransactOpts, _remark)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Execute(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "execute", _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Execute(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Execute(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Execute(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Initialize(&_GeneralUpdateVoting.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Initialize(&_GeneralUpdateVoting.TransactOpts, _registry)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) Veto(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "veto", _proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) Veto(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Veto(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) Veto(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.Veto(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) VoteAgainst(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "voteAgainst", _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteAgainst(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteAgainst is a paid mutator transaction binding the contract method 0x750e443a.
//
// Solidity: function voteAgainst(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) VoteAgainst(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteAgainst(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactor) VoteFor(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.contract.Transact(opts, "voteFor", _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteFor(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// VoteFor is a paid mutator transaction binding the contract method 0x86a50535.
//
// Solidity: function voteFor(uint256 _proposalId) returns()
func (_GeneralUpdateVoting *GeneralUpdateVotingTransactorSession) VoteFor(_proposalId *big.Int) (*types.Transaction, error) {
	return _GeneralUpdateVoting.Contract.VoteFor(&_GeneralUpdateVoting.TransactOpts, _proposalId)
}

// GeneralUpdateVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalCreatedIterator struct {
	Event *GeneralUpdateVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingProposalCreated)
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
		it.Event = new(GeneralUpdateVotingProposalCreated)
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
func (it *GeneralUpdateVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingProposalCreated represents a ProposalCreated event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalCreated struct {
	Id       *big.Int
	Proposal IVotingBaseProposal
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GeneralUpdateVotingProposalCreatedIterator, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingProposalCreatedIterator{contract: _GeneralUpdateVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e5.
//
// Solidity: event ProposalCreated(uint256 _id, (string,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),(uint256,uint256,uint256),bool) _proposal)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingProposalCreated)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseProposalCreated(log types.Log) (*GeneralUpdateVotingProposalCreated, error) {
	event := new(GeneralUpdateVotingProposalCreated)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalExecutedIterator struct {
	Event *GeneralUpdateVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingProposalExecuted)
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
		it.Event = new(GeneralUpdateVotingProposalExecuted)
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
func (it *GeneralUpdateVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingProposalExecuted represents a ProposalExecuted event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts, _proposalId []*big.Int) (*GeneralUpdateVotingProposalExecutedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingProposalExecutedIterator{contract: _GeneralUpdateVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed _proposalId)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingProposalExecuted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "ProposalExecuted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingProposalExecuted)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseProposalExecuted(log types.Log) (*GeneralUpdateVotingProposalExecuted, error) {
	event := new(GeneralUpdateVotingProposalExecuted)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingQuorumReachedIterator is returned from FilterQuorumReached and is used to iterate over the raw logs and unpacked data for QuorumReached events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingQuorumReachedIterator struct {
	Event *GeneralUpdateVotingQuorumReached // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingQuorumReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingQuorumReached)
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
		it.Event = new(GeneralUpdateVotingQuorumReached)
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
func (it *GeneralUpdateVotingQuorumReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingQuorumReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingQuorumReached represents a QuorumReached event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingQuorumReached struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumReached is a free log retrieval operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterQuorumReached(opts *bind.FilterOpts) (*GeneralUpdateVotingQuorumReachedIterator, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingQuorumReachedIterator{contract: _GeneralUpdateVoting.contract, event: "QuorumReached", logs: logs, sub: sub}, nil
}

// WatchQuorumReached is a free log subscription operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchQuorumReached(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingQuorumReached) (event.Subscription, error) {

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "QuorumReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingQuorumReached)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
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

// ParseQuorumReached is a log parse operation binding the contract event 0x878536ebf930768ad5274a079ba36028b128aeca4b9212fece414176c39e30f7.
//
// Solidity: event QuorumReached(uint256 id)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseQuorumReached(log types.Log) (*GeneralUpdateVotingQuorumReached, error) {
	event := new(GeneralUpdateVotingQuorumReached)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "QuorumReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingUserVotedIterator is returned from FilterUserVoted and is used to iterate over the raw logs and unpacked data for UserVoted events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingUserVotedIterator struct {
	Event *GeneralUpdateVotingUserVoted // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingUserVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingUserVoted)
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
		it.Event = new(GeneralUpdateVotingUserVoted)
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
func (it *GeneralUpdateVotingUserVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingUserVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingUserVoted represents a UserVoted event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingUserVoted struct {
	ProposalId   *big.Int
	VotingOption uint8
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserVoted is a free log retrieval operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterUserVoted(opts *bind.FilterOpts, _proposalId []*big.Int) (*GeneralUpdateVotingUserVotedIterator, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingUserVotedIterator{contract: _GeneralUpdateVoting.contract, event: "UserVoted", logs: logs, sub: sub}, nil
}

// WatchUserVoted is a free log subscription operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchUserVoted(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingUserVoted, _proposalId []*big.Int) (event.Subscription, error) {

	var _proposalIdRule []interface{}
	for _, _proposalIdItem := range _proposalId {
		_proposalIdRule = append(_proposalIdRule, _proposalIdItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "UserVoted", _proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingUserVoted)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
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

// ParseUserVoted is a log parse operation binding the contract event 0x5ac937fb2a69c6ddee38a23a1b04bbe8a7edb77cdc9bbfe2f9e26dd5a53166d4.
//
// Solidity: event UserVoted(uint256 indexed _proposalId, uint8 _votingOption, uint256 _amount)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseUserVoted(log types.Log) (*GeneralUpdateVotingUserVoted, error) {
	event := new(GeneralUpdateVotingUserVoted)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "UserVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralUpdateVotingVetoOccurredIterator is returned from FilterVetoOccurred and is used to iterate over the raw logs and unpacked data for VetoOccurred events raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingVetoOccurredIterator struct {
	Event *GeneralUpdateVotingVetoOccurred // Event containing the contract specifics and raw log

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
func (it *GeneralUpdateVotingVetoOccurredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralUpdateVotingVetoOccurred)
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
		it.Event = new(GeneralUpdateVotingVetoOccurred)
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
func (it *GeneralUpdateVotingVetoOccurredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralUpdateVotingVetoOccurredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralUpdateVotingVetoOccurred represents a VetoOccurred event raised by the GeneralUpdateVoting contract.
type GeneralUpdateVotingVetoOccurred struct {
	Id     *big.Int
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVetoOccurred is a free log retrieval operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) FilterVetoOccurred(opts *bind.FilterOpts, id []*big.Int, sender []common.Address) (*GeneralUpdateVotingVetoOccurredIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.FilterLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GeneralUpdateVotingVetoOccurredIterator{contract: _GeneralUpdateVoting.contract, event: "VetoOccurred", logs: logs, sub: sub}, nil
}

// WatchVetoOccurred is a free log subscription operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) WatchVetoOccurred(opts *bind.WatchOpts, sink chan<- *GeneralUpdateVotingVetoOccurred, id []*big.Int, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GeneralUpdateVoting.contract.WatchLogs(opts, "VetoOccurred", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralUpdateVotingVetoOccurred)
				if err := _GeneralUpdateVoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
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

// ParseVetoOccurred is a log parse operation binding the contract event 0x11e347c8ff2734bf22b4aaead8d3a24eb006d62ee60ab6bbb7adf2827a8e2204.
//
// Solidity: event VetoOccurred(uint256 indexed id, address indexed sender)
func (_GeneralUpdateVoting *GeneralUpdateVotingFilterer) ParseVetoOccurred(log types.Log) (*GeneralUpdateVotingVetoOccurred, error) {
	event := new(GeneralUpdateVotingVetoOccurred)
	if err := _GeneralUpdateVoting.contract.UnpackLog(event, "VetoOccurred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
