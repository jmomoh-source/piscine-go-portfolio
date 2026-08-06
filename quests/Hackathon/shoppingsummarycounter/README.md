# Quest10 — shoppingsummarycounter

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **map usage and string processing** in Go.  
The task: write a function `ShoppingSummaryCounter` that returns a summary of how many times each item appears in a given receipt string.

Rules:
- Input: a string containing item names separated by spaces.
- Output: a `map[string]int` where keys are item names and values are counts.
- Case sensitivity is preserved (e.g., `"Burger"` and `"burger"` are different).

## Instructions
- File to submit: `shoppingsummarycounter.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func ShoppingSummaryCounter(str string) map[string]int
```

## Implementation
`shoppingsummarycounter.go`:
```go
package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
    summary := make(map[string]int)
    items := strings.Fields(str)
    for _, item := range items {
        summary[item]++
    }
    return summary
}
```

### Explanation
- Use `strings.Fields` to split the input string by whitespace into a slice of items.
- Create a map to store counts.
- Iterate through the slice and increment the count for each item.
- Return the map.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
    for index, element := range piscine.ShoppingSummaryCounter(summary) {
        fmt.Println(index, "=>", element)
    }
}
```

Output:
```text
Burger => 2
Water => 4
Carrot => 4
Coffee => 1
Chips => 1
```

## Standard Library Equivalent
Go’s standard library provides `strings.Fields` for splitting strings and maps for counting.  
This solution demonstrates how to combine them to build a frequency counter.

## Skills Practiced
- String splitting
- Map usage for counting
- Iteration over slices
- Frequency analysis

## Notes
- This exercise reinforces how to process text data and summarize results.
- The approach is efficient: **O(n)** time complexity with **O(k)** space, where `k` is the number of distinct items.