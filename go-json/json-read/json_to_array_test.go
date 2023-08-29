package json_sample_test

import (
	"testing"

	jsonsample "github.com/loggar/go/go-json/json-read-write"

	"github.com/stretchr/testify/assert"
)

func TestReadArray(t *testing.T) {
	dataArray, errRead := jsonsample.ReadArray("data.ex.data.json")
	assert.NoError(t, errRead, "ReadArray() should not return an error")

	expectedLength := 8
	assert.Equal(t, expectedLength, len(dataArray), "len(dataArray) should be %d", expectedLength)

	item1 := dataArray[0]
	expectedValue1 := "the-new-value-1"
	assert.Equal(t, expectedValue1, item1.NewPath, "{%v}.NewPath should equal to %s", item1, expectedValue1)
}
