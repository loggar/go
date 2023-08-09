package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func checkIfPathExists() {
	file_name := "drafts/default.md"
	path, err := filepath.Abs(file_name)
	if err != nil {
		log.Println(err)
	} else {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Println("No, " + path + " does not exists")
		} else {
			log.Println("Yes, " + path + " exists")
		}
	}
}

func checkIfPathIsDir() {
	file_name := "drafts/default.md"
	path, err := filepath.Abs(file_name)
	if err != nil {
		log.Println(err)
	} else {
		if t, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatal("No, " + path + " does not exists")
		} else {
			log.Println(path)
			log.Println(t.IsDir())
		}
	}
}

func fileExtensionFromPath() {
	file_name := "default.md"
	dir, err := filepath.Abs(file_name)
	if err != nil {
		log.Println(err)
	} else {
		file_ext := path.Ext(dir)
		log.Println(file_ext)
	}
}

func fileNameFromPath() {
	file_name := "default.md"
	dir, err := filepath.Abs(file_name)
	if err != nil {
		log.Println(err)
	} else {
		file_ext := path.Ext(dir)
		log.Println(file_ext)
		log.Println(strings.TrimSuffix(dir, file_ext))
		log.Println(strings.TrimSuffix(file_name, file_ext))
	}
}

func checkIsAbsolute() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	log.Println(dir)
	log.Println(filepath.IsAbs(dir))

	dir = "../math"
	log.Println(dir)
	log.Println(filepath.IsAbs(dir))
}
