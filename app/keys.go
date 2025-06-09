package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v7/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	epochstypes "github.com/qubetics/qubetics-blockchain/v2/x/epochs/types"
	erc20types "github.com/qubetics/qubetics-blockchain/v2/x/erc20/types"
	evmtypes "github.com/qubetics/qubetics-blockchain/v2/x/evm/types"
	feemarkettypes "github.com/qubetics/qubetics-blockchain/v2/x/feemarket/types"
	vestingtypes "github.com/qubetics/qubetics-blockchain/v2/x/vesting/types"
	customminttypes "github.com/qubetics/qubetics-blockchain/v2/x/mint/types"
	oracletypes "github.com/qubetics/qubetics-blockchain/v2/x/oracle/types"
	vpntypes "github.com/qubetics/qubetics-blockchain/v2/x/vpn/types"


)

// StoreKeys returns the application store keys,
// the EVM transient store keys and the memory store keys
func StoreKeys() (
	map[string]*storetypes.KVStoreKey,
	map[string]*storetypes.MemoryStoreKey,
	map[string]*storetypes.TransientStoreKey,
) {
	storeKeys := []string{
		// SDK keys
		authtypes.StoreKey, banktypes.StoreKey, minttypes.StoreKey, stakingtypes.StoreKey,
		distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, capabilitytypes.StoreKey, consensusparamtypes.StoreKey,
		feegrant.StoreKey, authzkeeper.StoreKey,
		// ibc keys
		ibcexported.StoreKey, ibctransfertypes.StoreKey,
		// ica keys
		icahosttypes.StoreKey,
		// ibc rate-limit keys
		ratelimittypes.StoreKey,
		// ethermint keys
		evmtypes.StoreKey, feemarkettypes.StoreKey,
		// qubetics keys
		erc20types.StoreKey,
		epochstypes.StoreKey, vestingtypes.StoreKey,

	// Sentinel Hub keys
	customminttypes.StoreKey,
	oracletypes.StoreKey,
	vpntypes.StoreKey,
	
	}

	keys := sdk.NewKVStoreKeys(storeKeys...)

	// Add the EVM transient store key
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	return keys, memKeys, tkeys
}
