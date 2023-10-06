package date_and_time

import (
	"fmt"
	"time"
)

func time_isBefore() {
	t := time.Now()
	fmt.Println(t)

	afterOneHour := t.Add(time.Hour * 1)
	fmt.Println(afterOneHour)

	isNowAfter := t.After(afterOneHour)
	isOneAfter := afterOneHour.After(t)
	fmt.Println(isNowAfter)
	fmt.Println(isOneAfter)

	isNowBefore := t.Before(afterOneHour)
	isOneBefore := afterOneHour.Before(t)
	fmt.Println(isNowBefore)
	fmt.Println(isOneBefore)
}

func time_equal() {
	t := time.Now()
	fmt.Println(t)

	afterOneHour := t.Add(time.Hour * 1)
	fmt.Println(afterOneHour)

	isNowEqual := t.Equal(afterOneHour)
	isEqual := time.Now().Equal(t)
	fmt.Println(isNowEqual)
	fmt.Println(isEqual)

	londonTimeZone, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
	}
	londonTime := t.In(londonTimeZone)

	fmt.Println(t)
	fmt.Println(londonTime)
	fmt.Println(t.Equal(londonTime))
}
