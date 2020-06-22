package system

import (
	"math/big"

	"gitlab.com/q-dev/go-ethereum/contracts/validators/contract"

	"gitlab.com/q-dev/go-ethereum/accounts/abi/bind"
	"gitlab.com/q-dev/go-ethereum/common"
	"gitlab.com/q-dev/go-ethereum/core/types"
	"gitlab.com/q-dev/go-ethereum/crypto"

	"gitlab.com/q-dev/go-ethereum/ethclient"

	"gitlab.com/q-dev/go-ethereum/log"
)

// address 532c69263800e1f1cdb72acae555a85864146986
const privateKeyHex = "e4160a6c21bc6c70a93f07c42c540fcaaf4ab973b75e12a6dbe2b415a933cd6a"

type deployAndSaveFunc func(*bind.TransactOpts, bind.ContractBackend) (common.Address, *types.Transaction, error)

var contracts = []deployAndSaveFunc{deployAndSaveValidators}
var ValidatorContractAddress common.Address

func DeploySystemContracts(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Error("Failed to decode private key")
		panic(err)
	}

	// TODO add checking of deployed contracts (skip already deployed)

	// setup clef signer, create an abigen transactor and an RPC client
	transactor := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(privateKey.PublicKey),
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction,
		) (*types.Transaction, error) {
			return types.SignTx(tx, types.HomesteadSigner{}, privateKey)
		},
		Nonce:    big.NewInt(0),
		GasLimit: 123456791011, // Just to make it works on dev net
	}

	for _, sysContract := range contracts {
		address, tx, err := sysContract(transactor, client)
		if err != nil {
			log.Warn("Failed deploy to system contract", "error", err)
		} else {
			log.Info("System contract is deployed ", "address", address, "tx", tx.Hash().Hex())
		}

		transactor.Nonce.Add(transactor.Nonce, big.NewInt(1))
	}

}

func deployAndSaveValidators(transactor *bind.TransactOpts, backend bind.ContractBackend,
) (common.Address, *types.Transaction, error) {
	address, tx, _, err := contract.DeployValidators(transactor, backend)
	ValidatorContractAddress = address

	return address, tx, err
}
