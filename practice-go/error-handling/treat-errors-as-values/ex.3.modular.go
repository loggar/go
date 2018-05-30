package main

import (
	"fmt"
)

type cntWriter struct {
	count int
}

func (cw *cntWriter) write(s string) {
	if cw.count >= 32 {
		return
	}
	fmt.Printf("%s\n", s)
	cw.count += len(s)
}

func (cw *cntWriter) written() int {
	return cw.count
}

func main() {
	cw := &cntWriter{}
	cw.write("sample text 1")
	cw.write("sample text 2")
	cw.write("sample text 3")
	cw.write("sample text 4")

	fmt.Printf("Written length : %d\n", cw.written())
}
