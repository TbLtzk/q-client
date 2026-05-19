package governance

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/common"
)

const (
	governanceProposalPhaseUnknown  = "unknown"
	governanceProposalPhaseActive   = "active"
	governanceProposalPhaseDesired  = "desired"
	governanceProposalPhaseProposed = "proposed"
)

// GovernanceProposalThreshold reports signer progress against the active root set.
type GovernanceProposalThreshold struct {
	SignedCount      int  `json:"signedCount"`
	TotalRoots       int  `json:"totalRoots"`
	SignedPercent    int  `json:"signedPercent"`
	ThresholdPercent int  `json:"thresholdPercent"`
	MeetsThreshold   bool `json:"meetsThreshold"`
}

// GovernanceProposalStatus is the normalized status of a root-list or exclusion-list proposal hash.
type GovernanceProposalStatus struct {
	ProposalType   string                        `json:"proposalType"`
	Phase          string                        `json:"phase"`
	Hash           common.Hash                   `json:"hash"`
	Timestamp      uint64                        `json:"timestamp"`
	Signers        []common.Address              `json:"signers"`
	Threshold      GovernanceProposalThreshold   `json:"threshold"`
	NeedsSignature bool                          `json:"needsSignature"`
	QueriedSigner  common.Address                `json:"queriedSigner"`
}

func (a *GovernancePublicAPI) GetGovernanceProposalStatus(proposalType string, hash common.Hash, signer common.Address) (GovernanceProposalStatus, error) {
	switch proposalType {
	case common.GovernanceProposalTypeRootList:
		return a.gov.RootManager.rootListProposalStatus(hash, signer)
	case common.GovernanceProposalTypeExclusionList:
		return a.gov.RootManager.exclusionListProposalStatus(hash, signer)
	default:
		return GovernanceProposalStatus{}, fmt.Errorf("unsupported governance proposal type %q", proposalType)
	}
}

func (s *RootManager) rootListProposalStatus(hash common.Hash, queriedSigner common.Address) (GovernanceProposalStatus, error) {
	if hash == (common.Hash{}) {
		return GovernanceProposalStatus{}, errors.New("proposal hash is required")
	}

	s.rootLock.Lock()
	defer s.rootLock.Unlock()

	active := s.active
	if active != nil {
		active.aliases = s.getAliasesOfRoots(active.rootAddresses)
	}

	set, phase := s.findRootSetByHashLocked(hash)
	if set == nil {
		return GovernanceProposalStatus{
			ProposalType:  common.GovernanceProposalTypeRootList,
			Phase:         governanceProposalPhaseUnknown,
			Hash:          hash,
			QueriedSigner: queriedSigner,
		}, nil
	}

	return buildRootListProposalStatus(active, set, phase, queriedSigner), nil
}

func (s *RootManager) exclusionListProposalStatus(hash common.Hash, queriedSigner common.Address) (GovernanceProposalStatus, error) {
	if hash == (common.Hash{}) {
		return GovernanceProposalStatus{}, errors.New("proposal hash is required")
	}

	s.exLock.Lock()
	defer s.exLock.Unlock()

	s.rootLock.Lock()
	active := s.active
	if active != nil {
		active.aliases = s.getAliasesOfRoots(active.rootAddresses)
	}
	s.rootLock.Unlock()

	set, phase := s.findExclusionSetByHashLocked(hash)
	if set == nil {
		return GovernanceProposalStatus{
			ProposalType:  common.GovernanceProposalTypeExclusionList,
			Phase:         governanceProposalPhaseUnknown,
			Hash:          hash,
			QueriedSigner: queriedSigner,
		}, nil
	}

	return buildExclusionListProposalStatus(active, set, phase, queriedSigner), nil
}

func (s *RootManager) findRootSetByHashLocked(hash common.Hash) (*rootSet, string) {
	if s.active != nil && s.active.hash == hash {
		return s.active.copy(), governanceProposalPhaseActive
	}
	if s.desired != nil && s.desired.hash == hash {
		return s.desired.copy(), governanceProposalPhaseDesired
	}
	if s.proposed != nil && s.proposed.hash == hash {
		return s.proposed.copy(), governanceProposalPhaseProposed
	}
	return nil, governanceProposalPhaseUnknown
}

func (s *RootManager) findExclusionSetByHashLocked(hash common.Hash) (*exclusionSet, string) {
	if s.activeExSet != nil && s.activeExSet.hash == hash {
		return s.activeExSet.copy(), governanceProposalPhaseActive
	}
	if s.desiredExSet != nil && s.desiredExSet.hash == hash {
		return s.desiredExSet.copy(), governanceProposalPhaseDesired
	}
	if s.proposedExSet != nil && s.proposedExSet.hash == hash {
		return s.proposedExSet.copy(), governanceProposalPhaseProposed
	}
	return nil, governanceProposalPhaseUnknown
}

func buildRootListProposalStatus(active *rootSet, set *rootSet, phase string, queriedSigner common.Address) GovernanceProposalStatus {
	status := GovernanceProposalStatus{
		ProposalType:  common.GovernanceProposalTypeRootList,
		Phase:         phase,
		Hash:          set.hash,
		Timestamp:     set.timestamp,
		Signers:       uniqueKnownSigners(active, set.signers),
		QueriedSigner: queriedSigner,
	}

	if active != nil {
		signedCount, totalRoots, signedPercent, meetsThreshold := active.signerThresholdProgress(set.signers, rootListThresholdPercentage)
		status.Threshold = GovernanceProposalThreshold{
			SignedCount:      signedCount,
			TotalRoots:       totalRoots,
			SignedPercent:    signedPercent,
			ThresholdPercent: rootListThresholdPercentage,
			MeetsThreshold:   meetsThreshold,
		}
		status.NeedsSignature = active.needsSignatureFrom(set.signers, queriedSigner)
	}

	return status
}

func buildExclusionListProposalStatus(active *rootSet, set *exclusionSet, phase string, queriedSigner common.Address) GovernanceProposalStatus {
	status := GovernanceProposalStatus{
		ProposalType:  common.GovernanceProposalTypeExclusionList,
		Phase:         phase,
		Hash:          set.hash,
		Timestamp:     set.timestamp,
		Signers:       uniqueKnownSigners(active, set.signers),
		QueriedSigner: queriedSigner,
	}

	if active != nil {
		signedCount, totalRoots, signedPercent, meetsThreshold := active.signerThresholdProgress(set.signers, exclusionListThresholdPercentage)
		status.Threshold = GovernanceProposalThreshold{
			SignedCount:      signedCount,
			TotalRoots:       totalRoots,
			SignedPercent:    signedPercent,
			ThresholdPercent: exclusionListThresholdPercentage,
			MeetsThreshold:   meetsThreshold,
		}
		status.NeedsSignature = active.needsSignatureFrom(set.signers, queriedSigner)
	}

	return status
}

func uniqueKnownSigners(active *rootSet, proposalSigners map[common.Address][]byte) []common.Address {
	if active == nil {
		return nil
	}

	known := active.knownSigners(proposalSigners)
	seen := make(map[common.Address]struct{}, len(known))
	out := make([]common.Address, 0, len(known))
	for _, addr := range known {
		if _, exists := seen[addr]; exists {
			continue
		}
		seen[addr] = struct{}{}
		out = append(out, addr)
	}

	sort.Slice(out, func(i, j int) bool {
		return bytes.Compare(out[i].Bytes(), out[j].Bytes()) < 0
	})
	return out
}
