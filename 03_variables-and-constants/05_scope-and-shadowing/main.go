package main

import "fmt"

func main() {
	// variables are only visible untill the end of the block {} they're  declared in
	x := 10 // Scope: main function
	fmt.Println(x)

	// innner blocks can see outer blocks. outer blocks can't see inner blocks
	if true {
		x := 5 // Scope: this if-block only. This shadows the outer x
		fmt.Println(x) // 5
	}

	fmt.Println(x) // 10
}
