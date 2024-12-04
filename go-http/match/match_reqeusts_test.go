package helper

import (
	"net/http"
	"strings"
	"testing"
)

func TestCompareRequests(t *testing.T) {
	tests := []struct {
		name     string
		req1     *http.Request
		req2     *http.Request
		expected bool
	}{
		{
			name: "Equal requests",
			req1: func() *http.Request {
				req, _ := http.NewRequest("GET", "http://example.com", strings.NewReader("body1"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			req2: func() *http.Request {
				req, _ := http.NewRequest("GET", "http://example.com", strings.NewReader("body1"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			expected: true,
		},
		{
			name: "Different methods",
			req1: func() *http.Request {
				req, _ := http.NewRequest("GET", "http://example.com", strings.NewReader("body1"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			req2: func() *http.Request {
				req, _ := http.NewRequest("POST", "http://example.com", strings.NewReader("body1"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			expected: false,
		},
		{
			name: "Different bodies",
			req1: func() *http.Request {
				req, _ := http.NewRequest("GET", "http://example.com", strings.NewReader("body1"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			req2: func() *http.Request {
				req, _ := http.NewRequest("GET", "http://example.com", strings.NewReader("body2"))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CompareRequests(tt.req1, tt.req2); got != tt.expected {
				t.Errorf("CompareRequests() = %v, want %v", got, tt.expected)
			}
		})
	}
}
