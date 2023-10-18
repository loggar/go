package mydebug_test

import (
	"fmt"
	"os"
	"strings"
)

func checkGoDebug() {
	godebug := os.Getenv("GODEBUG")

	// Split the value by commas
	values := strings.Split(godebug, ",")

	// Check if x509sha1=1 is one of the values
	for _, value := range values {
		if value == "x509sha1=1" {
			fmt.Println("GODEBUG contains x509sha1=1")
			return
		}
	}

	fmt.Println("GODEBUG does not contain x509sha1=1")
}
