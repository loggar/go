package main

import "fmt"

var g = "global" // package level variables

func printLocal() {
	l := "local"
	fmt.Println(l)
}

func globalVars() {
	printLocal()
	fmt.Println(g)
}
