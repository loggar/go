package main

import "fmt"

func for_2d_2() {
	ints := [][]int{
		{0, 1, 2},
		{-1, -2, -3},
		{9, 8, 7},
	}

	for _, i := range ints {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
