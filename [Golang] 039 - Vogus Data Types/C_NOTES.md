# 🧠 1. Why He Calls Them “Bogus” Data Types

He calls them *bogus* because:

* People **memorize** them
* Don’t understand **why they exist**
* Use `int` everywhere without thinking

But in reality, these data types exist because of:

* **CPU architecture**
* **Memory efficiency**
* **Performance**
* **Real engineering decisions**

So they are **not bogus**.
They are **low-level power tools** ⚙️

---

# 🧠 2. First Principle: Bit → Byte → Memory Cell

Let’s lock this in.

```
1 bit  = 0 or 1
8 bits = 1 byte
```

Memory is made of **cells** (like boxes).

Depending on architecture:

* 32-bit system → memory cell = 32 bits
* 64-bit system → memory cell = 64 bits

📌 A **data type decides how many bits** of a memory cell it will occupy.

---

# 🧠 3. `int` vs `int8` (MOST IMPORTANT IDEA)

## 🔹 `int`

```go
var x int = 10
```

What happens?

* Go checks **CPU architecture**
* If CPU is:

  * 32-bit → `int` = 32 bits
  * 64-bit → `int` = 64 bits

📌 `int` is **architecture-dependent**

---

## 🔹 `int8`

```go
var x int8 = 10
```

What happens?

* Exactly **8 bits** are allocated
* No more, no less
* Remaining bits are left unused (handled by Go runtime)

📌 `int8` is **architecture-independent**

---

# 🧠 4. Why Not Always Use `int`?

Because **memory matters**.

Imagine:

* 1 million values
* Each value only needs range 0–100

Using:

* `int` → 8 bytes each (on 64-bit)
* `int8` → 1 byte each

📊 Memory usage:

```
int   → 8 MB
int8  → 1 MB
```

🔥 That’s **8× difference**

This is how:

* Mobile apps stay smooth
* Backend services scale
* Good engineers are separated from average ones

---

# 🧠 5. Range of Signed Integers (WHY LIMITS EXIST)

Signed integers store:

* Positive
* Negative
* Zero

Formula:

```
Signed range = -(2^(n-1))  to  (2^(n-1) - 1)
```

| Type  | Bits | Range            |
| ----- | ---- | ---------------- |
| int8  | 8    | -128 → 127       |
| int16 | 16   | -32,768 → 32,767 |
| int32 | 32   | ~ -2.1B → +2.1B  |
| int64 | 64   | INSANELY LARGE   |

### Example:

```go
var x int8 = 129 // ❌ ERROR
```

Why?

* Needs **9 bits**
* `int8` only has **8 bits**
* 👉 **Overflow**

---

# 🧠 6. Unsigned Integers (`uint` family)

## 🔹 Meaning of `u`

`u` = **unsigned**
Means:

* ❌ No negative numbers
* ✅ Only zero + positive numbers

---

## 🔹 Example: `uint8`

| Type  | Range      |
| ----- | ---------- |
| int8  | -128 → 127 |
| uint8 | 0 → 255    |

🔥 Same memory, **double positive range**

```go
var a uint8 = 255 // ✅ OK
var b uint8 = -1  // ❌ ERROR
```

---

## 🔹 When to use unsigned?

Use when:

* Negative values **never make sense**
* Example:

  * Age
  * Count
  * Size
  * Index
  * RGB color values

---

# 🧠 7. Boolean (`bool`)

```go
var flag bool = true
```

* Occupies **1 byte (8 bits)**
* Values:

  * `true`
  * `false`

📌 Interview question:

> How many bits does bool take in Go?

✅ **8 bits (1 byte)**

Wrong answer = instant rejection 😅

---

# 🧠 8. Floating Point Types (Decimal Numbers)

## 🔹 Why floats exist

Integers cannot represent:

```
10.25
3.14159
```

So we need **floating-point numbers**.

---

## 🔹 `float32`

* Uses **32 bits**
* Less precision
* Faster
* Smaller memory

## 🔹 `float64`

* Uses **64 bits**
* More precision
* Default choice in Go

```go
var a float32 = 10.23
var b float64 = 10.23456789
```

---

### 🧠 What if CPU is 32-bit and you use `float64`?

Go runtime:

* Uses **two memory cells**
* Combines them
* Manages alignment internally

📌 This is handled by **Go runtime (mini OS)**

---

# 🧠 9. Go Runtime = Mini Operating System

VERY IMPORTANT IDEA.

When a Go program runs:

* OS creates a **process**
* Inside that process:

  * Go runtime starts first
  * Your code runs **inside Go runtime**

Go runtime manages:

* Memory allocation
* Garbage collection
* Goroutines
* Stack growth
* Data layout

🔥 That’s why Go feels “smart”

---

# 🧠 10. `byte` — NOT a New Type

```go
type byte = uint8
```

* Alias (nickname)
* Same thing as `uint8`
* Used mainly for:

  * Files
  * Network data
  * Binary streams

📌 `byte` = **1 byte = 8 bits**

---

# 🧠 11. `rune` — INTERVIEW GOLD 🏆

```go
type rune = int32
```

### Purpose:

* Store **Unicode characters**
* Supports:

  * Bangla
  * Arabic
  * Chinese
  * Emojis ❤️

```go
var r rune = '❤'
```

---

### Why `int32`?

Unicode code points need:

* Large range
* Fixed size
* Signed integer

📌 `rune` = **32 bits = 4 bytes**

---

### Printing rune correctly

❌ Wrong:

```go
fmt.Println(r) // prints number
```

✅ Correct:

```go
fmt.Printf("%c\n", r)
```

---

# 🧠 12. `fmt.Printf` Formatting Cheatsheet (REVISION GOLD)

| Data       | Format |
| ---------- | ------ |
| int / uint | `%d`   |
| float      | `%f`   |
| bool       | `%t`   |
| rune       | `%c`   |
| string     | `%s`   |
| type       | `%T`   |
| newline    | `\n`   |

---

### Precision control

```go
fmt.Printf("%.2f", 10.23678)
// Output: 10.23
```

---

# 🧠 13. Checking Variable Type at Runtime

```go
fmt.Printf("%T\n", variable)
```

Example:

```go
var s string = "Hello"
fmt.Printf("%T\n", s)
// Output: string
```

🔥 Very useful in debugging & interviews

---

# 🧠 14. Core Engineering Lesson (MOST IMPORTANT)

Good engineer:

* Thinks about **memory**
* Thinks about **range**
* Thinks about **data representation**

Bad engineer:

* Uses `int` everywhere
* Doesn’t care about overflow
* Blames system when app is slow

---

# 🧠 FINAL INTERVIEW QUICK ANSWERS

### What is `int`?

> Architecture-dependent integer (32 or 64 bits).

### What is `int8`?

> Signed 8-bit integer (-128 to 127).

### What is `uint8`?

> Unsigned 8-bit integer (0 to 255).

### What is `byte`?

> Alias of `uint8`.

### What is `rune`?

> Alias of `int32`, used for Unicode characters.

### How many bits does `bool` use?

> 8 bits (1 byte).

### How to print rune?

> `%c`

---

# 🧠 Final Words

This lecture wasn’t about Go.
It was about **thinking like a computer scientist**.

You are now:

* Not memorizing
* Not guessing
* **Understanding**

📌 Save these notes
📌 Revisit before interviews
📌 You’ll thank yourself later

