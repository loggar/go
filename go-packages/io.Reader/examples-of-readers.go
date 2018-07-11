package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	example1()
}

/*
There are many kinds of Reader types available in the standard library, and you’ve almost certainly used at least one of them.

If you open a file for reading, the object returned is an `os.File`, which is a Reader (it implements the Read method):
*/

func example1() {
	var r io.Reader
	var err error
	r, err = os.Open("file.txt")

	fmt.Println(r)
	fmt.Println(err)
}

/*
You can also make a Reader from a normal string using `strings.NewReader`:
*/

func example2() {
	var r io.Reader
	r = strings.NewReader("Read will return these bytes")

	fmt.Println(r)
}

/*
The body data from an `http.Request` is a Reader:
*/

func example3(request http.Request) {
	var r io.Reader
	r = request.Body

	fmt.Println(r)
}

/*
A bytes.Buffer is a Reader:
*/

func example4(request http.Request) {
	var r io.Reader
	var buf bytes.Buffer
	r = &buf

	fmt.Println(r)
}
