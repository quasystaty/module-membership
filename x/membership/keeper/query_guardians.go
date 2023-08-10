package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/module-membership/x/membership/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Guardians(goCtx context.Context, req *types.QueryGuardiansRequest) (*types.QueryGuardiansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	guardians := k.GetGuardians(ctx)
	tw := math.LegacyZeroDec()

	return &types.QueryGuardiansResponse{
		Members:           guardians,
		TotalVotingWeight: &tw,
	}, nil
}
