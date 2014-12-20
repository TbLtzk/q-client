package types

import (
	"math/big"

	"gitlab.com/q-dev/q-client/state"
	"gitlab.com/q-dev/q-client/wire"
)

type BlockProcessor interface {
	Process(*Block) (*big.Int, state.Messages, error)
}

type Broadcaster interface {
	Broadcast(wire.MsgType, []interface{})
}
