# 📘 Golang 019 — Function Types & Standard (Named) Functions

> 🎯 **Big Goal of This Class**
> Understand that in Go:

* Functions are **first-class citizens**
* Functions have **types**
* Functions can be **named, passed, stored, and reused**

This is where Go starts showing its **functional-programming influence**.

---

## 1️⃣ Why This Topic Matters (Industry Perspective)

From the instructor’s experience (200+ interviews 👀):

💡 **Interviewers don’t test syntax**
They test:

* How you **think**
* Whether you understand **abstraction**
* Whether you understand **functions as values**

And Go uses functions **everywhere**:

* HTTP handlers
* Middleware
* Goroutines
* Callbacks
* Dependency injection

If you don’t understand function types →
you’ll struggle in **real backend Go code**.

---

## 2️⃣ What Is a “Standard / Named Function”?

This is the **classic function** you already know.

### Example

```go
func add(a int, b int) int {
    return a + b
}
```

### Breakdown (Very Important)

```text
func        → function keyword
add         → function name
(a int, b int) → input parameters
int         → return type
{ ... }     → function body
```

📌 This is called a **named function** because:

* It has a **name**
* The name points to a **function value**

---

## 3️⃣ Key Idea: Functions Have TYPES 😮

In Go, **every function has a type**.

### The type of `add` is:

```go
func(int, int) int
```

That means:

* Takes **two ints**
* Returns **one int**

👉 This function signature **IS A TYPE**.

---

## 4️⃣ Why Go Cares About Function Types

Because Go treats functions like:

* `int`
* `string`
* `struct`

💥 **Functions can be assigned, passed, and returned**

This enables:

* Clean architecture
* Testability
* Loose coupling
* Middleware chains

---

## 5️⃣ Storing a Function in a Variable

### Example

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    var f func(int, int) int
    f = add

    result := f(10, 20)
    fmt.Println(result)
}
```

### What Happened?

```text
add        → function
f          → variable holding a function
f(10, 20)  → calling the function
```

📌 `f` does NOT store data
📌 `f` stores **behavior**

---

## 6️⃣ Mental Model (VERY IMPORTANT)

Think of a function like this:

```text
Function Name
     ↓
Memory Address
     ↓
Executable Code
```

So when you do:

```go
f = add
```

You are saying:

> “Point `f` to the same function that `add` points to”

---

## 7️⃣ Function Type Matching Rule (STRICT)

Go is **very strict**.

❌ This will NOT work:

```go
var f func(int) int
f = add
```

Because:

```text
Expected: func(int) int
Got:      func(int, int) int
```

✔ Function **signature must match EXACTLY**

---

## 8️⃣ Why This Design Is Powerful

Because later you can do things like:

```go
type Operation func(int, int) int

func calculate(a int, b int, op Operation) int {
    return op(a, b)
}
```

Then:

```go
calculate(10, 20, add)
```

🔥 This is **real backend Go design**

---

## 9️⃣ Named Functions vs Anonymous Functions (Preview)

### Named Function

```go
func add(a int, b int) int {
    return a + b
}
```

### Anonymous Function (no name)

```go
func(a int, b int) int {
    return a + b
}
```

📌 Both have the **same type**
📌 One has a name, one doesn’t

We’ll go deep into this next class 👀

---

## 🔁 Comparison Table

| Feature            | Named Function |
| ------------------ | -------------- |
| Has name           | ✅              |
| Has type           | ✅              |
| Reusable           | ✅              |
| Can be passed      | ✅              |
| Stored in variable | ✅              |
| First-class        | ✅              |

---

## 🧠 One-Line Interview Answer

> **In Go, functions are first-class values with explicit types, meaning they can be assigned to variables, passed as arguments, and returned from other functions.**

---

## 10️⃣ Why Go Feels Different Here

Go is inspired by:

* **Functional programming** (functions as values)
* **Structured programming**
* Minimal OOP (no inheritance)

So Go code looks simple, but the **design power is huge**.

---

## 🚀 What You’re Ready For Next

After this lecture, you are ready to learn:

* Anonymous functions
* Callback patterns
* Closures
* Middleware
* Goroutines + functions
* Real backend patterns
