package main

import (
	"fmt"
	"path/filepath"
)

func absolutePath() {
	abs, err := filepath.Abs("./")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
}
