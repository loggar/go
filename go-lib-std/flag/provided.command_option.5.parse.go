package main

import (
	"flag"
	"fmt"
)

func main_5() {
	var port int
	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.Parse()
	fmt.Println(port)
	vals := flag.Args()
	fmt.Println(vals)
}

/*
$ go run flag.go -p 8123
8123
[]

$ go run flag.go -p 8123 1234 hello true
8123
[1234 hello true]

$ go run flag.go -p 8123 1234 hello true -p 9823 world
8123
[1234 hello true -p 9823 world]
*/
