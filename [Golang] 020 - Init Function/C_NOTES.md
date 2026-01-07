# 📘 Golang 020 — `init()` Function (Deep & Structured)

---

## 1️⃣ What Is the `init()` Function?

> **`init()` is a special function in Go that runs automatically before `main()`**

Key facts (non-negotiable rules):

* ❌ You **cannot call** `init()` manually
* ❌ `init()` **takes no parameters**
* ❌ `init()` **returns nothing**
* ✅ Go runtime **calls it automatically**
* ✅ It runs **after global variables initialize**
* ✅ It runs **before `main()`**

---

## 2️⃣ Why `init()` Exists (The WHY)

Go needed a way to:

* Prepare configuration
* Initialize state
* Validate environment
* Set defaults
* Run setup logic

👉 **Without polluting `main()`**

So Go designers said:

> “Let the runtime call a function automatically, exactly once, at startup.”

That function is `init()`.

---

## 3️⃣ The Execution Order (MOST IMPORTANT)

This is the **real truth** 👇
(Not the simplified version you usually hear)

### ✅ Actual Startup Order in Go

```
1️⃣ Global variables (top → bottom)
2️⃣ init() functions (file-wise, package-wise)
3️⃣ main()
```

📌 `init()` **always runs before `main()`**

---

## 4️⃣ Simple Example (Baseline)

```go
func init() {
    fmt.Println("I am the first function executed")
}

func main() {
    fmt.Println("Hello main function")
}
```

### Output:

```
I am the first function executed
Hello main function
```

✔ No call to `init()`
✔ Still executed
✔ Automatically

---

## 5️⃣ Why You Cannot Call `init()`

Try this:

```go
init()
```

❌ Compile error: **undefined: init**

Because `init()` is **runtime-controlled**, not programmer-controlled.

This prevents:

* Accidental multiple calls
* Dependency chaos
* Startup bugs

---

## 6️⃣ `init()` + Global Variables (Important Example)

### Code

```go
var a = 10

func init() {
    fmt.Println(a)
    a = 20
}

func main() {
    fmt.Println(a)
}
```

---

## 7️⃣ Step-by-Step Runtime Simulation 🧠

### 🔹 Step 1: Global Initialization

```text
Global Memory
-------------
a = 10
```

---

### 🔹 Step 2: `init()` Runs

Inside `init()`:

```go
fmt.Println(a)   // prints 10
a = 20
```

Now global state becomes:

```text
Global Memory
-------------
a = 20
```

📤 Output so far:

```
10
```

---

### 🔹 Step 3: `main()` Runs

```go
fmt.Println(a)
```

📤 Output:

```
20
```

---

### ✅ Final Output

```
10
20
```

---

## 8️⃣ Key Insight (VERY IMPORTANT)

> `init()` can **read and modify global variables**

And those changes are **visible to `main()`**.

This makes `init()` perfect for:

* Environment setup
* Configuration loading
* One-time preparation logic

---

## 9️⃣ Where `init()` Is Commonly Used (Real Projects)

You’ll see `init()` used in:

* Database connection setup
* Environment validation
* Logger initialization
* Registering plugins
* Config loading
* Dependency wiring

Example pattern:

```go
func init() {
    if os.Getenv("APP_ENV") == "" {
        panic("APP_ENV not set")
    }
}
```

---

## 🔟 Multiple `init()` Functions (Yes, Allowed!)

You can have:

* Multiple `init()` functions
* Across multiple files
* Across multiple packages

They run in:

```
imported packages → current package
file order → top to bottom
```

(We’ll go deep into this when packages are fully covered.)

---

## 1️⃣1️⃣ What `init()` Is NOT For ❌

Do **NOT** use `init()` for:

* Business logic
* Application flow
* Replacing `main()`
* Heavy processing
* Runtime decisions

Why?

Because:

* It’s hard to test
* Hard to control
* Hard to debug

---

## 1️⃣2️⃣ Mental Model (Remember Forever)

Think of `init()` like this:

```
🏗️ Setup phase (init)
🚀 Execution phase (main)
```

Or:

```
Global → init() → main()
```

---

## 1️⃣3️⃣ Interview One-Liner ✅

> **In Go, `init()` is a special function automatically executed after global initialization and before `main()`, used for one-time setup logic.**

---

## 1️⃣4️⃣ Common Interview Traps ⚠️

❓ Can `init()` take parameters?
➡ ❌ No

❓ Can `init()` return values?
➡ ❌ No

❓ Can we call `init()` manually?
➡ ❌ No

❓ Can it modify global variables?
➡ ✅ Yes

❓ Does it run before `main()`?
➡ ✅ Always

---

## 1️⃣5️⃣ Final Takeaways 🧠

* `init()` is **automatic**
* Order matters
* Globals → init → main
* Use for setup, not logic
* Powerful but dangerous if abused

