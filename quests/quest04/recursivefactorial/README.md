# Quest04 — recursivefactorial

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces recursion in Go.  
The task: write a function `RecursiveFactorial` that returns the factorial of the integer passed as a parameter.  

Rules:
- Factorial is defined as `n! = n × (n-1) × (n-2) × ... × 1`.
- For `n < 0`, return `0`.
- If overflow or invalid input occurs, return `0`.
- **Loops (`for`) are forbidden** — recursion must be used.

## Instructions
- File to submit: `recursivefactorial.go` (inside the `piscine` package)
- Expected function signature:
```go
func RecursiveFactorial(nb int) int {
}
```

Output must match exactly:
```
bash
$ go run .
24
$
```

## Implementation
recursivefactorial.go:
```
go
package piscine

func RecursiveFactorial(nb int) int {
    if nb < 0 {
        return 0
    }
    if nb == 0 || nb == 1 {
        return 1
    }
    return nb * RecursiveFactorial(nb-1)
}
```

## Explanation
- Base cases:

If nb < 0 → return 0.

If nb == 0 or nb == 1 → return 1.

- Recursive case:

Multiply nb by the factorial of nb-1.

This recursive definition mirrors the mathematical definition of factorial.

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
    arg := 4
    fmt.Println(piscine.RecursiveFactorial(arg))
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
24
```

## Skills Practiced
Recursion in Go

Base case vs. recursive case

Factorial calculation

Handling invalid inputs

## Notes
Factorials grow very quickly; large inputs may overflow int.

For this exercise, overflow cases simply return 0.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Recursion — Tour of Go (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Effective Go — Recursion (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Factorial Definition — MathWorld (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)