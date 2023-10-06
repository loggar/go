package date_and_time

import (
	"fmt"
	"time"
)

func time_zone_and_location() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Location())

	newYorkTimeZone, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	londonTimeZone, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
	}
	newYorkTime := t.In(newYorkTimeZone)
	londonTime := t.In(londonTimeZone)

	//local time
	fmt.Println(t)

	// london time
	fmt.Println(londonTimeZone)
	fmt.Println(londonTime)

	// new york time
	fmt.Println(newYorkTimeZone)
	fmt.Println(newYorkTime)
}
