package keeper

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/types"
)

// Helper functions
func areAllOptionsZero(voteOptions voteOptions) bool {
	for _, option := range voteOptions {
		if !option.IsZero() {
			return false
		}
	}
	return true
}

func areAllOptionsZeroDec(options map[govtypes_v1.VoteOption]math.LegacyDec) bool {
	for _, option := range options {
		if !option.IsZero() {
			return false
		}
	}
	return true
}

// isOnlyOneOptionSelected returns true if the voteOptions map contains only one option
func isOnlyOneOptionSelected(voteOptions voteOptions, option govtypes_v1.VoteOption) bool {
	// First make sure all the other options are zero
	for vote, voteValue := range voteOptions {
		if vote == option {
			continue
		}
		if !voteValue.IsZero() {
			return false
		}
	}
	// Now check the option we are interested in is equal to 1
	return voteOptions[option].Equal(math.NewInt(1))
}

// createMember creates a member with the given address
func createMember(address string) *types.Member {
	baseAccount := authtypes.NewBaseAccountWithAddress(sdk.AccAddress(address))
	member := types.NewMemberAccountWithDefaultMemberStatus(
		baseAccount,
	)
	return member
}

// mustCreateProposal creates a proposal and expects no errors
func mustCreateProposal(proposalId uint64, proposor string) *govtypes_v1.Proposal {
	member := createMember(proposor)
	// Time must be 30 seconds in the future
	depositEndTime := time.Now().Add(time.Second * 30)
	proposal, err := govtypes_v1.NewProposal([]sdk.Msg{},
		proposalId,
		time.Now(),
		depositEndTime,
		"Test Proposal Metadata",
		"Test Proposal Title",
		"Test Proposal Summary",
		member.GetAddress())

	if err != nil {
		panic(err)
	}
	return &proposal
}

// addVote adds a vote to the vote options
func addVote(voteOptions voteOptions, option govtypes_v1.VoteOption) {
	voteOptions[option] = voteOptions[option].Add(sdk.NewInt(1))
}

// printVoteOptionsToConsole prints the vote options to the console
func printVoteOptionsToConsole(results voteOptions) {
	for option, value := range results {
		println(option.String(), value.String())
	}
}

// printTallyResultsToConsole prints the tally results to the console
func printTallyResultsToConsole(results govtypes_v1.TallyResult) {
	println("Yes", results.GetYesCount())
	println("Abstain", results.GetAbstainCount())
	println("No", results.GetNoCount())
	println("NoWithVeto", results.GetNoWithVetoCount())
}
