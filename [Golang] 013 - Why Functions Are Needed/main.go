/*package main

import "fmt"

func main() {
	//print welcome meassage
	fmt.Println("Welcome to the Applicaton")

	//get user name as input
	fmt.Print("Enter Your Name: ")
	name := ""
	fmt.Scanln(&name)

	var num1 int
	var num2 int
	fmt.Print("Enter first number - ")
	fmt.Scanln(&num1)
	fmt.Print("Enter Second Number -")
	fmt.Scanln(&num2)

	sum := num1 + num2

	//Display results

	fmt.Println("Hello,", name)
	fmt.Println("Sum is =", sum)
	fmt.Println("Thank You for Using the Application")
}*/

package main

import "fmt"

func welcomeMessage() {
	fmt.Println("________________Welcome To The Application____________")
	fmt.Println()
	fmt.Println()
}

func getUserName() string {
	var name string
	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNUmbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Print("Enter your first Number - ")
	fmt.Scanln(&num1)
	fmt.Print("Enter yours Second NUmber - ")
	fmt.Scanln(&num2)
	fmt.Println()
	return num1, num2
}

func addTwoNumbers(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func Display(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Sum is : ", sum)
}

func GoodbyeMessage() {
	fmt.Println("Thanks for using the Application")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNUmbers()
	sum := addTwoNumbers(num1, num2)
	Display(name, sum)
	GoodbyeMessage()
}
