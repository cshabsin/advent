# Efficiency Analysis: Union-Find vs. Original Implementation

The original code was significantly **less efficient** than the standard Union-Find algorithm, primarily due to how it handled merging sets.

While the implementation was conceptually similar to the **"Quick-Find"** variation of the Disjoint Set algorithm (where finding the set ID is fast, but merging is slow), it had specific implementation details that made it much slower than necessary.

### 1. The "Re-indexing" Cost (O(N))

In the original code, circuits were stored in a slice `[][]int`. When two circuits were merged, one was deleted from the slice (`slices.Delete`). This shifted the indices of all subsequent circuits.
To keep the `inCircuit` map consistent, the code iterated over **every single point** in the map to decrement their circuit ID if it was greater than the deleted index:

```go
// This loop runs over all points (N) every time you merge!
for point, circuit := range c.inCircuit {
    if circuit > targetCircuit {
        c.inCircuit[point] = circuit - 1
    }
}
```

This makes a single union operation **O(N)** (where N is the total number of points), leading to a quadratic **O(N²)** complexity overall.

### 2. The "Relabeling" Cost (O(Size of Set))

Even without the slice shifting issue, the approach eagerly updated the membership of every node in the smaller set during a merge:

```go
// Updates every node in the target set
for _, p := range c.circuits[targetCircuit] {
    c.inCircuit[p] = sourceCircuit
}
```

This is a characteristic of "Quick-Find." While `Find` is O(1), `Union` is proportional to the size of the set being merged.

### Comparison with Union-Find

The **Union-Find** (Disjoint Set Union) data structure with **path compression** and **union by rank/size** is much faster:

1.  **Lazy Updates**: It doesn't update every node when sets merge. It only links the "root" of one set to the "root" of the other.
2.  **Complexity**: Both `Find` and `Union` operations are nearly **O(1)** (amortized).
3.  **No Re-indexing**: It doesn't need to shift or renumber sets globally.

### Summary

- **Original Code**: **O(N²)** (due to the global re-indexing loop).
- **Standard Union-Find**: Nearly **O(N)** (technically O(N \* α(N))).

For 1,000 points, O(N²) is roughly 1,000,000 operations, which is manageable. However, for larger inputs (e.g., N = 100,000), the original code would likely time out, whereas Union-Find would finish instantly.
