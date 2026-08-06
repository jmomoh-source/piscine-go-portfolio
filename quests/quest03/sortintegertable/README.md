# Quest03 — sortintegertable

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces sorting algorithms in Go.  
The task: write a function `SortIntegerTable` that reorders a slice of integers in ascending order.

## Instructions
- File to submit: `sortintegertable.go` (inside the `piscine` package)
- Expected function signature:
```go
func SortIntegerTable(table []int) {
}
```

Output must match exactly:
```
bash
$ go run .
[0 1 2 3 4 5]
$
```

## Implementation
sortintegertable.go:
```
go
package piscine

func SortIntegerTable(table []int) {
    n := len(table)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if table[j] > table[j+1] {
                table[j], table[j+1] = table[j+1], table[j]
            }
        }
    }
}
```

## Explanation
Uses a simple bubble sort algorithm:

Outer loop runs through the slice.

Inner loop compares adjacent elements.

If table[j] > table[j+1], swap them.

After completion, the slice is sorted in ascending order.

The function modifies the slice in place, so no return value is needed.

## Usage
Example test program:
```
go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    s := []int{5, 4, 3, 2, 1, 0}
    piscine.SortIntegerTable(s)
    fmt.Println(s)
}
```

Run it:
```
bash
go run .
```

Expected output:
```
text
[0 1 2 3 4 5]
```

## Skills Practiced
Implementing sorting algorithms in Go

Slice manipulation

Swapping values with multiple assignment

Understanding in‑place modifications

## Notes
Bubble sort is simple but not the most efficient for large slices.

For larger datasets, Go’s standard library provides sort.Ints, but here the goal is to practice manual sorting logic.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Slices — Tour of Go (go.dev in Bing) (bing.com in Bing)

Effective Go — Slices (go.dev in Bing) (bing.com in Bing)

Sorting Algorithms Overview (bing.com in Bing) (bing.com in Bing)