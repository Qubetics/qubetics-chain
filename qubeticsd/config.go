// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package qubeticsd

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	qubeticstypes "github.com/qubetics/qubetics-blockchain/v2/types"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/vm/types"
)

// EVMOptionsFn defines a function type for setting app options specifically for
// the Cosmos EVM app. The function should receive the chainID and return an error if
// any.
type EVMOptionsFn func(string) error

// NoOpEVMOptions is a no-op function that can be used when the app does not
// need any specific configuration.
func NoOpEVMOptions(_ string) error {
	return nil
}

var sealed = false

// QubeticsCoinInfo is the EVM coin of the Qubetics chain: tics, with 18 decimals.
var QubeticsCoinInfo = evmtypes.EvmCoinInfo{
	Denom:        qubeticstypes.AttoQubetics,
	DisplayDenom: qubeticstypes.DisplayDenom,
	Decimals:     evmtypes.EighteenDecimals,
}

// EvmAppOptions sets up the global Cosmos EVM configuration for the given
// chain ID (for example "qubetics_9087-1"), which also provides the EIP-155
// chain ID used by the EVM.
func EvmAppOptions(chainID string) error {
	if sealed {
		return nil
	}

	if err := setBaseDenom(QubeticsCoinInfo); err != nil {
		return err
	}

	baseDenom, err := sdk.GetBaseDenom()
	if err != nil {
		return err
	}

	ethCfg := evmtypes.DefaultChainConfig(chainID)

	err = evmtypes.NewEVMConfigurator().
		WithChainConfig(ethCfg).
		WithEVMCoinInfo(baseDenom, uint8(QubeticsCoinInfo.Decimals)).
		Configure()
	if err != nil {
		return err
	}

	sealed = true
	return nil
}

// setBaseDenom registers the display denom and base denom and sets the
// base denom for the chain.
func setBaseDenom(ci evmtypes.EvmCoinInfo) error {
	if err := sdk.RegisterDenom(ci.DisplayDenom, math.LegacyOneDec()); err != nil {
		return err
	}

	// sdk.RegisterDenom will automatically overwrite the base denom when the
	// new setBaseDenom() are lower than the current base denom's units.
	return sdk.RegisterDenom(ci.Denom, math.LegacyNewDecWithPrec(1, int64(ci.Decimals)))
}
