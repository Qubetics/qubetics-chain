package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	deposit "github.com/qubetics/qubetics-blockchain/v2/x/deposit/services"
	lease "github.com/qubetics/qubetics-blockchain/v2/x/lease/services"
	node "github.com/qubetics/qubetics-blockchain/v2/x/node/services"
	plan "github.com/qubetics/qubetics-blockchain/v2/x/plan/services"
	provider "github.com/qubetics/qubetics-blockchain/v2/x/provider/services"
	session "github.com/qubetics/qubetics-blockchain/v2/x/session/services"
	subscription "github.com/qubetics/qubetics-blockchain/v2/x/subscription/services"
	"github.com/qubetics/qubetics-blockchain/v2/x/vpn/keeper"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	deposit.RegisterServices(configurator, k.Deposit)
	lease.RegisterServices(configurator, k.Lease)
	node.RegisterServices(configurator, k.Node)
	plan.RegisterServices(configurator, k.Plan)
	provider.RegisterServices(configurator, k.Provider)
	session.RegisterServices(configurator, k.Session)
	subscription.RegisterServices(configurator, k.Subscription)
}
