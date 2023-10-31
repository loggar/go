package slice1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(slice []int) []int {
	length := len(slice)
	reversed := make([]int, length)
	copy(reversed, slice)

	for i := 0; i < length/2; i++ {
		reversed[i], reversed[length-i-1] = reversed[length-i-1], reversed[i]
	}

	return reversed
}
func TestSliceReverse(t *testing.T) {
	s := []int{5, 4, 3, 2, 1}
	r := reverse(s)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, r, "should equal")
}
