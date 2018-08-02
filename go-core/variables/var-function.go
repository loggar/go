package main

import (
	"fmt"
)

func varFn() {
	action := func() {
		fmt.Println("anonymous fn")
	}

	action()
}
