// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package distribution

const (
	// ErrSetWithdrawAddrAuth is raised when no authorization to set the withdraw address exists.
	ErrSetWithdrawAddrAuth = "set withdrawer address authorization for address %s does not exist"
	// ErrWithdrawDelRewardsAuth is raised when no authorization to withdraw delegation rewards exists.
	ErrWithdrawDelRewardsAuth = "withdraw delegation rewards authorization for address %s does not exist"
	// ErrWithdrawValCommissionAuth is raised when no authorization to withdraw validator commission exists.
	ErrWithdrawValCommissionAuth = "withdraw validator commission authorization for address %s does not exist"
	// ErrDifferentValidator is raised when the origin address is not the same as the validator address.
	ErrDifferentValidator = "origin address %s is not the same as validator address %s"
)
