// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// erc20 events
const (
	EventTypeConvertERC20           = "convert_erc20"
	EventTypeRegisterERC20          = "register_erc20"
	EventTypeToggleTokenConversion  = "toggle_token_conversion" // #nosec
	EventTypeRegisterERC20Extension = "register_erc20_extension"

	AttributeCoinSourceChannel = "source_channel"
	AttributeKeyCosmosCoin     = "cosmos_coin"
	AttributeKeyERC20Token     = "erc20_token" // #nosec
	AttributeKeyReceiver       = "receiver"
)

// LogTransfer Event type for Transfer(address from, address to, uint256 value)
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
