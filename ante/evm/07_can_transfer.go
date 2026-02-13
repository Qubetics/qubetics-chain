package evm

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
	"github.com/qubetics/qubetics-blockchain/v2/x/vm/statedb"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/vm/types"
)

// CanTransfer checks if the sender is allowed to transfer funds according to the EVM block
func CanTransfer(
	ctx sdk.Context,
	evmKeeper EVMKeeper,
	msg core.Message,
	baseFee *big.Int,
	ethCfg *params.ChainConfig,
	params evmtypes.Params,
	isLondon bool,
) error {
	if isLondon && msg.GasFeeCap().Cmp(baseFee) < 0 {
		return errorsmod.Wrapf(
			errortypes.ErrInsufficientFee,
			"max fee per gas less than block base fee (%s < %s)",
			msg.GasFeeCap(), baseFee,
		)
	}

	//  QBA-12: Reject transactions with a negative value to prevent underflows or malformed behavior
	if msg.Value().Sign() < 0 {
		return errorsmod.Wrapf(
			errortypes.ErrInvalidRequest,
			"transaction value must be non-negative: %s",
			msg.Value().String(),
		)
	}

	// Setup EVM configuration
	cfg := &statedb.EVMConfig{
		ChainConfig: ethCfg,
		Params:      params,
		CoinBase:    common.Address{},
		BaseFee:     baseFee,
	}

	// Create a new stateDB instance for the transaction
	stateDB := statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
	evm := evmKeeper.NewEVM(ctx, msg, cfg, evmtypes.NewNoOpTracer(), stateDB)

	// Check that the caller has enough balance to cover asset transfer for the topmost call
	if msg.Value().Sign() > 0 && !evm.Context.CanTransfer(stateDB, msg.From(), msg.Value()) {
		return errorsmod.Wrapf(
			errortypes.ErrInsufficientFunds,
			"failed to transfer %s from address %s using the EVM block context transfer function",
			msg.Value(),
			msg.From(),
		)
	}

	return nil
}
