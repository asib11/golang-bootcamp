package main

import "fmt"

var (
	a int = 10
	b int = 20
)

func add(x int , y int) int {
	return x + y
}

func main() {
	var p int = 40
	var q int = 60

	fmt.Println(add(p, q))
	fmt.Println(add(a, b))

	if p <= 40 {
		var r int = 100
		fmt.Println("Value of r is:", r)
	}

}