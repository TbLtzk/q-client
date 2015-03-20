package whisper

import (
	"testing"

	"gitlab.com/q-dev/q-client/common"
)

func TestSorting(t *testing.T) {
	m := map[int32]common.Hash{
		1: {1},
		3: {3},
		2: {2},
		5: {5},
	}
	exp := []int32{1, 2, 3, 5}
	res := sortKeys(m)
	for i, k := range res {
		if k != exp[i] {
			t.Error(k, "failed. Expected", exp[i])
		}
	}
}
