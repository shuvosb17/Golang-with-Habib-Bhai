# 📘 Golang Lecture 05 — Functions with Return Values & Types

---

## 1️⃣ The Problem with Printing Inside Functions

### Previous approach

```go
func add(a int, b int) {
    sum := a + b
    fmt.Println(sum)
}
```

### Why this is limiting

* Function **decides output format**
* You **can’t reuse** the result
* You **can’t store** the value
* You **can’t pass it elsewhere**

📌 In real software:

> Functions should **return data**, not just print it.

---

## 2️⃣ New Goal: Return the Result

We want:

* Function computes something
* Caller decides **what to do with the result**

This is **clean design**.

---

## 3️⃣ Function Anatomy (With Return Type)

### General structure

```go
func functionName(inputs) outputType {
    // logic
    return value
}
```

### Key idea

* Inputs → parameters
* Output → return type
* `return` sends value back to caller

---

## 4️⃣ Single Return Value Example (Addition)

### Function definition

```go
func add(numberOne int, numberTwo int) int {
    sum := numberOne + numberTwo
    return sum
}
```

### Read it in English

> “This function takes two integers and **returns one integer**.”

---

## 5️⃣ Why Return Type Is Written After Parameters

```go
func add(a int, b int) int
```

Go’s philosophy:

* Inputs first
* Output second
* Readable left → right

📌 This also allows **multiple return values** (you’ll see soon).

---

## 6️⃣ Calling a Function That Returns a Value

```go
result := add(a, b)
fmt.Println(result)
```

### What really happens

```
add(a, b) → returns 30
result := 30
```

📌 **Important mental model**

> The function call is replaced by the returned value.

---

## 7️⃣ Execution Flow (Very Important)

Let’s trace it step by step.

### Code

```go
a := 10
b := 20
sum := add(a, b)
```

### Memory & execution

```
RAM (main)
+--------+
| a = 10 |
| b = 20 |
+--------+

Call add(a, b)
```

---

## 8️⃣ Inside the Function (Stack Frame)

When `add()` is called:

```
add() memory
+------------------+
| numberOne = 10   |
| numberTwo = 20   |
| sum = 30         |
+------------------+
```

Then:

```go
return sum
```

➡️ sends `30` back to caller
➡️ function memory is destroyed

---

## 9️⃣ After Return (Back to Caller)

```
sum := add(a, b)
```

Becomes:

```
sum := 30
```

📌 This is why:

* You can store return values
* You can pass them to other functions
* You can chain logic

---

## 🔟 Important Rule About `main()`

```go
func main() {
}
```

### Why no return type?

* `main` is **entry point**
* OS does not expect a value
* Program just **starts and ends**

📌 `main` is also a function — just special.

---

## 1️⃣1️⃣ Multiple Return Values (Go Superpower)

Go allows:

```go
func getNumbers(a int, b int) (int, int)
```

Meaning:

> “This function returns **two integers**.”

📌 Multiple returns are **first-class citizens** in Go.

---

## 1️⃣2️⃣ Example: Sum & Multiplication Together

### Function definition

```go
func getNumbers(numberOne int, numberTwo int) (int, int) {
    sum := numberOne + numberTwo
    mul := numberOne * numberTwo
    return sum, mul
}
```

### Why parentheses?

* Required for **multiple return values**
* Makes output grouping explicit

---

## 1️⃣3️⃣ Calling a Function with Multiple Returns

```go
p, q := getNumbers(a, b)
```

### Meaning

```
p = sum
q = multiplication
```

📌 Order matters!

---

## 1️⃣4️⃣ Execution Flow (Multiple Returns)

### Call

```go
p, q := getNumbers(10, 20)
```

### Inside function

```
sum = 30
mul = 200
return 30, 200
```

### After return

```
p = 30
q = 200
```

---

## 1️⃣5️⃣ Printing the Results

```go
fmt.Println(p)
fmt.Println(q)
```

### Output

```
30
200
```

✔️ Clean
✔️ Reusable
✔️ Professional Go style

---

## 1️⃣6️⃣ Key Rules You Must Lock In

### 🔒 Return type must match

```go
func f() int {
    return "hello" // ❌
}
```

### 🔁 Return order matters

```go
return a, b  // maps left to right
```

### 🧠 Every return ends the function

* Code after `return` is unreachable

---

## 1️⃣7️⃣ Why Go Loves Multiple Returns

In real Go:

* Functions often return `(value, error)`
* This avoids exceptions
* Makes error handling explicit

📌 You’ll see this **everywhere** in Go.

---

## 🧠 Ultimate Mental Model (Must Remember)

```
Function call
   ↓
Create memory
   ↓
Execute logic
   ↓
Return value(s)
   ↓
Destroy memory
   ↓
Caller continues
```

---

## 🔁 Final Recap (Revision Perfect)

* Functions can return values
* Return type is mandatory
* Returned value replaces function call
* Go supports multiple return values
* `main()` returns nothing
* Memory is freed after return

---

## 🚀 What You’re Ready For Next

Now you’re ready to learn:

* Returning `error`
* Named return values
* Functions as building blocks
* Real backend-style Go code

