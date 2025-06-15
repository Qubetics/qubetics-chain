package keeper

import (
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	// return types.DefaultParams()
	return types.Params{
		AnnualRoi: k.GetAnnualRoi(ctx),
	}
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// GetAnnualRoi returns the annual ROI parameter
func (k Keeper) GetAnnualRoi(ctx sdk.Context) uint32 {
	var annualRoi uint32
	if k.paramstore.Has(ctx, types.KeyAnnualRoi) {
		k.paramstore.Get(ctx, types.KeyAnnualRoi, &annualRoi)
	} else {
		// Return default if not set
		annualRoi = types.DefaultParams().AnnualRoi
	}
	return annualRoi
}

// SetAnnualRoi sets the annual ROI parameter
func (k Keeper) SetAnnualRoi(ctx sdk.Context, annualRoi uint32) {
	k.paramstore.Set(ctx, types.KeyAnnualRoi, annualRoi)
}
