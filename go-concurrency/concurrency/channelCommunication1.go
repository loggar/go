package main

var (
	c = make(chan int, 10)
	a string
)

func f() {
	a = "hello, world"
	c <- 0
}

func main_0() {
	go f()
	<-c
	print(a)
}
