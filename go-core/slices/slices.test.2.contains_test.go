package slice1

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sliceContains(s []int, n int) bool {
	return slices.Contains(s, n)
}

func TestSlicesContains(t *testing.T) {
	numbers := []int{0, 42, 10, 8}

	assert.True(t, sliceContains(numbers, 10), "should be true")
	assert.False(t, sliceContains(numbers, 99), "should be false")
}
