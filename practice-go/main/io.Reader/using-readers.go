package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

/*
Now we have a few Reader types — let’s explore ways in which they can be used.

You can read from them directly (this turns out to be the least useful use case):
*/

func fn1(r io.Reader) {
	p := make([]byte, 256)
	n, err := r.Read(p)

	fmt.Println(n)
	fmt.Println(err)
}

/*
`ioutil.ReadAll` lets you read everything from a Reader, and get the raw []byte data:
*/

func fn2(r io.Reader) {
	b, err := ioutil.ReadAll(r)

	fmt.Println(b)
	fmt.Println(err)
}

/*
`io.Copy` lets you read ALL bytes from an io.Reader, and write it to an io.Writer:
*/

func fn3(r io.Reader, w io.Writer) {
	n, err := io.Copy(w, r)

	fmt.Println(n)
	fmt.Println(err)
}

/*
The JSON decoder lets you decode directly from a Reader:
*/

func fn4(r io.Reader, v interface{}) {
	err := json.NewDecoder(r).Decode(v)

	fmt.Println(err)
}

/*
If you’re reading bytes that have been gzipped, you can wrap the io.Reader in a gzip.Reader:
*/

func fn5(r io.Reader) {
	r, err := gzip.NewReader(r)

	fmt.Println(r)
	fmt.Println(err)
}
