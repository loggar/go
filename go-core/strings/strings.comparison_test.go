package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparison(t *testing.T) {
	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gopher"

	isEqual := s1 == s2

	//"gopher" is not same as "Gopher" and hence `false`
	assert.False(t, isEqual, "should be false")

	// "gopher" is not equal to "Gopher" and hence `true`
	assert.True(t, s1 != s2, "should be true")

	// "Gopher" comes first lexicographically than "gopher" so return true
	// G -> 71 in ASCII and g -> 103
	assert.True(t, s2 < s3, "should be true")

	// "Gopher" is not greater than "gopher" as `G` comes first in ASCII table
	// So G value is less than g i.e. 71 > 103 which is false
	assert.False(t, s2 > s3, "should be false")
}

func TestCompare(t *testing.T) {
	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gopher"

	assert.Equal(t, 1, strings.Compare(s1, s2), "should be true")
	assert.Equal(t, 0, strings.Compare(s1, s3), "should be true")
	assert.Equal(t, -1, strings.Compare(s2, s3), "should be true")
}

func TestEqualFold(t *testing.T) {
	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gophy"

	assert.True(t, strings.EqualFold(s1, s2), "should be true")
	assert.False(t, strings.EqualFold(s1, s3), "should be true")
	assert.False(t, strings.EqualFold(s2, s3), "should be true")
}
