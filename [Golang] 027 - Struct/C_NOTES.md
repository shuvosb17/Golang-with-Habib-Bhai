# 1️⃣ Why Do We Need `struct`?

Before `struct`, we had **single values**:

```go
var age int = 30
var name string = "Habib"
```

Problem ❌
These values are **related**, but stored **separately**.

We want:

> “One thing that represents a User”

That’s where `struct` comes in.

---

# 2️⃣ Struct = Custom Data Type (Very Important)

In Go:

```go
type User struct {
    Name string
    Age  int
}
```

### What did we just do?

We **created our own type**, just like:

* `int`
* `string`
* `float64`

But **custom-made**.

📌 **Key idea**

> `struct` is a **blueprint**, NOT a value.

---

# 3️⃣ Where Does a `struct` Definition Live in Memory?

This is where most people get confused 👇

### At **compile time**:

* `User` is a **type definition**
* It is **read-only**
* It lives in the **Code Segment**

```
CODE SEGMENT (read-only)
-----------------------
type User {
    Name string
    Age  int
}
```

📌 **Important**

* The struct definition is **NOT memory for data**
* It only tells Go:

  > “If someone creates a User, it must have these fields”

---

# 4️⃣ Creating a Struct Value (Instance / Object)

Now we do:

```go
var user1 User
```

### What happens?

* `user1` is a **variable**
* Its type is `User`
* Memory is allocated for:

  * `Name`
  * `Age`

📌 This is called:

* **Instance**
* or **Object** (OOP term)

---

# 5️⃣ Struct Instance Lives WHERE?

Inside a function (like `main`):

```go
func main() {
    var user1 User
}
```

Memory allocation happens in the **stack**:

```
STACK (main frame)
------------------
user1:
 ├── Name (string)
 └── Age  (int)
```

📌 Because:

* `user1` is local
* Not returned
* Not shared

---

# 6️⃣ Initializing a Struct (Assigning Values)

```go
user1 = User{
    Name: "Habib",
    Age:  30,
}
```

What actually gets stored?

❌ NOT the type
❌ NOT the definition

✅ Only the **values**

```
STACK (main frame)
------------------
user1:
 ├── Name → "Habib"
 └── Age  → 30
```

📌 The **struct definition stays in code segment forever**
📌 Each instance gets **its own memory**

---

# 7️⃣ Accessing Struct Fields (`.` operator)

```go
fmt.Println(user1.Name)
fmt.Println(user1.Age)
```

Execution steps:

1. Find `user1` in stack
2. Look inside its memory
3. Read the requested field

📌 `.` means:

> “Go inside this value”

---

# 8️⃣ Creating Multiple Instances (Very Important Concept)

```go
user2 := User{
    Name: "Rocky",
    Age:  16,
}
```

Now memory looks like:

```
STACK (main frame)
------------------
user1:
 ├── Name → "Habib"
 └── Age  → 30

user2:
 ├── Name → "Rocky"
 └── Age  → 16
```

📌 **They do NOT share memory**
📌 Same type ≠ same value

---

# 9️⃣ Most Common Beginner Mistake (DESTROY THIS CONFUSION ❌)

❌ Wrong thinking:

> “Struct type holds the data”

✅ Correct thinking:

> “Struct type only defines the shape”

### Analogy 🍽️

* Type = plate design
* Instance = food on the plate
* Same plate design, different food

---

# 🔁 Comparison with Basic Types

```go
var a int = 10
var b int = 20
```

Even though both are `int`:

* `a` has its own memory
* `b` has its own memory

Same rule applies to structs.

---

# 1️⃣0️⃣ Struct Fields Terminology

Inside a struct:

```go
type User struct {
    Name string
    Age  int
}
```

Fields are called:

* **Fields**
* **Properties**
* **Member variables**

All mean the same thing in Go context.

---

# 1️⃣1️⃣ Struct Creation Terms (Interview Gold)

| Term        | Meaning                     |
| ----------- | --------------------------- |
| Type        | Blueprint (`User`)          |
| Instance    | Actual value (`user1`)      |
| Instantiate | Creating an instance        |
| Field       | Name, Age                   |
| Object      | Same as instance (OOP term) |

📌 In Go interviews:

* “Instance” is safer than “object”

---

# 1️⃣2️⃣ Full Memory Flow Summary (This Is CRITICAL)

### Compile Time

```
CODE SEGMENT
------------
User type definition
main() function
```

### Run Time

```
STACK (main frame)
-----------------
user1 → {Habib, 30}
user2 → {Rocky, 16}
```

### End of Program

* Stack cleared
* Instances destroyed
* Program exits

---

# 1️⃣3️⃣ What Struct Is NOT ❌

* ❌ Not global memory
* ❌ Not shared automatically
* ❌ Not stored in code segment
* ❌ Not magical

Struct is **simple + predictable**.

---

# 1️⃣4️⃣ Why Structs Matter in Real Projects

Structs are everywhere:

* HTTP request/response
* Database models
* JSON decoding
* Configuration
* Domain models

If structs are weak → everything breaks.

---

# 1️⃣5️⃣ Mental Model (Keep This Forever 🧠)

> **Struct = Custom container**
> **Type = blueprint**
> **Instance = real container with values**
> **Each instance = separate memory**

If you remember ONLY this → you’ll never be confused again.

---

## 🚀 What Comes Next?

Next class: **Receiver Functions (Methods)**
That’s where:

* Struct + behavior combine
* Go’s OOP-style design appears
* Pointer vs value receivers become crucial

