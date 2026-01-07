# 📘 Golang 023 — Parameters, Arguments, First-Order & Higher-Order Functions

---

## PART 1️⃣ — Parameter vs Argument (Simple but Deadly in Interviews)

### 🔹 The Rule (One Line)

> **Argument is what you pass. Parameter is what you receive.**

---

### Example

```go
func add(a int, b int) {
    c := a + b
    fmt.Println(c)
}

func main() {
    add(2, 5)
}
```

### Breakdown

| Term           | What it is | Where               |
| -------------- | ---------- | ------------------- |
| **Arguments**  | `2`, `5`   | Function call       |
| **Parameters** | `a`, `b`   | Function definition |

### Flow (Mental Model)

```
Arguments  ──▶  Parameters
   2,5     ──▶   a,b
```

📌 **Memory trick (from the lecture):**

* **A = Argument = Before**
* **P = Parameter = After**

You **pass first**, you **receive later**.

---

### Why this matters?

* Team communication
* Clean technical discussion
* Very common **interview trap**

---

## PART 2️⃣ — What Is a First-Order Function?

### 🔹 Definition

> A **first-order function** works only with **simple data**
> (numbers, strings, booleans, structs, etc.)

---

### Examples (ALL First-Order)

* Named functions
* Anonymous functions
* IIFE (Immediately Invoked)
* Function expressions

```go
func add(a int, b int) int {
    return a + b
}
```

Why first-order?

* Parameters are **data**
* Returns **data**
* No function involved as input/output

📌 **Most functions you’ve written so far are first-order functions**

---

## PART 3️⃣ — Where “Order” Comes From (Big Picture)

This concept comes from:

```
Discrete Mathematics
   ↓
Logic (First-Order / Higher-Order)
   ↓
Functional Programming
   ↓
Go (inspired ideas)
```

You don’t need to *master math* — just understand the **idea**.

---

## PART 4️⃣ — Higher-Order Function (CORE CONCEPT 🔥)

### 🔹 Definition (Interview-Perfect)

A function is **higher-order** if **ANY ONE** of these is true:

1️⃣ Takes a function as a **parameter**
2️⃣ Returns a function
3️⃣ Does both

---

## PART 5️⃣ — Higher-Order Function (Function as Parameter)

### Example

```go
func processOperation(a int, b int, op func(int, int)) {
    op(a, b)
}

func add(x int, y int) {
    fmt.Println(x + y)
}

func main() {
    processOperation(2, 5, add)
}
```

### What’s happening?

| Role     | Value                  |
| -------- | ---------------------- |
| `a`, `b` | normal data            |
| `op`     | **function parameter** |
| `add`    | **callback function**  |

📌 **processOperation is a Higher-Order Function**

---

### Execution Flow

```
main()
 └─▶ processOperation(2,5,add)
      └─▶ add(2,5)
           └─▶ prints 7
```

---

## PART 6️⃣ — Callback Function (Very Important Term)

### 🔹 Definition

> A **callback function** is a function **passed as an argument** to another function.

In this example:

```go
add
```

is the **callback function**.

📌 Callback = “Call me back later”

---

## PART 7️⃣ — Higher-Order Function (Function as Return)

### Example

```go
func getAdder() func(int, int) {
    return func(a int, b int) {
        fmt.Println(a + b)
    }
}

func main() {
    sum := getAdder()
    sum(4, 3)
}
```

### What’s happening?

1. `getAdder()` returns a function
2. That function is assigned to `sum`
3. `sum(4,3)` executes it

📌 **Returning a function = Higher-Order Function**

---

## PART 8️⃣ — Why This Works: First-Class Functions

### 🔹 First-Class Citizen (General Rule)

A value is **first-class** if it can:

* Be stored in a variable
* Be passed as a parameter
* Be returned from a function

---

### In Go, these are first-class:

* int
* float
* string
* bool
* **function** ✅

```go
add := func(a int, b int) {
    fmt.Println(a + b)
}
```

📌 **Functions behave like data in Go**

---

## PART 9️⃣ — First-Class Function vs Higher-Order Function

⚠️ These are related but **not identical**

| Term                      | Meaning                            |
| ------------------------- | ---------------------------------- |
| **First-Class Function**  | Function can be treated like data  |
| **Higher-Order Function** | Function that uses other functions |

👉 **Higher-order functions exist because functions are first-class**

---

## PART 🔟 — Final Comparison Table (SAVE THIS)

| Concept               | Definition                       |
| --------------------- | -------------------------------- |
| Argument              | Value passed into a function     |
| Parameter             | Variable that receives the value |
| First-Order Function  | Works only with data             |
| Higher-Order Function | Takes/returns a function         |
| Callback Function     | Function passed as argument      |
| First-Class Function  | Function treated like data       |

---

## PART 1️⃣1️⃣ — Interview-Ready Answers 🎯

### ❓ Parameter vs Argument?

**Answer:**
Argument is passed, parameter receives.

---

### ❓ What is a higher-order function?

**Answer:**
A function that takes or returns another function.

---

### ❓ What is a callback function?

**Answer:**
A function passed to another function as an argument.

---

### ❓ Why are functions first-class in Go?

**Answer:**
Because they can be assigned, passed, and returned like variables.

---

## PART 1️⃣2️⃣ — One-Line Mental Model 🧠

> **First-order functions work with data.
> Higher-order functions work with behavior.**

---

## 🔥 Final Takeaway

This lesson is not about syntax.
It’s about **how Go lets you think in layers**:

* Data → Functions → Functions using functions
* Simple → Powerful → Expressive

If you truly understand this lesson,
**interview questions from this area become easy.**
