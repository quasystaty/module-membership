package keeper

import (
	"testing"

	"cosmossdk.io/math"
	govtypes_v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/suite"
)

type ScaleTallyResultsToIntegerMapTestSuite struct {
	suite.Suite
}

func (suite *ScaleTallyResultsToIntegerMapTestSuite) SetupTest() {
}

func (suite *ScaleTallyResultsToIntegerMapTestSuite) Test_EveryoneVotesYes() {
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
	scaled := scaleTallyResultsToIntegerMap(combined.results)
	// printVoteOptionsToConsole(scaled)
	// Assert that yes count is 1
	suite.Equal(int64(1), scaled[govtypes_v1.OptionYes].Int64())
	// All others are zero
	suite.Equal(int64(0), scaled[govtypes_v1.OptionNo].Int64())
	suite.Equal(int64(0), scaled[govtypes_v1.OptionAbstain].Int64())
	suite.Equal(int64(0), scaled[govtypes_v1.OptionNoWithVeto].Int64())
}

func (suite *ScaleTallyResultsToIntegerMapTestSuite) Test_GuardianVotesYesMemberVotesNo() {
	memberResults := NewEmptyVoteOptions()
	guardianResults := NewEmptyVoteOptions()
	totalVotingPower := math.LegacyMustNewDecFromStr("0.51")
	// Two members, one guardian, 51% voting power
	memberPower, guardianPower := calculateVotePower(2, 1, totalVotingPower)
	// Guardian votes yes
	addVote(guardianResults, govtypes_v1.OptionYes)
	// Member votes yes
	addVote(memberResults, govtypes_v1.OptionNo)

	// Execute test
	combined := calculateCombinedTallyResults(memberResults, guardianResults, memberPower, guardianPower)
	scaled := scaleTallyResultsToIntegerMap(combined.results)
	// printVoteOptionsToConsole(scaled)

	suite.Equal(int64(51), scaled[govtypes_v1.OptionYes].Int64())
	suite.Equal(int64(49), scaled[govtypes_v1.OptionNo].Int64())
	suite.Equal(int64(0), scaled[govtypes_v1.OptionAbstain].Int64())
	suite.Equal(int64(0), scaled[govtypes_v1.OptionNoWithVeto].Int64())
}

func TestScaleTallyResultsToIntegerMapTestSuite(t *testing.T) {
	suite.Run(t, new(ScaleTallyResultsToIntegerMapTestSuite))
}
