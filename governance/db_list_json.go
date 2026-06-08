package governance

import (
	"encoding/json"

	"gitlab.com/q-dev/q-client/common"
)

// rootListDB is the LevelDB JSON shape for root lists. Signatures use the legacy
// encoding/json [][]byte form (base64 elements), not govPub 0x hex.
type rootListDB struct {
	Timestamp  uint64           `json:"timestamp"`
	Nodes      []common.Address `json:"nodes"`
	Hash       common.Hash      `json:"hash"`
	Signatures [][]byte         `json:"signatures"`
}

// exclusionListDB is the LevelDB JSON shape for exclusion lists (legacy signatures).
type exclusionListDB struct {
	Timestamp  uint64                     `json:"timestamp"`
	Validators []common.ExcludedValidator `json:"validators"`
	Hash       common.Hash                `json:"hash"`
	Signatures [][]byte                   `json:"signatures"`
}

func rootSetToDBList(set *rootSet) rootListDB {
	list := set.makeList()
	sigs := make([][]byte, len(list.Signatures))
	for i, sig := range list.Signatures {
		sigs[i] = sig
	}
	return rootListDB{
		Timestamp:  list.Timestamp,
		Nodes:      list.Nodes,
		Hash:       list.Hash,
		Signatures: sigs,
	}
}

func exclusionSetToDBList(set *exclusionSet) exclusionListDB {
	list := set.makeList()
	sigs := make([][]byte, len(list.Signatures))
	for i, sig := range list.Signatures {
		sigs[i] = sig
	}
	return exclusionListDB{
		Timestamp:  list.Timestamp,
		Validators: list.Validators,
		Hash:       list.Hash,
		Signatures: sigs,
	}
}

func rootSetFromDBJSON(raw []byte) (*rootSet, error) {
	var stored rootListDB
	if err := json.Unmarshal(raw, &stored); err != nil {
		return nil, err
	}
	return newRootSet(&common.RootList{
		Timestamp:  stored.Timestamp,
		Nodes:      stored.Nodes,
		Hash:       stored.Hash,
		Signatures: common.HexSignaturesFromBytes(stored.Signatures),
	})
}

func exclusionSetFromDBJSON(raw []byte) (*exclusionSet, error) {
	var stored exclusionListDB
	if err := json.Unmarshal(raw, &stored); err != nil {
		return nil, err
	}
	return newExclusionSet(&common.ValidatorExclusionList{
		Timestamp:  stored.Timestamp,
		Validators: stored.Validators,
		Hash:       stored.Hash,
		Signatures: common.HexSignaturesFromBytes(stored.Signatures),
	})
}

func exclusionSetsFromQuarantineJSON(raw []byte) ([]exclusionSet, error) {
	var stored []exclusionListDB
	if err := json.Unmarshal(raw, &stored); err != nil {
		return nil, err
	}
	var sets []exclusionSet
	for i := range stored {
		set, err := newExclusionSet(&common.ValidatorExclusionList{
			Timestamp:  stored[i].Timestamp,
			Validators: stored[i].Validators,
			Hash:       stored[i].Hash,
			Signatures: common.HexSignaturesFromBytes(stored[i].Signatures),
		})
		if err != nil {
			continue
		}
		sets = append(sets, *set)
	}
	return sets, nil
}

func exclusionDBListsFromSets(sets []exclusionSet) []exclusionListDB {
	out := make([]exclusionListDB, len(sets))
	for i := range sets {
		out[i] = exclusionSetToDBList(&sets[i])
	}
	return out
}
