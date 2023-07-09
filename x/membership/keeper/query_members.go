package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/noria-net/module-membership/x/membership/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Members(goCtx context.Context, req *types.QueryMembersRequest) (*types.QueryMembersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var members types.Members
	ctx := sdk.UnwrapSDKContext(goCtx)
	guardians := k.GetGuardianAddresses(ctx)
	membersStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.MembersKeyPrefix)

	query.Paginate(membersStore, req.Pagination, func(key []byte, value []byte) error {
		var member types.Member
		if err := k.cdc.Unmarshal(value, &member); err != nil {
			return err
		}

		// Set Guardian flag
		member.IsGuardian = member.Status == types.MembershipStatus_MemberElectorate && types.IsGuardianAddressFromBech32(guardians, member.Address)

		members = append(members, &member)
		return nil
	})

	return &types.QueryMembersResponse{Members: members}, nil
}
