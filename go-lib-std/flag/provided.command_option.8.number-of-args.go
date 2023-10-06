package main

import (
	"flag"
	"fmt"
)

func main_8() {
	var port int
	var dir string

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&dir, "dir", "output_dir", "Directory")

	flag.Parse()

	fmt.Println(port)
	fmt.Println(dir)
	fmt.Println(flag.NArg())
}

/*
$ go run flag.go 1234
8000
output_dir
1

$ go run flag.go -p 8080 -dir dumps hello 1234
8080
dumps
2

$ go run flag.go -p 8080 hello 1234 false
8080
dumps
3
*/
