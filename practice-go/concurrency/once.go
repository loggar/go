package main

import "sync"

var a2 string
var once sync.Once

func setup() {
	a2 = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(a2)
}

func onceRun() {
	go doprint()
	go doprint()
}
