package keeper

import (
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

var _ types.QueryServer = Keeper{}
