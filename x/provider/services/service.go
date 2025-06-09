package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/qubetics/qubetics-blockchain/v2/x/provider/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/provider/services/v2"
	"github.com/qubetics/qubetics-blockchain/v2/x/provider/services/v3"
	v2types "github.com/qubetics/qubetics-blockchain/v2/x/provider/types/v2"
	v3types "github.com/qubetics/qubetics-blockchain/v2/x/provider/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))
	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}
