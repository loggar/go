package json_sample_test

import (
	"testing"

	json_sample "github.com/loggar/go/go-json/json-read-write"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstMatch(t *testing.T) {
	array1, errRead1 := json_sample.ReadArray("data.ex.data.json")
	assert.NoError(t, errRead1, "Read1() should not return an error")

	value1 := json_sample.FindFirstPathMatch(array1, "invalid-value")
	expectedValue1 := ""
	assert.Equal(t, expectedValue1, value1, "first match for %s should equal to %s", value1, expectedValue1)

	value2 := json_sample.FindFirstPathMatch(array1, "/def-156")
	expectedValue2 := "the-new-value-1"
	assert.Equal(t, expectedValue2, value2, "first match for %s should equal to %s", value2, expectedValue2)

	value3 := json_sample.FindFirstPathMatch(array1, "/jrh-629")
	expectedValue3 := "the-new-value-6"
	assert.Equal(t, expectedValue3, value3, "first match for %s should equal to %s", value3, expectedValue3)
}
