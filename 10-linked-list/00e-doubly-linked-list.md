## Doubly Linked List in Go: Complete Guide

A **doubly linked list (DLL)** allows traversal in both directions using two pointers in each node: one for the next node, one for the previous.

### 1. **Structure and Node Implementation in Go**

#### Node Definition

Each node has:
- `data`: The value to store.
- `prev`: Pointer to the previous node.
- `next`: Pointer to the next node.

```go
type DNode struct {
    data int
    prev *DNode
    next *DNode
}
```

#### Doubly Linked List Structure

- `head`: Points to the first node.
- Optionally, `tail` can be added for efficient tail insertions.

```go
type DoublyLinkedList struct {
    head *DNode
    tail *DNode // For O(1) tail operations
    length int
}
```

### 2. **Insertion**

#### Insert at Head

```go
func (dl *DoublyLinkedList) InsertAtHead(data int) {
    newNode := &DNode{data: data}
    if dl.head == nil {
        dl.head = newNode
        dl.tail = newNode
    } else {
        newNode.next = dl.head
        dl.head.prev = newNode
        dl.head = newNode
    }
    dl.length++
}
```

#### Insert at Tail

```go
func (dl *DoublyLinkedList) InsertAtTail(data int) {
    newNode := &DNode{data: data}
    if dl.tail == nil {
        dl.head = newNode
        dl.tail = newNode
    } else {
        dl.tail.next = newNode
        newNode.prev = dl.tail
        dl.tail = newNode
    }
    dl.length++
}
```

#### Insert at a Given Position (1-based index)

```go
func (dl *DoublyLinkedList) InsertAtPosition(data, pos int) {
    if pos <= 1 {
        dl.InsertAtHead(data)
        return
    }
    current := dl.head
    for i := 1; i < pos-1 && current != nil; i++ {
        current = current.next
    }
    if current == nil || current.next == nil {
        dl.InsertAtTail(data)
        return
    }
    newNode := &DNode{data: data}
    newNode.next = current.next
    newNode.prev = current
    current.next.prev = newNode
    current.next = newNode
    dl.length++
}
```

### 3. **Deletion**

#### Delete at Head

```go
func (dl *DoublyLinkedList) DeleteAtHead() {
    if dl.head == nil {
        return
    }
    if dl.head == dl.tail {
        dl.head = nil
        dl.tail = nil
    } else {
        dl.head = dl.head.next
        dl.head.prev = nil
    }
    dl.length--
}
```

#### Delete at Tail

```go
func (dl *DoublyLinkedList) DeleteAtTail() {
    if dl.tail == nil {
        return
    }
    if dl.head == dl.tail {
        dl.head = nil
        dl.tail = nil
    } else {
        dl.tail = dl.tail.prev
        dl.tail.next = nil
    }
    dl.length--
}
```

#### Delete at a Given Position (1-based)

```go
func (dl *DoublyLinkedList) DeleteAtPosition(pos int) {
    if pos <= 1 {
        dl.DeleteAtHead()
        return
    }
    current := dl.head
    for i := 1; i < pos && current != nil; i++ {
        current = current.next
    }
    if current == nil {
        return
    }
    if current == dl.tail {
        dl.DeleteAtTail()
        return
    }
    current.prev.next = current.next
    if current.next != nil {
        current.next.prev = current.prev
    }
    dl.length--
}
```

### 4. **Bidirectional Traversal**

#### Forward Traversal

```go
func (dl *DoublyLinkedList) PrintForward() {
    current := dl.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}
```

#### Backward Traversal

```go
func (dl *DoublyLinkedList) PrintBackward() {
    current := dl.tail
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.prev
    }
    fmt.Println()
}
```

### 5. **Reversing a Doubly Linked List**

#### Iterative Reversal

```go
func (dl *DoublyLinkedList) Reverse() {
    var prev, temp *DNode
    current := dl.head
    dl.tail = dl.head
    for current != nil {
        temp = current.prev
        current.prev = current.next
        current.next = temp
        prev = current
        current = current.prev
    }
    if prev != nil {
        dl.head = prev
    }
}
```

### 6. **Common Patterns and Practices**

- **Head/Tail Tracking**: Use both for O(1) insert/delete at both ends.
- **Boundary & Nil Checks**: Always check for empty list and single-node cases.
- **Bidirectional Traversal**: Useful for palindromic checks, reverse traversals, etc.
- **Updating Head/Tail**: After reversal or certain deletions, update both accordingly.

### 7. **Sample Usage**

```go
func main() {
    dl := &DoublyLinkedList{}
    dl.InsertAtHead(2)
    dl.InsertAtTail(3)
    dl.InsertAtHead(1)
    dl.InsertAtPosition(4, 3)
    dl.PrintForward()    // Output: 1 2 4 3
    dl.PrintBackward()   // Output: 3 4 2 1
    dl.Reverse()
    dl.PrintForward()    // Output: 3 4 2 1
    dl.DeleteAtHead()
    dl.PrintForward()    // Output: 4 2 1
}
```

Mastering these patterns, common operations, and edge cases will prepare you for advanced doubly linked list problems in competitive programming and real-world applications. Practice insertion, deletion, and traversal in both directions, and pay close attention to pointer manipulations for bug-free code!