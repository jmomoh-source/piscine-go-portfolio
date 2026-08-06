# Quest04 — fibonacci

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces recursion for sequence generation in Go.  
The task: write a function `Fibonacci` that returns the value at the given position `index` in the Fibonacci sequence.

Rules:
- The sequence starts: `0, 1, 1, 2, 3, 5, ...`
- Index `0` → `0`, index `1` → `1`.
- Negative indices return `-1`.
- **Loops (`for`) are forbidden** — recursion must be used.

## Instructions
- File to submit: `fibonacci.go` (inside the `piscine` package)
- Expected function signature:
```go
func Fibonacci(index int) int {
}
```

Output must match exactly:
```
bash
$ go run .
3
$
```

## Implementation
fibonacci.go:
```
go
package piscine

func Fibonacci(index int) int {
    if index < 0 {
        return -1
    }
    if index == 0 {
        return 0
    }
    if index == 1 {
        return 1
    }
    return Fibonacci(index-1) + Fibonacci(index-2)
}
```

## Explanation
- Base cases:

index < 0 → return -1.

index == 0 → return 0.

index == 1 → return 1.

- Recursive case:

Fibonacci(index) = Fibonacci(index-1) + Fibonacci(index-2).

This mirrors the mathematical definition of the Fibonacci sequence.

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
    arg1 := 4
    fmt.Println(piscine.Fibonacci(arg1))
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
3
```

## Skills Practiced
Recursion in Go

Base case vs. recursive case

Sequence generation

Handling invalid inputs

## Notes
Recursive Fibonacci is simple but inefficient for large indices due to repeated calculations.

For larger inputs, iterative or memoized approaches are more efficient, but recursion is required here.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Recursion — Tour of Go (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Effective Go — Recursion (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Fibonacci Sequence — MathWorld (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)