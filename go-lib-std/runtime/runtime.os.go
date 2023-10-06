package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())

	LineDelimiter := "\n"
	if runtime.GOOS == "windows" {
		LineDelimiter = "\r\n"
	}

	fmt.Println(LineDelimiter)
}
