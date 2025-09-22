package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/qubetics/qubetics-blockchain/v2/x/node/keeper"
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/services/v3"
	v3types "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}
