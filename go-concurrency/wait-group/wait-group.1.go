package main

import (
	"fmt"
	"sync"
)

func sendMail(address string, wg *sync.WaitGroup) {
	fmt.Println("Sending mail to", address)
	defer wg.Done()
	// Actual mail sending, smtp stuff
	// handle errors

	// client, err := smtp.Dial("smtp.example.com:587")
	// errr = client.Mail("sender@example.com")
	// err = client.Rcpt(address)
	// wc, err := client.Data()
	//_, err = wc.Write([]byte("This is the email body."))
	// err = wc.Close()
	// client.Quit()
}

func main_1() {
	emails := []string{
		"recipient1@example.com",
		"recipient2@example.com",
		"xyz@example.com",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(emails))

	for _, email := range emails {
		go sendMail(email, &wg)
	}
	wg.Wait()
	fmt.Println("All emails queued for sending")
	// Do other stuff
}
