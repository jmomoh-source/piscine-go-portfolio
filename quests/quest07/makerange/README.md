# Quest07 — makerange

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **slice creation with `make`** in Go.  
The task: write a function that takes `min` and `max` as parameters and returns a slice of integers containing all values between `min` (inclusive) and `max` (exclusive).

Rules:
- If `min >= max`, return `nil`.
- You must use `make` to allocate the slice.
- `append` is not allowed.

## Instructions
- File to submit: `makerange.go`
- Expected function signature:
```go
func MakeRange(min, max int) []int
```

## Implementation
`makerange.go`:
```go
package piscine

func MakeRange(min, max int) []int {
    if min >= max {
        return nil
    }
    result := make([]int, max-min)
    for i := min; i < max; i++ {
        result[i-min] = i
    }
    return result
}
```

### Explanation
- Check if `min >= max`; if so, return `nil`.
- Allocate a slice of length `max - min` using `make`.
- Fill the slice with values from `min` up to `max - 1`.
- Return the slice.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.MakeRange(5, 10))
    fmt.Println(piscine.MakeRange(10, 5))
}
```

Output:
```text
[5 6 7 8 9]
[]
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `append`:
```go
func MakeRangeStd(min, max int) []int {
    if min >= max {
        return nil
    }
    var result []int
    for i := min; i < max; i++ {
        result = append(result, i)
    }
    return result
}
```

⚠️ Note: `append` is more flexible, but the Piscine solution demonstrates how to pre‑allocate slices with `make` for efficiency.

## Skills Practiced
- Slice allocation with `make`
- Iteration
- Index arithmetic
- Conditional logic

## Notes
- This exercise demonstrates efficient slice allocation.
- For production code, prefer `make` when the final length is known, for performance.

## Resources
- Go Slices — Official Docs (go.dev in Bing)  
- Effective Go — Slices (go.dev in Bing)  
