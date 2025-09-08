package main

import "fmt"

type User struct {
	name string
	age  int
}

func PrintDetails(usr User) {
	fmt.Println("Name: ", usr.name)
	fmt.Println("Age : ", usr.age)
}

func main() {
	user1 := User{
		name: "Shuvo",
		age:  22,
	}

	user2 := User{
		name: "Tanvir",
		age:  22,
	}

	PrintDetails(user1)
	PrintDetails(user2)

}

func init() {
	fmt.Println("_____User Details______")
}
