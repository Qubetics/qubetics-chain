package keeper

import (
	"fmt"
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qubetics/qubetics-blockchain/v2/utils"
	"github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
)

// WhitelistedChains contains the list of chain IDs that are allowed to register ERC20 precompiles.
var WhitelistedChains = []string{
	"trusted-chain-1", // Example trusted chain ID
	"trusted-chain-2", // Another trusted chain ID
	"qubetics_9087-1",
}

// RegisterERC20Extension creates and adds an ERC20 precompile interface for an IBC Coin.
//
// It derives the ERC-20 address from the token denomination and registers the
// EVM extension as an active dynamic precompile.
//
// CONTRACT: This must ONLY be called if there is no existing token pair for the given denom.
func (k Keeper) RegisterERC20Extension(ctx sdk.Context, denom string) (*types.TokenPair, error) {
	// Check if the source chain is whitelisted
	sourceChainID := ctx.ChainID() // Get the chain ID of the current chain
	if !isChainWhitelisted(sourceChainID) {
		return nil, fmt.Errorf("chain %s is not whitelisted for ERC20 precompile creation", sourceChainID)
	}

	pair, err := k.CreateNewTokenPair(ctx, denom)
	if err != nil {
		return nil, err
	}
	// Add to existing EVM extensions
	err = k.EnableDynamicPrecompiles(ctx, pair.GetERC20Contract())
	if err != nil {
		return nil, err
	}
	return &pair, err
}

// isChainWhitelisted checks if the chain is in the list of whitelisted chains.
func isChainWhitelisted(chainID string) bool {
	for _, whitelistedChain := range WhitelistedChains {
		if chainID == whitelistedChain {
			return true
		}
	}
	return false
}

// EnableDynamicPrecompiles appends the addresses of the given Precompiles to the list
// of active dynamic precompiles.
func (k Keeper) EnableDynamicPrecompiles(ctx sdk.Context, addresses ...common.Address) error {
	// Get the current params and append the new precompiles
	params := k.GetParams(ctx)
	activePrecompiles := params.DynamicPrecompiles

	// Append and sort the new precompiles
	updatedPrecompiles, err := appendPrecompiles(activePrecompiles, addresses...)
	if err != nil {
		return err
	}

	// Update params
	params.DynamicPrecompiles = updatedPrecompiles
	k.Logger(ctx).Info("Added new precompiles", "addresses", addresses)
	return k.SetParams(ctx, params)
}

// appendPrecompiles append addresses to the existingPrecompiles and sort the resulting slice.
// The function returns an error if the two sets are overlapping.
func appendPrecompiles(existingPrecompiles []string, addresses ...common.Address) ([]string, error) {
	// check for duplicates
	hexAddresses := make([]string, len(addresses))
	for i := range addresses {
		addrHex := addresses[i].Hex()
		if slices.Contains(existingPrecompiles, addrHex) {
			return nil, fmt.Errorf("attempted to register a duplicate precompile address: %s", addrHex)
		}
		hexAddresses[i] = addrHex
	}

	existingLength := len(existingPrecompiles)
	updatedPrecompiles := make([]string, existingLength+len(hexAddresses))
	copy(updatedPrecompiles, existingPrecompiles)
	copy(updatedPrecompiles[existingLength:], hexAddresses)

	utils.SortSlice(updatedPrecompiles)
	return updatedPrecompiles, nil
}
