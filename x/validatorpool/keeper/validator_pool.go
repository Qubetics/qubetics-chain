package keeper

import (
	"fmt"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// AddToPool adds tokens to the validator pool
func (k Keeper) AddToPool(ctx sdk.Context, creator sdk.AccAddress, amount math.Int) error {
	// Validate input parameters
	if creator.Empty() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "creator address cannot be empty")
	}
	if amount.IsNil() || amount.IsZero() || amount.IsNegative() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", amount)
	}

	// Check if creator has enough balance
	balance := k.bankKeeper.GetBalance(ctx, creator, k.stakingKeeper.BondDenom(ctx))
	fmt.Println("Creator balance:", balance)
	if balance.Amount.LT(amount) {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient balance: %s", balance.Amount)
	}

	// Get module account
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "module account %s does not exist", types.ModuleName)
	}

	// Transfer tokens from creator to module account
	coins := sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), amount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coins); err != nil {
		return errorsmod.Wrapf(err, "failed to transfer tokens to module account")
	}

	// Emit event for successful transfer
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddToPool,
			sdk.NewAttribute(sdk.AttributeKeySender, creator.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.String()),
		),
	)

	return nil
}

// AllocateTokens calculates and distributes rewards to validators based on their voting power
func (k Keeper) AllocateTokens(ctx sdk.Context) error {

	fmt.Println("hello world========")
	// Get reward calculation parameters from params
	params := k.GetParams(ctx)
	annualRoi := sdk.NewDecFromInt(math.NewInt(int64(params.AnnualRoi))).Quo(sdk.NewDec(100)) // Convert basis points to decimal

	fmt.Println("annualROI=======", annualRoi)

	// Get current locked supply from module account
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	fmt.Println("moduleacccc=======", moduleAcc.GetName())
	if moduleAcc == nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "module account %s does not exist", types.ModuleName)
	}
	fmt.Println("mole acc addr=======",moduleAcc.GetAddress() )

	fmt.Println("bond denom =======", k.stakingKeeper.BondDenom(ctx))
	lockedSupply := k.bankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), k.stakingKeeper.BondDenom(ctx)).Amount
	fmt.Println("lockedsupply=======", lockedSupply)
	// Check if locked supply is zero, return early if true
	if lockedSupply.IsZero() {
		return nil
	}

	fmt.Println("annualRoi====", annualRoi)
	fmt.Println("lockedSupply====", lockedSupply)

	// Get blocks per year from mint module params
	mintParams := k.mintKeeper.GetParams(ctx)
	blocksPerYear := math.NewInt(int64(mintParams.BlocksPerYear))

	fmt.Println("blockperYear=======", blocksPerYear)

	// Get total bonded tokens
	totalBonded := k.stakingKeeper.TotalBondedTokens(ctx)
	if totalBonded.IsZero() {
		return nil
	}

	fmt.Println("totalBonded====", totalBonded)

	// Calculate base reward per block based on total bonded stake and ROI
	// Formula: (totalBonded * annualRoi) / (100 * blocksPerYear)
	baseReward := sdk.NewDecFromInt(totalBonded).
		Mul(annualRoi).
		Quo(sdk.NewDecFromInt(math.NewInt(100))). // Divide by 100 to get percentage
		Quo(sdk.NewDecFromInt(blocksPerYear))

	fmt.Println("Reward calculation details:")
	fmt.Println("totalBonded====", totalBonded)
	fmt.Println("annualRoi====", annualRoi)
	fmt.Println("blocksPerYear====", blocksPerYear)
	fmt.Println("baseReward calculation steps:")
	fmt.Println("1. totalBonded * annualRoi =", sdk.NewDecFromInt(totalBonded).Mul(annualRoi))
	fmt.Println("2. divided by 100 =", sdk.NewDecFromInt(totalBonded).Mul(annualRoi).Quo(sdk.NewDecFromInt(math.NewInt(100))))
	fmt.Println("3. final baseReward =", baseReward)

	// Calculate total rewards to be distributed
	totalRewards := baseReward.TruncateInt()
	if totalRewards.IsZero() {
		return nil
	}

	// Create coins for the total rewards
	rewardCoins := sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), totalRewards))

	// Get the fee collector module account
	feeCollector := k.accountKeeper.GetModuleAccount(ctx, "fee_collector")
	if feeCollector == nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "fee collector module account does not exist")
	}

	// Send coins from validatorpool module to fee collector
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, "fee_collector", rewardCoins); err != nil {
		return errorsmod.Wrapf(err, "failed to send coins to fee collector")
	}

	// Emit event for successful reward distribution
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAllocateTokens,
			sdk.NewAttribute(sdk.AttributeKeyAmount, totalRewards.String()),
		),
	)

	return nil
}
