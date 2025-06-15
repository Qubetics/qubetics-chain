package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	ctx, k, accountKeeper, bankKeeper, _, _ := setupTest(t.(*testing.T))

	// Set up initial balance for the creator
	creatorAddr := sdk.AccAddress([]byte("test-address-12345678901234567890"))
	initialBalance := sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000000)))
	bankKeeper.balances[creatorAddr.String()] = initialBalance

	// Set up module account
	moduleAcc := accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	require.NotNil(t, moduleAcc)

	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer_AddToPool(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)

	// Create a valid test address
	creatorAddr := sdk.AccAddress([]byte("test-address-12345678901234567890"))

	tests := []struct {
		name    string
		msg     *types.MsgAddToPool
		wantErr bool
	}{
		{
			name: "successful add to pool",
			msg: &types.MsgAddToPool{
				Creator: creatorAddr.String(),
				Amount:  "1000000stake",
			},
			wantErr: false,
		},
		{
			name: "invalid creator address",
			msg: &types.MsgAddToPool{
				Creator: "invalid-address",
				Amount:  "1000000stake",
			},
			wantErr: true,
		},
		{
			name: "invalid amount format",
			msg: &types.MsgAddToPool{
				Creator: creatorAddr.String(),
				Amount:  "invalid-amount",
			},
			wantErr: true,
		},
		{
			name: "multiple coins not allowed",
			msg: &types.MsgAddToPool{
				Creator: creatorAddr.String(),
				Amount:  "1000000stake,2000000token",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ms.AddToPool(ctx, tt.msg)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
			}
		})
	}
}
