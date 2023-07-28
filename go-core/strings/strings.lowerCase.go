package strings

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func sample_1() {
	ss := "Sammy Shark"
	fmt.Println(strings.ToUpper(ss))
	fmt.Println(strings.ToLower(ss))
}

func sample_2() {
	s1 := "Ubuntu 22"
	s2 := "meet"
	s3 := "IND"
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToLower(s2))
	fmt.Println(strings.ToLower(s3))

	fmt.Printf("\n")
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.ToUpper(s3))

	/*
		go mod init
		go get golang.org/x/text/cases
		go get golang.org/x/text/language
	*/
	fmt.Printf("\n")
	cases := cases.Title(language.English)
	fmt.Println(cases.String(s1))
	fmt.Println(cases.String(s2))
	fmt.Println(cases.String(s3))
	/*
		Ubuntu 22
		Meet
		Ind
	*/
}
