# 📘 Golang 021 — Anonymous Function & IIFE (Deep Explanation)

---

## 1️⃣ What Is an Anonymous Function?

### Meaning of *Anonymous*

> Anonymous = **no name**

So:

> **An anonymous function is a function without a name**

---

### Named (Standard) Function

```go
func add(a int, b int) {
    fmt.Println(a + b)
}
```

✔ Has a name: `add`
✔ Stored in memory
✔ Can be called later

This is called:

> **Named function / Standard function**

---

### Anonymous Function

```go
func(a int, b int) {
    fmt.Println(a + b)
}
```

❌ No name
✔ Has parameters
✔ Has a body
✔ Valid function

👉 This is an **anonymous function**

---

## 2️⃣ Why Anonymous Functions Exist (WHY)

Anonymous functions are used when:

* You need a function **once**
* You don’t want to pollute global scope
* You want **inline logic**
* You want **temporary behavior**

In short:

> **Use it when the function does not deserve a permanent name**

---

## 3️⃣ Important Rule in Go (Very Important)

🚫 **You cannot just write an anonymous function and leave it**

❌ This is invalid in Go:

```go
func(a int, b int) {
    fmt.Println(a + b)
}
```

Why?

Because:

* The function has **no name**
* Not stored anywhere
* Not called
* Not used

Go compiler says:

> “Why does this exist?”

---

## 4️⃣ Two Ways to Use Anonymous Functions

### ✅ Way 1: Assign to a Variable

```go
add := func(a int, b int) {
    fmt.Println(a + b)
}

add(5, 7)
```

Here:

* Function has no name
* But variable `add` **points to it**
* So it becomes callable

📌 This makes functions **first-class values**

---

### ✅ Way 2: Call Immediately (IIFE) ⭐

This is today’s main topic 👇

---

## 5️⃣ What Is IIFE?

### Full Form

> **Immediately Invoked Function Expression**

Short form:

> **IIFE** (pronounced *“iffy”*)

---

### Basic Structure

```go
func(a int, b int) {
    fmt.Println(a + b)
}(5, 7)
```

📌 Explanation:

* Function is **defined**
* Immediately **invoked**
* No name
* No storage
* Executes once

---

## 6️⃣ Why IIFE Is Needed

Think carefully 👇

Anonymous function:

* Has no name
* Has no reference
* Cannot be called later

So the only way to execute it is:

> **Call it immediately at the place of definition**

That’s why IIFE exists.

---

## 7️⃣ Step-by-Step Execution (Mental Simulation 🧠)

### Code

```go
func(a int, b int) {
    c := a + b
    fmt.Println(c)
}(5, 7)
```

### Runtime Thinking

1️⃣ Go sees a function definition
2️⃣ Immediately sees `(...)`
3️⃣ Invokes the function
4️⃣ Passes `5 → a`, `7 → b`
5️⃣ Executes body
6️⃣ Prints `12`
7️⃣ Function disappears

✔ No name
✔ No memory retention
✔ One-time execution

---

## 8️⃣ Why This Is Called an “Expression”

In programming:

> **An expression is anything that can be evaluated**

Examples:

```go
a := 10              // assignment expression
a + b               // arithmetic expression
if a > 0 { }         // conditional expression
add(2, 3)            // function invocation expression
```

So this:

```go
func(a int, b int) {
    fmt.Println(a + b)
}(5, 7)
```

✔ Is a **function expression**
✔ That is **immediately invoked**

Hence the name:

> **Immediately Invoked Function Expression**

---

## 9️⃣ Relation with `init()` (Important Connection)

From previous class:

```go
func init() {
    fmt.Println("I run first")
}
```

Difference:

| `init()`               | IIFE               |
| ---------------------- | ------------------ |
| Auto-called by runtime | Manually invoked   |
| Runs before `main()`   | Runs where written |
| No control             | Full control       |
| Package-level          | Any scope          |

---

## 🔟 Anonymous Function vs Named Function

| Feature          | Named Function | Anonymous Function  |
| ---------------- | -------------- | ------------------- |
| Has name         | ✅              | ❌                   |
| Stored in memory | ✅              | ❌ (unless assigned) |
| Reusable         | ✅              | ❌                   |
| One-time logic   | ❌              | ✅                   |
| Cleaner scope    | ❌              | ✅                   |

---

## 1️⃣1️⃣ Interview-Level Definitions (Memorize This)

### Anonymous Function

> A function without a name that can be assigned to a variable or invoked immediately.

### IIFE

> A function expression that is defined and executed immediately without being stored.

---

## 1️⃣2️⃣ Common Interview Questions ⚠️

❓ Can Go have anonymous functions?
➡ ✅ Yes

❓ Can Go have IIFE like JavaScript?
➡ ✅ Yes

❓ Why can’t anonymous functions exist alone in Go?
➡ Because they must be **used or invoked**

❓ Is IIFE stored in memory?
➡ ❌ No

❓ Can IIFE take parameters?
➡ ✅ Yes

---

## 1️⃣3️⃣ When You SHOULD Use IIFE

Use IIFE when:

* You want **temporary logic**
* You want **local execution**
* You want to avoid namespace pollution
* You want scoped initialization logic

Example:

```go
func() {
    temp := loadConfig()
    fmt.Println(temp)
}()
```

---

## 1️⃣4️⃣ Final Mental Model 🧠

Think of IIFE like:

```
Create → Execute → Destroy
```

No leftovers. No memory baggage.

---

## 1️⃣5️⃣ Final Takeaways 🔥

* Anonymous function = **no name**
* IIFE = **define + invoke immediately**
* Functions are **expressions** in Go
* Used heavily in interviews & advanced Go
* Clean, safe, powerful
