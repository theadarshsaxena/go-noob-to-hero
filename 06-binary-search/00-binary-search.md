# Binary search patterns

## 🧠 Mental Rules (Always Work):

1. **Binary Search is just a while loop** over a sorted condition. It doesn't always need sorted data; it just needs a monotonic function ✅
2. **Two types of searches:**

   * **Index-based** (use `l <= r`)
   * **Range-based** (use `l < r`)
3. If stuck, **draw a dry run** for a few values. Mid value movement makes things clear 🔍
4. **Always protect mid calculation** using: `mid := l + (r-l)/2`
5. Always check for the following cases when dry running
    - Array having repeating numbers e.g. 1,2,4,4,4,8,9
    - Array containing (or not containing) the target element
    - Array with target out of range of elements present in array e.g. if array is 2,3,4,5 and target is 1 or 6

---

## ✅ 1. Standard Binary Search

> 🔍 Search for a target in a sorted array

```go
l, r := 0, n-1
for l <= r {
    mid := l + (r-l)/2
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        l = mid + 1
    } else {
        r = mid - 1
    }
}
```

📌 **“Ye to hone hi hai” points:**

* `l <= r` (not `<`)
* Use `mid := l + (r-l)/2` to avoid overflow
* Return `-1` if not found

---
