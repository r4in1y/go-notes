package main

import "fmt"

func main() {
	var (
		numbers = [4]int{1, 2, 3, 4}
		fruits  = [...]string{"cherry", "apple", "orange"}
		unic    = [3]rune{'A', '#', '@'}
	)

	first := numbers[0]
	fmt.Println(first)
	fmt.Println(len(fruits))
	code := unic[2]
	fmt.Printf("%c\n", code)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}

}
