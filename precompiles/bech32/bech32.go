// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package bech32

import (
	"embed"
	"fmt"

	cmn "github.com/qubetics/qubetics-blockchain/v2/precompiles/common"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/core/vm"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/vm/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

// Precompile defines the precompiled contract for Bech32 encoding.
type Precompile struct {
	abi.ABI
	baseGas uint64
}

// NewPrecompile creates a new bech32 Precompile instance as a
// PrecompiledContract interface.
func NewPrecompile(baseGas uint64) (*Precompile, error) {
	newABI, err := cmn.LoadABI(f, "abi.json")
	if err != nil {
		return nil, err
	}

	if baseGas == 0 {
		return nil, fmt.Errorf("baseGas cannot be zero")
	}

	return &Precompile{
		ABI:     newABI,
		baseGas: baseGas,
	}, nil
}

// Address defines the address of the bech32 precompiled contract.
func (Precompile) Address() common.Address {
	return common.HexToAddress(evmtypes.Bech32PrecompileAddress)
}

// RequiredGas calculates the contract gas use.
func (p Precompile) RequiredGas(_ []byte) uint64 {
	return p.baseGas
}

// Run executes the precompiled contract bech32 methods defined in the ABI.
func (p Precompile) Run(_ *vm.EVM, contract *vm.Contract, _ bool) (bz []byte, err error) {
	// NOTE: This check avoid panicking when trying to decode the method ID
	if len(contract.Input) < 4 {
		return nil, vm.ErrExecutionReverted
	}

	methodID := contract.Input[:4]
	// NOTE: this function iterates over the method map and returns
	// the method with the given ID
	method, err := p.MethodById(methodID)
	if err != nil {
		return nil, err
	}

	argsBz := contract.Input[4:]
	args, err := method.Inputs.Unpack(argsBz)
	if err != nil {
		return nil, err
	}

	switch method.Name {
	case HexToBech32Method:
		bz, err = p.HexToBech32(method, args)
	case Bech32ToHexMethod:
		bz, err = p.Bech32ToHex(method, args)
	}

	if err != nil {
		return nil, err
	}

	return bz, nil
}
