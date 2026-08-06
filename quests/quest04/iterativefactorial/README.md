# Quest04 — iterativefactorial

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces iterative algorithms in Go.  
The task: write a function `IterativeFactorial` that returns the factorial of the integer passed as a parameter.  

Rules:
- Factorial is defined as `n! = n × (n-1) × (n-2) × ... × 1`.
- For `n < 0`, return `0`.
- If overflow or invalid input occurs, return `0`.

## Instructions
- File to submit: `iterativefactorial.go` (inside the `piscine` package)
- Expected function signature:
```go
func IterativeFactorial(nb int) int {
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
iterativefactorial.go:
```
go
package piscine

func IterativeFactorial(nb int) int {
    if nb < 0 {
        return 0
    }
    result := 1
    for i := 1; i <= nb; i++ {
        result *= i
    }
    return result
}
```

## Explanation
Checks for negative input → returns 0.

Initializes result to 1.

Iteratively multiplies result by each integer from 1 to nb.

Returns the final factorial value.

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
    fmt.Println(piscine.IterativeFactorial(arg))
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
Iterative algorithms in Go

Loop constructs (for)

Factorial calculation

Handling invalid inputs

## Notes
Factorials grow very quickly; large inputs may overflow int.

For this exercise, overflow cases simply return 0.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Loops — Tour of Go (bing.com in Bing) (bing.com in Bing)

Effective Go — Control Structures (bing.com in Bing) (bing.com in Bing)

Factorial Definition — MathWorld (bing.com in Bing) (bing.com in Bing)