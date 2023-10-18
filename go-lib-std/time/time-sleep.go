package datetime

import (
	"fmt"
	"time"
)

func time_sleep() {
	t1 := time.Now()
	time.Sleep(time.Second * 3)
	t2 := time.Now()
	duration := t2.Sub(t1)
	fmt.Println(duration)
}
