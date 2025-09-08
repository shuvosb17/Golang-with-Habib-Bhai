
## 📦 Structs and Functions: How to Print User Details

Let's see how Go lets you organize code to work with your own data types, and how memory is used when you call functions with structs!

---

## 👨‍💻 Your User Details Code

````go
package main

import "fmt"

// 🏗️ User struct: like a digital ID card
type User struct {
	name string
	age  int
}

// 🖨️ Regular function to print details (takes a User as input)
func PrintDetails(usr User) {
	fmt.Println("Name: ", usr.name)
	fmt.Println("Age : ", usr.age)
}

func main() {
	user1 := User{
		name: "Shuvo",
		age:  22,
	}

	user2 := User{
		name: "Tanvir",
		age:  22,
	}

	PrintDetails(user1) // 📤 Pass user1 to the function
	PrintDetails(user2) // 📤 Pass user2 to the function
}

func init() {
	fmt.Println("_____User Details______")
}
````

---

## 🏙️ Memory City for Functions and Structs

```
🏙️ MEMORY CITY
┌─────────────────────────────────────────────────────────────────────┐
│ 🏛️ CODE DISTRICT:                                                  │
│ • type User struct { ... } (ID card template)                       │
│ • func PrintDetails(), main(), init() (instructions)                │
├─────────────────────────────────────────────────────────────────────┤
│ 📚 STACK DISTRICT:                                                  │
│ • user1 and user2 cards (structs)                                   │
│ • Each PrintDetails() call gets a copy of the card                  │
│ • Local variables and parameters live here                          │
├─────────────────────────────────────────────────────────────────────┤
│ 🏗️ HEAP DISTRICT:                                                   │
│ • String contents ("Shuvo", "Tanvir")                               │
│ • Structs point to these strings                                    │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 🧩 How Does the Function Call Work?

### 1. **Pass By Value (Copying the Card)**

When you call `PrintDetails(user1)`:
- Go makes a **copy** of the `user1` card and gives it to the function.
- The function works with its own copy (`usr`), not the original.
- Any changes inside the function do **not** affect the original card.

### 2. **Stack Frame for Each Call**

- Each function call gets its own "workspace" (stack frame).
- When the function ends, this workspace is cleaned up.

### 3. **String Data on the Heap**

- The `name` field in the struct is a pointer to the actual string data on the heap.
- Both the original and the copy point to the same string content.

---

## 🧑‍🏫 Regular Function vs. Receiver Function

### Regular Function (your code):

```go
func PrintDetails(usr User) { ... }
PrintDetails(user1)
```
- You pass the struct as an argument.
- Function is separate from the struct.

### Receiver Function (method style):

```go
func (u User) PrintDetails() { ... }
user1.PrintDetails()
```
- Function is attached to the struct type.
- Syntax is more like "object-oriented" languages.

**Both ways:** The struct is copied when passed by value.

---

## 🧠 Key Concepts

- **Pass by value:** Go copies structs when passing to functions.
- **No side effects:** Changes inside the function don't affect the original struct.
- **Stack vs. Heap:** Structs live on the stack; string contents live on the heap.
- **Cleaner code:** Receiver functions (methods) make code more readable and organized.

---

## 🖨️ Program Output

```
_____User Details______
Name:  Shuvo
Age :  22
Name:  Tanvir
Age :  22
```

---

## 🧠 Beginner-Friendly Interview Questions

### Q1: What happens when you pass a struct to a function in Go?
**A:** Go makes a copy of the struct and gives it to the function (pass by value).

### Q2: What is a receiver function (method)?
**A:** It's a function attached to a struct type, letting you call it like `user1.PrintDetails()`.

### Q3: If you change the struct inside the function, does it affect the original?
**A:** No! The function only changes its own copy.

### Q4: Where is the string data stored?
**A:** The string content (like "Shuvo") is stored on the heap; the struct holds a pointer to it.

### Q5: Why use receiver functions?
**A:** They make your code more organized and easier to read, especially for larger programs.

---

## 🌟 Why This Matters

- 🗂️ **Organization:** Keep related code and data together.
- 🛡️ **Safety:** Avoid accidental changes to your original data.
- 🧩 **Scalability:** Methods make your code easier to grow and maintain.

Remember: Functions and structs in Go work together like a team—passing information safely and keeping your code neat and tidy! 🧑‍💻✨
