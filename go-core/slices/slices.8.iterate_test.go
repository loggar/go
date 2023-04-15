package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterate(t *testing.T) {
	code := [7]rune{'g', 'o', 'l', 'a', 'n', 'g'}

	s1, s2, s3 := "", "", ""

	// Using three statements for loop
	for i := 0; i < len(code); i++ {
		s1 += string(code[i])
	}

	// Using Range-based for loop
	for _, c := range code {
		s2 += string(c)
	}

	//Using for loop with range
	start, i, end := 2, 2, 5
	for range code[start:end] {
		fmt.Printf("Element at index %d = %c \n", i, code[i])
		s3 += string(code[i])
		i++
	}

	expected := "golang\x00"
	expectedRange := "lan"
	assert.Equal(t, expected, s1, "%s should equal %s", s1, expected)
	assert.Equal(t, expected, s2, "%s should equal %s", s2, expected)
	assert.Equal(t, expectedRange, s3, "%s should equal %s", s3, expectedRange)
}

/*
        	Error:      	Not equal:
        	            	expected: "golang"
        	            	actual  : "golang\x00"

There is an extra null character ('\x00') in the resulting string s1. This is likely because the code array has a fixed length of 7 elements, and the 7th element is implicitly initialized to the zero value for the rune type, which is the null character.
*/
