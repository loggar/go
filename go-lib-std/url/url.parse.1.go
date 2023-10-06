package main

import (
	"fmt"
	"net/url"
)

func main() {
	// simple url
	urlString := "http://www.google.com"
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
}
