# 🚀Lecture 011: **Understanding Return Values & Multiple Return Values in Go!**

---

## The Art of Function Output Management

In this lesson, we will explore **return values in functions** and **how Go allows functions to return multiple values**. Let's break it down step by step!

## 1️⃣ Returning a Single Value - The Basic Output 🔄

### 🔹 Single Return Value Anatomy

| Component | Purpose | Syntax Example |
|-----------|---------|----------------|
| **Return Type** | Declares what function outputs | `func add(...) int` |
| **Return Statement** | Sends value back to caller | `return sum` |
| **Value Assignment** | Captures returned value | `result := add(a, b)` |

### 🔹 Single Return Implementation

````go
package main

import "fmt"

// Function that adds two numbers and returns the sum
func add(number1 int, number2 int) int {
    sum := number1 + number2 // Perform addition
    return sum               // Return the result
}

func main() {
    a := 10
    b := 20

    sum := add(a, b)  // Store the returned value in 'sum'

    fmt.Println(sum)  // Print the result
}
````

### 🔹 Memory Flow for Single Return

```
Single Return Value Memory Flow:
┌─────────────────────────────────────┐
│            Main Function            │
├─────────────────────────────────────┤
│ a := 10  │ Memory: 0x1001 → 10      │
│ b := 20  │ Memory: 0x1002 → 20      │
└─────────────────┬───────────────────┘
                  │ add(a, b) called
                  ▼
┌─────────────────────────────────────┐
│           Add Function              │
├─────────────────────────────────────┤
│ number1 := 10 │ Memory: 0x2001 → 10 │
│ number2 := 20 │ Memory: 0x2002 → 20 │
│ sum := 30     │ Memory: 0x2003 → 30 │
│ return sum    │ Value: 30 returned  │ ◄── Value sent back
└─────────────────┬───────────────────┘
                  │ Return to main
                  ▼
┌─────────────────────────────────────┐
│        Value Assignment             │
├─────────────────────────────────────┤
│ sum := add(a,b) │ Memory: 0x1003 → 30 ◄── Returned value stored
│ fmt.Println(sum) │ Output: 30       │
└─────────────────────────────────────┘
```

### 🔹 Return vs Print Comparison

| Approach | Flexibility | Reusability | Testing |
|----------|-------------|-------------|---------|
| **Return Value** | High - can store, manipulate, pass | Excellent | Easy to test |
| **Direct Print** | Low - immediate output only | Limited | Hard to test |

### 🔹 Vending Machine Analogy 🥤

```
Vending Machine Process:
┌─────────────────────────────────────┐
│         Vending Machine             │
├─────────────────────────────────────┤
│ Input: $2 + Selection               │ ◄── Parameters (money, choice)
│ ┌─────────────────────────────────┐ │
│ │ Processing:                     │ │
│ │ 1. Validate payment             │ │ ◄── Function logic
│ │ 2. Check inventory              │ │
│ │ 3. Prepare drink                │ │
│ └─────────────────────────────────┘ │
│ Output: Cold Drink                  │ ◄── Return value
└─────────────────────────────────────┘
           │
           ▼
    Customer receives drink (return value)
```

---

## 2️⃣ Returning Multiple Values - The Power Feature 🎯

Go's unique ability to return multiple values from a single function call.

### 🔹 Multiple Return Value Structure

| Syntax Element | Purpose | Example |
|----------------|---------|---------|
| **Multiple Return Types** | Declare multiple outputs | `(int, int)` |
| **Multiple Return Values** | Return several results | `return sum, mul` |
| **Multiple Assignment** | Capture all returns | `p, q := getNumbers(a, b)` |

### 🔹 Multiple Return Implementation

````go
package main

import "fmt"

// Function that returns sum and multiplication of two numbers
func getNumbers(number1 int, number2 int) (int, int) {
    sum := number1 + number2
    mul := number1 * number2

    return sum, mul  // Returning two values
}

func main() {
    a := 10
    b := 20

    p, q := getNumbers(a, b) // Storing both returned values

    fmt.Println(p)  // Prints sum
    fmt.Println(q)  // Prints multiplication
}
````

### 🔹 Memory Flow for Multiple Returns

```
Multiple Return Values Memory Flow:
┌─────────────────────────────────────┐
│            Main Function            │
├─────────────────────────────────────┤
│ a := 10  │ Memory: 0x1001 → 10     │
│ b := 20  │ Memory: 0x1002 → 20     │
└─────────────────┬───────────────────┘
                  │ getNumbers(a, b) called
                  ▼
┌─────────────────────────────────────┐
│         GetNumbers Function         │
├─────────────────────────────────────┤
│ number1 := 10 │ Memory: 0x2001 → 10 │
│ number2 := 20 │ Memory: 0x2002 → 20 │
│ sum := 30     │ Memory: 0x2003 → 30 │
│ mul := 200    │ Memory: 0x2004 → 200│
│ return sum,mul│ Values: (30, 200)   │ ◄── Two values returned
└─────────────────┬───────────────────┘
                  │ Return to main
                  ▼
┌─────────────────────────────────────┐
│       Multiple Assignment           │
├─────────────────────────────────────┤
│ p, q := getNumbers(a,b)             │
│ p := 30       │ Memory: 0x1003 → 30 │ ◄── First return value
│ q := 200      │ Memory: 0x1004 → 200│ ◄── Second return value
└─────────────────────────────────────┘
```

### 🔹 Advanced Multiple Return Patterns

````go
// Named return values
func calculate(a, b int) (sum int, product int) {
    sum = a + b
    product = a * b
    return  // Automatically returns named variables
}

// Discarding unwanted return values with blank identifier
func processData(x, y int) (int, int, string) {
    return x + y, x * y, "processed"
}

func main() {
    sum, _ := processData(5, 3)  // Only capture sum, ignore product
    _, _, status := processData(5, 3)  // Only capture status
}
````

### 🔹 Multiple Return Use Cases

| Use Case | Example | Benefit |
|----------|---------|---------|
| **Error Handling** | `value, err := parseNumber("123")` | Safe error checking |
| **Min/Max Operations** | `min, max := findRange(numbers)` | Efficient computation |
| **Data Processing** | `result, metadata := processFile(path)` | Complete information |
| **Mathematical Operations** | `quotient, remainder := divide(10, 3)` | Related calculations |

### 🔹 Pizza Delivery App Analogy 🍕📱

```
Pizza Delivery Service:
┌─────────────────────────────────────┐
│        Delivery Service             │
├─────────────────────────────────────┤
│ Order: Pizza + Drink                │ ◄── Function call with parameters
│ ┌─────────────────────────────────┐ │
│ │ Processing:                     │ │
│ │ 1. Prepare pizza                │ │ ◄── Multiple operations
│ │ 2. Get drink from fridge        │ │
│ │ 3. Package both items           │ │
│ └─────────────────────────────────┘ │
│ Delivery: Pizza + Drink             │ ◄── Multiple return values
└─────────────────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│         Customer Receives           │
├─────────────────────────────────────┤
│ pizza, drink := orderDelivery()     │ ◄── Multiple assignment
│ enjoy(pizza)                        │
│ drink(drink)                        │
└─────────────────────────────────────┘
```

---

## 🔄 Return Value Patterns Comparison

| Pattern | Syntax | Use Case | Memory Efficiency |
|---------|--------|----------|-------------------|
| **Single Return** | `func add(a, b int) int` | Simple calculations | Optimal |
| **Multiple Return** | `func calc(a, b int) (int, int)` | Related operations | Very efficient |
| **Named Returns** | `func calc(a, b int) (sum, prod int)` | Clear documentation | Efficient |
| **Error Returns** | `func parse(s string) (int, error)` | Safe operations | Standard practice |

### 🔹 Best Practices for Return Values

````go
// ✅ GOOD: Clear, meaningful returns
func calculateCircle(radius float64) (area float64, circumference float64) {
    area = 3.14159 * radius * radius
    circumference = 2 * 3.14159 * radius
    return
}

// ✅ GOOD: Error handling pattern
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// ❌ AVOID: Too many return values
func processUser(id int) (string, int, string, bool, error) {
    // Too many returns make code hard to read
}
````

## 🚀 Final Summary

| Concept | Symbol | Key Advantage | Best Practice |
|---------|--------|---------------|---------------|
| **Single Return** | 🔄 | Simple, clear purpose | Use for focused operations |
| **Multiple Return** | 🎯 | Efficient related operations | Group logically related values |
| **Named Returns** | 📝 | Self-documenting code | Use for complex functions |
| **Error Returns** | ⚠️ | Safe error handling | Always handle potential errors |

### 🔹 Return Value Decision Guide

| Question | Single Return | Multiple Return |
|----------|---------------|-----------------|
| **Is it one logical result?** | ✅ Yes | ❌ No |
| **Are results related?** | ❌ Not necessarily | ✅ Yes |
| **Need error handling?** | Sometimes | Often with error as second return |
| **Function complexity?** | Simple | Moderate |

### 🔹 Memory Management Benefits

1. **Efficient Computation** - Calculate multiple related values in one pass
2. **Reduced Function Calls** - Get multiple results without separate calls
3. **Cleaner Code** - Avoid global variables or complex structures
4. **Type Safety** - Each return value has its specific type
5. **Performance** - Minimal memory overhead for multiple returns

Remember: Return values are your function's way of communicating results back to the caller. Single returns keep things simple, while multiple returns let you efficiently handle related operations - choose based on your specific needs!
