package main

import "fmt"

type User struct {
	name string
	age int
}


func (u User) greet() { // receiver function
	fmt.Println("Hello, my name is", u.name, "and I am", u.age, "years old.")
}

func (u User) isAdult(a int)  { // another receiver function with parameter
	fmt.Println(u.name, "is adult:", a )
}

func main() {
	u1 := User{
		name: "Asib",
		age: 28,
	}
	u1.greet()
	u1.isAdult(18)
}