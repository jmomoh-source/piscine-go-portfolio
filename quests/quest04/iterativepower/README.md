# Quest04 — iterativepower

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces iterative exponentiation in Go.  
The task: write a function `IterativePower` that returns the value of `nb` raised to the power of `power`.

Rules:
- Negative powers return `0`.
- Overflows do not need to be handled.
- The function must use iteration (loops).

## Instructions
- File to submit: `iterativepower.go` (inside the `piscine` package)
- Expected function signature:
```go
func IterativePower(nb int, power int) int {
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
iterativepower.go:
```
go
package piscine

func IterativePower(nb int, power int) int {
    if power < 0 {
        return 0
    }
    result := 1
    for i := 0; i < power; i++ {
        result *= nb
    }
    return result
}
```

## Explanation
If power < 0 → return 0.

Initialize result to 1.

Multiply result by nb, repeated power times.

Return the final result.

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
    fmt.Println(piscine.IterativePower(4, 3))
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
Iterative algorithms in Go

Loop constructs (for)

Exponentiation logic

Handling negative exponents

## Notes
Negative exponents are not supported in this simplified version → return 0.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Go Loops — Tour of Go (bing.com in Bing) (bing.com in Bing)

Effective Go — Control Structures (bing.com in Bing) (bing.com in Bing)

Exponentiation Basics — MathWorld (bing.com in Bing) (bing.com in Bing)