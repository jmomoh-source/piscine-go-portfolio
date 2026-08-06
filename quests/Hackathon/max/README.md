# Quest10 — max

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice traversal and conditional logic** in Go.  
The task: write a function `Max` that returns the maximum value in a slice of integers.  
- If the slice is empty, return `0`.

## Instructions
- File to submit: `max.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Max(a []int) int
```

## Implementation
`max.go`:
```go
package piscine

func Max(a []int) int {
    if len(a) == 0 {
        return 0
    }
    maxVal := a[0]
    for _, v := range a {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```

### Explanation
- If the slice is empty, return `0`.
- Initialize `maxVal` with the first element.
- Iterate through the slice:
  - If a value is greater than `maxVal`, update `maxVal`.
- Return the maximum value found.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    a := []int{23, 123, 1, 11, 55, 93}
    max := piscine.Max(a)
    fmt.Println(max)
}
```

Output:
```text
123
```

## Standard Library Equivalent
Go’s standard library does not provide a direct `Max` function for slices.  
This solution demonstrates how to implement maximum value search manually.

## Skills Practiced
- Slice traversal
- Conditional comparison
- Handling empty slices
- Efficient iteration

## Notes
- This exercise reinforces how to compute aggregate values from slices.
- The approach is efficient: **O(n)** time complexity, scanning each element once.