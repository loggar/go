package json_sample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ConvertDataToJsonString(data *JSONData) (string, error) {
	marshaledJSON, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Error marshalling JSON: %v", err)
	}
	return string(marshaledJSON), nil
}

func WriteDataToJsonFile(data *JSONData, filepath string) error {
	marshaledJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	err = ioutil.WriteFile(filepath, marshaledJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing JSON to file: %v", err)
	}

	return nil
}
