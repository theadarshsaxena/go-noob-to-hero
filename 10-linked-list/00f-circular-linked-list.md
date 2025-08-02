## Circular Linked List in Go: Detailed Guide

Circular linked lists form a loop—either singly (where the last node points back to the head) or doubly (where the list is linked in both directions, with the ends connected). These lists eliminate `nil` pointers, allowing infinite traversal and special use cases.

### 1. **Circular Singly Linked List (CSLL) Implementation**

#### Node Structure

```go
type CNode struct {
    data int
    next *CNode
}
```

#### Creating a Circular List

A CSLL is often represented with a `head` (or `tail`) pointer. Here’s a structure:

```go
type CircularLinkedList struct {
    head *CNode
    length int
}
```

#### Insertion at Tail (Common Pattern)

```go
func (cl *CircularLinkedList) InsertAtTail(data int) {
    newNode := &CNode{data: data}
    if cl.head == nil {
        newNode.next = newNode
        cl.head = newNode
    } else {
        current := cl.head
        for current.next != cl.head {
            current = current.next
        }
        current.next = newNode
        newNode.next = cl.head
    }
    cl.length++
}
```

#### Insertion at Head

```go
func (cl *CircularLinkedList) InsertAtHead(data int) {
    newNode := &CNode{data: data}
    if cl.head == nil {
        newNode.next = newNode
        cl.head = newNode
    } else {
        current := cl.head
        for current.next != cl.head {
            current = current.next
        }
        current.next = newNode
        newNode.next = cl.head
        cl.head = newNode
    }
    cl.length++
}
```

#### Deletion at Head

```go
func (cl *CircularLinkedList) DeleteAtHead() {
    if cl.head == nil {
        return
    }
    if cl.head.next == cl.head {
        cl.head = nil
    } else {
        current := cl.head
        for current.next != cl.head {
            current = current.next
        }
        current.next = cl.head.next
        cl.head = cl.head.next
    }
    cl.length--
}
```

### 2. **Circular Doubly Linked List (CDLL) Implementation**

#### Node Structure

```go
type CDNode struct {
    data int
    next *CDNode
    prev *CDNode
}
```

#### CDLL Structure

```go
type CircularDoublyLinkedList struct {
    head *CDNode
    length int
}
```

#### Insertion at Tail

```go
func (cdl *CircularDoublyLinkedList) InsertAtTail(data int) {
    newNode := &CDNode{data: data}
    if cdl.head == nil {
        newNode.next = newNode
        newNode.prev = newNode
        cdl.head = newNode
    } else {
        tail := cdl.head.prev
        tail.next = newNode
        newNode.prev = tail
        newNode.next = cdl.head
        cdl.head.prev = newNode
    }
    cdl.length++
}
```

#### Deletion at Head

```go
func (cdl *CircularDoublyLinkedList) DeleteAtHead() {
    if cdl.head == nil {
        return
    }
    if cdl.head.next == cdl.head {
        cdl.head = nil
    } else {
        tail := cdl.head.prev
        newHead := cdl.head.next
        tail.next = newHead
        newHead.prev = tail
        cdl.head = newHead
    }
    cdl.length--
}
```

### 3. **Traversal Logic**

#### CSLL Traversal

Avoid infinite loops by checking when you return to the head:

```go
func (cl *CircularLinkedList) PrintList() {
    if cl.head == nil {
        fmt.Println("List is empty")
        return
    }
    fmt.Print(cl.head.data, " ")
    for current := cl.head.next; current != cl.head; current = current.next {
        fmt.Print(current.data, " ")
    }
    fmt.Println()
}
```

#### CDLL Traversal (Forward and Backward)

```go
func (cdl *CircularDoublyLinkedList) PrintForward() {
    if cdl.head == nil {
        fmt.Println("List is empty")
        return
    }
    current := cdl.head
    for {
        fmt.Print(current.data, " ")
        current = current.next
        if current == cdl.head {
            break
        }
    }
    fmt.Println()
}

func (cdl *CircularDoublyLinkedList) PrintBackward() {
    if cdl.head == nil {
        fmt.Println("List is empty")
        return
    }
    tail := cdl.head.prev
    current := tail
    for {
        fmt.Print(current.data, " ")
        current = current.prev
        if current == tail {
            break
        }
    }
    fmt.Println()
}
```

### 4. **Common Patterns and Practices**

- **Always check** for a full loop to prevent infinite cycles during traversal.
- Handle the **single-node case** separately for insertion/deletion.
- Prefer keeping a **tail (or head.prev) reference** for O(1) tail inserts in circular lists.
- **Use cases:** CSLL/CDLL are useful in applications like round-robin scheduling, real-time systems (buffering), and some games.

### 5. **Sample Usage**

```go
func main() {
    cl := &CircularLinkedList{}
    cl.InsertAtTail(1)
    cl.InsertAtTail(2)
    cl.InsertAtHead(0)
    cl.PrintList() // Output: 0 1 2
    cl.DeleteAtHead()
    cl.PrintList() // Output: 1 2

    cdl := &CircularDoublyLinkedList{}
    cdl.InsertAtTail(5)
    cdl.InsertAtTail(6)
    cdl.PrintForward()  // Output: 5 6
    cdl.PrintBackward() // Output: 6 5
    cdl.DeleteAtHead()
    cdl.PrintForward()  // Output: 6
}
```

**Mastering circular lists—both singly and doubly—requires careful handling of looping pointers, edge cases, and effective traversal logic. Practice regular and edge case operations to develop robust intuition for these powerful data structures.**