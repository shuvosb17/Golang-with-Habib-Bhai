# 🧠 Golang 025 — Internal Memory (FINAL PIECE)

> **Core Question of this class**
> 👉 *If I declare a function expression inside another function, where does it live?*
> **Code Segment? Stack? Heap?**

This class answers that **once and for all**.

---

## 1️⃣ Two Phases of a Go Program (VERY IMPORTANT)

Every Go program runs in **two completely different phases**:

```
1️⃣ Compilation Phase
2️⃣ Execution Phase
```

You MUST separate these mentally.

---

## 🔹 Phase 1: Compilation Phase (go build)

### What happens here?

* Go compiler reads your `.go` files
* Converts everything into **binary (0 & 1)**
* Creates a **single executable file**

```bash
go build main.go
```

📁 Output:

```
./main   ← binary executable
```

---

### 💡 Key Truth

> **The computer does NOT understand Go code.
> It understands only binary.**

So all names (`main`, `add`, `call`) **disappear** at this stage.

---

## 2️⃣ What Goes Where (Compilation Phase)

### 🔹 Code Segment (inside the binary)

During compilation, Go puts **READ-ONLY things** here:

| Goes to Code Segment | Why          |
| -------------------- | ------------ |
| All functions        | Never change |
| Constant values      | Immutable    |
| Instructions         | Read-only    |

Example:

```go
const A = 10
func add(x, y int) {}
func call() {}
func main() {}
func init() {}
```

➡️ **ALL of these go to the Code Segment**

---

### 🔹 Data Segment (also inside the binary)

| Goes to Data Segment | Why     |
| -------------------- | ------- |
| Global variables     | Mutable |

```go
var p = 100
```

➡️ Stored in **Data Segment**

---

### 🚫 Important Correction (big lie before this class)

❌ “Global memory” is NOT a separate thing
✅ Global variables live in **Data Segment**

---

## 3️⃣ Binary File → RAM (Execution Phase)

Now comes execution:

```bash
./main
```

This is when **real memory allocation starts**.

---

### Runtime creates 4 regions in RAM:

```
┌──────────────────────────┐
│      Code Segment        │ ← loaded from binary
├──────────────────────────┤
│      Data Segment        │ ← globals
├──────────────────────────┤
│          Stack           │ ← function execution
├──────────────────────────┤
│          Heap            │ ← GC managed
└──────────────────────────┘
```

---

## 4️⃣ init() and main() (REAL TRUTH)

⚠️ **Correction from previous simplifications**

> `init()` and `main()` existence is checked
> **during compilation**, not runtime

If `main()` doesn’t exist → **compile error**

---

### Execution Order:

```
init()  → main()
```

---

## 5️⃣ Stack & Stack Frames (Execution Phase)

Whenever a function runs:

➡️ **A stack frame is created**

```
Stack Frame =
- Parameters
- Local variables
- References
```

When function ends → **frame is destroyed**

---

### Example

```go
func add(x, y int) {
    z := x + y
    fmt.Println(z)
}
```

Memory during execution:

```
Stack:
┌────────────┐
│ add frame  │ ← x, y, z
├────────────┤
│ main frame │
└────────────┘
```

After return → `add frame` is **popped**

---

## 6️⃣ The BIG QUESTION: Function Expression Inside Function

### Example Code

```go
func call() {
    add := func(x, y int) {
        fmt.Println(x + y)
    }

    add(5, 6)
}
```

### ❓ Where does this `add` function live?

---

## 🧨 FINAL ANSWER (THIS IS GOLD)

### 🔥 The function itself:

➡️ **Code Segment (created at compile time)**

### 🔥 The variable `add`:

➡️ **Stack frame of `call()`**

### 🔥 What does `add` store?

➡️ **A reference (pointer) to the function in Code Segment**

---

## 7️⃣ Visual Memory Model (IMPORTANT)

```
Code Segment (READ ONLY)
------------------------
[ func add(x,y) {...} ]  ← function body exists ONCE

Stack (Execution)
-----------------
call() stack frame
┌──────────────────────┐
│ add → ref(0x0042)    │  ← reference to code segment
└──────────────────────┘
```

> The function is NOT copied
> Only the **reference** is stored

---

## 8️⃣ Why Only `call()` Can Use `add`

Because:

* `add` variable exists **only inside call’s stack frame**
* When `call()` returns → stack frame destroyed
* Reference disappears
* Function still exists, but **no one can reach it**

📌 **This is lexical scoping + lifetime**

---

## 9️⃣ Why It’s NOT Stored in Stack or Heap

| Location       | Why NOT                    |
| -------------- | -------------------------- |
| Stack          | Stack is temporary         |
| Heap           | Function code is immutable |
| Data Segment   | It’s not global            |
| ✅ Code Segment | Correct                    |

---

## 10️⃣ Constants vs Variables (VERY IMPORTANT)

### Constant

```go
const A = 10
```

* Stored in **Code Segment**
* Inlined
* Cannot change

---

### Variable

```go
var p = 100
```

* Stored in **Data Segment**
* Mutable
* Accessed at runtime

---

## 11️⃣ Variable Lookup Order (FINAL)

When Go resolves a name:

```
1️⃣ Local stack frame
2️⃣ Parent stack frame
3️⃣ Data Segment (globals)
4️⃣ Code Segment (constants & functions)
```

---

## 12️⃣ Why Stack Is Faster Than Global

| Memory       | Speed     |
| ------------ | --------- |
| Stack        | ⚡ Fastest |
| Data Segment | Slower    |
| Heap         | Slowest   |

📌 That’s why Go prefers stack whenever possible.

---

## 13️⃣ Heap & GC (Short Recap)

* Heap used only when **escape analysis** decides
* Managed by **Garbage Collector**
* Stack is NOT GC managed

(Next classes go deep here)

---

## 14️⃣ Final Mental Model (SAVE THIS)

> **Functions & constants are born at compile time
> Variables live at runtime
> Stack is temporary
> Heap is managed
> References connect everything**

---

## 🎯 Interview-Ready Answer (Perfect)

> “Function expressions do not create new functions at runtime.
> The function body is stored once in the code segment during compilation,
> and the variable holds only a reference to it inside the stack frame.”

💥 This answer alone puts you in **top 10%**

---

## 🧠 Why This Lesson Matters So Much

Because now:

* Closures make sense
* Escape analysis makes sense
* GC behavior makes sense
* Performance decisions make sense
* Senior Go interviews become survivable

---

## 🚀 What Comes Next

From here:
➡️ Closures (real closure internals)
➡️ Heap vs Stack (escape analysis)
➡️ GC internals
➡️ Goroutines memory model

