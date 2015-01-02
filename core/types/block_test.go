package types

import (
	"bytes"
	"testing"

	"gitlab.com/q-dev/q-client/ethdb"
	"gitlab.com/q-dev/q-client/ethutil"
	"gitlab.com/q-dev/q-client/rlp"
)

func init() {
	ethutil.ReadConfig(".ethtest", "/tmp/ethtest", "")
	ethutil.Config.Db, _ = ethdb.NewMemDatabase()
}

func TestNewBlock(t *testing.T) {
	block := GenesisBlock()
	data := ethutil.Encode(block)

	var genesis Block
	err := rlp.Decode(bytes.NewReader(data), &genesis)
}
