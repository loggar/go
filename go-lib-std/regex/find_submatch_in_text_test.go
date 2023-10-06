package regex_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubmatchInTextFile(t *testing.T) {
	exp := regexp.MustCompile(
		`([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
			`|(mail|email)` +
			`|(\d{1,3})`,
	)

	email_file, err := os.ReadFile("data.ex.email.txt")
	assert.NoError(t, err, "read a file")

	match := exp.FindSubmatch(email_file)
	matches := exp.FindAllSubmatch(email_file, -1)

	assert.Equal(t, 6, len(match), "found sub-match")
	assert.Equal(t, 15, len(matches), "found sub-matches")
}
