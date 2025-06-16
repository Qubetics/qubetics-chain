package keeper

import (
	"context"

	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddToPool(goCtx context.Context, msg *types.MsgAddToPool) (*types.MsgAddToPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Parse creator address
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Parse coin string to get amount
	coins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount format: %s", err)
	}
	if len(coins) != 1 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "expected exactly one coin, got %d", len(coins))
	}
	amount := coins[0].Amount

	// Add tokens to the pool
	if err := k.Keeper.AddToPool(ctx, creator, amount); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to add tokens to pool")
	}

	return &types.MsgAddToPoolResponse{}, nil
}
