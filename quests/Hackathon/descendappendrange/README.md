# Quest10 — descendappendrange

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice building with append** in Go.  
The task: write a function `DescendAppendRange` that:
- Takes two integers: `max` and `min`.
- Returns a slice of integers containing all values between `max` and `min`.
- `max` is included, `min` is excluded.
- If `max <= min`, return an empty slice.
- **Restriction:** `make()` is not allowed; use `append` instead.

## Instructions
- File to submit: `descendappendrange.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func DescendAppendRange(max, min int) []int
```

## Implementation
`descendappendrange.go`:
```go
package piscine

func DescendAppendRange(max, min int) []int {
    if max <= min {
        return []int{}
    }
    result := []int{}
    for i := max; i > min; i-- {
        result = append(result, i)
    }
    return result
}
```

### Explanation
- If `max` is less than or equal to `min`, return an empty slice.
- Start from `max` and decrement until just above `min`.
- Append each value to the result slice.
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
    fmt.Println(piscine.DescendAppendRange(10, 5))
    fmt.Println(piscine.DescendAppendRange(5, 10))
}
```

Output:
```text
[10 9 8 7 6]
[]
```

## Standard Library Equivalent
Go’s standard library does not provide a direct “range builder” function.  
This solution demonstrates how to construct ranges manually using loops and `append`.

## Skills Practiced
- Slice building
- Looping with decrement
- Conditional logic
- Efficient use of `append`

## Notes
- This exercise reinforces how to build slices dynamically without `make`.
- Descending order is achieved by decrementing the loop counter.