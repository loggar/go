package main

import "sync"

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock()
}

func main() {
	l.Lock()
	go f()
	l.Lock()
	print(a)
}

/*
is guaranteed to print "hello, world". The first call to l.Unlock() (in f) happens before the second call to l.Lock() (in main) returns, which happens before the print.
*/
