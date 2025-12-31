# 📘 Golang Lecture 03 — `if`, `else if`, `else` & `switch`

---

## 1️⃣ Why Do We Need Decision Making?

### Real-life example (Instructor’s Facebook login)

When you log in:

* If email & password are correct → login
* Else → show error

👉 **Programs must decide what to do based on data**

This decision-making is done using:

* `if / else`
* `switch`

---

## 2️⃣ Core Idea: What Is a Condition?

A **condition** is an expression that results in:

```text
true  OR  false
```

Go does **not** accept anything else.

---

## 3️⃣ Basic `if` Structure

### Syntax

```go
if condition {
    // code runs if condition is true
}
```

### Mental model

```
Condition checked
   ↓
true  → execute block
false → skip block
```

---

## 4️⃣ Example: Age Check (Marriage Logic)

```go
age := 20

if age > 18 {
    fmt.Println("You are eligible to be married")
}
```

### What happens internally

```
RAM: age = 20
Check: 20 > 18 → true
→ execute print
```

---

## 5️⃣ `if - else if - else` (Multiple Decisions)

### Full structure

```go
if condition1 {
    // runs if condition1 is true
} else if condition2 {
    // runs if condition2 is true
} else {
    // runs if none matched
}
```

📌 **Only ONE block executes**

---

## 6️⃣ Detailed Age Example (All Paths)

```go
age := 18

if age > 18 {
    fmt.Println("Eligible to marry")
} else if age < 18 {
    fmt.Println("Not eligible, but can love")
} else {
    fmt.Println("Just a teenager")
}
```

### Decision flow

```
age > 18 ?  ❌
age < 18 ?  ❌
→ else runs
```

---

## 7️⃣ Comparison Operators (You MUST Remember)

| Operator | Meaning               |
| -------- | --------------------- |
| `>`      | Greater than          |
| `<`      | Less than             |
| `>=`     | Greater than or equal |
| `<=`     | Less than or equal    |
| `==`     | Equal to              |
| `!=`     | Not equal             |

📌 `=` is **assignment**
📌 `==` is **comparison**

---

## 8️⃣ Why `==` Uses Two Equal Signs?

```go
a = 10    // assignment
a == 10   // comparison
```

* `=` → put value in memory
* `==` → ask a question

---

## 9️⃣ Logical Operators (Combining Conditions)

### AND (`&&`)

```go
condition1 && condition2
```

✅ True only if **both are true**

### OR (`||`)

```go
condition1 || condition2
```

✅ True if **any one is true**

### NOT (`!`)

```go
!condition
```

🔁 Reverses the result

---

## 🔟 AND (`&&`) Example

```go
age := 20
gender := "male"

if age == 20 && gender == "male" {
    fmt.Println("Ready to marry")
}
```

### Evaluation

```
age == 20      → true
gender == male → true
true && true   → true
```

✔ Block executes

---

## 1️⃣1️⃣ OR (`||`) Example

```go
if age > 60 || gender == "male" {
    fmt.Println("Condition passed")
}
```

### Logic

```
false || true → true
```

✔ Executes

📌 OR means **“any one is enough”**

---

## 1️⃣2️⃣ NOT (`!`) Example

```go
isPretty := false

if !isPretty {
    fmt.Println("Condition matched")
}
```

### Logic

```
!false → true
```

✔ Executes

📌 NOT flips the truth

---

## 1️⃣3️⃣ Important Rule (How Go Evaluates)

Go evaluates conditions **left to right** and **stops early** if result is decided.

Example:

```go
false && anything → false
true || anything  → true
```

This is called **short-circuit evaluation**.

---

## 1️⃣4️⃣ Why Nothing Prints Sometimes?

```go
if age == 20 {
    fmt.Println("Hello")
}
```

If condition is `false`:

* Go **does nothing**
* This is **normal behavior**

---

## 1️⃣5️⃣ `switch` Statement (Alternative to if-else)

### When to use?

* When checking **one value**
* Against **many possible values**

---

## 1️⃣6️⃣ Basic `switch` Syntax

```go
switch value {
case 1:
    // code
case 2, 3:
    // code
default:
    // code
}
```

📌 No `break` needed in Go (unlike C/C++)

---

## 1️⃣7️⃣ Switch Example Explained

```go
a := 3

switch a {
case 1:
    fmt.Println("a is one")
case 2, 3:
    fmt.Println("a is either two or three")
default:
    fmt.Println("a is neither one nor two nor three")
}
```

### Evaluation flow

```
switch a → 3
case 1?     ❌
case 2, 3?  ✅
→ execute
```

---

## 1️⃣8️⃣ `default` Case

Runs when:

* No `case` matches

Equivalent to:

```go
else
```

---

## 1️⃣9️⃣ If vs Switch (Which Should You Use?)

| Situation         | Recommended |
| ----------------- | ----------- |
| Complex logic     | `if / else` |
| Many fixed values | `switch`    |
| Beginner friendly | `if / else` |

📌 **You can survive your whole life using only `if/else`**

---

## 🧠 Ultimate Mental Model (Very Important)

```
Condition → true / false
true  → execute block
false → skip block
```

For `if-else`:

```
First true block wins
Others are ignored
```

---

## 🔁 Final Recap (Perfect for Revision)

* Programs make decisions using conditions
* Conditions must return `true` or `false`
* `if / else if / else` checks sequentially
* Comparison operators compare values
* Logical operators combine conditions
* `switch` matches one value against many cases

---

## 🚀 You’re Now Ready For:

* Loops (`for`)
* Functions with conditions
* Real-world business logic
* Backend validation rules
