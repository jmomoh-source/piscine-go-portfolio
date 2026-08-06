# Quest10 — activebits

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **bitwise operations** in Go.  
The task: write a function `ActiveBits` that returns the number of active bits (bits with value `1`) in the binary representation of an integer.

Rules:
- Input: integer `n`.
- Output: number of `1` bits in its binary form.

## Instructions
- File to submit: `activebits.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
func ActiveBits(n int) int
```

## Implementation
`activebits.go`:
```go
package piscine

func ActiveBits(n int) int {
    count := 0
    for n != 0 {
        if n&1 == 1 {
            count++
        }
        n >>= 1
    }
    return count
}
```

### Explanation
- Initialize a counter.
- While `n` is not zero:
  - Use bitwise AND (`n & 1`) to check if the least significant bit is `1`.
  - Increment the counter if true.
  - Right‑shift `n` (`n >>= 1`) to check the next bit.
- Return the total count.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.ActiveBits(7))
}
```

Output:
```text
3
```

## Standard Library Equivalent
Go’s standard library does not provide a direct function for counting active bits.  
This solution demonstrates how to implement bit counting manually using bitwise operations.

## Skills Practiced
- Bitwise operations (`&`, `>>`)
- Looping until zero
- Binary representation understanding
- Efficient computation

## Notes
- This exercise reinforces how to manipulate and analyze binary data.
- The approach is efficient: **O(log n)** time complexity, proportional to the number of bits.