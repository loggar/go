package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {
	url := "https://github.com/loggar/go/raw/master/_in/go-files/csv/data.ex.1.csv"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	n, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
	}

	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
