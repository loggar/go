package date_and_time

import (
	"fmt"
	"time"
)

func time_ticker() {
	ticker := time.NewTicker(time.Second * 2)
	counter := 0
	for {
		select {
		case <-ticker.C:
			// api calls, call to database after specific time intervals
			counter++
			fmt.Println("Tick", counter)
		}
		if counter == 5 {
			ticker.Stop()
			return
		}
	}
}
