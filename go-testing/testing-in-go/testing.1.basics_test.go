package basics

import (
	"fmt"
	"testing"
)

func TestLoginTesting(t *testing.T) {
	fmt.Println("1. This will be displayed when running tests")
	t.Log("2. This will be displayed when running tests")
}

func TestTruth(t *testing.T) {
	if true != true {
		t.Error("everything I know is wrong")
	}
}
