package constraint

import (
	"errors"
	"fmt"
)

type Data Example

// Sort with comparable
// It was introduced in Go 1.18 as an interface implemented by comparable types such as structs, pointers, interfaces, channels, and so on.
//
// Note: Comparable is not used as a type of any variable.
func Sort[K comparable, T Data](values map[K]T) error {
	for k, t := range values {
		// sorting logic
		return errors.New(fmt.Sprintf("%v %v", k, t))
	}
	return nil
}
