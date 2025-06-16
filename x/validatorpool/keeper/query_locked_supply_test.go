package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
	"github.com/stretchr/testify/require"
)

func TestLockedSupplyQuery(t *testing.T) {
	ctx, k, mockAccountKeeper, mockBankKeeper, _, _ := setupTest(t)
	wctx := sdk.WrapSDKContext(ctx)

	// Set up module account
	moduleAcc := mockAccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	require.NotNil(t, moduleAcc)

	// Set up initial balance for the module account
	initialBalance := sdk.NewCoins(sdk.NewCoin("tics", sdk.NewInt(1000)))
	mockBankKeeper.balances[moduleAcc.GetAddress().String()] = initialBalance

	// Test case 1: Valid request
	response, err := k.LockedSupply(wctx, &types.QueryLockedSupplyRequest{})
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, sdk.NewInt(1000), response.LockedSupply)

	// Test case 2: Nil request
	_, err = k.LockedSupply(wctx, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid request")
}
