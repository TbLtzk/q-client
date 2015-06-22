package api

import (
	"gitlab.com/q-dev/q-client/rpc/shared"
)

// Merge multiple API's to a single API instance
func Merge(apis ...shared.EthereumApi) shared.EthereumApi {
	return newMergedApi(apis...)
}
