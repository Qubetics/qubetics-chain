package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v3 "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v3"
)

type SessionKeeper interface {
	SetParams(ctx sdk.Context, params v3.Params)
	Store(ctx sdk.Context) sdk.KVStore
}
