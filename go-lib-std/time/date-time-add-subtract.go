package datetime

import (
	"fmt"
	"time"
)

func add_time_and_date() {
	t := time.Now()
	fmt.Println(t)

	afterOneHour := t.Add(time.Hour * 1)
	fmt.Println(afterOneHour)

	afterOneDayTwoHours30minutes := t.AddDate(0, 0, 1).Add(time.Hour * 2).Add(time.Minute * 30)
	fmt.Println(afterOneDayTwoHours30minutes)

	afterThreeMonths15days := t.AddDate(0, 3, 15)
	fmt.Println(afterThreeMonths15days)
}

func subtract_time_and_date() {
	t := time.Now()
	fmt.Println(t)

	oneHourBack := t.Add(-1 * time.Hour)
	fmt.Println(oneHourBack)

	beforeOneYearTwoMonths := t.AddDate(-1, -2, 0)
	fmt.Println(beforeOneYearTwoMonths)
}
