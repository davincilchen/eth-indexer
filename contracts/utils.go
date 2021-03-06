// Copyright 2018 The eth-indexer Authors
// This file is part of the eth-indexer library.
//
// The eth-indexer library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The eth-indexer library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the eth-indexer library. If not, see <http://www.gnu.org/licenses/>.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	erc20ABI, _ = abi.JSON(strings.NewReader(ERC20TokenABI))
)

// DecimalsMsg returns the erc20 decimal message
func DecimalsMsg(contractAddress common.Address) *ethereum.CallMsg {
	method := "decimals"
	input, _ := erc20ABI.Pack(method)
	return &ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	}
}

// DecodeDecimals decodes the erc20 decimal message
func DecodeDecimals(data []byte) (uint8, error) {
	method := "decimals"
	result := new(uint8)
	err := erc20ABI.Unpack(result, method, data)
	if err != nil {
		return 0, err
	}
	return *result, nil
}

// BalanceOfMsg returns the erc20 balanceOf message
func BalanceOfMsg(contractAddress common.Address, account common.Address) *ethereum.CallMsg {
	method := "balanceOf"
	input, _ := erc20ABI.Pack(method, account)
	return &ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	}
}

// DecodeBalanceOf decodes the erc20 balanceOf message
func DecodeBalanceOf(data []byte) (*big.Int, error) {
	method := "balanceOf"
	result := new(*big.Int)
	err := erc20ABI.Unpack(result, method, data)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
