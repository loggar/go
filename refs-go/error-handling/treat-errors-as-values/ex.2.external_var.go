package main

import (
	"fmt"
	"os"
)

var count int

func write2(s string) {
	fmt.Printf("%s\n", s)
	count += len(s)
	if count >= 32 {
		fmt.Printf("Written length : %d\n", count)
		os.Exit(0)
	}
}

func external_var() {
	write2("sample text 1")
	write2("sample text 2")
	write2("sample text 3")
	write2("sample text 4")
}
