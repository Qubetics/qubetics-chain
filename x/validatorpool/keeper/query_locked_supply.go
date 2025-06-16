package keeper

import (
	"context"

	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LockedSupply(goCtx context.Context, req *types.QueryLockedSupplyRequest) (*types.QueryLockedSupplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get module account balance as locked supply
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		return nil, status.Error(codes.NotFound, "module account not found")
	}

	balance := k.bankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), k.stakingKeeper.BondDenom(ctx))
	return &types.QueryLockedSupplyResponse{
		LockedSupply: balance.Amount,
	}, nil
}
