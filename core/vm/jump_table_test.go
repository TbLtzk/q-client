package vm

import (
	"math/big"
	"testing"

	"gitlab.com/q-dev/q-client/params"
)

func TestInit(t *testing.T) {
	params.HomesteadBlock = big.NewInt(1)

	jumpTable := newJumpTable(big.NewInt(0))
	if jumpTable[DELEGATECALL].valid {
		t.Error("Expected DELEGATECALL not to be present")
	}

	for _, n := range []int64{1, 2, 100} {
		jumpTable := newJumpTable(big.NewInt(n))
		if !jumpTable[DELEGATECALL].valid {
			t.Error("Expected DELEGATECALL to be present for block", n)
		}
	}
}
