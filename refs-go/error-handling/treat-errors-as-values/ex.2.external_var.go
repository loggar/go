package main

import (
	"fmt"
	"os"
)

var count int

func write(s string) {
	fmt.Printf("%s\n", s)
	count += len(s)
	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		os.Exit(0)
	}
}

func main() {
	write("sample text 1")
	write("sample text 2")
	write("sample text 3")
	write("sample text 4")
}
