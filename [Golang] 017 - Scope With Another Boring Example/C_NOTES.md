# 📘 Golang 011 — Scope With Another (So-Called) Boring Example

> ⚠️ *This example looks boring because it’s SIMPLE.*
> But **simple things repeated deeply = mastery**.

---

## 🧠 What This Class Is REALLY About

This class answers **three critical questions**:

1. How functions find **variables**
2. How functions find **other functions**
3. Why **function order does NOT matter** in Go

And it does this using:

* Global scope
* Function scope
* Call stack simulation
* Memory allocation logic

---

## 🧱 The Code Structure (Mental Model)

We have:

### Global variables

```go
var a = 10
var b = 20
```

### Functions

```go
func printNumber(n int)
func add(x int, y int)
func main()
```

### Call chain

```
main → add → printNumber
```

---

## 🧩 What Each Function Does

### `printNumber`

* Takes **one integer**
* Prints it
* That’s it

### `add`

* Takes **two integers**
* Adds them
* Stores result
* Calls `printNumber(result)`

### `main`

* Calls `add(a, b)`

---

## 🖥️ Step-by-Step Execution (THIS IS THE GOLD)

### 1️⃣ Program Starts → Global Scope Created

When Go runs the program:

* A **global memory space** is allocated
* Everything outside functions goes here

Stored globally:

* `a = 10`
* `b = 20`
* `printNumber` (function definition)
* `add` (function definition)
* `main` (function definition)

📌 **Important**

> Functions are stored in memory **when declared**, not when called.

---

## 2️⃣ Go Looks for `main()`

Go scans global scope:

* ❌ No `main` → program fails
* ✅ `main` found → program starts

---

## 3️⃣ `main()` Is Executed

### Memory action:

* New memory block allocated for `main`
* This is **local to main**

Inside `main`:

```go
add(a, b)
```

---

## 4️⃣ Resolving `add(a, b)`

### Scope resolution rules (VERY IMPORTANT):

When Go sees `add`:

1. Check **local scope (main)** ❌
2. Check **global scope** ✅ found

When Go sees `a` and `b`:

1. Check **main scope** ❌
2. Check **global scope** ✅ found

So Go calls:

```go
add(10, 20)
```

---

## 5️⃣ `add()` Is Executed

### Memory action:

* New memory allocated for `add`
* Separate from `main`

Inside `add`:

```go
x = 10
y = 20
result = x + y  // 30
```

Now:

```go
printNumber(result)
```

---

## 6️⃣ Resolving `printNumber(result)`

Go tries to find `printNumber`:

1. Inside `add` ❌
2. Inside global scope ✅ found

Then:
Сhecks `result`:

* Exists inside `add`
* Value = `30`

So Go calls:

```go
printNumber(30)
```

---

## 7️⃣ `printNumber()` Is Executed

### Memory action:

* New memory allocated for `printNumber`

Inside:

```go
n = 30
fmt.Println(n)
```

📤 Output:

```
30
```

---

## 8️⃣ Function Cleanup (VERY IMPORTANT)

After execution:

1. `printNumber` finishes → memory **freed**
2. Control returns to `add`
3. `add` finishes → memory **freed**
4. Control returns to `main`
5. `main` finishes → program ends
6. Global memory released

📌 **This is the call stack lifecycle**

---

## 🔁 The BIG SURPRISE: Function Order Does NOT Matter

The instructor rearranged the code:

* `add` first
* `main` in the middle
* `printNumber` at the bottom

### Result?

✅ **Program still works perfectly**

### Why?

Because:

* Go **reads the entire file first**
* Builds the **global symbol table**
* Then executes `main`

📌 **Order matters in execution, NOT declaration**

---

## 🧠 Scope Resolution Rule (MASTER RULE)

Whenever Go sees a name:

1. Look in **current local scope**
2. Then parent function scope
3. Then **global scope**
4. If not found → ❌ compile error

---

## 🧪 Why This Example Matters

This example proves:

* Functions don’t need to be written before use
* Variables can come from parent scopes
* Functions can call other functions freely
* Scope resolution is **predictable & deterministic**

---

## 🧠 Core Takeaways (MEMORIZE THESE)

### ✅ Facts

* Functions are stored in global scope when declared
* Function calls create **new memory**
* Local variables die when function ends
* Go does NOT care about function order
* Scope lookup always goes **inside → outside**

---

## 📋 One-Page Scope Checklist

| Item                | Scope                              |
| ------------------- | ---------------------------------- |
| Global variables    | Global                             |
| Function parameters | Function                           |
| Local variables     | Block / function                   |
| Functions           | Global (if declared outside)       |
| Function order      | Irrelevant                         |
| Memory              | Allocated on call, freed on return |

---

## 🎯 Why the Instructor Called It “Boring”

Because:

* It repeats basics
* No new syntax
* No fancy features

But:

> **This repetition builds intuition**

That intuition is what lets you:

* Debug fast
* Read large Go codebases
* Pass interviews
* Write clean production code

---

## 🚀 What You’re Ready For Now

After lessons **014 → 017**, you now fully understand:

* Global scope
* Local scope
* Block scope
* Package scope
* Function scope
* Call stack behavior

👉 **You’re officially out of beginner territory**

