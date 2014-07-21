package main

import (
	"github.com/ethereum/eth-go"
	"gitlab.com/q-dev/q-client/ethereum/repl"
	"gitlab.com/q-dev/q-client/utils"
	"io/ioutil"
	"os"
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
	re := ethrepl.NewJSRE(ethereum)
	utils.RegisterInterrupt(func(os.Signal) {
		re.Stop()
	})
	re.Run(string(content))
}
