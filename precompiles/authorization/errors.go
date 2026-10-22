// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package authorization

const (
	// ErrAuthzDoesNotExistOrExpired is raised when the authorization does not exist.
	ErrAuthzDoesNotExistOrExpired = "authorization to %s for address %s does not exist or is expired"
	// ErrEmptyMethods is raised when the given methods array is empty.
	ErrEmptyMethods = "no methods defined; expected at least one message type url"
	// ErrEmptyStringInMethods is raised when the given methods array contains an empty string.
	ErrEmptyStringInMethods = "empty string found in methods array; expected no empty strings to be passed; got: %v"
	// ErrExceededAllowance is raised when the amount exceeds the set allowance.
	ErrExceededAllowance = "amount %s greater than allowed limit %s"
	// ErrInvalidGranter is raised when the granter address is not valid.
	ErrInvalidGranter = "invalid granter address: %v"
	// ErrInvalidGrantee is raised when the grantee address is not valid.
	ErrInvalidGrantee = "invalid grantee address: %v"
	// ErrInvalidMethods is raised when the given methods cannot be unpacked.
	ErrInvalidMethods = "invalid methods defined; expected an array of strings; got: %v"
	// ErrInvalidMethod is raised when the given method cannot be unpacked.
	ErrInvalidMethod = "invalid method defined; expected a string; got: %v"
	// ErrAuthzNotAccepted is raised when the authorization is not accepted.
	ErrAuthzNotAccepted = "authorization to %s for address %s is not accepted"
)
