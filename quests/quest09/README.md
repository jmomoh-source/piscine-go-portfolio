
---

## 📂 Project Structure
```text
quests/quest09/
├── foreach/
├── map/
├── any/
├── countif/
├── issorted/
├── doop/
├── sortwordarr/
└── advancedsortwordarr/
```

---

## Quest09 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Functional%20Programming%20Concepts-green`

### Overview
Quest09 introduces **functional programming patterns** in Go.  
The exercises focus on higher‑order functions, slice manipulation, and sorting.  
You’ll implement utilities that resemble common functional paradigms: `foreach`, `map`, `any`, `countif`, and sorting algorithms.

---

### Exercises Progression

| Exercise                 | Focus Area | Skills Practiced |
|---------------------------|------------|------------------|
| **foreach**              | Iteration | Apply a function to each element |
| **map**                  | Transformation | Transform slice elements |
| **any**                  | Predicate check | Return true if any element matches |
| **countif**              | Counting | Count elements satisfying a condition |
| **issorted**             | Order check | Verify slice ordering |
| **doop**                 | Operations | Apply arithmetic operations |
| **sortwordarr**          | Sorting | Implement basic word sorting |
| **advancedsortwordarr**  | Advanced sorting | Custom comparator sorting |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **foreach/map/any/countif** → `for` loops, `range`, `slices` package (Go 1.21+)
- **issorted** → `sort.SliceIsSorted`
- **doop** → basic arithmetic operators
- **sortwordarr/advancedsortwordarr** → `sort.Strings`, `sort.Slice`

---

### Skills Progression
Quest09 builds progressively:
1. **Iteration basics** → `foreach`
2. **Transformation** → `map`
3. **Predicate checks** → `any`, `countif`
4. **Order verification** → `issorted`
5. **Operations** → `doop`
6. **Sorting algorithms** → `sortwordarr`, `advancedsortwordarr`

By the end, you’ve practiced:
- Higher‑order function design
- Predicate logic
- Slice transformation
- Sorting algorithms
- Custom comparator functions

---

### Key Takeaway
Quest09 transforms slice and function exercises into a foundation for **functional programming in Go**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.