package mydebug_test

import (
	"os"
	"testing"
)

func GetEnvWithFallback(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func TestGetEnvWithFallback(t *testing.T) {
	// Backup the original value (if it's set) and restore it after the test.
	originalValue, isSet := os.LookupEnv("MY_TEST_ENV")

	t.Run("Environment variable set", func(t *testing.T) {
		os.Setenv("MY_TEST_ENV", "TestValue")
		defer os.Unsetenv("MY_TEST_ENV")

		got := GetEnvWithFallback("MY_TEST_ENV", "FallbackValue")
		if got != "TestValue" {
			t.Errorf("expected %q, got %q", "TestValue", got)
		}
	})

	t.Run("Environment variable not set", func(t *testing.T) {
		got := GetEnvWithFallback("MY_TEST_ENV", "FallbackValue")
		if got != "FallbackValue" {
			t.Errorf("expected %q, got %q", "FallbackValue", got)
		}
	})

	// Restore the original environment variable value (if it was set).
	if isSet {
		os.Setenv("MY_TEST_ENV", originalValue)
	} else {
		os.Unsetenv("MY_TEST_ENV")
	}
}
