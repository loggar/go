package main

import "fmt"

func varReassign() {
	i := 76
	fmt.Println(i)

	i = 42
	fmt.Println(i)

	// i = "Sammy"
	// cannot use "Sammy" (type string) as type int in assignment

	// i := 10
	// no new variables on left side of :=
}
