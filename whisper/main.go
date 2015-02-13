// +build none

package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/q-dev/q-client/crypto/secp256k1"
	"gitlab.com/q-dev/q-client/logger"
	"gitlab.com/q-dev/q-client/p2p"
	"gitlab.com/q-dev/q-client/whisper"
)

func main() {
	logger.AddLogSystem(logger.NewStdLogSystem(os.Stdout, log.LstdFlags, logger.InfoLevel))

	pub, _ := secp256k1.GenerateKeyPair()

	whisper := whisper.New()

	srv := p2p.Server{
		MaxPeers:   10,
		Identity:   p2p.NewSimpleClientIdentity("whisper-go", "1.0", "", string(pub)),
		ListenAddr: ":30300",
		NAT:        p2p.UPNP(),

		Protocols: []p2p.Protocol{whisper.Protocol()},
	}
	if err := srv.Start(); err != nil {
		fmt.Println("could not start server:", err)
		os.Exit(1)
	}

	select {}
}
