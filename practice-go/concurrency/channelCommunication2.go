package main

var ch1 = make(chan int)
var str1 string

func f1() {
	str1 = "hello, world"
	<-ch1
}

func channelCommunication2() {
	go f1()
	ch1 <- 0
	print(str1)
}
