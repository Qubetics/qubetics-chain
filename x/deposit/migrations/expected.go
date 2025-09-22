package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/qubetics/qubetics-blockchain/v2/x/deposit/types/v1"
)

type DepositKeeper interface {
	SetDeposit(ctx sdk.Context, deposit v1.Deposit)
	Store(ctx sdk.Context) sdk.KVStore
}
