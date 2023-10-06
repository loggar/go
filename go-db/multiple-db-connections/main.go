package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/loggar/go/go-db/multiple-db-connections/configs"
	"github.com/loggar/go/go-db/multiple-db-connections/dto"
	"github.com/loggar/go/go-db/multiple-db-connections/repo"
)

// Managing Multiple Database Connections in Golang
// https://github.com/chazool/go_medium_samples/tree/master/multi_db_coonnection_manager

func init() {
	// Initialize database connections during program startup.
	configs.InitDBConnections()
}

func main() {
	users := []dto.User{
		{
			UserName: "user1",
			Password: "test1",
		},
		{
			UserName: "user2",
			Password: "test2",
		},
	}

	// Create a new UserRepository with a specified connection name
	repo := repo.NewUserRepo("TEST_POSTGRES_CON")

	// save users
	err := repo.Save(users...)
	if err != nil {
		return
	}

	// get all users
	users, err = repo.FindAll()
	if err != nil {
		return
	}

	// print users
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	// Set up a channel to listen for OS signals (e.g., Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Wait for a signal (e.g., Ctrl+C) to gracefully exit the program
	<-c

	// Close database connections when the program terminates.
	defer configs.CloseDBConnections()
}
