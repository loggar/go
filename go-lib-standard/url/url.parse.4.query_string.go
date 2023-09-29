package main

import (
	"fmt"
	"net/url"
)

func main_4() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello&q=world&lang=en"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())
	for k, v := range parsedUrl.Query() {
		fmt.Println("KEY:", k)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}

	queryParams := parsedUrl.Query()
	queryParams.Add("name", "ferris")

	q := queryParams.Encode()
	fmt.Println(q)
}
