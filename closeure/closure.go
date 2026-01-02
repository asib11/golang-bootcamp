package main

import "fmt"

const a = 10
var p = 100

func outer() func() {
	money := 200
	age := 30

	fmt.Println("Age in outer:", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Init function called")
}


// Output:// Init function called
// Age in outer: 30
// 310
// 410
// Age in outer: 30
// 310
// 410

// Explanation:
// The code demonstrates the concept of closures in Go. The `outer` function defines local variables `money` and `age`, and returns an inner function `show` that modifies and prints the `money` variable. Each time `outer` is called, a new instance of `money` is created, allowing separate state for each closure. The `call` function invokes `outer` twice, creating two independent closures that maintain their own state for `money`. The `init` function is executed before `main`, printing a message when the program starts.	

// The code demonstrates the concept of closures in Go. The `outer` function defines local variables `money` and `age`, and returns an inner function `show` that modifies and prints the `money` variable. Each time `outer` is called, a new instance of `money` is created, allowing separate state for each closure. The `call` function invokes `outer` twice, creating two independent closures that maintain their own state for `money`. The `init` function is executed before `main`, printing a message when the program starts.

//escape analysis: money escapes to heap
