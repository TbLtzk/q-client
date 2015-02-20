package qwhisper

import (
	"testing"

	"gitlab.com/q-dev/q-client/whisper"
)

func TestHasIdentity(t *testing.T) {
	qw := New(whisper.New())
	id := qw.NewIdentity()
	if !qw.HasIdentity(id) {
		t.Error("expected to have identity")
	}
}
