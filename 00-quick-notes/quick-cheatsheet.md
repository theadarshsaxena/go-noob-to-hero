# Go Data Structures Cheat Sheet for Competitive Programming

## Arrays

### Initialization

```go
var arr [5]int                    // Zero-valued array
arr := [5]int{1, 2, 3, 4, 5}     // With values
arr := [...]int{1, 2, 3}         // Auto-size
arr := [5]int{0: 10, 4: 20}      // Specific indices
```

### Conditions

```go
if len(arr) > 0 { }               // Check length
if arr[0] == 5 { }                // Element condition
if arr == [5]int{} { }            // Compare arrays
```

### Loops

```go
for i := 0; i < len(arr); i++ { } // Index loop
for i, v := range arr { }         // Index + value
for _, v := range arr { }         // Value only
```

## Slices

### Initialization

```go
var s []int                       // Nil slice
s := []int{}                      // Empty slice
s := []int{1, 2, 3}              // With values
s := make([]int, 5)              // Length 5, zero values
s := make([]int, 5, 10)          // Length 5, capacity 10
```

### Conditions

```go
if len(s) == 0 { }               // Empty check
if s == nil { }                  // Nil check
if cap(s) > 10 { }               // Capacity check
if s[0] > 5 { }                  // Element condition
```

### Loops

```go
for i := 0; i < len(s); i++ { }  // Index loop
for i, v := range s { }          // Index + value
for _, v := range s { }          // Value only
// Reverse loop
for i := len(s) - 1; i >= 0; i-- { }
```

### Common Operations

```go
s = append(s, 1, 2, 3)           // Append elements
s = append(s, s2...)             // Append slice
s = s[:len(s)-1]                 // Remove last
s = s[1:]                        // Remove first
s = s[1:len(s)-1]               // Remove first & last
```

## Maps

### Initialization

```go
var m map[string]int             // Nil map
m = make(map[string]int)         // Empty map
m := map[string]int{}            // Empty map literal
m := map[string]int{             // With values
    "a": 1,
    "b": 2,
}
```

### Conditions

```go
if len(m) == 0 { }               // Empty check
if m == nil { }                  // Nil check
if v, ok := m["key"]; ok { }     // Key exists
if _, ok := m["key"]; !ok { }    // Key doesn't exist
if m["key"] > 5 { }              // Value condition
```

### Loops

```go
for k, v := range m { }          // Key + value
for k := range m { }             // Key only
for _, v := range m { }          // Value only
```

### Common Operations

```go
m["key"] = value                 // Set/update
delete(m, "key")                 // Delete key
v := m["key"]                    // Get (zero if missing)
v, ok := m["key"]                // Get with existence check
```

## Strings

### Initialization

```go
var s string                     // Empty string
s := "hello"                     // String literal
s := `multiline
string`                          // Raw string
s := string([]byte{65, 66})      // From bytes
```

### Conditions

```go
if len(s) == 0 { }               // Empty check
if s == "" { }                   // Empty check
if s[0] == 'a' { }               // First char (byte)
if strings.Contains(s, "sub") { } // Substring check
if strings.HasPrefix(s, "pre") { } // Prefix check
if strings.HasSuffix(s, "suf") { } // Suffix check
```

### Loops

```go
for i := 0; i < len(s); i++ { }  // Byte loop
for i, r := range s { }          // Rune loop (Unicode)
for _, r := range s { }          // Rune only
// Convert to slice for modification
for i, b := range []byte(s) { }  // Byte slice
```

### Common Operations

```go
s1 + s2                          // Concatenation
strings.Split(s, ",")            // Split to slice
strings.Join(slice, ",")         // Join slice
strings.ToLower(s)               // Lowercase
strings.ToUpper(s)               // Uppercase
```

## Runes (Characters)

### Initialization

```go
var r rune                       // Zero rune
r := 'a'                         // Rune literal
r := rune(65)                    // From int
runes := []rune(s)               // String to runes
```

### Conditions

```go
if r == 'a' { }                  // Character comparison
if r >= 'a' && r <= 'z' { }      // Range check
if unicode.IsDigit(r) { }        // Digit check
if unicode.IsLetter(r) { }       // Letter check
if unicode.IsSpace(r) { }        // Whitespace check
```

### Loops

```go
for _, r := range []rune(s) { }  // Rune slice
for i, r := range s { }          // String runes
```

## Structs

### Initialization

```go
type Point struct { X, Y int }
var p Point                      // Zero-valued
p := Point{1, 2}                 // Positional
p := Point{X: 1, Y: 2}           // Named fields
p := Point{X: 1}                 // Partial (Y = 0)
```

### Conditions

```go
if p.X > 0 { }                   // Field condition
if p == (Point{}) { }            // Zero value check
if p.X == p.Y { }                // Field comparison
```

### Loops

```go
points := []Point{{1,2}, {3,4}}
for _, p := range points { }     // Slice of structs
for i := range points { }        // Index only
```

## Pointers

### Initialization

```go
var p *int                       // Nil pointer
x := 42
p = &x                           // Address of x
p := new(int)                    // Allocated int pointer
```

### Conditions

```go
if p == nil { }                  // Nil check
if p != nil { }                  // Non-nil check
if *p > 5 { }                    // Dereference condition
```

### Loops

```go
ptrs := []*int{&x, &y, &z}
for _, p := range ptrs {
    if p != nil { *p = 0 }       // Safe dereference
}
```

## 2D Arrays

### Initialization

```go
var grid [3][4]int                // Fixed size 2D array
grid := [3][4]int{               // With values
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 10, 11, 12},
}
grid := [2][3]int{               // Partial initialization
    {1, 2},                      // Rest are zero
    {3, 4, 5},
}
```

### Conditions

```go
if len(grid) > 0 { }             // Check rows
if len(grid[0]) > 0 { }          // Check cols (if rows exist)
if grid[i][j] == 0 { }           // Element condition
if i >= 0 && i < len(grid) &&    // Bounds check
   j >= 0 && j < len(grid[0]) { }
```

### Loops

```go
// Standard nested loop
for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
        // grid[i][j]
    }
}

// Range loop
for i, row := range grid {
    for j, val := range row {
        // i, j, val
    }
}

// Row-wise processing
for i, row := range grid {
    for _, val := range row { }  // Process row
}
```

## 2D Slices

### Initialization

```go
var grid [][]int                 // Nil 2D slice

// Method 1: Make rows first
grid = make([][]int, rows)
for i := range grid {
    grid[i] = make([]int, cols)
}

// Method 2: All at once with values
grid = [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Method 3: Jagged array (different row sizes)
grid = [][]int{
    {1, 2},
    {3, 4, 5, 6},
    {7},
}

// Method 4: Pre-allocated with values
grid = make([][]int, rows)
for i := range grid {
    grid[i] = make([]int, cols)
    for j := range grid[i] {
        grid[i][j] = i*cols + j  // Initialize with values
    }
}
```

### Conditions

```go
if len(grid) == 0 { }            // Empty check
if grid == nil { }               // Nil check
if len(grid[i]) == 0 { }         // Empty row
if i < len(grid) && j < len(grid[i]) { } // Safe bounds
if grid[i][j] > 0 { }            // Element condition
```

### Loops

```go
// Standard nested loop
for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
        // grid[i][j]
    }
}

// Range loop with indices
for i, row := range grid {
    for j, val := range row {
        // i, j, val
    }
}

// Process each row
for _, row := range grid {
    for _, val := range row { }
}

// Reverse loops
for i := len(grid) - 1; i >= 0; i-- {
    for j := len(grid[i]) - 1; j >= 0; j-- {
        // grid[i][j]
    }
}
```

### Common Operations

```go
// Add row
grid = append(grid, []int{1, 2, 3})

// Add column (to each row)
for i := range grid {
    grid[i] = append(grid[i], newVal)
}

// Remove last row
if len(grid) > 0 {
    grid = grid[:len(grid)-1]
}

// Remove last column from each row
for i := range grid {
    if len(grid[i]) > 0 {
        grid[i] = grid[i][:len(grid[i])-1]
    }
}

// Copy 2D slice
newGrid := make([][]int, len(grid))
for i := range grid {
    newGrid[i] = make([]int, len(grid[i]))
    copy(newGrid[i], grid[i])
}
```

## Common Patterns for CP

### Frequency Map

```go
freq := make(map[rune]int)
for _, r := range s {
    freq[r]++
}
```

### Set Using Map

```go
set := make(map[int]bool)
set[1] = true
if set[1] { } // Check membership
delete(set, 1) // Remove
```

### Stack (using slice)

```go
stack := []int{}
stack = append(stack, x)         // Push
if len(stack) > 0 {
    top := stack[len(stack)-1]   // Peek
    stack = stack[:len(stack)-1] // Pop
}
```

### Queue (using slice)

```go
queue := []int{}
queue = append(queue, x)         // Enqueue
if len(queue) > 0 {
    front := queue[0]            // Front
    queue = queue[1:]            // Dequeue
}
```