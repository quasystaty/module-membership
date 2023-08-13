package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noria-net/module-membership/x/membership/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Member(goCtx context.Context, req *types.QueryMemberRequest) (*types.QueryMemberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Target member must have a valid address
	accAddress, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	// Get the member's account
	memberAccount, found := k.GetMemberAccount(ctx, accAddress)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrUnknownAddress, "member not found")
	}

	// Get the member's nickname
	nickname := k.GetMemberNickname(ctx, accAddress)

	// Return memberAccount inside the response
	return &types.QueryMemberResponse{
		Member: &types.Member{
			BaseAccount: memberAccount.BaseAccount,
			Status:      memberAccount.Status,
			Nickname:    nickname,
		},
	}, nil
}
