package v2

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"qubetics.subscription.v2.Subscription",
		(*Subscription)(nil),
		&NodeSubscription{},
		&PlanSubscription{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCancelRequest{},
		&MsgAllocateRequest{},
	)
}
