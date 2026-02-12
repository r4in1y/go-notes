package main

import (
	"fmt"
)

// struct definition
type User struct {
	UserName    string
	Email       string
	SignInCount int
	IsActive    bool
}

func main() {
	// initialization with field names
	user1 := User{
		UserName:    "alice",
		Email:       "alice@gmail.com",
		SignInCount: 5,
		IsActive:    true,
	}

	fmt.Println(user1)

	// initialization with default values. fields not specified are initialized to thier zero values of respecctive types
	user2 := User{
		UserName: "bob",
	}

	fmt.Println(user2)

	// initialization with pointers
	user3 := &User{
		UserName: "charile",
		Email:    "charlie@example.com",
	}

	fmt.Println(user3)
	user3.PrintInfo()
	user1.PrintInfo()
	user1.UpdateEmail("alicewonderland@fairy.com")
	fmt.Println(user1)
}

// Defining methods for structs
// value receiver
func (u User) PrintInfo() {
	fmt.Printf("Username: %s, Email: %s\n", u.UserName, u.Email)
}

// pointer receiver
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

type UserInfo interface {
	PrintInfo()
}

func ShowInfo(ui UserInfo) {
	ui.PrintInfo()
}
