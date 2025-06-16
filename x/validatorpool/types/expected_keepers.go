package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper defines the expected staking keeper interface
type StakingKeeper interface {
	// BondDenom returns the bondable coin denomination
	BondDenom(ctx sdk.Context) (res string)
	// GetBondedValidatorsByPower returns the list of validators sorted by power
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator
	// GetValidator returns the validator for the given address
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (stakingtypes.Validator, bool)
	// TotalBondedTokens returns the total bonded tokens
	TotalBondedTokens(ctx sdk.Context) math.Int
}

// DistributionKeeper defines the expected distribution keeper interface
type DistributionKeeper interface {
	GetValidatorAccumulatedCommission(ctx sdk.Context, val sdk.ValAddress) (commission distrtypes.ValidatorAccumulatedCommission)
	GetValidatorOutstandingRewards(ctx sdk.Context, val sdk.ValAddress) (rewards distrtypes.ValidatorOutstandingRewards)
	// GetDelegatorRewards(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (rewards sdk.DecCoins)
	AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins)
}

type MintKeeper interface {
	// Methods imported from mint should be defined here
	GetParams(ctx sdk.Context) minttypes.Params
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected bank keeper interface
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
