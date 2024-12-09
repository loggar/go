package array1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayNilLen(t *testing.T) {
	var arr1 []int
	var arr2 *[]int

	arr3 := &[]int{}

	arr4 := return_nil_array()
	arr5 := []int{}

	assert.Nil(t, arr1, "The array should be nil")
	assert.Len(t, arr1, 0, "The array length should be 0")

	assert.Nil(t, arr2, "The array should be nil")

	// if arr2 != nil {
	// 	assert.Len(t, arr2, 0, "The array length should be 0")
	// }

	assert.NotNil(t, arr3, "The array should not be nil")
	assert.Len(t, *arr3, 0, "The array length should be 0")

	assert.Nil(t, arr4, "The array should be nil")
	assert.Len(t, arr4, 0, "The array length should be 0")

	assert.NotNil(t, arr5, "The array should not be nil")
	assert.Len(t, arr5, 0, "The array length should be 0")
}

func return_nil_array() []int {
	return nil
}
