package indexer

import (
	"math/big"
	"reflect"

	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/consensus/clique"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/rpc"
)

type IndexerAPI struct {
	indexer *Indexer
}

func NewIndexerAPI(i *Indexer) *IndexerAPI {
	return &IndexerAPI{indexer: i}
}

type OutOfTurnStats struct {
	BlockNumber  uint64         `json:"blockNumber"`
	BlockHash    common.Hash    `json:"blockHash"`
	Difficulty   *big.Int       `json:"difficulty"`
	ActualSigner common.Address `json:"actualSigner"`
	MainAccount  common.Address `json:"mainAccount"`
	InTurnSigner common.Address `json:"inTurnSigner"`
}

type ValidatorBlocks struct {
	DueBlocks       uint64 `json:"dueBlocks"`
	InTurnBlocks    uint64 `json:"inTurnBlocks"`
	OutOfTurnBlocks uint64 `json:"outOfTurnBlocks"`
}

type ValidatorMetrics struct {
	CycleNumber  uint64          `json:"cycleNumber"`
	StartBlock   uint64          `json:"startBlock"`
	EndBlock     uint64          `json:"endBlock"`
	MinerAccount common.Address  `json:"minerAccount"`
	MainAccount  common.Address  `json:"mainAccount"`
	BlockMetrics ValidatorBlocks `json:"blockMetrics"`
	Status       string          `json:"status"`
}

// GetOutOfTurnStatsByNumber returns the stats by block:
// - the block number
// - the block hash
// - the difficulty
// - the actual signer
// - the inturn signer
func (api *IndexerAPI) GetOutOfTurnStatsByNumber(block *rpc.BlockNumber) (*OutOfTurnStats, error) {
	header := api.indexer.clique.ChainHeaderReader().GetHeaderByNumber(uint64(block.Int64()))
	snapshot, err := api.indexer.clique.GetSnapshot(block)
	if err != nil {
		return nil, err
	}
	return api.getOutOfTurnStatsFromSnapshot(header, snapshot)
}

// GetOutOfTurnStatsByHash returns the stats by hash.
// See function GetOutOfTurnStatsByNumber for return data.
func (api *IndexerAPI) GetOutOfTurnStatsByHash(hash common.Hash) (*OutOfTurnStats, error) {
	header := api.indexer.clique.ChainHeaderReader().GetHeaderByHash(hash)
	snapshot, err := api.indexer.clique.GetSnapshotAtHash(hash)
	if err != nil {
		return nil, err
	}
	return api.getOutOfTurnStatsFromSnapshot(header, snapshot)
}

func (api *IndexerAPI) getOutOfTurnStatsFromSnapshot(header *types.Header, snapshot *clique.Snapshot) (*OutOfTurnStats, error) {
	actualSigner, err := api.indexer.clique.Ecrecover(header, api.indexer.clique.Clique())
	if err != nil {
		return nil, err
	}
	origAccount := actualSigner
	if api.indexer.clique.ChainHeaderReader().Config().IsAthos(header.Number) {
		origAccount = api.indexer.clique.Clique().UnAliasAccounts([]common.Address{actualSigner},
			api.indexer.clique.ChainHeaderReader().Config().IsAthos(api.indexer.clique.ChainHeaderReader().CurrentHeader().Number))[0]
	}
	inTurnSigner := api.getInTurnSigner(snapshot)
	return &OutOfTurnStats{
		BlockNumber:  header.Number.Uint64(),
		BlockHash:    header.Hash(),
		Difficulty:   header.Difficulty,
		ActualSigner: actualSigner,
		MainAccount:  origAccount,
		InTurnSigner: inTurnSigner,
	}, nil
}

func (api *IndexerAPI) getIndexInAddressSlice(a []common.Address, x common.Address) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func (api *IndexerAPI) getInTurnSigner(snapshot *clique.Snapshot) common.Address {
	signers := snapshot.SignersList()
	if len(signers) == 0 {
		return common.Address{}
	}
	index := snapshot.Number % uint64(len(signers))
	inTurnSigner := signers[index]
	return inTurnSigner
}

func (api *IndexerAPI) GetConstitutionVotings(proposalCounter int64) ([]ConstitutionVoting, error) {
	var votings []ConstitutionVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().ConstitutionVoting()
	currentProposalCounter, err := provider.ProposalCounter(votingOpts)
	if err != nil {
		return []ConstitutionVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []ConstitutionVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []ConstitutionVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []ConstitutionVoting{}, err
		}
		voting := ConstitutionVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingSubType:   proposal.Classification,
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetGeneralUpdateVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().GeneralUpdateVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Params.VotingStartTime,
			VotingEndTime:   proposal.Params.VotingEndTime,
			VetoEndTime:     proposal.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEmergencyUpdateVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EmergencyUpdateVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Params.VotingStartTime,
			VotingEndTime:   proposal.Params.VotingEndTime,
			VetoEndTime:     proposal.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}

	return votings, nil
}

func (api *IndexerAPI) GetRootsVotings(proposalCounter int64, blockNumber uint64, lastBlockNumber uint64) ([]RootsVoting, error) {
	var votings []RootsVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	votingFilterOpts := &bind.FilterOpts{
		Start:   blockNumber,
		End:     &lastBlockNumber,
		Context: nil,
	}

	provider := api.indexer.clique.ContractsRegistry().RootsVoting()
	filterIter, err := provider.FilterProposalCreated(votingFilterOpts)
	if err != nil {
		return []RootsVoting{}, err
	}

	maxId := int64(0)
	for filterIter.Next() {
		maxId = filterIter.Event.Id.Int64()
	}
	currentProposalCounter := big.NewInt(maxId)

	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []RootsVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []RootsVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []RootsVoting{}, err
		}
		voting := RootsVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
			Candidate:       proposal.Candidate,
			ReplaceDest:     proposal.ReplaceDest,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetRootNodesSlashingVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().RootNodesSlashingVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetValidatorsSlashingVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().ValidatorsSlashingVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEpqfiMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EpqfiMembershipVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEpqfiParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EpqfiParametersVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEpdrMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EpdrMembershipVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEpdrParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EpdrParametersVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEprsMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EprsMembershipVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetEprsParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().EprsParametersVoting()
	currentProposalCounter, err := provider.ProposalCount(votingOpts)
	if err != nil {
		return []GeneralVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposal, err := provider.Proposals(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []GeneralVoting{}, err
		}
		voting := GeneralVoting{
			ProposalId:      cycleId,
			VotingType:      reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:    proposalStatus,
			CurrentQuorum:   proposalStats.CurrentQuorum,
			RequiredQuorum:  proposalStats.RequiredQuorum,
			VotingStartTime: proposal.Base.Params.VotingStartTime,
			VotingEndTime:   proposal.Base.Params.VotingEndTime,
			VetoEndTime:     proposal.Base.Params.VetoEndTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetContractRegistryAddressVotings(proposalCounter int64) ([]ContractRegistryVoting, error) {
	var votings []ContractRegistryVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().ContractRegistryAddressVoting()
	currentProposalCounter, err := provider.ProposalsCount(votingOpts)
	if err != nil {
		return []ContractRegistryVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		proposal, err := provider.GetProposal(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		voting := ContractRegistryVoting{
			ProposalId:        cycleId,
			VotingType:        reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:      proposalStatus,
			CurrentMajority:   proposalStats.CurrentMajority,
			RequiredMajority:  proposalStats.RequiredMajority,
			VotingStartTime:   proposal.VotingStartTime,
			VotingExpiredTime: proposal.VotingExpiredTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

func (api *IndexerAPI) GetContractRegistryUpgradeVotings(proposalCounter int64) ([]ContractRegistryVoting, error) {
	var votings []ContractRegistryVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.indexer.clique.ContractsRegistry().ContractRegistryUpgradeVoting()
	currentProposalCounter, err := provider.ProposalsCount(votingOpts)
	if err != nil {
		return []ContractRegistryVoting{}, err
	}
	for id := proposalCounter + 1; id < currentProposalCounter.Int64(); id++ {
		cycleId := big.NewInt(id)
		proposalStatus, err := provider.GetStatus(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		proposal, err := provider.GetProposal(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		proposalStats, err := provider.GetProposalStats(votingOpts, cycleId)
		if err != nil {
			return []ContractRegistryVoting{}, err
		}
		voting := ContractRegistryVoting{
			ProposalId:        cycleId,
			VotingType:        reflect.TypeOf(provider).Elem().Name(),
			VotingStatus:      proposalStatus,
			CurrentMajority:   proposalStats.CurrentMajority,
			RequiredMajority:  proposalStats.RequiredMajority,
			VotingStartTime:   proposal.VotingStartTime,
			VotingExpiredTime: proposal.VotingExpiredTime,
		}
		votings = append(votings, voting)
	}
	return votings, nil
}

type ConstitutionVoting struct {
	ProposalId      *big.Int `json:"proposalId"`
	VotingType      string   `json:"votingType"`
	VotingSubType   uint8    `json:"votingSubType"`
	VotingStatus    uint8    `json:"votingStatus"`
	VotingStartTime *big.Int `json:"votingStartTime"`
	CurrentQuorum   *big.Int `json:"currentQuorum"`
	RequiredQuorum  *big.Int `json:"requiredQuorum"`
	VotingEndTime   *big.Int `json:"votingEndTime"`
	VetoEndTime     *big.Int `json:"vetoEndTime"`
}

type GeneralVoting struct {
	ProposalId      *big.Int `json:"proposalId"`
	VotingType      string   `json:"votingType"`
	VotingStatus    uint8    `json:"votingStatus"`
	VotingStartTime *big.Int `json:"votingStartTime"`
	CurrentQuorum   *big.Int `json:"currentQuorum"`
	RequiredQuorum  *big.Int `json:"requiredQuorum"`
	VotingEndTime   *big.Int `json:"votingEndTime"`
	VetoEndTime     *big.Int `json:"vetoEndTime"`
}

type RootsVoting struct {
	ProposalId      *big.Int       `json:"proposalId"`
	VotingType      string         `json:"votingType"`
	VotingStatus    uint8          `json:"votingStatus"`
	VotingStartTime *big.Int       `json:"votingStartTime"`
	CurrentQuorum   *big.Int       `json:"currentQuorum"`
	RequiredQuorum  *big.Int       `json:"requiredQuorum"`
	VotingEndTime   *big.Int       `json:"votingEndTime"`
	VetoEndTime     *big.Int       `json:"vetoEndTime"`
	Candidate       common.Address `json:"candidate"`
	ReplaceDest     common.Address `json:"replaceDest"`
}

type ContractRegistryVoting struct {
	ProposalId        *big.Int `json:"proposalId"`
	VotingType        string   `json:"votingType"`
	VotingStatus      uint8    `json:"votingStatus"`
	VotingStartTime   *big.Int `json:"votingStartTime"`
	CurrentMajority   *big.Int `json:"currentQuorum"`
	RequiredMajority  *big.Int `json:"requiredQuorum"`
	VotingExpiredTime *big.Int `json:"votingExpiredTime"`
}
