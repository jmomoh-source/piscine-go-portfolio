# Quest10 — podiumposition

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice of slices manipulation** in Go.  
The task: write a function `PodiumPosition` that reorders podium positions correctly:
- Input: slice of slices of type `string` representing positions.
- Output: reordered slice so that positions are in ascending order: 1st, 2nd, 3rd, 4th.

## Instructions
- File to submit: `podiumposition.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func PodiumPosition(podium [][]string) [][]string
```

## Implementation
`podiumposition.go`:
```go
package piscine

func PodiumPosition(podium [][]string) [][]string {
    n := len(podium)
    result := make([][]string, n)
    for i := 0; i < n; i++ {
        result[i] = podium[n-1-i]
    }
    return result
}
```

### Explanation
- Get the length of the podium slice.
- Create a new slice of the same length using `make`.
- Fill it by reversing the order of the input slice.
- Return the corrected podium positions.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    p4 := []string{"4th Place"}
    p3 := []string{"3rd Place"}
    p2 := []string{"2nd Place"}
    p1 := []string{"1st Place"}

    position := [][]string{p4, p3, p2, p1}
    fmt.Println(piscine.PodiumPosition(position))
}
```

Output:
```text
[[1st Place] [2nd Place] [3rd Place] [4th Place]]
```

## Standard Library Equivalent
Go’s standard library does not provide a direct reverse function for slices of slices.  
This solution demonstrates how to implement reversal manually with indexing.

## Skills Practiced
- Slice of slices manipulation
- Index arithmetic
- Using `make` to allocate slices
- Algorithmic thinking

## Notes
- This exercise reinforces how to reorder nested slices.
- The approach is efficient: **O(n)** time complexity and **O(n)** space complexity.