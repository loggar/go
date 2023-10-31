package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapLen(t *testing.T) {
	var a map[string]int

	assert.Nil(t, a, "not initialized map should be nil")
	assert.Equal(t, 0, len(a), "not initialized map should have length 0")
}
func TestMapLiterals(t *testing.T) {
	char_freq := map[string]int{
		"M": 1,
		"e": 2,
		"t": 3,
	}

	assert.Equal(t, 1, char_freq["M"], "char_freq[%s] should equal %s", "M", 1)
	assert.Equal(t, 2, char_freq["e"], "char_freq[%s] should equal %s", "e", 2)
	assert.Equal(t, 3, char_freq["t"], "char_freq[%s] should equal %s", "t", 3)
}

func TestMapMake(t *testing.T) {
	marks := make(map[int]int)
	marks[65] = 7
	marks[95] = 11
	marks[80] = 13

	assert.Equal(t, 7, marks[65], "char_freq[%d] should equal %d", 65, 7)
	assert.Equal(t, 13, marks[80], "char_freq[%d] should equal %d", "e", 13)
}

func TestMapNew(t *testing.T) {
	name := new(map[byte]int)
	*name = map[byte]int{}
	name_map := *name

	name_map['m'] = 2
	name_map['e'] = 3
	name_map['t'] = 5

	assert.Equal(t, 2, name_map['m'], "char_freq[%b] should equal %d", 'm', 2)
	assert.Equal(t, 3, name_map['e'], "char_freq[%b] should equal %d", 'e', 3)
	assert.Equal(t, 5, name_map['t'], "char_freq[%b] should equal %d", 't', 5)
}
