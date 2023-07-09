package keeper

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/noria-net/module-membership/x/membership/types"
)

func (k Keeper) GetMemberAccount(ctx sdk.Context, address sdk.AccAddress) (member *types.Member, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.MemberKey(address)

	var b []byte
	if b = store.Get(key); b == nil {
		return member, false
	}

	k.cdc.Unmarshal(b, member)

	// Get member guardianship status
	member.IsGuardian = k.IsWhitelistedGuardian(ctx, address) &&
		member.Status == types.MembershipStatus_MemberElectorate

	return member, true
}

func (k Keeper) AppendMember(ctx sdk.Context, address sdk.AccAddress, newMember types.Member) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.MemberKey(address)

	// Fetch member counts
	memberCount := k.GetMemberCount(ctx)
	memberStatusCount := k.GetMemberStatusCount(ctx, newMember.Status)

	// Marshal and Set
	memberData := k.cdc.MustMarshal(&newMember)
	store.Set(key, memberData)

	// Bump member count
	k.SetMemberCount(ctx, memberCount+1)

	// Bump member status count
	k.SetMemberStatusCount(ctx, newMember.Status, memberStatusCount+1)
}

func (k Keeper) IsMemberByBech32Address(ctx sdk.Context, bech32Address string) bool {
	address := sdk.AccAddress(bech32Address)
	return k.IsMember(ctx, address)
}

func (k Keeper) IsMember(ctx sdk.Context, address sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.MemberKey(address)
	return store.Has(key)
}

func (k Keeper) GetMemberCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.MemberCountKey)
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetMemberCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := sdk.Uint64ToBigEndian(count)
	store.Set(types.MemberCountKey, bz)
}

func (k Keeper) GetMemberStatusCount(ctx sdk.Context, s types.MembershipStatus) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.MemberStatusCountKey(s))
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetMemberStatusCount(ctx sdk.Context, s types.MembershipStatus, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := sdk.Uint64ToBigEndian(count)
	store.Set(types.MemberStatusCountKey(s), bz)
}

func (k Keeper) UpdateMemberStatus(ctx sdk.Context, target sdk.AccAddress, s types.MembershipStatus) error {
	// Fetch the member
	member, found := k.GetMemberAccount(ctx, target)

	// Member must exist
	if !(found) {
		return errors.Wrapf(sdkerrors.ErrUnauthorized, "member not found: %s", target.String())
	}

	// Must be a valid status transition
	if !member.Status.CanTransitionTo(s) {
		return errors.Wrapf(types.ErrMembershipStatusChangeNotAllowed, "transition %s is not allowed", s.DescribeTransition(s))
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	key := types.MemberKey(target)

	// Fetch status counts
	oldStatus := member.Status
	newStatus := s
	oldStatusCount := k.GetMemberStatusCount(ctx, oldStatus)
	newStatusCount := k.GetMemberStatusCount(ctx, newStatus)

	// Update the member's status
	member.Status = s

	// Marshal and Set
	memberData := k.cdc.MustMarshal(member)
	store.Set(key, memberData)

	// Update the status counts
	k.SetMemberStatusCount(ctx, oldStatus, oldStatusCount-1)
	k.SetMemberStatusCount(ctx, newStatus, newStatusCount+1)

	// Publish an update event
	ctx.EventManager().EmitTypedEvent(
		// A member's citizenship status has changed
		&types.EventMemberStatusChanged{
			MemberAddress: target.String(),
			// TODO: Change this
			Operator:       "",
			Status:         types.MembershipStatus_MemberElectorate,
			PreviousStatus: types.MembershipStatus_MemberStatusEmpty,
		},
	)

	return nil
}
