package main

import "fmt"

type User struct {
	name string // member variable or field or property
	age int
}

func main() {
	user1 := User{ // instance or object
		name: "Asib", 
		age: 25,
		} 

	fmt.Println("User1:", user1)
	fmt.Println("User1 Name:", user1.name)
	fmt.Println("User1 Age:", user1.age)

	user2 := User{}
	user2.name = "John"
	user2.age = 30

	fmt.Println("User2:", user2)
	fmt.Println("User2 Name:", user2.name)
	fmt.Println("User2 Age:", user2.age)
}
