package validatorpool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/qubetics/qubetics-blockchain/v2/testutil/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/nullify"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ValidatorpoolKeeper(t)
	validatorpool.InitGenesis(ctx, *k, genesisState)
	got := validatorpool.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
