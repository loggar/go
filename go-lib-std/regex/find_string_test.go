package regex_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStringFromRegex(t *testing.T) {
	exp, err := regexp.Compile(`\b\d{5}(?:[-\s]\d{4})?\b`)
	assert.NoError(t, err, "compile regexp")

	pincode_file, err := os.ReadFile("data.ex.pincode.1.txt")
	assert.NoError(t, err, "read a file")

	match := exp.FindAll(pincode_file, -1)
	// log.Println(match) // [[49 50 51 52 53 45 49 50 51 52] [52 48 48 56 52]]
	matches := exp.FindAllString(string(pincode_file), -1)
	// log.Println(matches) //[12345-1234 40084]

	assert.Equal(t, 2, len(match), "found matches ([]byte)")
	assert.Equal(t, 2, len(matches), "found matches ([]string)")
}
