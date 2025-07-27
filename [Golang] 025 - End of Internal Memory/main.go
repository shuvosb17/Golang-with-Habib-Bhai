package main

import "fmt"

const a = 10 // Stored in Code Segment

var p = 20 // Stored in Data Segment

func call() { // Code Segment
	add := func(x int, y int) { // Function expression
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() { // Code Segment
	call()
	fmt.Println(a)
}

func init() { // Code Segment
	fmt.Println("Hi! This is Starting")
}
