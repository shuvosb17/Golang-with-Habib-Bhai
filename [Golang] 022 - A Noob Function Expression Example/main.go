package main

import "fmt"

func main() {

	add := func(a int, b int) {
		res := (a + b)
		fmt.Println(res)
	}
	add(2, 4)
}
