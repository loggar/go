package json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Unmarshalling(jsonString string) (*MyStruct, error) {
	var data MyStruct
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling JSON")
	}

	fmt.Printf("Unmarshalled data: %+v\n", data)
	return &data, nil
}

func TestUnmarshalling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantData *MyStruct
		wantErr  bool
	}{
		{
			name:     "Empty JSON string",
			jsonData: "{}",
			wantData: &MyStruct{},
			wantErr:  false,
		},
		{
			name:     "Invalid JSON string",
			jsonData: "",
			wantData: nil,
			wantErr:  true,
		},
		{
			name:     "Empty string",
			jsonData: "",
			wantData: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Unmarshalling(tt.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshalling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, data, tt.wantData, "Unmarshalled data")
		})
	}
}
