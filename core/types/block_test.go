package types

import (
	"fmt"
	"math/big"
	"testing"

	"gitlab.com/q-dev/q-client/common"
)

func TestConversion(t *testing.T) {
	var (
		parent   common.Hash
		coinbase common.Address
		hash     common.Hash
	)

	block := NewBlock(parent, coinbase, hash, big.NewInt(0), 0, "")
	fmt.Println(block)
}
