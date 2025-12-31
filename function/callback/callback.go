package main

import "fmt"

func call() func(int, int) {
	return add
}

func add(x int, y int) {
	z := x +y
	fmt.Println("Sum is: ", z)
}

func main( ) {
	sum := call()
	sum(10, 20)
}