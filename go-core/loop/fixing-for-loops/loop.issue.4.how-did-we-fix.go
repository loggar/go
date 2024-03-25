package main

import "fmt"

func main_4() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		i := i
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}

/*
$ go run go-dev/fixing-for-loops/loop.issue.3.how-did-we-fix.go
1
2
3
*/
