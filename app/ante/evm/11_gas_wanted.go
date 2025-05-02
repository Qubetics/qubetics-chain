package evm

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/qubetics/qubetics-blockchain/v2/types"
)

// GasWantedDecorator keeps track of the gasWanted amount on the current block in transient store
// for BaseFee calculation.
// NOTE: This decorator does not perform any validation
type GasWantedDecorator struct {
	evmKeeper       EVMKeeper
	feeMarketKeeper FeeMarketKeeper
}

// NewGasWantedDecorator creates a new NewGasWantedDecorator
func NewGasWantedDecorator(
	evmKeeper EVMKeeper,
	feeMarketKeeper FeeMarketKeeper,
) GasWantedDecorator {
	return GasWantedDecorator{
		evmKeeper,
		feeMarketKeeper,
	}
}

func (gwd GasWantedDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	evmParams := gwd.evmKeeper.GetParams(ctx)
	chainCfg := evmParams.GetChainConfig()
	ethCfg := chainCfg.EthereumConfig(gwd.evmKeeper.ChainID())

	blockHeight := big.NewInt(ctx.BlockHeight())
	isLondon := ethCfg.IsLondon(blockHeight)

	if err := CheckGasWanted(ctx, gwd.feeMarketKeeper, tx, isLondon); err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate)
}

// GetTxGasLimit is a hypothetical function that retrieves the transaction-specific gas limit.
// You will need to implement this based on your application's logic.
func GetTxGasLimit(tx sdk.Tx) uint64 {
	// Example: Assume we are retrieving the gas limit from the transaction's properties or some config
	// This could depend on the transaction type, smart contract, etc.
	// For now, we're just returning a fixed gas limit for demonstration purposes.
	return 1000000 // Example gas limit, this should be replaced with actual logic
}

func CheckGasWanted(ctx sdk.Context, feeMarketKeeper FeeMarketKeeper, tx sdk.Tx, isLondon bool) error {
	if !isLondon {
		return nil
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return nil
	}

	gasWanted := feeTx.GetGas()

	// Retrieve the transaction-specific gas limit (could be from a configuration, contract, etc.)
	txGasLimit := GetTxGasLimit(tx)  // Hypothetical function to get the tx-specific gas limit

	// Return error if the tx gas exceeds the tx gas limit
	if gasWanted > txGasLimit {
		return errorsmod.Wrapf(
			errortypes.ErrOutOfGas,
			"tx gas (%d) exceeds tx-specific gas limit (%d)",
			gasWanted,
			txGasLimit,
		)
	}

	isBaseFeeEnabled := feeMarketKeeper.GetBaseFeeEnabled(ctx)
	if !isBaseFeeEnabled {
		return nil
	}

	// Add total gasWanted to cumulative in block transientStore in FeeMarket module
	if _, err := feeMarketKeeper.AddTransientGasWanted(ctx, gasWanted); err != nil {
		return errorsmod.Wrapf(err, "failed to add gas wanted to transient store")
	}

	return nil
}
