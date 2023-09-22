package main

var str0 string

func fn0() {
	print("fn0: ", str0, "\n")
}

func main() {
	str0 = "Hello, world"
	print("main: ", str0, "\n")
	go fn0()
}
