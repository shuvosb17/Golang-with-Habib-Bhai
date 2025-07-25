/*package main

import "fmt"

func add(number1 int, number2 int) int {
	sum := number1 + number2

	return sum
}

func main() {
	a := 10
	b := 20

	sum := add(a, b)

	fmt.Println(sum)
}
*/

package main

import "fmt"

func getNumbers(number1 int, number2 int) (int, int) {
	sum := number1 + number2
	mul := number1 * number2

	return sum, mul
}

func main() {
	a := 10
	b := 20

	p, q := getNumbers(a, b)

	fmt.Println(p)
	fmt.Println(q)
}
