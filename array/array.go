package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5} // array decleration
	fmt.Println("Array:", arr)

	var arr2 = [3]string{"Go", "Python", "JavaScript"}

	fmt.Println("Array2:", arr2)

	var arr3 [4]int // default value array, default value is 0 for int type, float is 0.0, string is "", bool is false
	fmt.Println("Array3:", arr3)

	arr3[0] = 10
	arr3[1] = 20
	arr3[2] = 30
	arr3[3] = 40

	fmt.Println("Updated Array3:", arr3)
}
