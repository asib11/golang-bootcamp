package main

import "fmt"

func add(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	
	return sum, mul
}

func main() {
	var a int
	var b int

	fmt.Print("enter number a: ")
	fmt.Scan(&a)

	fmt.Print("enter number b: ")
	fmt.Scan(&b)

	p, q := add(a, b)
	fmt.Println("Sum is:", p)
	fmt.Println("Product is:", q)
}
