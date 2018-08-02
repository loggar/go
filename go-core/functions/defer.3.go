package main

import "fmt"

// defer rules.

/*
1. A deferred function's arguments are evaluated when the defer statement is evaluated.

In this example, the expression "i" is evaluated when the Println call is deferred. The deferred call will print "0" after the function returns.
*/

func a1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

/*
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.

This function prints "3210":
*/

func b1() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

/*
3. Deferred functions may read and assign to the returning function's named return values.

In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:
*/

func c1() (i int) {
	defer func() { i++ }()
	return 1
}

func c1test() {
	fmt.Print(c1())
}
