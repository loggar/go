package channelclosing

import "sync"

// https://go101.org/article/channel-closing.html

type MyChannel2[T any] struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel2[T any]() *MyChannel2[T] {
	return &MyChannel2[T]{C: make(chan T)}
}

func (mc *MyChannel2[T]) SafeClose2() {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
}

func (mc *MyChannel2[T]) IsClosed2() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
