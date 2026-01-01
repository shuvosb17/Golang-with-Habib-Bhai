# 📘 Golang 09 — Local Scope & Block Scope (Deep Explanation)

---

## 1️⃣ Quick Recap: Types of Scope in Go

In Go, **scope has three main levels**:

| Scope Type                 | Where it Exists                         |
| -------------------------- | --------------------------------------- |
| **Global Scope**           | Outside all functions                   |
| **Local (Function) Scope** | Inside a function                       |
| **Block Scope**            | Inside `{ }` like `if`, `switch`, `for` |

Today’s class focuses on **Local Scope & Block Scope**.

---

## 2️⃣ What Is a “Block”?

> **A block is any code wrapped inside `{ }` (curly braces)**

Examples of blocks:

```go
func main() { }      // function block
if condition { }    // if block
switch x { }        // switch block
for i := 0; i < n; i++ { } // loop block
```

📌 **Every block creates its own local scope**

---

## 3️⃣ Local Scope = Function Scope

### Example

```go
func main() {
    x := 18
}
```

* `x` is **local to `main()`**
* `x` is **NOT global**
* `x` is destroyed when `main()` ends

📌 **Variables declared inside a function live only inside that function**

---

## 4️⃣ Block Scope = Scope Inside `{ }` (Nested Scope)

Now look carefully 👀

```go
func main() {
    x := 18

    if x >= 18 {
        p := 10
        fmt.Println("I am matured boy")
        fmt.Println("I have", p, "girlfriends")
    }
}
```

### Scopes here:

```
Global
 └── main()
      └── if block
```

---

## 5️⃣ Variable Visibility Hierarchy (VERY IMPORTANT)

When Go tries to find a variable, it looks **step by step**:

```
1️⃣ Current block
2️⃣ Parent function
3️⃣ Global scope
4️⃣ Not found → ERROR
```

This rule explains **every scope behavior**.

---

## 6️⃣ Why `p` Works INSIDE the `if` Block ✅

```go
if x >= 18 {
    p := 10
    fmt.Println(p) // ✅ works
}
```

Why?

* `p` exists inside this `{ }`
* Code is still inside the same block

📌 **Block variables are visible only inside that block**

---

## 7️⃣ Why `p` FAILS Outside the Block ❌

```go
if x >= 18 {
    p := 10
}

fmt.Println(p) // ❌ error
```

### What happens internally:

1. `if` block ends
2. Block memory is destroyed
3. `p` no longer exists

📌 **When a block ends, its variables DIE**

---

## 8️⃣ Mental Memory Model (CRITICAL)

### During execution:

```
Main Scope:
x = 18

IF Block Scope:
p = 10
```

### After `}` of `if`:

```
Main Scope:
x = 18

IF Block Scope:
❌ destroyed
```

👉 `p` is gone forever.

---

## 9️⃣ Why This Is Called “Block Scope”

Because:

* Every `{ }` creates a **temporary mini-scope**
* Variables inside belong **only to that block**

That’s why Go calls it **block scope**.

---

## 🔟 Function Scope vs Block Scope (Side-by-Side)

| Feature    | Function Scope  | Block Scope       |
| ---------- | --------------- | ----------------- |
| Created by | `func`          | `{ }`             |
| Lifetime   | Whole function  | Until block ends  |
| Visibility | Entire function | Only inside block |
| Nested     | Yes             | Yes               |

---

## 1️⃣1️⃣ Nested Scope Example (Deep Understanding)

```go
func main() {
    a := 1

    if true {
        b := 2

        if true {
            c := 3
            fmt.Println(a, b, c) // ✅ all accessible
        }

        fmt.Println(c) // ❌ c died
    }

    fmt.Println(b) // ❌ b died
}
```

### Why?

* Inner blocks see outer variables
* Outer blocks **cannot see inside**

📌 **Scope only flows inward, never outward**

---

## 1️⃣2️⃣ Scope Search Rule (One-Line Rule)

> Go always searches from **closest scope outward**

This is why:

* Local overrides global
* Block overrides function
* Function overrides global

---

## 1️⃣3️⃣ Why Block Scope Is a GOOD Thing

Without block scope:

* Variables would clash
* Bugs would be invisible
* Code would be unsafe

Block scope:

* Prevents accidental misuse
* Limits variable lifetime
* Makes reasoning easier

---

## 1️⃣4️⃣ Real Engineering Principle Behind This

This follows **Single Responsibility Principle (SRP)**:

> Each block should manage only what it needs.

Same reason:

* Functions should do one job
* Blocks should keep data minimal

---

## 1️⃣5️⃣ Common Beginner Mistake 🚨

```go
if x > 0 {
    y := 5
}
fmt.Println(y) // ❌
```

❌ Thinking:

> “I already declared `y` above”

✅ Reality:

> `y` never existed outside the block

---

## 1️⃣6️⃣ Interview Question You WILL Get

> **What is block scope in Go?**

### Perfect Answer:

> Block scope refers to variables declared inside `{ }`, which are accessible only within that block and destroyed once the block ends.

---

## 1️⃣7️⃣ Final Golden Rules (LOCK THESE 🔒)

* `{ }` = new scope
* Inner scopes can see outer variables
* Outer scopes CANNOT see inner variables
* Variables die when scope ends
* Undefined variable = scope problem (99% of the time)

---

## 🧠 Final Visualization (Remember This Forever)

```
Global
 └── Function
      └── Block
```

You can see **upwards**, never **downwards**.

---

## 🚀 Why This Lecture Matters So Much

If you master:

* Local scope
* Block scope

Then:

* Closures become easy
* Loops make sense
* Bugs reduce drastically
* Interviews become trivial
