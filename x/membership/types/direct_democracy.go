package types

import (
	fmt "fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	MINIMUM_TOTAL_VOTING_WEIGHT = "0"
	MAXIMUM_TOTAL_VOTING_WEIGHT = "1"
)

// DirectDemocracy is the struct that contains the direct democracy state
func DefaultDirectDemocracy() DirectDemocracy {
	return DirectDemocracy{
		TotalVotingWeight: math.LegacyZeroDec(),
		Guardians:         []string{},
	}
}

func (dd DirectDemocracy) Validate() error {

	// totalVotingWeight must be between 0 and 100, inclusive
	if dd.TotalVotingWeight.LT(math.LegacyZeroDec()) || dd.TotalVotingWeight.GT(math.LegacyMustNewDecFromStr("100")) {
		return fmt.Errorf("total voting weight must be between 0 and 100, inclusive: %s", dd.TotalVotingWeight)
	}

	// Keep a temporary map of guardian addresses
	addresses := make(map[string]bool)

	for _, guardian := range dd.Guardians {
		// Every guardian address must be a valid address
		if _, err := sdk.AccAddressFromBech32(guardian); err != nil {
			return fmt.Errorf("invalid guardian address: %s", guardian)
		}

		// Cannot have duplicate guardian addresses
		if _, ok := addresses[guardian]; ok {
			return fmt.Errorf("duplicate guardian address: %s", guardian)
		}

		// add this address to the temporary map
		addresses[guardian] = true
	}

	return nil
}
