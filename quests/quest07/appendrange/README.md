# Quest07 — appendrange

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **slice creation without `make`** in Go.  
The task: write a function that takes `min` and `max` as parameters and returns a slice of integers containing all values between `min` (inclusive) and `max` (exclusive).

Rules:
- If `min >= max`, return `nil`.
- You must use `append` to build the slice.
- `make` is not allowed.

## Instructions
- File to submit: `appendrange.go`
- Expected function signature:
```go
func AppendRange(min, max int) []int
```

## Implementation
`appendrange.go`:
```go
package piscine

func AppendRange(min, max int) []int {
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

### Explanation
- Check if `min >= max`; if so, return `nil`.
- Initialize an empty slice (`var result []int`).
- Use a loop from `min` to `max - 1`.
- Append each integer to the slice.
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
    fmt.Println(piscine.AppendRange(5, 10))
    fmt.Println(piscine.AppendRange(10, 5))
}
```

Output:
```text
[5 6 7 8 9]
[]
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with `make`:
```go
func AppendRangeStd(min, max int) []int {
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

- ⚠️ Note: `make` is more efficient, but the Piscine solution demonstrates how to build slices manually with `append`.

## Skills Practiced
- Slice creation
- Iteration
- Using `append` to grow slices
- Conditional logic

## Notes
- This exercise demonstrates manual slice building.
- For production code, prefer `make` when the final length is known, for efficiency.

## Resources
- Go Slices — Official Docs (go.dev in Bing)  
- Effective Go — Slices (go.dev in Bing)  
