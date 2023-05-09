package main

import "fmt"

func main_4() {
	defer func() {
		fmt.Println("hello from the deferred function!")
	}()

	panic("oh no!")
}
