package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 10)
	}
}

// Test tables
func TestSumTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		r int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.r {
			t.Errorf("Sum of (%d + %d) was incorrect, got: %d, want: %d", table.x, table.y, total, table.r)
		}
	}
}
