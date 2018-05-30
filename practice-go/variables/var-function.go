package main

import (
	"fmt"
)

func main() {
	action := func() {
		fmt.Println("anonymous fn")
	}

	action()
}
