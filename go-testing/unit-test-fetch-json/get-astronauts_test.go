package main

import "testing"

type testWebRequest struct {
}

func (testWebRequest) FetchBytes(url string) []byte {
	return []byte(`{"number": 2}`)
}

func TestGetAstronauts(t *testing.T) {
	amount := GetAstronauts(testWebRequest{})
	expect := 1
	if amount != expect {
		t.Errorf("People in space, got: %d, want: %d.", amount, expect)
	}
}
