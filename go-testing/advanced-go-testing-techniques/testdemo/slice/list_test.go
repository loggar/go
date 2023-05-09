package slice

import (
	"testing"

	"github.com/loggar/go/go-testing/advanced-go-testing-techniques/testdemo/suite"
)

func TestSuite(t *testing.T) {
	loc := make(Local, 0)
	suite.Test(t, suite.Config{SP: &loc})
}
