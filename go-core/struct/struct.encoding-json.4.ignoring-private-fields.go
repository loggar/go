package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User_4 struct {
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

func main_a_4() {
	u := &User_4{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
