package main

import (
	"fmt"
)

func main_1() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}

/*
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
	<workspace>/src/github.com/loggar/go/go-error-handling/panic/panic.1.basic.go:13 +0x6c
exit status 2
Process exiting with code: 1
*/
