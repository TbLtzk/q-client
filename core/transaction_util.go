package core

import (
	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/logger"
	"gitlab.com/q-dev/q-client/logger/glog"
	"gitlab.com/q-dev/q-client/rlp"
)

func PutTransactions(db common.Database, block *types.Block, txs types.Transactions) {
	for i, tx := range block.Transactions() {
		rlpEnc, err := rlp.EncodeToBytes(tx)
		if err != nil {
			glog.V(logger.Debug).Infoln("Failed encoding tx", err)
			return
		}
		db.Put(tx.Hash().Bytes(), rlpEnc)

		var txExtra struct {
			BlockHash  common.Hash
			BlockIndex uint64
			Index      uint64
		}
		txExtra.BlockHash = block.Hash()
		txExtra.BlockIndex = block.NumberU64()
		txExtra.Index = uint64(i)
		rlpMeta, err := rlp.EncodeToBytes(txExtra)
		if err != nil {
			glog.V(logger.Debug).Infoln("Failed encoding tx meta data", err)
			return
		}
		db.Put(append(tx.Hash().Bytes(), 0x0001), rlpMeta)
	}
}

func PutReceipts(db common.Database, hash common.Hash, receipts types.Receipts) error {
	storageReceipts := make([]*types.ReceiptForStorage, len(receipts))
	for i, receipt := range receipts {
		storageReceipts[i] = (*types.ReceiptForStorage)(receipt)
	}

	bytes, err := rlp.EncodeToBytes(storageReceipts)
	if err != nil {
		return err
	}

	db.Put(append(receiptsPre, hash[:]...), bytes)

	return nil
}
