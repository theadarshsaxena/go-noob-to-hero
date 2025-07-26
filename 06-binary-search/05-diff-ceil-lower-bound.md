## 🔍 Difference Between Ceil/Floor and Lower/Upper Bound

### 🔧 Example

```go
arr := []int{1, 3, 5, 7, 9}
target := 6
```

```
                    target=6
					   ↓
Index:       0   1   2   3   4
Array:     [ 1 , 3 , 5 , 7 , 9 ]
                     ↑   ↑
                  floor  ceil
                     ↑    ↑
            lower_bound   upper_bound
```

### 🧠 TL;DR Mental Map:

| Concept         | What it returns | Condition                  | Use-case Style           | Requires Sorted Array |
| --------------- | --------------- | -------------------------- | ------------------------ | --------------------- |
| **Lower Bound** | Index `i`       | First `arr[i] ≥ target`    | Binary Search patterns   | ✅ Yes                 |
| **Ceil**        | Value `arr[i]`  | First `arr[i] ≥ target`    | Mathematical/logical ops | ✅ Yes                 |
| **Upper Bound** | Index `i`       | First `arr[i] > target`    | Binary Search patterns   | ✅ Yes                 |
| **Floor**       | Value `arr[i]`  | Greatest `arr[i] ≤ target` | Mathematical/logical ops | ✅ Yes                 |
 
---

### ✍️ Key Difference:

* `lower_bound()` → gives **index**, `ceil()` → gives **value**
* `upper_bound()` → gives **index**, `floor()` → gives **value**


### 🔹 Meaning:

* `Ceil` (first number ≥ 6) = **7**
* `Floor` (greatest number ≤ 6) = **5**
* `Lower Bound` (index of first ≥ 6) = **3**
* `Upper Bound` (index of first > 6) = **3**

If `target = 5`, then:

* `Ceil = 5`, `Floor = 5`
* `Lower Bound = 2`
* `Upper Bound = 3`

---

### 💡 Tips to Remember:

* Think of `lower_bound` like "📍where you’d insert to stay sorted".
* Think of `ceil` like "🎯closest number from right that’s not smaller".
* Upper Bound always skips equal → goes to strictly greater 🔼
* Floor always skips greater → sticks to less than or equal 🔽
