package main

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum"
	"gitlab.com/q-dev/q-client/cmd/ethereum/repl"
	"gitlab.com/q-dev/q-client/javascript"
	"gitlab.com/q-dev/q-client/utils"
)

func InitJsConsole(ethereum *eth.Ethereum) {
	repl := ethrepl.NewJSRepl(ethereum)
	go repl.Start()
	utils.RegisterInterrupt(func(os.Signal) {
		repl.Stop()
	})
}

func ExecJsFile(ethereum *eth.Ethereum, InputFile string) {
	file, err := os.Open(InputFile)
	if err != nil {
		logger.Fatalln(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Fatalln(err)
	}
	re := javascript.NewJSRE(ethereum)
	utils.RegisterInterrupt(func(os.Signal) {
		re.Stop()
	})
	re.Run(string(content))
}
