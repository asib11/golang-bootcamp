package main

import "fmt"
import "example.com/package/mathlib"

func main() {
	result := mathlib.Add(5, 10)
	fmt.Println("Result from mathlib.Add:", result)
}
