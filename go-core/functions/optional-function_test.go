package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SomeVerifier struct {
	CurrentTime time.Time
	Name        string
}

type SomeVerifierOptions struct {
	TimeOverride *time.Time
}

type SomeVerifierOptionSetter func(*SomeVerifierOptions)

func WithTimeOverride(t time.Time) SomeVerifierOptionSetter {
	return func(opt *SomeVerifierOptions) {
		opt.TimeOverride = &t
	}
}

func NewSomeVerifier(name string, setters ...SomeVerifierOptionSetter) (*SomeVerifier, error) {
	opts := &SomeVerifierOptions{}
	for _, setter := range setters {
		setter(opts)
	}

	var c time.Time
	if opts.TimeOverride != nil {
		c = *opts.TimeOverride
	} else {
		c = time.Now()
	}

	return &SomeVerifier{
		CurrentTime: c,
		Name:        name,
	}, nil
}

func WithMockTime(t time.Time) SomeVerifierOptionSetter {
	return func(opt *SomeVerifierOptions) {
		opt.TimeOverride = &t
	}
}

func TestNewSomeVerifier(t *testing.T) {
	mockCurrentTime := time.Date(2023, 10, 18, 5, 0, 0, 0, time.UTC)

	t.Run("SomeVerifier with mock time", func(t *testing.T) {
		sv, err := NewSomeVerifier("Mock", WithMockTime(mockCurrentTime))
		assert.NoError(t, err, "should not return an error")
		assert.Equal(t, mockCurrentTime, sv.CurrentTime, "should use mock time")
	})

	t.Run("SomeVerifier without mock time", func(t *testing.T) {
		sv, err := NewSomeVerifier("Mock")
		assert.NoError(t, err, "should not return an error")
		assert.NotEqual(t, mockCurrentTime, sv.CurrentTime, "should use actual current time")
	})
}
