package membership

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ABCITestSuite struct {
	suite.Suite
}

// Setup
func (suite *ABCITestSuite) SetupTest() {
}

// Teardown
func (suite *ABCITestSuite) TearDownTest() {
}

func (suite *ABCITestSuite) TestProcessActiveProposal() {

}

// Test suite runner
func TestABCITestSuite(t *testing.T) {
	suite.Run(t, new(ABCITestSuite))
}
