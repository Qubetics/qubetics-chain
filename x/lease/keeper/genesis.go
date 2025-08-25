package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/qubetics/qubetics-blockchain/v2/x/lease/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return v1.NewGenesisState(
		k.GetLeases(ctx),
		k.GetParams(ctx),
	)
}
