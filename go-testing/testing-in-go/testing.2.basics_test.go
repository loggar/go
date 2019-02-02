package basics

import "testing"

func TestBroken(t *testing.T) {
	if true != false {
		t.Error("expected", true, "got", false)
	}

	if 1+1 != 4 {
		t.Fatal("Can't proceed!", 1+1, 4)
	}
}

/*
$ go test -v ./testing-in-go/
*/
