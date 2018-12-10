package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetch() {
	url := "https://loggar.github.io/note/sample-res/sample.astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	/*
		There's one more important line I wanted to highlight. I've set a User-Agent in the HTTP request's header.
		This lets remote servers understand what kind of traffic it is receiving. Some sites will even reject empty or generic User-Agent strings.
	*/
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	/*
		It may not always make sense to attempt to parse the HTTP response, one case is if the HTTP code returned is a non-200 number.
		Here's how you can access the HTTP code from the response:
	*/
	fmt.Printf("HTTP: %s\n", res.Status)
	fmt.Println(people1.Number)
}
