package main

import "fmt"

func main() {
	// Function expression assigned to a variable
	sum := func(a int, b int) int {
		return a + b
	}

	result := sum(5, 10) // calling the function expression
	fmt.Println("Sum of 5 and 10 is:", result)
}
