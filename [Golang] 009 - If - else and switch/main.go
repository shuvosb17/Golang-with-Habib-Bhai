package main

import "fmt"

func main() {
	age := 20

	if age > 18 {
		fmt.Println("Your are ready to go!")
	} else if age == 20 {
		fmt.Println("Your are not ready to go but you can prepare!")
	} else {
		fmt.Println("Dont Go")
	}

}
