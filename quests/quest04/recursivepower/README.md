# Quest04 — recursivepower

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces recursion for exponentiation in Go.  
The task: write a function `RecursivePower` that returns the value of `nb` raised to the power of `power`.

Rules:
- Negative powers return `0`.
- Overflows do not need to be handled.
- **Loops (`for`) are forbidden** — recursion must be used.

## Instructions
- File to submit: `recursivepower.go` (inside the `piscine` package)
- Expected function signature:
```go
func RecursivePower(nb int, power int) int {
}
```

Output must match exactly:
```
bash
$ go run .
64
$
```

## Implementation
recursivepower.go:
```
go
package piscine

func RecursivePower(nb int, power int) int {
    if power < 0 {
        return 0
    }
    if power == 0 {
        return 1
    }
    return nb * RecursivePower(nb, power-1)
}
```

## Explanation
If power < 0 → return 0.

If power == 0 → return 1 (any number to the power of 0 is 1).

- Recursive case: multiply nb by the result of RecursivePower(nb, power-1) until the base case is reached.

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
    fmt.Println(piscine.RecursivePower(4, 3))
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
64
```

## Skills Practiced
Recursion in Go

Base case vs. recursive case

Exponentiation logic

Handling negative exponents

## Notes
Negative exponents are not supported in this simplified version → return 0.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Recursion — Tour of Go (bing.com in Bing) (bing.com in Bing)

Effective Go — Recursion (bing.com in Bing) (bing.com in Bing)

Exponentiation Basics — MathWorld (bing.com in Bing) (bing.com in Bing)