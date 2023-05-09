package json_sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Array1Element struct {
	NewPath string `json:"newPath"`
	Path    string `json:"path"`
}

type JSONData1 struct {
	Array1 []Array1Element `json:"array1"`
}

func ReadArray(filepath string) ([]Array1Element, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var jsonData JSONData1
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	array1 := jsonData.Array1

	array1JSON, err := json.Marshal(array1)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	fmt.Println("JSON1 Array:\n", string(array1JSON))

	return array1, nil
}
