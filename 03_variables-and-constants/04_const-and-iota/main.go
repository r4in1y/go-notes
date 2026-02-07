package main

import "fmt"

func main() {
	// Constants in Go are immutable values that are known at compile time
	// Single constant
	// The := operator can't b used with constants
	const PI = 3.14159

	// Multiple constants
	const (
		StatusOk    = 200
		StatusError = 500
	)

	// Typed constants
	const (
		MaxSize int  = 1024
		Debug   bool = false
	)

	// Untyped Constants
	const (
		Pi   = 3.142
		Name = "Bob"
		Yes  = true
	)

	// Constants can be used with different compatible types
	var f float64 = Pi
	var f32 float32 = Pi
	var c complex128 = Pi

	fmt.Println(f, f32, c)

	// iota
	// iota can only be used inside a const block
	const (
		Red    = iota // 0
		Green         // 1
		Blue          // 2
		Yellow        // 3
		Orange        // 4
	)

	fmt.Println(Red, Green, Blue, Yellow, Orange)
}
