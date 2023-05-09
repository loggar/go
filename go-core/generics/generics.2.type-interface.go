package main

import "fmt"

func main_2() {
	fmt.Printf("result %v\n", min2(1, 2))
	fmt.Printf("result %v\n", min2(0.2, 0.3))
	fmt.Printf("result %v\n", min2(1, 2))
	fmt.Printf("result %v\n", min2(1, 2))
	fmt.Printf("result %v\n", min2(1, 2))
	fmt.Printf("result %v\n", min2(1, 2))
}

type minTypes interface {
	~float64 | int
}

func min2[T minTypes](a T, b T) T {
	if a < b {
		return a
	}

	return b
}
