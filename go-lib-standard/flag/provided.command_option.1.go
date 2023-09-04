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
$ go run ./go-packages/path/provided.command_option.go --my-path $HOME/ws-loggar/go/src/github.com/loggar/go
$ go run ./go-packages/path/provided.command_option.go --my-path `pwd`
*/
