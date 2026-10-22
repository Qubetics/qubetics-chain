// Copyright Evmos Authors - Tharsis Labs Ltd - https://github.com/evmos/os/
// Copyright 2025, Interchain Labs - https://github.com/cosmos/evm/
// SPDX-License-Identifier: Apache-2.0
//
// Derived from cosmos/evm v0.1.0 (https://github.com/cosmos/evm), licensed under
// the Apache License, Version 2.0. Modified by Qubetics: module path renamed; backported to Cosmos SDK 0.47 / ibc-go v7.

package evm

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	anteinterfaces "github.com/qubetics/qubetics-blockchain/v2/ante/interfaces"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/statedb"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/vm/types"
	"github.com/ethereum/go-ethereum/common"
)

// VerifyAccountBalance checks that the account balance is greater than the total transaction cost.
// The account will be set to store if it doesn't exist, i.e. cannot be found on store.
// This method will fail if:
// - from address is NOT an EOA
// - account balance is lower than the transaction cost
func VerifyAccountBalance(
	ctx sdk.Context,
	accountKeeper anteinterfaces.AccountKeeper,
	account *statedb.Account,
	from common.Address,
	txData evmtypes.TxData,
) error {
	// Only EOA are allowed to send transactions.
	if account != nil && account.IsContract() {
		return errorsmod.Wrapf(
			errortypes.ErrInvalidType,
			"the sender is not EOA: address %s", from,
		)
	}

	if account == nil {
		acc := accountKeeper.NewAccountWithAddress(ctx, from.Bytes())
		accountKeeper.SetAccount(ctx, acc)
		account = statedb.NewEmptyAccount()
	}

	if err := keeper.CheckSenderBalance(sdkmath.NewIntFromBigInt(account.Balance), txData); err != nil {
		return errorsmod.Wrap(err, "failed to check sender balance")
	}

	return nil
}
