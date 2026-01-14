package main

import ("fmt")

func main() {
	var x []int // nil slice or empty slice or [], len=0, cap=0
	x = append(x, 19) // [19], len=1, cap=1
	x = append(x, 10) // [19,10], len=2, cap=2
	x = append(x, 25) // [19,10,25], len=3, cap=4

	y := x
	x = append(x, 30) // [19,10,25,30], len=4, cap=4
	y = append(y, 40) // [19,10,25,40], len=4, cap=4

	x[0] =10

	fmt.Println("x:", x) // x: [10 10 25 40]
	fmt.Println("y:", y) // y: [19 10 25 40]
}