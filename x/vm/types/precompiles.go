// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package types

const (
	P256PrecompileAddress   = "0x0000000000000000000000000000000000000100"
	Bech32PrecompileAddress = "0x0000000000000000000000000000000000000400"
)

const (
	StakingPrecompileAddress      = "0x0000000000000000000000000000000000000800"
	DistributionPrecompileAddress = "0x0000000000000000000000000000000000000801"
	ICS20PrecompileAddress        = "0x0000000000000000000000000000000000000802"
	VestingPrecompileAddress      = "0x0000000000000000000000000000000000000803"
	BankPrecompileAddress         = "0x0000000000000000000000000000000000000804"
	GovPrecompileAddress          = "0x0000000000000000000000000000000000000805"
	SlashingPrecompileAddress     = "0x0000000000000000000000000000000000000806"
	EvidencePrecompileAddress     = "0x0000000000000000000000000000000000000807"
)

// AvailableStaticPrecompiles defines the full list of all available EVM extension addresses.
//
// NOTE: To be explicit, this list does not include the dynamically registered EVM extensions
// like the ERC-20 extensions.
var AvailableStaticPrecompiles = []string{
	P256PrecompileAddress,
	Bech32PrecompileAddress,
	StakingPrecompileAddress,
	DistributionPrecompileAddress,
	ICS20PrecompileAddress,
	VestingPrecompileAddress,
	BankPrecompileAddress,
	GovPrecompileAddress,
	SlashingPrecompileAddress,
	EvidencePrecompileAddress,
}
