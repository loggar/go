package main

import (
	"fmt"
	"net/url"
)

func mai_7() {
	// url with fragments
	urlStr := "https://pkg.go.dev/math/rand#Norm Float64"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Fragment)
	fmt.Println(parsedUrl.RawFragment)
	fmt.Println(parsedUrl.EscapedFragment())
}
