package regex_test

import (
	"regexp"
	"testing"
)

// hasUnicodeAndSpecialChars checks if a string contains any unicode or some of special characters.
func hasUnicodeAndSpecialChars(str string) bool {
	// \p{L} matches any kind of letter from any script.
	re := regexp.MustCompile(`^[\p{L}\d \-\.\:\[\]()_,!?&@<>*$#'/]{1,64}$`)
	return re.MatchString(str)
}

func TestHasUnicodeAndSpecialChars(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"1", true},
		{"0123456789012345678901234567890123456789012345678901234567891234", true},
		{"Hello, world!", true},
		{"반갑다, 세계여!", true},
		{"こんにちは", true},
		{"-Saving-", true},
		{"This. and That.", true},
		{"This & that", true},
		{"()", true},
		{"<>@-_:,.#$&!?()*[]", true},
		{"<Custom>@-_:,.#$&!?*(label)", true},
		{"'", true},
		{"Hello/World", true},
		{"", false},
		{"01234567890123456789012345678901234567890123456789012345678912345", false},
		{"\\", false},
		{"|", false},
		{"Hello|World", false},
		{"Hello~World", false},
		{"Hello`World", false},
		{"Hello;World", false},
		{"Hello\"World", false},
	}

	for _, tc := range testCases {
		result := hasUnicodeAndSpecialChars(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %v for input '%s', got %v", tc.expected, tc.input, result)
		}
	}
}
