package main

import "fmt"

func main_3() {
	var creature = "shark"
	var pointer = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	fmt.Println("*pointer =", *pointer)

	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)

	fmt.Println("creature =", creature)
}
