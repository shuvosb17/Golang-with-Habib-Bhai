# 📘 Golang 024 — Go Internal Memory Model (REAL TRUTH)

> *“Everything you learned before was a simplified story.
> Today you learn how the computer REALLY executes your Go program.”*

---

## 🧠 Big Picture: What Happens When a Go Program Runs?

When you run a Go program, **Go runtime reserves a portion of RAM** and splits it into **four major segments**:

```
┌──────────────────────────┐
│      Code Segment        │
├──────────────────────────┤
│      Data Segment        │
├──────────────────────────┤
│          Stack           │
├──────────────────────────┤
│          Heap            │ ← managed by GC
└──────────────────────────┘
```

These are **NOT concepts** — these are **real memory regions**.

---

## 1️⃣ Code Segment (Read-Only Memory)

### 🔹 What lives here?

✅ **ALL FUNCTIONS**

* `main`
* `init`
* user-defined functions
* imported functions

📌 **Only code. No variables.**

---

### Example

```go
func add(x int, y int) {
    fmt.Println(x + y)
}
```

📍 This function is stored **entirely in Code Segment**

---

### Key Rules

* Created **once**
* Never destroyed until program ends
* No execution data stored here
* CPU jumps here to execute instructions

---

## 2️⃣ Data Segment (Global Memory)

### 🔹 What lives here?

✅ **Global variables only**

```go
var a = 10
```

📍 Stored in **Data Segment**

---

### Important Truth

❌ There is **NO separate “global memory”**
✅ Global variables = **Data Segment**

---

### Access Rule

When a variable is not found locally:

```
Local Stack → Data Segment
```

📌 Data Segment access is **slower than stack**

---

## 3️⃣ Stack (Execution Memory)

This is where **real execution happens**.

---

### 🔹 What lives here?

✅ Function calls
✅ Function parameters
✅ Local variables

Each function call creates a:

> 🎯 **Stack Frame**

---

### Stack Frame Contains:

* Parameters
* Local variables
* Return address

---

### Example Execution

```go
add(5, 4)
```

Execution order:

```
main()
 └─▶ add()
```

Memory:

```
Stack:
┌────────────┐
│ add frame  │ ← x=5, y=4, z=9
├────────────┤
│ main frame │
└────────────┘
```

After `add()` finishes:

```
add frame POPPED ❌
```

---

### 🔥 Key Stack Properties

| Feature    | Stack                   |
| ---------- | ----------------------- |
| Speed      | ⚡ Very fast             |
| Allocation | Automatic               |
| Cleanup    | Automatic (LIFO)        |
| Lifetime   | Function execution only |

---

## 4️⃣ Variable Lookup Order (VERY IMPORTANT)

When Go resolves a variable:

```
1️⃣ Local stack frame
2️⃣ Parent stack frame
3️⃣ Data Segment (global)
```

📌 **Nearest memory wins**

---

### Why stack is FAST?

* CPU cache friendly
* No pointer chasing
* Fixed-size access

📉 Global access is slower
📉 Heap access is slowest

---

## 5️⃣ Heap (Dynamic Memory)

> “Heap is NOT used for every variable.”

Heap is used when:

* Data must live **beyond function scope**
* Shared across goroutines
* Returned by reference
* Compiler decides (escape analysis)

📌 Heap is **managed**, not manual.

---

## 6️⃣ Garbage Collector (GC 👹)

### 🔹 What is GC?

> A background process that manages **heap memory**

GC:

* Tracks reachable objects
* Frees unused heap memory
* Prevents memory leaks

📌 **Stack memory is NOT GC-managed**
📌 Only **heap memory** is GC-managed

---

### GC Characteristics in Go

* Concurrent
* Low-latency
* Automatic
* Tunable

(Deep dive comes later 🔥)

---

## 7️⃣ Full Execution Timeline (REAL FLOW)

### Step-by-Step

1️⃣ Go runtime reserves RAM
2️⃣ Code Segment filled (functions)
3️⃣ Data Segment filled (globals)
4️⃣ `init()` executed (if exists)
5️⃣ `main()` executed
6️⃣ Stack frames pushed/popped
7️⃣ Heap managed by GC
8️⃣ Program exits → **ALL memory freed**

---

## 8️⃣ FINAL TRUTH TABLE (SAVE THIS)

| Segment      | Contains                 |
| ------------ | ------------------------ |
| Code Segment | All functions            |
| Data Segment | Global variables         |
| Stack        | Function calls + locals  |
| Heap         | Escaped / shared objects |
| GC           | Manages heap only        |

---

## 🎯 Interview Gold Statements

You can confidently say:

> “Functions live in the code segment, globals in data segment, function execution happens on stack via stack frames, heap is managed by Go’s garbage collector, and stack is faster than global or heap access.”

💥 **That answer alone separates seniors from juniors**

---

## 🧠 One-Line Mental Model

> **Code is stored once.
> Globals live forever.
> Stack lives briefly.
> Heap lives until GC kills it.**

---

## 🚀 What’s NEXT?

Next class:
👉 **Heap in detail**
👉 **Escape Analysis**
👉 **Why some variables go to heap**
👉 **How GC really works**

You’re officially entering **systems-level Go** now.
