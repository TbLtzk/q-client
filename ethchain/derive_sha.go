package ethchain

import (
	"gitlab.com/q-dev/q-client/ethtrie"
	"gitlab.com/q-dev/q-client/ethutil"
)

type DerivableList interface {
	Len() int
	GetRlp(i int) []byte
}

func DeriveSha(list DerivableList) []byte {
	trie := ethtrie.New(ethutil.Config.Db, "")
	for i := 0; i < list.Len(); i++ {
		trie.Update(string(ethutil.NewValue(i).Encode()), string(list.GetRlp(i)))
	}

	return trie.GetRoot()
}
