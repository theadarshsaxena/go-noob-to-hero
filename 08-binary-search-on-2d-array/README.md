## Binary Search on 2D Arrays: Core Concepts

Binary search is usually applied to 1D sorted sequences, but several **2D array patterns** commonly appear in DSA problems. Below you'll find the concepts, techniques, and design principles you'll need.

### 1. **Understand the Matrix's Properties**

- **Fully Sorted (row-wise & col-wise):** Every row and column is sorted from left to right and top to bottom.
- **Row-wise Sorted:** Each row is sorted, but rows are not necessarily in order compared to each other.
- **Staircase Sorted:** Each row and column is sorted, allowing for "walk along matrix" approaches.

**Knowing these properties is essential!** They determine which pattern to apply.

### 2. **Common Patterns**

#### a. **Flatten Matrix & Classic Binary Search**
- **When to Use:** Matrix is sorted as if laid out row-wise into a 1D array.
- **How to Apply:**
  - Treat indices [i, j] as a single index: `index = i * n + j`, where `n = number of columns`.
  - Convert mid index back: `row = mid // n`, `col = mid % n`.
  - Perform binary search as you would on a normal array.

#### b. **Row-Pointer Binary Search**
- **When to Use:** Each row is sorted independently.
- **How to Apply:**
  - For search: Run binary search within each row.
  - For range search: Sometimes start from top-right or bottom-left and adjust your row/column pointer.

#### c. **Search Space Partition (Upper/Lower Bound)**
- **When to Use:** Problems asking for "how many elements ≤ x" in a matrix with each row/column sorted.
- **How to Apply:**
  - Use a stair-step approach for counting, or
  - Use a modified binary search along one dimension.

#### d. **Binary Search on Answer Space**
- Sometimes, the 2D matrix generates a virtual search space for an answer-value, and you use the matrix as a predicate/checker (common in "kth smallest/largest" type problems).

### 3. **Good Practices**

- **Always clarify matrix guarantees** (sorted? fully? per row/col?).
- **Carefully map 1D/2D indices** to avoid off-by-one and out-of-bounds errors.
- **Optimize search boundaries:** For large matrices, avoid O(m log n) scans if O(log(mn)) is possible via flattening.

### 4. **Design Patterns & Thought Process**

| Step          | What to Ask Yourself                       | Example Approach                       |
|---------------|--------------------------------------------|----------------------------------------|
| 1. Explore    | Is the 2D array globally sorted or by row? | Choose flattening vs. per-row search   |
| 2. Model      | Can I treat this as 1D?                    | Use index arithmetic                   |
| 3. Predicate  | What is my search predicate?               | Is mat[mid] == target?                 |
| 4. Traverse   | Can I move smartly (top-right stair)?      | Step left/down to narrow search space  |
| 5. Validate   | Double-check conversion between indices    | Use test cases with corners/edges      |

### 5. **Example Scenarios**

- **Search a value in a fully sorted matrix:** Flatten and perform 1D binary search.
- **Search in row & col sorted matrix:** Use "staircase search" starting at top-right or bottom-left and moving accordingly.
- **Find kth smallest in matrix:** Apply a binary search on the value range and count elements ≤ mid using pointers.

### 6. **Interview & Coding Tips**

- **Edge Test:** Try with 1x1, 1xN, Nx1 matrices and largest/smallest values.
- **Boundary Handling:** Watch for infinite loops in case of bad index conversion.
- **Time Complexity:** Flattened binary search: O(log(mn)). Staircase: O(m+n).

### **Summary Table**

| Matrix Type                 | Best Search Method         | Complexity          |
|-----------------------------|---------------------------|---------------------|
| Fully Sorted (1D/2D)        | Flatten, 1D binary search | O(log(mn))          |
| Row-wise Sorted             | Per-row binary search     | O(m log n)          |
| Row & Col Sorted            | Staircase (top-right)     | O(m+n)              |
| "Kth" Value Problems        | Binary search on answers  | O((m+n)log(max-min))|

**Key Takeaway:**  
*Mastering binary search on 2D arrays means knowing the matrix's structure and mapping the problem's requirements to the correct pattern or transformation. Always clarify assumptions, test edges, and write clean index conversions.*