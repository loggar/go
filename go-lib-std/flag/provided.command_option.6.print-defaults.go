package main

import (
	"flag"
	"fmt"
)

func main_6() {
	var port int
	var help bool
	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		fmt.Println(port)
		vals := flag.Args()
		fmt.Println(vals)
	}
}

/*
$ go run help.go -h
Usage of /tmp/go-build121267600/b001/exe/help:
  -help
        Help
  -p int
        Provide a port number (default 8000)

$ go run help.go
8000
[]
*/
