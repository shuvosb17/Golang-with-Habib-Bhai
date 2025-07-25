package main

import "fmt"

func main() {
	//a := 10
	var a int = 10 //Alternative way to declare
	const p = 100
	fmt.Println(a)
	fmt.Println(p)
}


/*

// Package main defines a standalone executable program
package main

// Import the fmt package for formatted I/O operations
import "fmt"

// main is the entry point of the program
func main() {
    // Variable declaration and initialization
    // Short declaration: a := 10 (commented out)
    var a int = 10 // Explicit variable declaration with type and initialization
    
    // Constant declaration
    const p = 100 // Constant value that cannot be changed during execution
    
    // Output the values to console
    fmt.Println(a) // Print variable value: 10
    fmt.Println(p) // Print constant value: 100
}

*/
