package main

import (
	"fmt"
	"log"
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

func userHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(dir)
	}

	return dir
}
