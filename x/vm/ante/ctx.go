// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package ante

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

// BuildEvmExecutionCtx builds the context needed prior to executing an EVM transaction.
// It does the following:
// 1. Sets an empty KV gas config for gas to be calculated by opcodes
// and not kvstore actions
// 2. Setup an empty transient KV gas config for transient gas to be
// calculated by opcodes
func BuildEvmExecutionCtx(ctx sdktypes.Context) sdktypes.Context {
	// We need to setup an empty gas config so that the gas is consistent with Ethereum.
	return ctx.WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{})
}
