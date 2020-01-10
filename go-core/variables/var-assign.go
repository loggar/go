package main

import "fmt"

func varDeclare() {
	s := "Hello, World!"
	f := 45.06
	b := 5 > 9 // A Boolean value will return either true or false
	array := [4]string{"item_1", "item_2", "item_3", "item_4"}
	slice := []string{"one", "two", "three"}
	m := map[string]string{"letter": "g", "number": "seven", "symbol": "&"}

	fmt.Println(s)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(m)
}
