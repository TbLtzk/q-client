package helper

import (
	"log"
	"os"

	"gitlab.com/q-dev/q-client/ethlog"
	"gitlab.com/q-dev/q-client/ethutil"
)

var Logger ethlog.LogSystem
var Log = ethlog.NewLogger("TEST")

func init() {
	Logger = ethlog.NewStdLogSystem(os.Stdout, log.LstdFlags, ethlog.InfoLevel)
	ethlog.AddLogSystem(Logger)

	ethutil.ReadConfig(".ethtest", "/tmp/ethtest", "")
}
