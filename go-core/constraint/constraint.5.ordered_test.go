package constraint_test

import (
	"testing"

	"github.com/loggar/go/go-core/constraint"
	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	data := []int{1, 3, 4, 4, 5, 8, 7, 3, 2}     // sample array
	data32 := []int32{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array
	data64 := []int64{1, 3, 4, 4, 5, 8, 7, 3, 2} // sample array

	assert.True(t, constraint.FindDuplicate(data), "%v", data)
	assert.True(t, constraint.FindDuplicate(data32), "%v", data32)
	assert.True(t, constraint.FindDuplicate(data64), "%v", data64)
}
