package regex_test

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceString(t *testing.T) {
	str := "Gophy gophers go in the golang grounds"

	exp := regexp.MustCompile(`(Go|go)`)
	replaced_str := exp.ReplaceAllString(str, "to")
	assert.Equal(t, "tophy tophers to in the tolang grounds", replaced_str, "replace string")

	exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)

	assert.Equal(t, "hopphy hophop hop in the hoplang ghop", exp2.ReplaceAllString(str, "hop"), "replace string")
	assert.Equal(t, "Gophy go go in the golang g", exp2.ReplaceAllString(str, "$1"), "replace string")
	assert.Equal(t, "phy phers  in the lang g", exp2.ReplaceAllString(str, "$2"), "replace string")
	assert.Equal(t, "phy   in the lang grounds", exp2.ReplaceAllString(str, "$3"), "replace string")

	str = "Gophy gophers go in the golang grounds"
	assert.Equal(t, "Gophy gophers go in the golang g", exp2.ReplaceAllString(str, "$1$2"), "replace string")

	str = "Gophy gophers go in the golang cophers grounds"
	assert.Equal(t, "Gophy go go in the golang co grounds", exp2.ReplaceAllString(str, "$1$3"), "replace string")
}

func TestReplaceLiteralString(t *testing.T) {
	str := "Gophy gophers go in the golang cophers grounds"
	exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)

	// If we want to avoid expansion of the strings as we did in the previous example with $1 and other parameters:
	log.Println(exp2.ReplaceAllLiteralString(str, "$1"))
	assert.Equal(t, "$1phy $1$1 $1 in the $1lang co$1 g$1", exp2.ReplaceAllLiteralString(str, "$1"), "replace literal string")
}
