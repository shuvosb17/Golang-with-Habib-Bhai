# 🚀**[Golang] 022 - A Noob Function Expression Example**

---

## The Variable-Stored Function Revolution

Function expressions represent one of Go's most powerful features - treating functions as first-class citizens that can be stored in variables, passed around, and executed like any other value.

## 🔹 Complete Code Analysis

````go
package main

import "fmt"

func main() {
    add := func(a int, b int) {
        res := (a + b)
        fmt.Println(res)
    }
    add(2, 4)
}
````

---

## 🧠 Understanding Function Expressions

### 🔹 Function Expression Anatomy

| Component | Description | Example | Purpose |
|-----------|-------------|---------|---------|
| **Variable Assignment** | `add :=` assigns function to variable | `add` | Function storage |
| **Anonymous Function** | `func(a int, b int)` | No function name | Function definition |
| **Function Body** | `{ res := (a + b); ... }` | Logic implementation | Function behavior |
| **Function Invocation** | `add(2, 4)` | Call via variable | Function execution |

### 🔹 Function Expression vs Function Declaration

```
Function Types Comparison:
┌─────────────────────────────────────┐
│         Function Declaration        │
├─────────────────────────────────────┤
│ func add(a, b int) {                │ ◄── Named function
│   res := a + b                      │ ◄── Available throughout scope
│   fmt.Println(res)                  │ ◄── Can be called before definition
│ }                                   │
│ // Can call: add(2, 4) anywhere     │
├─────────────────────────────────────┤
│         Function Expression         │
├─────────────────────────────────────┤
│ add := func(a, b int) {             │ ◄── Anonymous function
│   res := a + b                      │ ◄── Stored in variable
│   fmt.Println(res)                  │ ◄── Must be defined before use
│ }                                   │
│ // Must call after: add(2, 4)       │
└─────────────────────────────────────┘
```

---

## 📊 Memory Allocation and Execution Flow

### Step 1: main() Function Initialization

```
Memory State: main() Function Start
┌─────────────────────────────────────┐
│           CODE SEGMENT              │
├─────────────────────────────────────┤
│ main() function loaded              │ ◄── Entry point ready
│ Anonymous function code compiled    │ ◄── Function code ready
├─────────────────────────────────────┤
│           STACK SEGMENT             │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │        main() Frame             │ │
│ │ Address: 0x7000                 │ │
│ │ Variable declaration in progress 
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

### Step 2: Function Expression Assignment

````go
add := func(a int, b int) {
    res := (a + b)
    fmt.Println(res)
}
````

```
Memory State: Function Expression Assigned
┌─────────────────────────────────────┐
│           HEAP MEMORY               │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │      Function Object            │ │
│ │ Address: 0x2000                 │ │
│ │ Code Pointer: Anonymous func    │ │ ◄── Points to function code
│ │ Closure Data: None (no captures)│ │ ◄── No captured variables
│ └─────────────────────────────────┘ │
├─────────────────────────────────────┤
│           STACK SEGMENT             │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │        main() Frame             │ │
│ │ Address: 0x7000                 │ │
│ │ Variables:                      │ │
│ │   add = 0x2000 (function ref)   │ │ ◄── Points to heap function object
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

### Step 3: Function Expression Invocation

````go
add(2, 4)  // Calling the function through the variable
````

```
Memory State: Function Expression Executing
┌─────────────────────────────────────┐
│           STACK SEGMENT             │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │    Anonymous Function Frame     │ │
│ │ Address: 0x7100                 │ │
│ │ Parameters:                     │ │
│ │   a = 2       │ Address: 0x7104 │ │ ◄── First parameter
│ │   b = 4       │ Address: 0x7108 │ │ ◄── Second parameter
│ │ Local Variables:                │ │
│ │   res = 6     │ Address: 0x710C │ │ ◄── Calculated result
│ │ Return Address: main() frame    │ │ ◄── Return to main
│ └─────────────────────────────────┘ │
│ ┌─────────────────────────────────┐ │
│ │        main() Frame             │ │
│ │ Address: 0x7000                 │ │
│ │ add = 0x2000 (function ref)     │ │ ◄── Function variable
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘

Console Output: 6
```

---

## 🚨 Declaration Order - Why It Matters

### 🔹 The Problem: Using Before Declaration

````go
// ❌ This will cause a compilation error
func main() {
    add(2, 4)  // Error: undefined: add
    
    add := func(a int, b int) {
        res := (a + b)
        fmt.Println(res)
    }
}
````

### 🔹 Compilation Error Analysis

| Aspect | Explanation | Technical Detail |
|--------|-------------|------------------|
| **Error Type** | Compile-time error | `undefined: add` |
| **Root Cause** | Variable used before declaration | Go requires declaration before use |
| **Scope Rules** | No variable hoisting in Go | Unlike JavaScript, no automatic lifting |
| **Static Typing** | Compile-time variable resolution | All variables must be known at compile time |

### 🔹 Declaration Order Visualization

```
Variable Lifecycle in Go:
┌─────────────────────────────────────┐
│           Before Declaration        │
├─────────────────────────────────────┤
│ Variable 'add': DOES NOT EXIST      │ ◄── Compiler has no knowledge
│ Attempting to use: COMPILATION ERROR│ ◄── Undefined variable
├─────────────────────────────────────┤
│           At Declaration            │
├─────────────────────────────────────┤
│ add := func(...) {...}              │ ◄── Variable comes into existence
│ Variable 'add': NOW EXISTS          │ ◄── Available for use
├─────────────────────────────────────┤
│           After Declaration         │
├─────────────────────────────────────┤
│ add(2, 4)  // ✅ Valid               ◄── Can be called safely
│ Variable 'add': ACCESSIBLE          │ ◄── Ready for invocation
└─────────────────────────────────────┘
```

### 🔹 Correct vs Incorrect Order

````go
// ❌ INCORRECT: Use before declaration
func incorrectOrder() {
    result := calculate(5, 3)  // Error: undefined: calculate
    
    calculate := func(x, y int) int {
        return x + y
    }
    
    fmt.Println(result)
}

// ✅ CORRECT: Declaration before use
func correctOrder() {
    calculate := func(x, y int) int {
        return x + y
    }
    
    result := calculate(5, 3)  // Works perfectly
    fmt.Println(result)
}
````

---

## 🎯 Practical Function Expression Applications

### 🔹 Function Expression as Variable

````go
func variableExample() {
    // Different mathematical operations
    var operation func(int, int) int
    
    // Assign different functions based on conditions
    useAddition := true
    
    if useAddition {
        operation = func(a, b int) int {
            return a + b
        }
    } else {
        operation = func(a, b int) int {
            return a * b
        }
    }
    
    result := operation(10, 5)
    fmt.Printf("Result: %d\n", result)  // Output: 15
}
````

### 🔹 Function Expression as Parameter

````go
func applyOperation(op func(int, int) int, a, b int) int {
    return op(a, b)
}

func parameterExample() {
    // Pass function expression as argument
    result1 := applyOperation(func(x, y int) int {
        return x + y
    }, 5, 3)
    
    result2 := applyOperation(func(x, y int) int {
        return x * y
    }, 5, 3)
    
    fmt.Printf("Addition: %d, Multiplication: %d\n", result1, result2)
}
````

### 🔹 Function Expression as Return Value

````go
func createCalculator(operation string) func(int, int) int {
    switch operation {
    case "add":
        return func(a, b int) int {
            return a + b
        }
    case "multiply":
        return func(a, b int) int {
            return a * b
        }
    case "subtract":
        return func(a, b int) int {
            return a - b
        }
    default:
        return func(a, b int) int {
            return 0
        }
    }
}

func returnValueExample() {
    adder := createCalculator("add")
    multiplier := createCalculator("multiply")
    
    fmt.Printf("Add: %d\n", adder(10, 5))        // 15
    fmt.Printf("Multiply: %d\n", multiplier(10, 5)) // 50
}
````

### 🔹 Closure with Function Expression

````go
func createCounter() func() int {
    count := 0  // Variable captured in closure
    
    return func() int {
        count++  // Modifies captured variable
        return count
    }
}

func closureExample() {
    counter1 := createCounter()
    counter2 := createCounter()  // Independent counter
    
    fmt.Printf("Counter1: %d\n", counter1()) // 1
    fmt.Printf("Counter1: %d\n", counter1()) // 2
    fmt.Printf("Counter2: %d\n", counter2()) // 1
    fmt.Printf("Counter1: %d\n", counter1()) // 3
}
````

---

## 🧠 Interview Questions & Comprehensive Answers

### 🔹 Q1: What is the difference between a function declaration and a function expression in Go?

**A: Comprehensive Comparison**

| Aspect | Function Declaration | Function Expression |
|--------|---------------------|-------------------|
| **Syntax** | `func name(params) { ... }` | `var := func(params) { ... }` |
| **Naming** | Has explicit name | Anonymous (stored in variable) |
| **Hoisting** | Available throughout scope | Must be declared before use |
| **Storage** | Function table | Variable holding function reference |
| **Scope** | Package or block level | Variable scope rules apply |

#### Detailed Examples

````go
// Function Declaration
func declaredAdd(a, b int) int {
    return a + b
}

// Can be called before definition (within same scope)
func exampleUsage() {
    result := declaredAdd(5, 3)  // ✅ Works
    fmt.Println(result)
}

// Function Expression
func expressionExample() {
    // Must declare before use
    var expressionAdd = func(a, b int) int {
        return a + b
    }
    
    result := expressionAdd(5, 3)  // ✅ Works after declaration
    fmt.Println(result)
}
````

#### Memory Allocation Differences

```
Memory Layout Comparison:

Function Declaration:
┌─────────────────────────────────────┐
│           CODE SEGMENT              │
├─────────────────────────────────────┤
│ declaredAdd: [function code]        │ ◄── Direct function storage
└─────────────────────────────────────┘

Function Expression:
┌─────────────────────────────────────┐
│           HEAP MEMORY               │
├─────────────────────────────────────┤
│ Function Object: [function code]    │ ◄── Heap-allocated function
├─────────────────────────────────────┤
│           STACK MEMORY              │
├─────────────────────────────────────┤
│ expressionAdd: [pointer to heap]    │ ◄── Variable stores reference
└─────────────────────────────────────┘
```

### 🔹 Q2: Can function expressions in Go create closures? How?

**A: Closure Mechanics with Function Expressions**

#### Basic Closure Example

````go
func createMultiplier(factor int) func(int) int {
    return func(value int) int {
        return value * factor  // Captures 'factor' from outer scope
    }
}

func closureDemo() {
    times2 := createMultiplier(2)
    times5 := createMultiplier(5)
    
    fmt.Println(times2(10))  // 20 (10 * 2)
    fmt.Println(times5(10))  // 50 (10 * 5)
}
````

#### Advanced Closure with State

````go
func createBankAccount(initialBalance float64) (func(float64), func(float64) bool, func() float64) {
    balance := initialBalance
    
    deposit := func(amount float64) {
        balance += amount
        fmt.Printf("Deposited $%.2f, Balance: $%.2f\n", amount, balance)
    }
    
    withdraw := func(amount float64) bool {
        if balance >= amount {
            balance -= amount
            fmt.Printf("Withdrew $%.2f, Balance: $%.2f\n", amount, balance)
            return true
        }
        fmt.Printf("Insufficient funds. Balance: $%.2f\n", balance)
        return false
    }
    
    getBalance := func() float64 {
        return balance
    }
    
    return deposit, withdraw, getBalance
}

func bankAccountExample() {
    deposit, withdraw, getBalance := createBankAccount(100.0)
    
    deposit(50.0)     // Deposited $50.00, Balance: $150.00
    withdraw(30.0)    // Withdrew $30.00, Balance: $120.00
    withdraw(200.0)   // Insufficient funds. Balance: $120.00
    
    fmt.Printf("Final balance: $%.2f\n", getBalance())  // $120.00
}
````

#### Closure Memory Model

```
Closure Memory Structure:
┌─────────────────────────────────────┐
│           HEAP MEMORY               │
├─────────────────────────────────────┤
│ ┌─────────────────────────────────┐ │
│ │        Closure Environment      │ │
│ │ Captured Variables:             │ │
│ │   factor = 2                    │ │ ◄── Variable from outer scope
│ │   balance = 120.0               │ │ ◄── Shared state
│ │ Function References:            │ │
│ │   times2: [function pointer]    │ │ ◄── Function with captured data
│ │   deposit: [function pointer]   │ │
│ │   withdraw: [function pointer]  │ │
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

### 🔹 Q3: What happens if you try to use a function expression before it is defined?

**A: Compilation Error Analysis**

#### The Error Scenario

````go
package main

import "fmt"

func main() {
    // ❌ This causes compilation error
    result := compute(5, 10)  // Error: undefined: compute
    
    compute := func(a, b int) int {
        return a + b
    }
    
    fmt.Println(result)
}
````

#### Compiler Error Details

| Error Component | Description | Technical Detail |
|----------------|-------------|------------------|
| **Error Message** | `undefined: compute` | Variable not in symbol table |
| **Error Phase** | Compile-time | Before program execution |
| **Root Cause** | Forward reference | Using undeclared identifier |
| **Go Behavior** | No implicit declarations | Explicit declaration required |

#### Variable Lifecycle in Go

```
Go Variable Resolution Process:
┌─────────────────────────────────────┐
│ Phase 1: Parsing                    │
├─────────────────────────────────────┤
│ • Scanner reads tokens              │
│ • Parser builds AST                 │
│ • Identifies all identifiers        │
├─────────────────────────────────────┤
│ Phase 2: Type Checking              │
├─────────────────────────────────────┤
│ • Resolves all variable references  │
│ • Checks declaration before use     │
│ • ERROR: 'compute' not declared     │ ◄── Compilation fails here
├─────────────────────────────────────┤
│ Phase 3: Code Generation (Skipped)  │
├─────────────────────────────────────┤
│ • Would generate machine code       │
│ • Never reached due to error        │
└─────────────────────────────────────┘
```

#### Comparison with Other Languages

````go
// JavaScript (hoisting behavior)
console.log(myFunc(5, 3));  // Works due to hoisting

function myFunc(a, b) {
    return a + b;
}

// Go (no hoisting)
func main() {
    result := myFunc(5, 3)  // ❌ Error: undefined: myFunc
    
    myFunc := func(a, b int) int {
        return a + b
    }
}
````

### 🔹 Q4: How would you modify the given code to make the function expression return the sum instead of printing it?

**A: Return Value Modification Patterns**

#### Original vs Modified Code

````go
// Original code (prints result)
func originalVersion() {
    add := func(a int, b int) {
        res := (a + b)
        fmt.Println(res)  // Prints but doesn't return
    }
    add(2, 4)  // Prints: 6
}

// Modified code (returns result)
func modifiedVersion() {
    add := func(a int, b int) int {
        res := (a + b)
        return res  // Returns the value
    }
    
    result := add(2, 4)  // Capture returned value
    fmt.Println(result)  // Print the captured result
}
````

#### Alternative Implementations

````go
// Approach 1: Direct return
func directReturn() {
    add := func(a, b int) int {
        return a + b  // Simplified version
    }
    
    fmt.Printf("Result: %d\n", add(2, 4))
}

// Approach 2: Multiple operations
func multipleOperations() {
    calculate := func(a, b int) (sum int, product int, average float64) {
        sum = a + b
        product = a * b
        average = float64(sum) / 2.0
        return  // Named return values
    }
    
    s, p, avg := calculate(2, 4)
    fmt.Printf("Sum: %d, Product: %d, Average: %.2f\n", s, p, avg)
}

// Approach 3: With error handling
func withErrorHandling() {
    safeDivide := func(a, b int) (int, error) {
        if b == 0 {
            return 0, fmt.Errorf("division by zero")
        }
        return a / b, nil
    }
    
    result, err := safeDivide(8, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Division result: %d\n", result)
}
````

### 🔹 Q5: Can you have recursive function expressions in Go? If so, how?

**A: Recursive Function Expression Patterns**

#### Method 1: Variable Declaration and Assignment

````go
func recursiveFactorial() {
    var factorial func(int) int  // Declare variable first
    
    factorial = func(n int) int {  // Then assign function
        if n <= 1 {
            return 1
        }
        return n * factorial(n-1)  // Recursive call
    }
    
    result := factorial(5)
    fmt.Printf("5! = %d\n", result)  // Output: 5! = 120
}
````

#### Method 2: Using Type Definition

````go
type RecursiveFunc func(int) int

func typeBasedRecursion() {
    var fibonacci RecursiveFunc
    
    fibonacci = func(n int) int {
        if n <= 1 {
            return n
        }
        return fibonacci(n-1) + fibonacci(n-2)
    }
    
    for i := 0; i < 10; i++ {
        fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
    }
}
````

#### Method 3: Closure-Based Recursion

````go
func createRecursiveCounter() func() int {
    count := 0
    
    var counter func() int
    counter = func() int {
        count++
        fmt.Printf("Count: %d\n", count)
        
        if count < 5 {
            return counter()  // Recursive call
        }
        
        return count
    }
    
    return counter
}

func closureRecursionExample() {
    recursiveCounter := createRecursiveCounter()
    finalCount := recursiveCounter()
    fmt.Printf("Final count: %d\n", finalCount)
}
````

#### Why Two-Step Declaration is Needed

```
Recursive Function Expression Memory Model:

Step 1: Variable Declaration
┌─────────────────────────────────────┐
│           STACK MEMORY              │
├─────────────────────────────────────┤
│ factorial: nil                      │ ◄── Variable exists but uninitialized
└─────────────────────────────────────┘

Step 2: Function Assignment
┌─────────────────────────────────────┐
│           HEAP MEMORY               │
├─────────────────────────────────────┤
│ Function Object: [recursive code]   │ ◄── Function can reference 'factorial'
├─────────────────────────────────────┤
│           STACK MEMORY              │
├─────────────────────────────────────┤
│ factorial: [pointer to heap]        │ ◄── Now points to function
└─────────────────────────────────────┘
```

---

## 🚀 Advanced Function Expression Patterns

### 🔹 Function Expression with Method Chaining

````go
type Calculator struct {
    value float64
}

func NewCalculator(initial float64) *Calculator {
    return &Calculator{value: initial}
}

func (c *Calculator) Add(op func(float64) float64) *Calculator {
    c.value = op(c.value)
    return c
}

func chainedCalculations() {
    result := NewCalculator(10.0).
        Add(func(x float64) float64 { return x + 5 }).    // Add 5
        Add(func(x float64) float64 { return x * 2 }).    // Multiply by 2
        Add(func(x float64) float64 { return x - 3 })     // Subtract 3
    
    fmt.Printf("Chained result: %.2f\n", result.value)  // (10 + 5) * 2 - 3 = 27
}
````

### 🔹 Function Expression Factory

````go
func createValidators() map[string]func(string) bool {
    return map[string]func(string) bool{
        "email": func(input string) bool {
            return strings.Contains(input, "@") && strings.Contains(input, ".")
        },
        "phone": func(input string) bool {
            return len(input) >= 10 && len(input) <= 15
        },
        "password": func(input string) bool {
            return len(input) >= 8
        },
    }
}

func validatorExample() {
    validators := createValidators()
    
    testData := map[string]string{
        "email":    "user@example.com",
        "phone":    "1234567890",
        "password": "securepass123",
    }
    
    for field, value := range testData {
        if validator, exists := validators[field]; exists {
            isValid := validator(value)
            fmt.Printf("%s '%s' is valid: %t\n", field, value, isValid)
        }
    }
}
````

---

## 🚀 Performance Considerations and Best Practices

### 🔹 Memory Efficiency Guidelines

| Pattern | Memory Impact | Performance | Recommendation |
|---------|---------------|-------------|----------------|
| **Simple Function Expression** | Heap allocation | Slight overhead | Use for flexibility |
| **Closure with Captured Variables** | Additional heap allocation | Memory retention | Use judiciously |
| **Recursive Function Expression** | Stack growth per call | Call overhead | Consider iterative alternatives |
| **Function Expression in Loop** | Multiple allocations | GC pressure | Extract to variable |

### 🔹 Best Practices

````go
// ✅ GOOD: Extract function expressions from loops
func goodLoopUsage() {
    transform := func(x int) int { return x * 2 }  // Define once
    
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        result := transform(num)  // Reuse function
        fmt.Println(result)
    }
}

// ❌ AVOID: Creating function expressions in loops
func badLoopUsage() {
    numbers := []int{1, 2, 3, 4, 5}
    for _, num := range numbers {
        transform := func(x int) int { return x * 2 }  // Created repeatedly
        result := transform(num)
        fmt.Println(result)
    }
}

// ✅ GOOD: Use function expressions for configuration
func configurationPattern() {
    processors := []func(string) string{
        func(s string) string { return strings.ToUpper(s) },
        func(s string) string { return strings.TrimSpace(s) },
        func(s string) string { return strings.ReplaceAll(s, " ", "_") },
    }
    
    input := "  hello world  "
    for _, processor := range processors {
        input = processor(input)
    }
    
    fmt.Printf("Processed: '%s'\n", input)  // "HELLO_WORLD"
}
````

## 🚀 Summary and Key Takeaways

| Concept | Symbol | Core Feature | Best Use Case |
|---------|--------|--------------|---------------|
| **Function Expression** | 🔧 | Anonymous function in variable | Dynamic behavior |
| **Declaration Order** | 📋 | Must declare before use | Compile-time safety |
| **Closure Capability** | 🔒 | Captures outer variables | State preservation |
| **First-Class Functions** | 🥇 | Functions as values | Functional programming |

### 🔹 When to Use Function Expressions

1. **Dynamic function assignment** - Different functions based on conditions
2. **Callback implementations** - Event handlers and function parameters
3. **Closure creation** - State preservation across function calls
4. **One-time complex logic** - Encapsulated operations
5. **Functional programming patterns** - Map, filter, reduce operations

### 🔹 Common Pitfalls to Avoid

| Pitfall | Problem | Solution |
|---------|---------|----------|
| **Use Before Declaration** | Compilation error | Always declare before use |
| **Function in Loop** | Memory waste | Extract to variable |
| **Deep Closures** | Memory leaks | Be mindful of captured variables |
| **Complex Logic** | Poor readability | Extract to named functions |

Remember: Function expressions provide powerful flexibility in Go, allowing you to treat functions as first-class citizens. Use them to create dynamic, configurable code while being mindful of declaration order and memory implications!
