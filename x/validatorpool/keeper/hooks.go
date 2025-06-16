package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Hooks implements the staking hooks interface
type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

// Create new validator pool hooks
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// AfterValidatorCreated implements the staking hooks interface
func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {
	return nil
}

// BeforeValidatorModified implements the staking hooks interface
func (h Hooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) error {
	return nil
}

// AfterValidatorRemoved implements the staking hooks interface
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	return nil
}

// AfterValidatorBonded implements the staking hooks interface
func (h Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	h.k.AllocateTokens(ctx)
	return nil
}

// AfterValidatorBeginUnbonding implements the staking hooks interface
func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	return nil
}

// BeforeDelegationCreated implements the staking hooks interface
func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

// BeforeDelegationSharesModified implements the staking hooks interface
func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

// BeforeDelegationRemoved implements the staking hooks interface
func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

// AfterDelegationModified implements the staking hooks interface
func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

// BeforeValidatorSlashed implements the staking hooks interface
func (h Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) error {
	return nil
}

// AfterUnbondingInitiated implements the staking hooks interface
func (h Hooks) AfterUnbondingInitiated(ctx sdk.Context, id uint64) error {
	return nil
}
