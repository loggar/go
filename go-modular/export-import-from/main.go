package main

import (
	"fmt"

	"export-import-from/innerpackage"
	"export-import-from/secondinnerpackage"
)

// go run .
// # or
// go run main.go same-package.go

func main() {
	fnSamePackage()

	innerpackage.FnExported()

	fmt.Println(innerpackage.VariableABC)
	fmt.Println(secondinnerpackage.VariableExported)
}
