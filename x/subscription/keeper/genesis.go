package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v3 "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	// Restore all subscriptions from genesis state
	for _, subscription := range state.Subscriptions {
		k.SetSubscription(ctx, subscription)
	}

	// Restore all allocations from genesis state
	for _, allocation := range state.Allocations {
		k.SetAllocation(ctx, allocation)
	}

	// Restore parameters
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	return v3.NewGenesisState(
		nil,
		k.GetSubscriptions(ctx),
		k.GetParams(ctx),
	)
}
