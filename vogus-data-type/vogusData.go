package main

import "fmt"

func main() {
	var a int8 = 127
	var b int8 = -127

	var v uint8 = 255 // unsigned (0 or positive only)
	var f float32 = 3.14
	var f2 float64 = 3.14159265358979323846

	var byte1 byte = 255 // byte is an alias for uint8 

	var x bool = true
	var r rune = '😀' // rune is an alias for int32, represents a Unicode code point

	var s string = "Hello, Vogus!" // string is a sequence of bytes
	fmt.Printf("a: %T %d\n", a, a)
	fmt.Printf("b: %T %d\n", b, b)
	fmt.Printf("v: %T %d\n", v, v)
	fmt.Printf("f: %T %f\n", f, f)
	fmt.Printf("f2: %T %f\n", f2, f2)
	fmt.Printf("byte: %T %d\n", byte1, byte1)
	fmt.Printf("r: %T %c\n", r, r)
	fmt.Printf("s: %T %s\n", s, s)
	fmt.Printf("x: %T %t\n", x, x)
}