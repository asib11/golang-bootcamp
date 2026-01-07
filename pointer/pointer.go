package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Asib",
		Age : 27,
	}

	p := &user // pointer to struct

	fmt.Println("User Name:", *p)
	fmt.Printf("address of user struct: %p\n", p) // Print the address of the struct using %p

	var a int = 58
	var p1 *int = &a // pointer to int

	fmt.Println("Value of a using pointer:", *p1, "Address of a:", p1)

	*p1 = 100 // modifying value using pointer

	fmt.Println("Modified value of a:", *p1, "Address of a:", p1)
}


/*In Go, the %p format specifier is used in fmt.Printf and related functions to print the memory address of a variable (i.e., a pointer). It outputs the address in hexadecimal format.

What does %p do?

%p is used to print the pointer's value, which is the address the pointer is pointing to, rather than the value stored at that address.

This format specifier shows the memory location of the object in hexadecimal notation.

Explanation:

p is a pointer to x, so fmt.Printf("Address of x: %p\n", p) prints the address that p is holding (the address of x).

&x directly gets the address of the variable x, and using %p prints that memory address.

Why use %p?

Pointer addresses are generally used for debugging purposes or when working with low-level memory management (e.g., dealing with large data structures, shared memory, or optimizing memory usage).

It's useful to see where an object is stored in memory, especially when you're passing references (pointers) around in your code.
*/