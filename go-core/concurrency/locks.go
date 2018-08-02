package main

import "sync"

var l1 sync.Mutex
var a1 string

func fn1() {
	a1 = "hello, world"
	l1.Unlock()
}

func locks() {
	l1.Lock()
	go fn1()
	l1.Lock()
	print(a1)
}

/*
is guaranteed to print "hello, world". The first call to l.Unlock() (in f) happens before the second call to l.Lock() (in main) returns, which happens before the print.
*/
