package main

import "fmt"

func main() {
	ints := [][]int{
		[]int{0, 1, 2},
		[]int{-1, -2, -3},
		[]int{9, 8, 7},
	}

	for _, i := range ints {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
