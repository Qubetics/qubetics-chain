// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/core/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	// GetSigningInfoMethod defines the ABI method name for the slashing SigningInfo query
	GetSigningInfoMethod = "getSigningInfo"
	// GetSigningInfosMethod defines the ABI method name for the slashing SigningInfos query
	GetSigningInfosMethod = "getSigningInfos"
	// GetParamsMethod defines the ABI method name for the slashing Params query
	GetParamsMethod = "getParams"
)

// GetSigningInfo implements the query to get a validator's signing info.
func (p *Precompile) GetSigningInfo(
	ctx sdk.Context,
	method *abi.Method,
	_ *vm.Contract,
	args []interface{},
) ([]byte, error) {
	req, err := ParseSigningInfoArgs(args)
	if err != nil {
		return nil, err
	}

	res, err := p.slashingKeeper.SigningInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	out := new(SigningInfoOutput).FromResponse(res)
	return method.Outputs.Pack(out.SigningInfo)
}

// GetSigningInfos implements the query to get signing info for all validators.
func (p *Precompile) GetSigningInfos(
	ctx sdk.Context,
	method *abi.Method,
	_ *vm.Contract,
	args []interface{},
) ([]byte, error) {
	req, err := ParseSigningInfosArgs(method, args)
	if err != nil {
		return nil, err
	}

	res, err := p.slashingKeeper.SigningInfos(ctx, req)
	if err != nil {
		return nil, err
	}

	out := new(SigningInfosOutput).FromResponse(res)
	return method.Outputs.Pack(out.SigningInfos, out.PageResponse)
}

// GetParams implements the query to get the slashing parameters.
func (p *Precompile) GetParams(
	ctx sdk.Context,
	method *abi.Method,
	_ *vm.Contract,
	_ []interface{},
) ([]byte, error) {
	res, err := p.slashingKeeper.Params(ctx, &types.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}

	out := new(ParamsOutput).FromResponse(res)
	return method.Outputs.Pack(out.Params)
}
