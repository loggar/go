package main

import (
	"flag"
	"fmt"
)

func main_7() {
	var port int
	var dir string

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&dir, "dir", "output_dir", "Directory")

	flag.Parse()

	fmt.Println(port)
	fmt.Println(dir)
	fmt.Println(flag.NFlag())
}

/*
$ go run flag.go
8000
output_dir
0

$ go run flag.go -p 8080 8999 false hello
8080
output_dir
1

$ go run flag.go -p 8080 -dir dumps hello 1234
8080
dumps
2
*/
