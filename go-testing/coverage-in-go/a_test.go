package main

import "testing"

func TestMax(t *testing.T) {
	if Max(1, 3) != 3 {
		t.Error("error, seharusnya 3")
	}

	// Jika pengecekan dibawah ini dihilangkan maka test coverage jadi 66.67%
	if Max(4, 3) != 4 {
		t.Error("error, seharusnya 4")
	}
}
