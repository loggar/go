package customerrors

import (
	"errors"
	"fmt"
)

// WrappedError ...
type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

// Wrap does..
func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func customErrors3() {
	err := errors.New("boom")
	err = Wrap(err, "main")

	fmt.Println(err)
}
