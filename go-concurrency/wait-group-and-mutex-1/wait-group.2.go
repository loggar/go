package main

import (
	"fmt"
	"sync"
)

func sendMail_no_need_wait_group_arg(address string) {
	fmt.Println("Sending mail to", address)
}

func main_2() {
	emails := []string{
		"recipient1@example.com",
		"recipient2@example.com",
		"xyz@example.com",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(emails))

	for _, email := range emails {
		mail := email
		go func(m string) {
			sendMail_no_need_wait_group_arg(m)
			// we don't have to change the signature of the function sendMail_no_need_wait_group_arg,
			// This keeps the functional logic of the function in its place and handles the concurrency on the go.
			wg.Done()
		}(mail)
	}
	wg.Wait()
	fmt.Println("All emails queued for sending")
	// Do other stuff
}
