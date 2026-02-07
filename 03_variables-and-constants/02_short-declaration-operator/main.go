package main

import "fmt"

func main() {
	// The short way to declare and initalize a variable is using the := operator.
	// It’s concise and handles the type inference for you automatically.
	name := "Bob"
	age := 32
	fruit := "cherry"
	city, weather := "London", 12.5
	isReady := true

	fmt.Println(name, age, fruit, city, weather, isReady)
}
