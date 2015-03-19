package types

import (
	"math/big"

	"gitlab.com/q-dev/q-client/state"
)

type BlockProcessor interface {
	Process(*Block) (*big.Int, state.Logs, error)
}
