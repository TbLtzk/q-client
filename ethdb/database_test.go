package ethdb

import (
	"os"
	"path"

	"gitlab.com/q-dev/q-client/common"
)

func newDb() *LDBDatabase {
	file := path.Join("/", "tmp", "ldbtesttmpfile")
	if common.FileExist(file) {
		os.RemoveAll(file)
	}

	db, _ := NewLDBDatabase(file)

	return db
}
