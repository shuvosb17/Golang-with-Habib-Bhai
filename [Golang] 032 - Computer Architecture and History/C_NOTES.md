## 🌱 Why This Lecture Exists (Big Picture)

Most people learn:

> “Write code → run program → get output”

But **real engineers** ask:

* *What happens inside the machine when I run code?*
* *Why does Go behave this way on my system?*
* *Why does memory overflow happen?*
* *Why does OS scheduling matter?*

👉 To answer these, you must understand **computer architecture**.

---

# 🧩 What Does “Architecture” Mean?

### 🧍 Human Architecture (Analogy)

A human has:

* Brain
* Heart
* Lungs
* Kidneys
* Bones

Together → **a human**

### 🖥️ Computer Architecture

A computer has:

* **CPU** (brain)
* **RAM** (short-term memory)
* **Hard Disk** (long-term memory)

Together → **a computer**

👉 Architecture = **what parts exist and how they work together**

---

## 🧠 What Is a Computer REALLY?

A computer is **NOT**:

* Keyboard ❌
* Mouse ❌
* Monitor ❌

Those are **input/output devices**.

### ✅ A computer is:

> **A machine that can calculate**

So the **real computer** is:

* **CPU**
* **RAM**
* **Storage**

Everything else is optional.

---

## 🧱 Core Components (Minimal View)

```
+-------------------+
|       CPU         |
+-------------------+
        ↑ ↓
+-------------------+
|       RAM         |
+-------------------+
        ↑ ↓
+-------------------+
|   Hard Disk       |
+-------------------+
```

---

## 💾 RAM vs Hard Disk (CRUCIAL DIFFERENCE)

### RAM

* Has **cells** (small boxes)
* Each cell stores binary data
* **Volatile** → loses data when power is off
* Very fast ⚡

### Hard Disk

* Also has cells
* **Non-volatile** → data survives power off
* Slower but cheap & huge 📦

| Feature     | RAM            | Hard Disk |
| ----------- | -------------- | --------- |
| Speed       | Very fast      | Slow      |
| Memory type | Volatile       | Permanent |
| Cost        | Expensive      | Cheap     |
| Role        | Working memory | Storage   |

---

## ⚡ Binary Reality: 0 and 1

Computers understand **only**:

* `0` → no electricity
* `1` → electricity present

That’s it.
Everything else is **illusion built on top**.

---

## 🧮 CPU: The Simplest Genius Ever

### CPU can do ONLY these operations:

1. Addition (+)
2. Subtraction (−)
3. Multiplication (×)
4. Division (÷)
5. AND
6. OR
7. NOT

👉 **Just 7 operations**.

And yet:

* Videos
* Games
* AI
* Internet
* Go programs
* Space rockets 🚀

🔥 This is the ART of computer science.

---

## 🧬 Inside the CPU

CPU has **two main parts**:

### 1️⃣ Processing Unit

* Performs math & logic
* Executes instructions

### 2️⃣ Register Set

* Tiny, ultra-fast memory
* Holds:

  * Data
  * Instructions
  * Addresses

```
CPU
├── Processing Unit (ALU)
└── Register Set
    ├── R1
    ├── R2
    ├── ...
    └── Instruction Pointer
```

---

## 👉 Instruction Pointer (Very Important)

Also called:

* Program Counter
* Pointing Register

### What it does:

* Stores **which memory cell to execute next**

Example:

```
IP = 0 → execute instruction at RAM[0]
IP = 1 → execute instruction at RAM[1]
```

This idea came from **very early computers** and still exists today.

---

# 📜 History: How Computers Were Born

Understanding history = understanding **why things are the way they are**.

---

## 🧮 Abacus (2700 BC)

![Image](https://images.computerhistory.org/revonline/images/xb93.80p-03-01.jpg?w=600)

![Image](https://home.csulb.edu/~cwallis/labs/computability/index.1a.gif)

* First counting machine
* Beads moved by hand
* Could **count**
* That alone qualifies it as a *computer*

---

## 🧠 Binary Logic (1600s)

Introduced by **Gottfried Wilhelm Leibniz**

* Invented **binary logic**
* Foundation of:

  * AND
  * OR
  * NOT

Without him:

* No modern computers
* No Go
* No internet

---

## ⚙️ First Programmable Computer (1830s)

### **Charles Babbage**

![Image](https://computerhistory.org/wp-content/uploads/2019/07/engine-structure.jpg)

![Image](https://coimages.sciencemuseumgroup.org.uk/370/611/medium_1905_0181_0052__0008_.jpg)

* Invented **Analytical Engine**
* Mechanical (no electricity)
* Used:

  * Gears
  * Steam power
  * Hand cranks

### Input & Output:

* **Punch cards**
* Holes = 1
* No hole = 0

---

## 🧑‍💻 First Programmer Ever

### **Ada Lovelace**

![Image](https://coimages.sciencemuseumgroup.org.uk/669/182/medium_smg00000912.jpg)

![Image](https://c02.purpledshub.com/uploads/sites/41/2018/08/47-Magdalen-College-Libraries-and-Archives-Daubeny-90.A.11-fc247cf.jpg?w=1200\&webp=1)

* Wrote first algorithm (1843)
* For Analytical Engine
* Calculated **Bernoulli Numbers**

👉 Programming existed **before electricity** 🤯

---

## 🧾 Punch Cards (Programming Hell)

![Image](https://upload.wikimedia.org/wikipedia/commons/thumb/2/26/Punched_card_program_deck.agr.jpg/250px-Punched_card_program_deck.agr.jpg)

![Image](https://www.ibm.com/adobe/dynamicmedia/deliver/dm-aid--af828863-76fb-4343-b39c-06dd286c7d37/punched-card-271.jpg?preferwebp=true)

* One instruction = one card
* 10,000 lines = 10,000 cards
* One mistake → entire program useless

Respect early programmers 🫡

---

## ⚡ Electronic Revolution (1945)

### **ENIAC**

![Image](https://www.simslifecycle.com/wp-content/uploads/sites/2/2022/01/Electronic-Numerical-Integrator-And-Computer.png)

![Image](https://upload.wikimedia.org/wikipedia/commons/thumb/0/01/Atanasoff-Berry_Computer_at_Durhum_Center.jpg/250px-Atanasoff-Berry_Computer_at_Durhum_Center.jpg)

* First electronic general-purpose computer
* Used **vacuum tubes**
* 17,468 tubes 😱
* Consumed massive power
* ~5000 operations/sec

Still revolutionary.

---

## 🔬 Transistor Revolution (1950s)

![Image](https://www.tubeampdoctor.com/magazin/wp-content/uploads/2023/01/tube_vs_transistor8.jpg)

![Image](https://upload.wikimedia.org/wikipedia/commons/b/b5/TRADIC_computer.jpg)

* Based on **quantum mechanics**
* Smaller
* Faster
* Less power

Replaced vacuum tubes.

---

## 🧠 Integrated Circuits (IC)

* Many transistors in one chip
* Enabled:

  * Smaller computers
  * Reliable systems

---

## 🖥️ Personal Computer Era

### **Intel**

* Intel 4004 (1971)
* 4-bit CPU

### **Apple**

* Popularized personal computers (1976)

### **IBM**

* IBM PC (1981)
* 16-bit architecture

---

## 📏 Bits, Bytes & 32/64-bit Truth

### Definitions:

* **1 bit** = 0 or 1
* **8 bits** = 1 byte
* **1024 bytes** = 1 KB
* **1024 KB** = 1 MB
* **1024 MB** = 1 GB

---

### 32-bit System

* Registers hold 32 bits
* Max addressable memory ≈ **4 GB**

### 64-bit System

* Registers hold 64 bits
* Address space ≈ **18 quintillion bytes**

👉 That’s why modern OS + Go prefer **64-bit**

---

## 🧠 Final Mental Model (LOCK THIS 🔐)

```
Program (Go code)
   ↓
Binary Instructions
   ↓
Stored in RAM cells
   ↓
Instruction Pointer selects next cell
   ↓
CPU registers load data
   ↓
ALU executes
   ↓
Result stored back
```

---

## 🎯 Why This Matters for Go & OS

Because:

* Goroutines depend on CPU scheduling
* Memory allocation depends on RAM architecture
* Garbage collector depends on address space
* 32/64-bit affects pointer size
* Performance depends on cache & registers

---

## 🌟 Final Words

This lecture is not about:
❌ memorizing dates
❌ passing exams

It is about:
✅ **seeing inside the machine**
✅ **feeling how code becomes reality**
✅ **becoming a real engineer**

