package main

import "fmt"

func main() {
	sammy := "Sammy"

	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}
}
