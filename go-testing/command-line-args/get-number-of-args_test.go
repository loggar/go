package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Keep backup of the real stdout.
	old := os.Stdout
	defer func() { os.Stdout = old }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set command line arguments
	os.Args = []string{"cmd", "arg1", "arg2"}

	main()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r) // Note the use of io.Copy here
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	out := <-outC

	expected := "Total Args: 2\n" // Corrected the expected output
	if out != expected {
		t.Errorf("Expected %q, got %q", expected, out)
	}
}
