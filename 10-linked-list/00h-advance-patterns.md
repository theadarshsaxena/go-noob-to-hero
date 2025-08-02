## Advanced Linked List Operations in Go: Concepts, Patterns, and Code

These advanced operations build on core linked list knowledge and are commonly featured in technical interviews and competitive programming challenges. Each section provides an overview, key patterns, and Go code examples.

### 1. **Copying a Linked List With Random Pointers**

A node has `next` and additional `random` pointers, pointing to any node (or nil).

#### **Pattern**:
- Use a map for O(n) time & space or interleave new nodes for O(1) space.

#### **Go Implementation (O(1) Space, Interleaving Pattern):**

```go
type RNode struct {
    data   int
    next   *RNode
    random *RNode
}

func CopyRandomList(head *RNode) *RNode {
    if head == nil {
        return nil
    }
    // Step 1: Interleave copied nodes
    for curr := head; curr != nil; {
        next := curr.next
        copy := &RNode{data: curr.data, next: next}
        curr.next = copy
        curr = next
    }
    // Step 2: Assign random pointers
    for curr := head; curr != nil; curr = curr.next.next {
        if curr.random != nil {
            curr.next.random = curr.random.next
        }
    }
    // Step 3: Unweave to separate lists
    orig, copyHead := head, head.next
    for orig != nil {
        copy := orig.next
        orig.next = copy.next
        if copy.next != nil {
            copy.next = copy.next.next
        }
        orig = orig.next
    }
    return copyHead
}
```
**Pattern**: Interleaving avoids extra space; always separate traversal and pointer assignment phases.

### 2. **Flattening a Multilevel Linked List**

Each node has `next` and `child` pointers representing sublists.

#### **Pattern**:
- Use DFS or an iterative stack to flatten depth-first.

#### **Go Implementation (Iterative, Stack-Based):**

```go
type MLNode struct {
    data  int
    next  *MLNode
    child *MLNode
}

func FlattenList(head *MLNode) *MLNode {
    if head == nil {
        return nil
    }
    stack := []*MLNode{}
    curr := head
    for curr != nil {
        if curr.child != nil {
            if curr.next != nil {
                stack = append(stack, curr.next)
            }
            curr.next = curr.child
            curr.child = nil
        } else if curr.next == nil && len(stack) > 0 {
            curr.next = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        curr = curr.next
    }
    return head
}
```
**Pattern**: Stack ensures you correctly maintain next pointers when jumping into child lists.

### 3. **Reversing Nodes in k-Groups**

Reverse every group of k nodes in the list.

#### **Pattern**:
- Use dummy nodes, in-place reversal, and pointer reconnection.

#### **Go Implementation:**

```go
func ReverseKGroup(head *Node, k int) *Node {
    dummy := &Node{next: head}
    groupPrev := dummy

    for {
        kth := groupPrev
        for i := 0; i < k && kth != nil; i++ {
            kth = kth.next
        }
        if kth == nil {
            break
        }

        groupNext := kth.next
        // Reverse group
        prev, curr := kth.next, groupPrev.next
        for curr != groupNext {
            temp := curr.next
            curr.next = prev
            prev = curr
            curr = temp
        }
        temp := groupPrev.next
        groupPrev.next = kth
        groupPrev = temp
    }
    return dummy.next
}
```
**Pattern**: Dummy head simplifies reconnections; always track the previous group's end and the next group's start.

### 4. **Group-wise Reversal (e.g., Reverse Alternate k Nodes)**

Reverse every other k-group (e.g., reverse k nodes, skip k nodes, repeat).

#### **Go Implementation (Reverse Alternate k):**

```go
func ReverseAlternateKGroup(head *Node, k int) *Node {
    current := head
    var prev *Node
    count := 0
    // Reverse k nodes
    for current != nil && count < k {
        next := current.next
        current.next = prev
        prev = current
        current = next
        count++
    }
    if head != nil {
        head.next = current
    }
    count = 0
    // Skip k nodes
    temp := current
    for temp != nil && count < k-1 {
        temp = temp.next
        count++
    }
    if temp != nil {
        temp.next = ReverseAlternateKGroup(temp.next, k)
    }
    return prev
}
```
**Pattern**: Clear division of phases—reverse, then skip, then recurse.

### 5. **Sorting a Linked List (Merge Sort and Insertion Sort)**

#### **Merge Sort (Preferred, O(n log n)):**

```go
func MergeSort(head *Node) *Node {
    if head == nil || head.next == nil {
        return head
    }
    // Split list
    mid := getMiddle(head)
    right := mid.next
    mid.next = nil
    left := MergeSort(head)
    right = MergeSort(right)
    // Merge
    return MergeSorted(left, right)
}
func getMiddle(head *Node) *Node {
    slow, fast := head, head
    for fast.next != nil && fast.next.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    return slow
}
```
- **Pattern:** Always use fast/slow pointers for splitting.

#### **Insertion Sort (O(n²), useful for nearly-sorted lists):**

```go
func InsertionSortList(head *Node) *Node {
    dummy := &Node{}
    current := head
    for current != nil {
        prev := dummy
        for prev.next != nil && prev.next.data < current.data {
            prev = prev.next
        }
        next := current.next
        current.next = prev.next
        prev.next = current
        current = next
    }
    return dummy.next
}
```
- **Pattern:** Leverage a dummy node to avoid frequent edge checks.

### 6. **LRU Cache (Using Doubly Linked List + Hashmap)**

Supports `Get`/`Put` in O(1) time with order updates.

#### **Pattern**:
- DLL maintains usage order; hashmap maps keys to nodes.

#### **Go Implementation (Skeleton):**

```go
type DLNode struct {
    key, value int
    prev, next *DLNode
}
type LRUCache struct {
    capacity int
    cache    map[int]*DLNode
    head     *DLNode // Most recently used
    tail     *DLNode // Least recently used
}

func NewLRUCache(cap int) *LRUCache {
    head, tail := &DLNode{}, &DLNode{}
    head.next = tail
    tail.prev = head
    return &LRUCache{capacity: cap, cache: make(map[int]*DLNode), head: head, tail: tail}
}

func (lru *LRUCache) Get(key int) int {
    if node, found := lru.cache[key]; found {
        lru.moveToFront(node)
        return node.value
    }
    return -1
}

func (lru *LRUCache) Put(key, value int) {
    if node, found := lru.cache[key]; found {
        node.value = value
        lru.moveToFront(node)
    } else {
        if len(lru.cache) == lru.capacity {
            lru.removeLRU()
        }
        newNode := &DLNode{key: key, value: value}
        lru.cache[key] = newNode
        lru.insertFront(newNode)
    }
}

// ... helper methods: moveToFront, insertFront, removeLRU, removeNode ...
```
- **Pattern:** Always move recently accessed/inserted nodes to the DLL head.

### 7. **Implementing Stacks and Queues Using Linked Lists**

#### **Stack (LIFO, Use Singly Linked List):**

```go
type Stack struct {
    top *Node
}
func (s *Stack) Push(val int) {
    s.top = &Node{data: val, next: s.top}
}
func (s *Stack) Pop() (int, bool) {
    if s.top == nil {
        return 0, false
    }
    val := s.top.data
    s.top = s.top.next
    return val, true
}
```

#### **Queue (FIFO, Use Both Head and Tail):**

```go
type Queue struct {
    head *Node
    tail *Node
}
func (q *Queue) Enqueue(val int) {
    newNode := &Node{data: val}
    if q.tail == nil {
        q.head = newNode
        q.tail = newNode
    } else {
        q.tail.next = newNode
        q.tail = newNode
    }
}
func (q *Queue) Dequeue() (int, bool) {
    if q.head == nil {
        return 0, false
    }
    val := q.head.data
    q.head = q.head.next
    if q.head == nil {
        q.tail = nil
    }
    return val, true
}
```
---

## **Common Patterns Across Problems**

- **Dummy nodes**: Simplify pointer manipulations and avoid multiple if/else branches.
- **Two-pointer traversal**: For group, alternate, or complex positional reversals and partitions.
- **Maps/Hashing**: For extra structure (as in LRU cache or random pointer copy).
- **Stacks**: Used in flattening and multi-level traversals.
- **Maintaining order with DLLs**: For caches or when removal from arbitrary positions is needed.

**Deliberate practice** on these operations—especially with clean pointer management and edge case handling—will grant you mastery in advanced linked list challenges often faced in programming interviews and contests.