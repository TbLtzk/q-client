package types

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/trie"
)

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveSha(list DerivableList) common.Hash {
	db, _ := ethdb.NewMemDatabase()
	trie := trie.New(nil, db)
	for i := 0; i < list.Len(); i++ {
		trie.Update(common.Encode(i), list.GetRlp(i))
	}

	return common.BytesToHash(trie.Root())
}
