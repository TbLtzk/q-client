package utils

import (
	"math/big"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
)

// SenderFromServer is a types.Signer that remembers the sender address returned by the RPC
// server. It is stored in the transaction's sender address cache to avoid an additional
// request in TransactionSender.
type SenderFromServer struct {
	Addr      common.Address
	Blockhash common.Hash
}

func SetSenderFromServer(tx *types.Transaction, addr common.Address, block common.Hash) {
	// Use types.Sender for side-effect to store our signer into the cache.
	types.Sender(&SenderFromServer{addr, block}, tx)
}

func (s *SenderFromServer) Equal(other types.Signer) bool {
	return true
}

func (s *SenderFromServer) Sender(tx *types.Transaction) (common.Address, error) {
	return s.Addr, nil
}

func (s *SenderFromServer) Hash(tx *types.Transaction) common.Hash {
	panic("can't sign with SenderFromServer")
}

func (s *SenderFromServer) SignatureValues(tx *types.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	panic("can't sign with SenderFromServer")
}

func (s *SenderFromServer) ChainID() *big.Int {
	panic("can't sign with SenderFromServer")
}
