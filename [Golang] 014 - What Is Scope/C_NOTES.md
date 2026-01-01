If **scope** is unclear:

* variables will feel random
* errors will feel unfair
* functions will feel magical (and scary)

If **scope is clear**:

* Go becomes logical
* errors become predictable
* debugging becomes easy

# 📘 Golang Lecture 08 — What Is Scope (Deep & Practical)

---

## 1️⃣ What “Scope” Actually Means (Plain English)

> **Scope = Where a variable or function is visible and usable**

If you can **access** it → it’s **in scope**
If you **cannot** access it → it’s **out of scope**

That’s it.
No magic. No theory. Just **visibility**.

---

## 2️⃣ The Three Main Scopes in Go (Core Concept)

Go has **three important scopes** you must understand:

```
┌───────────────┐
│ Global Scope  │  ← declared outside all functions
│               │
│  ┌─────────┐  │
│  │ main()  │  │  ← function scope
│  │         │  │
│  │  ┌───┐  │  │
│  │  │if │  │  │  ← block scope
│  │  └───┘  │  │
│  └─────────┘  │
└───────────────┘
```

---

## 3️⃣ Global Scope (Outside All Functions)

### Example

```go
var a = 20
var b = 30

func add(x int, y int) {
    z := x + y
    fmt.Println(z)
}

func main() {
    add(a, b)
}
```

### What is global here?

* `a`
* `b`
* `add`
* `main`

📌 **Anything declared outside functions → global scope**

---

## 4️⃣ Key Rule: Global Scope Visibility

> **Global variables and functions are accessible everywhere**

* Inside `main`
* Inside `add`
* Inside any function

That’s why:

```go
add(a, b)
```

works perfectly.

---

## 5️⃣ Function Scope (Inside a Function)

### Inside `add()`

```go
func add(x int, y int) {
    z := x + y
}
```

### Variables here:

* `x`
* `y`
* `z`

📌 These **exist ONLY inside `add()`**

---

## 6️⃣ Why This Fails ❌

```go
func main() {
    fmt.Println(z) // ❌ error
}
```

### Why?

Because:

* `z` was created inside `add`
* `main` is outside that scope

👉 **Child scope variables do NOT go upward**

---

## 7️⃣ Scope Lookup Rule (VERY IMPORTANT)

When Go sees a variable name, it searches in this order:

```
1️⃣ Current function / block
2️⃣ Parent function
3️⃣ Global scope
4️⃣ If not found → ERROR
```

This rule explains **everything**.

---

## 8️⃣ Example: Why `x + a` Works Inside `add()`

```go
func add(x int, y int) {
    z := x + a
}
```

### How Go resolves this:

1. Is `a` inside `add()`? ❌
2. Is `a` a parameter? ❌
3. Is `a` global? ✅

✔️ Works.

---

## 9️⃣ Example: Why `q` Does NOT Work ❌

```go
func main() {
    p := 30
    q := 40
    add(p, q)
}

func add(x int, y int) {
    fmt.Println(q) // ❌ error
}
```

### Why?

* `q` exists inside `main`
* `add()` is a **separate function**
* `main` variables are **not global**

📌 **Functions do NOT see each other’s local variables**

---

## 🔟 Function Parameters Create New Scope

```go
add(p, q)
```

Inside `add()`:

```go
x = p
y = q
```

📌 `x` and `y` are **new variables**, not aliases.

They:

* get values
* live inside `add`
* die when `add` ends

---

## 1️⃣1️⃣ Lifetime vs Scope (Important Distinction)

| Concept  | Meaning                   |
| -------- | ------------------------- |
| Scope    | Where variable is visible |
| Lifetime | How long variable exists  |

Example:

* `z` exists **only while `add()` runs**
* After function ends → memory freed

---

## 1️⃣2️⃣ Why `z` Cannot Be Used in `main`

```go
func add(x int, y int) {
    z := x + y
}

func main() {
    add(10, 20)
    fmt.Println(z) // ❌
}
```

### Explanation

* `z` was created in `add`
* `add` finished execution
* `z` was destroyed

📌 **Dead variables have no scope**

---

## 1️⃣3️⃣ Why This Line Is INVALID ❌

```go
add(b, z)
```

### Reason

* `z` exists only inside `add`
* `main` never had `z`
* global scope never had `z`

📌 **You cannot pass what you don’t own**

---

## 1️⃣4️⃣ Visual Memory Model (Critical)

### During program start

```
Global Scope:
a = 20
b = 30
add()
main()
```

---

### During `main()`

```
Main Scope:
p = 30
q = 40
```

---

### During `add(p, q)`

```
Add Scope:
x = 30
y = 40
z = 70
```

After function ends:

```
Add Scope destroyed
```

---

## 1️⃣5️⃣ Why This Design Is GOOD

If Go allowed everything everywhere:

* Bugs would explode
* Code would be unmanageable
* No safety

Scope:

* protects variables
* prevents accidental misuse
* makes code predictable

---

## 1️⃣6️⃣ Interview-Level Definition (Memorize This)

> **Scope determines where a variable or function can be accessed in a program.**

If it can be accessed → **in scope**
If not → **out of scope**

---

## 1️⃣7️⃣ Common Interview Trick Question

```go
var a = 10

func test() {
    var a = 20
    fmt.Println(a)
}
```

### Output?

```
20
```

📌 **Nearest scope wins**
(Local overrides global)

---

## 1️⃣8️⃣ Final Mental Model (LOCK THIS IN 🔒)

```
Scope is like rooms in a building:

Global = lobby
Function = private room
Block = cupboard

You can see:
- inside your room
- outside to lobby

But NOT:
- into another room
```

---

## 🔁 Final Recap (Perfect for Revision)

* Scope = visibility
* Go searches: local → parent → global
* Functions have isolated scopes
* Variables die when scope ends
* Global ≠ function-local
* If Go says “undefined” → scope problem

---

## 🚀 What This Unlocks Next

Now you’re ready for:

* Closures
* Shadowing
* Packages
* Real debugging
* Confident interviews

