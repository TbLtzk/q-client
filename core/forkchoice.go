// Copyright 2021 The go-ethereum Authors
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

package core

import (
	"errors"
	"math/big"

	"gitlab.com/q-dev/q-client/common"
	"gitlab.com/q-dev/q-client/common/math"
	"gitlab.com/q-dev/q-client/core/types"
	"gitlab.com/q-dev/q-client/log"
	"gitlab.com/q-dev/q-client/params"
)

// ChainReader defines a small collection of methods needed to access the local
// blockchain during header verification. It's implemented by both blockchain
// and lightchain.
type ChainReader interface {
	// Config retrieves the header chain's chain configuration.
	Config() *params.ChainConfig

	// GetTd returns the total difficulty of a local block.
	GetTd(common.Hash, uint64) *big.Int
}

// ForkChoice is the fork chooser based on the highest total difficulty of the
// chain(the fork choice used in the eth1) and the external fork choice (the fork
// choice used in the eth2). This main goal of this ForkChoice is not only for
// offering fork choice during the eth1/2 merge phase, but also keep the compatibility
// for all other proof-of-work networks.
type ForkChoice struct {
	chain ChainReader

	// preserve is a helper function used in td fork choice.
	// Miners will prefer to choose the local mined block if the
	// local td is equal to the extern one. It can be nil for light
	// client

	//As we use Clique as engine, we need to obey rule #3 of Eip3436
	//Choose the block whose validator had the least recent in-turn block assignment
	//This function is introduced because original preserve doesn't know anything about the external header
	preserve func(header *types.Header, externalHeader *types.Header) bool
}

func NewForkChoice(chainReader ChainReader, preserve func(header *types.Header, externalHeader *types.Header) bool) *ForkChoice {
	return &ForkChoice{
		chain:    chainReader,
		preserve: preserve,
	}
}

// ReorgNeeded returns whether the reorg should be applied
// based on the given external header and local canonical chain.
// In the td mode, the new head is chosen if the corresponding
// total difficulty is higher. In the extern mode, the trusted
// header is always selected as the head.
func (f *ForkChoice) ReorgNeeded(currentHeader *types.Header, externalHeader *types.Header) (bool, error) {
	var (
		localTD  = f.chain.GetTd(currentHeader.Hash(), currentHeader.Number.Uint64())
		externTd = f.chain.GetTd(externalHeader.Hash(), externalHeader.Number.Uint64())
	)
	if localTD == nil || externTd == nil {
		return false, errors.New("missing td")
	}
	// Accept the new header as the chain head if the transition
	// is already triggered. We assume all the headers after the
	// transition come from the trusted consensus layer.
	if ttd := f.chain.Config().TerminalTotalDifficulty; ttd != nil && ttd.Cmp(externTd) <= 0 {
		return true, nil
	}

	// If the total difficulty is higher than our known, add it to the canonical chain
	if diff := externTd.Cmp(localTD); diff > 0 {
		return true, nil
	} else if diff < 0 {
		return false, nil
	}
	// Local and external difficulty is identical.
	// Second clause in the if statement reduces the vulnerability to selfish mining.
	// Please refer to http://www.cs.cornell.edu/~ie53/publications/btcProcFC.pdf
	reorg := externTd.Cmp(localTD) > 0
	if !reorg && externTd.Cmp(localTD) == 0 {
		externalNum, currentNum := externalHeader.Number.Uint64(), currentHeader.Number.Uint64()
		if externalNum < currentNum {
		reorg = true
		} else if externalNum == currentNum {
			//Preserve check is modified in order to attempt of applying rule#3 from https://eips.ethereum.org/EIPS/eip-3436
			//If header numbers are the same, then choose the block whose validator had the least recent in-turn block assignment
		var currentPreserve, externPreserve bool
		if f.preserve != nil {
				currentPreserve, externPreserve = f.preserve(currentHeader, externalHeader), f.preserve(externalHeader, currentHeader)
		}
			//If both headers are from the same validator, then choose the block with the lower hash
			if currentPreserve && externPreserve {
				// EIP-3436 rule #4
				// Apply external chain if it has the lower hash
				reorg = externalHeader.Hash().Big().Cmp(currentHeader.Hash().Big()) < 0
			} else {
				reorg = false
			}
	}
	return reorg, nil
}
