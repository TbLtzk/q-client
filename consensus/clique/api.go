// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package clique

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"gitlab.com/q-dev/q-client/accounts/abi/bind"
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/hexutil"
	"gitlab.com/q-dev/q-client/consensus"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/rlp"
	"gitlab.com/q-dev/q-client/rpc"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-authority scheme.
type API struct {
	chain  consensus.ChainHeaderReader
	clique *Clique
}

// GetSnapshot retrieves the state snapshot at a given block.
func (api *API) GetSnapshot(number *rpc.BlockNumber) (*Snapshot, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.clique.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, true)
}

// GetSnapshotAtHash retrieves the state snapshot at a given block.
func (api *API) GetSnapshotAtHash(hash common.Hash) (*Snapshot, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	return api.clique.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, true)
}

// GetSigners retrieves the list of authorized signers at the specified block.
func (api *API) GetSigners(number *rpc.BlockNumber) ([]common.Address, error) {
	// Retrieve the requested block number (or current if none requested)
	var header *types.Header
	if number == nil || *number == rpc.LatestBlockNumber {
		header = api.chain.CurrentHeader()
	} else {
		header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
	}
	// Ensure we have an actually valid block and return the signers from its snapshot
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.clique.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, true)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

// GetSignersAtHash retrieves the list of authorized signers at the specified block.
func (api *API) GetSignersAtHash(hash common.Hash) ([]common.Address, error) {
	header := api.chain.GetHeaderByHash(hash)
	if header == nil {
		return nil, errUnknownBlock
	}
	snap, err := api.clique.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, true)
	if err != nil {
		return nil, err
	}
	return snap.signers(), nil
}

// Proposals returns the current proposals the node tries to uphold and vote on.
func (api *API) Proposals() map[common.Address]bool {
	api.clique.lock.RLock()
	defer api.clique.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range api.clique.proposals {
		proposals[address] = auth
	}
	return proposals
}

// Propose injects a new authorization proposal that the signer will attempt to
// push through.
func (api *API) Propose(address common.Address, auth bool) {
	api.clique.lock.Lock()
	defer api.clique.lock.Unlock()

	api.clique.proposals[address] = auth
}

// Discard drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
func (api *API) Discard(address common.Address) {
	api.clique.lock.Lock()
	defer api.clique.lock.Unlock()

	delete(api.clique.proposals, address)
}

type status struct {
	InturnPercent float64                `json:"inturnPercent"`
	SigningStatus map[common.Address]int `json:"sealerActivity"`
	NumBlocks     uint64                 `json:"numBlocks"`
}

// Status returns the status of the last N blocks,
// - the number of active signers,
// - the number of signers,
// - the percentage of in-turn blocks
func (api *API) Status() (*status, error) {
	var (
		numBlocks = uint64(64)
		header    = api.chain.CurrentHeader()
		diff      = uint64(0)
		optimals  = 0
	)
	snap, err := api.clique.snapshot(api.chain, header.Number.Uint64(), header.Hash(), nil, true)
	if err != nil {
		return nil, err
	}
	var (
		signers = snap.signers()
		end     = header.Number.Uint64()
		start   = end - numBlocks
	)
	if numBlocks > end {
		start = 1
		numBlocks = end - start
	}
	signStatus := make(map[common.Address]int)
	for _, s := range signers {
		signStatus[s] = 0
	}
	for n := start; n < end; n++ {
		h := api.chain.GetHeaderByNumber(n)
		if h == nil {
			return nil, fmt.Errorf("missing block %d", n)
		}
		if h.Difficulty.Cmp(diffInTurn) == 0 {
			optimals++
		}
		diff += h.Difficulty.Uint64()
		sealer, err := api.clique.Author(h)
		if err != nil {
			return nil, err
		}
		signStatus[sealer]++
	}
	return &status{
		InturnPercent: float64(100*optimals) / float64(numBlocks),
		SigningStatus: signStatus,
		NumBlocks:     numBlocks,
	}, nil
}

type blockNumberOrHashOrRLP struct {
	*rpc.BlockNumberOrHash
	RLP hexutil.Bytes `json:"rlp,omitempty"`
}

func (sb *blockNumberOrHashOrRLP) UnmarshalJSON(data []byte) error {
	bnOrHash := new(rpc.BlockNumberOrHash)
	// Try to unmarshal bNrOrHash
	if err := bnOrHash.UnmarshalJSON(data); err == nil {
		sb.BlockNumberOrHash = bnOrHash
		return nil
	}
	// Try to unmarshal RLP
	var input string
	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}
	blob, err := hexutil.Decode(input)
	if err != nil {
		return err
	}
	sb.RLP = blob
	return nil
}

// GetSigner returns the signer for a specific clique block.
// Can be called with either a blocknumber, blockhash or an rlp encoded blob.
// The RLP encoded blob can either be a block or a header.
func (api *API) GetSigner(rlpOrBlockNr *blockNumberOrHashOrRLP) (common.Address, error) {
	if len(rlpOrBlockNr.RLP) == 0 {
		blockNrOrHash := rlpOrBlockNr.BlockNumberOrHash
		var header *types.Header
		if blockNrOrHash == nil {
			header = api.chain.CurrentHeader()
		} else if hash, ok := blockNrOrHash.Hash(); ok {
			header = api.chain.GetHeaderByHash(hash)
		} else if number, ok := blockNrOrHash.Number(); ok {
			header = api.chain.GetHeaderByNumber(uint64(number.Int64()))
		}
		if header == nil {
			return common.Address{}, fmt.Errorf("missing block %v", blockNrOrHash.String())
		}
		return api.clique.Author(header)
	}
	block := new(types.Block)
	if err := rlp.DecodeBytes(rlpOrBlockNr.RLP, block); err == nil {
		return api.clique.Author(block.Header())
	}
	header := new(types.Header)
	if err := rlp.DecodeBytes(rlpOrBlockNr.RLP, header); err != nil {
		return common.Address{}, err
	}
	return api.clique.Author(header)
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

func (api *API) GetValidatorsMetricsForCycle(cycleSeqNumber uint64) ([]ValidatorMetrics, error) {
	epoch := api.clique.config.Epoch
	highestBlock := api.chain.CurrentHeader().Number.Uint64()
	if cycleSeqNumber == 0 ||
		(cycleSeqNumber-1)*epoch > api.chain.CurrentHeader().Number.Uint64() {
		return []ValidatorMetrics{}, errors.New("number of cycle cannot be null or haven't been mined yet")
	}
	endOfCycleBlock := cycleSeqNumber * epoch
	startOfCycleBlock := endOfCycleBlock - epoch + 1
	status := "Final"
	latestFullCycleBlock := highestBlock - (highestBlock % epoch)
	if endOfCycleBlock > latestFullCycleBlock {
		status = "Pending"
		endOfCycleBlock = highestBlock
	}
	transitionBlock := rpc.BlockNumber(endOfCycleBlock - epoch)
	snapshot, err := api.GetSnapshot(&transitionBlock)
	if err != nil {
		return []ValidatorMetrics{}, err
	}
	signers := snapshot.signers()
	validatorMetrics := make([]ValidatorMetrics, len(signers))
	blockMetrics, err := api.calcMetrics(startOfCycleBlock, endOfCycleBlock, *snapshot)
	if err != nil {
		return []ValidatorMetrics{}, err
	}
	mainAccounts := api.clique.unAliasAccounts(signers, api.chain.Config().IsAthos(api.chain.CurrentHeader().Number))
	for idx, signer := range signers {
		if entry, ok := blockMetrics[signer]; ok {
			entry.DueBlocks = epoch / uint64(len(signers))
			if uint64(api.getIndexInAddressSlice(signers, signer)) < epoch%uint64(len(signers)) {
				entry.DueBlocks++
			}
			blockMetrics[signer] = entry
		}
		validatorMetrics[idx] = ValidatorMetrics{
			CycleNumber:  cycleSeqNumber,
			StartBlock:   endOfCycleBlock - epoch + 1,
			EndBlock:     endOfCycleBlock,
			MinerAccount: signers[idx],
			MainAccount:  mainAccounts[idx],
			BlockMetrics: blockMetrics[signer],
			Status:       status,
		}
	}
	return validatorMetrics, nil
}

func (api *API) calcMetrics(startOfCycleBlock uint64, endOfCycleBlock uint64, snapshot Snapshot) (map[common.Address]ValidatorBlocks, error) {
	blockMetrics := make(map[common.Address]ValidatorBlocks)
	for _, signer := range snapshot.signers() {
		blockMetrics[signer] = ValidatorBlocks{}
	}
	for startOfCycleBlock <= endOfCycleBlock {
		blockOfCycle := rpc.BlockNumber(startOfCycleBlock)
		header := api.chain.GetHeaderByNumber(startOfCycleBlock)
		snapshotOfBlock, err := api.GetSnapshot(&blockOfCycle)
		if err != nil {
			return blockMetrics, err
		}
		actualSigner, err := ecrecover(header, snapshotOfBlock.sigcache)
		if err != nil {
			return blockMetrics, err
		}
		if entry, ok := blockMetrics[actualSigner]; ok {
			if snapshot.inturn(startOfCycleBlock, actualSigner) {
				entry.InTurnBlocks++
			} else {
				entry.OutOfTurnBlocks++
			}
			blockMetrics[actualSigner] = entry
		}
		startOfCycleBlock++
	}
	return blockMetrics, nil
}

func (api *API) GetEpochLength() uint64 {
	return api.clique.config.Epoch
}

// GetOutOfTurnStatsByNumber returns the stats by block:
// - the block number
// - the block hash
// - the difficulty
// - the actual signer
// - the inturn signer
func (api *API) GetOutOfTurnStatsByNumber(block *rpc.BlockNumber) (*OutOfTurnStats, error) {
	header := api.chain.GetHeaderByNumber(uint64(block.Int64()))
	snapshot, err := api.GetSnapshot(block)
	if err != nil {
		return nil, err
	}
	return api.getOutOfTurnStatsFromSnapshot(header, snapshot)
}

// GetOutOfTurnStatsByHash returns the stats by hash.
// See function GetOutOfTurnStatsByNumber for return data.
func (api *API) GetOutOfTurnStatsByHash(hash common.Hash) (*OutOfTurnStats, error) {
	header := api.chain.GetHeaderByHash(hash)
	snapshot, err := api.GetSnapshotAtHash(hash)
	if err != nil {
		return nil, err
	}
	return api.getOutOfTurnStatsFromSnapshot(header, snapshot)
}

func (api *API) getOutOfTurnStatsFromSnapshot(header *types.Header, snapshot *Snapshot) (*OutOfTurnStats, error) {
	actualSigner, err := ecrecover(header, api.clique.signatures)
	if err != nil {
		return nil, err
	}
	origAccount := actualSigner
	if api.chain.Config().IsAthos(header.Number) {
		origAccount = api.clique.unAliasAccounts([]common.Address{actualSigner}, api.chain.Config().IsAthos(api.chain.CurrentHeader().Number))[0]
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

func (api *API) getIndexInAddressSlice(a []common.Address, x common.Address) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func (api *API) getInTurnSigner(snapshot *Snapshot) common.Address {
	signers := snapshot.signers()
	index := snapshot.Number % uint64(len(signers))
	inTurnSigner := signers[index]
	return inTurnSigner
}

func (api *API) GetConstitutionVotings(proposalCounter int64) ([]ConstitutionVoting, error) {
	var votings []ConstitutionVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.ConstitutionVoting()
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

func (api *API) GetGeneralUpdateVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.GeneralUpdateVoting()
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

func (api *API) GetEmergencyUpdateVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EmergencyUpdateVoting()
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

func (api *API) GetRootsVotings(proposalCounter int64, blockNumber int64) ([]RootsVoting, error) {
	var votings []RootsVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	votingFilterOpts := &bind.FilterOpts{
		Start:   uint64(blockNumber),
		End:     nil,
		Context: nil,
	}

	provider := api.clique.registry.RootsVoting()
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

func (api *API) GetRootNodesSlashingVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.RootNodesSlashingVoting()
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

func (api *API) GetValidatorsSlashingVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.ValidatorsSlashingVoting()
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

func (api *API) GetEpqfiMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EpqfiMembershipVoting()
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

func (api *API) GetEpqfiParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EpqfiParametersVoting()
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

func (api *API) GetEpdrMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EpdrMembershipVoting()
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

func (api *API) GetEpdrParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EpdrParametersVoting()
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

func (api *API) GetEprsMembershipVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EprsMembershipVoting()
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

func (api *API) GetEprsParametersVotings(proposalCounter int64) ([]GeneralVoting, error) {
	var votings []GeneralVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.EprsParametersVoting()
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

func (api *API) GetContractRegistryAddressVotings(proposalCounter int64) ([]ContractRegistryVoting, error) {
	var votings []ContractRegistryVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.ContractRegistryAddressVoting()
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

func (api *API) GetContractRegistryUpgradeVotings(proposalCounter int64) ([]ContractRegistryVoting, error) {
	var votings []ContractRegistryVoting
	votingOpts := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	provider := api.clique.registry.ContractRegistryUpgradeVoting()
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
