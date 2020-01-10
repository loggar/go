package main

import "fmt"

// ByteSlice ...
type ByteSlice []byte

// Append ...
func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = slice
}

/*
we can do even better. If we modify our function so it looks like a standard Write method, like this,
*/

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	*p = slice
	return len(data), nil
}

/*
then the type *ByteSlice satisfies the standard interface io.Writer, which is handy. For instance, we can print into one.
*/

func writeUsage() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
}
