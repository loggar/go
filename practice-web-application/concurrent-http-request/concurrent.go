package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*
MakeRequest sends sample http-request
*/
func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go MakeRequest(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// λ go run concurrent.go http://localhost:8080/param1 http://localhost:8080/param2 http://localhost:8080/param3
// 0.01 elapsed with response length: 42 http://localhost:8080/param3
// 0.01 elapsed with response length: 42 http://localhost:8080/param1
// 0.01 elapsed with response length: 42 http://localhost:8080/param2
// 0.01s elapsed
