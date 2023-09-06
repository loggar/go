package sql_utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ReplaceNamedPlaceholders(query string, args map[string]interface{}) string {
	re := regexp.MustCompile(`:\w+`)
	return re.ReplaceAllStringFunc(query, func(m string) string {
		key := m[1:] // remove the leading `:`
		val, ok := args[key]
		if !ok {
			return m
		}
		// Convert the value to string safely
		// Add more types as needed
		var replacement string
		switch v := val.(type) {
		case string:
			replacement = "'" + v + "'"
		case int, int64:
			replacement = fmt.Sprintf("%d", v)
		default:
			replacement = fmt.Sprintf("%v", v)
		}
		return replacement
	})
}

func ReplaceSQLPlaceholders(query string, args []interface{}) string {
	for _, arg := range args {
		var replacement string
		switch v := arg.(type) {
		case string:
			replacement = "'" + v + "'"
		case int, int64:
			replacement = fmt.Sprintf("%d", v)
		// Add more types as needed
		default:
			replacement = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", replacement, 1)
	}
	return query
}
