# 📘 Golang Lecture 07 — Why Functions Are Needed (Core Software Engineering)

---

## 1️⃣ The Real Question Behind the Lecture

Many beginners think:

> “Why create functions?
> I can just write code line by line.”

**Technically?** Yes.
**Professionally?** Absolutely not.

This lecture explains **why**.

---

## 2️⃣ The Example Application (Big Picture)

The instructor builds a **mini console application**:

### App flow (human language)

1. Print welcome message
2. Ask user for name
3. Ask for two numbers
4. Add the numbers
5. Show result with user name
6. Say goodbye

This looks simple — but watch what happens when everything is written in **one function**.

---

## 3️⃣ The Problem: One Huge `main()` Function 😵

When everything is written in `main()`:

* Code becomes **long**
* Hard to read
* Hard to debug
* Hard to reuse
* Hard to maintain

### This is what the instructor calls:

> **“Hijibiji code”** (messy, tangled, unreadable)

Even if *you* understand it today,
**you won’t understand it after 2 weeks**.

---

## 4️⃣ Key Insight: Code Is Read More Than It Is Written

In real life:

* You write code once
* You read it **hundreds of times**
* Others read it too

So code must be:

* Clean
* Predictable
* Structured

This is where **functions** come in.

---

## 5️⃣ The Philosophy: Single Responsibility Principle (SRP)

The instructor introduces a **core software engineering principle**:

### 🔑 Single Responsibility Principle (SRP)

> **A function should do ONE thing, and do it well.**

Not two.
Not five.
Not “a bit of everything”.

---

## 6️⃣ Real-Life Analogy (Why SRP Makes Sense)

| Scenario   | Result       |
| ---------- | ------------ |
| One job    | High quality |
| Many jobs  | Confusion    |
| One role   | Focus        |
| Many roles | Mistakes     |

Same for code.

A function that:

* prints welcome
* takes input
* calculates
* prints output

❌ is doing **too much**

---

## 7️⃣ Refactoring the Application (The Right Way)

The instructor **breaks the app into small functions**.

Let’s walk through them one by one.

---

## 8️⃣ Function 1: Print Welcome Message

```go
func printWelcomeMessage() {
    fmt.Println("Welcome to the application")
}
```

### Responsibility

✅ Only prints welcome text
❌ Does nothing else

---

## 9️⃣ Function 2: Get User Name (Input + Return)

```go
func getUserName() string {
    var name string
    fmt.Println("Enter your name")
    fmt.Scanln(&name)
    return name
}
```

### Responsibility

* Ask for name
* Read input
* Return name

📌 **Does NOT print results**
📌 **Does NOT do calculations**

---

## 🔟 Function 3: Get Two Numbers

```go
func getTwoNumbers() (int, int) {
    var num1, num2 int
    fmt.Println("Enter first number")
    fmt.Scanln(&num1)
    fmt.Println("Enter second number")
    fmt.Scanln(&num2)
    return num1, num2
}
```

### Responsibility

* Input only
* No logic
* No printing results

---

## 1️⃣1️⃣ Function 4: Add Numbers

```go
func add(a int, b int) int {
    return a + b
}
```

### Responsibility

* Pure logic
* No input
* No output formatting

📌 This is a **pure function**
(very powerful concept)

---

## 1️⃣2️⃣ Function 5: Display Result

```go
func displayResult(name string, sum int) {
    fmt.Println("Hello", name)
    fmt.Println("Summation is", sum)
}
```

### Responsibility

* Display only
* No calculation
* No input

---

## 1️⃣3️⃣ Function 6: Print Goodbye Message

```go
func printGoodbyeMessage() {
    fmt.Println("Thank you for using the application")
    fmt.Println("Goodbye")
}
```

Simple. Clean. Focused.

---

## 1️⃣4️⃣ Final `main()` — Clean & Professional ✨

```go
func main() {
    printWelcomeMessage()

    name := getUserName()
    num1, num2 := getTwoNumbers()

    sum := add(num1, num2)

    displayResult(name, sum)

    printGoodbyeMessage()
}
```

### Look at this carefully 👀

* Reads like **English**
* Easy to understand
* Easy to change
* Easy to debug

This is **real software engineering**.

---

## 1️⃣5️⃣ What Changed (Before vs After)

### ❌ Before

* One giant function
* Hard to read
* Hard to modify
* No reuse

### ✅ After

* Small focused functions
* Reusable
* Testable
* Maintainable

---

## 1️⃣6️⃣ Why This Matters in Industry

In real projects:

* Millions of lines of code
* Thousands of functions
* Hundreds of developers

Without SRP:

> Chaos. Absolute chaos.

With SRP:

> Systems scale gracefully.

---

## 1️⃣7️⃣ Reusability (Huge Benefit)

Because logic is separated:

* You can reuse `add()` anywhere
* You can reuse `printGoodbyeMessage()`
* You can reuse `getUserName()` in another app

📌 **Reuse = power**

---

## 1️⃣8️⃣ Why Go Loves Functions

Go encourages:

* Small functions
* Clear intent
* Explicit behavior

That’s why:

* Go code looks simple
* Go code scales well
* Go is loved in backend & cloud systems

---

## 🧠 Ultimate Mental Model (Lock This In Forever)

```
One function
   →
One responsibility
   →
Clear code
   →
Scalable systems
```

---

## 🔁 Final Recap (Perfect for Revision)

* Functions exist to **organize behavior**
* One function = one job
* SRP reduces bugs
* Clean `main()` is a sign of good design
* Functions enable reuse & scalability
* This is how **real engineers write code**

---

## 🚀 Where You Are Now

At this point, you are no longer:
❌ “just learning syntax”

You are learning:
✅ **software engineering thinking**

