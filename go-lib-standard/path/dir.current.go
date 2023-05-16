package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func currentDir() {
	fmt.Println(filepath.Abs("./"))
}

func currentWorkDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}
