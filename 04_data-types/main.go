package main

import "fmt"

func main() {
	// integers
	var x int = 32    // this is signed(negative and positive numbers) int can be given a given bit (int8, int16, int32, int64)
	var y uint = 1032 // this is unsigned(positive numbers only) unit8, unit16, unit32, unit64

	// floating point
	var f float32
	f = 16.1234

	// complex numbers
	c := 2i
	c2 := 1 + 2i
	c3 := complex(12, 56)
	rl := real(c3)  // returns 12
	img := imag(c3) // returns 56

	// boolean type
	isReady := true
	active := false

	// rune represents unicode
	var (
		r rune = 'A'
		s rune = '世'
	)

	// strings
	var (
		name    = "Bob"
		message = `Hello
				   world in multiline`
	)
	length := len(name)
	char := name[0]
	slice := name[0:2]
	fmt.Println(x, y, f, c, c2, rl, img, isReady, active, r, s, length, char, slice, message)
}
