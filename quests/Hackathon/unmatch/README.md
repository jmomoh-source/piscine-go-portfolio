# Quest10 — unmatch

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **pair detection in slices** in Go.  
The task: write a function `Unmatch` that returns the element of the slice that does not have a corresponding pair.  
- If all numbers have a pair, return `-1`.

## Instructions
- File to submit: `unmatch.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Unmatch(a []int) int
```

## Implementation
`unmatch.go`:
```go
package piscine

func Unmatch(a []int) int {
    counts := make(map[int]int)
    for _, v := range a {
        counts[v]++
    }
    for k, v := range counts {
        if v%2 != 0 {
            return k
        }
    }
    return -1
}
```

### Explanation
- Create a map to count occurrences of each integer in the slice.
- Loop through the slice and increment counts.
- Loop through the map:
  - If a number’s count is odd, return that number (it has no pair).
- If all counts are even, return `-1`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    a := []int{1, 2, 3, 1, 2, 3, 4}
    unmatch := piscine.Unmatch(a)
    fmt.Println(unmatch) // 4
}
```

Output:
```text
4
```

## Standard Library Equivalent
Go’s standard library does not provide a direct “unmatched element” function.  
This solution demonstrates how to use maps for frequency counting and simple logic to detect odd occurrences.

## Skills Practiced
- Slice iteration
- Map usage for counting
- Odd/even logic
- Returning sentinel values (`-1`)

## Notes
- This exercise reinforces the use of maps for counting and detecting anomalies.
- The approach is efficient: **O(n)** time complexity with **O(k)** space, where `k` is the number of distinct elements.