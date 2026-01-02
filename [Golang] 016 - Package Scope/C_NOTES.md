# 📘 Golang 10 — Package Scope (Deep & Practical Explanation)

---

## 1️⃣ Why Package Scope Exists (Big Picture)

So far, everything lived in **one file** or **one function**.

But real Go programs:

* Have **multiple files**
* Have **multiple folders**
* Have **multiple packages**

👉 **Package scope controls what is visible across files and folders**

Without package scope:

* Large projects would be impossible
* Code reuse would break
* Libraries couldn’t exist

---

## 2️⃣ Rule #1 (VERY IMPORTANT)

> **All `.go` files inside the SAME folder must have the SAME package name**

### ❌ Wrong

```
main.go   → package main
add.go    → package habib
```

Go error:

```
found packages habib and main
```

### ✅ Correct

```
main.go   → package main
add.go    → package main
```

📌 **Folder = Package boundary**

---

## 3️⃣ Files vs Packages (Clear Difference)

| Concept | Meaning                |
| ------- | ---------------------- |
| File    | Just a file            |
| Folder  | Defines a package      |
| Package | Logical scope boundary |

👉 Files don’t create scope
👉 **Packages do**

---

## 4️⃣ Why `add()` Was “Undefined” Before ❌

You ran:

```bash
go run main.go
```

But `add()` lived in:

```
add.go
```

### What Go saw:

* Only `main.go`
* No `add()` in global scope

### Fix:

```bash
go run main.go add.go
```

📌 **Go only compiles files you explicitly include**

---

## 5️⃣ What Happens When You Run Multiple Files

When you run:

```bash
go run main.go add.go
```

Go does this internally:

1. Reads **all files**
2. Collects **global declarations**
3. Builds a **single package scope**
4. Then runs `main()`

That’s why `add()` suddenly works.

---

## 6️⃣ Package Scope = Shared Global Scope (Inside Package)

Inside **one package**:

* All global variables
* All global functions
* Are visible across files

Example:

```go
// add.go
package main

func add(a, b int) {
    fmt.Println(a + b)
}
```

```go
// main.go
package main

func main() {
    add(4, 7)
}
```

✔️ Works perfectly.

---

## 7️⃣ Now the REAL Package Scope (Different Folder)

Now comes the **real concept**.

### Folder structure:

```
project/
 ├── main.go
 ├── go.mod
 └── mathlibrary/
     └── math.go
```

This creates **two packages**:

* `main`
* `mathlibrary`

---

## 8️⃣ Why `mathlibrary` Was Not Found ❌

You tried:

```go
import "mathlibrary"
```

Go said:

```
cannot find package
```

Why?

Because:

* Go needs a **module path**
* Not just a folder name

---

## 9️⃣ What `go mod init` REALLY Does

When you ran:

```bash
go mod init example.com
```

Go created:

```go
module example.com
```

📌 This becomes the **root import path**

Now packages are referenced like:

```
example.com/mathlibrary
```

---

## 🔟 Correct Import Syntax (Custom Package)

```go
import "example.com/mathlibrary"
```

Then usage:

```go
mathlibrary.Add(4, 7)
```

---

## 1️⃣1️⃣ The MOST IMPORTANT Package Scope Rule 🔥

> **Only identifiers starting with CAPITAL letters are accessible outside the package**

This is Go’s visibility rule.

---

## 1️⃣2️⃣ Exported vs Unexported (CRITICAL)

| Name    | Visibility        |
| ------- | ----------------- |
| `add()` | ❌ package-private |
| `Add()` | ✅ exported        |
| `money` | ❌ private         |
| `Money` | ✅ exported        |

📌 **Capital letter = exported = visible across packages**

---

## 1️⃣3️⃣ Why `add()` Was Invisible ❌

```go
func add(a, b int) { }
```

From another package:

* ❌ Cannot see it
* ❌ Autocomplete won’t show it
* ❌ Compile error

### Fix:

```go
func Add(a, b int) { }
```

---

## 1️⃣4️⃣ Variables Follow SAME Rule

```go
var money = 100   // ❌ private
var Money = 100   // ✅ exported
```

Only `Money` is visible outside.

---

## 1️⃣5️⃣ Complete Visibility Table (Save This)

| Scope               | Visible Where     |
| ------------------- | ----------------- |
| Block               | Inside `{}` only  |
| Function            | Inside function   |
| Package (lowercase) | Same package only |
| Package (Capital)   | Any package       |
| Global              | Package-level     |

---

## 1️⃣6️⃣ Why Go Uses Capital Letters (Design Reason)

Unlike other languages:

* No `public`
* No `private`
* No `protected`

Go uses:

> **Capitalization as access control**

This makes Go:

* Simple
* Readable
* Enforced at compile time

---

## 1️⃣7️⃣ Mental Model (Very Important)

Think of packages like **buildings**:

* Lowercase names → private rooms
* Capital names → public entrance

You can only enter **public doors**.

---

## 1️⃣8️⃣ Why Package Scope Is Powerful

Package scope enables:

* Clean APIs
* Safe libraries
* Reusable code
* Industry-scale projects

This is how:

* `fmt.Println`
* `http.ListenAndServe`
* `json.Marshal`

all work.

---

## 1️⃣9️⃣ Common Interview Question

> **How does Go control visibility across packages?**

### Perfect Answer:

> Go uses capitalization. Identifiers starting with uppercase letters are exported and accessible from other packages; lowercase identifiers are package-private.

---

## 🔁 Final Recap (Revision Checklist)

* Same folder → same package
* Folder defines package
* `go mod init` defines module root
* Import uses `module/folder`
* Capital letter = exported
* Lowercase = private
* Package scope ≠ file scope

---

## 🚀 You’re Now Past Beginner Level

If you understand:

* Global scope
* Local & block scope
* Package scope
