package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GuardianAddresses is a slice of the guardian members
type GuardianAddresses []sdk.AccAddress

// IsGuardianAddressFromBech32 check if the given address is a whitelisted guardian
func IsGuardianAddressFromBech32(guardians GuardianAddresses, address string) bool {
	addr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	for _, guardian := range guardians {
		if guardian.Equals(addr) {
			return true
		}
	}
	return false
}

func ParseGuardianWhitelist(whitelist string) (res GuardianAddresses) {

	rawAddresses := strings.Split(whitelist, ",")

	// set res to an empty slice of sdk.AccAddress
	res = make([]sdk.AccAddress, 0, len(rawAddresses))

	// convert all addresses to sdk.AccAddress
	for _, rawAddress := range rawAddresses {
		// Skip empty addresses
		if len(rawAddress) == 0 {
			continue
		}
		// parse the string into an sdk.AccAddress
		address, err := sdk.AccAddressFromBech32(rawAddress)
		if err != nil {
			panic(err)
		}
		res = append(res, address)
	}
	return res
}
