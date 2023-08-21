package main

import "fmt"

func main_1() {
	fmt.Printf("result %v\n", min(1, 2))
	fmt.Printf("result %v\n", min(0.2, 0.3))
	fmt.Printf("result %v\n", min(1, 2))
	fmt.Printf("result %v\n", min(1, 2))
	fmt.Printf("result %v\n", min(1, 2))
	fmt.Printf("result %v\n", min(1, 2))
}

func min[T interface {
	~float64 | int
}](a T, b T) T {
	if a < b {
		return a
	}

	return b
}
