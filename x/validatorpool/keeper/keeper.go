package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           storetypes.StoreKey
		memKey             storetypes.StoreKey
		paramstore         paramtypes.Subspace
		authority          string
		bankKeeper         types.BankKeeper
		accountKeeper      types.AccountKeeper
		stakingKeeper      types.StakingKeeper
		mintKeeper         types.MintKeeper
		distributionKeeper types.DistributionKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authority string,
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	stakingKeeper types.StakingKeeper,
	mintKeeper types.MintKeeper,
	distributionKeeper types.DistributionKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		paramstore:         ps,
		authority:          authority,
		bankKeeper:         bankKeeper,
		accountKeeper:      accountKeeper,
		stakingKeeper:      stakingKeeper,
		mintKeeper:         mintKeeper,
		distributionKeeper: distributionKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
