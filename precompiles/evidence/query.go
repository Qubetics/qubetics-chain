// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package evidence

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	cmn "github.com/qubetics/qubetics-blockchain/v2/precompiles/common"
)

// Evidence implements the query logic for getting evidence by hash.
func (p *Precompile) Evidence(
	ctx sdk.Context,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 1, len(args))
	}

	evidenceHash, ok := args[0].([]byte)
	if !ok {
		return nil, errors.New(ErrInvalidEvidenceHash)
	}

	queryServer := p.evidenceKeeper
	res, err := queryServer.Evidence(ctx, &evidencetypes.QueryEvidenceRequest{
		Hash: strings.ToUpper(hex.EncodeToString(evidenceHash)),
	})
	if err != nil {
		return nil, err
	}

	// Convert the Any type to Equivocation
	equivocation, ok := res.Evidence.GetCachedValue().(*evidencetypes.Equivocation)
	if !ok {
		return nil, errors.New(ErrExpectedEquivocation)
	}

	// Convert to our Equivocation struct
	evidence := EquivocationData{
		Height:           equivocation.Height,
		Time:             uint64(equivocation.Time.Unix()), //nolint:gosec // G115
		Power:            equivocation.Power,
		ConsensusAddress: equivocation.ConsensusAddress,
	}

	return method.Outputs.Pack(evidence)
}

// GetAllEvidence implements the query logic for getting all evidence.
func (p *Precompile) GetAllEvidence(
	ctx sdk.Context,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, 1, len(args))
	}

	pageRequest, ok := args[0].(*query.PageRequest)
	if !ok {
		return nil, fmt.Errorf("invalid page request")
	}

	queryServer := p.evidenceKeeper
	res, err := queryServer.AllEvidence(ctx, &evidencetypes.QueryAllEvidenceRequest{
		Pagination: pageRequest,
	})
	if err != nil {
		return nil, err
	}

	evidenceList := make([]EquivocationData, len(res.Evidence))
	for i, evidence := range res.Evidence {
		equivocation, ok := evidence.GetCachedValue().(*evidencetypes.Equivocation)
		if !ok {
			return nil, fmt.Errorf("invalid evidence type at index %d: expected Equivocation", i)
		}

		evidenceList[i] = EquivocationData{
			Height:           equivocation.Height,
			Time:             uint64(equivocation.Time.Unix()), //nolint:gosec // G115
			Power:            equivocation.Power,
			ConsensusAddress: equivocation.ConsensusAddress,
		}
	}

	return method.Outputs.Pack(evidenceList, res.Pagination)
}
