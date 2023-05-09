package main

import "fmt"

func for_string_range() {
	sammy := "Sammy"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}
}
