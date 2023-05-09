package main

import "fmt"

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func fn1() {
	multi := func(i int, j int) int {
		return i * j
	}
	r1 := calc(multi, 10, 20)
	fmt.Println(r1)
}
