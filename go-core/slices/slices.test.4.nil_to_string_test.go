package slice1

import (
	"strings"
	"testing"
)

func GetLabelIds(i int) []string {
	if i >= 0 {
		return []string{}
	}

	return nil
}

func TestStringSlicesToString(t *testing.T) {
	slice1 := GetLabelIds(1)

	resultJoin := strings.Join(slice1, ", ")

	if resultJoin != "" {
		t.Errorf("Expected empty string, but got %s", resultJoin)
	}
}
