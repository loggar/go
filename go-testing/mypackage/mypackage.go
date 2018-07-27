package mypackage

import (
	"sync/atomic"
)

var count uint64

// Count safely gets a unique count number.
func Count() uint64 {
	atomic.AddUint64(&count, 1)
	return count
}
