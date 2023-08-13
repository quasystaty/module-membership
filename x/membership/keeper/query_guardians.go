package keeper

import (
	"context"

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
	totalVotingWeight := k.GetDirectDemocracySettings(ctx).TotalVotingWeight

	return &types.QueryGuardiansResponse{
		Members:           guardians,
		TotalVotingWeight: &totalVotingWeight,
	}, nil
}
