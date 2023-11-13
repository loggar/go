package main

import "fmt"

func main_1() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for range values {
		<-done
	}
}

/*
$ go run go-dev/fixing-for-loops/loop.issue.1.reference-end-of-iteration.go
c
c
c
*/
