package main

import "fmt"

const a = 10 // const = code segment
var p = 100 // var = data segment


func call() { // function = code segment -> stack segment
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Sum:", z)
	}

	add(a, p)
	add(20, 30)
}

func main() { // function = code segment -> stack segment
	call()
	fmt.Println("Main function executed")
}

func init() { // function = code segment -> stack segment
	fmt.Println("init function called")
}

// In this code, we have different segments of memory being utilized:
// 1. Code Segment: This is where the compiled code of the program resides. Functions like main, init, and call are stored here.
// 2. Data Segment: This segment holds global and static variables. The variable 'p' is stored here.
// 3. Stack Segment: This segment is used for function calls and local variables. When functions like main, init, and call are invoked, their local variables and parameters are stored in the stack segment.
// 4. Constant Segment: This segment holds constant values. The constant 'a' is stored here.

// fix things = code segement