// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package keeper

import (
	"context"
	"strings"

	"github.com/armon/go-metrics"
	"github.com/ethereum/go-ethereum/common"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
)

var _ types.MsgServer = Keeper{}

// Transfer defines a gRPC msg server method for the MsgTransfer message.
// This implementation overrides the default ICS20 transfer by converting
// the ERC20 tokens to their Cosmos representation if the token pair has been
// registered through governance.
func (k Keeper) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Determine if gas should be disabled — ONLY for precompile invocations.
	if isEVMPrecompileCall(ctx) {
		kvGasCfg := ctx.KVGasConfig()
		transientKVGasCfg := ctx.TransientKVGasConfig()
		ctx = ctx.
			WithKVGasConfig(storetypes.GasConfig{}).
			WithTransientKVGasConfig(storetypes.GasConfig{})

		defer func() {
			// Restore original gas config
			ctx = ctx.
				WithKVGasConfig(kvGasCfg).
				WithTransientKVGasConfig(transientKVGasCfg)
		}()
	}

	// use native denom or contract address
	denom := strings.TrimPrefix(msg.Token.Denom, erc20types.ModuleName+"/")

	pairID := k.erc20Keeper.GetTokenPairID(ctx, denom)
	if len(pairID) == 0 {
		// no-op: token is not registered so we can proceed with regular transfer
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	pair, _ := k.erc20Keeper.GetTokenPair(ctx, pairID)
	if !pair.Enabled {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	sender := sdk.MustAccAddressFromBech32(msg.Sender)

	if !k.erc20Keeper.IsERC20Enabled(ctx) {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	msg.Token.Denom = pair.Denom

	if !pair.IsNativeERC20() {
		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	balance := k.bankKeeper.GetBalance(ctx, sender, pair.Denom)
	if balance.Amount.GTE(msg.Token.Amount) {

		defer func() {
			telemetry.IncrCounterWithLabels(
				[]string{"erc20", "ibc", "transfer", "total"},
				1,
				[]metrics.Label{
					telemetry.NewLabel("denom", pair.Denom),
				},
			)
		}()

		return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
	}

	difference := msg.Token.Amount.Sub(balance.Amount)

	msgConvertERC20 := erc20types.NewMsgConvertERC20(
		difference,
		sender,
		pair.GetERC20Contract(),
		common.BytesToAddress(sender.Bytes()),
	)

	if _, err := k.erc20Keeper.ConvertERC20(sdk.WrapSDKContext(ctx), msgConvertERC20); err != nil {
		return nil, err
	}

	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"erc20", "ibc", "transfer", "total"},
			1,
			[]metrics.Label{
				telemetry.NewLabel("denom", pair.Denom),
			},
		)
	}()

	return k.Keeper.Transfer(sdk.WrapSDKContext(ctx), msg)
}

// isEVMPrecompileCall returns true if the current context is from a precompiled contract.
// You may replace or extend this logic based on how your EVM context flags are structured.
func isEVMPrecompileCall(ctx sdk.Context) bool {
	// Placeholder logic: customize to suit your precompile call detection
	// For example, check context values, Chain ID, or a flag set by the EVM module
	evmCaller := ctx.Value("evm_caller")
	return evmCaller != nil
}
