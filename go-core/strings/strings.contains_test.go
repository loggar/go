package strings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAndContainsAny(t *testing.T) {
	str := "javascript"
	substr := "script"
	s := "python"

	fmt.Println(strings.Contains(str, substr))
	fmt.Println(strings.Contains(str, s))

	fmt.Println(strings.ContainsAny(str, "joke"))
	fmt.Println(strings.ContainsAny(str, "xyz"))

	assert.False(t, strings.ContainsAny(str, ""), "should be false")
	assert.True(t, strings.ContainsAny(str, "xyzj"), "should be true")
}

func splitAfter(t *testing.T) {
	// Split method for splitting string into slice of string
	link := "meetgor.com/blog/golang/strings"
	fmt.Println(strings.Split(link, "/"))      // [meetgor.com blog golang strings]
	fmt.Println(strings.SplitAfter(link, "/")) // [meetgor.com/ blog/ golang/ strings]

	// SplitAfter method for splitting string into slice of string with the pattern
	data := "200kms50kms120kms"
	fmt.Println(strings.Split(data, "kms"))      // [200 50 120 ]
	fmt.Println(strings.SplitAfter(data, "kms")) // [200kms 50kms 120kms ]
}
