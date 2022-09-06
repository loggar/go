package main

import "fmt"

func main() {
	// create variable and assign value
	// var conferenceName = "Go Conference"
	// syntactic sugar
	conferenceName := "Go Conference"

	const conferenceTickets = 50
	var remainingTickets uint = 50

	// unsigned integer, assignment error
	// remainingTickets = -1

	fmt.Println(conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")

	fmt.Printf("%v/%v available\n", remainingTickets, conferenceTickets)
	// variable types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("conferenceName pointer %v", &conferenceName)

	var userName string
	var userTickets int

	userName = "Tom" // fmt.Scan(&userName)
	userTickets = 2  // fmt.Scan(userTickets)

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

	// type casting
	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
