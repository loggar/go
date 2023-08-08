package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parse_int(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}

func TestTypeSwitchInterface(t *testing.T) {
	t.Run("type switch interface as a function parameter", func(t *testing.T) {
		assert.Equal(t, 4, parse_int(4), "type switch interface")
		assert.Equal(t, 4, parse_int("4"), "type switch interface")
		assert.Equal(t, 4, parse_int(4.1234), "type switch interface")
	})
}
