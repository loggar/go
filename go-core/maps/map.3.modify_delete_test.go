package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapModifyItems(t *testing.T) {
	char_freq := map[string]int{
		"M": 1,
		"e": 2,
		"t": 3,
	}

	char_freq["M"] = 4
	char_freq["t"] = 5

	assert.Equal(t, 4, char_freq["M"], "char_freq[%s] should equal %s", "M", 4)
	assert.Equal(t, 2, char_freq["e"], "char_freq[%s] should equal %s", "e", 2)
	assert.Equal(t, 5, char_freq["t"], "char_freq[%s] should equal %s", "t", 5)
}

func TestMapDeleteKeys(t *testing.T) {
	char_freq := map[string]int{
		"M": 1,
		"e": 2,
		"t": 3,
	}

	key := "M"
	delete(char_freq, key)

	value, exist := char_freq[key]

	assert.False(t, exist, "%s should not exist", key)
	assert.Equal(t, 0, value, "value should equal to %d", 0)
}
