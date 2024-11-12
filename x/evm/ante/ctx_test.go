package ante_test

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/qubetics/qubetics-blockchain/v2/testutil/integration/qubetics/network"
	evmante "github.com/qubetics/qubetics-blockchain/v2/x/evm/ante"
)

func (suite *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New()

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	suite.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	suite.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
