## Regular Expressions in Go (Package `regexp`): Deep Dive & Tutorials

Let's master **regular expressions** (regex) in Go, focusing on practical usage of the powerful `regexp` package. You'll learn what regular expressions are, how to use all essential functions for searching, matching, extracting, and replacing—in simple, hands-on ways.

### **What is a Regular Expression?**

A **regular expression** is a sequence of characters that defines a search pattern. It's used for tasks like:
- Checking if a string matches a rule (validation)
- Extracting substrings (parsing)
- Replacing parts of text (transformation)

Go uses the **RE2 regex syntax**, similar to Python, Perl, and JavaScript.

## 1. **Getting Started: Import and Compile Patterns**

Import the package:
```go
import "regexp"
```

**Compile the regex:**  
- `Compile()` returns an error if the pattern's wrong—handle it.
- `MustCompile()` panics if the pattern is wrong—handy for static patterns.

```go
re := regexp.MustCompile(`\d{3}-\d{2}-\d{4}`) // Pattern for US SSN
```

> **Tip:** Use raw strings (backticks), so you don't have to escape every backslash.

## 2. **Matching: Does This String Fit the Pattern?**

- **Match entire string or substring?**
- Use `.MatchString()` or package-level `MatchString`.

```go
pattern := `^Go`
re := regexp.MustCompile(pattern)
fmt.Println(re.MatchString("Golang is cool")) // true
fmt.Println(re.MatchString("I love Go"))      // false
```

**Find a substring match:**
```go
hasDigit := regexp.MustCompile(`\d`).MatchString("abc123xyz")
fmt.Println(hasDigit) // true
```

## 3. **Finding and Extracting**

**- Find first match:**
```go
re := regexp.MustCompile(`[A-Z]\w+`)
fmt.Println(re.FindString("my ID is Xyz789")) // Xyz789
```

**- Find all matches:**
```go
re := regexp.MustCompile(`[a-zA-Z]+`)
words := re.FindAllString("Go is great!", -1)
fmt.Println(words) // ["Go" "is" "great"]
```

**- Submatches (capture groups):**
```go
re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
parts := re.FindStringSubmatch("id: 123-45-6789")
fmt.Println(parts) // ["123-45-6789" "123" "45" "6789"]
```

## 4. **Replacing Content**

**- Replace all matches:**
```go
re := regexp.MustCompile(`foo.?`)
result := re.ReplaceAllString("seafood fool", "bar")
fmt.Println(result) // "seabar barl"
```

**- Use ReplaceAllStringFunc for custom logic:**
```go
re := regexp.MustCompile(`[a-z]+`)
out := re.ReplaceAllStringFunc("hello 123 world", strings.ToUpper)
fmt.Println(out) // "HELLO 123 WORLD"
```

## 5. **Splitting Text**

Like `strings.Split`, but with a regex pattern:
```go
re := regexp.MustCompile(`\s+`)
words := re.Split("How are   you?\nFine!", -1)
fmt.Println(words) // ["How" "are" "you?" "Fine!"]
```

## 6. **Cheat Sheet: Core Methods**

| Purpose                      | Function                         | Signature                        | Uses Example                                      |
|------------------------------|----------------------------------|----------------------------------|---------------------------------------------------|
| Compile pattern              | `regexp.Compile`, `MustCompile`  | `Regexp` object                  | Compile, panic on error with MustCompile           |
| Match                        | `re.MatchString`                 | `bool`                           | Check presence of a pattern inside a string        |
| Find first match             | `re.FindString(s)`               | `string`                         | Get first matching substring                      |
| Find all matches             | `re.FindAllString(s, n)`         | `[]string`                       | Get all matches in string                         |
| Extract groups               | `re.FindStringSubmatch(s)`       | `[]string`                       | Extract captured (grouped) parts                  |
| Replace                      | `re.ReplaceAllString(s, r)`      | `string`                         | Replace matches with replacement                  |
| Replace with logic           | `re.ReplaceAllStringFunc`        | `string`                         | Apply function to each match                      |
| Split                       | `re.Split(s, n)`                 | `[]string`                       | Split string on pattern                           |

## 7. **Best Practices and Gotchas**

- Use **raw strings** for patterns: `` `\d+` ``.
- Always check for errors if you use `Compile`.
- Regex in Go works for **Unicode** text as well.
- For complex regex processing, combine with methods like `FindAllStringSubmatch`, `ReplaceAllStringFunc`, etc.
- Regex performance in Go is generally very good (uses RE2 engine).

## 8. **A Real-World Example**

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Validate and extract email
    re := regexp.MustCompile(`([\w\.\-]+)@([\w\-]+)\.([a-zA-Z]{2,})`)
    input := "Contact us at team@example.com or hello@world.io"
    emails := re.FindAllStringSubmatch(input, -1)
    for _, match := range emails {
        fmt.Println("Full:", match[0], "User:", match[1], "Domain:", match[2], "TLD:", match[3])
    }
}
```
**Output:**
```
Full: team@example.com User: team Domain: example TLD: com
Full: hello@world.io User: hello Domain: world TLD: io
```

Mastering Go's `regexp` package unlocks a vast range of text and data processing abilities: validation, parsing, transformation, and beyond. Play with sample code, learn the patterns, and use the power of regular expressions in your Go DSA journey