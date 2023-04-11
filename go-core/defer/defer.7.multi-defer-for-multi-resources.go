package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//lint:ignore U1000 // ignore unused variable warning//lint:ignore U1000 // ignore unused variable warning
func main_7() {
	if err := write2("sample.txt", "This file contains some sample text."); err != nil {
		log.Fatal("failed to create file")
	}

	if err := fileCopy("sample.txt", "sample-copy.txt"); err != nil {
		log.Fatal("failed to copy file")
	}
}

func write2(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	return file.Close()
}

func fileCopy(source string, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	fmt.Printf("Copied %d bytes from %s to %s\n", n, source, destination)

	if err := src.Close(); err != nil {
		return err
	}

	return dst.Close()
}
