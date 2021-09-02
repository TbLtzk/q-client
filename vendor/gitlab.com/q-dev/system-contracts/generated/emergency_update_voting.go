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

// EmergencyUpdateVotingABI is the input ABI used to generate the binding from.
const EmergencyUpdateVotingABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"_proposal\",\"type\":\"tuple\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIVoting.VotingOption\",\"name\":\"_votingOption\",\"type\":\"uint8\"}],\"name\":\"UserVoted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_remark\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVotesAgainst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteAgainst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumIVoting.ProposalStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"remark\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetoEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalExecutionP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredSQuorum\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"weightFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weightAgainst\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vetosCount\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingCounters\",\"name\":\"counters\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"internalType\":\"structIVoting.BaseProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentQuorum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentMajority\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVetoPercentage\",\"type\":\"uint256\"}],\"internalType\":\"structIVoting.VotingStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EmergencyUpdateVotingBin is the compiled bytecode used for deploying new contracts.
var EmergencyUpdateVotingBin = "0x608060405234801561001057600080fd5b50612002806100206000396000f3fe608060405234801561001057600080fd5b50600436106100af5760003560e01c8063013cf08b146100b4578063307a064f146100e057806349c2a1a6146101005780635277b4ae146101205780635c622a0e14610140578063750e443a1461016057806386a5053514610175578063b6f61f6614610188578063c4d66de81461019b578063c7f758a8146101ae578063d40b65eb146101ce578063da35c664146101e1578063fe0d94c1146101e9575b600080fd5b6100c76100c2366004611997565b6101fc565b6040516100d79493929190611b40565b60405180910390f35b6100f36100ee366004611997565b61032b565b6040516100d79190611f01565b61011361010e366004611902565b610623565b6040516100d79190611f3b565b61013361012e3660046119c7565b610ae8565b6040516100d79190611b13565b61015361014e366004611997565b610b08565b6040516100d79190611b1e565b61017361016e366004611997565b610e3e565b005b610173610183366004611997565b610f9c565b610113610196366004611997565b6110f6565b6101736101a93660046118aa565b61110b565b6101c16101bc366004611997565b6111cf565b6040516100d79190611eee565b6101136101dc366004611997565b611359565b61011361136e565b6101736101f7366004611997565b611374565b60016020818152600092835260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835292839183018282801561028e5780601f106102635761010080835404028352916020019161028e565b820191906000526020600020905b81548152906001019060200180831161027157829003601f168201915b50506040805161010081018252600187015481526002870154602080830191909152600388015482840152600488015460608084019190915260058901546080840152600689015460a0840152600789015460c0840152600889015460e08401528351908101845260098901548152600a89015491810191909152600b88015492810192909252600c9096015494959490935060ff169150859050565b610333611740565b81600061033f82610b08565b600781111561034a57fe5b14156103715760405162461bcd60e51b815260040161036890611c31565b60405180910390fd5b60008381526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a0908101909252608084018381529092849284919084018282801561040c5780601f106103e15761010080835404028352916020019161040c565b820191906000526020600020905b8154815290600101906020018083116103ef57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff16151591015290506104bc611740565b602080830180516080015183525160a001516040808401919091528301519081015190516000916104ed91906113f0565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb902719061052690600401611ec0565b60206040518083038186803b15801561053e57600080fd5b505afa158015610552573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057691906118c6565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b1580156105ae57600080fd5b505afa1580156105c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e691906119af565b905060006105f48383611451565b602085018190526040860151519091506000906106119085611451565b60608601525092979650505050505050565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061065790600401611ec0565b60206040518083038186803b15801561066f57600080fd5b505afa158015610683573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a791906118c6565b6001600160a01b031663a230c524336040518263ffffffff1660e01b81526004016106d29190611aff565b60206040518083038186803b1580156106ea57600080fd5b505afa1580156106fe573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072291906118e2565b61073e5760405162461bcd60e51b815260040161036890611d83565b60008054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb902719061077290600401611df7565b60206040518083038186803b15801561078a57600080fd5b505afa15801561079e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c291906118c6565b90506107cc61176f565b838152602081015142905260405162498bff60e81b81526001600160a01b0383169063498bff009061080090600401611d4e565b60206040518083038186803b15801561081857600080fd5b505afa15801561082c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085091906119af565b602080830151429290920191015260405162498bff60e81b81526001600160a01b0383169063498bff009061088790600401611ccb565b60206040518083038186803b15801561089f57600080fd5b505afa1580156108b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d791906119af565b60208201516080015260405162498bff60e81b81526001600160a01b0383169063498bff009061090990600401611d0c565b60206040518083038186803b15801561092157600080fd5b505afa158015610935573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095991906119af565b602082015160a0015260405162498bff60e81b81526001600160a01b0383169063498bff009061098b90600401611e89565b60206040518083038186803b1580156109a357600080fd5b505afa1580156109b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109db91906119af565b6020808301516060019190915260035460009081526001825260409020825180518493610a0c9284929101906117a3565b506020828101518051600184015580820151600284015560408082015160038086019190915560608084015160048701556080840151600587015560a0840151600687015560c0840151600787015560e0909301516008860155818601518051600987015593840151600a86015592810151600b850155930151600c909201805460ff1916921515929092179091555490517fa188b3e35b494a3dcb0a91f196c99377a74b06350898477006ed845cf90104e591610acb918490611f44565b60405180910390a15050600380546001810190915590505b919050565b600260209081526000928352604080842090915290825290205460ff1681565b60008181526001602081815260408084208151815460029581161561010002600019011694909404601f8101849004909302840160a090810190925260808401838152859493919284928491840182828015610ba55780601f10610b7a57610100808354040283529160200191610ba5565b820191906000526020600020905b815481529060010190602001808311610b8857829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602082810191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015481830152600b8601548185015292840192909252600c9093015460ff1615159290910191909152818101510151909150610c64576000915050610ae3565b806060015115610c78576005915050610ae3565b806020015160200151421015610c92576001915050610ae3565b604081015160208101519051600091610cab91906113f0565b60008054604051633fb9027160e01b81529293509091620100009091046001600160a01b031690633fb9027190610ce490600401611ec0565b60206040518083038186803b158015610cfc57600080fd5b505afa158015610d10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3491906118c6565b6001600160a01b031663de8fa4316040518163ffffffff1660e01b815260040160206040518083038186803b158015610d6c57600080fd5b505afa158015610d80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da491906119af565b90506000610db28383611451565b9050836020015160800151811015610dd1576002945050505050610ae3565b6000610de585604001516000015185611451565b9050846020015160a00151811015610e0557600295505050505050610ae3565b6020808601516060810151910151610e1c916113f0565b421115610e3157600795505050505050610ae3565b5060049695505050505050565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610e7190600401611ec0565b60206040518083038186803b158015610e8957600080fd5b505afa158015610e9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec191906118c6565b6001600160a01b031663a230c524336040518263ffffffff1660e01b8152600401610eec9190611aff565b60206040518083038186803b158015610f0457600080fd5b505afa158015610f18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3c91906118e2565b610f585760405162461bcd60e51b815260040161036890611d83565b806000610f6482610b08565b6007811115610f6f57fe5b1415610f8d5760405162461bcd60e51b815260040161036890611c31565b610f9882600261147c565b5050565b600054604051633fb9027160e01b8152620100009091046001600160a01b031690633fb9027190610fcf90600401611ec0565b60206040518083038186803b158015610fe757600080fd5b505afa158015610ffb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101f91906118c6565b6001600160a01b031663a230c524336040518263ffffffff1660e01b815260040161104a9190611aff565b60206040518083038186803b15801561106257600080fd5b505afa158015611076573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109a91906118e2565b6110b65760405162461bcd60e51b815260040161036890611d83565b8060006110c282610b08565b60078111156110cd57fe5b14156110eb5760405162461bcd60e51b815260040161036890611c31565b610f9882600161147c565b6000908152600160205260409020600a015490565b600054610100900460ff168061112457506111246115f5565b80611132575060005460ff16155b61116d5760405162461bcd60e51b815260040180806020018281038252602e815260200180611f7e602e913960400191505060405180910390fd5b600054610100900460ff16158015611198576000805460ff1961ff0019909116610100171660011790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610f98576000805461ff00191690555050565b6111d761176f565b8160006111e382610b08565b60078111156111ee57fe5b141561120c5760405162461bcd60e51b815260040161036890611c31565b6000838152600160208181526040928390208351815460029481161561010002600019011693909304601f8101839004909202830160a09081019094526080830182815292939092849290918491908401828280156112ac5780601f10611281576101008083540402835291602001916112ac565b820191906000526020600020905b81548152906001019060200180831161128f57829003601f168201915b50505091835250506040805161010081018252600184015481526002840154602080830191909152600385015482840152600485015460608084019190915260058601546080840152600686015460a0840152600786015460c0840152600886015460e0840152818501929092528251808301845260098601548152600a86015491810191909152600b8501548184015291830191909152600c9092015460ff1615159101529392505050565b60009081526001602052604090206009015490565b60035481565b600461137f82610b08565b600781111561138a57fe5b146113a75760405162461bcd60e51b815260040161036890611b84565b6000818152600160208190526040808320600c01805460ff19169092179091555182917f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f91a250565b600082820183811015611448576040805162461bcd60e51b815260206004820152601b60248201527a536166654d6174683a206164646974696f6e206f766572666c6f7760281b604482015290519081900360640190fd5b90505b92915050565b6000816114605750600061144b565b6114488261147661146f6115fb565b869061160b565b90611664565b600161148783610b08565b600781111561149257fe5b146114af5760405162461bcd60e51b815260040161036890611c76565b60008281526001602052604090206002015442106114df5760405162461bcd60e51b815260040161036890611be1565b600082815260026020908152604080832033845290915290205460ff16156115195760405162461bcd60e51b815260040161036890611e39565b600181600281111561152757fe5b14156115625760008281526001602081905260409091206009015461154b916113f0565b600083815260016020526040902060090155611593565b6000828152600160208190526040909120600a0154611580916113f0565b6000838152600160205260409020600a01555b600082815260026020908152604080832033845290915290819020805460ff191660011790555182907fe5a35d72c96d3a6952cdc06084bbe1e9ed2f42834f2f463c89110ab5777b84c5906115e9908490611b32565b60405180910390a25050565b303b1590565b6b033b2e3c9fd0803ce800000090565b60008261161a5750600061144b565b8282028284828161162757fe5b04146114485760405162461bcd60e51b8152600401808060200182810382526021815260200180611fac6021913960400191505060405180910390fd5b600061144883836040518060400160405280601a815260200179536166654d6174683a206469766973696f6e206279207a65726f60301b8152506000818361172a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156116ef5781810151838201526020016116d7565b50505050905090810190601f16801561171c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161173657fe5b0495945050505050565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b60405180608001604052806060815260200161178961182f565b8152602001611796611874565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826117d9576000855561181f565b82601f106117f257805160ff191683800117855561181f565b8280016001018555821561181f579182015b8281111561181f578251825591602001919060010190611804565b5061182b929150611895565b5090565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b60405180606001604052806000815260200160008152602001600081525090565b5b8082111561182b5760008155600101611896565b6000602082840312156118bb578081fd5b813561144881611f65565b6000602082840312156118d7578081fd5b815161144881611f65565b6000602082840312156118f3578081fd5b81518015158114611448578182fd5b60006020808385031215611914578182fd5b823567ffffffffffffffff8082111561192b578384fd5b818501915085601f83011261193e578384fd5b81358181111561194a57fe5b604051601f8201601f191681018501838111828210171561196757fe5b604052818152838201850188101561197d578586fd5b818585018683013790810190930193909352509392505050565b6000602082840312156119a8578081fd5b5035919050565b6000602082840312156119c0578081fd5b5051919050565b600080604083850312156119d9578081fd5b8235915060208301356119eb81611f65565b809150509250929050565b60008151808452815b81811015611a1b576020818501810151868301820152016119ff565b81811115611a2c5782602083870101525b50601f01601f19169290920160200192915050565b60006101a08251818552611a57828601826119f6565b9150506020830151611a6c6020860182611ab1565b506040830151611a80610120860182611a98565b50606083015115156101808501528091505092915050565b8051825260208082015190830152604090810151910152565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b6001600160a01b0391909116815260200190565b901515815260200190565b6020810160088310611b2c57fe5b91905290565b6020810160038310611b2c57fe5b60006101a0808352611b54818401886119f6565b915050611b646020830186611ab1565b611b72610120830185611a98565b82151561018083015295945050505050565b6020808252603f908201527f5b5145432d3033303030335d2d5468652070726f706f73616c20686173206e6f60408201527f74207061737365642c2074686520657865637574696f6e206661696c65642e00606082015260800190565b60208082526030908201527f5b5145432d3033303030315d2d566f74696e672069732066696e69736865642c60408201526f103330b4b632b2103a37903b37ba329760811b606082015260800190565b60208082526025908201527f5b5145432d3033303030355d2d50726f706f73616c20646f6573206e6f7420656040820152643c34b9ba1760d91b606082015260800190565b60208082526035908201527f5b5145432d3033303030305d2d50726f706f73616c20696e206e6f742070656e6040820152743234b73396103330b4b632b2103a37903b37ba329760591b606082015260800190565b60208082526021908201527f636f6e737469747574696f6e2e766f74696e672e656d675155706461746551526040820152604d60f81b606082015260800190565b60208082526022908201527f636f6e737469747574696f6e2e766f74696e672e656d6751557064617465524d60408201526120a560f11b606082015260800190565b6020808252818101527f636f6e737469747574696f6e2e766f74696e672e656d67515570646174655650604082015260600190565b6020808252604e908201527f5b5145432d3033303030345d2d5468652063616c6c657220646f6573206e6f7460408201527f2068617665206163636573732c206f6e6c792074686520726f6f74206e6f646560608201526d39903430bb329030b1b1b2b9b99760911b608082015260a00190565b60208082526022908201527f676f7665726e616e63652e636f6e737469747574696f6e2e706172616d657465604082015261727360f01b606082015260800190565b60208082526030908201527f5b5145432d3033303030325d2d5573657220616c726561647920766f7465642c60408201526f103330b4b632b2103a37903b37ba329760811b606082015260800190565b6020808252601f908201527f636f6e737469747574696f6e2e70726f706f73616c457865637574696f6e5000604082015260600190565b602080825260149082015273676f7665726e616e63652e726f6f744e6f64657360601b604082015260600190565b6000602082526114486020830184611a41565b600060a082019050825182526020830151602083015260408301516040830152606083015160608301526080830151608083015292915050565b90815260200190565b600083825260406020830152611f5d6040830184611a41565b949350505050565b6001600160a01b0381168114611f7a57600080fd5b5056fe496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212205af805022a03d21645994463994febd04f14afe396bff6e96d6f59ed498bace064736f6c63430007060033"

// DeployEmergencyUpdateVoting deploys a new Ethereum contract, binding an instance of EmergencyUpdateVoting to it.
func DeployEmergencyUpdateVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmergencyUpdateVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(EmergencyUpdateVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EmergencyUpdateVotingBin), backend)
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
func (_EmergencyUpdateVoting *EmergencyUpdateVotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_EmergencyUpdateVoting *EmergencyUpdateVotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(IVotingBaseProposal)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "getProposal", _proposalId)
	return *ret0, err
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
	var (
		ret0 = new(IVotingVotingStats)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "getProposalStats", _proposalId)
	return *ret0, err
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
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "getStatus", _proposalId)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "getVotesAgainst", _proposalId)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "getVotesFor", _proposalId)
	return *ret0, err
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "proposalCount")
	return *ret0, err
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
	ret := new(struct {
		Remark   string
		Params   IVotingVotingParams
		Counters IVotingVotingCounters
		Executed bool
	})
	out := ret
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
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
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EmergencyUpdateVoting.contract.Call(opts, out, "voted", arg0, arg1)
	return *ret0, err
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
	return event, nil
}
