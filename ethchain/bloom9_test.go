package ethchain

import (
	"testing"

	"gitlab.com/q-dev/q-client/vm"
)

func TestBloom9(t *testing.T) {
	testCase := []byte("testtest")
	bin := LogsBloom([]vm.Log{vm.Log{testCase, [][]byte{[]byte("hellohello")}, nil}}).Bytes()
	res := BloomLookup(bin, testCase)

	if !res {
		t.Errorf("Bloom lookup failed")
	}
}
