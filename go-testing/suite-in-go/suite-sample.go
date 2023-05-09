package suite_sample

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type ExampleTestSuite struct {
	suite.Suite
}

// This is an example test that will always succeed
func (suite *ExampleTestSuite) TestExample() {
	suite.Equal(true, true)
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
