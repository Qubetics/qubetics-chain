package validatorpool

import (
	"math/rand"

	"github.com/qubetics/qubetics-blockchain/v2/testutil/sample"
	validatorpoolsimulation "github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/simulation"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = validatorpoolsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgAddToPool = "op_weight_msg_add_to_pool"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddToPool int = 100

	opWeightMsgDistributeRewards = "op_weight_msg_distribute_rewards"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDistributeRewards int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	validatorpoolGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&validatorpoolGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddToPool int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddToPool, &weightMsgAddToPool, nil,
		func(_ *rand.Rand) {
			weightMsgAddToPool = defaultWeightMsgAddToPool
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddToPool,
		validatorpoolsimulation.SimulateMsgAddToPool(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDistributeRewards int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDistributeRewards, &weightMsgDistributeRewards, nil,
		func(_ *rand.Rand) {
			weightMsgDistributeRewards = defaultWeightMsgDistributeRewards
		},
	)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddToPool,
			defaultWeightMsgAddToPool,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				validatorpoolsimulation.SimulateMsgAddToPool(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),

		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
