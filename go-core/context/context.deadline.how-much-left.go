package main

import (
	"context"
	"fmt"
	"time"
)

func print() {
	// create a new context with a deadline
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

	// make sure to cancel the context when you're done
	defer cancel()

	// extract the deadline
	deadline, ok := ctx.Deadline()

	if ok {
		// calculate the time until the deadline
		timeLeft := time.Until(deadline)
		fmt.Println("Time until deadline:", timeLeft)
	} else {
		fmt.Println("No deadline set")
	}
}
