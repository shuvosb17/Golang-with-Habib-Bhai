# 🚀Lecture 012: **Comparing Different Types of Functions in Go!**

---

## The Complete Function Spectrum

In this lesson, we will compare **three different function types** in Go and understand when to use each pattern.

## 1️⃣ Function with No Parameters & No Return Values - The Autonomous Worker 🤖

### 🔹 No Parameters, No Return Structure

| Component | Purpose | Characteristics |
|-----------|---------|-----------------|
| **No Input** | Self-contained operation | Operates independently |
| **No Output** | Performs action only | Side effects (printing, file operations) |
| **Pure Action** | Does something, doesn't calculate | Event-driven functionality |

### 🔹 Implementation Example

````go
package main

import "fmt"

// Function without parameters and return value
func printSomething() {
    fmt.Println("Hello World!")
}

func main() {
    printSomething() // Calls the function
}
````

### 🔹 Memory and Execution Flow

```
No Parameters Function Flow:
┌─────────────────────────────────────┐
│            Main Function            │
├─────────────────────────────────────┤
│ printSomething() called             │ ◄── Simple function call
└─────────────────┬───────────────────┘
                  │ No data passed
                  ▼
┌─────────────────────────────────────┐
│       PrintSomething Function       │
├─────────────────────────────────────┤
│ No parameters to store              │
│ Execute: fmt.Println("Hello World!")│ ◄── Direct action
│ No return value                     │
└─────────────────┬───────────────────┘
                  │ Action completed
                  ▼
┌─────────────────────────────────────┐
│          Back to Main               │
├─────────────────────────────────────┤
│ Continue program execution          │ ◄── No value received
└─────────────────────────────────────┘
```

### 🔹 Real-world Use Cases

| Use Case | Example | Why This Pattern? |
|----------|---------|-------------------|
| **Initialization** | `setupDatabase()` | Self-contained setup |
| **Logging** | `printWelcomeMessage()` | Simple output tasks |
| **Cleanup** | `closeConnections()` | Independent maintenance |
| **Notifications** | `sendHeartbeat()` | Automated actions |

### 🔹 Morning Alarm Clock Analogy ⏰

```
Alarm Clock Behavior:
┌─────────────────────────────────────┐
│          Alarm Clock                │
├─────────────────────────────────────┤
│ Time: 7:00 AM                       │ ◄── Trigger condition
│ ┌─────────────────────────────────┐ │
│ │ ring() function:                │ │
│ │ • No input needed               │ │ ◄── Self-contained
│ │ • Just make sound               │ │
│ │ • No return value               │ │
│ └─────────────────────────────────┘ │
│ Output: BEEP BEEP BEEP!             │ ◄── Action performed
└─────────────────────────────────────┘
```

---

## 2️⃣ Function with Parameters but No Return Values - The Service Provider 🏪

### 🔹 Parameters Only Structure

| Component | Purpose | Characteristics |
|-----------|---------|-----------------|
| **Takes Input** | Receives data to process | Customizable behavior |
| **No Output** | Performs action based on input | Side effects with data |
| **Service Pattern** | Does something for you | Input-driven actions |

### 🔹 Implementation Example

````go
package main

import "fmt"

// Function with a parameter, but no return value
func getName(name string) {
    fmt.Println("Welcome to the course,", name)
}

func main() {
    getName("Shuvo") // Calls the function with "Shuvo" as input
}
````

### 🔹 Memory and Execution Flow

```
Parameters Only Function Flow:
┌─────────────────────────────────────┐
│            Main Function            │
├─────────────────────────────────────┤
│ getName("Shuvo") called             │
└─────────────────┬───────────────────┘
                  │ Pass "Shuvo"
                  ▼
┌─────────────────────────────────────┐
│         GetName Function            │
├─────────────────────────────────────┤
│ name := "Shuvo" │ Memory: 0x2001    │ ◄── Parameter stored
│ Process input:                      │
│ fmt.Println("Welcome...", name)     │ ◄── Action with input
│ No return statement                 │
└─────────────────┬───────────────────┘
                  │ Function completes
                  ▼
┌─────────────────────────────────────┐
│          Back to Main               │
├─────────────────────────────────────┤
│ Continue execution                  │ ◄── No value returned
│ Function memory cleaned up          │
└─────────────────────────────────────┘
```

### 🔹 Advanced Examples

````go
// Configuration function
func configureSettings(theme string, fontSize int) {
    fmt.Printf("Setting theme to %s with font size %d\n", theme, fontSize)
    // Apply settings logic here
}

// Logging function
func logMessage(level string, message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
}

// Data processing function
func processUserData(username string, email string) {
    fmt.Printf("Processing user: %s (%s)\n", username, email)
    // Database operations, validation, etc.
}
````

### 🔹 Cashier Service Analogy 🏪

```
Store Cashier Interaction:
┌─────────────────────────────────────┐
│            Customer                 │
├─────────────────────────────────────┤
│ Provides: Name "John"               │ ◄── Input parameter
└─────────────────┬───────────────────┘
                  │ Give name
                  ▼
┌─────────────────────────────────────┐
│             Cashier                 │
├─────────────────────────────────────┤
│ Receives: name = "John"             │
│ ┌─────────────────────────────────┐ │
│ │ greetCustomer(name):            │ │
│ │ • Process the name              │ │ ◄── Use input
│ │ • Say "Welcome, John!"          │ │
│ │ • No return needed              │ │
│ └─────────────────────────────────┘ │
│ Action: Verbal greeting             │ ◄── Service provided
└─────────────────────────────────────┘
```

---

## 3️⃣ Function with Parameters and Return Values - The Calculator 🧮

### 🔹 Full Function Structure

| Component | Purpose | Characteristics |
|-----------|---------|-----------------|
| **Takes Input** | Receives data to process | Data-driven |
| **Returns Output** | Provides calculated result | Value-producing |
| **Pure Function** | Input → Processing → Output | Predictable and testable |

### 🔹 Implementation Example

````go
package main

import "fmt"

// Function with parameters and a return value
func add(number1 int, number2 int) int {
    sum := number1 + number2
    return sum
}

func main() {
    a := 10
    b := 20

    sum := add(a, b) // Calls the function and stores returned value
    fmt.Println(sum)  // Prints the result
}
````

### 🔹 Memory and Execution Flow

```
Complete Function Flow:
┌─────────────────────────────────────┐
│            Main Function            │
├─────────────────────────────────────┤
│ a := 10, b := 20                   │
│ sum := add(a, b) called             │
└─────────────────┬───────────────────┘
                  │ Pass (10, 20)
                  ▼
┌─────────────────────────────────────┐
│           Add Function              │
├─────────────────────────────────────┤
│ number1 := 10   │ Memory: 0x2001   │ ◄── Parameters stored
│ number2 := 20   │ Memory: 0x2002   │
│ sum := 30       │ Memory: 0x2003   │ ◄── Calculation
│ return sum      │ Value: 30        │ ◄── Return result
└─────────────────┬───────────────────┘
                  │ Return 30
                  ▼
┌─────────────────────────────────────┐
│          Back to Main               │
├─────────────────────────────────────┤
│ sum := 30       │ Memory: 0x1003   │ ◄── Store returned value
│ fmt.Println(sum) │ Output: 30      │
└─────────────────────────────────────┘
```

### 🔹 Advanced Return Value Examples

````go
// Mathematical operations
func calculate(a, b float64, operation string) float64 {
    switch operation {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        if b != 0 {
            return a / b
        }
        return 0
    default:
        return 0
    }
}

// String processing
func formatName(firstName, lastName string) string {
    return fmt.Sprintf("%s, %s", lastName, firstName)
}

// Boolean logic
func isEven(number int) bool {
    return number%2 == 0
}
````

### 🔹 ATM Machine Analogy 🏦💳

```
ATM Transaction Process:
┌─────────────────────────────────────┐
│             Customer                │
├─────────────────────────────────────┤
│ Input: Card + PIN + Amount          │ ◄── Function parameters
└─────────────────┬───────────────────┘
                  │ Provide inputs
                  ▼
┌─────────────────────────────────────┐
│              ATM                    │
├─────────────────────────────────────┤
│ processWithdrawal(card, pin, amount)│
│ ┌─────────────────────────────────┐ │
│ │ • Validate card and PIN         │ │ ◄── Process inputs
│ │ • Check account balance         │ │
│ │ • Calculate new balance         │ │
│ │ • Prepare cash                  │ │
│ └─────────────────────────────────┘ │
│ Return: Cash + Receipt              │ ◄── Function return
└─────────────────┬───────────────────┘
                  │ Return values
                  ▼
┌─────────────────────────────────────┐
│         Customer Receives           │
├─────────────────────────────────────┤
│ cash, receipt := processWithdrawal()│ ◄── Capture returns
└─────────────────────────────────────┘
```

---

## 4️⃣ Function with Multiple Return Values - The Multi-Service Provider 🍔🥤

### 🔹 Multiple Return Implementation

````go
package main

import "fmt"

// Function that returns two values
func getNumbers(number1 int, number2 int) (int, int) {
    sum := number1 + number2
    mul := number1 * number2
    return sum, mul
}

func main() {
    a := 10
    b := 20

    p, q := getNumbers(a, b) // Storing two returned values
    fmt.Println(p) // Prints sum
    fmt.Println(q) // Prints multiplication
}
````

### 🔹 Restaurant Order Analogy 🍔🥤

```
Restaurant Multi-Item Order:
┌─────────────────────────────────────┐
│            Customer                 │
├─────────────────────────────────────┤
│ Order: "Burger and Drink"           │ ◄── Function call
└─────────────────┬───────────────────┘
                  │ Place order
                  ▼
┌─────────────────────────────────────┐
│           Restaurant                │
├─────────────────────────────────────┤
│ prepareOrder(burger, drink):        │
│ ┌─────────────────────────────────┐ │
│ │ • Cook burger                   │ │ ◄── Multiple operations
│ │ • Prepare drink                 │ │
│ │ • Package both items            │ │
│ └─────────────────────────────────┘ │
│ Return: Burger + Drink              │ ◄── Multiple returns
└─────────────────┬───────────────────┘
                  │ Deliver order
                  ▼
┌─────────────────────────────────────┐
│         Customer Receives           │
├─────────────────────────────────────┤
│ burger, drink := prepareOrder()     │ ◄── Multiple assignment
│ eat(burger)                         │
│ drink(drink)                        │
└─────────────────────────────────────┘
```

---

## 🔄 Function Type Comparison Matrix

| Function Type | Parameters | Return Values | Use Case | Memory Usage | Best For |
|---------------|------------|---------------|----------|--------------|----------|
| **No Params, No Return** | ❌ | ❌ | Autonomous actions | Minimal | Setup, cleanup, alerts |
| **Params, No Return** | ✅ | ❌ | Input-based actions | Low | Logging, display, processing |
| **Params & Return** | ✅ | ✅ | Calculations | Moderate | Math, transformations |
| **Multiple Return** | ✅ | ✅ Multiple | Related operations | Efficient | Complex calculations |

### 🔹 Function Selection Decision Tree

```
Function Type Selection:
┌─────────────────┐
│ Need input data?│
└────────┬────────┘
         │
    ┌────▼────┐
    │   NO    │              │   YES   │
    └────┬────┘              └────┬────┘
         │                        │
         ▼                        ▼
┌─────────────────┐      ┌─────────────────┐
│ Need to return  │      │ Need to return  │
│ values?         │      │ values?         │
└────────┬────────┘      └────────┬────────┘
         │                        │
    ┌────▼────┐              ┌────▼────┐
    │   NO    │    │   YES   │   NO    │    │   YES   │
    └────┬────┘    └────┬────┘    └────┬────┘    └────┬────┘
         │              │              │              │
         ▼              ▼              ▼              ▼
┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ No Params   │ │   Invalid   │ │ Params      │ │ Params &    │
│ No Return   │ │ Combination │ │ No Return   │ │ Return      │
└─────────────┘ └─────────────┘ └─────────────┘ └─────────────┘
```

## 🚀 Final Summary

| Pattern | Symbol | Core Purpose | Memory Pattern | Real-world Example |
|---------|--------|--------------|----------------|-------------------|
| **No Input/Output** | 🤖 | Autonomous actions | Self-contained | Alarm clock ringing |
| **Input Only** | 🏪 | Service provision | Input processing | Cashier greeting |
| **Input/Output** | 🧮 | Value calculation | Transform data | ATM transaction |
| **Multiple Output** | 🍔🥤 | Complex operations | Efficient multi-results | Restaurant order |

### 🔹 Best Practices by Function Type

1. **No Parameters, No Return**
   - Use for initialization, cleanup, notifications
   - Keep operations simple and focused
   - Avoid side effects when possible

2. **Parameters, No Return**
   - Perfect for logging, display, configuration
   - Validate input parameters
   - Use meaningful parameter names

3. **Parameters and Return**
   - Ideal for calculations and transformations
   - Make functions pure when possible
   - Handle edge cases in logic

4. **Multiple Returns**
   - Use for related operations
   - Consider error handling patterns
   - Keep return count reasonable (typically 2-3)

Remember: Choose your function type based on what you need to accomplish - simple actions need simple functions, while complex operations benefit from parameters and return values!
