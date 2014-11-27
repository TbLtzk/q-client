package types

import (
	"math/big"
	"gitlab.com/q-dev/q-client/state"
)

type BlockProcessor interface {
	ProcessWithParent(*Block, *Block) (*big.Int, state.Messages, error)
}
