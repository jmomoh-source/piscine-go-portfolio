# Quest09 — map

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **functional transformation of slices** in Go.  
The task: write a function `Map` that applies a function of type `func(int) bool` to each element of an integer slice and returns a slice of booleans.

Rules:
- The function must allocate a new slice of `bool` values.
- Each element of the input slice is transformed by the function `f`.
- The result slice must contain the return values of `f` in the same order.

## Instructions
- File to submit: `map.go`
- Expected function signature:
```go
func Map(f func(int) bool, a []int) []bool
```

## Implementation
`map.go`:
```go
package piscine

func Map(f func(int) bool, a []int) []bool {
    result := make([]bool, len(a))
    for i, v := range a {
        result[i] = f(v)
    }
    return result
}
```

### Explanation
- Allocate a slice of booleans with the same length as the input slice.
- Iterate through the input slice.
- Apply the function `f` to each element and store the result in the corresponding index.
- Return the result slice.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    result := piscine.Map(piscine.IsPrime, a)
    fmt.Println(result)
}
```

Output:
```text
[false true true false true false]
```

## Standard Library Equivalent
In Go’s standard library (Go 1.21+), you could achieve the same with `slices.Map`:
```go
import (
    "fmt"
    "slices"
)

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    result := slices.Map(a, func(v int) bool {
        return IsPrime(v)
    })
    fmt.Println(result)
}
```
⚠️ Note: `slices.Map` is concise and idiomatic.  
Your Piscine solution demonstrates manual slice transformation for deeper understanding.

## Skills Practiced
- Higher‑order functions
- Slice allocation with `make`
- Function application
- Boolean logic

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, prefer `slices.Map` (Go 1.21+) or explicit `for` loops for clarity.

## Resources
- Go `slices.Map` — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing)