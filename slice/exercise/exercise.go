package main

import ("fmt")

func changeSlice(s []int) []int {
	s[0] = 10
	s = append(s, 11)
	return s
}

func main() {
	x := []int{1, 2, 3, 4, 5} // len=5, cap=5
	x = append(x, 6)        // len=6, cap=10 (new underlying array)
	x = append(x, 7)        // len=7, cap=10
	a := x[4:]
	y := changeSlice(a)       // y is a new slice

	fmt.Println("x:", x) // x: [1 2 3 4 10 6 7]
	fmt.Println("y:", y) // y: [10 6 7 11]
}