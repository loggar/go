package main

var a string

func f() {
	print(a)
}

func main() {
	a = "Hello, world"
	go f()
}
