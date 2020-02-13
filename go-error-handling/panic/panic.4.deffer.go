package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("hello from the deferred function!")
	}()

	panic("oh no!")
}
