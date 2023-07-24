package keeper

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/suite"
)

const address_1 = "cosmos179774l7446j9jvvxnecg0qt6g3t00syyllmfgt"

type ProcessSingleVoteTestSuite struct {
	suite.Suite
}

type CalculateVoteResultsTestSuite struct {
	suite.Suite
	proposor  *types.Member
	proposal  *govtypes_v1.Proposal
	govParams govtypes_v1.Params
}

type CalculateCombinedTallyResultsTestSuite struct {
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

// Setup
func (suite *CalculateVoteResultsTestSuite) SetupTest() {
	proposalId := uint64(1)
	suite.proposor = createMember(address_1)
	suite.proposal = mustCreateProposal(proposalId, suite.proposor.Address)
	period := time.Duration(30) * time.Second

	suite.govParams = govtypes_v1.NewParams(
		// Minimum Deposit
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000000))),
		// Max deposit period
		period,
		// Voting period
		period,
		// Quorum
		"0.334",
		// Threshold
		"0.5",
		// Veto
		"0.334",
		// Minimum initial deposit ratio
		"0.334",
		// Burn proposal deposit
		false,
		// Burn vote quorum
		false,
		// Burn vote veto
		false,
	)
}

// Test Case: Majority votes yes and proposal passes
func (suite *CalculateVoteResultsTestSuite) Test_MajorityVotesYesAndProposalPasses() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(2, 1, math.LegacyNewDec(51))

	// Everyone votes yes
	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(guardianResults, govtypes_v1.OptionYes)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().True(passes)
	suite.Assert().Equal(suite.govParams.BurnProposalDepositPrevote, burnDeposits)
	suite.Assert().Equal("100", tallyResults.GetYesCount())
}

// Test Case: Guardians vote yes, members vote no, and proposal passes
func (suite *CalculateVoteResultsTestSuite) Test_GuardiansVoteYesMembersVoteNoAndProposalPasses() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(2, 1, math.LegacyNewDec(51))

	addVote(memberResults, govtypes_v1.OptionNo)
	addVote(guardianResults, govtypes_v1.OptionYes)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().True(passes)
	suite.Assert().Equal(suite.govParams.BurnProposalDepositPrevote, burnDeposits)
	suite.Assert().Equal("51", tallyResults.GetYesCount())
}

// Test Case: Guardians vote no, members vote yes, and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_GuardiansVoteNoMembersVoteYesAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(2, 1, math.LegacyNewDec(51))

	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(guardianResults, govtypes_v1.OptionNo)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().False(burnDeposits)
	suite.Assert().Equal("49", tallyResults.GetYesCount())
}

// Test Case: No votes are cast and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_NoVotesAreCastAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(2, 1, math.LegacyNewDec(51))

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().False(burnDeposits)
	suite.Assert().Equal("0", tallyResults.GetYesCount())
}

// Test Case: Threshold not met and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_ThresholdNotMetAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(4, 2, math.LegacyNewDec(51))

	// One Guardian votes Yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Everyone else votes no
	addVote(guardianResults, govtypes_v1.OptionNo)
	addVote(memberResults, govtypes_v1.OptionNo)
	addVote(memberResults, govtypes_v1.OptionNo)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().False(burnDeposits)
	suite.Assert().Equal("25", tallyResults.GetYesCount())
}

// Test Case: Veto threshold exceeded and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_VetoThresholdExceededAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(4, 2, math.LegacyNewDec(51))

	// A guardian and a member Vetos
	addVote(guardianResults, govtypes_v1.OptionNoWithVeto)
	addVote(memberResults, govtypes_v1.OptionNoWithVeto)

	// Everyone else votes Yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	addVote(memberResults, govtypes_v1.OptionYes)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().Equal(burnDeposits, suite.govParams.BurnVoteVeto)
	suite.Assert().Equal("50", tallyResults.GetYesCount())
}

// Test Case: All voters abstain and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_AllVotersAbstain() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(4, 2, math.LegacyNewDec(51))

	// Everyone votes to abstain
	addVote(guardianResults, govtypes_v1.OptionAbstain)
	addVote(guardianResults, govtypes_v1.OptionAbstain)
	addVote(memberResults, govtypes_v1.OptionAbstain)
	addVote(memberResults, govtypes_v1.OptionAbstain)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().False(burnDeposits)
	suite.Assert().Equal("0", tallyResults.GetYesCount())
}

// Test Case: Only members vote and quorum is not reached
func (suite *CalculateVoteResultsTestSuite) Test_OnlyMembersVoteAndQuorumIsNotReached() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	memberPower, guardianPower := calculateVotePower(4, 2, math.LegacyNewDec(51))

	// All members vote Yes
	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(memberResults, govtypes_v1.OptionYes)

	// TODO: Apply the Total Voting Power weighting to the quorum calculation

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	suite.Assert().False(passes)
	suite.Assert().Equal(burnDeposits, suite.govParams.BurnVoteQuorum)
	suite.Assert().Equal("49", tallyResults.GetYesCount())
}

// Run test suite
func TestCalculateVoteResultsTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateVoteResultsTestSuite))
}

// Setup
func (suite *CalculateCombinedTallyResultsTestSuite) SetupTest() {
}

// Test Case: No voting results, empty combined results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_NoVotingResultsEmptyCombinedResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, math.LegacyNewDec(51))

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	suite.Assert().True(areAllOptionsZeroDec(combined.results))
	suite.Assert().True(combined.votingPower.IsZero())
}

// Test Case: Only member votes, combined results are member results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_OnlyMemberVotesCombinedResultsAreMemberResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyNewDec(51)
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Member votes yes
	addVote(memberResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionYes, math.LegacyMustNewDecFromStr("100").Sub(totalVotingPower)))

	// All others must be zero
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNo))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionAbstain))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNoWithVeto))
	// Voting power should be 49%
	suite.Assert().True(combined.votingPower.Equal(math.LegacyNewDec(49)))
}

// Test Case: Only guardian votes, combined results are guardian results
func (suite *CalculateCombinedTallyResultsTestSuite) Test_OnlyGuardianVotesCombinedResultsAreGuardianResults() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyNewDec(51)
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionYes, totalVotingPower))
	// All others must be zero
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNo))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionAbstain))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNoWithVeto))
	// Voting power should be 51%
	suite.Assert().True(combined.votingPower.Equal(totalVotingPower))
}

// Test Case: Member and guardian vote yes, combined results are combined yes
func (suite *CalculateCombinedTallyResultsTestSuite) Test_MemberAndGuardianVoteYesCombinedResultsAreCombinedYes() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyNewDec(51)
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Member votes yes
	addVote(memberResults, govtypes_v1.OptionYes)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	// Only one vote cast
	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionYes, math.LegacyNewDec(100)))
	// All others must be zero
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNo))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionAbstain))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNoWithVeto))
	// Voting power should be 100%
	suite.Assert().True(combined.votingPower.Equal(math.LegacyNewDec(100)))
}

// Test Case: Member votes yes and guardian votes no
func (suite *CalculateCombinedTallyResultsTestSuite) Test_MemberVotesYesAndGuardianVotesNo() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyNewDec(51)
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Member votes no
	addVote(memberResults, govtypes_v1.OptionNo)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)

	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionYes, totalVotingPower))
	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionNo, math.LegacyNewDec(49)))
	// All others must be zero
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionAbstain))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNoWithVeto))
	// Voting power should be 100%
	suite.Assert().True(combined.votingPower.Equal(math.LegacyNewDec(100)))
}

// Test Case: Two guardians, two members, one guardian votes no, everyone votes yes
func (suite *CalculateCombinedTallyResultsTestSuite) Test_TwoGuardiansTwoMembersOneGuardianVotesNoEveryoneVotesYes() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyNewDec(51)
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

	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionYes, math.LegacyMustNewDecFromStr("74.5")))
	suite.Assert().True(isOptionValueEqualTo(combined.results, govtypes_v1.OptionNo, math.LegacyMustNewDecFromStr("25.5")))
	// All others must be zero
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionAbstain))
	suite.Assert().True(isOptionValueZero(combined.results, govtypes_v1.OptionNoWithVeto))
	// Voting power should be 100%
	suite.Assert().True(combined.votingPower.Equal(math.LegacyNewDec(100)))
}

// Run test suite
func TestCalculateVotePowerTestSuite(t *testing.T) {
	suite.Run(t, new(CalculateCombinedTallyResultsTestSuite))
}

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

func isOptionValueEqualTo(options map[govtypes_v1.VoteOption]math.LegacyDec, option govtypes_v1.VoteOption, value math.LegacyDec) bool {
	return options[option].Equal(value)
}

func isOptionValueZero(options map[govtypes_v1.VoteOption]math.LegacyDec, option govtypes_v1.VoteOption) bool {
	return options[option].IsZero()
}

func createMember(address string) *types.Member {
	baseAccount := authtypes.NewBaseAccountWithAddress(sdk.AccAddress(address))
	member := types.NewMemberAccountWithDefaultMemberStatus(
		baseAccount,
		"address_1",
	)
	return member
}

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

func addVote(voteOptions voteOptions, option govtypes_v1.VoteOption) {
	voteOptions[option] = voteOptions[option].Add(sdk.NewInt(1))
}
