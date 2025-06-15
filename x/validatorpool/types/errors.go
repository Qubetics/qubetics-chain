package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/validatorpool module sentinel errors
var (
	ErrSample = errorsmod.Register(ModuleName, 1100, "sample error")
)

// Validator pool module errors
var (
	ErrInvalidTotalSupply      = errorsmod.Register(ModuleName, 1, "invalid total supply")
	ErrInvalidLockedSupply     = errorsmod.Register(ModuleName, 2, "invalid locked supply")
	ErrInvalidDailyFees        = errorsmod.Register(ModuleName, 3, "invalid daily fees")
	ErrInvalidAnnualRoi        = errorsmod.Register(ModuleName, 4, "invalid annual ROI")
	ErrInvalidDeflationRate    = errorsmod.Register(ModuleName, 5, "invalid deflation rate")
	ErrPoolNotFound            = errorsmod.Register(ModuleName, 6, "validator pool not found")
	ErrNoValidators            = errorsmod.Register(ModuleName, 7, "no active validators found")
	ErrInsufficientPoolBalance = errorsmod.Register(ModuleName, 8, "insufficient pool balance")
	ErrInvalidRewardAmount     = errorsmod.Register(ModuleName, 9, "invalid reward amount")
)
