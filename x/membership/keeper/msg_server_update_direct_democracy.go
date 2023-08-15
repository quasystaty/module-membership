package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/module-membership/x/membership/types"
)

func (k msgServer) UpdateDirectDemocracy(goCtx context.Context, msg *types.MsgUpdateDirectDemocracy) (*types.MsgUpdateDirectDemocracyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateDirectDemocracyResponse{}, nil
}
