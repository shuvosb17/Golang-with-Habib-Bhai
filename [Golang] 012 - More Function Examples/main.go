package main

import "fmt"

func printSomething() {
	fmt.Println("Hello World!")
}

func getName(name string) {
	fmt.Println("Welcome to the course ,", name)
}

func main() {
	printSomething()
	getName("Shuvo")
}
