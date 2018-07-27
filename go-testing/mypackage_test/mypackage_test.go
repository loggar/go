package mypackage_test

import (
	"testing"

	"github.com/loggar/go/go-testing/mypackage"
)

/*
Put your tests in a different package

Go insists that files in the same folder belong to the same package, that is except for `_test.go` files. Moving your test code out of the package allows you to write tests as though you were a real user of the package. You cannot fiddle around with the internals, instead you focus on the exposed interface and are always thinking about any noise that you might be adding to your API.

This frees you up to change the internals however you like without having to adjust your test code.
*/

func TestCount(t *testing.T) {
	if mypackage.Count() != 1 {
		t.Error("Expected 1")
	}
	if mypackage.Count() != 2 {
		t.Error("Expected 2")
	}
	if mypackage.Count() != 3 {
		t.Error("Expected 3")
	}
}
