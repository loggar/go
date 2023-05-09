package json_sample_test

import (
	"testing"

	json_sample "github.com/loggar/go/go-json/json-read-write"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstMatch(t *testing.T) {
	array1, err := json_sample.ReadArray("data.ex.data.json")
	assert.NoError(t, err, "ReadArray() should not return an error")

	v1 := "invalid-value"
	value1 := json_sample.FindFirstPathMatch(array1, v1)
	expectedValue1 := ""
	assert.Equal(t, expectedValue1, value1, "first match for %s should equal to %s", v1, expectedValue1)

	v2 := "/def-156"
	value2 := json_sample.FindFirstPathMatch(array1, v2)
	expectedValue2 := "the-new-value-1"
	assert.Equal(t, expectedValue2, value2, "first match for %s should equal to %s", v2, expectedValue2)

	v3 := "/jrh-629"
	value3 := json_sample.FindFirstPathMatch(array1, v3)
	expectedValue3 := "the-new-value-6"
	assert.Equal(t, expectedValue3, value3, "first match for %s should equal to %s", v3, expectedValue3)
}
