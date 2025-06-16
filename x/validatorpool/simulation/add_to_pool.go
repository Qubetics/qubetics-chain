package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/keeper"
	"github.com/qubetics/qubetics-blockchain/v2/x/validatorpool/types"
)

func SimulateMsgAddToPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddToPool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddToPool simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddToPool simulation not implemented"), nil, nil
	}
}
