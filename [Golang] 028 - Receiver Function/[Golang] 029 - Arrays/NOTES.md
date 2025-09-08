
## 📦 Arrays in Go: Your Data Storage Boxes!

Arrays in Go are like **fixed-size boxes** 📦 that hold a set number of items of the same type. Let's see how they work, how they're stored in memory, and how you can use them!

---

## 🧑‍💻 Your Array Code (with Fixes)

````go
package main
import "fmt"

// 🗃️ Global array: always size 3, holds 3 strings
var arr2 = [3]string{"I", "Love", "You"}

func main() {
    // 📦 Array of 2 integers
    arr := [2]int{3, 6}
    fmt.Println(arr)        // [3 6]
    fmt.Println(arr2[1])    // Love

    // 📦 Array with size inferred (5 integers)
    arr3 := [...]int{10, 20, 30, 40, 50}
    fmt.Println("arr3:", arr3)

    // 📦 Array with specific indices set
    arr4 := [5]string{0: "Zero", 2: "Two", 4: "Four"}
    fmt.Println("arr4:", arr4) // [Zero  Two  Four]

    // 🛠️ Modify an array element
    arr[0] = 100
    fmt.Println("Modified arr:", arr) // [100 6]

    // 📦 Multi-dimensional array (array of arrays)
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println("Matrix:", matrix) // [[1 2 3] [4 5 6]]
}

func init() {
    fmt.Println("_____Array in Memory_____")
}
````

---

## 🏙️ Memory City for Arrays

```
🏙️ MEMORY CITY
┌─────────────────────────────────────────────────────────────┐
│ 🏛️ CODE DISTRICT:                                           
│ • func main(), func init()                                  │
│ • Array type definitions                                    │
├─────────────────────────────────────────────────────────────┤
│ 🏠 DATA DISTRICT:                                           
│ • arr2 = [3]string{"I", "Love", "You"}                     
│   (array headers here; string data on heap)                 │
├─────────────────────────────────────────────────────────────┤
│ 📚 STACK DISTRICT:                                          
│ • arr, arr3, arr4, matrix (local arrays in main)            │
│   (actual numbers/headers here)                             │
├─────────────────────────────────────────────────────────────┤
│ 🏗️ HEAP DISTRICT:                                           
│ • String contents: "I", "Love", "You", "Zero", "Two", "Four"│
└─────────────────────────────────────────────────────────────┘
```

---

## 🧩 How Arrays Work in Memory

- **Fixed Size:** The size is part of the type. `[2]int` is different from `[3]int`.
- **Contiguous Storage:** All elements are stored together in memory.
- **Value Type:** Assigning or passing an array copies the whole thing!
- **String Data:** For arrays of strings, the array holds pointers to the actual text, which lives on the heap.

---

## 🗂️ Array Examples Explained

| Example | What It Does | Memory |
|---------|--------------|--------|
| `arr := [2]int{3, 6}` | Box with 2 numbers | Stack |
| `arr2 = [3]string{"I", "Love", "You"}` | Box with 3 words | Data segment (headers), heap (text) |
| `arr3 := [...]int{10, 20, 30, 40, 50}` | Box with 5 numbers, size inferred | Stack |
| `arr4 := [5]string{0: "Zero", 2: "Two", 4: "Four"}` | Box with 5 slots, some filled | Stack (headers), heap (text) |
| `matrix := [2][3]int{...}` | 2 rows, 3 columns (like a table) | Stack |

---

## 🧠 Beginner-Friendly Interview Questions

### Q1: What is an array in Go?
**A:** A fixed-size box 📦 that holds items of the same type, like `[3]int` for 3 integers.

### Q2: Can you change the size of an array after creation?
**A:** No! Arrays are always the same size once created.

### Q3: What happens if you assign one array to another?
**A:** The entire array is copied (value type).

### Q4: Where is the actual string data stored for string arrays?
**A:** On the heap! The array holds pointers to the text.

### Q5: How do you access or change an array element?
**A:** Use an index: `arr[0] = 100` changes the first element.

### Q6: What is a multi-dimensional array?
**A:** An array of arrays, like a table: `[2][3]int` is 2 rows of 3 numbers.

---

## 🖨️ Program Output

```
_____Array in Memory_____
[3 6]
Love
arr3: [10 20 30 40 50]
arr4: [Zero  Two  Four]
Modified arr: [100 6]
Matrix: [[1 2 3] [4 5 6]]
```

---

## 🌟 Why Arrays Matter

- 📦 **Organized storage:** Keep related items together
- 🛡️ **Safe:** Fixed size prevents out-of-bounds errors
- 🚀 **Fast:** Contiguous memory for quick access
- 🧩 **Foundation:** Used to build more complex types (like slices)

Remember: Arrays are your **basic storage boxes** in Go—simple, fast, and reliable!
