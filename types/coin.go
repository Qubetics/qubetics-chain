package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoQubetics defines the default coin denomination used in Qubetics in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Qubetics.
	AttoQubetics string = "tics"

	// BaseDenomUnit defines the base denomination unit for Qubetics.
	// 1 qubetics = 1x10^{BaseDenomUnit} tics
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewQubeticsCoin is a utility function that returns an "tics" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewQubeticsCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoQubetics, amount)
}

// NewqubeticsdecCoin is a utility function that returns an "tics" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewqubeticsdecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoQubetics, amount)
}

// NewQubeticsCoinInt64 is a utility function that returns an "tics" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewQubeticsCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoQubetics, amount)
}
