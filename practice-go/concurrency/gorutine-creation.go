package main

var str0 string

func fn0() {
	print(str0)
}

func gorutineCreation() {
	str0 = "Hello, world"
	go fn0()
}
