package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/qubetics/qubetics-blockchain/v2/types"
	v1 "github.com/qubetics/qubetics-blockchain/v2/x/lease/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	// Validate the provided genesis state before mutating store
	if err := v1.ValidateGenesis(state); err != nil {
		panic(err)
	}

	// Initialize module parameters
	k.SetParams(ctx, state.Params)

	var maxID uint64
	for _, lease := range state.Leases {
		// Persist the lease
		k.SetLease(ctx, lease)

		// Build indexes
		provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
		if err != nil {
			panic(err)
		}
		nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
		if err != nil {
			panic(err)
		}

		k.SetLeaseForProvider(ctx, provAddr, lease.ID)
		k.SetLeaseForNodeByProvider(ctx, nodeAddr, provAddr, lease.ID)

		// Rebuild time-based queues. Helpers are no-ops on zero times
		k.SetLeaseForPayoutAt(ctx, lease.PayoutAt(), lease.ID)
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt(), lease.ID)
		k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

		if lease.ID > maxID {
			maxID = lease.ID
		}
	}

	// Initialize the lease counter to the highest existing ID to prevent collisions
	k.SetLeaseCount(ctx, maxID)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return v1.NewGenesisState(
		k.GetLeases(ctx),
		k.GetParams(ctx),
	)
}
