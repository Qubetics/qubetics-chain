package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/qubetics/qubetics-blockchain/v2/types"
	v1base "github.com/qubetics/qubetics-blockchain/v2/types/v1"
	v3 "github.com/qubetics/qubetics-blockchain/v2/x/session/types/v3"
)

// Default per-category caps to bound work per block.
const (
	DefaultMaxInactivePendingSessionsPerBlock = 1000
	DefaultMaxInactiveSessionsPerBlock        = 1000
)

// MaxInactivePendingSessionsPerBlock returns the maximum inactive pending sessions to process per block.
func (k *Keeper) MaxInactivePendingSessionsPerBlock(ctx sdk.Context) int {
	return DefaultMaxInactivePendingSessionsPerBlock
}

// MaxInactiveSessionsPerBlock returns the maximum inactive sessions to process per block.
func (k *Keeper) MaxInactiveSessionsPerBlock(ctx sdk.Context) int {
	return DefaultMaxInactiveSessionsPerBlock
}

// shouldHaltForGas returns true if the block gas meter indicates we've exceeded the limit.
func (k *Keeper) shouldHaltForGas(ctx sdk.Context) bool {
	if ctx.BlockGasMeter() == nil {
		return false
	}
	return ctx.BlockGasMeter().IsPastLimit()
}

// handleInactivePendingSessions processes pending sessions that have become inactive at the current block time.
func (k *Keeper) handleInactivePendingSessions(ctx sdk.Context) {
	limit := k.MaxInactivePendingSessionsPerBlock(ctx)
	processed := 0

	// Iterate through sessions that have become inactive at the current block time
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		// Skip the session if its status is not active
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive pending session", "id", item.GetID())

		// Create a message to cancel the inactive pending session
		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		// Get the appropriate handler for processing the cancel session message
		handler := k.router.Handler(msg)
		if handler == nil {
			// Log error instead of panicking to prevent chain halt
			k.Logger(ctx).Error("Handler not found for cancel session message", "id", item.GetID(), "type", sdk.MsgTypeURL(msg))
			return false
		}

		// Execute the handler to process the session cancel request
		resp, err := handler(ctx, msg)
		if err != nil {
			// Log error instead of panicking. The session remains in InactivePending state
			// and will be retried in the next block
			k.Logger(ctx).Error("Failed to cancel inactive pending session", "id", item.GetID(), "error", err.Error())
			return false
		}

		// Emit any events generated during the cancel process
		ctx.EventManager().EmitEvents(resp.GetEvents())

		processed++
		if processed >= limit {
			k.Logger(ctx).Info("Inactive pending session processing limit reached", "processed", processed)
			return true
		}
		if k.shouldHaltForGas(ctx) {
			k.Logger(ctx).Info("Inactive pending session processing halted due to block gas limit")
			return true
		}
		return false
	})
}

// handleInactiveSessions processes sessions that are in the inactive pending state.
func (k *Keeper) handleInactiveSessions(ctx sdk.Context) {
	limit := k.MaxInactiveSessionsPerBlock(ctx)
	processed := 0

	// Iterate through sessions that are inactive pending at the current block time
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		// Skip the session if its status is not inactive pending
		if !item.GetStatus().Equal(v1base.StatusInactivePending) {
			return false
		}

		k.Logger(ctx).Info("Handling inactive session", "id", item.GetID())

		// Perform pre-hook processing for sessions that are transitioning to inactive
		if err := k.SessionInactivePreHook(ctx, item.GetID()); err != nil {
			// Log error instead of panicking to prevent chain halt
			k.Logger(ctx).Error("Failed to process session inactive pre-hook", "id", item.GetID(), "error", err.Error())
			return false
		}

		// Convert account address and node address from Bech32 format.
		accAddr, err := sdk.AccAddressFromBech32(item.GetAccAddress())
		if err != nil {
			// Log error instead of panicking to prevent chain halt
			k.Logger(ctx).Error("Failed to parse account address for session", "id", item.GetID(), "error", err.Error())
			return false
		}

		nodeAddr, err := base.NodeAddressFromBech32(item.GetNodeAddress())
		if err != nil {
			// Log error instead of panicking to prevent chain halt
			k.Logger(ctx).Error("Failed to parse node address for session", "id", item.GetID(), "error", err.Error())
			return false
		}

		// Delete the session from the inactive queue and associated records.
		k.DeleteSession(ctx, item.GetID())
		k.DeleteSessionForAccount(ctx, accAddr, item.GetID())
		k.DeleteSessionForNode(ctx, nodeAddr, item.GetID())
		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		// Emit an event indicating the update of the session status to inactive
		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdateStatus{
				ID:          item.GetID(),
				AccAddress:  item.GetAccAddress(),
				NodeAddress: item.GetNodeAddress(),
				Status:      v1base.StatusInactive,
				StatusAt:    ctx.BlockTime().String(),
			},
		)

		processed++
		if processed >= limit {
			k.Logger(ctx).Info("Inactive session processing limit reached", "processed", processed)
			return true
		}
		if k.shouldHaltForGas(ctx) {
			k.Logger(ctx).Info("Inactive session processing halted due to block gas limit")
			return true
		}
		return false
	})
}

// BeginBlock is called at the beginning of each block to handle session-related operations.
func (k *Keeper) BeginBlock(ctx sdk.Context) {
	// Handle inactive pending sessions at the beginning of each block
	k.handleInactivePendingSessions(ctx)

	// Handle sessions that have become inactive at the beginning of each block
	k.handleInactiveSessions(ctx)
}
