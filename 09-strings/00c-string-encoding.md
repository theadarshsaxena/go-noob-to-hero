## String Encoding in Go (ASCII vs. UTF-8, Runes, and Iteration)

Let’s deeply explore string encoding in Go. We’ll clarify **ASCII vs. UTF-8**, what *runes* are, and how to work with Unicode in real-world Go code. You’ll finish with full mastery of how Go handles and represents characters.

### 1. **ASCII vs. UTF-8**

#### **ASCII**
- **ASCII** stands for American Standard Code for Information Interchange.
- Represents English letters, digits, and some symbols using **7-bit** codes (values 0–127).
- Each character fits in **1 byte**.

Example:
```go
ascii := "hello"
fmt.Println(ascii[0])         // Output: 104 (ASCII code for 'h')
fmt.Printf("%c\n", ascii[0])  // Output: h
```

#### **UTF-8**
- **UTF-8** is a universal encoding covering all Unicode characters.
- Each character (or *rune*) can use **1 to 4 bytes**.
- Most modern systems (including Go) default to UTF-8.

**Why does this matter?**
- With UTF-8 you can store letters, symbols, emojis, and scripts from around the world in strings.

Example:
```go
s := "gö 🦁" // 'ö' and '🦁' need multiple bytes!
fmt.Println([]byte(s)) // Shows byte representation
```

### 2. **Rune (Unicode Code Points) and How Go Handles Unicode**

#### **What is a Rune?**
- In Go, a `rune` is an alias for `int32`, representing a Unicode code point—a unique number for every character, regardless of encoding.
- Use *runes* to count or manipulate individual characters, not bytes.

#### **Go’s Approach**
- Strings are by default a sequence of bytes (`[]byte`).
- For Unicode, Go provides the `rune` type and makes it easy to work with Unicode characters.
- All string literals in Go are UTF-8 encoded unless otherwise specified.

**Example: Decoding bytes to runes**
```go
import "unicode/utf8"

str := "π" // Greek small letter pi
fmt.Println(len(str))                   // Output: 2 (bytes)
fmt.Println(utf8.RuneCountInString(str)) // Output: 1 (rune)
fmt.Printf("%U\n", []rune(str)[0])   // Output: U+03C0
```

### 3. **Iterating Strings by Bytes vs. Runes**

#### **By Bytes (Default)**
```go
s := "gö"
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %v\n", i, s[i])
}
// Output: The "ö" character is represented by two bytes
```
- This is fast but unsafe if there are non-ASCII characters—could break individual Unicode characters apart.

#### **By Runes (Correct, especially for Unicode)**
- Use the `for ... range` loop, which gives rune and its byte position:
```go
s := "gö 🦁"
for idx, runeVal := range s {
    fmt.Printf("byte idx:%d, rune:%c, codepoint:%U\n", idx, runeVal, runeVal)
}
// Output shows correct Unicode characters and their locations
```

**Key Takeaways:**
- Indexing or slicing by bytes is **risky** for non-ASCII content.
- Always iterate by runes when dealing with text that may contain Unicode.

## Code Demo: ASCII, UTF-8, and Runes All Together

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "Go π 🦁" // English, Greek, and Emoji

    fmt.Println("Raw bytes:", []byte(s))
    fmt.Println("len(s):", len(s)) // bytes
    fmt.Println("utf8.RuneCountInString(s):", utf8.RuneCountInString(s)) // runes

    fmt.Println("\nIterate by bytes:")
    for i := 0; i < len(s); i++ {
        fmt.Printf("%d: %x\n", i, s[i])
    }

    fmt.Println("\nIterate by runes:")
    for idx, r := range s {
        fmt.Printf("byte idx:%d, rune:%c, codepoint:%U\n", idx, r, r)
    }
}
```

## Summary Table

| Aspect                         | ASCII                          | UTF-8 (Unicode)                         | Rune (Go)                        |
|---------------------------------|--------------------------------|-----------------------------------------|-----------------------------------|
| Encoding                       | 7 bits per char, 1 byte        | 1-4 bytes per char, variable length     | Alias for int32, Unicode codepoint|
| Example                        | `A` → 65                       | `π` → 0xCF80 (`\u03C0`), 🦁 is 4 bytes  | `r := 'π'`                        |
| String Iteration (`for ...`)    | By bytes safe                  | By runes required for correctness       | `for _, r := range s {}`          |

### Essential Practice Exercise

Try using these ideas:
- Write a loop that prints bytes and runes for `"नमस्ते"` (Hindi "Namaste").
- Check and print the length by bytes and by runes for `"Go🚀Lang"`.

These fundamentals give you confidence working with any kind of textual data in **Go**, from passwords and usernames to multi-language, emoji-rich applications!