# Quest10 — abort

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **median calculation** in Go.  
The task: write a function `Abort` that returns the median of five integer arguments.  
- Median = the middle value when the numbers are sorted.
- With five numbers, the median is always the 3rd element after sorting.

## Instructions
- File to submit: `abort.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Abort(a, b, c, d, e int) int
```

## Implementation
`abort.go`:
```go
package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
    nums := []int{a, b, c, d, e}
    sort.Ints(nums)
    return nums[2] // median is the 3rd element (index 2)
}
```

### Explanation
- Place all five integers into a slice.
- Use `sort.Ints` from the Go standard library to sort them in ascending order.
- Return the element at index `2` (the third element), which is the median.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    middle := piscine.Abort(2, 3, 8, 5, 7)
    fmt.Println(middle) // 5
}
```

Output:
```text
5
```

## Standard Library Equivalent
Go’s standard library provides `sort.Ints` for sorting slices.  
This solution leverages it to quickly find the median without writing custom sorting logic.

## Skills Practiced
- Slice manipulation
- Sorting with the standard library
- Median calculation
- Understanding indexes in sorted data

## Notes
- With an odd number of elements, the median is always the middle element.
- This exercise reinforces how to use Go’s built‑in sorting utilities effectively.