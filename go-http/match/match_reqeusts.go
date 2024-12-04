package helper

import (
	"fmt"
	"net/http"
	"reflect"
)

// CompareRequests compares two http.Request objects and prints detailed differences.
// Returns true if they are equal, false otherwise.
func CompareRequests(req1, req2 *http.Request) bool {
	equal := true

	if !reflect.DeepEqual(req1.Method, req2.Method) {
		fmt.Printf("Method differs: %s != %s\n", req1.Method, req2.Method)
		equal = false
	}

	if !reflect.DeepEqual(req1.URL, req2.URL) {
		fmt.Printf("URL differs: %v != %v\n", req1.URL, req2.URL)
		equal = false
	}

	if !reflect.DeepEqual(req1.Header, req2.Header) {
		fmt.Printf("Header differs: %v != %v\n", req1.Header, req2.Header)
		equal = false
	}

	if !reflect.DeepEqual(req1.Body, req2.Body) {
		fmt.Printf("Body differs: %v != %v\n", req1.Body, req2.Body)
		equal = false
	}

	if !reflect.DeepEqual(req1.ContentLength, req2.ContentLength) {
		fmt.Printf("ContentLength differs: %d != %d\n", req1.ContentLength, req2.ContentLength)
		equal = false
	}

	if !reflect.DeepEqual(req1.TransferEncoding, req2.TransferEncoding) {
		fmt.Printf("TransferEncoding differs: %v != %v\n", req1.TransferEncoding, req2.TransferEncoding)
		equal = false
	}

	if !reflect.DeepEqual(req1.Host, req2.Host) {
		fmt.Printf("Host differs: %s != %s\n", req1.Host, req2.Host)
		equal = false
	}

	if !reflect.DeepEqual(req1.Trailer, req2.Trailer) {
		fmt.Printf("Trailer differs: %v != %v\n", req1.Trailer, req2.Trailer)
		equal = false
	}

	if !reflect.DeepEqual(req1.RemoteAddr, req2.RemoteAddr) {
		fmt.Printf("RemoteAddr differs: %s != %s\n", req1.RemoteAddr, req2.RemoteAddr)
		equal = false
	}

	if !reflect.DeepEqual(req1.RequestURI, req2.RequestURI) {
		fmt.Printf("RequestURI differs: %s != %s\n", req1.RequestURI, req2.RequestURI)
		equal = false
	}

	return equal
}
