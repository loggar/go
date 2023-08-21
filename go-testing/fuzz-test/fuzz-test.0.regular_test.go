package fuzztestingingo

import (
	"testing"

	bu "github.com/loggar/go/go-lib-src/byteutil"
)

// TestEqual is a simple regular test
func TestEqual(t *testing.T) {
	if !bu.Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Error("expected true, got false")
	}
}
