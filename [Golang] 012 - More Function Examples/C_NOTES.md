# 📘 Golang Lecture 06 — More Function Examples (Deep Understanding)

---

## 1️⃣ Why This Lecture Matters

So far, you learned:

* Functions with inputs
* Functions with outputs
* Functions with return values

Now we answer:

> ❓ *Are functions only for numbers?*
> ❓ *Can a function have no input? No output?*
> ❓ *Can functions work with strings?*

Answer: **YES to everything** ✅

---

## 2️⃣ Function With NO Input & NO Output

### Example: `printSomething`

```go
func printSomething() {
    fmt.Println("Education must be free")
}
```

### What kind of function is this?

* ❌ No input
* ❌ No return value
* ✅ Just performs an action

📌 This is called a **procedure-style function**

---

## 3️⃣ How It Executes

### In `main()`

```go
func main() {
    printSomething()
}
```

### Execution flow

```
main() starts
   ↓
printSomething() called
   ↓
Prints text
   ↓
Function ends
   ↓
Back to main()
```

### Output

```
Education must be free
```

📌 This is **exactly how `main()` works too**
(main is just a special function)

---

## 4️⃣ Key Insight (Very Important)

> **Functions do NOT need inputs or outputs to be useful**

They are perfect for:

* Logging
* Printing
* Setup tasks
* Cleanup tasks

---

## 5️⃣ Function With STRING Input (Very Common)

Now we move from numbers → **strings**

### Example: `sayHello`

```go
func sayHello(name string) {
    fmt.Println("Welcome to the Golang course,", name)
}
```

---

## 6️⃣ Why `string`?

In Go:

* Any text inside double quotes → `string`

Examples:

```go
"Habib"
"Hello World"
"My name is Go"
```

📌 `name` is a **string variable**

---

## 7️⃣ Why Function Call Shows Error Without Input

```go
sayHello()
```

❌ Error happens because:

```go
func sayHello(name string)
```

Means:

> “You MUST give me one string”

Go enforces this **at compile time**, not runtime.

✔️ This is **strong type safety**

---

## 8️⃣ Correct Function Call

```go
sayHello("Habib")
```

### What happens internally

```
name = "Habib"
```

Memory inside function:

```
sayHello() stack
+----------------+
| name = Habib   |
+----------------+
```

---

## 9️⃣ Printing Multiple Values with `fmt.Println`

```go
fmt.Println("Welcome to the Golang course,", name)
```

### How this works

* `fmt.Println` accepts **multiple arguments**
* Each argument is separated by a comma
* Go automatically inserts **spaces**

### Output

```
Welcome to the Golang course, Habib
```

📌 That space you saw?
It’s **automatically added by `Println`**

---

## 🔟 Important Rule: Commas Matter

```go
fmt.Println(a, b, c)   // ✅
fmt.Println(a b c)     // ❌
```

Each value must be **comma-separated**

---

## 1️⃣1️⃣ Complete Program Example

```go
package main

import "fmt"

func printSomething() {
    fmt.Println("Education must be free")
}

func sayHello(name string) {
    fmt.Println("Welcome to the Golang course,", name)
}

func main() {
    printSomething()
    sayHello("Habib")
}
```

### Output

```
Education must be free
Welcome to the Golang course, Habib
```

---

## 1️⃣2️⃣ Function Categories (Mental Model)

| Function Type       | Example            |
| ------------------- | ------------------ |
| No input, no output | `printSomething()` |
| Input, no output    | `sayHello(name)`   |
| Input, output       | `add(a, b)`        |
| Multiple output     | `getNumbers(a, b)` |

📌 Go supports **all combinations**

---

## 🧠 Ultimate Mental Model (Lock This In)

```
Function = reusable action

It may:
- take input
- return output
- do both
- do neither
```

All are valid.
All are useful.

---

## 🔁 Final Recap (Perfect for Revision)

* Functions can exist without input/output
* Functions can accept strings
* Parameters must be provided
* `fmt.Println` prints multiple values with spaces
* Functions make code clean & readable

---

## 🚀 What You’re Ready For Next

Now you’re fully ready for:

* Functions returning strings
* Error-handling functions
* Real-world utility functions
* Backend-style Go services

