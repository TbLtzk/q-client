package utils

import (
	"errors"
	"math/big"

	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/core/types"
)

// SenderFromServer is a types.Signer that remembers the sender address returned by the RPC
// server. It is stored in the transaction's sender address cache to avoid an additional
// request in TransactionSender.
type SenderFromServer struct {
	Addr      common.Address
	Blockhash common.Hash
}

var errNotCached = errors.New("sender not cached")

func SetSenderFromServer(tx *types.Transaction, addr common.Address, block common.Hash) {
	// Use types.Sender for side-effect to store our signer into the cache.
	types.Sender(&SenderFromServer{addr, block}, tx)
}

func (s *SenderFromServer) Equal(other types.Signer) bool {
	//os, ok := other.(*SenderFromServer)
	//return ok && os.blockhash == s.blockhash
	return true
}

func (s *SenderFromServer) Sender(tx *types.Transaction) (common.Address, error) {
	if s.Blockhash == (common.Hash{}) {
		return common.Address{}, errNotCached
	}
	return s.Addr, nil
}

func (s *SenderFromServer) Hash(tx *types.Transaction) common.Hash {
	panic("can't sign with SenderFromServer")
}
func (s *SenderFromServer) SignatureValues(tx *types.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	panic("can't sign with SenderFromServer")
}
