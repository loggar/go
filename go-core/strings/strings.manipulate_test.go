package strings

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprintf(t *testing.T) {
	name := "peter"
	domain := "telecom"
	service := "ceo"

	email := fmt.Sprintf("%s.%s@%s.com", service, name, domain)

	expected := "ceo.peter@telecom.com"
	assert.Equal(t, expected, email, "strings.join(%s) should result %s", email, expected)
}

func TestBuilder(t *testing.T) {
	// Using Builder function

	c := []string{"j", "a", "v", "a"}
	var builder strings.Builder
	for _, item := range c {
		builder.WriteString(item)
	}

	assert.Equal(t, "java", builder.String(), "should equal")

	b := []byte{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range b {
		builder.WriteByte(item)
	}

	assert.Equal(t, "javascript", builder.String(), "should equal")
}

func TestBytesBuffer(t *testing.T) {
	// Using bytes buffer method

	var buf bytes.Buffer

	for i := 0; i < 2; i++ {
		buf.WriteString("ja")
	}
	assert.Equal(t, "jaja", buf.String(), "should equal")

	buf.WriteByte('r')
	assert.Equal(t, "jajar", buf.String(), "should equal")

	k := []rune{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range k {
		buf.WriteRune(item)
	}
	assert.Equal(t, "jajarscript", buf.String(), "should equal")
}
