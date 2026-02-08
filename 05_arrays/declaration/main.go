package main

import "fmt"

func main() {
	// array declaration
	var my_arr [5]int
	// declaration and initailization
	var arr [5]int = [5]int{1, 2, 3}

	fmt.Println(my_arr, arr)
	// short declaration with initialization
	ages := [5]int{16, 17, 18, 19, 15}
	fruits := [3]string{"cherry", "mango", "apple"}

	fmt.Println(ages, fruits)

	// size determined by initializer
	days := [...]string{"mon", "tue", "wed"}
	fmt.Println(days)

	scores := [...]int{1: 3, 5: 5}
	fmt.Println(scores)

}
