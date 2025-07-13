## 🧩 Math Operations DSA Roadmap (in Go)

### ✅ 1. **Basic Number Theory**

> Start with the fundamentals — trust me, everything builds on this.
> 
- Prime Numbers & Sieve of Eratosthenes
- Divisibility rules
- Factors & Divisors
- Count of Divisors
- Sum of Divisors
- GCD (Euclidean Algorithm)
- LCM
- Co-prime check

📌 *Go Practice*: Write functions for GCD, LCM, prime check, factorize a number.

---

### ✅ 2. **Bit Manipulation Basics**

> This is where XOR fun begins! Super useful for coding rounds.
> 
- Binary Representation
- Bitwise Operators (`&`, `|`, `^`, `<<`, `>>`)
- Checking Even/Odd using Bits
- Set/Unset/Toggle a bit
- Counting set bits (`Brian Kernighan's Algo`)
- XOR Basics

📌 *Go Practice*:

```go
func isBitSet(n, pos int) bool {
	return (n & (1 << pos)) != 0
}

```

---

### ✅ 3. **XOR Trick Questions**

> XOR is 🔥 for sneaky interview tricks. Some of the most clever problems lie here.
> 
- XOR of 0 to N (pattern-based)
- Finding the unique element (when others appear twice)
- Find two non-repeating elements
- XOR swap (without temp var)
- XOR of subsets

🧪 *Classic Qs*:

- Find the missing number in array 0...n
- Find the odd one out

---

### ✅ 4. **Fast Exponentiation**

> Helps in time-sensitive problems, especially when dealing with large exponents.
> 
- Recursion-based exponentiation
- Binary exponentiation
- Modular Exponentiation

📌 *Use Case*: `a^b % mod` for large a, b

---

### ✅ 5. **Modular Arithmetic**

> 💯 necessary for competitive programming
> 
- Why modulo is needed
- Properties:
    
    `(a + b) % m = (a % m + b % m) % m`
    
    `(a * b) % m = (a % m * b % m) % m`
    
- Modular Inverse (for division mod)
- Fermat’s Little Theorem (for primes)

🧪 *Classic Qs*:

- Ways to reach stair (DP + mod)
- Combinatorics with mod (nCr % mod)

---

### ✅ 6. **Mathematical Patterns & Digit Tricks**

> Very common in math-heavy questions.
> 
- Sum of digits
- Digital root
- Reverse number
- Palindrome number
- Armstrong number
- Base conversions

---

### ✅ 7. **Combinatorics**

> High yield topics in problems like permutations, paths, arrangements
> 
- Factorials
- nCr (using Pascal's triangle, mod tricks)
- Permutations vs Combinations
- Inclusion-Exclusion Principle
- Catalan Numbers (if you go deep)

---

### ✅ 8. **Math + Binary Search Hybrid Problems**

> Yeah, math meets BS here — super common pattern.
> 
- Min time to finish task (LC style)
- Allocate pages, painters partition
- Aggressive cows 🐄 problem (classic!)

---

### ✅ 9. **Prime Factorization**

> Next-level stuff after basics.
> 
- Prime factorization using trial division
- Using sieve for factorization
- Count of distinct prime factors

---

### ✅ 10. **Number Theory Advanced**

> Not necessary for all interviews but epic for CP and deeper DSA
> 
- Euler’s Totient Function
- Sieve with SPF (Smallest Prime Factor)
- Extended Euclidean Algorithm
- Chinese Remainder Theorem (CRT)

---

## 🧠 Bonus: Problem-Solving Tips

- Use Go’s `math/big` when numbers go crazy large.
- Don't ignore recursion + mod + bitmask patterns.
- Learn to **derive patterns** (XOR of 1 to n etc.)
- Keep cheat sheets of mod tricks & binary operations

---

## ✨ TL;DR Learning Plan

| Phase | Topics |
| --- | --- |
| 🟢 Starter | GCD/LCM, Primes, Divisors |
| 🔵 Core | Bits, XOR, Fast Power, Modulo |
| 🟣 Advanced | Combinatorics, CRT, Euler's, Binary Search Hybrid |