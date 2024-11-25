package app

import (
	"cosmossdk.io/math"
)

var (
	// MainnetMinGasPrices defines 20B tics (or atqubetics) as the minimum gas price value on the fee market module.
	MainnetMinGasPrices = math.LegacyNewDec(20_000_000_000)
	// MainnetMinGasMultiplier defines the min gas multiplier value on the fee market module.
	// 50% of the leftover gas will be refunded
	MainnetMinGasMultiplier = math.LegacyNewDecWithPrec(5, 1)
)
