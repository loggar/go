package json_sample_test

import (
	"testing"

	jsonsample "github.com/loggar/go/go-json/json-read-write"

	"github.com/stretchr/testify/assert"
)

func TestReadStruct1(t *testing.T) {
	data1, errReadStruct := jsonsample.ReadStruct1("data.ex.target.json")
	assert.NoError(t, errReadStruct, "Read1() should not return an error")
	assert.NotNil(t, data1, "data should not be nil")

	topArray := data1.TopArrayElement
	expectedLength := 3
	assert.Equal(t, expectedLength, len(topArray), "len(topArray) should be %d", expectedLength)

	subArray := data1.TopArrayElement[1].SubArr
	expectedSubLen := 3
	assert.Equal(t, expectedSubLen, len(subArray), "len(subArray) should be %d", expectedSubLen)

	newValue1 := "updated-value-1"
	topArray[0].NewPath = newValue1
	assert.Equal(t, newValue1, data1.TopArrayElement[0].NewPath, "data should now be updated")

	newValue2 := "updated-value-2"
	subArray[0].NewPath = newValue2
	assert.Equal(t, newValue2, data1.TopArrayElement[1].SubArr[0].NewPath, "data should now be updated")
}
