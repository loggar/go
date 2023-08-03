package main

import "fmt"

func call_anonymous_fn() {
	draw := func() {
		fmt.Println("Drawing")
	}
	draw()
	draw()

	draw2 := func() int {
		fmt.Println("Drawing")
		return 1
	}()

	fmt.Println(draw2)
}
