package main

import (
	"fmt"
)

type User struct {
	ID       string
	Username string
	Age      int
}

func Find(username string) (*User, error) { return nil, nil }
func SetAge(user *User, age int) error    { return nil }

func FindUser(username string) (*User, error) {
	u, err := Find(username)
	if err != nil {
		return nil, fmt.Errorf("FindUser: failed executing db query: %w", err)
	}
	return u, nil
}

func SetUserAge(u *User, age int) error {
	if err := SetAge(u, age); err != nil {
		return fmt.Errorf("SetUserAge: failed executing db update: %w", err)
	}
	return nil
}

func FindAndSetUserAge(username string, age int) error {
	var user *User
	var err error

	user, err = FindUser(username)
	if err != nil {
		return fmt.Errorf("FindAndSetUserAge: %w", err)
	}

	if err = SetUserAge(user, age); err != nil {
		return fmt.Errorf("FindAndSetUserAge: %w", err)
	}

	return nil
}

func fnHandler3() {
	if err := FindAndSetUserAge("bob@example.com", 21); err != nil {
		fmt.Printf("failed finding or updating user: %s", err)
		return
	}

	fmt.Println("successfully updated user's age")
}
