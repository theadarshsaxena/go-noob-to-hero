## Key Patterns and Problems in Linked Lists (Go Implementation)

Mastering these core patterns and algorithms is crucial for competitive programming and deep DSA expertise. Here’s an in-depth explanation, with Go implementations and guidance for each problem:

### 1. **Finding the Middle of a Linked List (Fast & Slow Pointer)**

**Pattern**: Two pointers—`slow` moves one step at a time, `fast` moves two steps.

```go
func (l *LinkedList) FindMiddle() *Node {
    slow, fast := l.head, l.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    return slow
}
```
- **Tips**: Handles both even/odd lengths. Time: O(n), Space: O(1).

### 2. **Detecting Cycles (Floyd’s Cycle Detection Algorithm)**

**Pattern**: If `fast` and `slow` ever point to the same node, a cycle exists.

```go
func (l *LinkedList) HasCycle() bool {
    slow, fast := l.head, l.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            return true
        }
    }
    return false
}
```
- **Tips**: O(n) time, O(1) space. Classic interview question.

### 3. **Removing a Loop in Linked List**

**Pattern**: First, detect the loop (Floyd’s). Then, reset one pointer to head and advance both pointers one step at a time.

```go
func (l *LinkedList) RemoveCycle() {
    slow, fast := l.head, l.head
    loopExists := false
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
        if slow == fast {
            loopExists = true
            break
        }
    }
    if !loopExists {
        return
    }
    slow = l.head
    var prev *Node
    for slow != fast {
        prev = fast
        slow = slow.next
        fast = fast.next
    }
    prev.next = nil // Remove loop
}
```
- **Tips**: This severs the loop from the list.

### 4. **Finding the nth Node from the End (Two-pointer Technique)**

**Pattern**: Advance `first` by n steps, then move `first` and `second` together.

```go
func (l *LinkedList) NthFromEnd(n int) *Node {
    first, second := l.head, l.head
    for i := 0; i < n; i++ {
        if first == nil {
            return nil // n is larger than length
        }
        first = first.next
    }
    for first != nil {
        first = first.next
        second = second.next
    }
    return second
}
```
- **Tips**: O(n) time, O(1) space, single traversal.

### 5. **Merging Two Sorted Linked Lists**

**Pattern**: Use a dummy node; repeatedly link the smaller current node.

```go
func MergeSorted(l1, l2 *Node) *Node {
    dummy := &Node{}
    tail := dummy
    for l1 != nil && l2 != nil {
        if l1.data < l2.data {
            tail.next = l1
            l1 = l1.next
        } else {
            tail.next = l2
            l2 = l2.next
        }
        tail = tail.next
    }
    if l1 != nil {
        tail.next = l1
    } else {
        tail.next = l2
    }
    return dummy.next
}
```
- **Tips**: Handles all edge cases cleanly using a dummy head.

### 6. **Intersection of Two Linked Lists**

**Pattern**: Two pointers traverse both lists, switching to the other’s head upon reaching the end.

```go
func GetIntersectionNode(headA, headB *Node) *Node {
    if headA == nil || headB == nil {
        return nil
    }
    a, b := headA, headB
    for a != b {
        if a == nil {
            a = headB
        } else {
            a = a.next
        }
        if b == nil {
            b = headA
        } else {
            b = b.next
        }
    }
    return a // Intersection node or nil
}
```
- **Tips**: Equalizes distances covered, O(n+m) time, O(1) space.

### 7. **Checking if the List is a Palindrome**

**Pattern**: Find middle, reverse the second half, compare halves.

```go
func (l *LinkedList) IsPalindrome() bool {
    // Find middle (end of first half)
    slow, fast := l.head, l.head
    for fast != nil && fast.next != nil {
        slow = slow.next
        fast = fast.next.next
    }
    // Reverse second half
    var prev *Node
    current := slow
    for current != nil {
        next := current.next
        current.next = prev
        prev = current
        current = next
    }
    // Compare
    first, second := l.head, prev
    for second != nil {
        if first.data != second.data {
            return false
        }
        first = first.next
        second = second.next
    }
    return true
}
```
- **Tips**: Restore the list after check if needed.

### 8. **Partitioning a Linked List Around a Value**

**Pattern**: Use two dummy heads—one for nodes < x, one for nodes ≥ x.

```go
func Partition(head *Node, x int) *Node {
    before, after := &Node{}, &Node{}
    beforePtr, afterPtr := before, after
    current := head
    for current != nil {
        if current.data < x {
            beforePtr.next = current
            beforePtr = beforePtr.next
        } else {
            afterPtr.next = current
            afterPtr = afterPtr.next
        }
        current = current.next
    }
    afterPtr.next = nil
    beforePtr.next = after.next
    return before.next
}
```
- **Tips**: Preserves relative order; O(n) time, O(1) space.

### **Common Patterns to Master**

- **Fast/slow pointer**: For middle, cycles, or palindrome checks.
- **Dummy head node**: For cleaner edge/initial handling.
- **Two-pointer**: For distance/equalization tasks.
- **Edge Cases**: Always handle empty lists, single node, all nodes matching or not matching criteria.
- **In-place vs. new list**: Know when you must maintain original nodes vs. create new.
