package timeutil

import (
	"context"
	"fmt"
	"time"
)

func workWithContext(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second): // Simulate work taking 2 seconds
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled or timed out")
	}
}
