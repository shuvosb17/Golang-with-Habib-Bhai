package main

import "fmt"

/*
GO FUNCTION CONCEPTS - COMPREHENSIVE NOTES & EXAMPLES
=======================================================

This file covers:
1. Parameters vs Arguments
2. First-order Functions
   i.   Standard/Named Functions
   ii.  Anonymous Functions
   iii. IIFE (Immediately Invoked Function Expressions)
   iv.  Function Expressions
3. Higher-order Functions
4. Callback Functions
5. Functions as First-class Citizens
*/

// =================== 1. PARAMETERS VS ARGUMENTS ===================

// Parameters are the variables in the function definition
// Arguments are the actual values passed to the function when called
func add(x int, y int) int {
	return x + y
}

// =================== 2. FIRST-ORDER FUNCTIONS ===================

// i. Standard/Named Function - has a name, doesn’t take/return functions
func greet(name string) string {
	return "Hello, " + name
}

// =================== 3. HIGHER-ORDER FUNCTIONS ===================

// Takes a function as parameter (higher-order)
func processOperation(x int, y int, operation func(a int, b int) int) int {
	return operation(x, y)
}

// Returns a function (higher-order)
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// =================== 4. CALLBACK FUNCTIONS ===================

// Simulates an async data fetch with a callback function
func fetchData(id int, callback func(data string)) {
	data := fmt.Sprintf("Data for ID %d", id)
	callback(data)
}

// =================== MAIN FUNCTION ===================

func main() {
	fmt.Println("========== GO FUNCTION CONCEPTS ==========")

	// -------- 1. PARAMETERS VS ARGUMENTS --------
	fmt.Println("\n1. Parameters vs Arguments:")
	result := add(10, 5) // 10 and 5 are arguments
	fmt.Printf("  add(10, 5) = %d\n", result)

	// -------- 2. FIRST-ORDER FUNCTIONS --------
	fmt.Println("\n2. First-order Functions:")

	// i. Standard/Named Function
	fmt.Printf("  i. Named function: %s\n", greet("Alice"))

	// ii. Anonymous Function
	fmt.Println("  ii. Anonymous function:")
	func(message string) {
		fmt.Printf("    Anonymous says: %s\n", message)
	}("I have no name")

	// iii. IIFE (Immediately Invoked Function Expression)
	fmt.Println("  iii. IIFE:")
	iifeResult := func(x, y int) int {
		return x * y
	}(4, 5)
	fmt.Printf("    IIFE result: 4 * 5 = %d\n", iifeResult)

	// iv. Function Expression
	fmt.Println("  iv. Function Expression:")
	subtract := func(a, b int) int {
		return a - b
	}
	fmt.Printf("    subtract(10, 5) = %d\n", subtract(10, 5))

	// -------- 3. HIGHER-ORDER FUNCTIONS --------
	fmt.Println("\n3. Higher-order Functions:")

	// i. Passing function as parameter
	fmt.Println("  i. Function taking function as parameter:")
	multiply := func(a, b int) int {
		return a * b
	}
	processResult := processOperation(5, 3, multiply)
	fmt.Printf("    processOperation(5, 3, multiply) = %d\n", processResult)

	// ii. Returning a function
	fmt.Println("  ii. Function returning function:")
	doubler := createMultiplier(2)
	tripler := createMultiplier(3)
	fmt.Printf("    doubler(7) = %d\n", doubler(7))
	fmt.Printf("    tripler(7) = %d\n", tripler(7))

	// -------- 4. CALLBACK FUNCTIONS --------
	fmt.Println("\n4. Callback Functions:")
	fetchData(42, func(data string) {
		fmt.Printf("  Received: %s\n", data)
	})

	// -------- 5. FUNCTIONS AS FIRST-CLASS CITIZENS --------
	fmt.Println("\n5. Functions as First-class Citizens:")

	// i. Assigning function to variable
	fmt.Println("  i. Assign function to variable:")
	adder := add
	fmt.Printf("    adder(3, 4) = %d\n", adder(3, 4))

	// ii. Storing functions in data structures
	fmt.Println("  ii. Store functions in data structures:")
	operations := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
	}
	fmt.Printf("    operations[\"add\"](5, 3) = %d\n", operations["add"](5, 3))
	fmt.Printf("    operations[\"multiply\"](5, 3) = %d\n", operations["multiply"](5, 3))

	// iii. Passing function from structure
	fmt.Println("  iii. Pass function as argument:")
	fmt.Printf("    processOperation(10, 2, operations[\"subtract\"]) = %d\n",
		processOperation(10, 2, operations["subtract"]))

	fmt.Println("\n========== END OF EXAMPLES ==========")
}
