package keeper

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/noria-net/module-membership/x/membership/types"
	"github.com/stretchr/testify/suite"
)

type CalculateVoteResultsTestSuite struct {
	suite.Suite
	proposor  *types.Member
	proposal  *govtypes_v1.Proposal
	govParams govtypes_v1.Params
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
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingWeight)

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
	suite.Assert().Equal("1", tallyResults.GetYesCount())
}

// Test Case: Guardians vote yes, members vote no, and proposal passes
func (suite *CalculateVoteResultsTestSuite) Test_GuardiansVoteYesMembersVoteNoAndProposalPasses() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingWeight)

	addVote(memberResults, govtypes_v1.OptionNo)
	addVote(guardianResults, govtypes_v1.OptionYes)

	passes, burnDeposits, tallyResults := calculateVoteResults(*suite.proposal,
		suite.govParams,
		memberResults,
		guardianResults,
		memberPower,
		guardianPower)

	printTallyResultsToConsole(tallyResults)
	suite.Assert().True(passes, "proposal should pass, got %v", passes)
	suite.Assert().Equal(suite.govParams.BurnProposalDepositPrevote, burnDeposits)
	suite.Assert().Equal("51", tallyResults.GetYesCount())
}

// Test Case: Guardians vote no, members vote yes, and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_GuardiansVoteNoMembersVoteYesAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingWeight)

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
	printTallyResultsToConsole(tallyResults)
	suite.Assert().Equal("49", tallyResults.GetYesCount())
}

// Test Case: No votes are cast and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_NoVotesAreCastAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingWeight)

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
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(4, 2, totalVotingWeight)

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
	suite.Assert().Equal("255", tallyResults.GetYesCount())
}

// Test Case: Veto threshold exceeded and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_VetoThresholdExceededAndProposalFails() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(4, 2, totalVotingWeight)

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
	printTallyResultsToConsole(tallyResults)
	suite.Assert().Equal("5", tallyResults.GetYesCount())
	suite.Assert().Equal("5", tallyResults.GetNoWithVetoCount())
}

// Test Case: All voters abstain and proposal fails
func (suite *CalculateVoteResultsTestSuite) Test_AllVotersAbstain() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(4, 2, totalVotingWeight)

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
	totalVotingWeight := math.LegacyMustNewDecFromStr("0.51")
	memberPower, guardianPower := calculateVotePower(4, 2, totalVotingWeight)

	// All members vote Yes
	addVote(memberResults, govtypes_v1.OptionYes)
	addVote(memberResults, govtypes_v1.OptionYes)

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
