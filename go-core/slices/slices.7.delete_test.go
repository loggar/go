package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceDelete(t *testing.T) {
	marks := []int{80, 85, 90, 75, 60}
	fmt.Println(marks)

	var index = 3
	elem := marks[index]

	// append in such a way that the element to be removed is excluded
	marks = append(marks[:index], marks[index+1:]...)

	fmt.Printf("The element %d was deleted.\n", elem)
	fmt.Println(marks)

	assert.Equal(t, []int{80, 85, 90, 60}, marks, "assert %v after delete an element", marks)
}
