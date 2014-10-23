package ar

import (
	"math/big"

	"gitlab.com/q-dev/q-client/ethtrie"
)

type Block interface {
	Trie() *ethtrie.Trie
	Diff() *big.Int
}
