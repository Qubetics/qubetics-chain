// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/core/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// EmitEventArgs are the arguments required to emit an authorization event.
//
// The event type can be:
//   - ApprovalEvent
//   - GenericApprovalEvent
//   - AllowanceChangeEvent
//   - ...
type EmitEventArgs struct {
	Ctx            sdk.Context
	StateDB        vm.StateDB
	ContractAddr   common.Address
	ContractEvents map[string]abi.Event
	EventData      interface{}
}
