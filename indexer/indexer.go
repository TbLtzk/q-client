package indexer

import (
	"gitlab.com/q-dev/q-client/consensus/clique"
)

type Indexer struct {
	clique *clique.API
}

func New(clique *clique.API) *Indexer {
	return &Indexer{clique: clique}
}
