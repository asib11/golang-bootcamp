package main

import "fmt"

func processNumbers(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func add( x int, y int) {
	s := x + y
	fmt.Println("sum is: ", s)
}

func main() {
	processNumbers(5, 7, add)
}