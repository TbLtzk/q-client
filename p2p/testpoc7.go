// +build none

package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gitlab.com/q-dev/q-client/crypto/secp256k1"
	"gitlab.com/q-dev/q-client/logger"
	"gitlab.com/q-dev/q-client/p2p"
)

func main() {
	logger.AddLogSystem(logger.NewStdLogSystem(os.Stdout, log.LstdFlags, logger.DebugLevel))

	pub, _ := secp256k1.GenerateKeyPair()
	srv := p2p.Server{
		MaxPeers:   10,
		Identity:   p2p.NewSimpleClientIdentity("test", "1.0", "", string(pub)),
		ListenAddr: ":30303",
		NAT:        p2p.PMP(net.ParseIP("10.0.0.1")),
	}
	if err := srv.Start(); err != nil {
		fmt.Println("could not start server:", err)
		os.Exit(1)
	}

	// add seed peers
	seed, err := net.ResolveTCPAddr("tcp", "poc-7.ethdev.com:30303")
	if err != nil {
		fmt.Println("couldn't resolve:", err)
		os.Exit(1)
	}
	srv.SuggestPeer(seed.IP, seed.Port, nil)

	select {}
}
