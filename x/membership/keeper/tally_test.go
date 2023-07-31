package keeper

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestAllTestSuites(t *testing.T) {
	suite.Run(t, new(CalculateCombinedTallyResultsTestSuite))
	suite.Run(t, new(ScaleTallyResultsToIntegerMapTestSuite))
	suite.Run(t, new(ProcessSingleVoteTestSuite))
	suite.Run(t, new(CalculateVoteResultsTestSuite))
}
