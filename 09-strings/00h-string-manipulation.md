## Advanced String Manipulation Techniques: Custom Sorts, In-place Modification, Buffering \& Caching

Let's master the next set of powerful string techniques in Go: **custom string sorting**, efficient *in-place* modification using byte slices, and correct approaches to string *buffering* and *caching*. Each is vital for high-performance data structures, algorithms, and real-world Go development.

### 1. **Custom String Sorts**

Sorting strings with your own logic (e.g., by length, custom alphabet, case-insensitive) is often needed in algorithms and coding interviews.

#### **How Sorting Strings Works in Go**

Go’s standard library has a powerful `sort` package. For simple lexicographic sorting:

```go
import (
    "fmt"
    "sort"
)

func main() {
    arr := []string{"go", "Java", "python", "C"}
    sort.Strings(arr)           // Lexicographic (default)
    fmt.Println(arr)            // [C Java go python]
}
```


#### **Custom Sorts with `sort.Slice` or `sort.Sort`:**

- **Sort by length, custom rules, etc.**

```go
import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    arr := []string{"Go", "Gopher", "AI", "AlphaGo"}
    // Sort by string length (shortest first)
    sort.Slice(arr, func(i, j int) bool {
        return len(arr[i]) < len(arr[j])
    })
    fmt.Println(arr) // [Go AI Gopher AlphaGo]

    // Case-insensitive sort
    sort.Slice(arr, func(i, j int) bool {
        return strings.ToLower(arr[i]) < strings.ToLower(arr[j])
    })
    fmt.Println(arr)
}
```

> **Tip:** For very advanced needs, implement the `sort.Interface`.

### 2. **In-place String Modification (with Byte Slices)**

As you've learned, **strings are immutable** in Go. But you can *simulate in-place modification* by converting a string to a slice of bytes (when dealing with ASCII), or runes (for Unicode).

#### **Modify with `[]byte` (ASCII only):**

```go
s := "food"
b := []byte(s)
b[0] = 'g'           // Change 'f' to 'g'
s2 := string(b)
fmt.Println(s2)      // Output: "good"
```


#### **Unicode-safe Modification with `[]rune`:**

For any non-ASCII data, always use `[]rune`:

```go
s := "göph"
r := []rune(s)
r[0] = 'G'           // Change 'g' to 'G'
s2 := string(r)
fmt.Println(s2)      // Output: "Göph"
```


#### **Reverse a String In-place (Common DSA Interview):**

```go
func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    fmt.Println(reverseString("Go💡lang")) // Output: gnal💡oG
}
```

- When *in-place* appears in Go, think "use a mutable slice (byte/rune), then convert back if you need a string."


### 3. **String Buffering and Caching**

Efficient string handling sometimes involves **buffering** (storing temporary or incoming data) and **caching** (keeping ready-to-use results in memory).

#### **Buffering: Growing Data Efficiently**

For buffering, use `strings.Builder` for stringbuilding or `bytes.Buffer` for more generic byte-wise operations.

**Example with `bytes.Buffer`:**

```go
import (
    "bytes"
    "fmt"
)

func main() {
    var buf bytes.Buffer
    buf.WriteString("Hello, ")
    buf.Write([]byte("buffered world!"))
    out := buf.String()
    fmt.Println(out) // Output: Hello, buffered world!
}
```

**Why use it?**

- Useful when you’re dealing with raw bytes, or when you need to build data before writing to a file/network.


#### **String Caching with Maps**

To *cache* string processing results (e.g., in algorithms that reprocess the same substring), use Go’s built-in map type:

```go
import "fmt"

func expensiveComputation(s string) int {
    // Imagine this is slow
    return len(s) * 2
}

func main() {
    cache := make(map[string]int)
    s := "golang"
    if val, ok := cache[s]; ok {
        fmt.Println("Cached:", val)
    } else {
        val := expensiveComputation(s)
        cache[s] = val
        fmt.Println("Computed:", val)
    }
}
```

- **Pronounced in interviews and practice** for dynamic programming (memoization), repetitive parsing, and similar optimization cases.


#### **When is Buffering/Caching Used?**

- Building or reading large files/streams
- Parsing input for competitive programming
- Optimizing algorithms with subproblem repetition


### **Quick Recap Table**

| Technique | Tool/Pattern | Example Use Case | Key Practice |
| :-- | :-- | :-- | :-- |
| Custom String Sort | `sort.Slice/Sort` | Sort by length, custom | Provide custom less-func |
| In-place Modification | `[]byte`, `[]rune` | Char replace, reverse | Convert to slice, modify, back to string |
| Buffering | `strings.Builder`/`bytes.Buffer` | Big string building | Use for efficient appending |
| Caching | `map[string]valType` | Memoization, optimization | Store repeat results for quick use |

### **Summary**

- **Custom sorting** gives you flexibility for any string order or comparison logic.
- **In-place string modification** is possible via converting to slices—crucial in many DSA problems.
- **Efficient buffering and caching** help you handle big data and fast repeated access.

Practice each with your own code challenges—these are professional-grade tools! If you want DSA-problem-specific code samples or to see more edge cases, just ask.

