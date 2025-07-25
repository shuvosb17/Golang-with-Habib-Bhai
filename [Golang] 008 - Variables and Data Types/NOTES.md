# 🚀Lecture 008: **Understanding Variables and Data Types in Go!**

---

## The Building Blocks of Data Storage

In this class, we will explore **how Go manages data using variables and types**!

## 1️⃣ What is a Variable? - Your Data Container 🏦

A **variable** is a **container** that holds data in memory (RAM).

| Component | Description | Real-world Analogy |
|-----------|-------------|-------------------|
| **CPU** | The Brain 🧠 | The chef in a kitchen |
| **RAM** | Temporary Workspace 📄 | Kitchen counter space |
| **Hard Disk** | Long-Term Storage 📀 | Pantry/refrigerator |

### 🔹 Go Variable Creation

````go
var a int = 10  // Creates a variable 'a' of type int and assigns it 10
````

### 🔹 What happens behind the scenes?

```
Memory Allocation Process:
┌─────────────────────────┐
│     RAM Memory          │
├─────────────────────────┤
│ Address: 0x001234       │
│ Variable Name: a        │ ◄── Label to access the data
│ Type: int               │ ◄── Defines data format
│ Value: 10               │ ◄── Actual stored data
└─────────────────────────┘
```

### 🔹 Bank Locker Analogy 🏦

| Bank Locker Element | Programming Equivalent |
|---------------------|------------------------|
| **Locker Number** | Variable name (`a`) |
| **Money Inside** | Data value (`10`) |
| **Locker Size** | Data type (`int`) |
| **Bank Vault** | RAM memory |

---

## 2️⃣ Data Types in Go - Different Data Containers 🔢🔠

Go has **different data types** to store various kinds of data efficiently.

### 🔹 Numeric Data Types

| Type | Size | Purpose | Example | Range |
|------|------|---------|---------|-------|
| `int` | 32/64-bit | Whole numbers | `var age int = 25` | -2³¹ to 2³¹-1 (32-bit) |
| `float32` | 32-bit | Decimal numbers (single precision) | `var price float32 = 19.99` | ~6-7 decimal places |
| `float64` | 64-bit | Decimal numbers (double precision) | `var pi float64 = 3.14159265` | ~15-16 decimal places |

````go
var a int = 10        // Integer - whole numbers
var b float32 = 2.5   // Float (single precision)
var c float64 = 3.141 // Float (double precision)
````

### 🔹 Precision Comparison

```
Single vs Double Precision:
┌─────────────────┬─────────────────┐
│ Single (32-bit) │ Double (64-bit) │
├─────────────────┼─────────────────┤
│ 2.1234567       │ 2.123456789012345│
│ Less memory     │ More memory     │
│ Faster compute  │ Higher accuracy │
│ Graphics/ML     │ Scientific calc │
└─────────────────┴─────────────────┘
```

### 🔹 Boolean Data Type

````go
var isGoFun bool = true // Boolean type
````

| Value | Memory Representation | Usage |
|-------|----------------------|-------|
| `true` | 1 | Condition is met |
| `false` | 0 | Condition is not met |

### 🔹 String Data Type

````go
var message string = "Hello, Go!"
````

| Characteristic | Description |
|----------------|-------------|
| **Purpose** | Store text values |
| **Encoding** | UTF-8 by default |
| **Immutable** | Cannot be changed after creation |

### 🔹 Real-world Data Type Examples

| Data Type | Real-world Use Case | Example |
|-----------|-------------------|---------|
| **int** | Counting items 🍏 | Number of apples in a basket |
| **float** | Measurements 📏 | Height in meters (1.75m) |
| **bool** | Status checks 💡 | Light switch (ON/OFF) |
| **string** | Text data 🛂 | Name on a passport |

---

## 3️⃣ Declaring Variables in Go - Multiple Ways 📝

### 🔹 Standard Declaration (Explicit Type)

````go
var a int = 10
````

### 🔹 Type Inference (Shorter Syntax)

````go
a := 10        // Go detects 'int'
b := 2.10      // Go detects 'float64'
c := "Hello!"  // Go detects 'string'
d := true      // Go detects 'bool'
````

### 🔹 Declaration Methods Comparison

| Method | Syntax | When to Use |
|--------|--------|-------------|
| **Explicit** | `var name type = value` | When type clarity is important |
| **Inference** | `name := value` | For concise, readable code |
| **Zero Value** | `var name type` | When you need default values |

### 🔹 Phone Purchase Analogy 📱

```
Variable Declaration = Phone Purchase:
┌─────────────────────────────────────┐
│ Explicit Type Declaration           │
│ var phone iPhone13 = newPhone()     │ ◄── You specify exact model
└─────────────────────────────────────┘
                  vs
┌─────────────────────────────────────┐
│ Type Inference                      │
│ phone := newPhone()                 │ ◄── Go figures out the model
└─────────────────────────────────────┘
```

---

## 4️⃣ Constants - Unchangeable Values 🔒

A **constant (`const`)** is a **fixed value** that cannot be modified after declaration.

### 🔹 Constant Declaration

````go
const pi = 3.14159 // Mathematical constant
const maxUsers = 100 // Application limit
````

### 🔹 Constants vs Variables

| Aspect | Constants | Variables |
|--------|-----------|-----------|
| **Mutability** | ❌ Cannot change | ✅ Can change |
| **Memory** | Compile-time allocation | Runtime allocation |
| **Use Case** | Fixed values (PI, limits) | Dynamic data |

### 🔹 Attempting to Change Constants

````go
const p = 100 // Defines 'p' as a constant
p = 10       // ❌ ERROR! Constants cannot be changed
````

### 🔹 Real-world Constant Examples

| Constant Type | Example | Why Constant? |
|---------------|---------|---------------|
| **Physical Laws** | Speed of light (299,792,458 m/s) 🌟 | Universal constants |
| **Personal Data** | Date of birth 🎂 | Never changes |
| **App Config** | Maximum file size 📁 | Business rules |

---

## 🔄 Variable Lifecycle

```
Variable Lifecycle:
┌─────────────────┐
│   Declaration   │ ◄── var a int OR a := 10
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Initialization │ ◄── Assign initial value
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Usage       │ ◄── Read/modify throughout program
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Scope Ends     │ ◄── Memory automatically freed
└─────────────────┘
```

## 🚀 Final Summary

| Concept | Symbol | Purpose | Key Takeaway |
|---------|--------|---------|--------------|
| **Variables** | 🏦 | Store data temporarily in RAM | Containers for your data |
| **Data Types** | 🔢🔠 | Define data format and operations | int, float, bool, string |
| **Declaration** | 📝 | Create variables (`var` or `:=`) | Explicit vs inference |
| **Constants** | 🔒 | Fixed values that never change | Immutable data |

### 🔹 Best Practices

1. **Use meaningful variable names** - `userAge` instead of `a`
2. **Choose appropriate data types** - `int` for counts, `float64` for precision
3. **Use constants for fixed values** - Configuration limits, mathematical constants
4. **Prefer type inference** - Cleaner, more readable code

Remember: Variables are the foundation of data manipulation in Go. Understanding them well will make everything else easier!
