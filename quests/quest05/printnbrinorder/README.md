# Quest05 — printnbrinorder

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **digit extraction and sorting** in Go.  
The task: write a function `PrintNbrInOrder` that prints the digits of an integer in ascending order.

Rules:
- Input is a non‑negative integer (`int` type).
- Extract all digits, sort them, and print them in ascending order.
- If `n == 0`, print `0`.
- Conversion to `int64` is not allowed.
- Use `z01.PrintRune` for output.

## Instructions
- File to submit: `printnbrinorder.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintNbrInOrder(n int) {
}
```

Output must match exactly:
```
bash
$ go run . | cat -e
1230123$
$
```

## Implementation
printnbrinorder.go:
```
go
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
    if n == 0 {
        z01.PrintRune('0')
        return
    }

    digits := []int{}
    for n > 0 {
        digits = append(digits, n%10)
        n /= 10
    }

    // Simple insertion sort
    for i := 1; i < len(digits); i++ {
        key := digits[i]
        j := i - 1
        for j >= 0 && digits[j] > key {
            digits[j+1] = digits[j]
            j--
        }
        digits[j+1] = key
    }

    for _, d := range digits {
        z01.PrintRune(rune(d + '0'))
    }
}
```

## Explanation
Extract digits using modulo (n % 10) and division (n /= 10).

Store digits in a slice.

Sort the slice (here using insertion sort).

Print each digit as a rune ('0' + digit).

Special case: if n == 0, print '0'.

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintNbrInOrder(321)
    piscine.PrintNbrInOrder(0)
    piscine.PrintNbrInOrder(321)
}
```

Run it:
```
bash
go run . | cat -e
```

Expected output:
```
text
1230123$
```

## Standard Library Equivalent
In Go’s standard library, you could achieve similar behavior using:
```
go
import (
    "fmt"
    "sort"
    "strconv"
)

func PrintNbrInOrderStd(n int) {
    s := strconv.Itoa(n)
    digits := []rune(s)
    sort.Slice(digits, func(i, j int) bool { return digits[i] < digits[j] })
    fmt.Print(string(digits))
}
```

- ⚠️ Note: This uses strconv.Itoa and sort.Slice, which are part of the standard library.
Your manual implementation shows how to achieve the same result without relying on these helpers.

## Skills Practiced
Digit extraction from integers

Sorting algorithms

Rune printing with z01.PrintRune

Understanding Go’s sort and strconv equivalents

## Notes
Manual sorting demonstrates algorithmic thinking.

For production code, prefer sort.Ints or sort.Slice for efficiency and readability.

## Resources
Go sort Package — Official Docs (go.dev in Bing)

Go strconv.Itoa — Official Docs (go.dev in Bing)

Effective Go — Runes and Strings (go.dev in Bing)