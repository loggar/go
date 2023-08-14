package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func for_slice_range() int {
	sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

	j := 0

	for i, shark := range sharks {
		fmt.Println(i, shark)
		j++
	}

	return j
}

func Test_for_slice_range(t *testing.T) {
	count := for_slice_range()
	assert.Equal(t, 6, count, "loop using for-range")
}
