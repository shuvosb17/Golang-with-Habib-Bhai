package main

import "fmt"

var (
	x = 30
	y = 40
)

func add(p int, q int) int {
	sum := p + q
	return sum
}

func display(num int) {
	fmt.Println(num)
}

func main() {
	a := 20
	b := 20

	sum1 := add(a, b)
	display(sum1)

	sum2 := add(x, b)
	display(sum2)
}
