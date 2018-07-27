package mypackage

import (
	"testing"
)

/*
Internal tests go in a different file

If you do need to unit test some internals, create another file with `_internal_test.go` as the suffix. Internal tests will necessarily be more brittle than your interface tests — but they’re a great way to ensure internal components are behaving, and are especially useful if you do test-driven development.
*/

func TestCount(t *testing.T) {
	if Count() != 1 {
		t.Error("expected 1")
	}

	if count != 1 { // test internals
		t.Error("expected 1 for count too")
	}
}
