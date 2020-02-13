package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := write("readme.txt", "This is a readme file"); err != nil {
		log.Fatal("failed to write file:", err)
	}
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}
	return nil
}
