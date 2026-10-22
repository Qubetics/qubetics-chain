// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package ics20

const (
	// ErrInvalidSourcePort is raised when the source port is invalid.
	ErrInvalidSourcePort = "invalid source port"
	// ErrInvalidSourceChannel is raised when the source channel is invalid.
	ErrInvalidSourceChannel = "invalid source port"
	// ErrInvalidSender is raised when the sender is invalid.
	ErrInvalidSender = "invalid sender: %s"
	// ErrInvalidReceiver is raised when the receiver is invalid.
	ErrInvalidReceiver = "invalid receiver: %s"
	// ErrInvalidTimeoutTimestamp is raised when the timeout timestamp is invalid.
	ErrInvalidTimeoutTimestamp = "invalid timeout timestamp: %d"
	// ErrInvalidMemo is raised when the memo is invalid.
	ErrInvalidMemo = "invalid memo: %s"
	// ErrInvalidHash is raised when the hash is invalid.
	ErrInvalidHash = "invalid hash: %s"
	// ErrNoMatchingAllocation is raised when no matching allocation is found.
	ErrNoMatchingAllocation = "no matching allocation found for source port: %s, source channel: %s, and denom: %s"
	// ErrDifferentOriginFromSender is raised when the origin address is not the same as the sender address.
	ErrDifferentOriginFromSender = "origin address %s is not the same as sender address %s"
	// ErrTraceNotFound is raised when the denom trace for the specified request does not exist.
	ErrTraceNotFound = "denomination trace not found"
)
