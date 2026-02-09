package main

import "fmt"

func main() {
	var n []int // empty slice
	nums := []int{12, 21, 30, 40}

	s := make([]int, 2)     // len = 5
	t := make([]int, 5, 10) // len = 5; capacity = 10

	arr := [5]int{1, 2, 3, 4, 5}
	sub_arr := arr[1:4] // slice from an array
	fmt.Println(nums, s, t, n, sub_arr)
}
