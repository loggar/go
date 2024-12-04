package helper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func printRequestDetails(req *http.Request) {
	// Print the request method, URL, and protocol
	fmt.Printf("Method: %s\n", req.Method)
	fmt.Printf("URL: %s\n", req.URL.String())
	fmt.Printf("Protocol: %s\n", req.Proto)

	// Print the request headers
	fmt.Println("Headers:")
	for name, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}

	// Print the request body
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v\n", err)
		} else {
			fmt.Printf("Body: %s\n", string(bodyBytes))
			// Restore the body for further use
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
	}

	// Print the request dump
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Printf("Error dumping request: %v\n", err)
	} else {
		fmt.Printf("Full Request Dump:\n%s\n", string(dump))
	}
}
