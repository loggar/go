package go_embed_test

import (
	_ "embed"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var error1 = errors.New("ERROR CODE 0")

func aFunc(i int32) (int32, error) {
	if i < 1 {
		return 0, nil
	}
	return 1, fmt.Errorf("wrapped error: %w", error1)
}

func TestErrors(t *testing.T) {
	err1 := errors.New("ERROR CODE 0")
	assert.EqualError(t, err1, error1.Error(), "Equals error")
	assert.Equal(t, err1.Error(), error1.Error(), "assert error")

	_, err2 := aFunc(1)
	assert.True(t, errors.Is(err2, error1), "errors Is")

	_, err3 := aFunc(0)
	assert.NoError(t, err3, "assert no error")
	assert.True(t, err3 == nil, "assert nil error")
}
