package channelclosing

import "sync"

// https://go101.org/article/channel-closing.html

type MyChannel[T any] struct {
	C    chan T
	once sync.Once
}

func NewMyChannel[T any]() *MyChannel[T] {
	return &MyChannel[T]{C: make(chan T)}
}

func (mc *MyChannel[T]) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}
