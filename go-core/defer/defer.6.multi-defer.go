package main

import "fmt"

//nolint:go-staticcheck
func main_6() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
