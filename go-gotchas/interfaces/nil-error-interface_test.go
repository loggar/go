package main

import (
	"fmt"
	"os"
	"testing"
)

func bar1() error {
	var err error // nil
	return err
}

func bar2() error {
	var err *os.PathError // a pointer of nil
	return err
}

// Test test
func Test(t *testing.T) {
	t.Run("compare two type of error interfaces", func(t *testing.T) {
		err1 := bar1()
		fmt.Println(err1 == nil) // true

		if err1 != nil {
			t.Errorf("got %v, want %v", err1 != nil, err1 == nil)
		}

		err2 := bar2()
		fmt.Println(err2 == nil) // false

		if err2 == nil {
			t.Errorf("got %v, want %v", err2 == nil, err2 != nil)
		}
	})

}
