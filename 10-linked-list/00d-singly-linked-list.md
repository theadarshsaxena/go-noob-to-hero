## Singly Linked List in Go: Complete Guide

Mastering singly linked lists is essential for data structures and competitive programming. This guide walks you through all critical operations, code patterns, and common pitfalls, focusing on Go implementations.

### 1. **Creating a Singly Linked List**

Each node stores data and a pointer to the next node.

```go
type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
    length int // optional, tracks size
}
```

#### Initializing an Empty List

```go
func NewLinkedList() *LinkedList {
    return &LinkedList{}
}
```

### 2. **Traversal**

#### Iterative Traversal

Walk through the list from head to tail.

```go
func (l *LinkedList) TraverseIterative() {
    current := l.head
    for current != nil {
        fmt.Println(current.data)
        current = current.next
    }
}
```

#### Recursive Traversal

A recursive approach to print each node’s data.

```go
func (l *LinkedList) TraverseRecursive(n *Node) {
    if n == nil {
        return
    }
    fmt.Println(n.data)
    l.TraverseRecursive(n.next)
}
```
Call with: `ll.TraverseRecursive(ll.head)`

### 3. **Insertion**

#### Insert at Head

```go
func (l *LinkedList) InsertAtHead(data int) {
    newNode := &Node{data: data, next: l.head}
    l.head = newNode
    l.length++
}
```

#### Insert at Tail

```go
func (l *LinkedList) InsertAtTail(data int) {
    newNode := &Node{data: data}
    if l.head == nil {
        l.head = newNode
    } else {
        current := l.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    l.length++
}
```

#### Insert at Given Position (1-based)

```go
func (l *LinkedList) InsertAtPosition(data, pos int) {
    if pos <= 1 {
        l.InsertAtHead(data)
        return
    }
    current := l.head
    for i := 1; i < pos-1 && current != nil; i++ {
        current = current.next
    }
    if current == nil {
        l.InsertAtTail(data)
        return
    }
    newNode := &Node{data: data, next: current.next}
    current.next = newNode
    l.length++
}
```

### 4. **Deletion**

#### Delete at Head

```go
func (l *LinkedList) DeleteAtHead() {
    if l.head != nil {
        l.head = l.head.next
        l.length--
    }
}
```

#### Delete at Tail

```go
func (l *LinkedList) DeleteAtTail() {
    if l.head == nil {
        return
    }
    if l.head.next == nil {
        l.head = nil
    } else {
        current := l.head
        for current.next.next != nil {
            current = current.next
        }
        current.next = nil
    }
    l.length--
}
```

#### Delete at Given Position (1-based, assume valid position)

```go
func (l *LinkedList) DeleteAtPosition(pos int) {
    if pos <= 1 || l.head == nil {
        l.DeleteAtHead()
        return
    }
    current := l.head
    for i := 1; i < pos-1 && current.next != nil; i++ {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
        l.length--
    }
}
```

### 5. **Searching**

#### Iterative Search

```go
func (l *LinkedList) SearchIterative(target int) bool {
    current := l.head
    for current != nil {
        if current.data == target {
            return true
        }
        current = current.next
    }
    return false
}
```

#### Recursive Search

```go
func (l *LinkedList) SearchRecursive(n *Node, target int) bool {
    if n == nil {
        return false
    }
    if n.data == target {
        return true
    }
    return l.SearchRecursive(n.next, target)
}
```
Call with: `ll.SearchRecursive(ll.head, value)`

### 6. **Reversing a Singly Linked List**

#### Iterative Method

```go
func (l *LinkedList) ReverseIterative() {
    var prev *Node = nil
    current := l.head
    var next *Node
    for current != nil {
        next = current.next
        current.next = prev
        prev = current
        current = next
    }
    l.head = prev
}
```

#### Recursive Method

```go
func (l *LinkedList) ReverseRecursive(n *Node) *Node {
    if n == nil || n.next == nil {
        return n
    }
    rest := l.ReverseRecursive(n.next)
    n.next.next = n
    n.next = nil
    return rest
}

// To reverse, call: ll.head = ll.ReverseRecursive(ll.head)
```

### 7. **Common Patterns To Follow**

- **Dummy Node Pattern:** For complex insertions and deletions, use a dummy (sentinel) node to avoid handling head separately.
- **Length Tracking:** Keep a `length` field for quick list size queries.
- **Boundary Checks:** Always check for `nil` before dereferencing.
- **Recursion:** Practice both recursive and iterative for interview/problem-solving flexibility.
- **Edge Cases:** Always consider empty list, single element, insert/delete at head/tail.

### 8. **Sample Usage**

```go
func main() {
    ll := NewLinkedList()
    ll.InsertAtHead(3)
    ll.InsertAtTail(5)
    ll.InsertAtPosition(4, 2)
    ll.TraverseIterative() // 3 4 5

    ll.DeleteAtPosition(2)
    ll.TraverseIterative() // 3 5

    fmt.Println(ll.SearchIterative(5)) // true

    ll.ReverseIterative()
    ll.TraverseIterative() // 5 3
}
```

By consistently practicing these operations—head/tail/position insertions and deletions, iterative and recursive patterns, and handling of edge cases—you’ll become proficient in both interview and competitive programming settings for singly linked lists in Go.