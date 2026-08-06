# Quest10 — compact

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice manipulation and pointer usage** in Go.  
The task: write a function `Compact` that:
- Deletes zero‑value elements (`""`) from a slice of strings.
- Returns the number of non‑zero elements.
- Operates directly on the slice via a pointer.

## Instructions
- File to submit: `compact.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func Compact(ptr *[]string) int
```

## Implementation
`compact.go`:
```go
package piscine

func Compact(ptr *[]string) int {
    original := *ptr
    compacted := []string{}
    for _, v := range original {
        if v != "" {
            compacted = append(compacted, v)
        }
    }
    *ptr = compacted
    return len(compacted)
}
```

### Explanation
- Dereference the pointer to access the slice.
- Create a new slice `compacted` to store non‑empty strings.
- Iterate through the original slice:
  - Append only non‑empty values.
- Replace the original slice with the compacted one.
- Return the length of the compacted slice.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

const N = 6

func main() {
    a := make([]string, N)
    a[0] = "a"
    a[2] = "b"
    a[4] = "c"

    for _, v := range a {
        fmt.Println(v)
    }

    fmt.Println("Size after compacting:", piscine.Compact(&a))

    for _, v := range a {
        fmt.Println(v)
    }
}
```

Output:
```text
a

b

c

Size after compacting: 3
a
b
c
```

## Standard Library Equivalent
Go’s standard library does not provide a direct “compact” function for slices.  
This solution demonstrates how to manually filter and reassign slices using pointers.

## Skills Practiced
- Slice manipulation
- Pointer usage
- Filtering values
- Returning computed results

## Notes
- This exercise reinforces how to work with slices by reference.
- Compacting is a common operation in data cleaning and memory optimization.