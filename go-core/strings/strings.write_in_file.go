package strings

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndWriteIfFail(t *testing.T) {
	expected := "  string1\n"
	actual := "  string1  \n"

	// Use assert.Equal to check equality; it returns a boolean status.
	if !assert.Equal(t, expected, actual, "Output Test") {
		// Write them to files:
		_ = os.WriteFile("test_fail_expected.txt", []byte(expected), 0644)
		_ = os.WriteFile("test_fail_actual.txt", []byte(actual), 0644)
	}
}
