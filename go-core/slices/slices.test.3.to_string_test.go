package slice1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlicesJoin(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	result := strings.Join(slice, ", ")

	assert.Equal(t, "apple, banana, cherry", result, "[]string join]")

	sliceEmpty := []string{}
	joinedEmptySlice := strings.Join(sliceEmpty, ", ")

	assert.Equal(t, "", joinedEmptySlice, "empty []string join]")
}
