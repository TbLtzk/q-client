package indexer

import (
	"gitlab.com/q-dev/q-client/consensus/clique"
)

type Indexer struct {
	clique *clique.API
}

func New(cliqueApi *clique.API) *Indexer {
	return &Indexer{clique: cliqueApi}
}
