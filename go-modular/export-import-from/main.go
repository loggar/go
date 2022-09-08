package main

import (
	"export-import-from/innerpackage"
	"fmt"
)

func main() {
	fnSamePackage()

	innerpackage.FnExported()

	fmt.Println(innerpackage.VariableABC)
}
