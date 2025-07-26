package main

import "fmt"

func main() {
	//anonymous function
	//Immediately invoked function expression

	func(a int, b int) {
		res := a + b
		fmt.Println(res)
	}(5, 10)
}
