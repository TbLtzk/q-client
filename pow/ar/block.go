package ar

import (
	"math/big"

	"gitlab.com/q-dev/q-client/trie"
)

type Block interface {
	Trie() *trie.Trie
	Diff() *big.Int
}
