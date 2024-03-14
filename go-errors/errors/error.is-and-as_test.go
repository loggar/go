package main

import (
	"errors"
	"fmt"
	"testing"
)

// CustomError defines a type for demonstration purposes.
type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

// triggerError returns an error wrapped with additional context.
func triggerError() error {
	baseErr := &CustomError{"base error"}
	// Wrap the base error with context using github.com/pkg/errors or fmt.Errorf in newer Go versions.
	wrappedErr := fmt.Errorf("wrapped error: %w", baseErr)
	return wrappedErr
}

func TestErrorsIs(t *testing.T) {
	err := triggerError()

	// Check if the error is exactly a CustomError using errors.Is
	targetErr := &CustomError{"base error"}

	fmt.Println(err)
	fmt.Println(targetErr)

	if errors.Is(err, targetErr) {
		t.Errorf("Expected the error not to be targetError, but it was")
	}
}

func TestErrorsAs(t *testing.T) {
	err := triggerError()

	// Attempt to find a CustomError in the error chain using errors.As
	var targetErr *CustomError
	if !errors.As(err, &targetErr) {
		t.Errorf("Expected to find a CustomError in the error chain, but none was found")
	} else if targetErr.Error() != "base error" {
		t.Errorf("Expected the found error to be 'base error', got '%s'", targetErr.Error())
	}
}
