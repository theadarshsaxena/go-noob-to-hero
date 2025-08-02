## String Conversion in Go: Full Explanation

Understanding how to convert strings to various types—and back—is essential in Go. Here’s a deep, clear look at **all conversion patterns** you’ll encounter, complete with code and best-practices.

### 1. **String to Integer / Float and Vice Versa**

These are common when dealing with input/output, file processing, or parsing data.

#### **String → Integer**

```go
import (
    "fmt"
    "strconv"
)

s := "42"
n, err := strconv.Atoi(s)
if err != nil {
    fmt.Println("Invalid integer!")
} else {
    fmt.Println(n) // 42 (int type)
}

// Or for different integer bases
s := "101"
n, err := strconv.ParseInt(s, 2, 64) // binary to int64
fmt.Println(n) // 5
```

#### **String → Float**

```go
f, err := strconv.ParseFloat("3.14", 64)
fmt.Println(f) // 3.14
```

#### **Integer / Float → String**

```go
si := strconv.Itoa(99) // int to string: "99"
sf := strconv.FormatFloat(7.892, 'f', 2, 64) // "7.89": 2 decimals
```

**Note:** All these conversions return an error! Always check `err`.

### 2. **String to Byte Slice (`[]byte`) and Rune Slice (`[]rune`)**

#### **String → []byte**
- Each element is a byte
- Good for low-level manipulation, writing to files, etc.

```go
s := "Go!"
b := []byte(s)
fmt.Println(b) // [71 111 33]
```

#### **String → []rune**
- Each element is a Unicode code point (rune, supports emoji/foreign characters)
- Lets you slice or change Unicode safely

```go
uni := "gö 🦁"
r := []rune(uni) // [103 246 32 129421]
fmt.Println(r)
fmt.Println(string(r[1:3])) // "ö "
```

### 3. **Byte and Rune Conversions**

#### **[]byte ↔ string**

```go
// []byte to string
b := []byte{104, 105}
s := string(b) // "hi"

// string to []byte
msg := "hello"
bs := []byte(msg)
```

#### **[]rune ↔ string**

```go
// []rune to string
r := []rune{'ग', 'ो', '-', '🚀'}
s := string(r) // "गो-🚀"

// string to []rune
sp := "go🚀lang"
rs := []rune(sp)
```

#### **byte ↔ rune**
- **byte** is 8 bits: good for ASCII and binary data.
- **rune** is 32 bits: needed for Unicode.

You can cast an ASCII byte to a rune:

```go
var b byte = 'A'
r := rune(b)
fmt.Printf("%c", r) // Output: A
```

But for Unicode, always use runes directly.

### **Example: Practice with Complete Conversions**

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"
    n, _ := strconv.Atoi(s)
    fmt.Println(n + 7)    // 130

    f, _ := strconv.ParseFloat("19.87", 64)
    fmt.Println(f * 2)    // 39.74

    // To string
    sFromInt := strconv.Itoa(41)
    sFromFloat := strconv.FormatFloat(1.23, 'f', 2, 64)
    fmt.Println(sFromInt, sFromFloat) // "41 1.23"

    // To bytes and runes
    st := "Go🚀"
    bytes := []byte(st)
    runes := []rune(st)
    fmt.Println(bytes) // [71 111 ...]
    fmt.Println(runes) // [71 111 128640]

    // Back to string
    fmt.Println(string(bytes))
    fmt.Println(string(runes))
}
```

### **Quick Reference Table**

| Operation                 | Example Code                                    | Output or Note        |
|---------------------------|-------------------------------------------------|-----------------------|
| String → int              | `strconv.Atoi("17")`                            | 17                    |
| int → String              | `strconv.Itoa(17)`                              | "17"                  |
| String → float            | `strconv.ParseFloat("3.25", 64)`                | 3.25                  |
| float → String            | `strconv.FormatFloat(3.25, 'f', 2, 64)`         | "3.25"                |
| String → []byte           | `[]byte("Go!")`                                 | [71 111 33]           |
| []byte → String           | `string([]byte{71, 111, 33})`                   | "Go!"                 |
| String → []rune           | `[]rune("गो")`                                  | [2327 2379]           |
| []rune → String           | `string([]rune{71, 111, 128640})`               | "Go🚀"                |

### **Tips and Gotchas**
- **Always check errors** from string-to-number conversions.
- **Unicode:** Use `[]rune` or range loops for any international/emoji data.
- ASCII? Simple—bytes suffice.

If you practice converting strings and understanding bytes vs runes, you’ll never be confused dealing with text, files, or network streams in Go DSA!