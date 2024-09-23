package basics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginTesting(t *testing.T) {
	fmt.Println("1. This will be displayed when running tests")
	t.Log("2. This will be displayed when running tests")
}

func TestLogfInTesting(t *testing.T) {
	expected := "expected"
	actual := "actual"

	// Use assert.Equal to check equality; it returns a boolean status.
	if assert.NotEqual(t, expected, actual, "Output") {
		t.Logf("Expected: %q", expected)
		t.Logf("Actual: %q", actual)
	}
}

func TestTruth(t *testing.T) {
	i := 1 + 1

	if i != 2 {
		t.Error("everything I know is wrong")
	}
}
