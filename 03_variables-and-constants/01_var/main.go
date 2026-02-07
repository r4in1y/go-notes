package main

import "fmt"

func main() {
	// A variable declaration creates one or more variable
	// A varibale can be be declared without initialization
	var ( // Grouped declaration
		i    int
		j, k float64
	)

	var fruit string

	// initalizing the varibales
	i = 5
	j, k = 6, 7
	fruit = "cherry"

	// declaring and initialization
	var name string = "Bob Rob"
	var age int = 32
	var height_in_cm float64 = 171.2

	// A varibale declared without a type is given the type of the corresponding initailization
	var city = "New York" // city is string
	var weather = 4.3     // weather is float64
	var cold = true       // cold is bool
	var n = 4             // k is int

	fmt.Print(i, j, k, fruit, name, age, height_in_cm, city, weather, cold, n)
}
