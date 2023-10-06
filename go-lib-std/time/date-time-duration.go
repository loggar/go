package date_and_time

import (
	"fmt"
	"time"
)

func parse_duration_from_string() {
	screen_time, err := time.ParseDuration("6h30m")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", screen_time)
	fmt.Println(screen_time.Hours())
	fmt.Println(screen_time.Minutes())
}

func parse_duration_from_times() {
	newYear := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// current time is 2023-06-18 15:27:12 +0000 UTC
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Since(newYear).Hours())

	nextNewYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(nextNewYear.Sub(newYear).Hours())
	fmt.Println(nextNewYear.Sub(newYear).String())
}
