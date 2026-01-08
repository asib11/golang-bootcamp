package main

import ("fmt")


func main() {

	// 1. slice from an existing array
	arr := [6]string{"this", "is", "a", "very", "simple", "array"}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	S := arr[1:4] // Length = 4-1=3, Capacity = 6-1=5
	fmt.Println(S)
	fmt.Println(len(S))
	fmt.Println(cap(S))

	S1 := S[1:2] // Length = 2-1=1, Capacity = 5-1=4
	fmt.Println(S1)
	fmt.Println(len(S1))
	fmt.Println(cap(S1))
}




/*


* 1. slice from an existing array
* 2. slice literal
* 3. slice from a slice

OUTPUT:
[this is a very simple array]
6
6
[is a very]
3
5
[a]
1
4


*/
