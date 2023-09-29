package main

import (
	"fmt"
	"net/url"
)

func main_5() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())

	queryParams := parsedUrl.Query()

	fmt.Println(queryParams.Get("q"))

	fmt.Println(queryParams.Has("q"))

	if queryParams.Has("lang") {
		fmt.Println(queryParams.Get("lang"))
	}

	queryParams.Add("q", "ferris")
	fmt.Println(queryParams)

	queryParams.Set("q", "books")
	fmt.Println(queryParams)

	queryParams.Del("q")
	fmt.Println(queryParams)
}
