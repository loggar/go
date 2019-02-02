package basics

import (
	"encoding/hex"
	"testing"
)

func TestError(t *testing.T) {
	dest := make([]byte, 0)
	if _, err := hex.Decode(dest, []byte{8}); err != nil {
		t.Error(err)
	}
}

/*
$ go test -v -run Error ./testing-in-go/

=== RUN   TestError
--- FAIL: TestError (0.00s)
        testing.3.basics_test.go:11: encoding/hex: odd length hex string
FAIL
exit status 1
FAIL    github.com/loggar/go/go-testing/testing-in-go   0.068s
*/
