# Quest09 — advancedsortwordarr

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **custom comparator sorting** in Go.  
The task: write a function `AdvancedSortWordArr` that sorts a slice of strings based on a comparator function passed as a parameter.

Rules:
- The comparator function `f func(a, b string) int` must return:
  - Positive if `a > b`
  - Zero if `a == b`
  - Negative if `a < b`
- Sorting must be done in place (mutating the slice).
- No return value is required.

## Instructions
- File to submit: `advancedsortwordarr.go`
- Expected function signature:
```go
func AdvancedSortWordArr(a []string, f func(a, b string) int)
```

## Implementation
`advancedsortwordarr.go`:
```go
package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if f(a[i], a[j]) > 0 {
                a[i], a[j] = a[j], a[i]
            }
        }
    }
}
```

### Explanation
- Use a nested loop (selection sort style).
- Compare strings using the provided comparator function `f`.
- Swap elements if out of order.
- The slice is sorted in place according to the comparator.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func Compare(a, b string) int {
    return int([]rune(a)[0] - []rune(b)[0])
}

func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    piscine.AdvancedSortWordArr(result, Compare)
    fmt.Println(result)
}
```

Output:
```text
[1 2 3 A B C a b c]
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `sort.Slice` and a custom comparator:
```go
import (
    "fmt"
    "sort"
)

func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })
    fmt.Println(result)
}
```
⚠️ Note: `sort.Slice` is concise and idiomatic.  
Your Piscine solution demonstrates manual comparator‑based sorting for deeper understanding.

## Skills Practiced
- Higher‑order functions
- Custom comparator design
- Slice manipulation
- Sorting algorithms
- In‑place mutation

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, prefer `sort.Slice` for efficiency and clarity.

## Resources
- Go `sort.Slice` — Official Docs (go.dev in Bing)  
- Effective Go — Sorting (go.dev in Bing)  