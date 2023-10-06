package main

import (
	"flag"
	"fmt"
)

func main_3() {
	var port int
	var dir string

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&dir, "dir", "output_dir", "Directory")

	flag.Parse()

	fmt.Println(port)
	fmt.Println(dir)
	flag.Set("dir", "dumps")
	fmt.Println(dir)
}

/*
$ go run provided.command_option.set.go -p 8080
8080
output_dir
dumps
*/
