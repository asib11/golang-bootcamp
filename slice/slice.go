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

	// slice from a slice
	S1 := S[1:2] // Length = 2-1=1, Capacity = 5-1=4
	fmt.Println(S1)
	fmt.Println(len(S1))
	fmt.Println(cap(S1))

	// 2. slice literal
	sliceLiteral := []int{10, 20, 30, 40, 50}
	fmt.Println(sliceLiteral)
	fmt.Println(len(sliceLiteral))
	fmt.Println(cap(sliceLiteral))

	// 4. function slice (make function)
	scliceMake := make([]int, 3)
	fmt.Println(scliceMake)
	fmt.Println(len(scliceMake))
	fmt.Println(cap(scliceMake))

	// 5. make function with len and cap
	s3 := make([]int, 2, 5)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))
	s3[1] = 100
	fmt.Println(s3)

	// 6. empty slice or nil slice
	var s4 []int
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	var s5 []int

	s5 = append(s5, 10)
	s5 = append(s5, 20)
	fmt.Println(s5)
	fmt.Println(len(s5))
	fmt.Println(cap(s5))
	
}




/*


* 1. slice from an existing array
* 2. slice literal
* 3. slice from a slice
* 4. function slice (make function)
* 5. make function with len and cap
* 6. empty slice or nil slice
* 7. slcie underlying array rule => 1024 -> 100% or doble the capacity increase, 1024 < increase by 25%

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
