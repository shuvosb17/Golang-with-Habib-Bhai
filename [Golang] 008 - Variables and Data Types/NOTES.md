# 📘 Golang Lecture 02 — Variables & Data Types (Deep Dive)

---

## 1️⃣ What Is a Variable? (Core Idea)

### Simple definition

A **variable** is a **container** that holds **data**.

### Real-world analogy

| Real World     | Programming      |
| -------------- | ---------------- |
| Cup            | Variable         |
| Water / Tea    | Data             |
| Pour new drink | Assign new value |

> The cup stays the same — only the content changes.

---

## 2️⃣ Why Variables Exist (Computer Science Truth)

At its core:

> **Programming = storing, changing, and moving data**

A program without data is useless.

So:

* Data needs a **place**
* That place needs a **name**
* That named place is called a **variable**

---

## 3️⃣ Where Variables Live (Inside the Computer)

Let’s zoom inside a computer 🖥️

```
Computer
 ├── CPU        → does calculations
 ├── RAM        → temporary memory (FAST)
 └── Disk       → permanent storage (SLOW)
```

📌 **Variables live in RAM**

---

## 4️⃣ RAM Mental Model (Very Important)

Think of RAM as many small boxes:

```
RAM
+-----+-----+-----+-----+
|     |     |     |     |
+-----+-----+-----+-----+
```

Each box:

* Can store **one value**
* Has an **address**
* Can be given a **name** (variable)

---

## 5️⃣ Declaring a Variable (Short Form)

### Code

```go
a := 10
```

### What Go does internally

```
Step 1: Find empty RAM cell
Step 2: Name it "a"
Step 3: Store value 10
```

```
RAM
+-----+-----+-----+
| a=10|     |     |
+-----+-----+-----+
```

---

## 6️⃣ Printing a Variable

```go
fmt.Println(a)
```

### Execution flow

```
Look for variable "a"
→ Go to RAM
→ Read value (10)
→ Print 10
```

📌 **No quotes**
Because `a` is a variable, not text.

---

## 7️⃣ Why Order Matters (Undefined Error Explained)

### ❌ Wrong

```go
fmt.Println(a)
a := 10
```

### Why error?

At print time:

* Variable **does not exist yet**
* RAM cell not created

📌 **Rule**

> A variable must be declared **before it is used**

---

## 8️⃣ What Is a Data Type?

A **data type** tells Go:

* What kind of data is stored
* How much memory to allocate
* What operations are allowed

---

## 9️⃣ Main Data Types (You Actually Need)

### 🔢 Numeric

* `int` → whole numbers
* `float32`, `float64` → decimal numbers

### 🔤 String

* Text inside double quotes

### ✅ Boolean

* `true`
* `false`

---

## 🔟 Examples of Data Types

| Value           | Type     |
| --------------- | -------- |
| `10`            | `int`    |
| `40.34`         | `float`  |
| `"Hello World"` | `string` |
| `true`          | `bool`   |

---

## 1️⃣1️⃣ Explicit Declaration (Long Form)

```go
var x int = 10
```

### Meaning

```
var     → declare variable
x       → variable name
int     → data type
10      → value
```

📌 Use this when:

* You want **clarity**
* You care about **type control**

---

## 1️⃣2️⃣ Type Inference (Go Is Smart)

```go
a := 10
```

Go automatically infers:

```
10 → integer → int
```

📌 This is called **type inference**

---

## 1️⃣3️⃣ Changing Variable Values

```go
a := true
a = false
```

### RAM behavior

```
First:  a → true
Then:   a → false
```

Old value is **replaced**, not duplicated.

---

## 1️⃣4️⃣ Why `:=` Works Only Once

```go
a := 10   // declaration
a = 20    // assignment
```

❌ This is illegal:

```go
a := 20
```

### Why?

* `:=` means **declare + assign**
* After declaration, variable already exists

📌 Think of it like:

> You can name a baby once — after that, you just talk to them.

---

## 1️⃣5️⃣ Type Safety (Why Go Refuses Mixed Types)

```go
a := true
a = "Habib" // ❌ error
```

### Why Go rejects this

* `a` was declared as `bool`
* `"Habib"` is a `string`
* Go **does not allow type changes**

📌 This prevents:

* Bugs
* Memory corruption
* Runtime crashes

---

## 1️⃣6️⃣ Constants (`const`)

```go
const pi = 100
```

### Meaning

* Value **cannot change**
* Read-only memory

❌ Illegal:

```go
pi = 50
```

📌 Use constants for:

* Fixed values
* Configuration
* Mathematical constants

---

## 1️⃣7️⃣ Execution Order Matters (Important Example)

```go
a := 100
a = 50
fmt.Println(a)
a = 109
```

### Output

```
50
```

### Why?

```
Line 1 → a = 100
Line 2 → a = 50
Line 3 → print a (50)
Line 4 → a = 109 (too late)
```

📌 Go runs **top to bottom**

---

## 1️⃣8️⃣ Comments (Ignored by Go)

```go
// single-line comment

/*
multi-line
comment
*/
```

Used for:

* Explanation
* Debugging
* Documentation

---

## 1️⃣9️⃣ Types You Saw (But Don’t Need Yet)

| Type                              | Meaning                 |
| --------------------------------- | ----------------------- |
| `int8`, `int16`, `int32`, `int64` | Memory size control     |
| `uint`                            | Unsigned integers       |
| `float64`                         | High precision decimals |
| `complex`                         | Complex numbers         |

📌 These matter later (OS, memory, performance)

---

## 🧠 Ultimate Mental Model (Remember This)

```
Variable = Named RAM cell

RAM cell:
- Has a name
- Has a fixed type
- Holds one value at a time
```

---

## 🔁 Final Recap (Perfect for Revision)

* Variables store data in **RAM**
* Data has **types**
* Go is **strict but safe**
* `:=` declares once
* Type cannot change
* Constants never change
* Execution is **top → bottom**

---

## 🚀 You Are Now Ready For:

* Conditions (`if`)
* Loops
* Functions
* Real Go programs
