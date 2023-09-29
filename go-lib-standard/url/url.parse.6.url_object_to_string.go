package main

import (
	"fmt"
	"net/url"
)

func main_6() {
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	fmt.Println("URL:", urlStr)
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	queryParams := parsedUrl.Query()
	queryParams.Add("name", "ferris")

	q := queryParams.Encode()
	fmt.Println(q)
	parsedUrl.RawQuery = q
	newUrl := parsedUrl.String()
	fmt.Println("New URL:", newUrl)
}
