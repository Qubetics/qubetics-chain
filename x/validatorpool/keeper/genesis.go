package keeper

import (
	"fmt"

	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	fmt.Println("Starting InitGenesis with state:", genState)

	// Ensure module account exists
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("module account %s does not exist", types.ModuleName))
	}
	fmt.Println("Module account address:", moduleAcc.GetAddress())

	// Mint initial locked supply to module account
	initialLockedSupply, ok := sdk.NewIntFromString("298529031000000000000000000") // 298,529,031 with 18 decimals
	if !ok {
		panic("invalid initial locked supply")
	}
	fmt.Println("Initial locked supply:", initialLockedSupply)

	// Check if we need to mint coins
	moduleBalance := k.bankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), "tics")
	fmt.Println("Initial module balance:", moduleBalance)

	// Always mint the initial locked supply
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("tics", initialLockedSupply)))
	if err != nil {
		panic(fmt.Sprintf("failed to mint initial locked supply: %v", err))
	}

	// Verify the minting
	moduleBalance = k.bankKeeper.GetBalance(ctx, moduleAcc.GetAddress(), "tics")
	fmt.Println("Module account balance after minting:", moduleBalance)
	if !moduleBalance.Amount.Equal(initialLockedSupply) {
		panic(fmt.Sprintf("module account balance mismatch: expected %s, got %s", initialLockedSupply, moduleBalance.Amount))
	}

	// Double check the balance after a short delay
	ctx.Logger().Info("Verifying module account balance after initialization",
		"module_address", moduleAcc.GetAddress().String(),
		"expected_balance", initialLockedSupply.String(),
		"actual_balance", moduleBalance.Amount.String(),
	)
}
