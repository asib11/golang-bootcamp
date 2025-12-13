package main

import "fmt"

var a int = 10

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Value of a in main:", a)
}

func init(){
	fmt.Println("Init function called, a =", a)
	a = 20
}
