package ui

import "gitlab.com/q-dev/q-client/core/types"

type Interface interface {
	UnlockAccount(address []byte) bool
	ConfirmTransaction(tx *types.Transaction) bool
}
