package regex_test

import (
	"regexp"
	"testing"
	"unicode"
)

// Using regexp
func validateWithRegex(s string) bool {
	pattern := regexp.MustCompile(`^[\p{L}]+$`)
	return pattern.MatchString(s)
}

// Using unicode package (more efficient)
func validateWithUnicode(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

func TestValidateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "english only",
			input:    "Hello",
			expected: true,
		},
		{
			name:     "korean characters",
			input:    "안녕하세요",
			expected: true,
		},
		{
			name:     "chinese characters",
			input:    "你好",
			expected: true,
		},
		{
			name:     "japanese characters",
			input:    "こんにちは",
			expected: true,
		},
		{
			name:     "mixed letters",
			input:    "안녕하세요Helloこんにちは",
			expected: true,
		},
		{
			name:     "with numbers",
			input:    "Hello123",
			expected: false,
		},
		{
			name:     "with special chars",
			input:    "Hello!",
			expected: false,
		},
		{
			name:     "with space",
			input:    "Hello World",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "arabic characters",
			input:    "مرحبا",
			expected: true,
		},
		{
			name:     "cyrillic characters",
			input:    "привет",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test regex version
			result := validateWithRegex(tt.input)
			if result != tt.expected {
				t.Errorf("validateWithRegex(%q) = %v, want %v", tt.input, result, tt.expected)
			}

			// Test unicode version
			result = validateWithUnicode(tt.input)
			if result != tt.expected {
				t.Errorf("validateWithUnicode(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
