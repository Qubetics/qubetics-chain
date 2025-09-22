package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/qubetics/qubetics-blockchain/v2/x/lease/types/v1"
)

type LeaseKeeper interface {
	SetParams(ctx sdk.Context, params v1.Params)
}
