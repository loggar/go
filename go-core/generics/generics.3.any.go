package main

import "fmt"

type SomeStruct struct {
	Name string
	Type string
}

func main_3() {
	func1("abc")
	func1(1)
	func1(1.0)
	func1(SomeStruct{
		Name: "some",
		Type: "struct",
	})

	func1([]int{4, 2})
}

func func1[T any](a T) {
	fmt.Printf("%T %v\n", a, a)
}
