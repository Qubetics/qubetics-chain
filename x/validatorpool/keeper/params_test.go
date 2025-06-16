package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/qubetics/qubetics-blockchain/v2/testutil/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ValidatorpoolKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
