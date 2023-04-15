package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceModify(t *testing.T) {
	word := []byte{'f', 'u', 'z', 'z', 'y'}

	fmt.Printf("%s\n", word)
	word[0] = 'b'
	word[len(word)-1] = 'z'
	fmt.Printf("%s\n", word)

	assert.Equal(t, "buzzz", string(word), "assert string(%v)", word)
}
