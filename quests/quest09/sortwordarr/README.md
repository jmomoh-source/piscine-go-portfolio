# Quest09 — sortwordarr

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **sorting slices of strings** in Go.  
The task: write a function `SortWordArr` that sorts a slice of strings in ascending ASCII order.

Rules:
- Sorting must be done in place (mutating the slice).
- Order is determined by ASCII values, not by locale or case‑insensitive rules.
- No return value is required.

## Instructions
- File to submit: `sortwordarr.go`
- Expected function signature:
```go
func SortWordArr(a []string)
```

## Implementation
`sortwordarr.go`:
```go
package piscine

func SortWordArr(a []string) {
    for i := 0; i < len(a)-1; i++ {
        for j := i + 1; j < len(a); j++ {
            if a[i] > a[j] {
                a[i], a[j] = a[j], a[i]
            }
        }
    }
}
```

### Explanation
- Use a simple nested loop (bubble/selection sort style).
- Compare strings directly with `>` which uses ASCII ordering.
- Swap elements if out of order.
- The slice is sorted in place.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    piscine.SortWordArr(result)
    fmt.Println(result)
}
```

Output:
```text
[1 2 3 A B C a b c]
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `sort.Strings`:
```go
import (
    "fmt"
    "sort"
)

func main() {
    result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
    sort.Strings(result)
    fmt.Println(result)
}
```
⚠️ Note: `sort.Strings` is concise and idiomatic.  
Your Piscine solution demonstrates manual sorting for deeper understanding.

## Skills Practiced
- Slice manipulation
- String comparison
- Sorting algorithms
- In‑place mutation

## Notes
- This exercise demonstrates manual sorting logic.
- For production code, prefer `sort.Strings` or `sort.Slice` for efficiency and clarity.

## Resources
- Go `sort.Strings` — Official Docs (go.dev in Bing)  
- Effective Go — Sorting (go.dev in Bing)  