// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package staking

const (
	// ErrDecreaseAmountTooBig is raised when the amount by which the allowance should be decreased is greater
	// than the authorization limit.
	ErrDecreaseAmountTooBig = "amount by which the allowance should be decreased is greater than the authorization limit: %s > %s"
	// ErrDifferentOriginFromDelegator is raised when the origin address is not the same as the delegator address.
	ErrDifferentOriginFromDelegator = "origin address %s is not the same as delegator address %s"
	// ErrNoDelegationFound is raised when no delegation is found for the given delegator and validator addresses.
	ErrNoDelegationFound = "delegation with delegator %s not found for validator %s"
	// ErrDifferentOriginFromValidator is raised when the origin address is not the same as the validator address.
	ErrDifferentOriginFromValidator = "origin address %s is not the same as validator operator address %s"
	// ErrCannotCallFromContract is raised when a function cannot be called from a smart contract.
	ErrCannotCallFromContract = "this method can only be called directly to the precompile, not from a smart contract"
)
