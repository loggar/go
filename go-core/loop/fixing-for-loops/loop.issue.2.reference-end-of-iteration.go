package main

import "fmt"

func main_2() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}

/*
$ go run go-dev/fixing-for-loops/loop.issue.2.reference-end-of-iteration.go
4
4
4
*/
