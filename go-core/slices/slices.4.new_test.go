package slice1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceNew(t *testing.T) {
	langs := new([3]string)[0:2]

	langs[0], langs[1] = "Python", "Go"
	assert.Equal(t, 2, len(langs), "assert len(%v)", langs)
	assert.Equal(t, 3, cap(langs), "assert cap(%v)", langs)

	langs = append(langs, "Java", "Kotlin", "PHP")

	assert.Equal(t, 5, len(langs), "assert len(%v)", langs)
	assert.Equal(t, 6, cap(langs), "assert cap(%v)", langs)
}
