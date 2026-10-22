// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

//go:build test
// +build test

// This file is used to allow the testing of EVM configuration initialization
// without the need to introduce testing requirements in the final binary. In
// this case, the file provides the possibility to restore the EIP activator
// functions to the initial state without the need to compile ResetActivators
// in the final binary.

package vm

var originalActivators = make(map[string]func(*JumpTable))

func init() {
	keys := GetActivatorsEipNames()

	originalActivators = make(map[string]func(*JumpTable), len(keys))

	for _, k := range keys {
		originalActivators[k] = activators[k]
	}
}

// ResetActivators resets activators to the original go ethereum activators map
func ResetActivators() {
	activators = make(map[string]func(*JumpTable))
	for k, v := range originalActivators {
		activators[k] = v
	}
}
