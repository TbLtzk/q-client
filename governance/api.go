package governance

import (
	"math/big"
	"sort"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/log"
)

// GovernanceAPI. (shouldn't be opened via http/ws)
type GovernanceAPI struct {
	*GovernancePublicAPI
}

// GovernanceExtAPI
type GovernancePublicAPI struct {
	gov *Governance
}

// NewGovernanceAPI.
func NewGovernanceAPI(back *Governance, extApi *GovernancePublicAPI) *GovernanceAPI {
	return &GovernanceAPI{GovernancePublicAPI: extApi}
}

// NewGovernancePublicAPI
func NewGovernancePublicAPI(back *Governance) *GovernancePublicAPI {
	return &GovernancePublicAPI{gov: back}
}

func (a *GovernancePublicAPI) ActiveRootList() *RootList {
	return newRootList(a.gov.RootManager.getActiveRootSet(true))
}

func (a *GovernancePublicAPI) DesiredRootList() *RootList {
	return newRootList(a.gov.RootManager.getDesiredRootSet(true))
}

func (a *GovernancePublicAPI) ProposedRootList() *RootList {
	return newRootList(a.gov.RootManager.getProposedRootSet(true))
}

func (a *GovernancePublicAPI) OnchainRootList() *RootList {
	return newRootList(a.gov.RootManager.getOnchainRootSet(true))
}

func (a *GovernanceAPI) ProposeRootListUpdate(list common.RootList) (common.Hash, error) {
	set, err := newRootSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid root list")
	}
	set.updateAliases(a.gov.RootManager.getAliasesOfRoots(set.rootAddresses))

	set, err = a.gov.RootManager.proposeRootSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) ProposeOnchainRootList() (common.Hash, error) {
	set := a.gov.RootManager.getOnchainRootSet(true)
	if set == nil {
		return common.Hash{}, errors.New("can't get on-cain root set")
	}

	set, err := a.gov.RootManager.proposeRootSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedRootList() error {
	return a.gov.RootManager.acceptProposedRootList(true)
}

func (a *GovernancePublicAPI) DiffRootList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffRootListByName(nameA, nameB, true)
}

func (a *GovernancePublicAPI) ActiveExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getActiveExclusionSet())
}

func (a *GovernancePublicAPI) ActiveExclusionListPrettify() string {
	return printPretiffiedList(newExclusionListPrettify(a.gov.RootManager.getActiveExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64()))
}

func (a *GovernancePublicAPI) DesiredExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getDesiredExclusionSet())
}

func (a *GovernancePublicAPI) DesiredExclusionListPrettify() string {
	return printPretiffiedList(newExclusionListPrettify(a.gov.RootManager.getDesiredExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64()))
}

func (a *GovernancePublicAPI) ProposedExclusionList() *ExclusionList {
	return newExclusionList(a.gov.RootManager.getProposedExclusionSet())
}

func (a *GovernancePublicAPI) ProposedExclusionListPrettify() string {
	return printPretiffiedList(newExclusionListPrettify(a.gov.RootManager.getProposedExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64()))
}

func (a *GovernancePublicAPI) IsInExclusionList(address string) string {
	return printPrettifiedSearch(checkExclusionLists(common.HexToAddress(address), a.gov.RootManager.getActiveExclusionSet(),
		a.gov.RootManager.getDesiredExclusionSet(), a.gov.RootManager.getProposedExclusionSet(), a.gov.RootManager.bc.CurrentBlock().Number().Int64()))
}

func (a *GovernanceAPI) ProposeExclusionListUpdate(list common.ValidatorExclusionList) (common.Hash, error) {
	set, err := newExclusionSet(&list)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "invalid exclusion list")
	}

	set, err = a.gov.RootManager.proposeExclusionSet(set)
	if err != nil {
		return common.Hash{}, err
	}

	return set.hash, nil
}

func (a *GovernanceAPI) AcceptProposedExclusionList() error {
	return a.gov.RootManager.acceptProposedExclusionList(true)
}

func (a *GovernancePublicAPI) DiffExclusionList(nameA, nameB string) ([]DiffEntry, error) {
	return a.gov.RootManager.diffExclusionListByName(nameA, nameB)
}

func (a *GovernanceAPI) AddConstitutionFile(filename string, constitutionHash *common.Hash) error {
	return a.gov.ConstitutionManager.addConstitutionFile(filename)
}

// RequestForConstitutionFile creates request for specific constitution file from it's peers.
// Once this request created, node will be asking its peers for this file until it succeeded
func (a *GovernanceAPI) RequestForConstitutionFile(constitutionHash *common.Hash) error {
	hash, err := a.gov.ConstitutionManager.addConstitutionFileRequest(constitutionHash)

	if err != nil {
		return err
	}

	var hashes []common.Hash
	hashes = append(hashes, *hash)
	newReq := common.ConstitutionFilesRequest{Hashes: hashes}

	a.gov.handler.broadcastConstitutionRequest(&newReq)

	log.Info("Request for the constitution file with hash " + constitutionHash.String() + " created. Once one of your peers has file with the required hash - you'll be informed.")

	return nil
}

func (a *GovernanceAPI) ConstitutionFileRequests() ([]common.Hash, error) {
	return a.gov.ConstitutionManager.db.getConstitutionFileRequests()
}

func (a *GovernanceAPI) ConstitutionFiles() ([]common.ConstitutionFile, error) {
	return a.gov.ConstitutionManager.db.getConstitutionFiles()
}

func (a *GovernanceAPI) GetRootNodeApprovals(blockNumber *big.Int, hash *common.Hash) (*[]common.RootNodeApproval, error) {
	list, err := a.gov.RootManager.getActiveApprovalList(blockNumber, hash)
	var res []common.RootNodeApproval
	if list != nil {
		res = list.Approvals
	}
	return &res, err
}

func (a *GovernanceAPI) GetLatestTransitionBlocks(amount int) (*[]TransitionBlocksWithApproval, error) {
	currentBlockNumber := a.gov.RootManager.bc.CurrentBlock().Number().Uint64()
	if uint64(amount) > currentBlockNumber/101 {
		log.Error("Not enough transition blocks!")
		return nil, nil
	}
	transitionBlockNumber := currentBlockNumber - currentBlockNumber%101
	transitionBlocks := make([]TransitionBlocksWithApproval, amount)
	for amount > 0 {
		transitionBlock := a.gov.RootManager.bc.GetHeaderByNumber(transitionBlockNumber)
		rNApprovals, err := a.GetRootNodeApprovals(transitionBlock.Number, nil)
		if err != nil {
			return nil, err
		}
		var approvals []Approval
		for _, approval := range *rNApprovals {
			if approval.Hash == transitionBlock.Hash() {
				approvals = append(approvals, Approval{
					Signer:    approval.Signer,
					Signature: approval.Signature,
				})
			}
		}
		transitionBlocks[amount-1] = TransitionBlocksWithApproval{
			BlockNumber: transitionBlock.Number,
			Hash:        transitionBlock.Hash(),
			Approvals:   approvals,
		}
		transitionBlockNumber = transitionBlockNumber - 101
		amount--
	}
	return &transitionBlocks, nil
}

type Approval struct {
	Signer    common.Address `json:"signer"`
	Signature []byte         `json:"signature"`
}

type TransitionBlocksWithApproval struct {
	BlockNumber *big.Int    `json:"blockNumber"`
	Hash        common.Hash `json:"hash"`
	Approvals   []Approval  `json:"approvals"`
}

type RootList struct {
	Timestamp uint64           `json:"timestamp"`
	Nodes     []common.Address `json:"nodes"`
	Hash      common.Hash      `json:"hash"`

	Signers []common.Address `json:"signers"`
}

func newRootList(set *rootSet) *RootList {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	return &RootList{
		Timestamp: set.timestamp,
		Hash:      set.hash,
		Nodes:     set.rootAddresses,
		Signers:   signers,
	}
}

type ExclusionList struct {
	Timestamp  uint64                     `json:"timestamp"`
	Validators []common.ExcludedValidator `json:"validators"`
	Hash       common.Hash                `json:"hash"`

	Signers []common.Address `json:"signers"`
}

type ExclusionListPrettify struct {
	Hash       common.Hash                 `json:"hash"`
	Timestamp  string                      `json:"timestamp (created at)"`
	Signers    []common.Address            `json:"signers"`
	Validators map[common.Address][]string `json:"validators"`
}

func newExclusionList(set *exclusionSet) *ExclusionList {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	l := set.makeList()
	return &ExclusionList{
		Timestamp:  set.timestamp,
		Hash:       set.hash,
		Validators: l.Validators,
		Signers:    signers,
	}
}

func newExclusionListPrettify(set *exclusionSet, currentBlock int64) *ExclusionListPrettify {
	if set == nil {
		return nil
	}

	var signers []common.Address
	for addr := range set.signers {
		signers = append(signers, addr)
	}

	formattedBlockRanges := make(map[common.Address][]string)

	for address, blockRangesByAddress := range set.blockRanges {
		formattedBlockRanges[address] = getBanStates(blockRangesByAddress, currentBlock)
	}

	timeUnix := time.Unix(int64(set.timestamp), 0)
	return &ExclusionListPrettify{
		Timestamp:  strconv.FormatInt(int64(set.timestamp), 10) + " (" + timeUnix.String() + ")",
		Hash:       set.hash,
		Signers:    signers,
		Validators: formattedBlockRanges,
	}
}

func getBanStates(blockRangesByAddress []common.BlockRange, currentBlock int64) []string {
	var endBlockString string
	var stringBlockRanges []string
	for _, br := range blockRangesByAddress {
		startBlock := int64(br.StartAddress)
		endBlock := int64(br.EndAddress)
		switch {
		case endBlock == 0:
			endBlockString = "\t(active, open-ended)"
		case endBlock < currentBlock && startBlock < currentBlock:
			endBlockString = " - #" + strconv.FormatInt(endBlock, 10) + "\t(expired)"
		case endBlock > currentBlock && startBlock > currentBlock:
			endBlockString = " - #" + strconv.FormatInt(endBlock, 10) + "\t(scheduled at block #" + strconv.FormatInt(startBlock, 10) + ")"
		case endBlock > currentBlock && startBlock < currentBlock:
			endBlockString = " - #" + strconv.FormatInt(endBlock, 10) +
				"\t(active, ends at block #" + strconv.FormatInt(endBlock, 10) + ")"
		}
		stringBlockRange := "#" + strconv.FormatInt(startBlock, 10) + endBlockString
		stringBlockRanges = append(stringBlockRanges, stringBlockRange)
	}
	return stringBlockRanges
}

func checkExclusionLists(address common.Address, activeExSet *exclusionSet, desiredExSet *exclusionSet, proposedExSet *exclusionSet, currentBlock int64) map[string][]string {
	searchResults := make(map[string][]string)
	var searchResultsByList []string

	activeString := "Active exclusion list"
	activeString, searchResultsByList = searchAddressInList(activeString, activeExSet, address, currentBlock)
	searchResults[activeString] = searchResultsByList

	desiredString := "Desired exclusion list"
	desiredString, searchResultsByList = searchAddressInList(desiredString, desiredExSet, address, currentBlock)
	searchResults[desiredString] = searchResultsByList

	proposedString := "Proposed exclusion list"
	proposedString, searchResultsByList = searchAddressInList(proposedString, proposedExSet, address, currentBlock)
	searchResults[proposedString] = searchResultsByList

	return searchResults
}

func searchAddressInList(setString string, set *exclusionSet, address common.Address, currentBlock int64) (string, []string) {
	if set != nil && set.blockRanges[address] != nil && getBanStates(set.blockRanges[address], currentBlock) != nil {
		setString += " (" + set.hash.String() + ", " + strconv.FormatInt(int64(set.timestamp), 10) + "), on block range(s)"
		return setString, getBanStates(set.blockRanges[address], currentBlock)
	} else {
		setString += ", provided address is absent"
		return setString, nil
	}
}

func printPretiffiedList(prettified *ExclusionListPrettify) string {
	res := "Hash: " + prettified.Hash.String() + "\n"
	res += "Timestamp(created at): " + prettified.Timestamp + "\n"
	res += "Signers: \n"
	for _, s := range prettified.Signers {
		res += "\t" + s.String() + "\n"
	}
	res += "Validators: \n"
	for address, branges := range prettified.Validators {
		res += "\t" + address.String() + ": \n"
		sort.Strings(branges)
		for _, br := range branges {
			res += "\t\t" + br + "\n"
		}
	}
	return res
}

func printPrettifiedSearch(searchResults map[string][]string) string {
	res := ""
	for listType, branges := range searchResults {
		res += listType + "\n"
		sort.Strings(branges)
		for _, br := range branges {
			res += "\t" + br + "\n"
		}
	}
	return res
}
