package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapLength(t *testing.T) {
	char_freq := map[string]int{
		"M": 1,
		"e": 2,
		"t": 3,
	}

	assert.Equal(t, 3, len(char_freq), "len should equal to %d", "M", 3)
}

func TestMapKey(t *testing.T) {
	name_map := map[byte]int{
		'm': 2,
		'e': 3,
		't': 5,
	}

	var key byte = 't'
	value, exist := name_map[key]

	assert.True(t, exist, "%s should exist", key)
	assert.Equal(t, 5, value, "value should equal to %d", 5)
}
