# 🚀 **[Golang] 027 - Struct**

---

## Building Your First User Card! 🎴

Think of a struct as a **digital index card** or **ID card** where you can store related information about something. Let's learn how to create and use these special cards in Go!

## 🆔 Your User Card Code

````go
package main

import "fmt"

// 🏗️ Creating the blueprint for a User card
type User struct {
    name string  // 📝 Name field (like writing space for name)
    age  int     // 🎂 Age field (like writing space for age)
}

func main() {
    // 📄 Method 1: Create blank card, then fill it
    var user1 User  // Creates empty card with default values

    user1 = User{   // Fill the blank card with information
        name: "Shuvo",  // ✏️ Write name
        age:  22,       // ✏️ Write age
    }
    
    fmt.Println(user1.name) // 📖 Read the name
    fmt.Println(user1.age)  // 📖 Read the age

    // 📄 Method 2: Create and fill card in one step
    user2 := User{
        name: "Tanvir", // ✏️ Write name immediately
        age:  23,       // ✏️ Write age immediately
    }
    
    fmt.Println(user2.name) // 📖 Read the name
    fmt.Println(user2.age)  // 📖 Read the age
}

func init() {
    fmt.Println("_______User Prototype_____")
}
````

---

## 🎴 What Is a Struct? (Simple Explanation)

A **struct** is like a **custom container** or **filing folder** that holds related information together:

```
🗂️ STRUCT = CUSTOM CONTAINER
┌─────────────────────────────────────────────────────────────────────┐
│ Think of it like:                                                   
│ 🎴 ID Card - with spaces for name, age, photo                      │
│ 📋 Form - with fields to fill out                                  │
│ 🗃️ File Folder - with sections for different documents             │
│ 📦 Package - with compartments for different items                 │
│                                                                     
│ A struct groups related data together! 👥                          │
└─────────────────────────────────────────────────────────────────────┘
```

### 🔍 Real-Life Analogy: School ID Card

```
🎓 SCHOOL ID CARD (User Struct)
┌─────────────────────────────────────────────────────────────────────┐
│                    STUDENT ID CARD                                  │
├─────────────────────────────────────────────────────────────────────┤
│ 📝 Name: [Shuvo        ]  ← name string field                      
│ 🎂 Age:  [22          ]  ← age int field                           
│ 📸 Photo: [    😊     ]  ← (we could add more fields!)             │
├─────────────────────────────────────────────────────────────────────┤
│ Why together? Because they all describe the SAME person! 👤         
└─────────────────────────────────────────────────────────────────────┘
```

---

## 🏙️ Memory City for Structs

Let's see how structs live in our memory city:

```
🏙️ MEMORY CITY FOR STRUCTS
┌─────────────────────────────────────────────────────────────────────┐
│ 🏛️ CODE DISTRICT (Blueprint Storage)                                │
│ • type User struct { ... } (the ID card template) 📋                │
│ • func main(), func init() (instruction manuals) 📚                 
├─────────────────────────────────────────────────────────────────────┤
│ 🏠 DATA DISTRICT (Global Storage)                                   │
│ • (empty in this example) 🏠                                        
├─────────────────────────────────────────────────────────────────────┤
│ 📚 STACK DISTRICT (Temporary Desk Space)                            │
│ • 🎴 user1 card: name="Shuvo", age=22                               │
│ • 🎴 user2 card: name="Tanvir", age=23                              │
│ • Both cards are sitting on the desk! 📋                            │
├─────────────────────────────────────────────────────────────────────┤
│ 🏗️ HEAP DISTRICT (String Storage)                                  
│ • "Shuvo" text (actual letters) 📝                                  │
│ • "Tanvir" text (actual letters) 📝                                 │
│ • String contents live here! 📄                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 🎬 Step-by-Step Card Creation Journey

### 🛡️ Step 1: Office Opens (init)

````go
func init() {
    fmt.Println("_______User Prototype_____")
}
````

```
📢 OUTPUT: "_______User Prototype_____"
🏢 User card creation office is now open!
```

### 🏢 Step 2: Main Office Starts Work

````go
func main() {
    // Starting card creation process...
}
````

```
📚 STACK DISTRICT - Floor 1
┌─────────────────────────────────────────────────────────────────────┐
│ 🏢 main() Office                                                    │
│ • Task: Create user cards                                           
│ • Desk space ready for card creation! 📋                            │
└─────────────────────────────────────────────────────────────────────┘
```

### 📄 Step 3: Create First Card (Blank Method)

````go
var user1 User  // Create blank card
````

```
📚 STACK DISTRICT - User1 Card Created
┌─────────────────────────────────────────────────────────────────────┐
│ 🎴 user1 Card (Blank)                                               
│ ┌─────────────────────────────────────────────────────────────────┐ 
│ │ 📝 Name: [        ] (empty string "")                          │ │
│ │ 🎂 Age:  [   0    ] (zero value)                               │ │
│ └─────────────────────────────────────────────────────────────────┘ 
│ Status: Blank card ready for filling! ✏️                           │
└─────────────────────────────────────────────────────────────────────┘
```

### ✏️ Step 4: Fill First Card

````go
user1 = User{
    name: "Shuvo",
    age:  22,
}
````

```
📚 STACK DISTRICT - User1 Card Filled
┌─────────────────────────────────────────────────────────────────────┐
│ 🎴 user1 Card (Completed)                                           
│ ┌─────────────────────────────────────────────────────────────────┐ │
│ │ 📝 Name: [Shuvo  ] → points to heap 🏗️                        │ │
│ │ 🎂 Age:  [  22   ] → stored directly 💯                       │ │
│ └─────────────────────────────────────────────────────────────────┘ │
│ Status: Card completely filled! ✅                                   
└─────────────────────────────────────────────────────────────────────┘

🏗️ HEAP DISTRICT - String Storage
┌─────────────────────────────────────────────────────────────────────┐
│ 📝 "Shuvo" (actual letters stored here)                             │
│ • user1.name points to this location 👉                             │
└─────────────────────────────────────────────────────────────────────┘
```

### 📖 Step 5: Read First Card

````go
fmt.Println(user1.name) // Read name
fmt.Println(user1.age)  // Read age
````

```
📖 READING user1 Card:
• Look at name field → follow pointer to heap → find "Shuvo"
📢 OUTPUT: "Shuvo"

• Look at age field → read directly from card → find 22
📢 OUTPUT: "22"
```

### 🎴 Step 6: Create Second Card (Quick Method)

````go
user2 := User{
    name: "Tanvir",
    age:  23,
}
````

```
📚 STACK DISTRICT - Now Has Two Cards!
┌─────────────────────────────────────────────────────────────────────┐
│ 🎴 user1 Card                 │ 🎴 user2 Card                       │
│ ┌─────────────────────────────┐ │ ┌─────────────────────────────────┐ │
│ │ 📝 Name: [Shuvo  ]          │ │ │ 📝 Name: [Tanvir ]              │ │
│ │ 🎂 Age:  [  22   ]          │ │ │ 🎂 Age:  [  23   ]              │ │
│ └─────────────────────────────┘ │ └─────────────────────────────────┘ │
│ Status: ✅ Complete             │ Status: ✅ Complete               │
└─────────────────────────────────────────────────────────────────────┘

🏗️ HEAP DISTRICT - String Storage Updated
┌─────────────────────────────────────────────────────────────────────┐
│ 📝 "Shuvo" ← user1.name points here                                 │
│ 📝 "Tanvir" ← user2.name points here                                │
│ Both strings stored separately! 🔐                                  │
└─────────────────────────────────────────────────────────────────────┘
```

### 📖 Step 7: Read Second Card

````go
fmt.Println(user2.name) // Read name
fmt.Println(user2.age)  // Read age
````

```
📖 READING user2 Card:
• Look at name field → follow pointer to heap → find "Tanvir"
📢 OUTPUT: "Tanvir"

• Look at age field → read directly from card → find 23
📢 OUTPUT: "23"
```

---

## 🔍 Key Struct Concepts Explained Simply

### 🎯 1. **Two Ways to Create Cards**

| Method | Code | Description |
|--------|------|-------------|
| **📄 Blank then Fill** | `var user1 User` <br> `user1 = User{...}` | Create empty card first, fill later |
| **🎴 Create and Fill** | `user2 := User{...}` | Create and fill in one step |

### 🎯 2. **Where Data Lives**

```
📍 DATA LOCATION MAP
┌─────────────────────────────────────────────────────────────────────┐
│ Struct Card Itself:                                                 │
│ 📚 STACK DISTRICT (main's desk)                                     
│ • user1 and user2 cards sit here                                    │
│ • Contains: age numbers + pointers to names                         │
│                                                                     │
│ String Text Content:                                                │
│ 🏗️ HEAP DISTRICT (text storage)                                     
│ • "Shuvo" and "Tanvir" stored here                                  │
│ • Cards point to these locations                                    │
└─────────────────────────────────────────────────────────────────────┘
```

### 🎯 3. **Field Access with Dots**

```
🔍 DOT NOTATION = CARD READING
┌─────────────────────────────────────────────────────────────────────┐
│ user1.name  →  "Look at user1 card, find name field"               │
│ user1.age   →  "Look at user1 card, find age field"                │
│                                                                     
│ Think of the dot (.) as "look at" or "go to"! 👀                   │
└─────────────────────────────────────────────────────────────────────┘
```

### 🎯 4. **Zero Values (Default Values)**

```
📋 BLANK CARD DEFAULT VALUES
┌─────────────────────────────────────────────────────────────────────┐
│ When you create var user1 User:                                     
│ • name string → "" (empty string) 📝                                │
│ • age int     → 0 (zero number) 🔢                                  │
│                                                                     
│ Go automatically fills blanks with safe default values! 🛡️          │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 🎯 Complete Program Output

```
_______User Prototype_____  ← 🛡️ Office opens
Shuvo                      ← 📖 Reading user1.name
22                         ← 📖 Reading user1.age
Tanvir                     ← 📖 Reading user2.name
23                         ← 📖 Reading user2.age
```

---

## 🧠 Beginner-Friendly Interview Questions

### Q1: What is a struct in simple terms?
**A:** A struct is like an **ID card** 🎴 or **form** 📋 that groups related information together, like putting someone's name and age on the same card.

### Q2: Why use structs instead of separate variables?
**A:** Because related information belongs together! 👥 It's like keeping all your school documents in one folder instead of scattered everywhere.

### Q3: What does the dot (.) in user1.name mean?
**A:** The dot means **"look at"** 👀 - so `user1.name` means "look at user1's name field."

### Q4: What are zero values in structs?
**A:** When you create an empty struct, Go fills it with **safe default values** 🛡️ - empty string `""` for text, `0` for numbers.

### Q5: Where do structs live in memory?
**A:** The struct cards live on the 📚 **Stack** (your desk), but string contents live in the 🏗️ **Heap** (text storage).

### Q6: Can you change struct values after creation?
**A:** Yes! ✅ You can change them like `user1.age = 25` - just like erasing and rewriting on your card.

### Q7: What's the difference between the two creation methods?
**A:** 
- **Method 1**: `var user1 User` → Create blank card first 📄
- **Method 2**: `user2 := User{...}` → Create and fill immediately 🎴

### Q8: Can structs have different types of fields?
**A:** Absolutely! ✅ You can mix strings, numbers, booleans - like having name (text), age (number), and married (true/false) on the same card.

---

## 🎪 Fun Real-Life Struct Examples

### 1. 🐕 Pet Information Card
````go
type Pet struct {
    name    string
    species string  // "dog", "cat", etc.
    age     int
    hungry  bool    // true or false
}
````

### 2. 📱 Phone Contact Card
````go
type Contact struct {
    name   string
    phone  string
    email  string
    friend bool
}
````

### 3. 🎮 Game Character Card
````go
type Character struct {
    name   string
    health int
    level  int
    alive  bool
}
````

### 4. 📚 Book Information Card
````go
type Book struct {
    title  string
    author string
    pages  int
    read   bool
}
````

---

## 🌟 Why Structs Are Awesome

1. 🗂️ **Organization**: Keep related data together
2. 📋 **Structure**: Create your own custom data types
3. 🔍 **Easy Access**: Use dots to read/write fields
4. 🛡️ **Safety**: Automatic zero values prevent errors
5. 🎯 **Clarity**: Code becomes more readable and meaningful
6. 🏗️ **Building Blocks**: Foundation for more complex programs

Remember: Structs are like creating your own **custom containers** 📦 to organize information in a way that makes sense for your program! They're the building blocks for creating organized, readable code! 🏗️✨
