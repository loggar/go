package main

import (
	"fmt"
	"log"
	"strconv"
)

func numberToString() {
	a := strconv.Itoa(12)
	fmt.Printf("%q\n", a)

	user := "Sammy"
	lines := 50

	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	fmt.Println(fmt.Sprint(421.034))

	f := 5524.53
	fmt.Println(fmt.Sprint(f))

}

func stringToNumbers() {
	linesYesterday := "50"
	linesToday := "108"

	yesterday, err := strconv.Atoi(linesYesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(linesToday)
	if err != nil {
		log.Fatal(err)
	}
	linesMore := today - yesterday

	fmt.Println(linesMore)
}

func stringsBytes() {
	a := "my string"

	b := []byte(a)

	c := string(b)

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)
}
