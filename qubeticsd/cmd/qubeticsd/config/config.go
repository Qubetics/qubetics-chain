// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	qubeticstypes "github.com/qubetics/qubetics-blockchain/v2/types"
)

const (
	// Bech32Prefix defines the Bech32 prefix used for accounts on the Qubetics chain.
	Bech32Prefix = qubeticstypes.Bech32MainPrefix

	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = qubeticstypes.DisplayDenom
	// BaseDenom defines to the default denomination used on the Qubetics chain.
	BaseDenom = qubeticstypes.AttoQubetics
	// BaseDenomUnit defines the precision of the base denomination.
	BaseDenomUnit = qubeticstypes.BaseDenomUnit
)

// SetBech32Prefixes sets the global prefixes used when serializing addresses and
// public keys to Bech32 strings, including the dVPN provider and node prefixes.
func SetBech32Prefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(qubeticstypes.Bech32PrefixAccAddr, qubeticstypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(qubeticstypes.Bech32PrefixValAddr, qubeticstypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(qubeticstypes.Bech32PrefixConsAddr, qubeticstypes.Bech32PrefixConsPub)

	qubeticsConfig := qubeticstypes.GetConfig()
	qubeticsConfig.SetBech32PrefixForProvider(qubeticstypes.Bech32PrefixProvAddr, qubeticstypes.Bech32PrefixProvPub)
	qubeticsConfig.SetBech32PrefixForNode(qubeticstypes.Bech32PrefixNodeAddr, qubeticstypes.Bech32PrefixNodePub)
}

// SetBip44CoinType sets the Ethereum coin type (60) so keys derive the same
// addresses as Ethereum wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(qubeticstypes.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)                          // Shared
	config.SetFullFundraiserPath(qubeticstypes.BIP44HDPath) //nolint: staticcheck
}
