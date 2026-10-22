// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/qubetics/qubetics-blockchain/v2/precompiles/common"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/core/vm"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	// EventTypeValidatorUnjailed defines the event type for validator unjailing
	EventTypeValidatorUnjailed = "ValidatorUnjailed"
)

// Add this struct after the existing constants
type EventValidatorUnjailed struct {
	Validator common.Address
}

// EmitValidatorUnjailedEvent emits the ValidatorUnjailed event
func (p Precompile) EmitValidatorUnjailedEvent(ctx sdk.Context, stateDB vm.StateDB, validator common.Address) error {
	// Prepare the event topics
	event := p.ABI.Events[EventTypeValidatorUnjailed]
	topics := make([]common.Hash, 2)

	// The first topic is always the signature of the event
	topics[0] = event.ID

	var err error
	topics[1], err = cmn.MakeTopic(validator)
	if err != nil {
		return err
	}

	stateDB.AddLog(&ethtypes.Log{
		Address:     p.Address(),
		Topics:      topics,
		BlockNumber: uint64(ctx.BlockHeight()), //nolint:gosec // G115
	})

	return nil
}
