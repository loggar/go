package regex_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndexFromRegex(t *testing.T) {
	exp, err := regexp.Compile(`\b\d{5}(?:[-\s]\d{4})?\b`)
	assert.NoError(t, err, "compile regexp")

	pincode_file, err := os.ReadFile("data.ex.pincode.1.txt")
	assert.NoError(t, err, "read a file")

	match_index := exp.FindIndex(pincode_file)
	// log.Printf("%T\n", match_index) // []int
	// log.Println(match_index) //[9 19]
	assert.Equal(t, 2, len(match_index), "found matches ([]int)")

	var sliced_string []byte
	if len(match_index) > 0 {
		// Get the slice of the original string from start to end index
		sliced_string = pincode_file[match_index[0]:match_index[1]]
	}
	assert.Equal(t, []byte("12345-1234"), sliced_string, "get the text in the source string")
}

func TestFindTextFromIndex(t *testing.T) {
	exp, err := regexp.Compile(`\b\d{5}(?:[-\s]\d{4})?\b`)
	assert.NoError(t, err, "compile regexp")

	pincode_file, err := os.ReadFile("data.ex.pincode.2.txt")
	assert.NoError(t, err, "read a file")

	match_indexes := exp.FindAllIndex(pincode_file, -1)
	// log.Printf("%T\n", match_indexes) // [][]int
	// log.Println(match_indexes) // [[9 19] [232 237]]

	assert.Equal(t, 2, len(match_indexes), "found matches ([][]int)")
}
