package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// infinite loop
	// for {
	// 	fmt.Println("This will run forever")
	// }

	var age int = 18
	for {
		fmt.Println("checking...")

		if age >= 18 {
			break
		}
	}

	// whille loop
	for age < 20 {
		fmt.Println(age)
		age++

	}

	//
}
