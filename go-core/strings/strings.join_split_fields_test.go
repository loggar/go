package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	sa := []string{"sharks", "crustaceans", "plankton"}
	saJoin := strings.Join(sa, ",")

	expected := "sharks,crustaceans,plankton"
	assert.Equal(t, expected, saJoin, "should equal")
}

func TestSplit(t *testing.T) {
	s := "sharks crustaceans plankton"
	ss := strings.Split(s, " ")

	expected := []string{"sharks", "crustaceans", "plankton"}
	assert.Equal(t, expected, ss, "should equal")
}

func TestFields(t *testing.T) {
	data := "  username password     email  date"
	fields := strings.Fields(data)

	expected := []string{"username", "password", "email", "date"}
	assert.Equal(t, expected, fields, "should equal")
}
