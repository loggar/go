package array1

import "fmt"

//nolint:unused // ignore unused variable warning
func arrayMultiDimensional() {
	var a [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	fmt.Printf("%q", a)
	// [["row 1 - column 1" "row 1 - column 2" "row 1 - column 3"]
	// ["row 2 - column 1" "row 2 - column 2" "row 2 - column 3"]]
}
