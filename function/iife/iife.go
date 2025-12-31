package main

import "fmt"

func main() {
	//anonymous function
	// Immediately Invoked Function Expression (IIFE) in Go
	func(a int, b int) {
		fmt.Println("Hello from IIFE!")
		fmt.Println("Sum of a and b is:", a+b)
	}(5, 10) // passing arguments 5 and 10 to the IIFE
}

func init() {
	fmt.Println("Init function called before main")
}