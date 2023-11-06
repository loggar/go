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

	var key2 byte = 'x'
	value2, exist2 := name_map[key2]
	assert.False(t, exist2, "%s should not exist", key)
	assert.Equal(t, 0, value2)
}

// MappingInfo is used to test struct arrays in a map
type MappingInfo struct {
	Path    string
	Version string
}

func TestMapStructValue(t *testing.T) {
	name_map := map[byte]MappingInfo{
		'm': {
			Path:    "a",
			Version: "1",
		},
	}

	var key byte = 'm'
	value, exist := name_map[key]

	assert.True(t, exist, "%s should exist", key)
	assert.Equal(t, "a", value.Path, "value should equal")

	var key2 byte = 'x'
	value2, exist2 := name_map[key2]
	assert.False(t, exist2, "should not exist")
	assert.Equal(t, MappingInfo{
		Path:    "",
		Version: "",
	}, value2)
}

func TestMapStructArrayValue(t *testing.T) {
	name_map := map[byte][]MappingInfo{
		'm': {
			{
				Path:    "a",
				Version: "1",
			},
		},
	}

	var key byte = 'm'
	value, exist := name_map[key]

	assert.True(t, exist, "%s should exist", key)
	assert.Equal(t, "a", value[0].Path, "value should equal")

	var key2 byte = 'x'
	value2, exist2 := name_map[key2]
	assert.False(t, exist2, "should not exist")
	assert.Nil(t, value2)
}
