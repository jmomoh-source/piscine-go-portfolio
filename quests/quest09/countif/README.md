# Quest09 — countif

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **predicate counting on slices** in Go.  
The task: write a function `CountIf` that returns the number of elements in a string slice for which the function `f func(string) bool` returns `true`.

Rules:
- The function must iterate through the slice.
- Apply the predicate function `f` to each element.
- Increment a counter whenever `f` returns `true`.
- Return the final count.

## Instructions
- File to submit: `countif.go`
- Expected function signature:
```go
func CountIf(f func(string) bool, tab []string) int
```

## Implementation
`countif.go`:
```go
package piscine

func CountIf(f func(string) bool, tab []string) int {
    count := 0
    for _, v := range tab {
        if f(v) {
            count++
        }
    }
    return count
}
```

### Explanation
- Initialize a counter to `0`.
- Iterate through the slice of strings.
- Apply the predicate function `f` to each element.
- If `f(v)` returns `true`, increment the counter.
- Return the counter after the loop.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    tab1 := []string{"Hello", "how", "are", "you"}
    tab2 := []string{"This", "1", "is", "4", "you"}

    answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
    answer2 := piscine.CountIf(piscine.IsNumeric, tab2)

    fmt.Println(answer1)
    fmt.Println(answer2)
}
```

Output:
```text
0
2
```

## Standard Library Equivalent
In Go’s standard library, you would typically use a `for` loop directly.  
From Go 1.21+, you could use `slices.IndexFunc` or `slices.ContainsFunc` for related checks, but counting requires a loop:
```go
func CountIfStd(f func(string) bool, tab []string) int {
    count := 0
    for _, v := range tab {
        if f(v) {
            count++
        }
    }
    return count
}
```
⚠️ Note: Go does not have a built‑in `countIf` function.  
Your Piscine solution demonstrates manual predicate counting for deeper understanding.

## Skills Practiced
- Higher‑order functions
- Predicate logic
- Slice iteration
- Counting elements

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, explicit `for` loops are often clearer, but higher‑order functions can improve abstraction.

## Resources
- Go Functions — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing) 