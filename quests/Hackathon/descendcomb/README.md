# Quest10 — descendcomb

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **nested loops and formatted output** in Go.  
The task: write a function `DescendComb` that prints all possible combinations of two different two‑digit numbers in **descending order**, separated by a comma and a space.

Rules:
- Each number must be printed as two digits (e.g., `01`, `09`, `99`).
- Combinations are printed in descending order: highest pairs first.
- The last combination should not end with a trailing comma.

## Instructions
- File to submit: `descendcomb.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func DescendComb()
```

## Implementation
`descendcomb.go`:
```go
package piscine

import "fmt"

func DescendComb() {
    for a := 99; a >= 0; a-- {
        for b := a - 1; b >= 0; b-- {
            fmt.Printf("%02d %02d", a, b)
            if !(a == 1 && b == 0) {
                fmt.Print(", ")
            }
        }
    }
    fmt.Println()
}
```

### Explanation
- Use two nested loops:
  - Outer loop: `a` goes from 99 down to 0.
  - Inner loop: `b` goes from `a-1` down to 0.
- Print each pair as `aa bb` using `fmt.Printf("%02d %02d", ...)` to ensure two‑digit formatting.
- Add a comma and space after each pair except the last (`01 00`).
- Finally, print a newline.

## Usage
Example test program:
```go
package main

import "piscine"

func main() {
    piscine.DescendComb()
}
```

Output (truncated for brevity):
```text
99 98, 99 97, 99 96, ..., 03 01, 03 00, 02 01, 02 00, 01 00
```

## Standard Library Equivalent
This solution uses Go’s **`fmt`** package for formatted printing.  
It demonstrates how to generate and format combinations manually using loops and standard library functions.

## Skills Practiced
- Nested loops
- Formatted printing with `fmt`
- String formatting without external libraries
- Handling edge cases (avoiding trailing comma)

## Notes
- This exercise reinforces control flow and output formatting.
- The descending order requirement makes it slightly trickier than the usual ascending combination problems.