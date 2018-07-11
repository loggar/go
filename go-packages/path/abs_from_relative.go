package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abs, err := filepath.Abs("./")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
}
