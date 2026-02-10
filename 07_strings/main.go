package main

import "fmt"

func main() {
	message := "Hello world" // this is a string literal
	fmt.Println(message)

	txt := `this is a raw string literal
	it can have multiple lines without the need for /n`

	fmt.Println(txt)

	// string concatenation
	firstName := "Jhon"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
}
