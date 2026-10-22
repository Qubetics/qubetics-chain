// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package qubeticsd

import (
	qubeticstypes "github.com/qubetics/qubetics-blockchain/v2/types"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
)

// WQubeticsContract is the address of the wrapped native token (WERC20
// precompile) for tics. It is the default native precompile address used by
// Cosmos EVM.
const WQubeticsContract = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517"

// QubeticsTokenPairs contains the token pair for the chain's native denom.
var QubeticsTokenPairs = []erc20types.TokenPair{
	{
		Erc20Address:  WQubeticsContract,
		Denom:         qubeticstypes.AttoQubetics,
		Enabled:       true,
		ContractOwner: erc20types.OWNER_MODULE,
	},
}
