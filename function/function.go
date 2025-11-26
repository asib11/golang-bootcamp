package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println( "Sum is:", sum)
}

func main() {
	var a int
	var b int

	fmt.Print("enter number a: ")
	fmt.Scan(&a)

	fmt.Print("enter number b: ")
	fmt.Scan(&b)

	add(a, b)
}
