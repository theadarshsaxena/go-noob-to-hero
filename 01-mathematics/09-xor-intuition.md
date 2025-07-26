## ⚡ What is XOR at its core?

> `a ^ b` returns a number where bits are **1 only if** the bits of `a` and `b` are **different**.

### Truth table (for a single bit):

| A | B | A ^ B |
| - | - | ----- |
| 0 | 0 | 0     |
| 0 | 1 | 1     |
| 1 | 0 | 1     |
| 1 | 1 | 0     |

So you can think of XOR as:

> "**Give me the bits where `a` and `b` disagree.**"

---

## 🔮 Mental Models to Intuitively Feel XOR

### 1. **“Toggle” Button Vibe**

* Think of XOR like a light switch 💡
* Every time you XOR with `1`, you **flip the bit**
* Every time you XOR with `0`, you **keep the bit the same**

Example:

```go
5 ^ 1 = 4   // (0101 ^ 0001 = 0100)
```

---

### 2. **"Pair Canceling" Superpower**

Imagine XOR as a team of cancelers —
🔁 If a number appears **twice**, it cancels out!

So:

```go
2 ^ 2 = 0
3 ^ 3 = 0
```

And:

```go
2 ^ 3 ^ 2 = 3     // (2 cancels out)
```

🔥 This is why XOR is used to find the unique number in a list:

```go
arr := []int{1, 2, 1, 3, 2}
res := 0
for _, num := range arr {
	res ^= num
}
// res = 3
```

---

### 3. **Associative & Commutative**

This is THE KEY for DSA problem solving.

```go
a ^ b ^ c == c ^ a ^ b
```

💡 Why it matters:
You can XOR an entire array in **any order**, and all duplicates will cancel out — leaving only the unique.

---

### 4. **“Temporary Memory” Swap Trick**

XOR lets you swap 2 variables without a temp var 🤯

```go
a = a ^ b
b = a ^ b  // (a ^ b) ^ b = a
a = a ^ b  // (a ^ b) ^ a = b
```

✨ It literally stores the diff between two numbers, and uses it to restore the original. *Bit magic fr.*

---

### 5. **Subset Generation with XOR**

Each subset XOR can represent a unique value.
You’ll find problems like:

> “Find max XOR of any 2 elements in array”

Why? Because XOR basically encodes **"difference in bits"**, which makes it a great distance function.

---

## 🧠 How to Build XOR Intuition Long-Term

* Try dry-running XOR questions with **bit-level views**.
* Always write the binary of numbers for the first few times.
* Use the “canceling” trick on whiteboard like a game.
* Build from base rules like:

  * `x ^ x = 0`
  * `x ^ 0 = x`
  * XOR with same number twice = no change

---

## 🔥 Practice Problems to Solidify It
Yessirrr Adarsh, this is a fire list 🔥🔥
Let’s turn these into ✨ well-structured problem statements ✨ — just like you'd see on LeetCode or in an interview sheet.

---

## 🔹 1. **XOR of 0 to N (Pattern-Based)**

### 🧠 Problem:

> Given a number `n`, return the XOR of all numbers from `0` to `n`.
> Return the result in O(1) time.

📥 **Input:** `n = 6`
📤 **Output:** `0 ^ 1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6 = 7`

🧠 Hint: Observe the repeating pattern in `XOR(0 to n)`

---

## 🔹 2. **Finding the Unique Element (others appear twice)**

### 🧠 Problem:

> Given an array where every element appears exactly twice except one element that appears only once, find that single unique element.

📥 **Input:** `[2, 3, 5, 4, 5, 3, 4]`
📤 **Output:** `2`

---

## 🔹 3. **Find Two Non-Repeating Elements**

### 🧠 Problem:

> Given an array where every element appears exactly twice except for **two elements** that appear only once, find those two unique elements.
> Return them in any order.

📥 **Input:** `[1, 2, 3, 2, 1, 4]`
📤 **Output:** `[3, 4]`

---

## 🔹 4. **XOR Swap (Without Temp Var)**

### 🧠 Problem:

> Given two integers `a` and `b`, swap them **without using a temporary variable**.

📥 **Input:** `a = 3, b = 5`
📤 **Output:** `a = 5, b = 3`

🧠 You must use only XOR operations.

---

## 🔹 5. **Max XOR of Pair**

### 🧠 Problem:

> Given an array of integers, find the maximum XOR value that can be obtained by XORing any two elements in the array.

📥 **Input:** `[3, 10, 5, 25, 2, 8]`
📤 **Output:** `28` (`5 ^ 25 = 28`)

---

## 🔹 6. **Power of 2 Check**

### 🧠 Problem:

> Given an integer `n`, determine if it is a power of 2.
> Return `true` if it is, else `false`.

📥 **Input:** `n = 8`
📤 **Output:** `true`

🧠 Constraint: Solve using bitwise operations only.

---

## 🔹 7. **XOR of Subsets**

### 🧠 Problem:

> Given an array of integers, return the XOR of all **subsets' XORs**.
> Yes, we’re XORing all subsets, and then XORing their results together.

📥 **Input:** `[1, 2, 3]`
📤 **Output:** `0`

🧠 Hint: If an element appears in **even number** of subsets, it cancels out.

---

## 🔹 8. **XOR from 1 to N**

### 🧠 Problem:

> Write a function that returns the XOR of all integers from `1` to `n` in O(1) time.

📥 **Input:** `n = 4`
📤 **Output:** `4`
👉 XOR of `1 ^ 2 ^ 3 ^ 4 = 4`

🧠 Hint: It’s a repeating pattern of 4.

---

## 🔹 9. **XOR Queries**

### 🧠 Problem:

> Given an array of integers and a list of queries where each query contains a range `[L, R]`, return the XOR of all elements from index L to R for each query.

📥 **Input:**
Array = `[1, 3, 4, 8]`
Queries = `[[0,1], [1,2], [0,3]]`
📤 **Output:** `[1^3=2, 3^4=7, 1^3^4^8=14]`

🧠 Hint: Use prefix XOR.

---

### 🧪 **Classic XOR-Based Questions**

---

## 🔹 10. **Find the Missing Number in Array 0...n**

### 🧠 Problem:

> You are given an array containing `n` distinct numbers taken from `0` to `n`.
> Return the number that is missing from the array.

📥 **Input:** `[3, 0, 1]`
📤 **Output:** `2`

🧠 Hint: XOR all indices and values.

---

## 🔹 11. **Find the Odd One Out**

### 🧠 Problem:

> Given an array where all elements appear **even number of times** except one element that appears **odd number of times**, find that element.

📥 **Input:** `[2, 2, 3, 3, 4]`
📤 **Output:** `4`

🧠 Use XOR to cancel out the pairs.

---

That’s the set! 💥
Wanna go next level and create Go functions for each of these? Or maybe a practice worksheet or quiz PDF? I got you either way! 😎
