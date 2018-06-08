package main

import (
	"flag"
	"fmt"
)

func main() {
	var myPath string
	flag.StringVar(&myPath, "my-path", "/", "Provide project path as an absolute path")
	flag.Parse()

	fmt.Printf("provided path was %s\n", myPath)
}

/*
$ cd /Users/Arman/go/src/github.com/rmaan/project/
$ go run main.go  --my-path /Users/Arman/go/src/github.com/rmaan/project/
provided path was /Users/Arman/go/src/github.com/rmaan/project/
$ # or more easily
$ go run main.go  --my-path `pwd`
provided path was /Users/Arman/go/src/github.com/rmaan/project/
*/
