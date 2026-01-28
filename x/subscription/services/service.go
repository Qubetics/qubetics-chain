package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/qubetics/qubetics-blockchain/v2/x/subscription/keeper"
	v2 "github.com/qubetics/qubetics-blockchain/v2/x/subscription/services/v2"
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/subscription/services/v3"
	v2types "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v2"
	v3types "github.com/qubetics/qubetics-blockchain/v2/x/subscription/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))
	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}
