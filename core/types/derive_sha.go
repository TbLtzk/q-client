package types

import (
	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/ptrie"
)

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveSha(list DerivableList) []byte {
	db, _ := ethdb.NewMemDatabase()
	trie := ptrie.New(nil, db)
	for i := 0; i < list.Len(); i++ {
		trie.Update(ethutil.Encode(i), list.GetRlp(i))
	}

	return trie.Root()
}
