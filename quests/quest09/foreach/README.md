# Quest09 — foreach

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **higher‑order functions** in Go.  
The task: write a function `ForEach` that applies a given function to each element of an integer slice.

Rules:
- The function must accept another function as a parameter (`f func(int)`).
- The function must iterate over the slice and apply `f` to each element.
- No return value is required.

## Instructions
- File to submit: `foreach.go`
- Expected function signature:
```go
func ForEach(f func(int), a []int)
```

## Implementation
`foreach.go`:
```go
package piscine

func ForEach(f func(int), a []int) {
    for _, v := range a {
        f(v)
    }
}
```

### Explanation
- `ForEach` takes a function `f` and a slice `a`.
- It iterates over the slice using `range`.
- For each element, it calls `f(v)`.

## Usage
Example test program:
```go
package main

import "piscine"

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    piscine.ForEach(piscine.PrintNbr, a)
}
```

Output:
```text
123456
```

## Standard Library Equivalent
In Go’s standard library, you would typically use a `for` loop directly:
```go
func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    for _, v := range a {
        fmt.Print(v)
    }
}
```
⚠️ Note: Go does not have built‑in higher‑order functions like `map` or `forEach` in other languages.  
Your Piscine solution demonstrates how to implement them manually.

## Skills Practiced
- Higher‑order functions
- Function types
- Slice iteration
- Applying functions to data

## Notes
- This exercise demonstrates how Go supports functional programming concepts via first‑class functions.
- For production code, explicit `for` loops are often clearer, but higher‑order functions can improve abstraction.

## Resources
- Go Functions — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing)