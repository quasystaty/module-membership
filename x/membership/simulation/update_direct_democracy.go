package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/noria-net/module-membership/x/membership/keeper"
	"github.com/noria-net/module-membership/x/membership/types"
)

func SimulateMsgUpdateDirectDemocracy(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateDirectDemocracy{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateDirectDemocracy simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateDirectDemocracy simulation not implemented"), nil, nil
	}
}
