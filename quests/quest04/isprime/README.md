# Quest04 — isprime

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces prime number checking in Go.  
The task: write a function `IsPrime` that returns `true` if the integer passed as a parameter is a prime number, otherwise returns `false`.

Rules:
- Only positive numbers can be prime.
- `1` is **not** considered prime.
- The function must be optimized to avoid timeouts with large inputs.

## Instructions
- File to submit: `isprime.go` (inside the `piscine` package)
- Expected function signature:
```go
func IsPrime(nb int) bool {
}
```

Output must match exactly:
```
bash
$ go run .
true
false
$
```

## Implementation
isprime.go:
```
go
package piscine

func IsPrime(nb int) bool {
    if nb <= 1 {
        return false
    }
    if nb == 2 {
        return true
    }
    if nb%2 == 0 {
        return false
    }
    for i := 3; i*i <= nb; i += 2 {
        if nb%i == 0 {
            return false
        }
    }
    return true
}
```

## Explanation
Rejects numbers ≤ 1 → not prime.

- Special case: 2 is prime.

Rejects even numbers greater than 2.

Checks divisibility only up to √nb:

If any divisor is found, return false.

Otherwise, return true.

This optimization avoids unnecessary checks and ensures efficiency.

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
    fmt.Println(piscine.IsPrime(5)) // true
    fmt.Println(piscine.IsPrime(4)) // false
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
true
false
```

## Skills Practiced
Prime number logic

Efficient iteration up to square root

Conditional checks and optimization

Handling edge cases (≤ 1, even numbers)

## Notes
Prime checking is fundamental in number theory and cryptography.

This implementation is efficient enough for typical integer ranges.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Prime Numbers — MathWorld (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Go Loops — Tour of Go (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Efficient Prime Checking Algorithms (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)