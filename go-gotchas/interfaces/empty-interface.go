package main

import (
	"fmt"
)

func foo1() interface{} {
	foo := int64(42)
	return foo
}

func foo2() []interface{} {
	foo := []int{1, 2, 3}
	// return foo
	// [go] cannot use foo (type []int) as type []interface {} in return argument

	r := make([]interface{}, len(foo))
	for i := range foo {
		r[i] = foo[i]
	}
	return r
}

func main() {
	i := foo1()
	fmt.Println(i) // 42

	sliceI := foo2()
	fmt.Println(sliceI) // [1 2 3]
}
