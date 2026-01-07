# 📘 Golang 022 — Function Expression (Assigning Functions to Variables)

---

## 1️⃣ What Is a Function Expression? (Core Idea)

### Definition (Very Important)

> **A function expression is when you assign a function to a variable.**

Example:

```go
add := func(a int, b int) {
    fmt.Println(a + b)
}
```

Here:

* `add` → variable
* value of `add` → a function

This is why it’s called a **function expression**.

---

## 2️⃣ Why Is It Called an “Expression”?

In programming:

> **An expression is anything that produces a value**

Examples:

```go
x := 10              // value expression
y := x + 5           // arithmetic expression
add := func(...) {}  // function expression
```

So:

* `10` is a value
* `func(...) {}` is ALSO a value
* That value happens to be executable

---

## 3️⃣ Anonymous Function → Error (If Not Used Properly)

### ❌ This Is INVALID in Go

```go
func(a int, b int) {
    fmt.Println(a + b)
}
```

Why?

* No name
* Not stored
* Not called

Go does **not** allow “floating” anonymous functions.

---

## 4️⃣ Fixing the Error: Assign Function to a Variable ✅

```go
add := func(a int, b int) {
    c := a + b
    fmt.Println(c)
}
```

Now:

* Function has a **reference**
* Stored in memory
* Callable

---

## 5️⃣ Calling a Function Expression

```go
add(2, 3)
```

Execution:

* `a = 2`
* `b = 3`
* `c = 5`
* Output → `5`

---

## 6️⃣ Mental Model: Function as a Value 🧠

Think like this:

```
add ─────▶ func(a, b) { ... }
```

So:

* `add` behaves like a function
* But technically it’s just a **variable holding a function**

👉 **Functions are first-class citizens in Go**

---

## 7️⃣ Execution Order Matters (THIS IS THE TRAP ⚠️)

### ❌ This Will FAIL

```go
add(2, 3)

add := func(a int, b int) {
    fmt.Println(a + b)
}
```

### Why?

Because:

* Code executes **top to bottom**
* At the time of calling `add(2, 3)`
* `add` does **not exist yet**

📌 Error: `undefined: add`

---

## 8️⃣ Named Function vs Function Expression (Critical Difference)

### ✅ Named Function (Works)

```go
add(2, 3)

func add(a int, b int) {
    fmt.Println(a + b)
}
```

### ❌ Function Expression (Fails)

```go
add(2, 3)

add := func(a int, b int) {
    fmt.Println(a + b)
}
```

### Why the difference?

| Feature                    | Named Function | Function Expression |
| -------------------------- | -------------- | ------------------- |
| Known at compile time      | ✅              | ❌                   |
| Hoisted-like behavior      | ✅              | ❌                   |
| Depends on execution order | ❌              | ✅                   |

---

## 9️⃣ Scope Rules Apply (Local vs Global)

### ✅ Global Function Expression (Works)

```go
var add = func(a int, b int) {
    fmt.Println(a + b)
}

func main() {
    add(2, 3)
}
```

Why?

* Global scope is initialized **before `main()`**
* Memory already allocated

---

### ❌ Local Scope Call Before Definition (Fails)

```go
func main() {
    add(2, 3) // ❌ undefined

    add := func(a int, b int) {
        fmt.Println(a + b)
    }
}
```

📌 **Golden Rule**:

> A function expression behaves EXACTLY like a variable.

---

## 🔟 Deep Execution Flow (RAM Simulation Simplified)

### Step-by-step:

1. Compiler reads the file
2. Global scope functions/variables are registered
3. `init()` (if any) runs
4. `main()` starts execution
5. Inside `main`, code runs line-by-line
6. Function expressions are created **only when execution reaches them**
7. Call works **only after creation**

---

## 1️⃣1️⃣ Function Expression + Shadowing ⚠️

```go
func add(a int, b int) {
    fmt.Println(a + b)
}

func main() {
    add := func(a int, b int) {
        fmt.Println(a * b)
    }

    add(2, 3)
}
```

### Output:

```
6
```

Why?

* Local `add` **shadows** global `add`
* Same rules as variable shadowing

---

## 1️⃣2️⃣ Calling From Another Function

### This WORKS:

```go
func add(a int, b int) {
    fmt.Println(a + b)
}

func sum() {
    add(2, 4)
}

func main() {
    sum()
}
```

### But this DOES NOT (local function expression):

```go
func sum() {
    add(2, 4) // ❌ undefined

    add := func(a int, b int) {
        fmt.Println(a + b)
    }
}
```

📌 Because:

* `add` doesn’t exist yet at call time

---

## 1️⃣3️⃣ The “Child Not Born Yet” Analogy 👶 (Very Accurate)

Calling a function expression before its definition is like:

> Trying to call your child before they’re born.

* Future existence doesn’t matter
* Present memory allocation does

---

## 1️⃣4️⃣ Key Interview Takeaways 🎯

### Q: What is a function expression?

**A:** Assigning a function to a variable.

---

### Q: Is a function expression hoisted in Go?

**A:** ❌ No.

---

### Q: Can function expressions be global?

**A:** ✅ Yes.

---

### Q: Why does calling before definition fail?

**A:** Because the variable hasn’t been created yet.

---

### Q: Are function expressions affected by scope?

**A:** ✅ Yes, fully.

---

## 1️⃣5️⃣ Final Golden Rules (Write These Down ✍️)

1. **Function expression = variable behavior**
2. **Order matters**
3. **Local scope ≠ global scope**
4. **Defined first → then call**
5. **Shadowing rules apply**
6. **Named functions ≠ function expressions**

---

## 🧠 One-Line Summary

> **If a function is stored in a variable, Go treats it like a variable—no magic, no hoisting, no forgiveness.**

