// Deepak Ahuja
// https://levelup.gitconnected.com/get-a-taste-of-concurrency-in-go-625e4301810f

package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for msg := range c {
		fmt.Println(msg)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
		return
	}
	c <- link + " is up!"
}
