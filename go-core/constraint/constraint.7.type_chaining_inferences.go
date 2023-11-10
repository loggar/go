package constraint

import "fmt"

// ToChan: Example of Type Chaining
// Constraint type chaining and type inferences
func ToChan[U ~[]T, T any](t U) <-chan T {
	c := make(chan T)

	// ...

	return c
}

func Usage() {
	c := ToChan([]int{1, 2, 3})
	fmt.Printf("c: %v\n", c)
}
