package slice1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesComparison(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 3, 2}

	assert.Equal(t, a, b, "The two slices should be equal")
	assert.NotEqual(t, a, c, "The two slices should not be equal")
}
