package accounts

import (
	"testing"

	"gitlab.com/q-dev/q-client/crypto"
	"gitlab.com/q-dev/q-client/crypto/randentropy"
)

func TestAccountManager(t *testing.T) {
	ks := crypto.NewKeyStorePlain(crypto.DefaultDataDir())
	am := NewAccountManager(ks)
	pass := "" // not used but required by API
	a1, err := am.NewAccount(pass)
	toSign := randentropy.GetEntropyCSPRNG(32)
	_, err = am.Sign(a1, pass, toSign)
	if err != nil {
		t.Fatal(err)
	}
}
