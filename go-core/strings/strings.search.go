package strings

import (
	"fmt"
	"strings"
)

func search() {
	ss := "Sammy Shark"

	fmt.Println(strings.HasPrefix(ss, "Sammy"))
	fmt.Println(strings.HasSuffix(ss, "Shark"))

	fmt.Println(strings.Contains(ss, "Sh"))
	fmt.Println(strings.Count(ss, "S"))
	fmt.Println(strings.Count(ss, "s"))
}
