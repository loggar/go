package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	name          string
	password      string
	preferredFish []string
	createdAt     time.Time
}

func main() {
	u := &User{
		name:      "Sammy the Shark",
		password:  "fisharegreat",
		createdAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
