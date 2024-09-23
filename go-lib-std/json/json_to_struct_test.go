package json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Unmarshalling(jsonString string) error {
	var data MyStruct
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling JSON")
	}

	fmt.Printf("Unmarshalled data: %+v\n", data)
	return nil
}

func TestUnmarshalling(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "Empty JSON string",
			jsonData: "{}",
			wantErr:  false,
		},
		{
			name:     "Invalid JSON string",
			jsonData: "",
			wantErr:  true,
		},
		{
			name:     "Empty string",
			jsonData: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unmarshalling(tt.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshalling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
