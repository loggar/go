package main

import (
	"log"
	"os"
	"path/filepath"
)

func getRelativePathFromBaseToTargetPath() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err = filepath.Abs("plugins/")
	s, err := filepath.Abs("static/projects/")
	if err != nil {
		log.Println(err)
	}

	log.Println(s)
	log.Println(dir)
	log.Println(filepath.Rel(s, dir))
}
