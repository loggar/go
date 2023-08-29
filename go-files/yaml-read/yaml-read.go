package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// go get gopkg.in/yaml.v3
// go run go-files/yaml-read/yaml-read.go

func main() {
	f, err := os.ReadFile("_in/go-files/yaml/data.ex.1.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}
}
