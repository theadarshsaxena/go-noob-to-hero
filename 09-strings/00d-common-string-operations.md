## Common String Operations in Go: Complete Guide

Let’s work through all the essential string operations in Go, with easy explanations and code for each. These are the building blocks for solving many real-world and DSA problems.

### 1. **Substrings (Slicing)**

- **Strings in Go can be sliced**, like arrays or slices, but *only by byte indexes* (watch out for Unicode!).
- Slicing returns a new string referring to the original bytes.

```go
s := "Hello, World!"
sub := s[7:12] // from byte 7 to 11
fmt.Println(sub) // Output: World
```

> **Caution:** Slicing works **correctly for ASCII**, but may break for multi-byte Unicode characters. For Unicode “safe” substrings, use `[]rune` conversion.

```go
uni := "göpher"      // 'ö' is two bytes
runed := []rune(uni)
fmt.Println(string(runed[1:3])) // Output: "öp"
```

### 2. **String Comparison**

- **Using `==` or `!=` for comparison:**
  - Compares byte-by-byte if strings are of the same length.

```go
a := "apple"
b := "Apple"
c := "apple"
fmt.Println(a == b) // Output: false
fmt.Println(a == c) // Output: true
```

- **Case-insensitive comparison:** Use `strings.EqualFold`

```go
import "strings"
fmt.Println(strings.EqualFold("GoLang", "golang")) // Output: true
```

### 3. **Searching within Strings**

- Use the `strings` package for most search tasks.

```go
import "strings"

str := "hello, world"

// Contains substring?
fmt.Println(strings.Contains(str, "world")) // true

// Find index, -1 if not found
fmt.Println(strings.Index(str, "o")) // 4 (first 'o')
fmt.Println(strings.LastIndex(str, "o")) // 8 (last 'o')

// Has prefix or suffix?
fmt.Println(strings.HasPrefix(str, "hello")) // true
fmt.Println(strings.HasSuffix(str, "world")) // true
```

### 4. **Modifying Strings (Building New Strings, Not In-Place)**

- **Strings are immutable**: create a new string for "modifications".
- **Replace substrings:** Use `strings.Replace` (or related).

```go
import "strings"

s := "banana apple banana"
replaced := strings.Replace(s, "banana", "orange", 1) // Only first
fmt.Println(replaced) // orange apple banana

// Replace all
replacedAll := strings.ReplaceAll(s, "banana", "orange")
fmt.Println(replacedAll) // orange apple orange

// To insert/delete/change, use concatenation or slices.
ss := "kite"
newstr := ss[:1] + "i" + ss[2:] // Change 'k' to 'i'
fmt.Println(newstr) // "iite"
```

### 5. **Splitting and Joining Strings**

- **Splitting:** Break a string into a slice of substrings.

```go
import "strings"

csv := "a,b,c"
fields := strings.Split(csv, ",") // ["a" "b" "c"]
fmt.Println(fields)

// Split at white space
input := "go for go"
words := strings.Fields(input) // ["go" "for" "go"]

// Limit splits: SplitN
a := "x-y-z"
parts := strings.SplitN(a, "-", 2) // ["x" "y-z"]
```

- **Joining:** Combine slices into a single string.

```go
joined := strings.Join(fields, ";") // "a;b;c"
fmt.Println(joined)
```

### 6. **Trimming and Padding**

- Use `strings.Trim*` for cleaning up unwanted characters:

```go
import "strings"

s := "  hello \t\n"
fmt.Printf("[%s]\n", strings.TrimSpace(s)) // [hello]

// Trim specific bytes
addr := "!!!house!!!"
fmt.Println(strings.Trim(addr, "!")) // house

// Pad (add to length): since Go 1.10+
padded := s + strings.Repeat("!", 5-len(s)) // crude right pad
fmt.Println(padded)

import "fmt"

// Left/Right Pad helper function
func padRight(str string, length int, padChar string) string {
    if len(str) >= length {
        return str
    }
    return str + strings.Repeat(padChar, length-len(str))
}

fmt.Println(padRight("go", 5, "_")) // go___
```

## Quick Reference Table

| Operation         | Function/Example                                    | Key Usage                              |
|-------------------|-----------------------------------------------------|----------------------------------------|
| Substring         | `s[2:5]`, `string([]rune(s)[2:5])`                  | Slice by bytes/runes                   |
| Comparison        | `==`, `!=`, `strings.EqualFold(a, b)`               | Direct and case-insensitive            |
| Search            | `strings.Contains(s, "t")`                          | Substring check                        |
| Search Index      | `strings.Index/LastIndex(s, "o")`                   | Position of substring                  |
| Prefix/Suffix     | `strings.HasPrefix/HasSuffix(s, "pre")`             | Test beginning/end                     |
| Replace           | `strings.Replace/ReplaceAll(s, old, new, n)`        | Replace substrings                     |
| Split             | `strings.Split(s, ",")`, `strings.Fields(s)`        | Into slice of substrings               |
| Join              | `strings.Join(arr, sep)`                            | From slice to string                   |
| Trim              | `strings.TrimSpace/Trim(s, chars)`                  | Remove unwanted chars                  |
| Pad               | Custom with `strings.Repeat`                        | Add fixed chars to length              |

**Practice Tip:**  
Work through these operations by writing tiny Go programs. Try working with Unicode (e.g., “Go🐍Lang”), and edge cases (empty strings, absent substrings, etc.) to master all behaviors!