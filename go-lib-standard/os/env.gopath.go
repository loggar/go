package main

import (
	"fmt"
	"go/build"
	"os"
)

func goPath() {
	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fmt.Println(gopath)
}
