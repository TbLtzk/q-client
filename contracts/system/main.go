package system

import (
	"math/big"

	"gitlab.com/q-dev/go-ethereum/contracts/validators/contract"

	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/core/types"
	"gitlab.com/q-dev/go-ethereum/crypto"

	"gitlab.com/q-dev/go-ethereum/ethclient"

	"gitlab.com/q-dev/go-ethereum/cmd/utils"

	"gitlab.com/q-dev/go-ethereum/log"
)

// address 532c69263800e1f1cdb72acae555a85864146986
const privateKeyHex = "e4160a6c21bc6c70a93f07c42c540fcaaf4ab973b75e12a6dbe2b415a933cd6a"

type deployFunc func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, *contract.Validators, error)

var contracts = []deployFunc{contract.DeployValidators}

func DeploySystemContracts(client *ethclient.Client) {

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		utils.Fatalf("Failed to decode private key %v", err)
	}

	// TODO add checking of deployed contracts (skip already deployed)

	// setup clef signer, create an abigen transactor and an RPC client
	transactor := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(privateKey.PublicKey),
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction,
		) (*types.Transaction, error) {
			return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
		},
		Nonce: big.NewInt(0),
	}

	for _, sysContract := range contracts {
		address, tx, _, err := sysContract(transactor, client)
		if err != nil {
			log.Warn("Failed deploy to system contract", "error", err)
		}
		log.Info("Deployed system contract", "address", address, "tx", tx.Hash().Hex())
		// TODO save somewhere addresses of created contracts (global var?)

		transactor.Nonce.Add(transactor.Nonce, big.NewInt(1))
	}

}
