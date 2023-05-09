package main

import (
	"bytes"
	"fmt"
	"sync"
)

/*
SyncedBuffer ...
*/
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main_1() {
	p := new(SyncedBuffer) // type *SyncedBuffer
	var v SyncedBuffer     // type  SyncedBuffer

	fmt.Println(p)
	fmt.Println(&v)
}
