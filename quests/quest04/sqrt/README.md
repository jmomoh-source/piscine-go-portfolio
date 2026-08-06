# Quest04 — sqrt

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces integer square root calculation in Go.  
The task: write a function `Sqrt` that returns the square root of the integer passed as a parameter **if the square root is a whole number**. Otherwise, return `0`.

Rules:
- Only perfect squares return their square root.
- Non‑perfect squares return `0`.
- Negative inputs return `0`.

## Instructions
- File to submit: `sqrt.go` (inside the `piscine` package)
- Expected function signature:
```go
func Sqrt(nb int) int {
}
```

Output must match exactly:
```
bash
$ go run .
2
0
$
```

## Implementation
sqrt.go:
```
go
package piscine

func Sqrt(nb int) int {
    if nb < 0 {
        return 0
    }
    for i := 1; i*i <= nb; i++ {
        if i*i == nb {
            return i
        }
    }
    return 0
}
```

## Explanation
If nb < 0 → return 0.

Loop through integers starting from 1:

If i*i == nb, return i (perfect square root found).

If i*i > nb, stop searching.

If no exact square root is found, return 0.

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
    fmt.Println(piscine.Sqrt(4))
    fmt.Println(piscine.Sqrt(3))
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
2
0
```

## Skills Practiced
Integer arithmetic in Go

Loop constructs (for)

Square root logic

Handling invalid inputs

## Notes
This implementation only works for perfect squares.

For non‑perfect squares, the function returns 0.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Loops — Tour of Go (bing.com in Bing) (bing.com in Bing)

Effective Go — Control Structures (bing.com in Bing) (bing.com in Bing)

Square Root Definition — MathWorld (bing.com in Bing) (bing.com in Bing)