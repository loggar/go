package mapstring

import (
	"testing"

	"github.com/loggar/go/go-testing/advanced-go-testing-techniques/testdemo/suite"
)

func TestSuite(t *testing.T) {
	rem := make(Remote)
	suite.Test(t, suite.Config{SP: rem})
}
