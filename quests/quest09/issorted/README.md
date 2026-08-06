# Quest09 — issorted

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **custom comparator functions** in Go.  
The task: write a function `IsSorted` that returns `true` if a slice of integers is sorted according to a comparator function, otherwise returns `false`.

Rules:
- The comparator function `f func(a, b int) int` must return:
  - Positive if `a > b`
  - Zero if `a == b`
  - Negative if `a < b`
- You must use this comparator to check ordering.
- Write your own comparator function for testing.

## Instructions
- File to submit: `issorted.go`
- Expected function signature:
```go
func IsSorted(f func(a, b int) int, a []int) bool
```

## Implementation
`issorted.go`:
```go
package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
    if len(a) < 2 {
        return true
    }

    ascending := true
    descending := true

    for i := 0; i < len(a)-1; i++ {
        cmp := f(a[i], a[i+1])
        if cmp > 0 {
            ascending = false
        }
        if cmp < 0 {
            descending = false
        }
    }

    return ascending || descending
}
```

### Explanation
- If the slice has fewer than 2 elements, it is trivially sorted.
- Track two flags: `ascending` and `descending`.
- Iterate through the slice comparing adjacent elements with `f`.
  - If any pair violates ascending order, set `ascending = false`.
  - If any pair violates descending order, set `descending = false`.
- Return `true` if either ascending or descending order holds.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func f(a, b int) int {
    return a - b
}

func main() {
    a1 := []int{0, 1, 2, 3, 4, 5}
    a2 := []int{0, 2, 1, 3}

    result1 := piscine.IsSorted(f, a1)
    result2 := piscine.IsSorted(f, a2)

    fmt.Println(result1)
    fmt.Println(result2)
}
```

Output:
```text
true
false
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `sort.SliceIsSorted`:
```go
import (
    "fmt"
    "sort"
)

func main() {
    a := []int{0, 1, 2, 3, 4, 5}
    sorted := sort.SliceIsSorted(a, func(i, j int) bool {
        return a[i] < a[j]
    })
    fmt.Println(sorted) // true
}
```
⚠️ Note: `sort.SliceIsSorted` is concise and idiomatic.  
Your Piscine solution demonstrates manual comparator logic for deeper understanding.

## Skills Practiced
- Higher‑order functions
- Custom comparator design
- Slice iteration
- Sorting checks

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, prefer `sort.SliceIsSorted` for clarity and efficiency.

## Resources
- Go `sort.SliceIsSorted` — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing)  