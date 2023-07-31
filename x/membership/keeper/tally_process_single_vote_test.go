package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/suite"
)

const address_1 = "cosmos179774l7446j9jvvxnecg0qt6g3t00syyllmfgt"

type ProcessSingleVoteTestSuite struct {
	suite.Suite
}

// Setup
func (suite *ProcessSingleVoteTestSuite) SetupTest() {
}

// Voter is not a denom member
func (suite *ProcessSingleVoteTestSuite) Test_VoterIsNotDenomMember() {
	member := createMember(address_1)
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	option := govtypes_v1.NewWeightedVoteOption(govtypes_v1.OptionYes, sdk.NewDec(1))
	vote := govtypes_v1.NewVote(1, member.GetAddress(), govtypes_v1.WeightedVoteOptions{option}, "")
	err := processSingleVote(vote, member, false, memberResults, guardianResults)

	// Verify that an error is returned and that the vote is not counted
	suite.Assert().ErrorContains(err, "voter is not a member of the electorate")
	suite.Assert().True(areAllOptionsZero(memberResults))
	suite.Assert().True(areAllOptionsZero(guardianResults))
}

// Voter is not eligible to vote
func (suite *ProcessSingleVoteTestSuite) Test_VoterIsNotEligibleToVote() {
	option := govtypes_v1.NewWeightedVoteOption(govtypes_v1.OptionYes, sdk.NewDec(1))

	var member *types.Member
	var memberResults voteOptions
	var guardianResults voteOptions
	var vote govtypes_v1.Vote

	// Test all membership statuses
	for _, status := range types.MembershipStatus_value {

		// skip this loop if the status is electorate
		if status == int32(types.MembershipStatus_MemberElectorate) {
			continue
		}

		member = createMember(address_1)
		memberResults = NewEmptyVoteOptions()
		guardianResults = NewEmptyVoteOptions()
		vote = govtypes_v1.NewVote(1, member.GetAddress(), govtypes_v1.WeightedVoteOptions{option}, "")
		member.Status = types.MembershipStatus(status)

		err := processSingleVote(vote, member, true, memberResults, guardianResults)

		// Verify that an error is returned and that the vote is not counted
		suite.Assert().ErrorContains(err, "member is not eligible to vote")
		suite.Assert().True(areAllOptionsZero(memberResults))
		suite.Assert().True(areAllOptionsZero(guardianResults))
	}
}

// Voting weight is invalid
func (suite *ProcessSingleVoteTestSuite) Test_VotingWeightIsInvalid() {
	member := createMember(address_1)
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	option_yes := govtypes_v1.NewWeightedVoteOption(govtypes_v1.OptionYes, sdk.NewDec(1))
	option_no := govtypes_v1.NewWeightedVoteOption(govtypes_v1.OptionNo, sdk.NewDec(1))
	vote := govtypes_v1.NewVote(1,
		member.GetAddress(),
		govtypes_v1.WeightedVoteOptions{option_yes, option_no},
		"")
	err := processSingleVote(vote, member, true, memberResults, guardianResults)

	// Verify that an error is returned and that the vote is not counted
	suite.Assert().ErrorContains(err, "invalid voting weight")
	suite.Assert().True(areAllOptionsZero(memberResults))
	suite.Assert().True(areAllOptionsZero(guardianResults))
}

// Valid vote is processed
func (suite *ProcessSingleVoteTestSuite) Test_ValidVoteIsProcessed() {
	member := createMember(address_1)
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	option := govtypes_v1.NewWeightedVoteOption(govtypes_v1.OptionYes, sdk.NewDec(1))
	vote := govtypes_v1.NewVote(1, member.GetAddress(), govtypes_v1.WeightedVoteOptions{option}, "")
	err := processSingleVote(vote, member, true, memberResults, guardianResults)

	// Verify that an error is returned and that the vote is not counted
	suite.Assert().NoError(err)
	suite.Assert().True(isOnlyOneOptionSelected(memberResults, govtypes_v1.OptionYes))
	suite.Assert().True(areAllOptionsZero(guardianResults))
}

// Run test suite
func TestProcessSingleVoteTestSuite(t *testing.T) {
	suite.Run(t, new(ProcessSingleVoteTestSuite))
}
