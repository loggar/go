package array1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterateArray(t *testing.T) {
	code := [7]rune{'#', '5', 'g', 't', 'm', 'y', '6'}

	s := ""

	for i := 0; i < len(code); i++ {
		s = s + string(code[i])
	}

	expected := "#5gtmy6"
	assert.Equal(t, expected, s, "%s should equal %s", s, expected)

	s2 := ""

	for _, c := range code {
		s2 = s2 + string(c)
	}

	assert.Equal(t, expected, s2, "%s should equal %s", s, expected)
}
