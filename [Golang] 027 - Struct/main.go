package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	var user1 User

	user1 = User{
		name: "Shuvo",
		age:  22,
	}
	fmt.Println(user1.name)
	fmt.Println(user1.age)

	user2 := User{
		name: "Tanvir",
		age:  23,
	}
	fmt.Println(user2.name)
	fmt.Println(user2.age)
}

func init() {
	fmt.Println("_______User Prototype_____")
}
