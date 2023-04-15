package slice1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceLenCap(t *testing.T) {
	var marks []int

	assert.Equal(t, 0, len(marks), "len(%v) should equal %d", marks, 0)
	assert.Equal(t, 0, cap(marks), "cap(%v) should equal %d", marks, 0)

	frameworks := []string{"Django", "Laravel", "Flask", "Rails"}

	assert.Equal(t, 4, len(frameworks), "len(%v) should equal %d", frameworks, 4)
	assert.Equal(t, 4, cap(frameworks), "cap(%v) should equal %d", frameworks, 4)

	var langs = make([]string, 3, 5)
	langs[0], langs[1], langs[2] = "Python", "Go", "Javascript"

	assert.Equal(t, 3, len(langs), "assert len(%v)", langs)
	assert.Equal(t, 5, cap(langs), "assert cap(%v)", langs)
}
