# Quest10 — reversemenuindex

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice reversal** in Go.  
The task: write a function `ReverseMenuIndex` that takes a slice of strings and returns another slice with the elements in reverse order.

Rules:
- Input: slice of strings.
- Output: slice reversed.
- Restriction: `append()` is not allowed, but `make()` is allowed.

## Instructions
- File to submit: `reversemenuindex.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func ReverseMenuIndex(menu []string) []string
```

## Implementation
`reversemenuindex.go`:
```go
package piscine

func ReverseMenuIndex(menu []string) []string {
    n := len(menu)
    reversed := make([]string, n)
    for i := 0; i < n; i++ {
        reversed[i] = menu[n-1-i]
    }
    return reversed
}
```

### Explanation
- Get the length of the input slice.
- Create a new slice of the same length using `make`.
- Fill it by iterating through the input slice and assigning elements in reverse order.
- Return the reversed slice.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
```

Output:
```text
[starters drinks mains desserts]
```

## Standard Library Equivalent
Go’s standard library does not provide a direct reverse function for slices.  
This solution demonstrates how to implement reversal manually with indexing.

## Skills Practiced
- Slice manipulation
- Index arithmetic
- Using `make` to allocate slices
- Algorithmic thinking

## Notes
- This exercise reinforces how to work with slices without relying on `append`.
- The approach is efficient: **O(n)** time complexity and **O(n)** space complexity.