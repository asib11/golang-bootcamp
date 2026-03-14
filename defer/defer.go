package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first", result) 

	show := func() {
		result = result + 10
		fmt.Println("defer", result) 
	}
	defer show() // defer will execute after the surrounding function returns
	
	result = 5

	p := func(a int) {
		fmt.Println("inside p", a)
	}
	
	defer p(result)
	defer fmt.Println(result) 

	fmt.Println("second", result)
	return
}

func main() {
	a := calculate()
	fmt.Println("main", a)
}