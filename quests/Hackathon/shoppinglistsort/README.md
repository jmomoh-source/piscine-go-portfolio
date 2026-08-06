# Quest10 — shoppinglistsort

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **custom sorting logic** in Go.  
The task: write a function `ShoppingListSort` that sorts a slice of strings by their length in ascending order.

Rules:
- Input: slice of strings.
- Output: slice sorted by string length.
- Strings in the input slice are guaranteed to have different lengths.

## Instructions
- File to submit: `shoppinglistsort.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func ShoppingListSort(slice []string) []string
```

## Implementation
`shoppinglistsort.go`:
```go
package piscine

func ShoppingListSort(slice []string) []string {
    // Simple bubble sort by length
    n := len(slice)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if len(slice[j]) > len(slice[j+1]) {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
    return slice
}
```

### Explanation
- Use a bubble sort algorithm since only `len` is allowed.
- Compare lengths of adjacent strings.
- Swap if the current string is longer than the next.
- Continue until sorted in ascending order.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
    fmt.Println(piscine.ShoppingListSort(slice))
}
```

Output:
```text
[Tea Milk Honey Pepper Mushroom Pineapple]
```

## Standard Library Equivalent
Normally, Go’s `sort.Slice` would be used for custom sorting.  
Here, the restriction forces us to implement sorting manually using only `len`.

## Skills Practiced
- Slice manipulation
- Bubble sort algorithm
- Custom comparison logic
- Understanding constraints

## Notes
- This exercise reinforces algorithmic thinking under restrictions.
- Bubble sort is not the most efficient, but it’s simple and valid for this task.