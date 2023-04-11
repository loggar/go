package main

import (
	"io"
	"log"
	"os"
)

//nolint:go-staticcheck
func main_5() {
	if err := write1("readme.txt", "This is a readme file"); err != nil {
		log.Fatal("failed to write file:", err)
	}
}

func write1(fileName string, text string) error {
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
