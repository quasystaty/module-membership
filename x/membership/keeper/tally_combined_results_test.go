package keeper

import (
	"testing"

	"cosmossdk.io/math"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/suite"
)

type CalculateCombinedTallyResultsTestSuite struct {
	suite.Suite
}

// Setup
func (suite *CalculateCombinedTallyResultsTestSuite) SetupTest() {
}

// Test Case: No voting results, empty combined results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_NoVotingResultsEmptyCombinedResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	suite.Assert().True(areAllOptionsZeroDec(combined.results))
	suite.Equal(math.LegacyNewDec(0), combined.votingPower)
}

// Test Case: Only member votes, combined results are member results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_OnlyMemberVotesCombinedResultsAreMemberResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Member votes yes
	addVote(memberResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Equal(math.LegacyNewDec(1).Sub(totalVotingPower), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_YES])

	// All others must be zero
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_ABSTAIN])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO_WITH_VETO])
	// Voting power should be 49%
	suite.Equal(math.LegacyMustNewDecFromStr("0.49"), combined.votingPower)
}

// Test Case: Only guardian votes, combined results are guardian results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_OnlyGuardianVotesCombinedResultsAreGuardianResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Equal(totalVotingPower, combined.results[govtypes_v1.VoteOption_VOTE_OPTION_YES])
	// All others must be zero
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_ABSTAIN])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO_WITH_VETO])
	// Voting power should be 51%
	suite.Equal(math.LegacyMustNewDecFromStr("0.51"), combined.votingPower)
}

// Test Case: Member and guardian vote yes, combined results are combined yes
func (suite *CalculateCombinedTallyResultsTestSuite) Test_MemberAndGuardianVoteYesCombinedResultsAreCombinedYes() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Member votes yes
	addVote(memberResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Equal(math.LegacyNewDec(1), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_YES])
	// All others must be zero
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_ABSTAIN])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO_WITH_VETO])
	// Voting power should be 100%
	suite.Equal(math.LegacyNewDec(1), combined.votingPower)
}

// Test Case: Member votes yes and guardian votes no
func (suite *CalculateCombinedTallyResultsTestSuite) Test_MemberVotesYesAndGuardianVotesNo() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Member votes no
	addVote(memberResults, govtypes_v1.OptionNo)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	suite.Equal(totalVotingPower, combined.results[govtypes_v1.VoteOption_VOTE_OPTION_YES])
	suite.Equal(math.LegacyNewDec(1).Sub(totalVotingPower), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO])
	// All others must be zero
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_ABSTAIN])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO_WITH_VETO])
	// Voting power should be 100%
	suite.Equal(math.LegacyNewDec(1), combined.votingPower)
}

// Test Case: Two guardians, two members, one guardian votes no, everyone votes yes
func (suite *CalculateCombinedTallyResultsTestSuite) Test_TwoGuardiansTwoMembersOneGuardianVotesNoEveryoneVotesYes() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(4, 2, totalVotingPower)
	// One guardian votes no
	addVote(guardianResults, govtypes_v1.OptionNo)
	// Everyone else votes yes
	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(guardianResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	suite.Equal(math.LegacyMustNewDecFromStr("0.745"), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_YES])
	suite.Equal(math.LegacyMustNewDecFromStr("0.255"), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO])
	// All others must be zero
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_NO_WITH_VETO])
	suite.Equal(math.LegacyNewDec(0), combined.results[govtypes_v1.VoteOption_VOTE_OPTION_ABSTAIN])
	// Voting power should be 100%
	suite.Equal(math.LegacyNewDec(1), combined.votingPower)
}

// Run test suite
func TestCalculateCombinedTallyResultsTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateCombinedTallyResultsTestSuite))
}
