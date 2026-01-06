# 📘 Golang 012 — Variable Shadowing (Deep & Clear)

---

## 1️⃣ What Is Variable Shadowing? (Plain English)

> **Variable Shadowing means:**
> A variable declared in a **smaller scope** has the **same name** as a variable in a **bigger scope**, and the smaller one **hides** the bigger one.

📌 The outer variable still exists —
📌 But it becomes **invisible (shadowed)** inside the inner scope.

---

## 2️⃣ The Core Example (Interview Favorite)

### Code

```go
var a = 10

func main() {
    age := 30

    if age > 18 {
        a := 47
        fmt.Println(a)
    }

    fmt.Println(a)
}
```

---

## 3️⃣ Before Thinking About Output — THINK ABOUT SCOPE

### Where variables live:

| Variable   | Scope            |
| ---------- | ---------------- |
| `a = 10`   | Global scope     |
| `age = 30` | `main()` scope   |
| `a := 47`  | `if` block scope |

⚠️ **Important:**
`a := 47` does **NOT modify** the global `a`
It creates a **new variable** with the same name.

---

## 4️⃣ Step-by-Step Mental Simulation (THIS IS KEY)

---

### 🔹 Step 1: Global Scope Created

```text
Global Memory
-------------
a = 10
main()
```

---

### 🔹 Step 2: `main()` Starts → Local Scope

```text
Main Memory
-----------
age = 30
```

---

### 🔹 Step 3: `if age > 18`

* `age = 30`
* `30 > 18` → ✅ TRUE
* Enter the **if block**

---

### 🔹 Step 4: Inside `if` Block → New Block Scope

```go
a := 47
```

⚠️ This creates a **new variable `a`**

```text
Block Memory (if)
-----------------
a = 47   ← shadows global a
```

---

### 🔹 Step 5: First Print

```go
fmt.Println(a)
```

Where does Go look?

1. Current block → ✅ found `a = 47`
2. Stops searching

📤 Output:

```
47
```

---

### 🔹 Step 6: Exit `if` Block

* Block memory destroyed
* `a = 47` **dies**

---

### 🔹 Step 7: Second Print

```go
fmt.Println(a)
```

Where does Go look now?

1. Current block → ❌
2. `main()` scope → ❌
3. Global scope → ✅ `a = 10`

📤 Output:

```
10
```

---

## 5️⃣ Final Output (Confirmed)

```
47
10
```

---

## 6️⃣ Why This Confuses People (Interview Trap)

People think:

> “I changed `a` to 47”

❌ **WRONG**

What actually happened:

* You created **another `a`**
* That `a` lived only inside the block
* It **shadowed** the outer `a`

---

## 7️⃣ Shadowing ≠ Overwriting (CRITICAL DIFFERENCE)

### ❌ Shadowing (new variable)

```go
a := 47
```

### ✅ Overwriting (same variable)

```go
a = 47
```

⚠️ Using `:=` **creates a new variable**
⚠️ Using `=` **updates existing variable**

---

## 8️⃣ Scope Lookup Rule (MEMORIZE THIS)

Whenever Go sees a variable name:

```
1️⃣ Current block
2️⃣ Parent function
3️⃣ Global scope
4️⃣ ERROR if not found
```

Go **stops immediately** at the first match.

---

## 9️⃣ Why It’s Called “Shadowing”

Think of sunlight ☀️

* Global variable = person standing
* Inner variable = shadow covering the person
* Person still exists
* But you **can’t see them**

That’s **shadowing**.

---

## 🔥 Real-World Bug Example (VERY COMMON)

```go
err := doSomething()

if err != nil {
    err := logError(err)
}
```

❌ Outer `err` is NOT updated
❌ Bug stays hidden
❌ Debugging nightmare

---

## ✅ Best Practices (PRODUCTION RULES)

### ✔ Avoid shadowing unless intentional

### ✔ Don’t reuse variable names across scopes

### ✔ Be extra careful with `:=`

### ✔ Use clear naming in nested blocks

---

## 🧠 One-Line Definition (Interview-Ready)

> **Variable shadowing occurs when a variable declared in an inner scope has the same name as one in an outer scope, hiding the outer variable within that scope.**

---

## 📌 Final Takeaways

* Shadowing creates **new variables**
* It does **not modify outer variables**
* Scope decides which variable is visible
* `:=` is the most common cause
* Interviews LOVE this topic

---

## 🚀 What You’re Ready For Next

After **018**, you now understand:

* Global vs local variables
* Block lifetime
* Shadowing dangers
* Real debugging intuition

