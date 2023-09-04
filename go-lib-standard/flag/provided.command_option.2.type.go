package main

import (
	"flag"
	"fmt"
)

func main_2() {
	var port int
	var dir string
	var publish bool

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&dir, "dir", "output_dir", "Directory")
	flag.BoolVar(&publish, "publish", false, "Publish the article")

	flag.Parse()

	fmt.Println(port)
	fmt.Println(dir)
	fmt.Println(publish)

	if publish {
		fmt.Println("Publishing article...")
	} else {
		fmt.Println("Article saved as Draft!")
	}
}

/*
$ go run provided.command_option.type.go
8000
output_dir
false
Article saved as Draft!

$ go run provided.command_option.type.go -p 1234
1234
output_dir
false
Article saved as Draft!

$ go run provided.command_option.type.go -p 1234 -dir site_out
1234
site_out
false
Article saved as Draft!

$ go run provided.command_option.type.go -publish
8000
output_dir
true
Publishing article...
*/
