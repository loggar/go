package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceAccess(t *testing.T) {
	scores := []int{80, 85, 90, 75, 60, 56, 83}

	fmt.Println(scores)
	fmt.Println("From index 2 to 4", scores[2:5])
	fmt.Println("From index 0 to 2", scores[:3])
	fmt.Println("From index 3 to 5", scores[3:])

	assert.Equal(t, []int{90, 75, 60}, scores[2:5], "assert %v[2:5]", scores)
	assert.Equal(t, []int{80, 85, 90}, scores[:3], "assert %v[:3]", scores)
	assert.Equal(t, []int{75, 60, 56, 83}, scores[3:], "assert %v[3:]", scores)
}
