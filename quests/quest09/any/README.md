# Quest09 — any

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **predicate checks on slices** in Go.  
The task: write a function `Any` that returns `true` if, when a string slice is passed through a function `f func(string) bool`, at least one element returns `true`.

Rules:
- The function must iterate through the slice.
- Apply the predicate function `f` to each element.
- Return `true` immediately if any element satisfies the condition.
- Return `false` if no elements satisfy the condition.

## Instructions
- File to submit: `any.go`
- Expected function signature:
```go
func Any(f func(string) bool, a []string) bool
```

## Implementation
`any.go`:
```go
package piscine

func Any(f func(string) bool, a []string) bool {
    for _, v := range a {
        if f(v) {
            return true
        }
    }
    return false
}
```

### Explanation
- Iterate through the slice of strings.
- Apply the predicate function `f` to each element.
- If any call returns `true`, return `true` immediately.
- If none return `true`, return `false`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    a1 := []string{"Hello", "how", "are", "you"}
    a2 := []string{"This", "is", "4", "you"}

    result1 := piscine.Any(piscine.IsNumeric, a1)
    result2 := piscine.Any(piscine.IsNumeric, a2)

    fmt.Println(result1)
    fmt.Println(result2)
}
```

Output:
```text
false
true
```

## Standard Library Equivalent
In Go’s standard library (Go 1.21+), you could achieve the same with `slices.ContainsFunc`:
```go
import (
    "fmt"
    "slices"
)

func main() {
    a := []string{"This", "is", "4", "you"}
    result := slices.ContainsFunc(a, func(s string) bool {
        return IsNumeric(s)
    })
    fmt.Println(result) // true
}
```
⚠️ Note: `slices.ContainsFunc` is concise and idiomatic.  
Your Piscine solution demonstrates manual predicate checks for deeper understanding.

## Skills Practiced
- Higher‑order functions
- Predicate logic
- Slice iteration
- Early return optimization

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, prefer `slices.ContainsFunc` (Go 1.21+) or explicit `for` loops for clarity.

## Resources
- Go `slices.ContainsFunc` — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing) 