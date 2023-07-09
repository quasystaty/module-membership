package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/module-membership/x/membership/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GetGuardianWhitelist(ctx),
		k.GetTotalVotingWeight(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	// validate the list of guardians
	k.paramstore.SetParamSet(ctx, &params)
}

// Returns the guardian addresses defined in param
// NOTE: these addresses must still be validated against the electorate
func (k Keeper) GetGuardianWhitelist(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyGuardians, &res)
	return res
}

func (k Keeper) GetTotalVotingWeight(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyTotalVotingWeight, &res)
	return res
}
