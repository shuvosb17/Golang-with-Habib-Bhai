# 📘 Golang 022 — Function Expression (Assigning Functions to Variables)

---

## 1️⃣ What Is a Function Expression?

### Simple Definition

> **A function expression is when you assign a function to a variable.**

Just like:

```go
x := 10
```

You can do:

```go
add := func(a int, b int) {
    fmt.Println(a + b)
}
```

👉 Here:

* `add` is a **variable**
* The value stored inside it is a **function**

That’s why it’s called a **function expression**.

---

## 2️⃣ Why Is It Called an “Expression”?

In programming:

> **An expression is anything that produces a value**

Examples:

```go
x := 10                  // value expression
y := x + 5               // arithmetic expression
add := func(...) { }     // function expression
```

So this line:

```go
add := func(a int, b int) { ... }
```

✔ Assigns a **value**
✔ That value happens to be a **function**

Hence:

> **Function Expression**

---

## 3️⃣ Anonymous Function + Variable = Function Expression

### Anonymous Function Alone ❌

```go
func(a int, b int) {
    fmt.Println(a + b)
}
```

🚫 Invalid in Go
Because:

* No name
* Not stored
* Not invoked

---

### Assign It to a Variable ✅

```go
add := func(a int, b int) {
    c := a + b
    fmt.Println(c)
}
```

✔ Now valid
✔ Function has a **reference**
✔ Callable

---

## 4️⃣ Calling a Function Expression

Once assigned:

```go
add(2, 3)
```

Execution:

* `2 → a`
* `3 → b`
* Output: `5`

---

## 5️⃣ Memory & Mental Model 🧠

Think like this:

```
Variable  ─────▶  Function Body
   add     ─────▶  func(a, b) { ... }
```

So:

* `add` behaves **like a function**
* But technically it’s a **variable**

This is why Go treats functions as **first-class citizens**.

---

## 6️⃣ VERY IMPORTANT RULE (Interview Gold ⚠️)

### ❌ This Will NOT Work

```go
add(2, 3)

add := func(a int, b int) {
    fmt.Println(a + b)
}
```

### Why?

Because:

> **Go executes top to bottom inside a scope**

At the moment you call `add(2, 3)`:

* `add` **does not exist yet**
* Compiler says: **undefined**

---

## 7️⃣ Function Expression Is NOT Hoisted ❌

### Named Functions (Hoisted-like behavior)

```go
add(2, 3)

func add(a int, b int) {
    fmt.Println(a + b)
}
```

✔ Works
Because named functions are **known during compile time**

---

### Function Expressions (NOT hoisted)

```go
add(2, 3) // ❌ ERROR

add := func(a int, b int) {
    fmt.Println(a + b)
}
```

❌ Error: `undefined: add`

📌 **Key Difference**

| Feature                    | Named Function | Function Expression |
| -------------------------- | -------------- | ------------------- |
| Hoisted                    | ✅ Yes          | ❌ No                |
| Stored at compile time     | ✅              | ❌                   |
| Depends on execution order | ❌              | ✅                   |

---

## 8️⃣ Scope Matters A LOT 🔥

### Local Scope Example

```go
func main() {
    add := func(a int, b int) {
        fmt.Println(a + b)
    }

    add(4, 5)
}
```

✔ Works

---

### Calling Before Definition (Local Scope)

```go
func main() {
    add(4, 5) // ❌ undefined

    add := func(a int, b int) {
        fmt.Println(a + b)
    }
}
```

❌ Error

📌 Reason:

> Local variables (including function expressions) **exist only after execution reaches their declaration**

---

## 9️⃣ Global Scope Function Expression (Allowed)

Yes — **this is valid in Go** 👇

```go
var add = func(a int, b int) {
    fmt.Println(a + b)
}

func main() {
    add(2, 3)
}
```

✔ Works
✔ Because global scope is initialized **before `main()`**

---

## 🔟 Shadowing with Function Expressions ⚠️

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
* Inside `main`, local version wins

📌 **Same shadowing rules as variables**

---

## 1️⃣1️⃣ Nested Function Calls (Simulation Insight)

Functions can call:

* Global functions
* Other function expressions
* Functions defined earlier in scope

But:

> **Function expressions must exist BEFORE invocation**

---

## 1️⃣2️⃣ The Golden Rule 🧠 (Memorize This)

> 🔥 **A function expression behaves exactly like a variable.**

So ask yourself:

> “Would this work if `add` was a normal variable?”

If NO → it will fail
If YES → it will work

---

## 1️⃣3️⃣ Interview Questions & Answers 🎯

### Q: What is a function expression?

**A:** Assigning a function to a variable.

---

### Q: Are function expressions hoisted in Go?

**A:** ❌ No.

---

### Q: Can a function expression be global?

**A:** ✅ Yes.

---

### Q: Why does calling before definition fail?

**A:** Because the variable doesn’t exist yet.

---

### Q: Can function expressions be shadowed?

**A:** ✅ Yes, like normal variables.

---

## 1️⃣4️⃣ When Should You Use Function Expressions?

Use when:

* Passing functions as arguments
* Returning functions
* Closures
* Callbacks
* Temporary behavior
* Cleaner scoped logic

---

## 1️⃣5️⃣ Final Summary 🔥

* Function expression = **function stored in a variable**
* Anonymous function + assignment = function expression
* NOT hoisted
* Order matters
* Scope rules apply
* Same behavior as normal variables
