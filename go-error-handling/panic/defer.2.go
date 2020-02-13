package main

import (
	"io"
	"os"
)

// CopyFileProblematic ref from https://blog.golang.org/defer-panic-and-recover
func CopyFileProblematic(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

/*
This works, but there is a bug. If the call to os.Create fails, the function will return without closing the source file. This can be easily remedied by putting a call to src.Close before the second return statement, but if the function were more complex the problem might not be so easily noticed and resolved. By introducing defer statements we can ensure that the files are always closed:
*/

// CopyFile with defer statement
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
