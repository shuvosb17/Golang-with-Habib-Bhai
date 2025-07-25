# 🚀Lecture 009: **Understanding Conditional Statements in Go! (if...else & switch)**

---

## The Decision-Making Engine of Programming

In this class, we will explore how Go makes **decisions** using `if...else` and `switch` statements.

## 1️⃣ If...Else Statement - The Decision Maker 🏗️

An `if...else` statement **executes different blocks of code depending on a condition**.

| Component | Purpose | Syntax |
|-----------|---------|--------|
| **if** | Primary condition check | `if condition { }` |
| **else if** | Alternative condition | `else if condition { }` |
| **else** | Default fallback | `else { }` |

### 🔹 Basic If...Else Structure

````go
package main

import "fmt"

func main() {
    age := 18  // Test variable

    if age >= 18 {
        fmt.Println("You are eligible to be married")
    } else if age < 18 {
        fmt.Println("You are not eligible to be married, but you can love someone")
    } else {
        fmt.Println("You are just a teenager, not eligible to be married")
    }
}
````

### 🔹 Execution Flow Diagram

```
Condition Evaluation Flow:
┌─────────────────┐
│   age := 18     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ if age >= 18?   │ ◄── First condition checked
└────────┬────────┘
         │ YES
         ▼
┌─────────────────┐
│ Execute if block│ ◄── Print: "You are eligible..."
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Skip else blocks│ ◄── Remaining conditions ignored
└─────────────────┘
```

### 🔹 Common Mistake: Order Matters!

| Problem | Solution |
|---------|----------|
| **Broad condition first** | **Specific condition first** |

````go
// ❌ WRONG: Broad condition first
if age > 18 {          // 20 > 18 → true (stops here)
    fmt.Println("Ready!")
} else if age == 20 {  // Never reached!
    fmt.Println("Special case!")
}

// ✅ CORRECT: Specific condition first
if age == 20 {         // Check specific case first
    fmt.Println("Special case!")
} else if age > 18 {   // Then check broader condition
    fmt.Println("Ready!")
}
````

### 🔹 Multiple Conditions vs If-Else Chain

| Approach | When to Use | Result |
|----------|-------------|---------|
| **If-Else Chain** | One condition should execute | Only first true condition runs |
| **Separate If Statements** | Multiple conditions can be true | All true conditions run |

````go
// Multiple separate conditions (all can execute)
if age > 18 {
    fmt.Println("You are ready to go!")
}
if age == 20 {
    fmt.Println("You can prepare!")
}
if age < 25 {
    fmt.Println("You are young!")
}
````

### 🔹 Rollercoaster Analogy 🎢

```
Theme Park Entry Requirements:
┌─────────────────────────────┐
│        Height Check         │
├─────────────────────────────┤
│ if height >= 140cm          │ ◄── Can ride big rides
│   "Enjoy all attractions!"  │
│ else if height >= 120cm     │ ◄── Can ride some rides
│   "Some rides available"    │
│ else                        │ ◄── Kids area only
│   "Kids area only"          │
└─────────────────────────────┘
```

---

## 2️⃣ Logical Operators - Combining Conditions 🤔

Go supports **logical operators** to create complex conditions:

| Operator | Name | Description | Example |
|----------|------|-------------|---------|
| `&&` | AND | Both conditions must be true | `age >= 18 && hasLicense` |
| `||` | OR | At least one condition must be true | `isWeekend || isHoliday` |
| `!` | NOT | Negates/reverses a condition | `!isRaining` |

### 🔹 AND Operator (`&&`) Example

````go
package main

import "fmt"

func main() {
    age := 20
    sex := "male"

    if age >= 20 && sex == "male" {
        fmt.Println("You are ready to be married!")
    } else {
        fmt.Println("You are not ready yet.")
    }
}
````

### 🔹 Logical Operators Truth Table

| A | B | A && B | A \|\| B | !A |
|---|---|--------|----------|----| 
| true | true | true | true | false |
| true | false | false | true | false |
| false | true | false | true | true |
| false | false | false | false | true |

### 🔹 Real-world Logical Operations

| Scenario | Condition | Logic |
|----------|-----------|-------|
| **Driving License** 🚗 | `age >= 18 && passedTest` | Both required |
| **Weekend Fun** 🎉 | `isWeekend || isHoliday` | Either works |
| **Indoor Activity** 🏠 | `!isSunny` | Opposite condition |

### 🔹 Complex Condition Example

````go
// Bank loan approval system
if (age >= 21 && age <= 65) && (income >= 30000 || hasCollateral) && !hasBadCredit {
    fmt.Println("Loan approved!")
} else {
    fmt.Println("Loan denied.")
}
````

---

## 3️⃣ Switch Statement - The Multiple Choice Master 🔄

A `switch` statement **checks multiple values** efficiently and executes the matching case.

| Component | Purpose | Syntax |
|-----------|---------|--------|
| **switch** | Variable to test | `switch variable {` |
| **case** | Possible values | `case value:` |
| **default** | Fallback option | `default:` |

### 🔹 Basic Switch Structure

````go
package main

import "fmt"

func main() {
    a := 3

    switch a {
    case 1:
        fmt.Println("a is 1")
    case 2, 3: // Multiple cases together
        fmt.Println("a is either 2 or 3")
    case 4:
        fmt.Println("a is 4")
    default:
        fmt.Println("a is neither 1, 2, 3, nor 4")
    }
}
````

### 🔹 Switch vs If-Else Comparison

| Aspect | Switch | If-Else |
|--------|--------|---------|
| **Best for** | Checking one variable against multiple values | Complex conditions with different variables |
| **Readability** | Clean for many cases | Better for logical combinations |
| **Performance** | Optimized for value matching | General purpose |

### 🔹 Advanced Switch Features

````go
// Switch without expression (acts like if-else)
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18 && age < 65:
    fmt.Println("Adult")
default:
    fmt.Println("Senior")
}

// Switch with fallthrough
switch grade {
case "A":
    fmt.Println("Excellent!")
    fallthrough  // Continue to next case
case "B":
    fmt.Println("Good job!")
default:
    fmt.Println("Keep trying!")
}
````

### 🔹 Pizza Ordering System Analogy 🍕

```
Pizza Menu Selection:
┌─────────────────────────────┐
│    Customer Choice: 3       │
├─────────────────────────────┤
│ case 1: Cheese Pizza        │
│ case 2: Pepperoni Pizza     │
│ case 3: BBQ Pizza          │ ◄── Matches here!
│ case 4: Veggie Pizza        │
│ default: Invalid choice     │
└─────────────────────────────┘
           │
           ▼
    "One BBQ Pizza coming up!"
```

### 🔹 Real-world Switch Examples

| Use Case | Example |
|----------|---------|
| **Day of Week** | `case "Monday":` → "Back to work!" |
| **HTTP Status** | `case 200:` → "Success", `case 404:` → "Not Found" |
| **User Menu** | `case 1:` → "View Profile", `case 2:` → "Settings" |

---

## 🔄 Decision Flow Comparison

```
If-Else vs Switch Decision Flow:

If-Else Chain:                    Switch Statement:
┌─────────────┐                  ┌─────────────┐
│ Condition 1 │ ──true──►        │   Value 1   │ ──match──►
└─────┬───────┘                  └─────┬───────┘
      │false                           │no match
      ▼                                ▼
┌─────────────┐                  ┌─────────────┐
│ Condition 2 │ ──true──►        │   Value 2   │ ──match──►
└─────┬───────┘                  └─────┬───────┘
      │false                           │no match
      ▼                                ▼
┌─────────────┐                  ┌─────────────┐
│    Else     │                  │   Default   │
└─────────────┘                  └─────────────┘
```

## 🚀 Final Summary

| Statement Type | Symbol | Best Use Case | Key Feature |
|----------------|--------|---------------|-------------|
| **If-Else** | 🏗️ | Complex conditions with logical operators | Flexible condition evaluation |
| **Logical Operators** | 🤔 | Combining multiple boolean expressions | `&&`, `||`, `!` for complex logic |
| **Switch** | 🔄 | Testing one variable against multiple values | Clean, optimized value matching |

### 🔹 Decision-Making Best Practices

1. **Use if-else for complex conditions** - Multiple variables, ranges, logical operations
2. **Use switch for value matching** - Single variable against known values
3. **Order conditions properly** - Specific cases before general ones
4. **Consider readability** - Choose the structure that makes code clearer
5. **Use logical operators wisely** - Combine related conditions effectively

### 🔹 When to Use What?

| Scenario | Recommended Approach |
|----------|---------------------|
| **Age ranges** | If-else with `&&` operators |
| **Menu selections** | Switch statement |
| **Multiple flags** | If-else with `||` operators |
| **Single value check** | Switch or simple if |

Remember: Conditional statements are the brain of your program - they make decisions that determine how your code behaves in different situations!
