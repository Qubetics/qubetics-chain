package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"

	tmdb "github.com/cometbft/cometbft-db"
	cmttypeslog "github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

// MockAccountKeeper is a mock of AccountKeeper interface
type MockAccountKeeper struct {
	accounts map[string]authtypes.AccountI
}

func NewMockAccountKeeper() *MockAccountKeeper {
	return &MockAccountKeeper{
		accounts: make(map[string]authtypes.AccountI),
	}
}

func (m *MockAccountKeeper) GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI {
	if acc, ok := m.accounts[moduleName]; ok {
		return acc.(authtypes.ModuleAccountI)
	}
	return nil
}

func (m *MockAccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI {
	if acc, ok := m.accounts[addr.String()]; ok {
		return acc
	}
	return nil
}

// MockBankKeeper is a mock of BankKeeper interface
type MockBankKeeper struct {
	balances          map[string]sdk.Coins
	ModuleAccountAddr string
}

func NewMockBankKeeper() *MockBankKeeper {
	return &MockBankKeeper{
		balances: make(map[string]sdk.Coins),
	}
}

func (m *MockBankKeeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	if coins, ok := m.balances[addr.String()]; ok {
		amt := coins.AmountOf(denom)
		return sdk.NewCoin(denom, amt)
	}
	return sdk.NewCoin(denom, math.ZeroInt())
}

func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	// Update sender balance
	if coins, ok := m.balances[senderAddr.String()]; ok {
		m.balances[senderAddr.String()] = coins.Sub(amt...)
	}
	// Update module balance using ModuleAccountAddr
	if coins, ok := m.balances[m.ModuleAccountAddr]; ok {
		m.balances[m.ModuleAccountAddr] = coins.Add(amt...)
	} else {
		m.balances[m.ModuleAccountAddr] = amt
	}
	return nil
}

func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	// Update sender module balance
	if coins, ok := m.balances[senderModule]; ok {
		m.balances[senderModule] = coins.Sub(amt...)
	}

	// Update recipient module balance
	if coins, ok := m.balances[recipientModule]; ok {
		m.balances[recipientModule] = coins.Add(amt...)
	} else {
		m.balances[recipientModule] = amt
	}
	return nil
}

func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	if coins, ok := m.balances[senderModule]; ok {
		m.balances[senderModule] = coins.Sub(amt...)
	}
	if coins, ok := m.balances[recipientAddr.String()]; ok {
		m.balances[recipientAddr.String()] = coins.Add(amt...)
	} else {
		m.balances[recipientAddr.String()] = amt
	}
	return nil
}

func (m *MockBankKeeper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	if coins, ok := m.balances[moduleName]; ok {
		m.balances[moduleName] = coins.Add(amt...)
	} else {
		m.balances[moduleName] = amt
	}
	return nil
}

// MockStakingKeeper is a mock of StakingKeeper interface
type MockStakingKeeper struct {
	validators []stakingtypes.Validator
	bondDenom  string
}

func NewMockStakingKeeper() *MockStakingKeeper {
	return &MockStakingKeeper{
		bondDenom: "stake",
	}
}

func (m *MockStakingKeeper) BondDenom(ctx sdk.Context) string {
	return m.bondDenom
}

func (m *MockStakingKeeper) GetBondedValidatorsByPower(ctx sdk.Context) []stakingtypes.Validator {
	return m.validators
}

func (m *MockStakingKeeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (stakingtypes.Validator, bool) {
	for _, val := range m.validators {
		if val.GetOperator().Equals(addr) {
			return val, true
		}
	}
	return stakingtypes.Validator{}, false
}

func (m *MockStakingKeeper) TotalBondedTokens(ctx sdk.Context) math.Int {
	total := math.ZeroInt()
	for _, val := range m.validators {
		total = total.Add(val.GetBondedTokens())
	}
	return total
}

// MockMintKeeper is a mock of MintKeeper interface
type MockMintKeeper struct {
	blocksPerYear uint64
}

func NewMockMintKeeper() *MockMintKeeper {
	return &MockMintKeeper{
		blocksPerYear: 6311520, // Default blocks per year
	}
}

func (m *MockMintKeeper) GetParams(ctx sdk.Context) minttypes.Params {
	return minttypes.Params{
		BlocksPerYear: m.blocksPerYear,
	}
}

func setupTest(t *testing.T) (sdk.Context, *keeper.Keeper, *MockAccountKeeper, *MockBankKeeper, *MockStakingKeeper, *MockMintKeeper) {
	// Create in-memory database
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	// Create store keys
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	// Mount stores
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	// Create context
	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, cmttypeslog.NewNopLogger())

	// Create codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Create params subspace
	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"ValidatorpoolParams",
	)

	// Create mock keepers
	mockAccountKeeper := NewMockAccountKeeper()
	mockBankKeeper := NewMockBankKeeper()
	mockStakingKeeper := NewMockStakingKeeper()
	mockMintKeeper := NewMockMintKeeper()

	// Create validatorpool keeper
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		"",
		mockBankKeeper,
		mockAccountKeeper,
		mockStakingKeeper,
		mockMintKeeper,
		nil, // distribution keeper not needed for these tests
	)

	// Set up module accounts
	validatorPoolAcc := authtypes.NewEmptyModuleAccount(types.ModuleName)
	feeCollectorAcc := authtypes.NewEmptyModuleAccount("fee_collector")
	mockAccountKeeper.accounts[types.ModuleName] = validatorPoolAcc
	mockAccountKeeper.accounts["fee_collector"] = feeCollectorAcc
	// Set up balances using address string as key
	mockBankKeeper.ModuleAccountAddr = validatorPoolAcc.GetAddress().String()
	mockBankKeeper.balances[validatorPoolAcc.GetAddress().String()] = sdk.NewCoins(sdk.NewCoin("stake", math.ZeroInt()))
	mockBankKeeper.balances[feeCollectorAcc.GetAddress().String()] = sdk.NewCoins(sdk.NewCoin("stake", math.ZeroInt()))

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return ctx, k, mockAccountKeeper, mockBankKeeper, mockStakingKeeper, mockMintKeeper
}

func TestAddToPool(t *testing.T) {
	ctx, k, mockAccountKeeper, mockBankKeeper, _, _ := setupTest(t)

	// Create test accounts
	creator := sdk.AccAddress([]byte("creator"))
	amount := math.NewInt(1000)

	// Set up initial balances
	mockBankKeeper.balances[creator.String()] = sdk.NewCoins(sdk.NewCoin("stake", math.NewInt(2000)))

	tests := []struct {
		name          string
		creator       sdk.AccAddress
		amount        math.Int
		expectedError bool
	}{
		{
			name:          "successful add to pool",
			creator:       creator,
			amount:        amount,
			expectedError: false,
		},
		{
			name:          "empty creator address",
			creator:       sdk.AccAddress{},
			amount:        amount,
			expectedError: true,
		},
		{
			name:          "zero amount",
			creator:       creator,
			amount:        math.ZeroInt(),
			expectedError: true,
		},
		{
			name:          "negative amount",
			creator:       creator,
			amount:        math.NewInt(-1000),
			expectedError: true,
		},
		{
			name:          "insufficient balance",
			creator:       creator,
			amount:        math.NewInt(3000),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := k.AddToPool(ctx, tt.creator, tt.amount)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// Verify balances after transfer
				creatorBalance := mockBankKeeper.GetBalance(ctx, tt.creator, "stake")
				poolBalance := mockBankKeeper.GetBalance(ctx, mockAccountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress(), "stake")
				require.Equal(t, math.NewInt(2000).Sub(tt.amount), creatorBalance.Amount)
				require.Equal(t, tt.amount, poolBalance.Amount)
			}
		})
	}
}

func TestAllocateTokens(t *testing.T) {
	ctx, k, _, mockBankKeeper, mockStakingKeeper, _ := setupTest(t)

	// Set up initial pool balance
	validatorPoolAcc := authtypes.NewEmptyModuleAccount(types.ModuleName)
	feeCollectorAcc := authtypes.NewEmptyModuleAccount("fee_collector")
	initialPoolBalance := sdk.NewCoin("stake", math.NewInt(1000000))
	mockBankKeeper.balances[validatorPoolAcc.GetAddress().String()] = sdk.NewCoins(initialPoolBalance)
	mockBankKeeper.balances[feeCollectorAcc.GetAddress().String()] = sdk.NewCoins(sdk.NewCoin("stake", math.ZeroInt()))

	// Set up validators
	val1 := stakingtypes.Validator{
		OperatorAddress: "val1",
		Tokens:          math.NewInt(1000),
		Status:          stakingtypes.Bonded,
	}
	val2 := stakingtypes.Validator{
		OperatorAddress: "val2",
		Tokens:          math.NewInt(2000),
		Status:          stakingtypes.Bonded,
	}
	mockStakingKeeper.validators = []stakingtypes.Validator{val1, val2}

	tests := []struct {
		name          string
		setup         func()
		expectedError bool
	}{
		{
			name: "successful token allocation",
			setup: func() {
				// Set up initial balances
				mockBankKeeper.balances[types.ModuleName] = sdk.NewCoins(initialPoolBalance)
				mockBankKeeper.balances["fee_collector"] = sdk.NewCoins(sdk.NewCoin("stake", math.ZeroInt()))
			},
			expectedError: false,
		},
		{
			name: "empty pool balance",
			setup: func() {
				mockBankKeeper.balances[types.ModuleName] = sdk.NewCoins(sdk.NewCoin("stake", math.ZeroInt()))
			},
			expectedError: false,
		},
		{
			name: "zero total bonded tokens",
			setup: func() {
				mockBankKeeper.balances[types.ModuleName] = sdk.NewCoins(initialPoolBalance)
				mockStakingKeeper.validators = []stakingtypes.Validator{}
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			err := k.AllocateTokens(ctx)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
