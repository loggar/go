package datetime

import (
	"fmt"
	"time"
)

func parse_from_string() {
	customDate := "2023-04-26"
	t, err := time.Parse("2006-01-02", customDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customDate)
	fmt.Println(t)

	customDate = "2023-0426"
	t, err = time.Parse("2006-01-02", customDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customDate)
	fmt.Println(t)
}
