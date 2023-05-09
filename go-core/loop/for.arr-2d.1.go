package main

import "fmt"

func for_2d_1() {
	ints := [][]int{
		{0, 1, 2},
		{-1, -2, -3},
		{9, 8, 7},
	}

	for _, i := range ints {
		fmt.Println(i)
	}
}
