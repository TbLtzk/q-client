package governance

import (
	"fmt"
	"testing"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/crypto"

	"github.com/pkg/errors"
	"gitlab.com/q-dev/q-client/params"
)

func TestRoot(t *testing.T) {
	path := "gov"
	db, err := newDatabase(path)
	if err != nil {
		t.Fatal(err)
	}

	list := &params.DevnetRootNodes
	list.Nodes = append(list.Nodes, common.BytesToAddress(crypto.Keccak256Hash([]byte("yoba")).Bytes()))

	set, _ := newRootSet(list)
	db.saveCurrentRootSet(set)

	got, err := db.getCurrentRootSet()
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to get set"))
	}

	fmt.Println(got.hash.Hex())
}
