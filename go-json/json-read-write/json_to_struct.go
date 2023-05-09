package json_sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SubArrElement struct {
	NewPath string `json:"newPath"`
	Path    string `json:"path"`
}

type TopArrayElement struct {
	NewPath string          `json:"newPath"`
	Path    string          `json:"path"`
	SubArr  []SubArrElement `json:"subArr"`
}

type JSONData struct {
	TopKey          map[string]string `json:"topKey"`
	TopArrayElement []TopArrayElement `json:"topArrayElement"`
}

func ReadStruct1(filepath string) (*JSONData, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var jsonData JSONData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	fmt.Println("JSON2:\n", jsonData)

	return &jsonData, nil
}
