package main

import (
	"fmt"
)

/*
The return or result "parameters" of a Go function can be given names and used as regular variables, just like the incoming parameters.
When named, they are initialized to the zero values for their types when the function begins;
if the function executes a return statement with no arguments, the current values of the result parameters are used as the returned values.

The names are not mandatory but they can make code shorter and clearer:
they're documentation. If we name the results of nextInt it becomes obvious which returned int is which.
*/

func nextDigit(s string, pos int) (value, nextPos int) {
	return
}

func main() {
	var v, p int = nextDigit("string", 0)
	fmt.Printf("%d, %d\n", v, p)
}
