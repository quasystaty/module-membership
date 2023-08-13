package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noria-net/module-membership/x/membership/types"

	"cosmossdk.io/errors"
)

func (k msgServer) Enroll(goCtx context.Context, msg *types.MsgEnroll) (*types.MsgEnrollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Must have a valid address
	enrollee := sdk.MustAccAddressFromBech32(msg.Creator)

	// Must have a valid nickname length (if set)
	if len(msg.Nickname) > types.NicknameMaxLength {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "nickname too long")
	}

	// Save it to the store
	err := k.AppendMember(ctx, enrollee); if err != nil {
		return nil, err
	}

	// Set the user's nickname
	k.SetMemberNickname(ctx, enrollee, msg.Nickname)

	// Publish events
	err = ctx.EventManager().EmitTypedEvents(
		// A new member was enrolled
		&types.EventMemberEnrolled{MemberAddress: enrollee.String()},
		// A member's citizenship status has changed
		&types.EventMemberStatusChanged{
			MemberAddress:  enrollee.String(),
			Operator:       enrollee.String(),
			Status:         types.MembershipStatus_MemberElectorate,
			PreviousStatus: types.MembershipStatus_MemberStatusEmpty,
		},
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgEnrollResponse{}, nil
}
