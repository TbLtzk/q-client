package helper

import (
	"log"
	"os"

	logpkg "gitlab.com/q-dev/q-client/logger"
)

var Logger logpkg.LogSystem
var Log = logpkg.NewLogger("TEST")

func init() {
	Logger = logpkg.NewStdLogSystem(os.Stdout, log.LstdFlags, logpkg.InfoLevel)
	logpkg.AddLogSystem(Logger)
}
