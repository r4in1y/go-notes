package main

import "fmt"

func main() {
	var m map[string]int // m = nil, not initialized
	// m["Age"] = 21 // this causes an erro
	n := make(map[string]int) // initilaized. n points to an initialized map. m does not.
	c := map[string]int{}     // c!=nil. initialized
	c["age"] = 42
	d := c["age"]
	n["status"] = 200
	i := n["status"]
	fruitsCount := map[string]string{
		"Apple":      "5",
		"Orange":     "16",
		"Banana":     "12",
		"Cherry":     "7",
		"Strawberry": "8",
	}

	for key, value := range fruitsCount {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println(len(fruitsCount), m, n, i, c, d)
}
