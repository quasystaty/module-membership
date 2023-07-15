package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noria-net/module-membership/x/membership/types"
)

// IsGuardian check if the given address is a whitelisted guardian
// NOTE: It does not check if this address represents a member or if
// the member has a status of MemberElectorate
func (k Keeper) IsWhitelistedGuardian(ctx sdk.Context, addr sdk.AccAddress) bool {
	whitelist := k.GetGuardianWhitelist(ctx)
	allGuardians := types.ParseGuardianWhitelist(whitelist)
	for _, guardianAddress := range allGuardians {
		if guardianAddress.Equals(addr) {
			return true
		}
	}
	return false
}

// GetGuardians returns all guardians of the electorate
// NOTE: Only valid members with membership status of MemberElectorate are returned
func (k Keeper) GetGuardians(ctx sdk.Context) []*types.Member {
	whitelist := k.GetGuardianWhitelist(ctx)

	// Initialise guardians to an empty slice
	guardians := make([]*types.Member, 0)

	// Exclude any addresses that are not members
	for _, addr := range types.ParseGuardianWhitelist(whitelist) {
		// Get the member
		member, found := k.GetMemberAccount(ctx, addr)
		// Guardian must be a member with a status of MemberElectorate
		if found && member.Status == types.MembershipStatus_MemberElectorate {
			member.IsGuardian = true
			guardians = append(guardians, &member)
		}
	}

	return guardians
}

// GetGuardianAddresses returns all guardian addresses
func (k Keeper) GetGuardianAddresses(ctx sdk.Context) []sdk.AccAddress {
	whitelist := k.GetGuardianWhitelist(ctx)
	return types.ParseGuardianWhitelist(whitelist)
}
