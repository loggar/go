package regex

import (
	"regexp"
	"strings"
)

func normalizeWhitespace(s string) string {
	// Trim leading/trailing whitespace
	s = strings.TrimSpace(s)

	// Collapse multiple spaces/newlines into a single space
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	return s
}
