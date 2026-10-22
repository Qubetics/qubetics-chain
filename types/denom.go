package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
)

const (
	// AttoQubetics is the base denomination of the chain, used for staking,
	// governance deposits, fees and EVM gas (1 qubetics = 10^18 tics).
	AttoQubetics = "tics"

	// DisplayDenom is the denomination shown to users.
	DisplayDenom = "qubetics"

	// BaseDenomUnit is the number of decimals between DisplayDenom and AttoQubetics.
	BaseDenomUnit = 18
)

// PowerReduction converts staked tics into consensus power (1 qubetics = 1 power).
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))
