package main

import "fmt"

func main() {
	// age := 20

	// if age <18 {
	// 	fmt.Println("Minor")
	// } else if age >=18 && age <65 {
	// 	fmt.Println("Adult")
	// } else {
	// 	fmt.Println("Senior Citizen")
	// }

	a := 3
	switch a {
	case 1:
		fmt.Println("One")
	case 2, 3 :
		fmt.Println("Two or Three")
	default:
		fmt.Println("Other Number")
	}
}
