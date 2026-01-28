package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v3 "github.com/qubetics/qubetics-blockchain/v2/x/node/types/v3"
	sessiontypes "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v3"
)

// UpdateSessionMaxValues checks the session type and updates its values if needed.
//
// NOTE: This is a stub implementation that currently performs no updates. The subscription keeper
// has the actual implementation (x/subscription/keeper/session.go:UpdateSessionMaxValues) which
// calculates and updates session max bytes and duration based on allocation details.
//
// The node keeper's version is required to satisfy the SessionKeeper interface defined in
// x/session/keeper/expected.go. Once node-specific session logic is needed, this function
// should be implemented here. For now, it delegates to subscription keeper's implementation
// via the session module orchestration.
//
// TODO: implement node-specific session max value updates
// if the subscription keeper's implementation is sufficient for all use cases.
func (k *Keeper) UpdateSessionMaxValues(_ sdk.Context, session sessiontypes.Session) error {
	// Check if the session is of type v3.Session.
	_, ok := session.(*v3.Session)
	if !ok {
		return nil
	}
	// Return nil as no update needed.
	return nil
}
