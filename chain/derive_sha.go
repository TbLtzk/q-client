package chain

import (
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/trie"
)

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveSha(list DerivableList) []byte {
	trie := trie.New(ethutil.Config.Db, "")
	for i := 0; i < list.Len(); i++ {
		trie.Update(string(ethutil.NewValue(i).Encode()), string(list.GetRlp(i)))
	}

	return trie.GetRoot()
}
