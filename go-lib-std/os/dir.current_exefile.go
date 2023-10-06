package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func currentExecutableFile() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}
