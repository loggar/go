package mydebug_test

import (
	"os"
	"strings"
	"testing"
)

func isX509SHA1Enabled() bool {
	value := os.Getenv("GODEBUG")
	return strings.Contains(value, "x509sha1=1")
}

func TestIsX509SHA1Enabled(t *testing.T) {
	t.Run("GODEBUG=x509sha1=1", func(t *testing.T) {
		os.Setenv("GODEBUG", "x509sha1=1")
		defer os.Unsetenv("GODEBUG") // cleanup after the test

		if !isX509SHA1Enabled() {
			t.Fatal("expected x509sha1 to be enabled")
		}
	})

	t.Run("GODEBUG=x509sha1=0", func(t *testing.T) {
		os.Setenv("GODEBUG", "x509sha1=0")
		defer os.Unsetenv("GODEBUG") // cleanup after the test

		if isX509SHA1Enabled() {
			t.Fatal("expected x509sha1 to be disabled")
		}
	})
}
