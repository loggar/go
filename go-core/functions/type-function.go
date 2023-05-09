package main

import "fmt"

type calculatorNum func(int, int) int

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func fn3() {
	multi := func(i int, j int) int {
		return i * j
	}

	r1 := calNum(multi, 10, 20)
	fmt.Println(r1)
}
