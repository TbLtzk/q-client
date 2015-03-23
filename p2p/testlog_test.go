package p2p

import (
	"testing"

	"gitlab.com/q-dev/q-client/logger"
)

type testLogger struct{ t *testing.T }

func testlog(t *testing.T) testLogger {
	logger.Reset()
	l := testLogger{t}
	logger.AddLogSystem(l)
	return l
}

func (l testLogger) LogPrint(msg logger.LogMsg) {
	l.t.Logf("%s", msg.String())
}

func (testLogger) detach() {
	logger.Flush()
	logger.Reset()
}
