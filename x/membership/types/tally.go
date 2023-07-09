package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// TallyStats is a struct containing the stats of a tally
type TallyStats struct {
	// NumVotes is the number of votes cast
	NumVotes uint64
	// NumMembers is the number of members who voted
	NumMembers uint64
	// NumGuardians is the number of guardians who voted
	NumGuardians uint64
	// NumEligibleMembers is the number of members who are eligible to vote
	NumEligibleMembers uint64
	// NumEligibleVotes is the number of votes we can include in the tally
	NumEligibleVotes uint64
	// TotalVotingPower is the total weight of all votes (including higher-weight individuals, called Guardians)
	TotalVotingPower sdk.Dec
}

// NewEmptyTallyStats creates a new TallyStats object with all values set to zero
func NewEmptyTallyStats() TallyStats {
	return TallyStats{
		NumVotes:           0,
		NumMembers:         0,
		NumGuardians:       0,
		NumEligibleMembers: 0,
		NumEligibleVotes:   0,
		TotalVotingPower:   sdk.ZeroDec(),
	}
}
