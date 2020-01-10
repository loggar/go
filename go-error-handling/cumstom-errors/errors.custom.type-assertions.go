package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// RequestError2 ...
type RequestError2 struct {
	StatusCode int

	Err error
}

func (r *RequestError2) Error() string {
	return r.Err.Error()
}

// Temporary does ...
func (r *RequestError2) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func doRequest2() error {
	return &RequestError2{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func customError2() {
	err := doRequest2()
	if err != nil {
		fmt.Println(err)
		re, ok := err.(*RequestError2)
		if ok {
			if re.Temporary() {
				fmt.Println("This request can be tried again")
			} else {
				fmt.Println("This request cannot be tried again")
			}
		}
		os.Exit(1)
	}

	fmt.Println("success!")
}
