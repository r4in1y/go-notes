package main

import "fmt"

func main() {
	var i int     // zero value is 0
	var j float64 // zero value is 0.0
	var b bool    // zero value is false
	var s string  // zero value is ""(empty string)

	fmt.Printf("%v %v %v %q\n", i, j, b, s)
}
