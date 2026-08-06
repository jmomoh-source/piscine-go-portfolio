# Quest04 — findnextprime

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise builds on prime number logic.  
The task: write a function `FindNextPrime` that returns the first prime number that is equal to or greater than the integer passed as a parameter.

Rules:
- Only positive numbers can be prime.
- The function must be optimized to avoid timeouts with large inputs.
- `1` is not considered prime.

## Instructions
- File to submit: `findnextprime.go` (inside the `piscine` package)
- Expected function signature:
```go
func FindNextPrime(nb int) int {
}
```

Output must match exactly:
```
bash
$ go run .
5
5
$
```

## Implementation
findnextprime.go:
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

func FindNextPrime(nb int) int {
    if nb <= 2 {
        return 2
    }
    for {
        if IsPrime(nb) {
            return nb
        }
        nb++
    }
}
```

## Explanation
- IsPrime helper function efficiently checks primality:

Rejects numbers ≤ 1.

Special case: 2 is prime.

Rejects even numbers > 2.

Checks divisibility only up to √nb.

- FindNextPrime:

If nb ≤ 2, return 2.

Otherwise, increment nb until a prime is found.

Returns the first prime ≥ nb.

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
    fmt.Println(piscine.FindNextPrime(5)) // 5
    fmt.Println(piscine.FindNextPrime(4)) // 5
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
5
5
```

## Skills Practiced
Prime number logic

Efficient iteration up to square root

Combining helper functions

Handling edge cases (≤ 2)

## Notes
This implementation is efficient enough for typical integer ranges.

Output must be exactly as specified, with no extra spaces or characters.

## Resources
Prime Numbers — MathWorld (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Go Loops — Tour of Go (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)

Efficient Prime Checking Algorithms (bing.com in Bing) (bing.com in Bing) (bing.com in Bing)