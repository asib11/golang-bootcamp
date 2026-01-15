package main

import ("fmt")

func print(numbers ...int) {
	fmt.Println("Numbers:", numbers)
}

func main() {
	print(1, 2, 3, 4, 5)
}