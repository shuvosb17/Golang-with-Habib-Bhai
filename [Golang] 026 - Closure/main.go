package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	age := 30
	money := 100

	fmt.Println(age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}
func main() {
	call()
}

func init() {
	fmt.Println("_______Bank_________")
}
