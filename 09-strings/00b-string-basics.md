## String Basics in Go: Masterclass

Let’s break down **everything you must know** about strings in Go, from the very foundation to the nuanced details that pros use. Each concept includes simple explanations and well-commented code samples so nothing feels mysterious.

### 1. **String Declaration and Initialization**

- **String** is a sequence of bytes (typically, but not always, representing characters).
- Declaring strings is straightforward.

```go
// Declaration without initialization (default is empty string)
var name string
fmt.Println(name) // Output: ""

// Declaration with initialization
var greeting string = "Hello, Go!"

// Short-hand declaration (common in Go)
city := "Mumbai"

fmt.Println(greeting) // Output: Hello, Go!
fmt.Println(city)     // Output: Mumbai
```


### 2. **Immutability of Strings**

- **Strings are immutable in Go.**
    - You can’t change a string after it’s created.
    - If you “modify” a string, you are creating a new string in memory.

```go
str := "hello"
// str[0] = 'H' // ❌ Not allowed! Compile-time error
// str = "somethingElse" // ❌ Also not valid
// str := "valid" // This is allowed

// To "change" just the first character, convert to a slice:
b := []byte(str)
b[0] = 'H'
str2 := string(b)
fmt.Println(str2) // Output: Hello
```

**Teaching tip:** Since strings are immutable, avoid repeated concatenation in loops (it’s inefficient)! More on that later.

### 3. **String Literals (Interpreted vs Raw)**

- **Interpreted String Literals:** Use double quotes `""`. Special characters (like `\n`) are processed.
- **Raw String Literals:** Use backticks `````. Everything is literal (no escape sequences).

```go
// Interpreted
strA := "First line\nSecond line"
fmt.Println(strA)
// First line
// Second line

// Raw (good for multi-line or patterns)
strB := `C:\Users\Go\file.txt
Second line\with\backslash`
fmt.Println(strB)
// Output exactly as written, including backslashes and newlines
```

Choose which makes your code clearer. Use raw strings for regex, file paths, or multi-line text.

### 4. **Accessing String Characters (Bytes vs Runes)**

#### - By Byte (Default)

```go
str := "GoLang"
fmt.Println(str[0]) // Output: 71 (ASCII code for 'G')
fmt.Printf("%c\n", str[1]) // Output: o

// Loop over bytes
for i := 0; i < len(str); i++ {
    fmt.Printf("%c ", str[i])
}
// Output: G o L a n g
```


#### - By Rune (for Unicode)

- A *rune* is an alias for `int32` and represents a Unicode code point.
- Useful for non-ASCII (e.g., emojis, foreign languages).

```go
unicodeStr := "gö 🦁" // g, o with umlaut, space, lion emoji

for i, runeValue := range unicodeStr {
    fmt.Printf("Index: %d, Rune: %c\n", i, runeValue)
}
// Output: Index: 0, Rune: g
//         Index: 1, Rune: ö
//         Index: 3, Rune:  
//         Index: 4, Rune: 🦁
```

**Notice:** Byte indexes differ from character (rune) indexes if the string contains Unicode!

### 5. **Length of a String (`len()` function)**

- `len(string)` returns **bytes**, not runes (Unicode characters).

```go
a := "hello"
b := "gö 🦁" // 3 runes, but more bytes

fmt.Println(len(a)) // Output: 5
fmt.Println(len(b)) // Output: Could be 7, because 'ö' and '🦁' are multi-byte

// Count runes (characters) correctly:
import "unicode/utf8"
fmt.Println(utf8.RuneCountInString(b)) // Output: 4
```


### 6. **String Concatenation and Formatting**

#### - Concatenation

**Use + for joining:**

```go
s1 := "Go"
s2 := "Lang"
result := s1 + " " + s2
fmt.Println(result) // Output: Go Lang
```

- For many strings or in a loop, use `strings.Builder` for efficiency.


#### - Formatting

**Use `fmt.Sprintf` to compose strings like templates:**

```go
name := "Divya"
age := 19
message := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
fmt.Println(message) // Output: My name is Divya and I am 19 years old.
```

- Common formatting verbs:
    - `%s` – string
    - `%d` – integer
    - `%f` – float
    - `%v` – any value


## Quick Practice Example

```go
// Try combining all concepts!
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "He💜Go"
    fmt.Printf("Original: %s\n", s)

    // Length in bytes vs runes
    fmt.Printf("Bytes: %d\n", len(s))
    fmt.Printf("Runes: %d\n", utf8.RuneCountInString(s))

    // Print each rune
    fmt.Print("Runes: ")
    for i, r := range s {
        fmt.Printf("%c ", r)
    }
    fmt.Println()

    // Concatenation and formatting
    part1 := "Go"
    part2 := "pher"
    full := part1 + part2
    fmt.Println("Concatenated: " + full)

    // Formatted string
    msg := fmt.Sprintf("%s is %d bytes, %d runes", s, len(s), utf8.RuneCountInString(s))
    fmt.Println(msg)
}
```


## Summary Table

| Concept | Code Example Sketched | Key Point |
| :-- | :-- | :-- |
| Declaration \& Initialization | `name := "Divya"` | Use `:=` for short-hand. |
| Immutability | `str = 'M'` (Not allowed) | Strings can’t be changed in-place. |
| Literals: Interpreted vs Raw | `"Hello\nWorld"` vs <code>\`Hello\nWorld\`</code> | Raw=backtick, literal, multi-line. |
| Access: Bytes vs Runes | `str[i]` vs `for _, rune := range str {}` | Use runes for Unicode correctness. |
| Length | `len(str)`, `utf8.RuneCountInString(str)` | `len` gives bytes, not characters! |
| Concatenation \& Formatting | `a + b`, `fmt.Sprintf("Hi, %s", name)` | Use `fmt.Sprintf` for templates. |
