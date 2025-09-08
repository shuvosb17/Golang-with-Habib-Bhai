package main

import "fmt"

var arr2 = [3]string{"I", "Love", "You"}

func main() {
	// Fixed syntax error: use curly braces for array initialization
	arr := [2]int{3, 6}
	fmt.Println(arr)
	fmt.Println(arr2[1])

	// Additional array examples
	// 1. Array with inferred size
	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println("arr3:", arr3)

	// 2. Array initialization with specific indices
	arr4 := [5]string{0: "Zero", 2: "Two", 4: "Four"}
	fmt.Println("arr4:", arr4) // Other elements get zero values ("")

	// 3. Modifying array elements
	arr[0] = 100
	fmt.Println("Modified arr:", arr)

	// 4. Multi-dimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Matrix:", matrix)
}

func init() {
	fmt.Println("_____Array in Memory_____")
}
