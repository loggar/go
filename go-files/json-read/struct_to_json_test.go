package json_sample_test

import (
	"fmt"
	"testing"

	jsonsample "github.com/loggar/go/go-files/json-read"

	"github.com/stretchr/testify/assert"
)

func updateData(data *jsonsample.JSONData) {
	topArray := data.TopArrayElement
	subArray := data.TopArrayElement[1].SubArr

	topArray[0].NewPath = "updated-value-1"
	subArray[0].NewPath = "updated-value-2"
}

func TestConvertDataToJsonString(t *testing.T) {
	data1, errReadStruct := jsonsample.ReadStruct1("data.ex.target.json")
	assert.NoError(t, errReadStruct, "ReadArray() should not return an error")
	assert.NotNil(t, data1, "data should not be nil")

	updateData(data1)

	jsonStr, errConvert := jsonsample.ConvertDataToJsonString(data1)
	assert.NoError(t, errConvert, "ConvertDataToJsonString() should not return an error")
	fmt.Println("JSON2_NEW:\n", jsonStr)
}

func TestWriteDataToJsonFile(t *testing.T) {
	data1, errReadStruct := jsonsample.ReadStruct1("data.ex.target.json")
	assert.NoError(t, errReadStruct, "ReadArray() should not return an error")
	assert.NotNil(t, data1, "data should not be nil")

	updateData(data1)

	outFilePath := "data.out.test.json"
	errWrite := jsonsample.WriteDataToJsonFile(data1, outFilePath)
	assert.NoError(t, errWrite, "WriteDataToJsonFile() should not return an error")

	data2, errReadStruct := jsonsample.ReadStruct1(outFilePath)
	assert.NoError(t, errReadStruct, "ReadArray() should not return an error")
	assert.Equal(t, data1, data2, "Error")
}
