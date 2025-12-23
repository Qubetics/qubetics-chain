package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	// Restore all nodes from genesis state
	for _, node := range state.Nodes {
		k.SetNode(ctx, node)
	}

	// Restore parameters
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	return v3.NewGenesisState(
		k.GetNodes(ctx),
		k.GetParams(ctx),
	)
}
