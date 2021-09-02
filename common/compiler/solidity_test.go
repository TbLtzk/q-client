// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package compiler

import (
	"os/exec"
	"strings"
	"testing"
)

const (
	testSource = `
pragma solidity >0.0.0;
contract test {
   /// @notice Will multiply ` + "`a`" + ` by 7.
   function multiply(uint a) public returns(uint d) {
       return a * 7;
   }
}
`
)

func skipWithoutSolc(t *testing.T) {
	if _, err := exec.LookPath("solc"); err != nil {
		t.Skip(err)
	}
}

func TestSolidityCompiler(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString("", testSource)
	if err != nil {
		t.Fatalf("error compiling source. result %v: %v", contracts, err)
	}
	if len(contracts) != 1 {
		t.Errorf("one contract expected, got %d", len(contracts))
	}
	c, ok := contracts["test"]
	if !ok {
		c, ok = contracts["<stdin>:test"]
		if !ok {
			t.Fatal("info for contract 'test' not present in result")
		}
	}
	if c.Code == "" {
		t.Error("empty code")
	}
	if c.Info.Source != testSource {
		t.Error("wrong source")
	}
	if c.Info.CompilerVersion == "" {
		t.Error("empty version")
	}
}

func TestSolidityCompileError(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString("", testSource[4:])
	if err == nil {
		t.Errorf("error expected compiling source. got none. result %v", contracts)
	}
	t.Logf("error: %v", err)
}

func TestParseTruffleJSON(t *testing.T) {
	abi, bin, err := ParseTruffleJSON([]byte(truffleTestData))
	if err != nil {
		t.Fatal("failed to parse truffle output", "err", err)
	}

	if len(*abi) == 0 {
		t.Error("empty abi")
	}

	if len(*bin) == 0 {
		t.Error("empty bytecode")
	}

	if strings.HasPrefix(*bin, "0x") {
		t.Error("bytecode has 0x prefix")
	}
}

const truffleTestData = `
{
  "contractName": "Bailout",
  "abi": [
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_registry",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "receiver",
          "type": "address"
        }
      ],
      "name": "SuccessTransfer",
      "type": "event"
    },
    {
      "stateMutability": "payable",
      "type": "receive"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "request",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ],
  "bytecode": "0x608060405234801561001057600080fd5b506040516104633803806104638339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506103cf806100946000396000f3fe6080604052600436106100225760003560e01c8063d845a4b31461002e57610029565b3661002957005b600080fd5b34801561003a57600080fd5b506100676004803603602081101561005157600080fd5b810190808035906020019092919050505061007f565b60405180821515815260200191505060405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fb902716040518163ffffffff1660e01b815260040180806020018281038252601c8152602001807f746f6b656e65636f6e6f6d6963732e73797374656d526573657276650000000081525060200191505060206040518083038186803b15801561012457600080fd5b505afa158015610138573d6000803e3d6000fd5b505050506040513d602081101561014e57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d836040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156101b157600080fd5b505af11580156101c5573d6000803e3d6000fd5b505050506040513d60208110156101db57600080fd5b810190808051906020019092919050505061025e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f52657475726e656420776974682066616c736520726573756c7400000000000081525060200191505060405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff168360405180600001905060006040518083038185875af1925050503d80600081146102be576040519150601f19603f3d011682016040523d82523d6000602084013e6102c3565b606091505b505090508061033a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f5472616e73666572206661696c6564000000000000000000000000000000000081525060200191505060405180910390fd5b7f6bdc666bbe47cadd8bfc4dbb947317b58ab01c4ecfd733e8ea0509e0a76afdbf8333604051808381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1600191505091905056fea2646970667358221220ab8d7dfe9274e2ce429e7eade1642953b39add5e9a3e8f0e3c28442d06c69b9264736f6c63430007010033"
}`
